package srcparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type FileParser struct {
	PackageName  string
	filename     string
	astFile      *ast.File
	importMap    map[string]*ImportSpec
	interfaceMap map[string]*InterfaceSpec
}

func NewFileParser(filename string) *FileParser {
	return &FileParser{
		filename:     filename,
		importMap:    make(map[string]*ImportSpec),
		interfaceMap: make(map[string]*InterfaceSpec),
	}
}

func (f *FileParser) Parse() error {
	if err := f.loadFile(); err != nil {
		return err
	}
	f.extractPackage()
	f.extractImports()
	if err := f.extractInterfaceList(); err != nil {
		return err
	}
	for _, v := range f.interfaceMap {
		v.syncImport(f.importMap)
	}
	return nil
}

func (f *FileParser) GetInterfaceList() []*InterfaceSpec {
	result := make([]*InterfaceSpec, 0, len(f.interfaceMap))
	for _, v := range f.interfaceMap {
		result = append(result, v)
	}
	return result
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
		f.importMap[importSpec.GetName()] = importSpec
	}
}

func (f *FileParser) extractInterfaceList() error {
	for _, v := range f.astFile.Decls {
		genD, ok := v.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genD.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			_, ok = currType.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}
			interSpec := newInterfaceSpec(f.PackageName, currType.Name.Name, genD)
			err := interSpec.Parse()
			if err != nil {
				return fmt.Errorf("unable parse interface %s: %w", currType.Name.Name, err)
			}
			f.interfaceMap[currType.Name.Name] = interSpec
		}
	}
	return nil
}
