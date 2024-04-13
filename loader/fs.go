package loader

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"gown/component"
)

type fsLoader struct {
	projectPath string
	fset        *token.FileSet
}

func NewFsLoader(projectPath string) fsLoader {
	return fsLoader{
		projectPath: filepath.Clean(projectPath),
		fset:        token.NewFileSet(),
	}
}

func (l *fsLoader) Load() (*component.Project, error) {
	app, err := l.loadApp()

	if err != nil {
		return nil, err
	}

	return &component.Project{
		App: app,
	}, nil
}

func (l *fsLoader) loadApp() (*component.Application, error) {
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

func (l *fsLoader) loadPackage(packagePath ...string) ([]component.SourceFile, []string, error) {
	pth := []string{l.projectPath}
	pth = append(pth, packagePath...)
	packageDir := filepath.Join(pth...)

	files := []component.SourceFile{}
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

func (l *fsLoader) loadFile(packageDir string, fileName string) (component.SourceFile, error) {
	filePath := filepath.Join(packageDir, fileName)
	file, err := parser.ParseFile(l.fset, filePath, nil, parser.ParseComments)

	if err != nil {
		return component.SourceFile{}, err
	}

	return component.SourceFile{
		Path: filePath,
		File: file,
	}, nil
}
