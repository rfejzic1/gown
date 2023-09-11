package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "./main.go", nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		funcDecl, ok := n.(*ast.FuncDecl)
		if ok && funcDecl.Name.String() != "main" {
			// line := fset.Position(funcDecl.Pos()).Line
			// fmt.Printf("%v declared on line %d\n", funcDecl.Name, line)

			funcDecl.Body = &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("fmt"),
								Sel: ast.NewIdent("Println"),
							},
							Args: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: fmt.Sprintf("\"Hello, there\""),
								},
							},
							Ellipsis: token.NoPos,
						},
					},
				},
			}
		}
		return true
	})

	format.Node(os.Stdout, fset, file)
}

func test() {
	fmt.Println("Hello, original")
}

func test2() {}
