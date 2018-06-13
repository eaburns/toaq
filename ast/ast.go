// Package ast contains an abstract syntax tree and Toaq parser.
package ast

import "github.com/eaburns/toaq/tone"

// A note on node naming:
// Besides the Node interface, all nodes belong to a more specific interface.
// For example, Predicate, Argument, Relative, and so on.
// The names are nodes are such that the interface name comes last
// and the specifier describing what the node is comes first.
// This way, node names read as a noun phrase
// with the noun being the node's interface and an adjective
// describing, more specifically, what the node is.
// For example, a PredicateArgument is an Argument-type node
// formed from a Predicate.

// Node is implemented by all AST nodes.
type Node interface {
	// Start returns the byte-offset into the text of the start of the node.
	Start() int
	// End returns the exclusive byte-offset into the text of the end of the node.
	End() int

	modNode(*Mod) Node
}

//
// Text
//

// A Text is a complete Toaq text.
type Text struct {
	Leading   Mod
	Discourse Discourse
}

func (n *Text) Start() int         { return n.Discourse.Start() }
func (n *Text) End() int           { return n.Discourse.End() }
func (n Text) modNode(m *Mod) Node { n.Discourse = *n.Discourse.mod(m); return &n }

// A Discourse is a non-empty sequence of sentences and/or fragments.
type Discourse []Node

func (n *Discourse) Start() int { return (*n)[0].Start() }
func (n *Discourse) End() int   { return (*n)[len(*n)-1].End() }

func (n Discourse) modNode(m *Mod) Node { return n.mod(m) }

func (n Discourse) mod(m *Mod) *Discourse {
	l := len(n)
	c := make(Discourse, l)
	copy(c, n)
	if s, ok := c[l-1].(Sentence); ok {
		c[l-1] = s.modSentence(m)
	} else {
		c[l-1] = c[l-1].(Fragment).modFragment(m)
	}
	return &c
}

//
// Sentences.
//

// Sentence is implemented by all sentence nodes.
type Sentence interface {
	Node

	modSentence(*Mod) Sentence
}

// A StatementSentence is a statement that is a sentence.
type StatementSentence struct {
	JE        *Word
	Statement Statement
	DA        *Word
}

func (n *StatementSentence) Start() int         { return start(n.JE, n.Statement) }
func (n *StatementSentence) End() int           { return end(n.DA, n.Statement) }
func (n StatementSentence) modNode(m *Mod) Node { return n.modSentence(m) }

func (n StatementSentence) modSentence(m *Mod) Sentence {
	if n.DA != nil {
		n.DA = n.DA.mod(m)
	} else {
		n.Statement = n.Statement.modStatement(m)
	}
	return &n
}

// A CoPSentence is a connector phrase sententce.
type CoPSentence CoP

func (n *CoPSentence) Start() int                 { return (*CoP)(n).Start() }
func (n *CoPSentence) End() int                   { return (*CoP)(n).End() }
func (n CoPSentence) modNode(m *Mod) Node         { return n.mod(m) }
func (n CoPSentence) modSentence(m *Mod) Sentence { return n.mod(m) }
func (n CoPSentence) mod(m *Mod) *CoPSentence     { return (*CoPSentence)(CoP(n).mod(m)) }

//
// Fragments
//

// Fragment is implemented by all fragment nodes.
type Fragment interface {
	Node
	modFragment(*Mod) Fragment
}

// A Prenex is a prenex.
type Prenex struct {
	Terms Terms
	BI    Word
}

func (n *Prenex) Start() int                 { return n.Terms.Start() }
func (n *Prenex) End() int                   { return n.BI.End() }
func (n Prenex) modNode(m *Mod) Node         { return n.mod(m) }
func (n Prenex) modFragment(m *Mod) Fragment { return n.mod(m) }

func (n Prenex) mod(m *Mod) *Prenex {
	n.BI = *n.BI.mod(m)
	return &n
}

//
// Statement nodes.
//

// Statement is implemented by all statement nodes.
type Statement interface {
	Node
	modStatement(*Mod) Statement
}

// A PrenexStatement is a statement with a prenex.
type PrenexStatement struct {
	Prenex    Prenex
	Statement Statement
}

func (n *PrenexStatement) Start() int                   { return n.Prenex.Start() }
func (n *PrenexStatement) End() int                     { return n.Statement.End() }
func (n PrenexStatement) modNode(m *Mod) Node           { return n.mod(m) }
func (n PrenexStatement) modStatement(m *Mod) Statement { return n.mod(m) }

func (n PrenexStatement) mod(m *Mod) *PrenexStatement {
	n.Statement = n.Statement.modStatement(m)
	return &n
}

// A Predication is a statement with a predicate and terms.
type Predication struct {
	Predicate Predicate
	Terms     Terms
	NA        *Word
}

func (n *Predication) Start() int                   { return n.Predicate.Start() }
func (n *Predication) End() int                     { return end(n.NA, n.Terms) }
func (n Predication) modNode(m *Mod) Node           { return n.mod(m) }
func (n Predication) modStatement(m *Mod) Statement { return n.mod(m) }

func (n Predication) mod(m *Mod) *Predication {
	if n.NA != nil {
		n.NA = n.NA.mod(m)
	} else if n.Terms != nil {
		n.Terms = n.Terms.mod(m)
	} else {
		n.Predicate = n.Predicate.modPredicate(m)
	}
	return &n
}

// A CoPStatement is a connector phrase statement.
type CoPStatement CoP

func (n *CoPStatement) Start() int                   { return (*CoP)(n).Start() }
func (n *CoPStatement) End() int                     { return (*CoP)(n).End() }
func (n CoPStatement) modNode(m *Mod) Node           { return n.modStatement(m) }
func (n CoPStatement) modStatement(m *Mod) Statement { return (*CoPStatement)(CoP(n).mod(m)) }

//
// Predicate nodes
//

// The Predicate interface is implemented by all predicate nodes.
type Predicate interface {
	Node
	modPhrase(*Mod) Phrase
	modPredicate(*Mod) Predicate
}

// A PrefixedPredicate is a predicate with a prefix word.
type PrefixedPredicate struct {
	MU        Word
	Predicate Predicate
}

func (n *PrefixedPredicate) Start() int             { return n.MU.Start() }
func (n *PrefixedPredicate) End() int               { return n.Predicate.End() }
func (n PrefixedPredicate) modNode(m *Mod) Node     { return n.modPredicate(m) }
func (n PrefixedPredicate) modPhrase(m *Mod) Phrase { return n.modPredicate(m) }

func (n PrefixedPredicate) modPredicate(m *Mod) Predicate {
	n.Predicate = n.Predicate.modPredicate(m)
	return &n
}

// A SerialPredicate is a serial predicate phrase.
type SerialPredicate struct {
	Left, Right Predicate
}

func (n *SerialPredicate) Start() int             { return n.Left.Start() }
func (n *SerialPredicate) End() int               { return n.Right.End() }
func (n SerialPredicate) modNode(m *Mod) Node     { return n.modPredicate(m) }
func (n SerialPredicate) modPhrase(m *Mod) Phrase { return n.modPredicate(m) }

func (n SerialPredicate) modPredicate(m *Mod) Predicate {
	n.Right = n.Right.modPredicate(m)
	return &n
}

// A WordPredicate is a word that is a predicate.
type WordPredicate Word

func (n WordPredicate) PrettyPrint() string           { return `WordPredicate("` + n.T + `")` }
func (n *WordPredicate) Start() int                   { return (*Word)(n).Start() }
func (n *WordPredicate) End() int                     { return (*Word)(n).End() }
func (n WordPredicate) modNode(m *Mod) Node           { return n.modPredicate(m) }
func (n WordPredicate) modPhrase(m *Mod) Phrase       { return n.modPredicate(m) }
func (n WordPredicate) modPredicate(m *Mod) Predicate { return (*WordPredicate)(Word(n).mod(m)) }

// Phrase is implemented by all nodes
// that can be the complement of a MI phrase.
type Phrase interface {
	Node
	modPhrase(*Mod) Phrase
}

// A MIPredicate is a predicate beginning with word-class MI.
type MIPredicate struct {
	MI     Word
	GA     *Word
	Phrase Phrase
}

func (n *MIPredicate) Start() int             { return n.MI.Start() }
func (n *MIPredicate) End() int               { return end(n.GA, n.Phrase) }
func (n MIPredicate) modNode(m *Mod) Node     { return n.modPredicate(m) }
func (n MIPredicate) modPhrase(m *Mod) Phrase { return n.modPredicate(m) }

func (n MIPredicate) modPredicate(m *Mod) Predicate {
	if n.GA != nil {
		n.GA = n.GA.mod(m)
	} else {
		n.Phrase = n.Phrase.modPhrase(m)
	}
	return &n
}

// A POPredicate is a predicate beginning with word-class PO.
type POPredicate struct {
	PO       Word
	GA       *Word
	Argument Argument
}

func (n *POPredicate) Start() int             { return n.PO.Start() }
func (n *POPredicate) End() int               { return end(n.GA, n.Argument) }
func (n POPredicate) modNode(m *Mod) Node     { return n.modPredicate(m) }
func (n POPredicate) modPhrase(m *Mod) Phrase { return n.modPredicate(m) }

func (n POPredicate) modPredicate(m *Mod) Predicate {
	if n.GA != nil {
		n.GA = n.GA.mod(m)
	} else {
		n.Argument = n.Argument.modArgument(m)
	}
	return &n
}

// A MOPredicate is a predicate beginning with word-class MO.
type MOPredicate struct {
	MO, TEO   Word
	Discourse Discourse
}

func (n *MOPredicate) Start() int                   { return n.MO.Start() }
func (n *MOPredicate) End() int                     { return n.TEO.End() }
func (n MOPredicate) modNode(m *Mod) Node           { return n.modPredicate(m) }
func (n MOPredicate) modPhrase(m *Mod) Phrase       { return n.modPredicate(m) }
func (n MOPredicate) modPredicate(m *Mod) Predicate { n.TEO = *n.TEO.mod(m); return &n }

// A LUPredicate is a LU phrase predicate.
type LUPredicate LUPhrase

func (n *LUPredicate) Start() int                   { return (*LUPhrase)(n).Start() }
func (n *LUPredicate) End() int                     { return (*LUPhrase)(n).End() }
func (n LUPredicate) modNode(m *Mod) Node           { return n.modPredicate(m) }
func (n LUPredicate) modPhrase(m *Mod) Phrase       { return n.modPredicate(m) }
func (n LUPredicate) modPredicate(m *Mod) Predicate { return (*LUPredicate)(LUPhrase(n).mod(m)) }

// A CoPPredicate is a connector predicate phrase.
type CoPPredicate CoP

func (n *CoPPredicate) Start() int                   { return (*CoP)(n).Start() }
func (n *CoPPredicate) End() int                     { return (*CoP)(n).End() }
func (n CoPPredicate) modNode(m *Mod) Node           { return n.mod(m) }
func (n CoPPredicate) modPhrase(m *Mod) Phrase       { return n.mod(m) }
func (n CoPPredicate) modPredicate(m *Mod) Predicate { return n.mod(m) }
func (n CoPPredicate) mod(m *Mod) *CoPPredicate      { return (*CoPPredicate)(CoP(n).mod(m)) }

//
// Terms
//

// Term is implemented by all term nodes.
type Term interface {
	Node
	modTerm(*Mod) Term
}

// A LinkedTerm is an argument preceeded by a linking word.
type LinkedTerm struct {
	GO       Word
	Argument Argument
}

func (n *LinkedTerm) Start() int         { return n.GO.Start() }
func (n *LinkedTerm) End() int           { return n.Argument.End() }
func (n LinkedTerm) modNode(m *Mod) Node { return n.modTerm(m) }

func (n LinkedTerm) modTerm(m *Mod) Term {
	n.Argument = n.Argument.modArgument(m)
	return &n
}

// A Terms is a silce of Terms.
type Terms []Term

func (n Terms) Start() int                  { return n[0].Start() }
func (n Terms) End() int                    { return n[len(n)-1].End() }
func (n Terms) modNode(m *Mod) Node         { return n.mod(m) }
func (n Terms) modFragment(m *Mod) Fragment { return n.mod(m) }

func (n Terms) mod(m *Mod) Terms {
	if m == nil {
		return n
	}
	l := len(n)
	c := make(Terms, l)
	copy(c, n)
	c[l-1] = c[l-1].modTerm(m)
	return c
}

// Prepend returns a copy of the receiver with the given Terms prepended.
func (n Terms) Prepend(ts ...Term) Terms {
	c := make(Terms, len(n)+len(ts))
	copy(c[len(ts):], n)
	copy(c, ts)
	return c
}

// A TermSet is a connector of terms.
type TermSet CoP

func (n *TermSet) Start() int         { return (*CoP)(n).Start() }
func (n *TermSet) End() int           { return (*CoP)(n).End() }
func (n TermSet) modNode(m *Mod) Node { return n.mod(m) }
func (n TermSet) modTerm(m *Mod) Term { return n.mod(m) }
func (n TermSet) mod(m *Mod) *TermSet { return (*TermSet)(CoP(n).mod(m)) }

//
// Argument nodes.
//

// Argument is implemented by all argument nodes.
type Argument interface {
	Node
	modTerm(*Mod) Term
	modPhrase(*Mod) Phrase
	modArgument(*Mod) Argument
}

// A PredicateArgument is a predicate that is an argument.
type PredicateArgument struct {
	Focus      *Word
	Quantifier *Word
	Predicate  Predicate
	// Relative is nil if there is no relative clause.
	Relative Relative
}

func (n *PredicateArgument) Start() int {
	if n.Focus != nil {
		return n.Focus.Start()
	}
	return start(n.Quantifier, n.Predicate)
}

func (n *PredicateArgument) End() int {
	if n.Relative != nil {
		return n.Relative.End()
	}
	return n.Predicate.End()
}

func (n PredicateArgument) modNode(m *Mod) Node         { return n.mod(m) }
func (n PredicateArgument) modTerm(m *Mod) Term         { return n.mod(m) }
func (n PredicateArgument) modPhrase(m *Mod) Phrase     { return n.mod(m) }
func (n PredicateArgument) modArgument(m *Mod) Argument { return n.mod(m) }

func (n PredicateArgument) mod(m *Mod) *PredicateArgument {
	if n.Relative != nil {
		n.Relative = n.Relative.modRelative(m)
	} else {
		n.Predicate = n.Predicate.modPredicate(m)
	}
	return &n
}

// An CoPArgument is an argument connector phrase.
type CoPArgument CoP

func (n *CoPArgument) Start() int                 { return (*CoP)(n).Start() }
func (n *CoPArgument) End() int                   { return (*CoP)(n).End() }
func (n CoPArgument) modNode(m *Mod) Node         { return n.modArgument(m) }
func (n CoPArgument) modTerm(m *Mod) Term         { return n.modArgument(m) }
func (n CoPArgument) modPhrase(m *Mod) Phrase     { return n.modArgument(m) }
func (n CoPArgument) modArgument(m *Mod) Argument { return (*CoPArgument)(CoP(n).mod(m)) }

//
// Relative clasue nodes.
//

// Relative is implemented by all relative clause nodes.
type Relative interface {
	Node
	modFragment(*Mod) Fragment
	modRelative(*Mod) Relative
}

// A PredicationRelative is a statement relative clause.
type PredicationRelative struct{ Predication }

func (n PredicationRelative) modNode(m *Mod) Node         { return n.modRelative(m) }
func (n PredicationRelative) modFragment(m *Mod) Fragment { return n.modRelative(m) }

func (n PredicationRelative) modRelative(m *Mod) Relative {
	n.Predication = *n.Predication.mod(m)
	return &n
}

// A LURelative is a LU phrase relative clause.
type LURelative LUPhrase

func (n *LURelative) Start() int                 { return (*LUPhrase)(n).Start() }
func (n *LURelative) End() int                   { return (*LUPhrase)(n).End() }
func (n LURelative) modNode(m *Mod) Node         { return n.modRelative(m) }
func (n LURelative) modFragment(m *Mod) Fragment { return n.modRelative(m) }
func (n LURelative) modRelative(m *Mod) Relative { return (*LURelative)(LUPhrase(n).mod(m)) }

// An CoPRelative is an argument connector phrase.
type CoPRelative CoP

func (n *CoPRelative) Start() int                 { return (*CoP)(n).Start() }
func (n *CoPRelative) End() int                   { return (*CoP)(n).End() }
func (n CoPRelative) modNode(m *Mod) Node         { return n.modRelative(m) }
func (n CoPRelative) modFragment(m *Mod) Fragment { return n.modRelative(m) }
func (n CoPRelative) modRelative(m *Mod) Relative { return (*CoPRelative)(CoP(n).mod(m)) }

//
// Adverb nodes.
//

// Adverb is implement by all adverb nodes.
type Adverb interface {
	Node
	modTerm(*Mod) Term
	modPhrase(*Mod) Phrase
	modAdverb(*Mod) Adverb
}

// An PredicateAdverb is a predicate functioning as an adverb.
type PredicateAdverb struct{ Predicate }

func (n *PredicateAdverb) modNode(m *Mod) Node     { return n.modAdverb(m) }
func (n *PredicateAdverb) modTerm(m *Mod) Term     { return n.modAdverb(m) }
func (n *PredicateAdverb) modPhrase(m *Mod) Phrase { return n.modAdverb(m) }

func (n PredicateAdverb) modAdverb(m *Mod) Adverb {
	n.Predicate = n.Predicate.modPredicate(m)
	return &n
}

// An CoPAdverb is an argument connector phrase.
type CoPAdverb CoP

func (n *CoPAdverb) Start() int             { return (*CoP)(n).Start() }
func (n *CoPAdverb) End() int               { return (*CoP)(n).End() }
func (n CoPAdverb) modNode(m *Mod) Node     { return n.modAdverb(m) }
func (n CoPAdverb) modTerm(m *Mod) Term     { return n.modAdverb(m) }
func (n CoPAdverb) modPhrase(m *Mod) Phrase { return n.modAdverb(m) }
func (n CoPAdverb) modAdverb(m *Mod) Adverb { return (*CoPAdverb)(CoP(n).mod(m)) }

//
// Preposition nodes.
//

// Preposition is implemented by all prepositional phrase nodes.
//
// Note that "preposition" alone typically refers to just the predicate, not the entire phrase.
// But in this case, the predicate is already just called a Predicate,
// and "PrepositionalPhrase" is too long (and "PP" is too short).
type Preposition interface {
	Node
	modTerm(*Mod) Term
	modPhrase(*Mod) Phrase
	modPreposition(*Mod) Preposition
}

// A PredicationPreposition is a prepositional predicate and an argument.
type PredicationPreposition struct {
	Predicate Predicate
	Argument  Argument
}

func (n *PredicationPreposition) Start() int             { return n.Predicate.Start() }
func (n *PredicationPreposition) End() int               { return n.Argument.End() }
func (n PredicationPreposition) modNode(m *Mod) Node     { return n.modPreposition(m) }
func (n PredicationPreposition) modTerm(m *Mod) Term     { return n.modPreposition(m) }
func (n PredicationPreposition) modPhrase(m *Mod) Phrase { return n.modPreposition(m) }

func (n PredicationPreposition) modPreposition(m *Mod) Preposition {
	n.Argument = n.Argument.modArgument(m)
	return &n
}

// An CoPPreposition is an argument connector phrase.
type CoPPreposition CoP

func (n *CoPPreposition) Start() int                       { return (*CoP)(n).Start() }
func (n *CoPPreposition) End() int                         { return (*CoP)(n).End() }
func (n CoPPreposition) modNode(m *Mod) Node               { return n.mod(m) }
func (n CoPPreposition) modTerm(m *Mod) Term               { return n.mod(m) }
func (n CoPPreposition) modPhrase(m *Mod) Phrase           { return n.mod(m) }
func (n CoPPreposition) modPreposition(m *Mod) Preposition { return n.mod(m) }
func (n CoPPreposition) mod(m *Mod) *CoPPreposition        { return (*CoPPreposition)(CoP(n).mod(m)) }

//
// Content nodes.
//

// Content is implemented by all content clause nodes.
type Content interface {
	Node
	modPredicate(*Mod) Predicate
	modPhrase(*Mod) Phrase
	modContent(*Mod) Content
}

// A PredicationContent is a statement content clause.
type PredicationContent struct{ Predication }

func (n PredicationContent) modNode(m *Mod) Node           { return n.mod(m) }
func (n PredicationContent) modPredicate(m *Mod) Predicate { return n.mod(m) }
func (n PredicationContent) modPhrase(m *Mod) Phrase       { return n.mod(m) }
func (n PredicationContent) modContent(m *Mod) Content     { return n.mod(m) }

func (n PredicationContent) mod(m *Mod) *PredicationContent {
	n.Predication = *n.Predication.mod(m)
	return &n
}

// A LUContent is a LU phrase adverb.
type LUContent LUPhrase

func (n *LUContent) Start() int                   { return (*LUPhrase)(n).Start() }
func (n *LUContent) End() int                     { return (*LUPhrase)(n).End() }
func (n LUContent) modNode(m *Mod) Node           { return n.modContent(m) }
func (n LUContent) modPredicate(m *Mod) Predicate { return n.modContent(m) }
func (n LUContent) modPhrase(m *Mod) Phrase       { return n.modContent(m) }
func (n LUContent) modContent(m *Mod) Content     { return (*LUContent)(LUPhrase(n).mod(m)) }

// An CoPContent is an argument connector phrase.
type CoPContent CoP

func (n *CoPContent) Start() int                   { return (*CoP)(n).Start() }
func (n *CoPContent) End() int                     { return (*CoP)(n).End() }
func (n CoPContent) modNode(m *Mod) Node           { return n.modContent(m) }
func (n CoPContent) modPredicate(m *Mod) Predicate { return n.modContent(m) }
func (n CoPContent) modPhrase(m *Mod) Phrase       { return n.modContent(m) }
func (n CoPContent) modContent(m *Mod) Content     { return (*CoPContent)(CoP(n).mod(m)) }

//
// Mod nodes
//

// Mod is implemented by all free modifier nodes.
type Mod interface {
	Node
	mod(*Mod) Mod
}

// A Parenthetical is a parenthetical  free modifier.
type Parenthetical struct {
	KIO, KI   Word
	Discourse Discourse
}

func (n *Parenthetical) Start() int         { return n.KIO.Start() }
func (n *Parenthetical) End() int           { return n.KI.End() }
func (n Parenthetical) modNode(m *Mod) Node { return n.mod(m) }

func (n Parenthetical) mod(m *Mod) Mod {
	n.Discourse = *n.Discourse.mod(m)
	return &n
}

// A Incidental is an incidental free modifier.
type Incidental struct {
	JU        Word
	Statement Statement
}

func (n *Incidental) Start() int         { return n.JU.Start() }
func (n *Incidental) End() int           { return n.Statement.End() }
func (n Incidental) modNode(m *Mod) Node { return n.mod(m) }

func (n Incidental) mod(m *Mod) Mod {
	n.Statement = n.Statement.modStatement(m)
	return &n
}

// A Vocative is a vocative free modifier.
type Vocative struct {
	HU       Word
	Argument Argument
}

func (n *Vocative) Start() int         { return n.HU.Start() }
func (n *Vocative) End() int           { return n.Argument.End() }
func (n Vocative) modNode(m *Mod) Node { return n.mod(m) }

func (n Vocative) mod(m *Mod) Mod {
	n.Argument = n.Argument.modArgument(m)
	return &n
}

// An Interjection is an interjection word.
type Interjection Word

func (n Interjection) PrettyPrint() string { return `Interjection("` + n.T + `"})` }
func (n *Interjection) Start() int         { return (*Word)(n).Start() }
func (n *Interjection) End() int           { return (*Word)(n).End() }
func (n Interjection) modNode(m *Mod) Node { return n.mod(m) }
func (n Interjection) mod(m *Mod) Mod      { return (*Interjection)(Word(n).mod(m)) }

// A Space is a whitespace-only word.
type Space Word

func (n Space) PrettyPrint() string { return `Space("` + n.T + `")` }
func (n *Space) Start() int         { return (*Word)(n).Start() }
func (n *Space) End() int           { return (*Word)(n).End() }
func (n Space) modNode(m *Mod) Node { return n.mod(m) }
func (n Space) mod(m *Mod) Mod      { return (*Space)(Word(n).mod(m)) }

//
// General nodes
//

// CoP is a connector phrase used for both forethought and afterthought
// and for all connectable node types.
type CoP struct {
	TO0, TO1    *Word
	RU          Word
	Left, Right Node
}

func (n *CoP) Start() int         { return start(n.TO0, n.Left) }
func (n *CoP) End() int           { return n.Right.End() }
func (n CoP) modNode(m *Mod) Node { return n.mod(m) }

func (n CoP) mod(m *Mod) *CoP {
	n.Right = n.Right.modNode(m)
	return &n
}

// A LUPhrase is a phrase beginning with word-class LU.
type LUPhrase struct {
	LU        Word
	Statement Statement
}

func (n *LUPhrase) Start() int         { return n.LU.Start() }
func (n *LUPhrase) End() int           { return n.Statement.End() }
func (n LUPhrase) modNode(m *Mod) Node { return n.mod(m) }

func (n LUPhrase) mod(m *Mod) *LUPhrase {
	n.Statement = n.Statement.modStatement(m)
	return &n
}

// A Word is a single word of text.
type Word struct {
	T    string
	S, E int
	M    Mod
}

func (n Word) PrettyPrint() string { return `"` + n.T + `"` }
func (n *Word) Start() int         { return n.S }
func (n *Word) End() int           { return n.E }

func (n Word) modNode(m *Mod) Node { return n.mod(m) }

func (n Word) mod(m *Mod) *Word {
	if m == nil {
		return &n
	}
	if n.M == nil {
		n.M = *m
	} else {
		n.M = n.M.mod(m)
	}
	return &n
}

func start(w *Word, n Node) int {
	if w != nil {
		return w.Start()
	}
	return n.Start()
}

func end(w *Word, n Node) int {
	if w != nil {
		return w.End()
	}
	return n.End()
}

// ToASCII converts all words of an AST to their standard ASCII form,
// with a tone marker following each syllable.
func ToASCII(node Node) {
	Visit(node, FuncVisitor(func(n Node) {
		if w, ok := n.(*Word); ok {
			w.T = tone.ToASCII(w.T)
		}
	}))
}
