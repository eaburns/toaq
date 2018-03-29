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
	"github.com/eaburns/toaq/parser"
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
	p := parser.New(input)
	parseTree, err := p.Tree()
	if err != nil {
		failTree := err.(parser.Error).Tree()
		peg.DedupFails(failTree)
		peg.PrettyWrite(os.Stdout, failTree)
		os.Stdout.WriteString("\n")
		fmt.Println(err)
		os.Exit(1)
	}
	peg.PrettyWrite(os.Stdout, parseTree)
	os.Stdout.WriteString("\n")
	text, err := p.Text()
	if err != nil {
		panic(err) // can't fail since ParseTree succeeded.
	}
	pretty.Print(text)
	fmt.Println("")
	fmt.Println(toString(text))
	for i := 0; i < len(input); {
		fmt.Print(i, ": ")
		w, d, err := p.Word(i)
		if err != nil {
			failTree := err.(parser.Error).Tree()
			peg.DedupFails(failTree)
			peg.PrettyWrite(os.Stdout, failTree)
			os.Stdout.WriteString("\n")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(toString(w))
		i += d
	}
}

func toString(node ast.Node) string {
	var s strings.Builder
	ast.Visit(node, func(n ast.Node) bool {
		if w, ok := n.(*ast.Word); ok {
			s.WriteString(w.T)
		}
		return true
	})
	return s.String()
}
