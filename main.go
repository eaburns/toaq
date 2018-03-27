// This is a temporary program to test the parser.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/eaburns/peggy/peg"
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
	p := parser.New(string(data))
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
	tree, err := p.Text()
	if err != nil {
		panic(err) // can't fail since ParseTree succeeded.
	}
	fmt.Println(tree)
}
