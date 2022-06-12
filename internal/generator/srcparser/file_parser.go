package srcparser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type FileParser struct {
	PackageName   string
	filename string
	astFile *ast.File
	importSpec    map[string]*ImportSpec
}

func NewFileParser(filename string) *FileParser {
	return &FileParser{
		filename: filename,
		importSpec: make(map[string]*ImportSpec),
	}
}

func (f *FileParser) Parse() error {
	if err := f.loadFile(); err != nil {
		return err
	}
	return nil
}

func (f *FileParser) loadFile() error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, f.filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	f.astFile = node
	return nil
}

func (f *FileParser) extractPackage() {
	f.PackageName = f.astFile.Name.Name
}

func (f *FileParser) extractImports() {
	for _, v := range f.astFile.Imports {
		importSpec := newImportSpec(v)
		f.importSpec[importSpec.GetName()] = importSpec
	}
}
