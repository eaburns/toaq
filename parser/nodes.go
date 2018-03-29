package parser

import "github.com/eaburns/toaq/ast"

func text(s0 *ast.Mod, d *[]ast.Node) ast.Text {
	var txt ast.Text
	if s0 != nil {
		txt.Leading = *s0
	}
	if d != nil {
		txt.Discourse = *d
	}
	return txt
}

func prenexStmt(n *ast.Prenex, p ast.Predication) ast.Predication {
	p.Prenex = n
	return p
}

func prenex(ts *ast.Terms, s *ast.Mod, bi ast.Word) *ast.Prenex {
	return &ast.Prenex{Terms: *ts.Mod(s), BI: bi}
}

func predication(p ast.Predicate, s *ast.Mod, ts **ast.Terms) ast.Predication {
	st := ast.Predication{Predicate: p.ModPredicate(s)}
	if ts != nil {
		st.Terms = *ts
	}
	return st
}

func linkedTerm(g **ast.Word, a ast.Argument) ast.Term {
	if g == nil {
		return a
	}
	return &ast.LinkedTerm{GO: **g, Argument: a}
}

func focusedArgument(ku *ast.Word, a ast.PredicateArgument) *ast.PredicateArgument {
	a.Focus = ku
	return &a
}

func quantifiedArgument(tu *ast.Word, a ast.PredicateArgument) ast.PredicateArgument {
	a.Quantifier = tu
	return a
}

func argumentRelative(a ast.PredicateArgument, s *ast.Mod, rc *ast.Relative) ast.PredicateArgument {
	a = *a.Mod(s)
	if rc != nil {
		a.Relative = *rc
	}
	return a
}

func prepPhrase(p ast.Predicate, s *ast.Mod, a ast.Argument) *ast.PredicationPreposition {
	return &ast.PredicationPreposition{Predicate: p.ModPredicate(s), Argument: a}
}

// endPredication terminates a statement.
func endPredication(st ast.Predication, s *ast.Mod, na *ast.Word) *ast.Predication {
	st = *st.Mod(s)
	st.NA = na
	return &st
}

func prefixed(mu ast.Word, s *ast.Mod, p ast.Predicate) *ast.PrefixedPredicate {
	return &ast.PrefixedPredicate{MU: *mu.Mod(s), Predicate: p}
}

func serial(left ast.Predicate, s *ast.Mod, right ast.Predicate) *ast.SerialPredicate {
	return &ast.SerialPredicate{Left: left.ModPredicate(s), Right: right}
}

func copLeft(l ast.Node, s *ast.Mod, bar ast.CoP) *ast.CoP {
	bar.Left = l.ModNode(s)
	return &bar
}

func copRight(ru ast.Word, s *ast.Mod, r ast.Node) ast.CoP {
	return ast.CoP{RU: *ru.Mod(s), Right: r}
}

func foreCoP(con ast.CoP, s *ast.Mod, cop ast.CoP) *ast.CoP {
	cop.RU = *con.RU.Mod(s)
	cop.TO0 = con.TO0
	return &cop
}

func foreCoPLeft(l ast.Node, s *ast.Mod, bar ast.CoP) ast.CoP {
	bar.Left = l.ModNode(s)
	return bar
}

func foreCoPRight(to ast.Word, s *ast.Mod, r ast.Node) ast.CoP {
	return ast.CoP{TO1: to.Mod(s), Right: r}
}

func connector(to ast.Word, s *ast.Mod, ru ast.Word) ast.CoP {
	return ast.CoP{TO0: to.Mod(s), RU: ru}
}

func miEnd(mi ast.MIPredicate, s *ast.Mod, ga *ast.Word) *ast.MIPredicate {
	mi.Phrase = mi.Phrase.ModPhrase(s)
	mi.GA = ga
	return &mi
}

func poEnd(po ast.POPredicate, s *ast.Mod, ga *ast.Word) *ast.POPredicate {
	po.Argument = po.Argument.ModArgument(s)
	po.GA = ga
	return &po
}

func moEnd(mo ast.MOPredicate, s *ast.Mod, teo ast.Word) *ast.MOPredicate {
	mo.Discourse = *mo.Discourse.Mod(s)
	mo.TEO = teo
	return &mo
}

func endParen(p ast.Parenthetical, s *ast.Mod, ki ast.Word) *ast.Parenthetical {
	p.KI = ki
	p.Discourse = *p.Discourse.Mod(s)
	return &p
}
