package component

import "go/ast"

type Package struct {
	Projct  *Project
	Name    string
	Path    string
	Sources []SourceFile
}

func (p *Package) Add(files ...SourceFile) {
	if p.Sources == nil {
		p.Sources = []SourceFile{}
	}

	p.Sources = append(p.Sources, files...)
}

type SourceFile struct {
	Package *Package
	Name    string
	Path    string
	Node    ast.Node
}
