// Package ast contains an abstract syntax tree for Toaq.
package ast

import "strings"

// A Node is any node in the AST.
type Node interface {
	// Start returns the byte-offset into the text of the start of the node.
	Start() int
	// End returns the exclusive byte-offset into the text of the end of the node.
	End() int

	// Mod returns a copy of the reciver with a free modifier appended.
	// If the argument is nil, nothing is appended
	// and the receiver may be returned without copying.
	Mod(*Node) Node

	// buildString appends the string representation
	// of the subtree rooted at the receiver
	// to the string builder.
	buildString(*strings.Builder)
}

// A Word is a single word of text.
type Word struct {
	T    string
	S, E int
	M    Node
}

func (n Word) String() string {
	var s strings.Builder
	n.buildString(&s)
	return s.String()
}

func (n Word) buildString(s *strings.Builder) {
	s.WriteString(n.T)
	if n.M != nil {
		n.M.buildString(s)
	}
}

// Start returns the byte-offset of the start of the word.
func (n Word) Start() int { return n.S }

// End retturns the (exclusive) byte-offset of the end of the word.
func (n Word) End() int { return n.E }

func (n Word) Mod(m *Node) Node {
	if m == nil {
		return n
	}
	if n.M == nil {
		n.M = *m
	} else {
		n.M = n.M.Mod(m)
	}
	return n
}
