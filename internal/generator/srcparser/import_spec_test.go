package srcparser

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImportSpec_GetName(t *testing.T) {
	t.Run("simple import name", func(t *testing.T) {
		astSpec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Value: "\"fmt\"",
			},
		}
		spec := newImportSpec(astSpec)
		require.Equal(t, "fmt", spec.GetName())
	})

	t.Run("composite import name", func(t *testing.T) {
		astSpec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Value: "\"go/ast\"",
			},
		}
		spec := newImportSpec(astSpec)
		require.Equal(t, "ast", spec.GetName())
	})

	t.Run("alias import name", func(t *testing.T) {
		astSpec := &ast.ImportSpec{
			Name: &ast.Ident{
				Name: "alias",
			},
			Path: &ast.BasicLit{
				Value: "\"fmt\"",
			},
		}
		spec := newImportSpec(astSpec)
		require.Equal(t, "alias", spec.GetName())
	})
}

func Test_newImportSpec(t *testing.T) {
	astSpec := &ast.ImportSpec{}
	spec := newImportSpec(astSpec)
	require.NotNil(t, spec)
}
