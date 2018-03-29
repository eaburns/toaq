// Package ast contains an abstract syntax tree for Toaq.
package ast

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

	ModNode(*Mod) Node
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
func (n Text) ModNode(m *Mod) Node { n.Discourse = *n.Discourse.Mod(m); return &n }

// A Discourse is a non-empty sequence of sentences and/or fragments.
type Discourse []Node

func (n *Discourse) Start() int { return (*n)[0].Start() }
func (n *Discourse) End() int   { return (*n)[len(*n)-1].End() }

func (n Discourse) ModNode(m *Mod) Node { return n.Mod(m) }

func (n Discourse) Mod(m *Mod) *Discourse {
	l := len(n)
	c := make(Discourse, l)
	copy(c, n)
	if s, ok := c[l-1].(Sentence); ok {
		c[l-1] = s.ModSentence(m)
	} else {
		c[l-1] = c[l-1].(Fragment).ModFragment(m)
	}
	return &c
}

//
// Sentences.
//

// Sentence is implemented by all sentence nodes.
type Sentence interface {
	Node

	ModSentence(*Mod) Sentence
}

// A StatementSentence is a statement that is a sentence.
type StatementSentence struct {
	JE        *Word
	Statement Statement
	DA        *Word
}

func (n *StatementSentence) Start() int         { return start(n.JE, n.Statement) }
func (n *StatementSentence) End() int           { return end(n.DA, n.Statement) }
func (n StatementSentence) ModNode(m *Mod) Node { return n.ModSentence(m) }

func (n StatementSentence) ModSentence(m *Mod) Sentence {
	if n.DA != nil {
		n.DA = n.DA.Mod(m)
	} else {
		n.Statement = n.Statement.ModStatement(m)
	}
	return &n
}

// A CoPSentence is a connector phrase sententce.
type CoPSentence CoP

func (n *CoPSentence) Start() int                 { return (*CoP)(n).Start() }
func (n *CoPSentence) End() int                   { return (*CoP)(n).End() }
func (n CoPSentence) ModNode(m *Mod) Node         { return n.Mod(m) }
func (n CoPSentence) ModSentence(m *Mod) Sentence { return n.Mod(m) }
func (n CoPSentence) Mod(m *Mod) *CoPSentence     { return (*CoPSentence)(CoP(n).Mod(m)) }

//
// Fragments
//

// Fragment is implemented by all fragment nodes.
type Fragment interface {
	Node
	ModFragment(*Mod) Fragment
}

// A Prenex is a prenex.
type Prenex struct {
	Terms Terms
	BI    Word
}

func (n *Prenex) Start() int                 { return n.Terms.Start() }
func (n *Prenex) End() int                   { return n.BI.End() }
func (n Prenex) ModNode(m *Mod) Node         { return n.Mod(m) }
func (n Prenex) ModFragment(m *Mod) Fragment { return n.Mod(m) }

func (n Prenex) Mod(m *Mod) *Prenex {
	n.BI = *n.BI.Mod(m)
	return &n
}

//
// Statement nodes.
//

// Statement is implemented by all statement nodes.
type Statement interface {
	Node
	ModStatement(*Mod) Statement
}

// A Predication is a statement with a predicate and terms.
type Predication struct {
	Prenex    *Prenex
	Predicate Predicate
	Terms     *Terms
	NA        *Word
}

func (n *Predication) Start() int {
	if n.Prenex != nil {
		return n.Prenex.Start()
	}
	return n.Predicate.Start()
}

func (n *Predication) End() int                     { return end(n.NA, n.Terms) }
func (n Predication) ModNode(m *Mod) Node           { return n.Mod(m) }
func (n Predication) ModStatement(m *Mod) Statement { return n.Mod(m) }

func (n Predication) Mod(m *Mod) *Predication {
	if n.NA != nil {
		n.NA = n.NA.Mod(m)
	} else if n.Terms != nil {
		n.Terms = n.Terms.Mod(m)
	} else {
		n.Predicate = n.Predicate.ModPredicate(m)
	}
	return &n
}

// A CoPStatement is a connector phrase statement.
type CoPStatement CoP

func (n *CoPStatement) Start() int                   { return (*CoP)(n).Start() }
func (n *CoPStatement) End() int                     { return (*CoP)(n).End() }
func (n CoPStatement) ModNode(m *Mod) Node           { return n.ModStatement(m) }
func (n CoPStatement) ModStatement(m *Mod) Statement { return (*CoPStatement)(CoP(n).Mod(m)) }

//
// Predicate nodes
//

// The Predicate interface is implemented by all predicate nodes.
type Predicate interface {
	Node
	ModPhrase(*Mod) Phrase
	ModPredicate(*Mod) Predicate
}

// A PrefixedPredicate is a predicate with a prefix word.
type PrefixedPredicate struct {
	MU        Word
	Predicate Predicate
}

func (n *PrefixedPredicate) Start() int             { return n.MU.Start() }
func (n *PrefixedPredicate) End() int               { return n.Predicate.End() }
func (n PrefixedPredicate) ModNode(m *Mod) Node     { return n.ModPredicate(m) }
func (n PrefixedPredicate) ModPhrase(m *Mod) Phrase { return n.ModPredicate(m) }

func (n PrefixedPredicate) ModPredicate(m *Mod) Predicate {
	n.Predicate = n.Predicate.ModPredicate(m)
	return &n
}

// A SerialPredicate is a serial predicate phrase.
type SerialPredicate struct {
	Left, Right Predicate
}

func (n *SerialPredicate) Start() int             { return n.Left.Start() }
func (n *SerialPredicate) End() int               { return n.Right.End() }
func (n SerialPredicate) ModNode(m *Mod) Node     { return n.ModPredicate(m) }
func (n SerialPredicate) ModPhrase(m *Mod) Phrase { return n.ModPredicate(m) }

func (n SerialPredicate) ModPredicate(m *Mod) Predicate {
	n.Right = n.Right.ModPredicate(m)
	return &n
}

// A WordPredicate is a word that is a predicate.
type WordPredicate Word

func (n WordPredicate) PrettyPrint() string           { return `WordPredicate("` + n.T + `")` }
func (n *WordPredicate) Start() int                   { return (*Word)(n).Start() }
func (n *WordPredicate) End() int                     { return (*Word)(n).End() }
func (n WordPredicate) ModNode(m *Mod) Node           { return n.ModPredicate(m) }
func (n WordPredicate) ModPhrase(m *Mod) Phrase       { return n.ModPredicate(m) }
func (n WordPredicate) ModPredicate(m *Mod) Predicate { return (*WordPredicate)(Word(n).Mod(m)) }

// Phrase is implemented by all nodes
// that can be the complement of a MI phrase.
type Phrase interface {
	Node
	ModPhrase(*Mod) Phrase
}

// A MIPredicate is a predicate beginning with word-class MI.
type MIPredicate struct {
	MI     Word
	GA     *Word
	Phrase Phrase
}

func (n *MIPredicate) Start() int             { return n.MI.Start() }
func (n *MIPredicate) End() int               { return end(n.GA, n.Phrase) }
func (n MIPredicate) ModNode(m *Mod) Node     { return n.ModPredicate(m) }
func (n MIPredicate) ModPhrase(m *Mod) Phrase { return n.ModPredicate(m) }

func (n MIPredicate) ModPredicate(m *Mod) Predicate {
	if n.GA != nil {
		n.GA = n.GA.Mod(m)
	} else {
		n.Phrase = n.Phrase.ModPhrase(m)
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
func (n POPredicate) ModNode(m *Mod) Node     { return n.ModPredicate(m) }
func (n POPredicate) ModPhrase(m *Mod) Phrase { return n.ModPredicate(m) }

func (n POPredicate) ModPredicate(m *Mod) Predicate {
	if n.GA != nil {
		n.GA = n.GA.Mod(m)
	} else {
		n.Argument = n.Argument.ModArgument(m)
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
func (n MOPredicate) ModNode(m *Mod) Node           { return n.ModPredicate(m) }
func (n MOPredicate) ModPhrase(m *Mod) Phrase       { return n.ModPredicate(m) }
func (n MOPredicate) ModPredicate(m *Mod) Predicate { n.TEO = *n.TEO.Mod(m); return &n }

// A LUPredicate is a LU phrase predicate.
type LUPredicate LUPhrase

func (n *LUPredicate) Start() int                   { return (*LUPhrase)(n).Start() }
func (n *LUPredicate) End() int                     { return (*LUPhrase)(n).End() }
func (n LUPredicate) ModNode(m *Mod) Node           { return n.ModPredicate(m) }
func (n LUPredicate) ModPhrase(m *Mod) Phrase       { return n.ModPredicate(m) }
func (n LUPredicate) ModPredicate(m *Mod) Predicate { return (*LUPredicate)(LUPhrase(n).Mod(m)) }

// A CoPPredicate is a connector predicate phrase.
type CoPPredicate CoP

func (n *CoPPredicate) Start() int                   { return (*CoP)(n).Start() }
func (n *CoPPredicate) End() int                     { return (*CoP)(n).End() }
func (n CoPPredicate) ModNode(m *Mod) Node           { return n.Mod(m) }
func (n CoPPredicate) ModPhrase(m *Mod) Phrase       { return n.Mod(m) }
func (n CoPPredicate) ModPredicate(m *Mod) Predicate { return n.Mod(m) }
func (n CoPPredicate) Mod(m *Mod) *CoPPredicate      { return (*CoPPredicate)(CoP(n).Mod(m)) }

//
// Terms
//

// Term is implemented by all term nodes.
type Term interface {
	Node
	ModTerm(*Mod) Term
}

// A LinkedTerm is an argument preceeded by a linking word.
type LinkedTerm struct {
	GO       Word
	Argument Argument
}

func (n *LinkedTerm) Start() int         { return n.GO.Start() }
func (n *LinkedTerm) End() int           { return n.Argument.End() }
func (n LinkedTerm) ModNode(m *Mod) Node { return n.ModTerm(m) }

func (n LinkedTerm) ModTerm(m *Mod) Term {
	n.Argument = n.Argument.ModArgument(m)
	return &n
}

// A Terms is a linked list of Term nodes.
type Terms struct {
	Term  Term
	Terms *Terms
}

func (n *Terms) Start() int { return n.Term.Start() }

func (n *Terms) End() int {
	if n.Terms != nil {
		return n.Terms.End()
	}
	return n.Term.End()
}

func (n Terms) ModNode(m *Mod) Node         { return n.Mod(m) }
func (n Terms) ModFragment(m *Mod) Fragment { return n.Mod(m) }

func (n Terms) Mod(m *Mod) *Terms {
	if n.Terms != nil {
		n.Terms = n.Terms.Mod(m)
	} else {
		n.Term = n.Term.ModTerm(m)
	}
	return &n
}

// A TermSet is a connector of terms.
type TermSet CoP

func (n *TermSet) Start() int         { return (*CoP)(n).Start() }
func (n *TermSet) End() int           { return (*CoP)(n).End() }
func (n TermSet) ModNode(m *Mod) Node { return n.Mod(m) }
func (n TermSet) ModTerm(m *Mod) Term { return n.Mod(m) }
func (n TermSet) Mod(m *Mod) *TermSet { return (*TermSet)(CoP(n).Mod(m)) }

//
// Argument nodes.
//

// Argument is implemented by all argument nodes.
type Argument interface {
	Node
	ModTerm(*Mod) Term
	ModPhrase(*Mod) Phrase
	ModArgument(*Mod) Argument
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

func (n PredicateArgument) ModNode(m *Mod) Node         { return n.Mod(m) }
func (n PredicateArgument) ModTerm(m *Mod) Term         { return n.Mod(m) }
func (n PredicateArgument) ModPhrase(m *Mod) Phrase     { return n.Mod(m) }
func (n PredicateArgument) ModArgument(m *Mod) Argument { return n.Mod(m) }

func (n PredicateArgument) Mod(m *Mod) *PredicateArgument {
	if n.Relative != nil {
		n.Relative = n.Relative.ModRelative(m)
	} else {
		n.Predicate = n.Predicate.ModPredicate(m)
	}
	return &n
}

// An CoPArgument is an argument connector phrase.
type CoPArgument CoP

func (n *CoPArgument) Start() int                 { return (*CoP)(n).Start() }
func (n *CoPArgument) End() int                   { return (*CoP)(n).End() }
func (n CoPArgument) ModNode(m *Mod) Node         { return n.ModArgument(m) }
func (n CoPArgument) ModTerm(m *Mod) Term         { return n.ModArgument(m) }
func (n CoPArgument) ModPhrase(m *Mod) Phrase     { return n.ModArgument(m) }
func (n CoPArgument) ModArgument(m *Mod) Argument { return (*CoPArgument)(CoP(n).Mod(m)) }

//
// Relative clasue nodes.
//

// Relative is implemented by all relative clause nodes.
type Relative interface {
	Node
	ModFragment(*Mod) Fragment
	ModRelative(*Mod) Relative
}

// A PredicationRelative is a statement relative clause.
type PredicationRelative struct{ Predication }

func (n PredicationRelative) ModNode(m *Mod) Node         { return n.ModRelative(m) }
func (n PredicationRelative) ModFragment(m *Mod) Fragment { return n.ModRelative(m) }

func (n PredicationRelative) ModRelative(m *Mod) Relative {
	n.Predication = *n.Predication.Mod(m)
	return &n
}

// A LURelative is a LU phrase relative clause.
type LURelative LUPhrase

func (n *LURelative) Start() int                 { return (*LUPhrase)(n).Start() }
func (n *LURelative) End() int                   { return (*LUPhrase)(n).End() }
func (n LURelative) ModNode(m *Mod) Node         { return n.ModRelative(m) }
func (n LURelative) ModFragment(m *Mod) Fragment { return n.ModRelative(m) }
func (n LURelative) ModRelative(m *Mod) Relative { return (*LURelative)(LUPhrase(n).Mod(m)) }

// An CoPRelative is an argument connector phrase.
type CoPRelative CoP

func (n *CoPRelative) Start() int                 { return (*CoP)(n).Start() }
func (n *CoPRelative) End() int                   { return (*CoP)(n).End() }
func (n CoPRelative) ModNode(m *Mod) Node         { return n.ModRelative(m) }
func (n CoPRelative) ModFragment(m *Mod) Fragment { return n.ModRelative(m) }
func (n CoPRelative) ModRelative(m *Mod) Relative { return (*CoPRelative)(CoP(n).Mod(m)) }

//
// Adverb nodes.
//

// Adverb is implement by all adverb nodes.
type Adverb interface {
	Node
	ModTerm(*Mod) Term
	ModPhrase(*Mod) Phrase
	ModAdverb(*Mod) Adverb
}

// An PredicateAdverb is a predicate functioning as an adverb.
type PredicateAdverb struct{ Predicate }

func (n *PredicateAdverb) ModNode(m *Mod) Node     { return n.ModAdverb(m) }
func (n *PredicateAdverb) ModTerm(m *Mod) Term     { return n.ModAdverb(m) }
func (n *PredicateAdverb) ModPhrase(m *Mod) Phrase { return n.ModAdverb(m) }

func (n PredicateAdverb) ModAdverb(m *Mod) Adverb {
	n.Predicate = n.Predicate.ModPredicate(m)
	return &n
}

// An CoPAdverb is an argument connector phrase.
type CoPAdverb CoP

func (n *CoPAdverb) Start() int             { return (*CoP)(n).Start() }
func (n *CoPAdverb) End() int               { return (*CoP)(n).End() }
func (n CoPAdverb) ModNode(m *Mod) Node     { return n.ModAdverb(m) }
func (n CoPAdverb) ModTerm(m *Mod) Term     { return n.ModAdverb(m) }
func (n CoPAdverb) ModPhrase(m *Mod) Phrase { return n.ModAdverb(m) }
func (n CoPAdverb) ModAdverb(m *Mod) Adverb { return (*CoPAdverb)(CoP(n).Mod(m)) }

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
	ModTerm(*Mod) Term
	ModPhrase(*Mod) Phrase
	ModPreposition(*Mod) Preposition
}

// A PredicationPreposition is a prepositional predicate and an argument.
type PredicationPreposition struct {
	Predicate Predicate
	Argument  Argument
}

func (n *PredicationPreposition) Start() int             { return n.Predicate.Start() }
func (n *PredicationPreposition) End() int               { return n.Argument.End() }
func (n PredicationPreposition) ModNode(m *Mod) Node     { return n.ModPreposition(m) }
func (n PredicationPreposition) ModTerm(m *Mod) Term     { return n.ModPreposition(m) }
func (n PredicationPreposition) ModPhrase(m *Mod) Phrase { return n.ModPreposition(m) }

func (n PredicationPreposition) ModPreposition(m *Mod) Preposition {
	n.Argument = n.Argument.ModArgument(m)
	return &n
}

// An CoPPreposition is an argument connector phrase.
type CoPPreposition CoP

func (n *CoPPreposition) Start() int                       { return (*CoP)(n).Start() }
func (n *CoPPreposition) End() int                         { return (*CoP)(n).End() }
func (n CoPPreposition) ModNode(m *Mod) Node               { return n.Mod(m) }
func (n CoPPreposition) ModTerm(m *Mod) Term               { return n.Mod(m) }
func (n CoPPreposition) ModPhrase(m *Mod) Phrase           { return n.Mod(m) }
func (n CoPPreposition) ModPreposition(m *Mod) Preposition { return n.Mod(m) }
func (n CoPPreposition) Mod(m *Mod) *CoPPreposition        { return (*CoPPreposition)(CoP(n).Mod(m)) }

//
// Content nodes.
//

// Content is implemented by all content clause nodes.
type Content interface {
	Node
	ModPredicate(*Mod) Predicate
	ModPhrase(*Mod) Phrase
	ModContent(*Mod) Content
}

// A PredicationContent is a statement content clause.
type PredicationContent struct{ Predication }

func (n PredicationContent) ModNode(m *Mod) Node           { return n.Mod(m) }
func (n PredicationContent) ModPredicate(m *Mod) Predicate { return n.Mod(m) }
func (n PredicationContent) ModPhrase(m *Mod) Phrase       { return n.Mod(m) }
func (n PredicationContent) ModContent(m *Mod) Content     { return n.Mod(m) }

func (n PredicationContent) Mod(m *Mod) *PredicationContent {
	n.Predication = *n.Predication.Mod(m)
	return &n
}

// A LUContent is a LU phrase adverb.
type LUContent LUPhrase

func (n *LUContent) Start() int                   { return (*LUPhrase)(n).Start() }
func (n *LUContent) End() int                     { return (*LUPhrase)(n).End() }
func (n LUContent) ModNode(m *Mod) Node           { return n.ModContent(m) }
func (n LUContent) ModPredicate(m *Mod) Predicate { return n.ModContent(m) }
func (n LUContent) ModPhrase(m *Mod) Phrase       { return n.ModContent(m) }
func (n LUContent) ModContent(m *Mod) Content     { return (*LUContent)(LUPhrase(n).Mod(m)) }

// An CoPContent is an argument connector phrase.
type CoPContent CoP

func (n *CoPContent) Start() int                   { return (*CoP)(n).Start() }
func (n *CoPContent) End() int                     { return (*CoP)(n).End() }
func (n CoPContent) ModNode(m *Mod) Node           { return n.ModContent(m) }
func (n CoPContent) ModPredicate(m *Mod) Predicate { return n.ModContent(m) }
func (n CoPContent) ModPhrase(m *Mod) Phrase       { return n.ModContent(m) }
func (n CoPContent) ModContent(m *Mod) Content     { return (*CoPContent)(CoP(n).Mod(m)) }

//
// Mod nodes
//

// Mod is implemented by all free modifier nodes.
type Mod interface {
	Node
	Mod(*Mod) Mod
}

// A Parenthetical is a parenthetical  free modifier.
type Parenthetical struct {
	KIO, KI   Word
	Discourse Discourse
}

func (n *Parenthetical) Start() int         { return n.KIO.Start() }
func (n *Parenthetical) End() int           { return n.KI.End() }
func (n Parenthetical) ModNode(m *Mod) Node { return n.Mod(m) }

func (n Parenthetical) Mod(m *Mod) Mod {
	n.Discourse = *n.Discourse.Mod(m)
	return &n
}

// A Incidental is an incidental free modifier.
type Incidental struct {
	JU        Word
	Statement Statement
}

func (n *Incidental) Start() int         { return n.JU.Start() }
func (n *Incidental) End() int           { return n.Statement.End() }
func (n Incidental) ModNode(m *Mod) Node { return n.Mod(m) }

func (n Incidental) Mod(m *Mod) Mod {
	n.Statement = n.Statement.ModStatement(m)
	return &n
}

// A Vocative is a vocative free modifier.
type Vocative struct {
	HU       Word
	Argument Argument
}

func (n *Vocative) Start() int         { return n.HU.Start() }
func (n *Vocative) End() int           { return n.Argument.End() }
func (n Vocative) ModNode(m *Mod) Node { return n.Mod(m) }

func (n Vocative) Mod(m *Mod) Mod {
	n.Argument = n.Argument.ModArgument(m)
	return &n
}

// An Interjection is an interjection word.
type Interjection Word

func (n Interjection) PrettyPrint() string { return `Interjection("` + n.T + `"})` }
func (n *Interjection) Start() int         { return (*Word)(n).Start() }
func (n *Interjection) End() int           { return (*Word)(n).End() }
func (n Interjection) ModNode(m *Mod) Node { return n.Mod(m) }
func (n Interjection) Mod(m *Mod) Mod      { return (*Interjection)(Word(n).Mod(m)) }

// A Space is a whitespace-only word.
type Space Word

func (n Space) PrettyPrint() string { return `Space("` + n.T + `")` }
func (n *Space) Start() int         { return (*Word)(n).Start() }
func (n *Space) End() int           { return (*Word)(n).End() }
func (n Space) ModNode(m *Mod) Node { return n.Mod(m) }
func (n Space) Mod(m *Mod) Mod      { return (*Space)(Word(n).Mod(m)) }

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

func (n *CoP) Start() int { return start(n.TO0, n.Left) }
func (n *CoP) End() int   { return n.Right.End() }

func (n CoP) Mod(m *Mod) *CoP {
	n.Right = n.Right.ModNode(m)
	return &n
}

// A LUPhrase is a phrase beginning with word-class LU.
type LUPhrase struct {
	LU        Word
	Statement Statement
}

func (n *LUPhrase) Start() int { return n.LU.Start() }
func (n *LUPhrase) End() int   { return n.Statement.End() }

func (n LUPhrase) Mod(m *Mod) *LUPhrase {
	n.Statement = n.Statement.ModStatement(m)
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

func (n Word) ModNode(m *Mod) Node { return n.Mod(m) }

func (n Word) Mod(m *Mod) *Word {
	if m == nil {
		return &n
	}
	if n.M == nil {
		n.M = *m
	} else {
		n.M = n.M.Mod(m)
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
