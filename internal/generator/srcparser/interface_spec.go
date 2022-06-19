package srcparser

import (
	"errors"
	"fmt"
	"go/ast"
)

type InterfaceSpec struct {
	PackageName string
	Name        string
	ast         *ast.GenDecl
	FuncList    []*FuncSpec

	ImportList map[string]*ImportSpec
}

func newInterfaceSpec(pkgName, name string, target *ast.GenDecl) *InterfaceSpec {
	return &InterfaceSpec{
		PackageName: pkgName,
		Name:        name,
		ast:         target,
		ImportList:  make(map[string]*ImportSpec),
	}
}

func (i *InterfaceSpec) Parse() error {
	for _, spec := range i.ast.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			return errors.New("unable get type spec")
		}
		tp, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			return errors.New("unable get interface type")
		}
		err := i.parseInterfaceMethods(tp)
		if err != nil {
			return err
		}

		// field.ImportList = mergeMapImportSpec(field.ImportList, importList)
	}
	return nil
}

func (i *InterfaceSpec) parseInterfaceMethods(s *ast.InterfaceType) error {
	for _, f := range s.Methods.List {
		if !ast.IsExported(f.Names[0].Name) {
			continue
		}
		fn, ok := f.Type.(*ast.FuncType)
		if !ok {
			return errors.New("unable cast func to ast.FuncType")
		}
		method := newFuncSpec(i.PackageName, f.Names[0].Name, fn)
		i.FuncList = append(i.FuncList, method)
		if err := method.Parse(); err != nil {
			return fmt.Errorf("parse func: %w", err)
		}
	}

	return nil
}

func (i *InterfaceSpec) syncImport(importList map[string]*ImportSpec) {
	nameList := make([]string, 0)
	for _, v := range i.FuncList {
		nameList = append(nameList, v.ImportList...)
	}
	nameList = filterUniqStr(nameList)
	for _, name := range nameList {
		if _, ok := importList[name]; ok {
			i.ImportList[name] = importList[name]
		}
	}
}
