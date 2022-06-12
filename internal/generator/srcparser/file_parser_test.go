package srcparser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testFile = "./data/example.go"

func TestFileParser_Parse(t *testing.T) {

}

func TestFileParser_extractImports(t *testing.T) {
	parser := loadTestParser(t)
	parser.extractImports()
}

func TestFileParser_extractPackage(t *testing.T) {
	parser := loadTestParser(t)
	parser.extractPackage()
	require.Equal(t, "data", parser.PackageName)
}

func TestFileParser_loadFile(t *testing.T) {
	parser := NewFileParser(testFile)
	require.NotNil(t, parser)
	err := parser.loadFile()
	require.NoError(t, err)
	require.NotNil(t, parser.astFile)
}

func TestNewFileParser(t *testing.T) {
	parser := NewFileParser(testFile)
	require.NotNil(t, parser)
}

func loadTestParser(t *testing.T) *FileParser {
	parser := NewFileParser(testFile)
	require.NotNil(t, parser)
	err := parser.loadFile()
	require.NoError(t, err)
	require.NotNil(t, parser.astFile)
	return parser
}