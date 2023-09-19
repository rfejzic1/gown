package gown

import "go/ast"

type Project struct {
	app *Application
}

type Application struct {
	modules []Module
	files   []File
}

func NewApplication(modules []Module, files []File) (*Application, error) {
	return &Application{
		modules: modules,
		files:   files,
	}, nil
}

type Module struct {
	name  string
	files []File
}

func NewModule(name string, files []File) Module {
	return Module{
		name:  name,
		files: files,
	}
}

type File struct {
	path string
	name string
	node ast.Node
}
