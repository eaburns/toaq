package ast

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

const (
	openBraces  = "({[<"
	closeBraces = ")}]>"
)

// BracesString returns a string representation of the AST,
// showing nesting with (), {}, [], and <>.
// Nodes with only a single child are not shown,
// nor is the original spacing between words.
func BracesString(n Node) string {
	var s strings.Builder
	if err := braces(&s, 0, n); err != nil {
		// strings.Builder should never return an error.
		panic(err.Error())
	}
	return s.String()
}

func braces(w io.Writer, cur int, n Node) error {
	switch n := n.(type) {
	case nil:
		return nil

	case *Text:
		return braceKids(w, cur, n.Leading, &n.Discourse)

	case *Discourse:
		kids := make([]interface{}, len(*n))
		for i, k := range *n {
			kids[i] = k
		}
		return braceKids(w, cur, kids...)

	case *StatementSentence:
		return braceKids(w, cur, n.JE, n.Statement, n.DA)

	case *CoPSentence:
		return braces(w, cur, (*CoP)(n))

	case *Prenex:
		return braceKids(w, cur, n.Terms, &n.BI)

	case *PrenexStatement:
		return braceKids(w, cur, &n.Prenex, n.Statement)

	case *Predication:
		return braceKids(w, cur, n.Predicate, n.Terms, n.NA)

	case *CoPStatement:
		return braces(w, cur, (*CoP)(n))

	case *PrefixedPredicate:
		return braceKids(w, cur, &n.MU, n.Predicate)

	case *SerialPredicate:
		return braceKids(w, cur, n.Left, n.Right)

	case *WordPredicate:
		return braces(w, cur, (*Word)(n))

	case *MIPredicate:
		return braceKids(w, cur, &n.MI, n.Phrase, n.GA)

	case *POPredicate:
		return braceKids(w, cur, &n.PO, n.Argument, n.GA)

	case *MOPredicate:
		return braceKids(w, cur, &n.MO, &n.Discourse, &n.TEO)

	case *LUPredicate:
		return braces(w, cur, (*LUPhrase)(n))

	case *CoPPredicate:
		return braces(w, cur, (*CoP)(n))

	case *LinkedTerm:
		return braceKids(w, cur, &n.GO, n.Argument)

	case Terms:
		kids := make([]interface{}, len(n))
		for i, k := range n {
			kids[i] = k
		}
		return braceKids(w, cur, kids...)

	case *TermSet:
		return braces(w, cur, (*CoP)(n))

	case *PredicateArgument:
		return braceKids(w, cur, n.Focus, n.Quantifier, n.Predicate, n.Relative)

	case *CoPArgument:
		return braces(w, cur, (*CoP)(n))

	case *PredicationRelative:
		return braces(w, cur, &n.Predication)

	case *LURelative:
		return braces(w, cur, (*LUPhrase)(n))

	case *CoPRelative:
		return braces(w, cur, (*CoP)(n))

	case *PredicateAdverb:
		return braces(w, cur, n.Predicate)

	case *CoPAdverb:
		return braces(w, cur, (*CoP)(n))

	case *PredicationPreposition:
		return braceKids(w, cur, n.Predicate, n.Argument)

	case *CoPPreposition:
		return braces(w, cur, (*CoP)(n))

	case *PredicationContent:
		return braces(w, cur, &n.Predication)

	case *LUContent:
		return braces(w, cur, (*LUPhrase)(n))

	case *CoPContent:
		return braces(w, cur, (*CoP)(n))

	case *Parenthetical:
		return braceKids(w, cur, &n.KIO, &n.Discourse, &n.KI)

	case *Incidental:
		return braceKids(w, cur, &n.JU, n.Statement)

	case *Vocative:
		return braceKids(w, cur, &n.HU, n.Argument)

	case *Interjection:
		return braces(w, cur, (*Word)(n))

	case *Space:
		return braces(w, cur, n.M)

	case *CoP:
		if n.TO0 == nil {
			return braceKids(w, cur, n.Left, &n.RU, n.Right)
		}
		return braceKids(w, cur, n.TO0, &n.RU, n.Left, n.TO1, n.Right)

	case *LUPhrase:
		return braceKids(w, cur, &n.LU, n.Statement)

	case *Word:
		if n == nil {
			return nil
		}
		return braceKids(w, cur, n.T, n.M)

	default:
		panic(fmt.Sprintf("unknown node type %T", n))
	}
}

// braceKids prints braces around a set of nodes.
// Each element of xs must be either a Node, a *Node, or a string.
// Nils are ignored. This includes nil *Node, nil Node,
// and a Node that contains a nil pointer or slice value.
// If fewer than two elements of xs is non-nil as per the above rules,
// then no braces are printed around the arguments.
func braceKids(w io.Writer, cur int, xs ...interface{}) error {
	nodes := make([]interface{}, 0, len(xs))
	for _, x := range xs {
		v := reflect.ValueOf(x)
		if !v.IsValid() {
			continue
		}

		if s, ok := x.(string); ok {
			nodes = append(nodes, s)
			continue
		}

		// If it's a pointer that is not a node,
		// it must be a pointer to a node.
		// Dereference it.
		if _, ok := x.(Node); !ok && v.Kind() == reflect.Ptr {
			if v.IsNil() {
				continue
			}
			v = v.Elem()
		}

		// Skip spaces by replacing them with their mod.
		n := v.Interface().(Node)
		for {
			s, ok := n.(*Space)
			if !ok {
				break
			}
			n = s.M
		}

		// Now, it must be a node.
		// If it's not nil, and it's not a pointer to nil,
		// add it to the list.
		v = reflect.ValueOf(n)
		k := v.Kind()
		if n != nil && ((k != reflect.Ptr && k != reflect.Slice) || !v.IsNil()) {
			nodes = append(nodes, n)
		}
	}
	next := cur
	if len(nodes) > 1 {
		if _, err := io.WriteString(w, openBraces[cur:cur+1]); err != nil {
			return err
		}
		next = (cur + 1) % len(openBraces)
	}
	for i, n := range nodes {
		if s, ok := n.(string); ok {
			if _, err := io.WriteString(w, s); err != nil {
				return err
			}
		} else {
			if err := braces(w, next, n.(Node)); err != nil {
				return err
			}
		}
		if i < len(nodes)-1 {
			if _, err := io.WriteString(w, " "); err != nil {
				return err
			}
		}
	}
	if len(nodes) > 1 {
		if _, err := io.WriteString(w, closeBraces[cur:cur+1]); err != nil {
			return err
		}
	}
	return nil
}
