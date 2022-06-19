package srcparser

import (
	"fmt"
	"go/ast"
	"log"
)

type FuncSpec struct {
	Name        string
	PackageName string
	ast         *ast.FuncType
	ImportList  []string
	ArgList     []*VarSpec
	ResultList  []*VarSpec
}

func newFuncSpec(pkgName, name string, ast *ast.FuncType) *FuncSpec {
	return &FuncSpec{
		PackageName: pkgName,
		Name:        name,
		ast:         ast,
		ImportList:  make([]string, 0),
	}
}

func (f *FuncSpec) Parse() error {
	args, argImport, err := f.extractParams(f.ast.Params.List)
	if err != nil {
		return err
	}
	f.ArgList = args
	result, resImport, err := f.extractParams(f.ast.Results.List)
	if err != nil {
		return err
	}
	f.ImportList = filterUniqStr(append(argImport, resImport...))
	f.ResultList = result

	return nil
}

func (f *FuncSpec) extractParams(fn []*ast.Field) ([]*VarSpec, []string, error) {
	var result []*VarSpec
	importList := make([]string, 0)
	for _, param := range fn {
		var arg *VarSpec
		switch sel := param.Type.(type) {
		case *ast.SelectorExpr:
			pkg := ""
			if ident, ok := sel.X.(*ast.Ident); ok {
				pkg = ident.Name
			}
			arg = NewVarSpec(pkg, sel.Sel.Name, false)
		case *ast.StarExpr:
			switch ptr := sel.X.(type) {
			case *ast.Ident:
				arg = NewVarSpec(f.PackageName, ptr.Name, true)
			case *ast.SelectorExpr:
				ident, ok := ptr.X.(*ast.Ident)
				if !ok {
					log.Fatalln("unable cast to ast.Ident package argument from *ast.SelectorExpr")
				}
				importList = append(importList, ident.Name)
				arg = NewVarSpec(ident.Name, ptr.Sel.Name, true)
			default:
				log.Fatalln("unable cast param to ast")
			}
		case *ast.Ident:
			arg = NewVarSpec("", sel.Name, false)
		case *ast.Ellipsis:
			// skip return nil, nil, fmt.Errorf("multiple args, unsupported now")
		default:
			return nil, nil, fmt.Errorf("unable cast arg to unexpected type")
		}
		if arg != nil {
			result = append(result, arg)
		}
	}
	return result, importList, nil
}
