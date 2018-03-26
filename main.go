// This is a temporary program to test the parser.
package main

import (
	"fmt"
	"os"

	"github.com/eaburns/peggy/peg"
	"github.com/eaburns/toaq/parse"
)

func main() {
	var err error
	var text *parse.Text
	if len(os.Args) > 1 {
		text, err = parse.File(os.Args[1])
	} else {
		text, err = parse.Input(os.Stdin)
	}
	if err != nil {
		parseErr, ok := err.(parse.Error)
		if !ok {
			fmt.Println("failed to read input:", err)
			os.Exit(1)
		}
		failTree := parseErr.FailTree()
		peg.DedupFails(failTree)
		peg.PrettyWrite(os.Stdout, failTree)
		os.Stdout.WriteString("\n")
		fmt.Println(parseErr)
		os.Exit(1)
	}
	peg.PrettyWrite(os.Stdout, text.ParseTree())
	os.Stdout.WriteString("\n")
	os.Stdout.WriteString(text.String())
}
