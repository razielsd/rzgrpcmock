package data

import (
	"fmt"
	"go/ast"
	alias "strings"
)

func FakeUseImport() {
	a := ast.ImportSpec{}
	fmt.Sprintf(alias.ToLower("asd"), a.Path) //nolint
}
