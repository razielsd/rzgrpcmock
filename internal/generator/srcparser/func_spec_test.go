package srcparser

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuncSpec_Parse(t *testing.T) {
}

func TestFuncSpec_extractParams(t *testing.T) {
}

func Test_newFuncSpec(t *testing.T) {
	node := &ast.FuncType{}
	actual := newFuncSpec("pkgname", "funcName", node)
	require.NotNil(t, actual)
}
