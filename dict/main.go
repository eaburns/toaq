// +build ignore

// Main loads the words from https://uakci.pl/toadua/dict and
// writes a go source file of all the words in a variable called "Words".
// This is intended to be used with go generate.
package main

import (
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
	"time"

	"github.com/eaburns/toaq/dict"
)

const url = "https://uakci.pl/toadua/dict"

var out = flag.String("o", "", "The output file.")

func main() {
	flag.Parse()

	ws, err := dict.Load(url)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	f, err := os.Create(*out)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	src := fmt.Sprintf("package dict\n"+
		"// Words contains the dictionary of words.\n"+
		"// Last updated: %v.\n"+
		"var Words = %#v\n", time.Now(), ws)
	src = strings.Replace(src, "dict.", "", -1)
	if err := gofmt(f, src); err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}

func gofmt(w io.Writer, s string) error {
	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, "", s, parser.ParseComments)
	if err != nil {
		return err
	}
	if err := format.Node(w, fset, root); err != nil {
		return err
	}
	return nil
}
