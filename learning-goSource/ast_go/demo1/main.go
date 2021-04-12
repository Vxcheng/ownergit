package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return
	}
	ast.Print(fset, f)
}

// src is the input for which we want to print the AST.
const src = `
package main
func main() {
	println("Hello, World!")
}
`
