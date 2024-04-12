package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	// writeSource()
	modifySource()
}

// Use this to create AST and write to source file
func writeSource() {
	fset := token.NewFileSet()

	f := &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: ast.NewIdent("main"),
				Type: &ast.FuncType{
					Params:  &ast.FieldList{},
					Results: &ast.FieldList{},
				},
				Body: &ast.BlockStmt{},
			},
		},
	}

	if err := format.Node(os.Stdout, fset, f); err != nil {
		log.Fatal(err)
	}
}

// Use this to modify existing source file
func modifySource() {
	fset := token.NewFileSet()

	const hello = `package main
import "fmt"

func main() {
    fmt.Println("Hello, world")
}`

	f, err := parser.ParseFile(fset, "hello.go", hello, 0)

	myFunctionDecl := &ast.FuncDecl{
		Name: ast.NewIdent("myFunction"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("myParam")},
						Type:  ast.NewIdent("string"),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("bool"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.IDENT,
							Value: "true",
						},
					},
				},
			},
		},
	}

	f.Decls = append(f.Decls, myFunctionDecl)

	if err != nil {
		log.Fatal(err)
	}

	if err := format.Node(os.Stdout, fset, f); err != nil {
		log.Fatal(err)
	}
}
