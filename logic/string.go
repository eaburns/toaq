package logic

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

// String returns a string representation of the logic expression.
func String(n Node) string {
	var s strings.Builder
	if err := print(&printer{w: &s}, n); err != nil {
		panic("impossible strings.Builder error")
	}
	return s.String()
}

// PrettyString returns a string representation of the logic expression.
// However, it uses some codepoints that are pretty, but less commonly available.
// Specifically,
// 	For definite quantification, it uses turned greek small letter iota, instead of ι.
// 	For variables it uses mathematical italic capital letter [A-Z], instead of [A-Z].
func PrettyString(n Node) string {
	var s strings.Builder
	if err := print(&printer{w: &s, pretty: true}, n); err != nil {
		panic("impossible strings.Builder error")
	}
	return s.String()
}

func (n *Binding) print(pr *printer) error {
	if n.Restriction == nil {
		return pr.write(n.Quantifier, variable(n.N), " ", n.Scope)
	}
	return pr.write("[", n.Quantifier, variable(n.N), " : ", n.Restriction, "] ", n.Scope)
}

func (n *Connection) print(pr *printer) error {
	lo, lc := parens(n, n.Left)
	ro, rc := parens(n, n.Right)
	return pr.write(lo, n.Left, lc, " ", n.Connector, " ", ro, n.Right, rc)
}

func parens(n *Connection, arg Statement) (string, string) {
	switch arg := arg.(type) {
	case *Predication:
		return "", ""
	case *Connection:
		if arg.Connector == n.Connector {
			return "", ""
		}
	}
	return "(", ")"
}

func (n *Predication) print(pr *printer) error {
	if p, ok := n.Predicate.(*Pred); ok && len(n.Arguments) == 2 && isBinOp(p) {
		return pr.write(n.Arguments[0], p.Name, n.Arguments[1])
	}
	return pr.write(n.Predicate, "(", n.Arguments, ")")
}

func isBinOp(n *Pred) bool {
	return n.Name == Equal || n.Name == Among
}

func (n *Variable) print(pr *printer) error {
	return pr.write(variable(n.N))
}

func (n *Plurality) print(pr *printer) error {
	return pr.write("⌊", n.Arguments, "⌉")
}

func (n *Quote) print(pr *printer) error {
	if n.FullText {
		return pr.write("«", n.Text, "»")
	}
	return pr.write("‹", n.Text, "›")
}

func (n *Pred) print(pr *printer) error {
	return pr.write(n.Name)
}

func (n *Abstraction) print(pr *printer) error {
	return pr.write("{", n.Statement, "}")
}

func (n *Serial) print(pr *printer) error {
	return pr.write(n.Left, "∘", n.Right)
}

func (n *Swap) print(pr *printer) error {
	return pr.write("(mu ", n.Predicate, ")")
}

// A printer prints logic expressions.
type printer struct {
	pretty   bool
	varNames []string
	w        io.Writer
}

func print(pr *printer, n Node) error {
	if n == nil {
		return nil
	}
	runes := make(map[string][]int)
	seen := make(map[int]bool)
	max := -1
	n.Visit(func(n Node) {
		v, ok := n.(*Variable)
		if !ok || seen[v.N] {
			return
		}
		seen[v.N] = true
		if v.N > max {
			max = v.N
		}
		r := varRune(v)
		runes[r] = append(runes[r], v.N)
	})
	if max >= len(seen) {
		// The variable numbers aren't in sequence.
		// This is probably a partial expression;
		// maybe you are debugging?
		// Let's just give up on pretty names.
		pr.varNames = nil
		return n.print(pr)
	}
	pr.varNames = make([]string, len(seen))
	for r, vs := range runes {
		switch {
		case r == "X" && len(vs) == 3:
			pr.varNames[vs[2]] = "Z"
			fallthrough
		case r == "X" && len(vs) == 2:
			pr.varNames[vs[1]] = "Y"
			pr.varNames[vs[0]] = "X"
		case len(vs) == 1:
			pr.varNames[vs[0]] = r
		default:
			for i, v := range vs {
				pr.varNames[v] = r + strconv.Itoa(i)
			}
		}
	}
	if err := n.print(pr); err != nil {
		return err
	}
	// TODO: For now we write a '.' at the end
	// so unit tests expect to end in a period.
	// However, once sentences are added to the logic tree,
	// each sentence will output a punctuation for its illocution:
	// 	da|ε .
	// 	ba !
	// 	ka <dunno>
	// 	moq ?
	_, err := io.WriteString(pr.w, ".")
	return err
}

func varRune(v *Variable) string {
	if v.Def == nil {
		return "U"
	}
	pred := leftPredication(v.Def.Restriction)
	if pred == nil {
		return "X"
	}
	p := leftPred(pred.Predicate)
	if p == nil || isBinOp(p) {
		return "X"
	}
	_, w := utf8.DecodeRuneInString(p.Name)
	return strings.ToUpper(p.Name[:w])
}

// If the left-most leaf statement is a *Predication, it is returned.
func leftPredication(st Statement) *Predication {
	switch st := st.(type) {
	case *Predication:
		return st
	case *Connection:
		return leftPredication(st.Left)
	default:
		return nil
	}
}

func leftPred(p Predicate) *Pred {
	switch p := p.(type) {
	case *Pred:
		return p
	case *Serial:
		l := leftPred(p.Left)
		if l != nil && l.Name == "shi" {
			return leftPred(p.Right)
		}
		return l
	default:
		return nil
	}
}

type variable int

func (pr *printer) write(vals ...interface{}) error {
	var err error
	for _, v := range vals {
		switch v := v.(type) {
		case nil:
			continue
		case string:
			_, err = io.WriteString(pr.w, v)
		case Quantifier:
			if pr.pretty && v == "ι" {
				v = "\u2129" // turned greek small letter iota
			}
			_, err = io.WriteString(pr.w, string(v))
		case Connector:
			_, err = io.WriteString(pr.w, string(v))
		case []Argument:
			err = writeArguments(pr, v)
		case variable:
			var n string
			if pr.varNames == nil || int(v) >= len(pr.varNames) {
				n = "X" + strconv.Itoa(int(v))
			} else {
				n = pr.varNames[v]
			}
			if pr.pretty {
				n = mathItalicCapitalLetters(n)
			}
			_, err = io.WriteString(pr.w, n)
		case Node:
			err = v.print(pr)
		default:
			panic(fmt.Sprintf("impossible type: %T", v))
		}
		if err != nil {
			break
		}
	}
	return err
}

func writeArguments(pr *printer, args []Argument) error {
	for i, arg := range args {
		comma := ", "
		if i == len(args)-1 {
			comma = ""
		}
		var err error
		if arg == nil {
			err = pr.write("·", comma)
		} else {
			err = pr.write(arg, comma)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func mathItalicCapitalLetters(str string) string {
	const A = 0x1D434
	var s strings.Builder
	for _, r := range str {
		if 'A' <= r && r <= 'Z' {
			r = r - 'A' + A
		}
		s.WriteRune(r)
	}
	return s.String()
}
