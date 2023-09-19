package gown

import (
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"
)

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
	app, err := l.loadApp()

	if err != nil {
		return nil, err
	}

	return &Project{
		app: app,
	}, nil
}

func (l *loader) loadApp() (*Application, error) {
	appFiles, subdirs, err := l.loadPackage("app")

	if err != nil {
		return nil, err
	}

	modules := []Module{}

	for _, dirname := range subdirs {
		files, _, err := l.loadPackage("app", dirname)

		if err != nil {
			return nil, err
		}

		module := NewModule(dirname, files)
		modules = append(modules, module)
	}

	return &Application{
		modules: modules,
		files:   appFiles,
	}, nil
}

func (l *loader) loadPackage(packagePath ...string) ([]File, []string, error) {
	pth := []string{l.projectPath}
	pth = append(pth, packagePath...)
	packageDir := path.Join(pth...)

	files := []File{}
	subdirs := []string{}

	entries, err := os.ReadDir(packageDir)

	if err != nil {
		return nil, nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subdirs = append(subdirs, entry.Name())
			continue
		}

		isGoFile := entry.Type().IsRegular() && strings.HasSuffix(entry.Name(), ".go")

		if isGoFile {
			file, err := l.loadFile(packageDir, entry.Name())

			if err != nil {
				return nil, nil, err
			}

			files = append(files, file)
		}
	}

	return files, subdirs, nil
}

func (l *loader) loadFile(packageDir string, fileName string) (File, error) {
	filePath := path.Join(packageDir, fileName)
	node, err := parser.ParseFile(l.fset, filePath, nil, parser.ParseComments)

	if err != nil {
		return File{}, err
	}

	return File{
		path: filePath,
		name: fileName,
		node: node,
	}, nil
}
