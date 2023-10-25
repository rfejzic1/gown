package loader

import (
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"

	"gown/component"
)

type directoryLoader struct {
	projectPath string
	fset        *token.FileSet
}

func NewDirectoryLoader(projectPath string) directoryLoader {
	return directoryLoader{
		projectPath: path.Clean(projectPath),
		fset:        token.NewFileSet(),
	}
}

func (l *directoryLoader) Load() (*component.Project, error) {
	app, err := l.loadApp()

	if err != nil {
		return nil, err
	}

	return &component.Project{
		App: app,
	}, nil
}

func (l *directoryLoader) loadApp() (*component.Application, error) {
	appFiles, subdirs, err := l.loadPackage("app")

	if err != nil {
		return nil, err
	}

	modules := []component.Module{}

	for _, dirname := range subdirs {
		files, _, err := l.loadPackage("app", dirname)

		if err != nil {
			return nil, err
		}

		module := component.NewModule(dirname, files)
		modules = append(modules, module)
	}

	return &component.Application{
		Modules: modules,
		Files:   appFiles,
	}, nil
}

func (l *directoryLoader) loadPackage(packagePath ...string) ([]component.File, []string, error) {
	pth := []string{l.projectPath}
	pth = append(pth, packagePath...)
	packageDir := path.Join(pth...)

	files := []component.File{}
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

func (l *directoryLoader) loadFile(packageDir string, fileName string) (component.File, error) {
	filePath := path.Join(packageDir, fileName)
	node, err := parser.ParseFile(l.fset, filePath, nil, parser.ParseComments)

	if err != nil {
		return component.File{}, err
	}

	return component.File{
		Path: filePath,
		Name: fileName,
		Node: node,
	}, nil
}
