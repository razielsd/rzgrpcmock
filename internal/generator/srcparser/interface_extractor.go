package srcparser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

type InterfaceExtractor struct {
	PackageName   string
	Name          string
	importSpec    map[string]*ast.ImportSpec
	InterfaceList []*InterfaceField
}

type InterfaceField struct {
	Name       string
	Target     *ast.GenDecl
	MethodList []*InterfaceMethod
	ImportList map[string]*ast.ImportSpec
}

type InterfaceMethod struct {
	Name   string
	Args   []*VarField
	Result []*VarField
}

type VarField struct {
	Name      string
	Package   string
	IsPointer bool
}

func newInterfaceField(name string, target *ast.GenDecl) *InterfaceField {
	return &InterfaceField{
		Name:       name,
		Target:     target,
		ImportList: make(map[string]*ast.ImportSpec),
	}
}

func NewVarField(pkg, name string, isPointer bool) *VarField {
	return &VarField{
		Name:      name,
		Package:   pkg,
		IsPointer: isPointer,
	}
}

func (m InterfaceMethod) String() string {
	args := make([]string, 0, len(m.Args))
	for _, a := range m.Args {
		args = append(args, a.String())
	}
	result := make([]string, 0, len(m.Result))
	for _, r := range m.Result {
		result = append(result, r.String())
	}
	s := fmt.Sprintf("%s(%s) (%s)", m.Name, strings.Join(args, ", "), strings.Join(result, ", "))
	return s
}

func (v VarField) String() string {
	s := ""
	if v.IsPointer {
		s = "*"
	}
	if v.Package != "" {
		s += v.Package + "."
	}
	s += v.Name
	return s
}

func (v VarField) FullName() string {
	s := ""
	if v.Package != "" {
		s += v.Package + "."
	}
	s += v.Name
	return s
}

func NewInterfaceExtractor() *InterfaceExtractor {
	return &InterfaceExtractor{}
}

func (i *InterfaceExtractor) Parse(filename string) error {
	i.importSpec = make(map[string]*ast.ImportSpec)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	i.extractPackage(node)
	i.extractImports(node)
	for _, f := range node.Decls {
		var field *InterfaceField
		genD, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genD.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			_, ok = currType.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}
			field = newInterfaceField(currType.Name.Name, genD)
		}
		if field != nil {
			err := i.extractInterface(field)
			if err != nil {
				return err
			}
			i.InterfaceList = append(i.InterfaceList, field)
		}
	}

	return nil
}

func (i *InterfaceExtractor) extractPackage(node ast.Node) {
	x, ok := node.(*ast.File)
	if ok {
		i.PackageName = x.Name.Name
	}
}

func (i *InterfaceExtractor) extractImports(node ast.Node) {
	x, ok := node.(*ast.File)
	if ok {
		for _, v := range x.Imports {
			if v.Name != nil {
				i.importSpec[v.Name.String()] = v
			}
		}
	}
}

func (i *InterfaceExtractor) extractInterface(field *InterfaceField) error {
	for _, spec := range field.Target.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			return errors.New("unable get type spec")
		}
		tp, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			return errors.New("unable get interface type")
		}

		methodList, importList, err := i.parseInterfaceMethods(tp)
		if err != nil {
			return err
		}
		field.ImportList = mergeMapImportSpec(field.ImportList, importList)
		field.MethodList = append(field.MethodList, methodList...)
	}
	return nil
}

func (i *InterfaceExtractor) parseInterfaceMethods(s *ast.InterfaceType) ([]*InterfaceMethod, map[string]*ast.ImportSpec, error) {
	var methodList []*InterfaceMethod
	importList := make(map[string]*ast.ImportSpec)
	for _, f := range s.Methods.List {
		method := &InterfaceMethod{
			Name: f.Names[0].Name,
		}
		if !ast.IsExported(method.Name) {
			continue
		}
		fn, ok := f.Type.(*ast.FuncType)
		if !ok {
			log.Fatalln("unable cast func to ast.FuncType")
		}

		args, paramImport, err := i.extractParams(fn.Params.List)
		if err != nil {
			return nil, nil, err
		}
		importList = mergeMapImportSpec(importList, paramImport)
		method.Args = args
		result, paramImport, err := i.extractParams(fn.Results.List)
		if err != nil {
			return nil, nil, err
		}
		importList = mergeMapImportSpec(importList, paramImport)
		method.Result = result
		methodList = append(methodList, method)
	}
	return methodList, importList, nil
}

func (i *InterfaceExtractor) extractParams(fn []*ast.Field) ([]*VarField, map[string]*ast.ImportSpec, error) {
	var result []*VarField
	importList := make(map[string]*ast.ImportSpec)
	for _, param := range fn {
		var arg *VarField
		switch sel := param.Type.(type) {
		case *ast.SelectorExpr:
			pkg := ""
			if ident, ok := sel.X.(*ast.Ident); ok {
				pkg = ident.Name
			}
			arg = NewVarField(pkg, sel.Sel.Name, false)
		case *ast.StarExpr:
			switch ptr := sel.X.(type) {
			case *ast.Ident:
				arg = NewVarField(i.PackageName, ptr.Name, true)
			case *ast.SelectorExpr:
				ident, ok := ptr.X.(*ast.Ident)
				if !ok {
					log.Fatalln("unable cast to ast.Ident package argument from *ast.SelectorExpr")
				}
				spec, ok := i.importSpec[ident.Name]
				if ok {
					importList[ident.Name] = spec
				}
				arg = NewVarField(ident.Name, ptr.Sel.Name, true)
			default:
				log.Fatalln("unable cast param to ast")
			}
		case *ast.Ident:
			arg = NewVarField("", sel.Name, false)
		case *ast.Ellipsis:
			// multiple args, unsupported now
		default:
			return nil, nil, fmt.Errorf("unable cast arg to unexpected type")
		}
		if arg != nil {
			result = append(result, arg)
		}
	}
	return result, importList, nil
}

func mergeMapImportSpec(m1, m2 map[string]*ast.ImportSpec) map[string]*ast.ImportSpec {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
