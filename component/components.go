package component

import "go/ast"

type Project struct {
	App *Application
}

type Application struct {
	Modules []Module
	Files   []File
}

func NewApplication(modules []Module, files []File) (*Application, error) {
	return &Application{
		Modules: modules,
		Files:   files,
	}, nil
}

type Module struct {
	Name string
	File []File
}

func NewModule(name string, files []File) Module {
	return Module{
		Name: name,
		File: files,
	}
}

type File struct {
	Path string
	Name string
	Node ast.Node
}
