package logic

import "sort"

func simplify(n Statement) Statement {
	n.addHoa(nil)
	n = n.simplifyStatement(newState(n))
	sequenceVarNums(n)
	return n
}

func sequenceVarNums(n Statement) {
	var i int
	n.Visit(func(n Node) {
		switch n := n.(type) {
		case *Binding:
			n.N = i
			i++
		case *Variable:
			if n.Def != nil {
				n.N = n.Def.N
			} else {
				n.N = i
				i++
			}
		}
	})
}

func (n *Binding) addHoa(b *Binding) []int {
	var used []int
	if n.Restriction != nil {
		used = n.Restriction.addHoa(n)
	}
	return merge(used, n.Scope.addHoa(nil))
}

func (n *Connection) addHoa(b *Binding) []int {
	return merge(n.Left.addHoa(b), n.Right.addHoa(b))
}

func (n *Predication) addHoa(b *Binding) []int {
	used := n.Predicate.addHoa(nil)
	for _, arg := range n.Arguments {
		if arg != nil {
			used = merge(used, arg.addHoa(nil))
		}
	}
	if b == nil || len(n.Arguments) > 0 {
		return used
	}
	for _, u := range used {
		if u == b.N {
			return used
		}
	}
	n.Arguments = []Argument{&Variable{N: b.N, Def: b}}
	return used
}

func (n *Variable) addHoa(b *Binding) []int { return []int{n.N} }

func (n *Plurality) addHoa(b *Binding) []int {
	var used []int
	for _, arg := range n.Arguments {
		if arg != nil {
			used = merge(used, arg.addHoa(nil))
		}
	}
	return used
}

func (n *Abstraction) addHoa(b *Binding) []int { return n.Statement.addHoa(nil) }

func (n *Serial) addHoa(b *Binding) []int {
	return merge(n.Left.addHoa(nil), n.Right.addHoa(nil))
}

func (n *Swap) addHoa(b *Binding) []int { return n.Predicate.addHoa(nil) }

func (n *Quote) addHoa(b *Binding) []int { return nil }

func (n *Pred) addHoa(b *Binding) []int { return nil }

// merge merges two sorted slices of ints, in-place,
// into a single sorted slice with no dups.
func merge(l, r []int) []int {
	if len(l) == 0 {
		return r
	}
	l = append(l, r...)
	sort.Ints(l)
	j := 0
	for _, x := range l[1:] {
		if x > l[j] {
			j++
			l[j] = x
		}
	}
	return l
}

type state struct {
	sub      map[int]Argument
	predVars map[int]bool
}

func newState(n Statement) *state {
	s := &state{
		sub:      make(map[int]Argument),
		predVars: make(map[int]bool),
	}
	n.Visit(func(n Node) {
		p, ok := n.(*Predication)
		if !ok {
			return
		}
		v, ok := p.Predicate.(*Variable)
		if !ok {
			return
		}
		s.predVars[v.N] = true
	})
	return s
}

func (n *Binding) simplifyStatement(s *state) Statement {
	if n.Restriction != nil {
		n.Restriction = n.Restriction.simplifyStatement(s)
	}
	if sub(s, n) {
		return n.Scope.simplifyStatement(s)
	}
	n.Scope = n.Scope.simplifyStatement(s)
	return n
}

func sub(s *state, n *Binding) bool {
	pred, ok := n.Restriction.(*Predication)
	if !ok ||
		(n.Quantifier != The && n.Quantifier != All) ||
		len(pred.Arguments) != 2 ||
		pred.Arguments[0] == nil ||
		pred.Arguments[1] == nil {
		return false
	}
	p, ok := pred.Predicate.(*Pred)
	if !ok || p.Name != Equal {
		return false
	}
	left, right := pred.Arguments[0], pred.Arguments[1]
	l, ok := left.(*Variable)
	if !ok {
		if r, ok := right.(*Variable); ok {
			l, left, right = r, right, left
		} else {
			return false
		}
	}
	if s.predVars[l.N] {
		// This variable is used as a predicate; we cannot substitute it.
		return false
	}
	switch right.(type) {
	case *Quote, *Plurality, *Abstraction:
		s.sub[l.N] = right
	default:
		return false
	}
	return true
}

func (n *Connection) simplifyStatement(s *state) Statement {
	n.Left = n.Left.simplifyStatement(s)
	n.Right = n.Right.simplifyStatement(s)
	return n
}

func (n *Predication) simplifyStatement(s *state) Statement {
	for i, a := range n.Arguments {
		if a != nil {
			n.Arguments[i] = a.simplifyArgument(s)
		}
	}

	switch pred := n.Predicate.(type) {
	case *Swap:
		return swap(s, n, pred)
	case *Abstraction:
		if sub := lambdaSub(s, n, pred); sub != nil {
			return sub
		}
		break
	}
	n.Predicate = n.Predicate.simplifyPredicate(s)
	return n
}

func swap(s *state, n *Predication, swap *Swap) Statement {
	args := n.Arguments
	if swap.Pos >= len(n.Arguments) {
		args = make([]Argument, swap.Pos+1)
		copy(args, n.Arguments)
	}
	args[0], args[swap.Pos] = args[swap.Pos], args[0]
	n.Arguments = args
	n.Predicate = swap.Predicate
	return n.simplifyStatement(s)
}

func lambdaSub(s *state, n *Predication, pred *Abstraction) Statement {
	var nargs int
	for _, a := range n.Arguments {
		if a != nil {
			nargs++
		}
	}
	var nlambda int
	b, ok := pred.Statement.(*Binding)
	for ok && b.Quantifier == Lambda && b.Restriction == nil {
		nlambda++
		b, ok = b.Scope.(*Binding)
	}
	if nlambda == 0 || nargs != nlambda {
		return nil
	}

	var i int
	st := pred.Statement
	for i = 0; i < nlambda; i++ {
		b := st.(*Binding)
		s.sub[b.N] = n.Arguments[i]
		st = b.Scope
	}
	return st.simplifyStatement(s)
}

func (n *Variable) simplifyPredicate(s *state) Predicate { return n }

func (n *Variable) simplifyArgument(s *state) Argument {
	if a, ok := s.sub[n.N]; ok {
		return a
	}
	return n
}

func (n *Plurality) simplifyArgument(s *state) Argument {
	var args []Argument
	for _, a := range n.Arguments {
		if a == nil {
			continue
		}
		a = a.simplifyArgument(s)
		if plural, ok := a.(*Plurality); ok {
			args = append(args, plural.Arguments...)
		} else {
			args = append(args, a)
		}
	}
	n.Arguments = args
	return n
}

func (n *Abstraction) simplifyArgument(s *state) Argument {
	return n.simplify(s)
}

func (n *Abstraction) simplifyPredicate(s *state) Predicate {
	return n.simplify(s)
}

func (n *Abstraction) simplify(s *state) *Abstraction {
	n.Statement = n.Statement.simplifyStatement(s)
	return n
}

func (n *Serial) simplifyPredicate(s *state) Predicate {
	n.Left = n.Left.simplifyPredicate(s)
	n.Right = n.Right.simplifyPredicate(s)
	return n
}

func (n *Swap) simplifyPredicate(s *state) Predicate {
	n.Predicate = n.Predicate.simplifyPredicate(s)
	return n
}

func (n *Pred) simplifyPredicate(s *state) Predicate { return n }
func (n *Quote) simplifyArgument(s *state) Argument  { return n }
