package srcparser

import (
	"go/ast"
	"strings"
)

type ImportSpec struct {
	spec *ast.ImportSpec
}

func newImportSpec(spec *ast.ImportSpec) *ImportSpec {
	return &ImportSpec{spec: spec}
}

func (i *ImportSpec) GetName() string {
	if i.spec.Name != nil {
		return i.spec.Name.String()
	}
	return i.extractImportName(i.spec.Path.Value)
}

func (i *ImportSpec) extractImportName(path string) string {
	path = strings.Trim(path, `"`)
	parts := strings.Split(path, `/`)
	return parts[len(parts) - 1]
}