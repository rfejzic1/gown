package gown

import (
	"fmt"
	"go/ast"
	"io"
)

type Project struct {
	app *Application
}

type Application struct {
	modules []Module
	files   []File
}

type Module struct {
	name  string
	files []File
}

func NewModule(name string, files []File) Module {
	return Module{
		// TODO: Sanitize name
		name:  name,
		files: files,
	}
}

type File struct {
	path string
	name string
	node ast.Node
}

func PrintProjectStructure(p *Project, w io.Writer) {
	fmt.Fprintf(w, "app:\n")
	for _, module := range p.app.modules {
		fmt.Fprintf(w, "  %s:\n", module.name)

		for _, file := range module.files {
			fmt.Fprintf(w, "    %s\n", file.name)
		}
	}

	for _, file := range p.app.files {
		fmt.Fprintf(w, "  %s\n", file.name)
	}
}
