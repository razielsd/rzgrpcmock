package srcparser

import "go/ast"

type InterfaceSpec struct {
	Name       string
	ast        *ast.GenDecl
	MethodList []*MethodSpec
	ImportList map[string]*ImportSpec
}

func newInterfaceSpec(name string, target *ast.GenDecl) *InterfaceSpec {
	return &InterfaceSpec{
		Name:       name,
		ast:        target,
		ImportList: make(map[string]*ImportSpec),
	}
}
