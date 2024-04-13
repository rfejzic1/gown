package component

import (
	"go/ast"
)

type Package struct {
	Projct *Project
	Name   string
	Path   string
	Source *SourceFile
}

type SourceFile struct {
	Path string
	File *ast.File
}
