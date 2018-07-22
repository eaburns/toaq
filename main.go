// This is a temporary program to test the parser.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/eaburns/peggy/peg"
	"github.com/eaburns/pretty"
	"github.com/eaburns/toaq/ast"
)

func main() {
	var path string
	var in io.Reader = os.Stdin
	if len(os.Args) > 1 {
		path = os.Args[1]
		f, err := os.Open(path)
		if err != nil {
			fmt.Printf("failed to open %s: %v", path, err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	}
	data, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Printf("failed to read input: %v", err)
		os.Exit(1)
	}
	input := string(data)
	p := ast.NewParser(input)
	parseTree, err := p.Tree()
	if err != nil {
		failTree := err.(ast.Error).Tree()
		peg.DedupFails(failTree)
		peg.PrettyWrite(os.Stdout, failTree)
		os.Stdout.WriteString("\n")
		fmt.Println(path + err.Error())
		os.Exit(1)
	}
	text, err := p.Text()
	if err != nil {
		panic(err) // can't fail since ParseTree succeeded.
	}
	fmt.Println(peg.Pretty(parseTree))
	fmt.Println(pretty.String(text))
	fmt.Println(toString(text))
}

func toString(node ast.Node) string {
	var s strings.Builder
	ast.Visit(node, ast.FuncVisitor(func(n ast.Node) {
		if w, ok := n.(*ast.Word); ok {
			s.WriteString(w.T)
		}
	}))
	return s.String()
}
