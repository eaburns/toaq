package ast

import (
	"fmt"
	"reflect"
)

// Visit calls f for each node in the tree in pre-, depth-first order
// (first f is called on the node itself, then its children).
//
// For nodes that are also Words, Visit is called on the node as a Word
// as though the Word were its single child;
// Mods are considered children of the Word node.
//
// If f return false for any node, the traversal stops and Visit return false.
func Visit(n Node, f func(Node) bool) bool {
	if n == nil {
		return true
	}
	if v := reflect.ValueOf(n); v.Kind() == reflect.Ptr && v.IsNil() {
		return true
	}
	switch n := n.(type) {
	case *Text:
		return f(n) && Visit(n.Leading, f) && Visit(&n.Discourse, f)

	case *Discourse:
		if !f(n) {
			return false
		}
		for _, k := range *n {
			if !Visit(k, f) {
				return false
			}
		}
		return true

	case *StatementSentence:
		return f(n) && Visit(n.JE, f) && Visit(n.Statement, f) && Visit(n.DA, f)

	case *CoPSentence:
		return f(n) && visitCoP((*CoP)(n), f)

	case *Prenex:
		return f(n) && Visit(&n.Terms, f) && Visit(&n.BI, f)

	case *Predication:
		return f(n) && Visit(n.Prenex, f) && Visit(n.Predicate, f) && Visit(n.Terms, f) && Visit(n.NA, f)

	case *CoPStatement:
		return f(n) && visitCoP((*CoP)(n), f)

	case *PrefixedPredicate:
		return f(n) && Visit(&n.MU, f) && Visit(n.Predicate, f)

	case *SerialPredicate:
		return f(n) && Visit(n.Left, f) && Visit(n.Right, f)

	case *WordPredicate:
		return f(n) && Visit((*Word)(n), f)

	case *MIPredicate:
		return f(n) && Visit(&n.MI, f) && Visit(n.Phrase, f) && Visit(n.GA, f)

	case *POPredicate:
		return f(n) && Visit(&n.PO, f) && Visit(n.Argument, f) && Visit(n.GA, f)

	case *MOPredicate:
		return f(n) && Visit(&n.MO, f) && Visit(&n.Discourse, f) && Visit(&n.TEO, f)

	case *LUPredicate:
		return f(n) && Visit(&n.LU, f) && Visit(n.Statement, f)

	case *CoPPredicate:
		return f(n) && visitCoP((*CoP)(n), f)

	case *LinkedTerm:
		return f(n) && Visit(&n.GO, f) && Visit(n.Argument, f)

	case *Terms:
		return f(n) && Visit(n.Term, f) && Visit(n.Terms, f)

	case *TermSet:
		return f(n) && visitCoP((*CoP)(n), f)

	case *PredicateArgument:
		return f(n) && Visit(n.Focus, f) && Visit(n.Quantifier, f) && Visit(n.Predicate, f) && Visit(n.Relative, f)

	case *CoPArgument:
		return f(n) && visitCoP((*CoP)(n), f)

	case *PredicationRelative:
		return f(n) && Visit(&n.Predication, f)

	case *LURelative:
		return f(n) && Visit(&n.LU, f) && Visit(n.Statement, f)

	case *CoPRelative:
		return f(n) && visitCoP((*CoP)(n), f)

	case *PredicateAdverb:
		return f(n) && Visit(n.Predicate, f)

	case *CoPAdverb:
		return f(n) && visitCoP((*CoP)(n), f)

	case *PredicationPreposition:
		return f(n) && Visit(n.Predicate, f) && Visit(n.Argument, f)

	case *CoPPreposition:
		return f(n) && visitCoP((*CoP)(n), f)

	case *PredicationContent:
		return f(n) && Visit(&n.Predication, f)

	case *LUContent:
		return f(n) && Visit(&n.LU, f) && Visit(n.Statement, f)

	case *CoPContent:
		return f(n) && visitCoP((*CoP)(n), f)

	case *Parenthetical:
		return f(n) && Visit(&n.KIO, f) && Visit(&n.Discourse, f) && Visit(&n.KI, f)

	case *Incidental:
		return f(n) && Visit(&n.JU, f) && Visit(n.Statement, f)

	case *Vocative:
		return f(n) && Visit(&n.HU, f) && Visit(n.Argument, f)

	case *Interjection:
		return f(n) && Visit((*Word)(n), f)

	case *Space:
		return f(n) && Visit((*Word)(n), f)

	case *Word:
		return n == nil || f(n) && (n.M == nil || Visit(n.M, f))

	default:
		panic(fmt.Sprintf("unknown node type %T", n))
	}
}

func visitCoP(n *CoP, f func(Node) bool) bool {
	if n.TO0 == nil {
		return Visit(n.Left, f) && Visit(&n.RU, f) && Visit(n.Right, f)
	}
	return Visit(n.TO0, f) && Visit(&n.RU, f) && Visit(n.Left, f) && Visit(n.TO1, f) && Visit(n.Right, f)
}
