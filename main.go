package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main(){
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset,"testdata/a.go",nil,0)
	if err != nil{
		panic(err)
	}

	for _, d := range f.Decls {
		switch d := d.(type){
		case *ast.FuncDecl:
			fmt.Println(d.Name)
		}
	}
}
