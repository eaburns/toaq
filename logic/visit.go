package logic

import "reflect"

func (n *Binding) Visit(f func(Node)) {
	f(n)
	visit(f, n.Restriction, n.Scope)
}

func (n *Connection) Visit(f func(Node)) {
	f(n)
	visit(f, n.Left, n.Right)
}

func (n *Predication) Visit(f func(Node)) {
	f(n)
	visit(f, n.Predicate)
	for _, a := range n.Arguments {
		visit(f, a)
	}
}

func (n *Variable) Visit(f func(Node)) { f(n) }

func (n *Plurality) Visit(f func(Node)) {
	f(n)
	for _, a := range n.Arguments {
		visit(f, a)
	}
}

func (n *Quote) Visit(f func(Node)) { f(n) }

func (n *Pred) Visit(f func(Node)) { f(n) }

func (n *Abstraction) Visit(f func(Node)) {
	f(n)
	visit(f, n.Statement)
}

func (n *Serial) Visit(f func(Node)) {
	f(n)
	visit(f, n.Left, n.Right)
}

func (n *Swap) Visit(f func(Node)) {
	f(n)
	visit(f, n.Predicate)
}

func visit(f func(Node), nodes ...Node) {
	for _, n := range nodes {
		if n != nil {
			n.Visit(f)
		}
	}
}

// NillOutAST sets all AST fields to nil.
// This is useful for de-cluttering a tree before printing it.
func NillOutAST(n Node) {
	n.Visit(func(n Node) {
		if n == nil {
			return
		}
		v := reflect.ValueOf(n)
		if k := v.Kind(); k == reflect.Ptr || k == reflect.Interface {
			if v.IsNil() {
				return
			}
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return
		}
		f := v.FieldByName("AST")
		if !f.IsValid() || f.Kind() != reflect.Interface {
			return
		}
		f.Set(reflect.Zero(f.Type()))
	})
}
