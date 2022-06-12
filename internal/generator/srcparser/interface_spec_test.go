package srcparser

import (
	"github.com/stretchr/testify/require"
	ast2 "go/ast"
	"testing"
)

func Test_newInterfaceSpec(t *testing.T) {
	ast := &ast2.GenDecl{}
	spec := newInterfaceSpec("test", ast)
	require.NotNil(t, spec)
}
