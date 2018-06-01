package ast

import (
	"fmt"
	"reflect"
)

// A Visitor has its Visit method called on each node of the AST by the Visit function.
type Visitor interface {
	Visit(Node) Visitor
}

// A FuncVisitor is a Visitor that calls a function for each node in the AST.
type FuncVisitor func(Node)

// Visit implements the Visitor interface for FuncVisitor.
func (v FuncVisitor) Visit(n Node) Visitor {
	v(n)
	return v
}

// Visit traverses the AST in depth-first order.
// For each node, it calls v.Visit(node) with a non-nil node.
// If the returned visitor is non-nil, then Visit recurs,
// calling Visit on all children of the node with the returned visitor,
// and finally calling Visit on the returned visitor with nil.
//
// Visit considers all Word nodes, CoP nodes, and LU nodes
// to have a single child of type Word, CoP, and LUPhrase respectively.
// So, for example, if Visit recurs on a WordPredicate node,
// the next call to visit will be for a node of type Word.
func Visit(n Node, v Visitor) {
	if n == nil {
		return
	}
	if v := reflect.ValueOf(n); v.Kind() == reflect.Ptr && v.IsNil() {
		return
	}
	if v = v.Visit(n); v == nil {
		return
	}
	defer v.Visit(nil)

	switch n := n.(type) {
	case *Text:
		Visit(n.Leading, v)
		Visit(&n.Discourse, v)

	case *Discourse:
		for _, k := range *n {
			Visit(k, v)
		}

	case *StatementSentence:
		Visit(n.JE, v)
		Visit(n.Statement, v)
		Visit(n.DA, v)

	case *CoPSentence:
		Visit((*CoP)(n), v)

	case *Prenex:
		Visit(n.Terms, v)
		Visit(&n.BI, v)

	case *PrenexStatement:
		Visit(&n.Prenex, v)
		Visit(n.Statement, v)

	case *Predication:
		Visit(n.Predicate, v)
		Visit(n.Terms, v)
		Visit(n.NA, v)

	case *CoPStatement:
		Visit((*CoP)(n), v)

	case *PrefixedPredicate:
		Visit(&n.MU, v)
		Visit(n.Predicate, v)

	case *SerialPredicate:
		Visit(n.Left, v)
		Visit(n.Right, v)

	case *WordPredicate:
		Visit((*Word)(n), v)

	case *MIPredicate:
		Visit(&n.MI, v)
		Visit(n.Phrase, v)
		Visit(n.GA, v)

	case *POPredicate:
		Visit(&n.PO, v)
		Visit(n.Argument, v)
		Visit(n.GA, v)

	case *MOPredicate:
		Visit(&n.MO, v)
		Visit(&n.Discourse, v)
		Visit(&n.TEO, v)

	case *LUPredicate:
		Visit((*LUPhrase)(n), v)

	case *CoPPredicate:
		Visit((*CoP)(n), v)

	case *LinkedTerm:
		Visit(&n.GO, v)
		Visit(n.Argument, v)

	case Terms:
		for _, t := range n {
			Visit(t, v)
		}

	case *TermSet:
		Visit((*CoP)(n), v)

	case *PredicateArgument:
		Visit(n.Focus, v)
		Visit(n.Quantifier, v)
		Visit(n.Predicate, v)
		Visit(n.Relative, v)

	case *CoPArgument:
		Visit((*CoP)(n), v)

	case *PredicationRelative:
		Visit(&n.Predication, v)

	case *LURelative:
		Visit((*LUPhrase)(n), v)

	case *CoPRelative:
		Visit((*CoP)(n), v)

	case *PredicateAdverb:
		Visit(n.Predicate, v)

	case *CoPAdverb:
		Visit((*CoP)(n), v)

	case *PredicationPreposition:
		Visit(n.Predicate, v)
		Visit(n.Argument, v)

	case *CoPPreposition:
		Visit((*CoP)(n), v)

	case *PredicationContent:
		Visit(&n.Predication, v)

	case *LUContent:
		Visit((*LUPhrase)(n), v)

	case *CoPContent:
		Visit((*CoP)(n), v)

	case *Parenthetical:
		Visit(&n.KIO, v)
		Visit(&n.Discourse, v)
		Visit(&n.KI, v)

	case *Incidental:
		Visit(&n.JU, v)
		Visit(n.Statement, v)

	case *Vocative:
		Visit(&n.HU, v)
		Visit(n.Argument, v)

	case *Interjection:
		Visit((*Word)(n), v)

	case *Space:
		Visit((*Word)(n), v)

	case *CoP:
		if n.TO0 == nil {
			Visit(n.Left, v)
			Visit(&n.RU, v)
			Visit(n.Right, v)
			return
		}
		Visit(n.TO0, v)
		Visit(&n.RU, v)
		Visit(n.Left, v)
		Visit(n.TO1, v)
		Visit(n.Right, v)

	case *LUPhrase:
		Visit(&n.LU, v)
		Visit(n.Statement, v)

	case *Word:
		Visit(n.M, v)

	default:
		panic(fmt.Sprintf("unknown node type %T", n))
	}
}
