package ast

import (
	"unicode/utf8"

	"github.com/eaburns/toaq/tone"
)

func text(s0 *Mod, d *[]Node) Text {
	var txt Text
	if s0 != nil {
		txt.Leading = *s0
	}
	if d != nil {
		txt.Discourse = *d
	}
	return txt
}

func prenexStmt(p *Prenex, st Statement) Statement {
	if p == nil {
		return st
	}
	return &PrenexStatement{Prenex: *p, Statement: st}
}

func prenex(ts Terms, s *Mod, bi Word) *Prenex {
	return &Prenex{Terms: ts.mod(s), BI: bi}
}

func predication(p Predicate, s *Mod, ts *Terms) Predication {
	st := Predication{Predicate: p.modPredicate(s)}
	if ts != nil {
		st.Terms = *ts
	}
	return st
}

func linkedTerm(g **Word, a Argument) Term {
	if g == nil {
		return a
	}
	return &LinkedTerm{GO: **g, Argument: a}
}

func focusedArgument(ku *Word, a PredicateArgument) *PredicateArgument {
	a.Focus = ku
	return &a
}

func quantifiedArgument(tu *Word, a PredicateArgument) PredicateArgument {
	a.Quantifier = tu
	return a
}

func argumentRelative(a PredicateArgument, s *Mod, rc *Relative) PredicateArgument {
	a = *a.mod(s)
	if rc != nil {
		a.Relative = *rc
	}
	return a
}

func prepPhrase(p Predicate, s *Mod, a Argument) *PredicationPreposition {
	return &PredicationPreposition{Predicate: p.modPredicate(s), Argument: a}
}

// endPredication terminates a statement.
func endPredication(st Predication, s *Mod, na *Word) *Predication {
	st = *st.mod(s)
	st.NA = na
	return &st
}

func serial(left Predicate, s *Mod, right Predicate) *SerialPredicate {
	return &SerialPredicate{Left: left.modPredicate(s), Right: right}
}

func copLeft(l Node, s *Mod, bar CoP) *CoP {
	bar.Left = l.modNode(s)
	return &bar
}

func copRight(ru Word, s *Mod, r Node) CoP {
	return CoP{RU: *ru.mod(s), Right: r}
}

func foreCoP(con CoP, s *Mod, cop CoP) *CoP {
	cop.RU = *con.RU.mod(s)
	cop.TO0 = con.TO0
	return &cop
}

func foreCoPLeft(l Node, s *Mod, bar CoP) CoP {
	bar.Left = l.modNode(s)
	return bar
}

func foreCoPRight(to Word, s *Mod, r Node) CoP {
	return CoP{TO1: to.mod(s), Right: r}
}

func connector(to Word, s *Mod, ru Word) CoP {
	return CoP{TO0: to.mod(s), RU: ru}
}

func miEnd(mi MIPredicate, s *Mod, ga *Word) *MIPredicate {
	mi.Phrase = mi.Phrase.modPhrase(s)
	mi.GA = ga
	return &mi
}

func poEnd(po POPredicate, s *Mod, ga *Word) *POPredicate {
	po.Argument = po.Argument.modArgument(s)
	po.GA = ga
	return &po
}

func moEnd(mo MOPredicate, s *Mod, teo Word) *MOPredicate {
	mo.Discourse = *mo.Discourse.mod(s)
	mo.TEO = teo
	return &mo
}

func endParen(p Parenthetical, s *Mod, ki Word) *Parenthetical {
	p.KI = ki
	p.Discourse = *p.Discourse.mod(s)
	return &p
}

// addTone swaps the first letter of des with its diacritic variant given the tone rune.
func addTone(des, ton string) string {
	t, _ := utf8.DecodeRuneInString(ton)
	_, w := utf8.DecodeRuneInString(des)
	return tone.Diacritic[t][des[:w]][0] + des[w:]
}
