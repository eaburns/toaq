package parse

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/eaburns/peggy/peg"
)

//go:generate peggy -o toaq.go toaq.peg

// A Text is a successfully parsed Toaq text.
type Text struct {
	*_Parser
}

// ParseTree returns the full parse tree of the text.
func (txt *Text) ParseTree() *peg.Node {
	_, t := _full_textNode(txt._Parser, 0)
	// Reset so a subsequent call returns a different tree.
	// This allows the caller to modify the returned tree
	// without affecting the tree made by subsequent calls.
	txt._Parser.node = make(map[_key]*peg.Node)
	return t
}

func (txt *Text) String() string {
	_, str := _full_textAction(txt._Parser, 0)
	return *str
}

// An Error is a parse error.
type Error interface {
	error
	FailTree() *peg.Fail
}

type parseError struct {
	fail *peg.Fail
	err  peg.Error
}

func (err *parseError) Error() string       { return err.err.Error() }
func (err *parseError) FailTree() *peg.Fail { return err.fail }

// String parses a string of text as a full Toaq discourse
// and returns either a successfully parsed *Text
// or a parse Error.
func String(text string) (*Text, Error) {
	p := _NewParser(text)
	pos, perr := _full_textAccepts(p, 0)
	if pos < 0 {
		_, failTree := _full_textFail(p, 0, perr)
		return nil, &parseError{
			fail: failTree,
			err:  peg.SimpleError(text, failTree),
		}
	}
	return &Text{_Parser: p}, nil
}

// Input parses the contents of an io.Reader
// as a full Toaq discourse
// and returns either a successfully parsed *Text,
// an error representing failure reading the input,
// or an Error representing a parse error.
func Input(in io.Reader) (*Text, error) {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	return String(string(data))
}

// File parses the contents of a file as a full Toaq discourse
// and returns either a successfully parsed *Text,
// an error representing failure on the input file,
// or an Error representing a parse error.
func File(path string) (*Text, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	text, err := Input(bufio.NewReader(f))
	if parseErr, ok := err.(*parseError); ok {
		parseErr.err.FilePath = path
		return nil, parseErr
	}
	return text, err
}
