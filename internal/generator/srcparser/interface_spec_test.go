package srcparser

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newInterfaceSpec(t *testing.T) {
	node := &ast.GenDecl{}
	spec := newInterfaceSpec("pkgname", "test", node)
	require.NotNil(t, spec)
}
