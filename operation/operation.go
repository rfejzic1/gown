package operation

import (
	"go/ast"
	"go/format"
	"go/token"
	"gown/component"
	"os"
	"path/filepath"
)

const filePermission = 0744
const dirPermission = 0755

// TODO: Create a gown.toml (or other config) file to store information about the project
//		 Create builder functions and query functions to search and manipulate the AST
//		 Create helper functons to create arbitrary source files (such as air.toml or .gitingore or .Makerile or README.md)
//		 Create go.mod file at the root of the project

// Creates a directory with the given path relative to the project.
func CreateDirectory(p *component.Project, dirPath ...string) (string, error) {
	relPath := filepath.Join(dirPath...)
	fullPath := filepath.Join(p.Path, relPath)
	err := os.Mkdir(fullPath, dirPermission)
	return relPath, err
}

// Creates a libray package (non-executable) with a single go file inside.
func CreatePackage(p *component.Project, packagePath ...string) (*component.Package, error) {
	relPath, err := CreateDirectory(p, packagePath...)

	if err != nil {
		return nil, err
	}

	packageName := filepath.Base(relPath)

	pkg := &component.Package{
		Projct: p,
		Name:   packageName,
		Path:   relPath,
	}

	node := &ast.File{
		Name: ast.NewIdent(packageName),
	}

	sourceFile, err := WriteSourceFile(p, node, pkg.Path, pkg.Name)

	if err != nil {
		return nil, err
	}

	pkg.Source = sourceFile

	return pkg, nil
}

// Creates a package "main" that is executable.
func CreatePackageMain(p *component.Project, packagePath ...string) (*component.Package, error) {
	pkg, err := CreatePackage(p, packagePath...)

	if err != nil {
		return nil, err
	}

	mainDecl := &ast.FuncDecl{
		Name: ast.NewIdent("main"),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{},
			Results: &ast.FieldList{},
		},
		Body: &ast.BlockStmt{},
	}

	source := pkg.Source.File

	source.Name = ast.NewIdent("main")

	if source.Decls == nil {
		source.Decls = []ast.Decl{}
	}

	source.Decls = append(source.Decls, mainDecl)

	sourceFile, err := WriteSourceFile(p, pkg.Source.File, pkg.Path, pkg.Name)

	if err != nil {
		return nil, err
	}

	pkg.Source = sourceFile

	return pkg, nil
}

// Deletes a package and all it's subpackages. It doesn't refactor the rest of the source to remove references.
func DeletePackage(p *component.Project, packagePath ...string) error {
	relPath := filepath.Join(packagePath...)
	fullPath := filepath.Join(p.Path, relPath)
	return os.RemoveAll(fullPath)
}

// Creates or updates a go source file.
func WriteSourceFile(p *component.Project, source *ast.File, filePath ...string) (*component.SourceFile, error) {
	relPath := filepath.Join(filePath...) + ".go"
	fullPath := filepath.Join(p.Path, relPath)
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, filePermission)

	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	err = format.Node(file, fset, source)

	if err != nil {
		return nil, err
	}

	return &component.SourceFile{
		Path: relPath,
		File: source,
	}, nil
}

func WriteFile(p *component.Project, content []byte, filePath ...string) error {
	relPath := filepath.Join(filePath...)
	fullPath := filepath.Join(p.Path, relPath)

	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, filePermission)

	if err != nil {
		return err
	}

	_, err = file.Write(content)
	return err
}
