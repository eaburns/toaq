package ast

import "fmt"

// Equal returns whether the sub-trees rooted at two nodes are equal.
//
// Note that two Words are considered equal when their T fields are ==,
// and when their M fields are Equal(); the S and E fields are ignored.
func Equal(a, b Node) bool {
	if a == nil {
		return b == nil
	}
	switch a := a.(type) {
	case *Text:
		b, ok := b.(*Text)
		return ok &&
			Equal(a.Leading, b.Leading) &&
			Equal(&a.Discourse, &b.Discourse)

	case *Discourse:
		b, ok := b.(*Discourse)
		if !ok || len(*a) != len(*b) {
			return false
		}
		for i := range *a {
			if !Equal((*a)[i], (*b)[i]) {
				return false
			}
		}
		return true

	case *StatementSentence:
		b, ok := b.(*StatementSentence)
		return ok &&
			wordEqual(a.JE, b.JE) &&
			wordEqual(a.DA, b.DA) &&
			Equal(a.Statement, b.Statement)

	case *CoPSentence:
		b, ok := b.(*CoPSentence)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *Prenex:
		b, ok := b.(*Prenex)
		return ok && wordEqual(&a.BI, &b.BI) && Equal(&a.Terms, &b.Terms)

	case *Predication:
		b, ok := b.(*Predication)
		return ok &&
			wordEqual(a.NA, b.NA) &&
			Equal(a.Predicate, b.Predicate) &&
			(a.Terms == nil) == (b.Terms == nil) &&
			(a.Terms == nil || Equal(a.Terms, b.Terms)) &&
			(a.Prenex == nil) == (b.Prenex == nil) &&
			(a.Prenex == nil || Equal(a.Prenex, b.Prenex))

	case *CoPStatement:
		b, ok := b.(*CoPStatement)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *PrefixedPredicate:
		b, ok := b.(*PrefixedPredicate)
		return ok && wordEqual(&a.MU, &b.MU) && Equal(a.Predicate, b.Predicate)

	case *SerialPredicate:
		b, ok := b.(*SerialPredicate)
		return ok && Equal(a.Left, b.Left) && Equal(a.Right, b.Right)

	case *WordPredicate:
		b, ok := b.(*WordPredicate)
		return ok && wordEqual((*Word)(a), (*Word)(b))

	case *MIPredicate:
		b, ok := b.(*MIPredicate)
		return ok && wordEqual(&a.MI, &b.MI) && wordEqual(a.GA, b.GA) && Equal(a.Phrase, b.Phrase)

	case *POPredicate:
		b, ok := b.(*POPredicate)
		return ok && wordEqual(&a.PO, &b.PO) && wordEqual(a.GA, b.GA) && Equal(a.Argument, b.Argument)

	case *MOPredicate:
		b, ok := b.(*MOPredicate)
		return ok && wordEqual(&a.MO, &b.MO) && wordEqual(&a.TEO, &b.TEO) && Equal(&a.Discourse, &b.Discourse)

	case *LUPredicate:
		b, ok := b.(*LUPredicate)
		return ok && wordEqual(&a.LU, &b.LU) && Equal(a.Statement, b.Statement)

	case *CoPPredicate:
		b, ok := b.(*CoPPredicate)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *LinkedTerm:
		b, ok := b.(*LinkedTerm)
		return ok && wordEqual(&a.GO, &b.GO) && Equal(a.Argument, b.Argument)

	case *Terms:
		b, ok := b.(*Terms)
		return ok && Equal(a.Term, b.Term) &&
			(a.Terms == nil) == (b.Terms == nil) &&
			(a.Terms == nil || Equal(a.Terms, b.Terms))

	case *TermSet:
		b, ok := b.(*TermSet)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *PredicateArgument:
		b, ok := b.(*PredicateArgument)
		return ok &&
			wordEqual(a.Focus, b.Focus) &&
			wordEqual(a.Quantifier, b.Quantifier) &&
			Equal(a.Predicate, b.Predicate) &&
			(a.Relative == nil) == (b.Relative == nil) &&
			(a.Relative == nil || Equal(a.Relative, b.Relative))

	case *CoPArgument:
		b, ok := b.(*CoPArgument)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *PredicationRelative:
		b, ok := b.(*PredicationRelative)
		return ok && Equal(&a.Predication, &b.Predication)

	case *LURelative:
		b, ok := b.(*LURelative)
		return ok && wordEqual(&a.LU, &b.LU) && Equal(a.Statement, b.Statement)

	case *CoPRelative:
		b, ok := b.(*CoPRelative)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *PredicateAdverb:
		b, ok := b.(*PredicateAdverb)
		return ok && Equal(a.Predicate, b.Predicate)

	case *CoPAdverb:
		b, ok := b.(*CoPAdverb)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *PredicationPreposition:
		b, ok := b.(*PredicationPreposition)
		return ok && Equal(a.Predicate, b.Predicate) && Equal(a.Argument, b.Argument)

	case *CoPPreposition:
		b, ok := b.(*CoPPreposition)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *PredicationContent:
		b, ok := b.(*PredicationContent)
		return ok && Equal(&a.Predication, &b.Predication)

	case *LUContent:
		b, ok := b.(*LUContent)
		return ok && wordEqual(&a.LU, &b.LU) && Equal(a.Statement, b.Statement)

	case *CoPContent:
		b, ok := b.(*CoPContent)
		return ok && copEqual((*CoP)(a), (*CoP)(b))

	case *Parenthetical:
		b, ok := b.(*Parenthetical)
		return ok && wordEqual(&a.KI, &b.KI) && wordEqual(&a.KIO, &b.KIO) && Equal(&a.Discourse, &b.Discourse)

	case *Incidental:
		b, ok := b.(*Incidental)
		return ok && wordEqual(&a.JU, &b.JU) && Equal(a.Statement, b.Statement)

	case *Vocative:
		b, ok := b.(*Vocative)
		return ok && wordEqual(&a.HU, &b.HU) && Equal(a.Argument, b.Argument)

	case *Interjection:
		b, ok := b.(*Interjection)
		return ok && wordEqual((*Word)(a), (*Word)(b))

	case *Space:
		b, ok := b.(*Space)
		return ok && wordEqual((*Word)(a), (*Word)(b))

	default:
		panic(fmt.Sprintf("unknown node type %T", a))
	}
}

func copEqual(a, b *CoP) bool {
	return wordEqual(a.TO0, b.TO0) &&
		wordEqual(a.TO1, b.TO1) &&
		wordEqual(&a.RU, &b.RU) &&
		Equal(a.Left, b.Left) &&
		Equal(a.Right, b.Right)
}

func wordEqual(a, b *Word) bool {
	if a == nil {
		return b == nil
	}
	return b != nil && a.T == b.T && Equal(a.M, b.M)
}
