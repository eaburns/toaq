package ast

import "github.com/eaburns/peggy/peg"

//go:generate peggy -o toaq.go toaq.peggy
//go:generate peggy -pretty -o toaq.peg toaq.peggy

// Error is a parse error.
type Error interface {
	error
	Tree() *peg.Fail
	// ByteOffset returns the byte-offset of the deepest error in the parse.
	ByteOffset() int
}

type parseError struct {
	fail *peg.Fail
	err  peg.Error
}

func (err *parseError) Error() string { return err.err.Error() }

// Tree returns the tree of all failed backtracks that lead to a deepest-error.
func (err *parseError) Tree() *peg.Fail { return err.fail }
func (err *parseError) ByteOffset() int { return byteOffset(err.fail) }

func byteOffset(fail *peg.Fail) int {
	if len(fail.Kids) == 0 {
		return fail.Pos
	}
	max := -1
	for _, k := range fail.Kids {
		if o := byteOffset(k); o > max {
			max = o
		}
	}
	return max

}

// A Parser parses a given Toaq text.
type Parser struct {
	// FilePath is a file path used for error reporting.
	FilePath string
	*_Parser
}

// NewParser returns a new parser for a Toaq text.
func NewParser(text string) *Parser {
	return &Parser{_Parser: _NewParser(text)}
}

// Tree returns the full parse tree of the text.
func (p *Parser) Tree() (*peg.Node, Error) {
	if err := p.parse(_full_textAccepts, 0); err != nil {
		return nil, err
	}
	_, t := _full_textNode(p._Parser, 0)
	// Reset so a subsequent call returns a different tree.
	// This allows the caller to modify the returned tree
	// without affecting the tree made by subsequent calls.
	p._Parser.node = make(map[_key]*peg.Node)
	return t, nil
}

// Word returns the word at the given byte-offset into the text
// with any trailing whitespace attached and the number of bytes parsed.
func (p *Parser) Word(offs int) (*Word, int, Error) {
	if err := p.parse(_wordAccepts, offs); err != nil {
		return nil, 0, err
	}
	pos, t := _wordAction(p._Parser, offs)
	// Reset so a subsequent call returns a different tree.
	// This allows the caller to modify the returned tree
	// without affecting the tree made by subsequent calls.
	p._Parser.node = make(map[_key]*peg.Node)
	return *t, pos - offs, nil
}

// Text returns the AST of the text.
func (p *Parser) Text() (*Text, Error) {
	if err := p.parse(_full_textAccepts, 0); err != nil {
		return nil, err
	}
	_, txt := _full_textAction(p._Parser, 0)
	// Reset so subsequent calls return a different tree.
	// This allows the caller to modify the returned tree
	// without affecting the tree made by subsequent calls.
	p._Parser.act = make(map[_key]interface{})
	return txt, nil
}

func (p *Parser) parse(rule func(*_Parser, int) (int, int), start int) Error {
	pos, perr := rule(p._Parser, start)
	if pos >= 0 {
		return nil
	}
	_, failTree := _full_textFail(p._Parser, 0, perr)
	return &parseError{
		fail: failTree,
		err:  peg.SimpleError(p._Parser.text, failTree),
	}
}
