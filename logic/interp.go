package logic

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/eaburns/toaq/ast"
	"github.com/eaburns/toaq/tone"
)

// Interpret returns a logic interpretation of a Toaq text.
// If the text consists of only fragments, then nil is returned.
func Interpret(text *ast.Text) Statement {
	s := newInterp().text(text)
	if s == nil {
		// The text only contained fragments.
		return nil
	}
	return simplify(s)
}

type interpreter struct {
	// n is the next variable number.
	n int

	// theScope holds bindings lifted to the top-most scope,
	// in increasing order of the closing of their scopes.
	theScope []*Binding

	// hiScope is a stack of subordinate-statement-top scopes,
	// containing lifted hi-bindings in increasing order
	// of the closing of their scopes.
	hiScope [][]*Binding

	// liftedScope maps from arguments to lifted bindings.
	// This is to ensure that lifted bindings are not duplicated
	// in the face of branching
	// For example **mai? ru pai? ji/** should not have 2 ι-bindigs for ji
	// even after branching into **mai? ji/ na ru pai? ji/ na**.
	liftedScope map[ast.Argument]*Binding

	// hoaScope is a stack of hoa substitutions.
	// The top is the variable that currently substitutes hoa.
	// If the binding is nil, then the hoa becomes a λx.
	hoaScope []*Binding

	// varScope is a stack of mappings from input variables to unique numbers.
	varScope []varScope

	// predications maps from a predicate to its containing predication,
	// adverb, or preposition node.
	// This is used to ensure that logic tree AST fields
	// point to the original AST even in the face of re-writes.
	predications map[ast.Predicate]ast.Node
}

func newInterp() *interpreter {
	return &interpreter{
		liftedScope:  make(map[ast.Argument]*Binding),
		predications: make(map[ast.Predicate]ast.Node),
	}
}

type varScope struct {
	v   string
	def *Binding
}

func (interp *interpreter) newVar() int {
	v := interp.n
	interp.n++
	return v
}

// pushVar pushes a new variable onto the scope.
func (interp *interpreter) pushVar(s string, def *Binding) int {
	v := interp.newVar()
	interp.varScope = append(interp.varScope, varScope{s, def})
	return v
}

// lookupVar returns the new name of an input variable and true,
// if the variable is in scope.
// Otherwise, it returns a new name for the variable and false.
// The new name is in the given tone.
func (interp *interpreter) lookupVar(s string) (int, *Binding) {
	for i := len(interp.varScope) - 1; i >= 0; i-- {
		b := interp.varScope[i]
		if b.v == s {
			return b.def.N, b.def
		}
	}
	return interp.newVar(), nil
}

func (interp *interpreter) pushHoa(def *Binding) {
	interp.hoaScope = append(interp.hoaScope, def)
}

func (interp *interpreter) popHoa() {
	interp.hoaScope = interp.hoaScope[:len(interp.hoaScope)-1]
}

func (interp *interpreter) text(n *ast.Text) Statement {
	return interp.discourse(&n.Discourse)
}

func (interp *interpreter) discourse(n *ast.Discourse) Statement {
	var st Statement
	for _, k := range *n {
		switch k := k.(type) {
		case ast.Fragment:
			// Ignored.
		case ast.Sentence:
			if st == nil {
				st = interp.sentence(k)
			} else {
				st = &Connection{
					Connector: And,
					Left:      st,
					Right:     interp.sentence(k),
					AST:       n,
				}
			}
		default:
			panic(fmt.Sprintf("unknown discourse item type %T", k))
		}
	}
	for i := len(interp.theScope) - 1; i >= 0; i-- {
		st = bindScope(interp.theScope[i], st)
	}
	return st
}

func (interp *interpreter) sentence(n ast.Sentence) Statement {
	switch n := n.(type) {
	case *ast.StatementSentence:
		return interp.subordinateStatement(n.Statement)
	case *ast.CoPSentence:
		if !hasText(n.RU, "ru") {
			panic("bad sentence connector: " + n.RU.T)
		}
		return &Connection{
			Connector: And,
			Left:      interp.sentence(n.Left.(ast.Sentence)),
			Right:     interp.sentence(n.Right.(ast.Sentence)),
			AST:       n,
		}
	default:
		panic(fmt.Sprintf("unknown sentence type %T", n))
	}
}

func (interp *interpreter) subordinateStatement(n ast.Statement) Statement {
	return interp.subordinate(func() Statement {
		return interp.statement(nil, n)
	})
}

// subordinate wraps a statement interpreting function
// with the necessary hi-binding handling for subordinate statements.
func (interp *interpreter) subordinate(f func() Statement) Statement {
	defer func(s [][]*Binding) { interp.hiScope = s }(interp.hiScope)
	interp.hiScope = append(interp.hiScope, nil)

	st := f()

	scope := interp.varScope
	hiScope := interp.hiScope[len(interp.hiScope)-1]
	for i := len(hiScope) - 1; i >= 0; i-- {
		hi := hiScope[i]
		scope = append(scope, varScope{"", hi})
		interp.clearHiFreeVarDefs(scope, hi)
		st = bindScope(hi, st)
	}
	return st
}

// clearHiFreeVarDefs clears the .Def field
// and assigns a new, unique number
// to all Variables over which the hi-binding was lifted.
func (interp *interpreter) clearHiFreeVarDefs(scope []varScope, hi *Binding) {
	hi.Visit(func(n Node) {
		v, ok := n.(*Variable)
		if !ok {
			return
		}
		for _, s := range scope {
			if v.Def == s.def {
				return
			}
		}
		v.Def = nil
		v.N = interp.n
		interp.n++
	})
}

func (interp *interpreter) statement(args arguments, n ast.Statement) Statement {
	defer func(s []varScope) { interp.varScope = s }(interp.varScope)

	switch n := n.(type) {
	case *ast.PrenexStatement:
		termsCopy := append(ast.Terms{}, n.Prenex.Terms...)
		return interp.prenexStatement(args, n, termsCopy)
	case *ast.Predication:
		if n.Terms == nil {
			return interp.predication(n.Predicate, args)
		}

		// We re-write the predication into a prenex statement,
		// but the AST field of the eventual predication or predications
		// should point to the original Predication node.
		interp.predications[n.Predicate] = n

		return interp.statement(args, &ast.PrenexStatement{
			Prenex: ast.Prenex{
				Terms: n.Terms,
				BI:    ast.Word{T: "pa"},
			},
			Statement: &ast.Predication{Predicate: n.Predicate, NA: n.NA},
		})
	case *ast.CoPStatement:
		return &Connection{
			Connector: connector(n.RU),
			Left:      interp.statement(args, n.Left.(ast.Statement)),
			Right:     interp.statement(args, n.Right.(ast.Statement)),
			AST:       n,
		}
	default:
		panic(fmt.Sprintf("unknown statement type %T", n))
	}
}

// An argument is the argument of a predication
// with an optional pos field that gives the position
// asserted by a linked term (go and cu).
// If pos=0 then the argument fills the next empty slot
// after all pos>0 arguments have filled their slots.
type argument struct {
	Argument
	pos int
}

type arguments []argument

func (args arguments) add(arg Argument) arguments {
	return append(args, argument{Argument: arg})
}

func (args arguments) addAt(arg Argument, pos int) arguments {
	return append(args, argument{Argument: arg, pos: pos})
}

func (interp *interpreter) predication(n ast.Predicate, args arguments) Statement {
	return interp.distributePredication(n, interp.predicate(n), args)
}

func (interp *interpreter) distributePredication(astPred ast.Predicate, n interface{}, args arguments) Statement {
	switch n := n.(type) {
	case *predConnection:
		return &Connection{
			Connector: n.con,
			Left:      interp.distributePredication(astPred, n.left, args),
			Right:     interp.distributePredication(astPred, n.right, args),
			AST:       n.AST,
		}
	case Predicate:
		// Use the original AST node for the AST field
		// in the case that this was re-written.
		astNode := ast.Node(astPred)
		if p, ok := interp.predications[astPred]; ok {
			astNode = p
		}
		return &Predication{
			Predicate: n,
			Arguments: orderArguments(args),
			AST:       astNode,
		}
	default:
		panic(fmt.Sprintf("unknown predicate type %T", n))
	}
}

func orderArguments(args arguments) []Argument {
	nargs := len(args)
	for _, a := range args {
		if a.pos > nargs {
			nargs = a.pos
		}
	}
	largs := make([]Argument, nargs)
	for _, a := range args {
		if a.pos > 0 {
			largs[a.pos-1] = a.Argument
		}
	}
	var i int
	for _, a := range args {
		if a.pos > 0 {
			continue
		}
		for largs[i] != nil {
			i++
		}
		largs[i] = a.Argument
	}
	return largs
}

// predConnection is like a Connection, but on Predicates instead of Statements.
// It is a temporary struct used only during interpreting predications,
// but is never present in the output of the interpreter,
// as all predConnections are converted to statements.
type predConnection struct {
	con   Connector
	left  interface{}
	right interface{}
	AST   ast.Node
}

// predicate interpretes a predicate.
// It returns either a Predicate or a *predConnection.
func (interp *interpreter) predicate(n ast.Predicate) interface{} {
	switch n := n.(type) {
	case *ast.WordPredicate:
		if _, s, ok := isVariablePredicate(n); ok {
			x, def := interp.lookupVar(s)
			return &Variable{N: x, Def: def, AST: n}
		}
		if isHoaPredicate(n) {
			if len(interp.hoaScope) == 0 {
				return &Variable{N: interp.newVar(), AST: n}
			}
			def := interp.hoaScope[len(interp.hoaScope)-1]
			return &Variable{N: def.N, Def: def, AST: n}
		}
		return &Pred{Name: wordPred(n), AST: n}
	case *ast.SerialPredicate:
		return interp.serialPredicate(n)
	case *ast.MIPredicate:
		quote := &Quote{
			Text:     toString(n.Phrase),
			FullText: false,
			AST:      n.Phrase,
		}
		return interp.makeLambda(n, miPred(&n.MI), quote)
	case *ast.POPredicate:
		return interp.poPredicate(n)
	case *ast.MOPredicate:
		equal := &Pred{Name: Equal, AST: n}
		quote := &Quote{
			Text:     toString(&n.Discourse),
			FullText: true,
			AST:      &n.Discourse,
		}
		return interp.makeLambda(n, equal, quote)
	case *ast.LUPredicate:
		var b *Binding
		if hasText(n.LU, "lu") {
			b = &Binding{N: interp.newVar(), Quantifier: Lambda, AST: n}
			interp.pushHoa(b)
			defer interp.popHoa()
		}
		st := interp.subordinateStatement(n.Statement)
		// lu? and li? return the abstraction as a λ-function.
		//
		// TODO: ma? and tio? return some predicate of the abstraction,
		// indicating their relationship to it. For example,
		// ma?: (λX X = ?{...})  // could be (λX mane(X, {...}))
		// tio?: (λX X = #{...})  // could be (λX ne(X, {...}))
		// If we keep lo?, it's just a content: (λX X = {...})
		return &Abstraction{Statement: bindScope(b, st), AST: n}
	case *ast.CoPPredicate:
		return &predConnection{
			con:   connector(n.RU),
			left:  interp.predicate(n.Left.(ast.Predicate)),
			right: interp.predicate(n.Right.(ast.Predicate)),
			AST:   n,
		}
	// Content clauses can be predicates when content argument
	// is converted into variable form:
	// **mai^** becomes **ta do/ maiX**,
	// where X is some non-existant tone denoting
	// a relative predication that is a content clause.
	// This is impossible in actual Toaq,
	// but it's possible as an intermediate state during interpretation.
	// Here, we reconcile it by converting to {abstraction}(·) form.
	case *ast.CoPContent:
		return &predConnection{
			con:   connector(n.RU),
			left:  interp.predicate(n.Left.(ast.Predicate)),
			right: interp.predicate(n.Right.(ast.Predicate)),
			AST:   n,
		}
	case *ast.LUContent:
		sub := interp.subordinateStatement(n.Statement)
		return interp.makeContentPred(n, sub)
	case *ast.PredicationContent:
		sub := interp.subordinateStatement(&n.Predication)
		return interp.makeContentPred(n, sub)
	default:
		panic(fmt.Sprintf("unknown predicate type %T", n))
	}
}

func (interp *interpreter) makeContentPred(n ast.Node, st Statement) Predicate {
	equal := &Pred{Name: Equal, AST: n}
	content := &Abstraction{Statement: st, AST: n}
	return interp.makeLambda(n, equal, content)
}

func (interp *interpreter) distributePredicatePrefix(n *ast.PrefixedPredicate, pred interface{}) interface{} {
	switch pred := pred.(type) {
	case *predConnection:
		return &predConnection{
			con:   pred.con,
			left:  interp.distributePredicatePrefix(n, pred.left),
			right: interp.distributePredicatePrefix(n, pred.right),
			AST:   pred.AST,
		}
	case Predicate:
		if !hasText(n.MU, "mu") {
			panic("unknown MU: " + n.MU.T)
		}
		return &Swap{Predicate: pred, Pos: 1, AST: n}
	default:
		panic(fmt.Sprintf("unknown predicate type %T", n))
	}
}

func (interp *interpreter) serialPredicate(n *ast.SerialPredicate) interface{} {
	l := interp.predicate(n.Left.(ast.Predicate))
	r := interp.predicate(n.Right.(ast.Predicate))
	return interp.serialPredicateLeft(n, l, r)
}

func (interp *interpreter) serialPredicateLeft(astNode ast.Node, l, r interface{}) interface{} {
	switch l := l.(type) {
	case *predConnection:
		return &predConnection{
			con:   l.con,
			left:  interp.serialPredicateLeft(astNode, l.left, r),
			right: interp.serialPredicateLeft(astNode, l.right, r),
			AST:   l.AST,
		}
	case Predicate:
		return interp.serialPredicateRight(astNode, l, r)
	default:
		panic(fmt.Sprintf("unknown predicate type %T", l))
	}
}

func (interp *interpreter) serialPredicateRight(astNode ast.Node, l Predicate, r interface{}) interface{} {
	switch r := r.(type) {
	case *predConnection:
		return &predConnection{
			con:   r.con,
			left:  interp.serialPredicateLeft(astNode, l, r.left),
			right: interp.serialPredicateLeft(astNode, l, r.right),
			AST:   r.AST,
		}
	case Predicate:
		return &Serial{Left: l, Right: r, AST: astNode}
	default:
		panic(fmt.Sprintf("unknown predicate type %T", r))
	}
}

func (interp *interpreter) poPredicate(n *ast.POPredicate) Predicate {
	pred := &Pred{Name: poPred(n.PO), AST: n}
	x := interp.newVar()
	arg0 := &Variable{N: x, AST: n}
	lambda := &Binding{
		N:          x,
		Quantifier: Lambda,
		Scope:      interp.poArgument(n, pred, arg0, n.Argument),
		AST:        n,
	}
	arg0.Def = lambda
	return &Abstraction{Statement: lambda, AST: n}
}

func (interp *interpreter) poArgument(n ast.Node, pred Predicate, arg0 Argument, arg1 ast.Argument) Statement {
	defer func(s []varScope) { interp.varScope = s }(interp.varScope)

	var b *Binding
	var a Argument
	switch arg1 := arg1.(type) {
	case *ast.PredicateArgument:
		b, a = interp.predicateArgument(arg1)
		break
	case *ast.CoPArgument:
		if hasText(arg1.RU, "roi") {
			b, a = interp.roi(arg1)
			break
		}
		return &Connection{
			Connector: connector(arg1.RU),
			Left:      interp.poArgument(n, pred, arg0, arg1.Left.(ast.Argument)),
			Right:     interp.poArgument(n, pred, arg0, arg1.Right.(ast.Argument)),
			AST:       n,
		}
	default:
		panic(fmt.Sprintf("unknown argument type %T", n))
	}
	st := &Predication{
		Predicate: pred,
		Arguments: []Argument{arg0, a},
		AST:       n,
	}
	return bindScope(b, st)
}

// prenexStatement interprets a prenex statement using the given term list.
// This function modifies the contents of the terms slice.
func (interp *interpreter) prenexStatement(args arguments, n *ast.PrenexStatement, terms ast.Terms) Statement {
	defer func(s []varScope) { interp.varScope = s }(interp.varScope)

	if len(terms) == 0 {
		if !hasText(n.Prenex.BI, "pa") {
			args = nil
		}
		return interp.statement(args, n.Statement)
	}
	switch t := terms[0].(type) {
	case *ast.TermSet:
		leftTerms := terms[1:].Prepend(t.Left.(ast.Terms)...)
		rightTerms := terms[1:].Prepend(t.Right.(ast.Terms)...)
		return &Connection{
			Connector: connector(t.RU),
			Left:      interp.prenexStatement(args, n, leftTerms),
			Right:     interp.prenexStatement(args, n, rightTerms),
			AST:       t,
		}
	case *ast.CoPArgument:
		if hasText(t.RU, "roi") {
			b, a := interp.roi(t)
			st := interp.prenexStatement(args.add(a), n, terms[1:])
			return bindScope(b, st)
		}
		l, r := t.Left.(ast.Term), t.Right.(ast.Term)
		return &Connection{
			Connector: connector(t.RU),
			Left:      interp.prenexStatement(args, n, withTerm0(l, terms)),
			Right:     interp.prenexStatement(args, n, withTerm0(r, terms)),
			AST:       t,
		}
	case *ast.PredicateArgument:
		b, a := interp.predicateArgument(t)
		if t.Focus != nil && b != nil && !hasText(n.Prenex.BI, "pa") {
			// The argument will not be used, but it had focus.
			// We don't want to lose that there was a focus word,
			// so put it on the prenex term instead.
			b.Focus = true
		}
		st := interp.prenexStatement(args.add(a), n, terms[1:])
		return bindScope(b, st)
	case *ast.LinkedTerm:
		var b *Binding
		var a Argument
		switch arg := t.Argument.(type) {
		case *ast.CoPArgument:
			if hasText(arg.RU, "roi") {
				b, a = interp.roi(arg)
				break
			}
			l := &ast.LinkedTerm{GO: t.GO, Argument: arg.Left.(ast.Argument)}
			r := &ast.LinkedTerm{GO: t.GO, Argument: arg.Right.(ast.Argument)}
			return &Connection{
				Connector: connector(arg.RU),
				Left:      interp.prenexStatement(args, n, withTerm0(l, terms)),
				Right:     interp.prenexStatement(args, n, withTerm0(r, terms)),
				AST:       arg,
			}
		case *ast.PredicateArgument:
			b, a = interp.predicateArgument(arg)
			break
		default:
			panic(fmt.Sprintf("unknown argument type %T", n))
		}
		args = args.addAt(a, goPos(t.GO))
		st := interp.prenexStatement(args, n, terms[1:])
		return bindScope(b, st)
	case ast.Preposition:
		sub := interp.prenexStatement(args, n, terms[1:])
		if hasFreeVars(sub) {
			return interp.preposition(t, &statementTail{args, n, terms[1:]})
		}
		x := interp.newVar()
		b := &Binding{N: x, Quantifier: The, AST: t}
		a := &Variable{N: x, Def: b, AST: t}
		b.Restriction = &Predication{
			Predicate: interp.makeContentPred(n, sub),
			Arguments: []Argument{a},
			AST:       t,
		}
		interp.theScope = append(interp.theScope, b)
		return interp.preposition(t, a)
	case ast.Adverb:
		sub := interp.prenexStatement(args, n, terms[1:])
		x := interp.newVar()
		b := &Binding{N: x, AST: t}
		a := &Variable{N: x, Def: b, AST: t}
		b.Restriction = &Predication{
			Predicate: interp.makeContentPred(n, sub),
			Arguments: []Argument{a},
			AST:       t,
		}
		if hasFreeVars(sub) {
			b.Quantifier = All
			return bindScope(b, interp.adverb(t, a))
		}
		b.Quantifier = The
		interp.theScope = append(interp.theScope, b)
		return interp.adverb(t, a)
	default:
		panic(fmt.Sprintf("unknown term type %T", n))
	}
}

type statementTail struct {
	args   arguments
	prenex *ast.PrenexStatement
	terms  ast.Terms
}

// arg0 is one of two things depending on whether
// the tail of the statement being subordinated
// by the preposition contains free variables.
//
// If it has no free variables mean that it can be ι-bound at top-scope.
// In this case, arg0 is fully-interpreted before the call to preposition;
// it is a *Variable ι-bound to the subordinated sentence.
//
// If it does have free variables, then it cannot be moved to top-scope,
// because that may move it over the binding of its free variables.
// Instead, it must be Λ-bound and in the current scope.
// In this case, arg0 is a *statementTail with the continuation
// of a call to prenexStatement, which must be in the scope
// of the arguments of the preposition.
// In other words, the continuation must be interpreted and Λ-bound
// only after interpreting the preposition arguments.
func (interp *interpreter) preposition(n ast.Preposition, arg0 interface{}) Statement {
	defer func(s []varScope) { interp.varScope = s }(interp.varScope)

	switch n := n.(type) {
	case *ast.PredicationPreposition:
		// The preposition will boil down to a predication or predications.
		// The AST field should point to the Preposition node.
		interp.predications[n.Predicate] = n

		return interp.predicationPreposition(n.Predicate, arg0, n.Argument)
	case *ast.CoPPreposition:
		return &Connection{
			Connector: connector(n.RU),
			Left:      interp.preposition(n.Left.(ast.Preposition), arg0),
			Right:     interp.preposition(n.Right.(ast.Preposition), arg0),
			AST:       n,
		}
	default:
		panic(fmt.Sprintf("unknown preposition type %T", n))
	}
}

func (interp *interpreter) predicationPreposition(p ast.Predicate, arg0 interface{}, arg1 ast.Argument) Statement {
	defer func(s []varScope) { interp.varScope = s }(interp.varScope)

	var b *Binding
	var a Argument
	switch arg1 := arg1.(type) {
	case *ast.PredicateArgument:
		b, a = interp.predicateArgument(arg1)
	case *ast.CoPArgument:
		if hasText(arg1.RU, "roi") {
			b, a = interp.roi(arg1)
			break
		}
		return &Connection{
			Connector: connector(arg1.RU),
			Left:      interp.predicationPreposition(p, arg0, arg1.Left.(ast.Argument)),
			Right:     interp.predicationPreposition(p, arg0, arg1.Left.(ast.Argument)),
			AST:       arg1,
		}
	default:
		panic(fmt.Sprintf("unknown argument type %T", arg1))
	}

	switch arg0 := arg0.(type) {
	case *Variable:
		st := interp.makePredication(p, arg0, a)
		return bindScope(b, st)

	case *statementTail:
		args, prenex, terms := arg0.args, arg0.prenex, arg0.terms
		sub := interp.prenexStatement(args, prenex, terms)
		x := interp.newVar()
		b1 := &Binding{N: x, Quantifier: All, AST: prenex}
		a0 := &Variable{N: x, Def: b1, AST: prenex}
		b1.Restriction = &Predication{
			Predicate: interp.makeContentPred(prenex, sub),
			Arguments: []Argument{a0},
			AST:       prenex,
		}
		st := interp.makePredication(p, a0, a)
		return bindScope(b, bindScope(b1, st))
	default:
		panic(fmt.Sprintf("unknown arg0 type %T", arg0))
	}
}

func (interp *interpreter) adverb(n ast.Adverb, arg Argument) Statement {
	defer func(s []varScope) { interp.varScope = s }(interp.varScope)

	switch n := n.(type) {
	case *ast.PredicateAdverb:
		// The adverb will boil down to a predication.
		// The AST field should point to the Adverb node.
		interp.predications[n.Predicate] = n

		return interp.makePredication(n.Predicate, arg)
	case *ast.CoPAdverb:
		return &Connection{
			Connector: connector(n.RU),
			Left:      interp.adverb(n.Left.(ast.Adverb), arg),
			Right:     interp.adverb(n.Right.(ast.Adverb), arg),
			AST:       n,
		}
	default:
		panic(fmt.Sprintf("unknown adverb type %T", n))
	}
}

func (interp *interpreter) roi(n *ast.CoPArgument) (*Binding, Argument) {
	x := interp.newVar()
	b := &Binding{N: x, Quantifier: The, AST: n}
	a := &Variable{N: x, Def: b, AST: n}
	b.Restriction = interp.roiLeft(a, n.Left.(ast.Argument), n.Right.(ast.Argument))
	return b, a
}

func (interp *interpreter) roiLeft(x *Variable, left, right ast.Argument) Statement {
	var a Argument
	var b *Binding
	switch n := left.(type) {
	case *ast.CoPArgument:
		if hasText(n.RU, "roi") {
			b, a = interp.roi(n)
			break
		}
		return &Connection{
			Connector: connector(n.RU),
			Left:      interp.roiLeft(x, n.Left.(ast.Argument), right),
			Right:     interp.roiLeft(x, n.Right.(ast.Argument), right),
			AST:       n,
		}
	case *ast.PredicateArgument:
		b, a = interp.predicateArgument(n)
	default:
		panic(fmt.Sprintf("unknown argument type %T", n))
	}
	return bindScope(b, interp.roiRight(x, a, right))
}

func (interp *interpreter) roiRight(x *Variable, left Argument, right ast.Argument) Statement {
	var a Argument
	var b *Binding
	switch n := right.(type) {
	case *ast.CoPArgument:
		if hasText(n.RU, "roi") {
			b, a = interp.roi(n)
			break
		}
		return &Connection{
			Connector: connector(n.RU),
			Left:      interp.roiRight(x, left, n.Left.(ast.Argument)),
			Right:     interp.roiRight(x, left, n.Right.(ast.Argument)),
			AST:       n,
		}
	case *ast.PredicateArgument:
		b, a = interp.predicateArgument(n)
	default:
		panic(fmt.Sprintf("unknown argument type %T", n))
	}
	plural := &Plurality{
		Arguments: []Argument{left, a},
		AST:       x.AST,
	}
	st := &Predication{
		Predicate: &Pred{Name: Equal, AST: x.AST},
		Arguments: []Argument{x, plural},
		AST:       x.AST,
	}
	return bindScope(b, st)
}

func (interp *interpreter) predicateArgument(n *ast.PredicateArgument) (*Binding, Argument) {
	if b, ok := interp.liftedScope[n]; ok {
		// This was a ι- or hi-quantified predicate
		// that already had its scope lifted
		// by a previous conditional branch.
		return nil, &Variable{N: b.N, Def: b, AST: n}
	}
	if _, _, ok := isVariablePredicate(n.Predicate); ok {
		return interp.varArgument(n)
	}
	if isHoaPredicate(n.Predicate) {
		return interp.hoaArgument(n)
	}

	x := interp.newVar()
	b := &Binding{N: x, Quantifier: quantifier(n.Quantifier), AST: n}
	a := &Variable{N: x, Def: b, Focus: n.Focus != nil, AST: n}
	b.Restriction = interp.makePredication(n.Predicate, a)
	if n.Relative != nil {
		b.Restriction = &Connection{
			Connector: And,
			Left:      b.Restriction,
			Right:     interp.relativeClause(b, n.Relative),
			AST:       n,
		}
	}

	if b.Quantifier == The {
		if hasFreeVars(b) {
			b.Quantifier = All
			return b, a
		}
		interp.liftedScope[n] = b
		interp.theScope = append(interp.theScope, b)
		return nil, a
	}

	if b.Quantifier == Hi {
		// Lift it to the top scope of the current statement.
		// If this is lifted over the bindings of vars referenced within,
		// they will be changed to un-bound variables
		// later when we push it onto the top-scope
		// of the containing subordinate.
		// (See: clearHiFreeVarDefs)
		interp.liftedScope[n] = b
		n := len(interp.hiScope) - 1
		interp.hiScope[n] = append(interp.hiScope[n], b)
		return nil, a
	}

	return b, a
}

func (interp *interpreter) varArgument(n *ast.PredicateArgument) (*Binding, Argument) {
	_, s, ok := isVariablePredicate(n.Predicate)
	if !ok {
		panic("impossible")
	}

	if n.Quantifier != nil {
		// A quantified variable creates a new binding.
		// We re-name the variable with a new, unique name,
		// and emit a term for its binding and an argument.
		b := &Binding{Quantifier: quantifier(n.Quantifier), AST: n}
		x := interp.pushVar(s, b)
		b.N = x
		a := &Variable{N: x, Def: b, AST: n}
		if n.Relative != nil {
			b.Restriction = interp.relativeClause(b, n.Relative)
		}
		return b, a
	}

	x, def := interp.lookupVar(s)
	ax := &Variable{N: x, Def: def, AST: n}
	if n.Relative == nil {
		return nil, ax
	}

	y := interp.newVar()
	b := &Binding{N: y, Quantifier: All, AST: n}
	ay := &Variable{N: y, Def: b, AST: n}
	b.Restriction = &Connection{
		Connector: And,
		Left: &Predication{
			Predicate: &Pred{Name: Among, AST: n},
			Arguments: []Argument{ay, ax},
			AST:       n,
		},
		Right: interp.relativeClause(b, n.Relative),
	}
	return b, ay
}

func (interp *interpreter) hoaArgument(n *ast.PredicateArgument) (*Binding, Argument) {
	if !isHoaPredicate(n.Predicate) {
		panic("impossible")
	}
	if len(interp.hoaScope) == 0 {
		// This is an un-bound, would-be λ-bound variable.
		return nil, &Variable{N: interp.newVar(), AST: n}
	}

	def := interp.hoaScope[len(interp.hoaScope)-1]
	x := def.N
	ax := &Variable{N: x, Def: def, AST: n}
	if n.Quantifier == nil && n.Relative == nil {
		return nil, ax
	}

	y := interp.newVar()
	b := &Binding{N: y, Quantifier: All, AST: n}
	ay := &Variable{N: y, Def: b, AST: n}
	b.Restriction = &Predication{
		Predicate: &Pred{Name: Among, AST: n},
		Arguments: []Argument{ay, ax},
		AST:       n,
	}
	if n.Relative != nil {
		b.Restriction = &Connection{
			Connector: And,
			Left:      b.Restriction,
			Right:     interp.relativeClause(b, n.Relative),
		}
	}
	return b, ay
}

func isHoaPredicate(n ast.Predicate) bool {
	w, ok := n.(*ast.WordPredicate)
	return ok && strings.ToLower(tone.WithTone(w.T, tone.None)) == "hoa"
}

func isVariablePredicate(n ast.Predicate) (rune, string, bool) {
	w, ok := n.(*ast.WordPredicate)
	if !ok {
		return 0, "", false
	}
	return splitVar(w.T)
}

func (interp *interpreter) relativeClause(b *Binding, n ast.Relative) Statement {
	defer func(s []varScope) { interp.varScope = s }(interp.varScope)

	if b != nil {
		interp.pushHoa(b)
		defer interp.popHoa()
	}

	switch n := n.(type) {
	case *ast.PredicationRelative:
		return interp.subordinateStatement(&n.Predication)
	case *ast.LURelative:
		return interp.subordinateStatement(n.Statement)
	case *ast.CoPRelative:
		return &Connection{
			Connector: connector(n.RU),
			Left:      interp.relativeClause(nil, n.Left.(ast.Relative)),
			Right:     interp.relativeClause(nil, n.Right.(ast.Relative)),
			AST:       n,
		}
	default:
		panic(fmt.Sprintf("unknown relative clause type %T", n))
	}
}

func (interp *interpreter) makePredication(pred ast.Predicate, args ...Argument) Statement {
	var as arguments
	for _, a := range args {
		as = as.add(a)
	}
	return interp.predication(pred, as)
}

func (interp *interpreter) makeLambda(n ast.Node, pred Predicate, arg1 Argument) *Abstraction {
	x := interp.newVar()
	arg0 := &Variable{N: x, AST: n}
	lambda := &Binding{
		N:          x,
		Quantifier: Lambda,
		Scope: &Predication{
			Predicate: pred,
			Arguments: []Argument{arg0, arg1},
		},
		AST: n,
	}
	arg0.Def = lambda
	return &Abstraction{Statement: lambda, AST: n}
}

func bindScope(bind *Binding, stmt Statement) Statement {
	if bind == nil {
		return stmt
	}
	if bind.Scope == nil {
		bind.Scope = stmt
	} else {
		bind.Scope = bindScope(bind.Scope.(*Binding), stmt)
	}
	return bind
}

func wordPred(w *ast.WordPredicate) string {
	var b strings.Builder
	// TODO: V tone isn't correctly removed.
	for _, r := range strings.ToLower(tone.ToASCII(w.T)) {
		switch {
		case r == 'v':
			// A tone marker; skip it.
			break
		case r == 'ı':
			b.WriteRune('i')
		case unicode.IsLetter(r):
			b.WriteRune(r)
		}
	}
	switch name := b.String(); name {
	case "mea":
		return Among
	case "jeqtu":
		return Equal
	default:
		return name
	}
}

func poPred(w ast.Word) string {
	switch {
	case hasText(w, "po"):
		return "raq"
	case hasText(w, "jei"):
		return Equal
	}
	panic("impossible: " + w.T)
}

func miPred(w *ast.Word) Predicate {
	switch {
	case hasText(*w, "mi"):
		return &Swap{Predicate: &Pred{Name: "chua", AST: w}, Pos: 1, AST: w}
	case hasText(*w, "shu"):
		return &Pred{Name: Equal, AST: w}
	}
	panic("impossible: " + w.T)
}

func quantifier(w *ast.Word) Quantifier {
	switch {
	case w == nil || hasText(*w, "ta"):
		return The
	case hasText(*w, "ja"):
		return Lambda
	case hasText(*w, "sa"):
		return Some
	case hasText(*w, "tu"):
		return Each
	case hasText(*w, "sia"):
		return None
	case hasText(*w, "hi"):
		return Hi
	}
	panic("impossible: " + w.T)
}

func connector(w ast.Word) Connector {
	switch {
	case hasText(w, "ru"):
		return And
	case hasText(w, "ra"):
		return Or
	case hasText(w, "ri"):
		return Ri
	case hasText(w, "roi"):
		panic("bad connector: " + w.T)
	}
	panic("impossible: " + w.T)
}

func goPos(w ast.Word) int {
	switch {
	case hasText(w, "go"):
		return 2
	case hasText(w, "cu"):
		return 3
	}
	panic("impossible: " + w.T)
}

// withTerm0 sets ts[0] to t and returns ts.
func withTerm0(t ast.Term, ts ast.Terms) ast.Terms {
	ts[0] = t
	return ts
}

// hasFreeVars returns whether a logic-sub-tree has un-bound variables.
func hasFreeVars(n Node) bool {
	free := false
	binds := make(map[*Binding]bool)
	n.Visit(func(n Node) {
		switch n := n.(type) {
		case *Binding:
			binds[n] = true
		case *Variable:
			if n.Def != nil && n.Def.Quantifier != The && !binds[n.Def] {
				free = true
			}
		}
	})
	return free
}

// hasText returns whether the text of a word equals (case insensitive) a string.
// If the word or string has tone markers, they are assumed to be of the same format
// (both ASCII, both diacritics, both using compose runes, and so on).
func hasText(w ast.Word, b string) bool {
	a := tone.WithTone(w.T, tone.None)
	for len(a) > 0 && len(b) > 0 {
		ra, wa := utf8.DecodeRuneInString(a)
		rb, wb := utf8.DecodeRuneInString(b)
		if ra == 'ı' {
			ra = 'i'
		}
		if rb == 'ı' {
			rb = 'i'
		}
		if unicode.ToLower(ra) != unicode.ToLower(rb) {
			return false
		}
		a, b = a[wa:], b[wb:]
	}
	return len(a) == len(b)
}

// oLetters contains all of the diacritic variants of the letter 'o'.
// oTones contains the tone marker of the 'o' at the same index in oLetter.
var oLetters, oTones = func() ([]rune, []rune) {
	var o, ot []rune
	for t, asciis := range tone.Diacritic {
		for ascii, d := range asciis {
			if ascii != "o" && ascii != "O" {
				continue
			}
			r, _ := utf8.DecodeRuneInString(d)
			o = append(o, r)
			ot = append(ot, t)
		}
	}
	return o, ot
}()

// splitVar splits the tone off of a do-variable,
// returning the tone, the un-toned variable, and true.
// If the text does not begin with "do" in any tone,
// then false is returned.
func splitVar(txt string) (rune, string, bool) {
	if len(txt) < 1 || txt[0] != 'd' || txt[0] == 'D' {
		return tone.None, "", false
	}
	r, w := utf8.DecodeRuneInString(txt[1:])
	i := -1
	for j, o := range oLetters {
		if o == r {
			i = j
			break
		}
	}
	if i < 0 {
		return tone.None, "", false
	}
	w++ // 'd'
	if w < len(txt) && strings.ContainsRune("aeiıouAEIOUqQ", rune(txt[w])) {
		// If o is no the end of the syllable, then this is not a variable.
		return tone.None, "", false
	}
	return oTones[i], strings.ToLower("do" + txt[w:]), true
}

func toString(node ast.Node) string {
	var s strings.Builder
	ast.Visit(node, ast.FuncVisitor(func(n ast.Node) {
		if w, ok := n.(*ast.Word); ok {
			s.WriteString(w.T)
		}
	}))
	return strings.TrimSpace(s.String())
}
