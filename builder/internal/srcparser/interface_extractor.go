package srcparser

import (
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
	InterfaceList []*InterfaceField
}

type InterfaceField struct {
	Name       string
	Target     *ast.GenDecl
	MethodList []*InterfaceMethod
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
		s = s + v.Package + "."
	}
	s = s + v.Name
	return s
}

func (v VarField) FullName() string {
	s := ""
	if v.Package != "" {
		s = s + v.Package + "."
	}
	s = s + v.Name
	return s
}

func NewInterfaceExtractor() *InterfaceExtractor {
	return &InterfaceExtractor{}
}

func (i *InterfaceExtractor) Parse(filename string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	if err := i.extractPackage(node); err != nil {
		return fmt.Errorf("unable parse package name: %w", err)
	}

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
			field = &InterfaceField{Name: currType.Name.Name, Target: genD}
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

func (i *InterfaceExtractor) extractPackage(node ast.Node) error {
	switch x := node.(type) {
	case *ast.File:
		i.PackageName = x.Name.Name
		fmt.Println(i.PackageName)
	}
	return nil
}

func (i *InterfaceExtractor) extractInterface(field *InterfaceField) error {
	for _, spec := range field.Target.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			fmt.Println("Error 1")
		}
		tp, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			fmt.Println("Error 1")
		}

		methodList, err := i.parseInterfaceMethods(tp)
		if err != nil {
			return err
		}
		field.MethodList = append(field.MethodList, methodList...)
	}
	return nil
}

func (i *InterfaceExtractor) parseInterfaceMethods(s *ast.InterfaceType) ([]*InterfaceMethod, error) {
	var methodList []*InterfaceMethod
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

		args, err := i.extractParams(fn.Params.List)
		if err != nil {
			return nil, err
		}
		method.Args = args
		result, err := i.extractParams(fn.Results.List)
		if err != nil {
			return nil, err
		}
		method.Result = result
		methodList = append(methodList, method)
	}
	return methodList, nil
}

func (i *InterfaceExtractor) extractParams(fn []*ast.Field) ([]*VarField, error) {
	var result []*VarField
	for _, param := range fn {
		var arg *VarField
		switch sel := param.Type.(type) {
		case *ast.SelectorExpr:
			pkg := ""
			if ident, ok := sel.X.(*ast.Ident); ok == true {
				pkg = ident.Name
			}
			arg = NewVarField(pkg, sel.Sel.Name, false)
		case *ast.StarExpr:
			ident, ok := sel.X.(*ast.Ident)
			if !ok {
				log.Fatalln("unable cast to ast.Ident package argument")
			}
			arg = NewVarField(i.PackageName, ident.Name, true)
		case *ast.Ident:
			arg = NewVarField("", sel.Name, false)
		case *ast.Ellipsis:
			// multiple args, unsupported now
		default:
			return nil, fmt.Errorf("unable cast arg to unexpected type")
		}
		if arg != nil {
			result = append(result, arg)
		}
	}
	return result, nil
}
