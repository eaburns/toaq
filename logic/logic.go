// Package logic contains a simple logic representation of Toaq.
package logic

import (
	"strconv"

	"github.com/eaburns/toaq/ast"
)

// Node is implemented by all nodes in the Logic expression tree.
type Node interface {
	ASTNode() ast.Node

	// Visit calls a function for each node in the tree
	// in depth-first pre-order.
	Visit(func(Node))

	// print writes a string representation of the node
	// using a standard mathematical notation.
	print(*printer) error

	// addHoa adds implicit hoa arguments
	// to top-level restriction predications of a binding,
	// where the bound variable is un-used in the predication,
	// and the predication has no arguments.
	//
	// The argument is the binding
	// in which the current node is a restriction,
	// or nil.
	addHoa(*Binding) []int
}

// A Statement is a logical statement.
type Statement interface {
	Node
	simplifyStatement(*state) Statement
}

// A Quantifier is a quantifier type.
type Quantifier string

const (
	// The is ι quantifier.
	The Quantifier = "ι"
	// Lambda is the λ quantifier.
	Lambda = "λ"
	// Each is the ∀ quantifier.
	Each = "∀"
	// All is the Λ quantifier ([∃x : tuq(x, {λy P(…)(y)})]).
	All = "Λ"
	// Some is the ∃ quantifier.
	Some = "∃"
	// None is the ∄ quantifier.
	None = "∄"
	// Hi is the quantifier question quantifier.
	Hi = "¿"
)

// A Binding is tha binding of a new variable.
type Binding struct {
	Quantifier  Quantifier
	N           int
	Focus       bool
	Restriction Statement
	Scope       Statement
	AST         ast.Node
}

// A Connector is a logical connector like ∧ and ∨.
type Connector string

const (
	// And is the ∧ connector.
	And Connector = "∧"
	// Or is the ∨ connector.
	Or = "∨"
	// Ri is the connector question connector.
	Ri = "?"
)

// A Connection is a connection of two statements by a logical connector.
type Connection struct {
	Connector   Connector
	Left, Right Statement
	AST         ast.Node
}

// A Predication is the instantiation of a predicate with arguments.
type Predication struct {
	Predicate Predicate
	Arguments []Argument
	AST       ast.Node
}

// An Argument is a single argument of a predication.
type Argument interface {
	Node
	simplifyArgument(*state) Argument
}

// A Variable is a variable argument.
type Variable struct {
	N     int
	Def   *Binding
	Focus bool
	AST   ast.Node
}

// A Plurality is a plural constant.
type Plurality struct {
	Arguments []Argument
	AST       ast.Node
}

// A Quote is a quotation literal.
type Quote struct {
	Text string
	// FullText indicates whether this is a full-text quote
	// or just a quote of a single phrase.
	FullText bool
	AST      ast.Node
}

// A Predicate is the predicate of a predication.
type Predicate interface {
	Node
	simplifyPredicate(*state) Predicate
}

// A Pred is a word predicate.
type Pred struct {
	Name string
	AST  ast.Node
}

const (
	// Equal is the = operator.
	Equal = "="
	// Among is the ≺ operator.
	Among = "≺"
)

// An Abstraction is a predicate or an argument representing a statement.
type Abstraction struct {
	Statement Statement
	AST       ast.Node
}

// A Serial is a serial composition of two predicates.
type Serial struct {
	Left, Right Predicate
	AST         ast.Node
}

// A Swap is modification of an underlying predicate,
// that swaps the place of two of its arguments.
type Swap struct {
	Predicate Predicate
	Pos       int
	AST       ast.Node
}

func (n *Variable) PrettyPrint() string {
	var f string
	if n.Focus {
		f = "*"
	}
	return f + "Variable(" + strconv.Itoa(n.N) + ")"
}

func (n *Binding) ASTNode() ast.Node     { return n.AST }
func (n *Connection) ASTNode() ast.Node  { return n.AST }
func (n *Predication) ASTNode() ast.Node { return n.AST }
func (n *Variable) ASTNode() ast.Node    { return n.AST }
func (n *Plurality) ASTNode() ast.Node   { return n.AST }
func (n *Quote) ASTNode() ast.Node       { return n.AST }
func (n *Pred) ASTNode() ast.Node        { return n.AST }
func (n *Abstraction) ASTNode() ast.Node { return n.AST }
func (n *Serial) ASTNode() ast.Node      { return n.AST }
func (n *Swap) ASTNode() ast.Node        { return n.AST }

func (n *Quote) PrettyPrint() string { return `Quote("` + n.Text + `")` }
func (n *Pred) PrettyPrint() string  { return `"` + n.Name + `"` }
