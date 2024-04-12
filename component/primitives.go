package component

import (
	"go/ast"
	"os"
)

type Package struct {
	Projct *Project
	Name   string
	Path   string
}

type SourceFile struct {
	Path string
	Node ast.Node
	File *os.File
}
