package gown

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"
)

type file struct {
	path string
	node ast.Node
}

type loader struct {
	projectPath string
	fset        *token.FileSet
}

func NewLoader(projectPath string) loader {
	return loader{
		projectPath: path.Clean(projectPath),
		fset:        token.NewFileSet(),
	}
}

func (l *loader) LoadProject() (*Project, error) {
	// TODO: Load all the packages in the project

	return &Project{}, nil
}

func (l *loader) loadPackage(packagePath string) ([]file, error) {
	packageDir := path.Join(l.projectPath, packagePath)

	files := []file{}

	entries, err := os.ReadDir(packageDir)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		isGoFile := entry.Type().IsRegular() && strings.HasSuffix(entry.Name(), ".go")

		if isGoFile {
			filePath := path.Join(packageDir, entry.Name())
			file, err := l.loadFile(filePath)

			if err != nil {
				return nil, err
			}

			files = append(files, file)
		}
	}

	return files, nil
}

func (l *loader) loadFile(filePath string) (file, error) {
	node, err := parser.ParseFile(l.fset, filePath, nil, parser.ParseComments)

	if err != nil {
		return file{}, err
	}

	return file{
		path: filePath,
		node: node,
	}, nil
}
