package ast

import (
	"strings"
	"testing"
)

func TestEq(t *testing.T) {
	tests := []struct {
		n, eq Node
		ne    []Node
	}{
		{
			n: &Text{},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Text{Leading: &Space{T: " "}},
				&Text{Discourse: Discourse{(*Interjection)(&Word{T: "m"})}},
			},
		},
		{
			n: &Text{Leading: &Space{T: " "}},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Text{},
				&Text{Discourse: Discourse{(*Interjection)(&Word{T: "m"})}},
			},
		},
		{
			n: &StatementSentence{
				Statement: &Predication{Predicate: p("hio?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&StatementSentence{
					Statement: &Predication{Predicate: p("mai?")},
				},
				&StatementSentence{
					Statement: &CoPStatement{
						Left:  &Predication{Predicate: p("mai?")},
						Right: &Predication{Predicate: p("pai?")},
					},
				},
				&StatementSentence{
					JE:        &Word{T: "je"},
					Statement: &Predication{Predicate: p("hio?")},
				},
				&StatementSentence{
					DA:        &Word{T: "da"},
					Statement: &Predication{Predicate: p("hio?")},
				},
			},
		},
		{
			n: &CoPSentence{
				TO0: &Word{T: "to"},
				TO1: &Word{T: "to"},
				RU:  Word{T: "ru"},
				Left: &StatementSentence{
					Statement: &Predication{Predicate: p("mai?")},
				},
				Right: &StatementSentence{
					Statement: &Predication{Predicate: p("pai?")},
				},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&StatementSentence{
					Statement: &Predication{Predicate: p("mai?")},
				},
				&StatementSentence{
					Statement: &CoPStatement{
						Left:  &Predication{Predicate: p("mai?")},
						Right: &Predication{Predicate: p("pai?")},
					},
				},
				&CoPSentence{
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &StatementSentence{
						Statement: &Predication{Predicate: p("mai?")},
					},
					Right: &StatementSentence{
						Statement: &Predication{Predicate: p("pai?")},
					},
				},
				&CoPSentence{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &StatementSentence{
						Statement: &Predication{Predicate: p("mai?")},
					},
					Right: &StatementSentence{
						Statement: &Predication{Predicate: p("pai?")},
					},
				},
				&CoPSentence{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ri"},
					Left: &StatementSentence{
						Statement: &Predication{Predicate: p("mai?")},
					},
					Right: &StatementSentence{
						Statement: &Predication{Predicate: p("pai?")},
					},
				},
				&CoPSentence{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Right: &StatementSentence{
						Statement: &Predication{Predicate: p("de?")},
					},
					Left: &StatementSentence{
						Statement: &Predication{Predicate: p("pai?")},
					},
				},
				&CoPSentence{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Right: &StatementSentence{
						Statement: &Predication{Predicate: p("mai?")},
					},
					Left: &StatementSentence{
						Statement: &Predication{Predicate: p("de?")},
					},
				},
				&CoPSentence{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Right: &StatementSentence{
						Statement: &Predication{Predicate: p("mai?")},
					},
					Left: &StatementSentence{
						Statement: &Predication{Predicate: p("pai?")},
					},
				},
			},
		},
		{
			n: &Prenex{
				Terms: Terms{&PredicateArgument{Predicate: p("hio/")}},
				BI:    Word{T: "bi"},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Prenex{
					Terms: Terms{&PredicateArgument{Predicate: p("jai/")}},
					BI:    Word{T: "bi"},
				},
				&Prenex{
					Terms: Terms{&PredicateArgument{Predicate: p("hio/")}},
					BI:    Word{T: "pa"},
				},
			},
		},
		{
			n: &PrenexStatement{
				Prenex: Prenex{
					Terms: Terms{&PredicateArgument{Predicate: p("hio/")}},
					BI:    Word{T: "bi"},
				},
				Statement: &Predication{Predicate: p("mai/")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&PrenexStatement{
					Prenex: Prenex{
						Terms: Terms{&PredicateArgument{Predicate: p("hio/")}},
						BI:    Word{T: "pa"},
					},
					Statement: &Predication{Predicate: p("mai/")},
				},
				&PrenexStatement{
					Prenex: Prenex{
						Terms: Terms{&PredicateArgument{Predicate: p("hio/")}},
						BI:    Word{T: "bi"},
					},
					Statement: &Predication{Predicate: p("pai/")},
				},
			},
		},
		{
			n: &Predication{
				Predicate: p("mai?"),
				Terms:     Terms{&PredicateArgument{Predicate: p("ji/")}},
				NA:        &Word{T: "na"},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Predication{
					Predicate: p("pai?"),
					Terms:     Terms{&PredicateArgument{Predicate: p("ji/")}},
					NA:        &Word{T: "na"},
				},
				&Predication{
					Predicate: p("mai?"),
					Terms:     Terms{&PredicateArgument{Predicate: p("suq/")}},
					NA:        &Word{T: "na"},
				},
				&Predication{
					Predicate: p("mai?"),
					NA:        &Word{T: "na"},
				},
				&Predication{
					Predicate: p("mai?"),
					Terms:     Terms{&PredicateArgument{Predicate: p("suq/")}},
				},
			},
		},
		{
			n: &CoPStatement{
				TO0:   &Word{T: "to"},
				TO1:   &Word{T: "to"},
				RU:    Word{T: "ru"},
				Left:  &Predication{Predicate: p("mai?")},
				Right: &Predication{Predicate: p("pai?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&CoPStatement{
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPStatement{
					TO0:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPStatement{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ri"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPStatement{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("mea?heo-")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPStatement{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("mea?heo-")},
				},
				&CoPStatement{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Right: &Predication{Predicate: p("mai?")},
					Left:  &Predication{Predicate: p("pai?")},
				},
			},
		},
		{
			n: &SerialPredicate{Left: p("mai?"), Right: p("pai?")},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&SerialPredicate{Left: p("mea?heo-"), Right: p("pai?")},
				&SerialPredicate{Left: p("mai?"), Right: p("mea?heo-")},
				&SerialPredicate{Right: p("mai?"), Left: p("pai?")},
			},
		},
		{
			n: (*WordPredicate)(&Word{T: "pai?"}),
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				(*WordPredicate)(&Word{T: "mai?"}),
			},
		},
		{
			n: &MIPredicate{
				MI:     Word{T: "mi/"},
				Phrase: p("Toaq?"),
				GA:     &Word{T: "ga"},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				(*WordPredicate)(&Word{T: "mai?"}),
				&MIPredicate{
					MI:     Word{T: "mi/"},
					Phrase: p("Hoaq?"),
					GA:     &Word{T: "ga"},
				},
				&MIPredicate{
					MI:     Word{T: "mi/"},
					Phrase: p("Toaq?"),
				},
			},
		},
		{
			n: &POPredicate{
				PO:       Word{T: "po/"},
				Argument: &PredicateArgument{Predicate: p("ji/")},
				GA:       &Word{T: "ga"},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				(*WordPredicate)(&Word{T: "mai?"}),
				&POPredicate{
					PO:       Word{T: "po/"},
					Argument: &PredicateArgument{Predicate: p("suq/")},
					GA:       &Word{T: "ga"},
				},
				&POPredicate{
					PO:       Word{T: "po/"},
					Argument: &PredicateArgument{Predicate: p("ji/")},
				},
			},
		},
		{
			n: &MOPredicate{
				MO:        Word{T: "mo/"},
				Discourse: Discourse{(*Interjection)(&Word{T: "m"})},
				TEO:       Word{T: "teo"},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&MOPredicate{
					MO:        Word{T: "momo/"},
					Discourse: Discourse{(*Interjection)(&Word{T: "m"})},
					TEO:       Word{T: "teo"},
				},
				&MOPredicate{
					MO:        Word{T: "mo/"},
					Discourse: Discourse{(*Interjection)(&Word{T: "HA"})},
					TEO:       Word{T: "teo"},
				},
				&MOPredicate{
					MO:        Word{T: "mo/"},
					Discourse: Discourse{(*Interjection)(&Word{T: "m"})},
					TEO:       Word{T: "teoteo-"},
				},
			},
		},
		{
			n: &LUPredicate{
				LU:        Word{T: "lu/"},
				Statement: &Predication{Predicate: p("hio?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&LUPredicate{
					LU:        Word{T: "lu/lu-"},
					Statement: &Predication{Predicate: p("hio?")},
				},
				&LUPredicate{
					LU:        Word{T: "lu/"},
					Statement: &Predication{Predicate: p("mai?")},
				},
			},
		},
		{
			n: &CoPPredicate{
				TO0:   &Word{T: "to"},
				TO1:   &Word{T: "to"},
				RU:    Word{T: "ru"},
				Left:  p("mai?"),
				Right: p("pai?"),
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&CoPPredicate{
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  p("mai?"),
					Right: p("pai?"),
				},
				&CoPPredicate{
					TO0:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  p("mai?"),
					Right: p("pai?"),
				},
				&CoPPredicate{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ri"},
					Left:  p("mai?"),
					Right: p("pai?"),
				},
				&CoPPredicate{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  p("de?"),
					Right: p("pai?"),
				},
				&CoPPredicate{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  p("mai?"),
					Right: p("de?"),
				},
				&CoPPredicate{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Right: p("mai?"),
					Left:  p("pai?"),
				},
			},
		},
		{
			n: &LinkedTerm{
				GO:       Word{T: "go"},
				Argument: &PredicateArgument{Predicate: p("ji/")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&LinkedTerm{
					GO:       Word{T: "gogo-"},
					Argument: &PredicateArgument{Predicate: p("ji/")},
				},
				&LinkedTerm{
					GO:       Word{T: "go"},
					Argument: &PredicateArgument{Predicate: p("suq/")},
				},
			},
		},
		{
			n: &TermSet{
				TO0:   &Word{T: "to"},
				TO1:   &Word{T: "to"},
				RU:    Word{T: "ru"},
				Left:  &PredicateArgument{Predicate: p("mai?")},
				Right: &PredicateArgument{Predicate: p("pai?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&TermSet{
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&TermSet{
					TO0:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&TermSet{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ri"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&TermSet{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("de?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&TermSet{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("de?")},
				},
				&TermSet{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Right: &PredicateArgument{Predicate: p("mai?")},
					Left:  &PredicateArgument{Predicate: p("pai?")},
				},
			},
		},
		{
			n: &PredicateArgument{
				Focus:      &Word{T: "ku"},
				Quantifier: &Word{T: "tu"},
				Predicate:  p("mai/"),
				Relative: &PredicationRelative{
					Predication: Predication{Predicate: p("paiV")},
				},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&PredicateArgument{
					Focus:      &Word{T: "kuku"},
					Quantifier: &Word{T: "tu"},
					Predicate:  p("mai/"),
					Relative: &PredicationRelative{
						Predication: Predication{Predicate: p("paiV")},
					},
				},
				&PredicateArgument{
					Focus:      &Word{T: "ku"},
					Quantifier: &Word{T: "ke"},
					Predicate:  p("mai/"),
					Relative: &PredicationRelative{
						Predication: Predication{Predicate: p("paiV")},
					},
				},
				&PredicateArgument{
					Focus:      &Word{T: "ku"},
					Quantifier: &Word{T: "tu"},
					Predicate:  p("pai/"),
					Relative: &PredicationRelative{
						Predication: Predication{Predicate: p("paiV")},
					},
				},
				&PredicateArgument{
					Focus:      &Word{T: "ku"},
					Quantifier: &Word{T: "tu"},
					Predicate:  p("mai/"),
					Relative: &PredicationRelative{
						Predication: Predication{Predicate: p("jaiV")},
					},
				},
			},
		},
		{
			n: &CoPArgument{
				TO0:   &Word{T: "to"},
				TO1:   &Word{T: "to"},
				RU:    Word{T: "ru"},
				Left:  &PredicateArgument{Predicate: p("mai?")},
				Right: &PredicateArgument{Predicate: p("pai?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&CoPArgument{
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&CoPArgument{
					TO0:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&CoPArgument{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ri"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&CoPArgument{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("de?")},
					Right: &PredicateArgument{Predicate: p("pai?")},
				},
				&CoPArgument{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateArgument{Predicate: p("mai?")},
					Right: &PredicateArgument{Predicate: p("de?")},
				},
				&CoPArgument{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Right: &PredicateArgument{Predicate: p("mai?")},
					Left:  &PredicateArgument{Predicate: p("pai?")},
				},
			},
		},
		{
			n: &PredicationRelative{
				Predication: Predication{Predicate: p("maiV")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&PredicationRelative{
					Predication: Predication{Predicate: p("paiV")},
				},
			},
		},
		{
			n: &LURelative{
				LU:        Word{T: "luV"},
				Statement: &Predication{Predicate: p("mai?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&LURelative{
					LU:        Word{T: "luVlu-"},
					Statement: &Predication{Predicate: p("mai?")},
				},
				&LURelative{
					LU:        Word{T: "luV"},
					Statement: &Predication{Predicate: p("pai?")},
				},
			},
		}, {
			n: &CoPRelative{
				TO0:   &Word{T: "to"},
				TO1:   &Word{T: "to"},
				RU:    Word{T: "ru"},
				Left:  &Predication{Predicate: p("mai?")},
				Right: &Predication{Predicate: p("pai?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&CoPRelative{
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPRelative{
					TO0:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPRelative{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ri"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPRelative{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("de?")},
					Right: &Predication{Predicate: p("pai?")},
				},
				&CoPRelative{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &Predication{Predicate: p("mai?")},
					Right: &Predication{Predicate: p("de?")},
				},
				&CoPRelative{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Right: &Predication{Predicate: p("mai?")},
					Left:  &Predication{Predicate: p("pai?")},
				},
			},
		},
		{
			n: &PredicateAdverb{Predicate: p("mai~")},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&PredicateAdverb{Predicate: p("pai~")},
			},
		},
		{
			n: &CoPAdverb{
				TO0:   &Word{T: "to"},
				TO1:   &Word{T: "to"},
				RU:    Word{T: "ru"},
				Left:  &PredicateAdverb{Predicate: p("mai?")},
				Right: &PredicateAdverb{Predicate: p("pai?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&CoPAdverb{
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateAdverb{Predicate: p("mai?")},
					Right: &PredicateAdverb{Predicate: p("pai?")},
				},
				&CoPAdverb{
					TO0:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateAdverb{Predicate: p("mai?")},
					Right: &PredicateAdverb{Predicate: p("pai?")},
				},
				&CoPAdverb{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ri"},
					Left:  &PredicateAdverb{Predicate: p("mai?")},
					Right: &PredicateAdverb{Predicate: p("pai?")},
				},
				&CoPAdverb{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateAdverb{Predicate: p("de?")},
					Right: &PredicateAdverb{Predicate: p("pai?")},
				},
				&CoPAdverb{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &PredicateAdverb{Predicate: p("mai?")},
					Right: &PredicateAdverb{Predicate: p("de?")},
				},
				&CoPAdverb{
					TO0:   &Word{T: "to"},
					TO1:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Right: &PredicateAdverb{Predicate: p("mai?")},
					Left:  &PredicateAdverb{Predicate: p("pai?")},
				},
			},
		},
		{
			n: &PredicationPreposition{
				Predicate: p(`mai\`),
				Argument:  &PredicateArgument{Predicate: p("ji/")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&PredicationPreposition{
					Predicate: p(`pai\`),
					Argument:  &PredicateArgument{Predicate: p("ji/")},
				},
				&PredicationPreposition{
					Predicate: p(`mai\`),
					Argument:  &PredicateArgument{Predicate: p("suq/")},
				},
			},
		},
		{
			n: &CoPPreposition{
				TO0: &Word{T: "to"},
				TO1: &Word{T: "to"},
				RU:  Word{T: "ru"},
				Left: &PredicationPreposition{
					Predicate: p(`mai\`),
					Argument:  &PredicateArgument{Predicate: p("suq/")},
				},
				Right: &PredicationPreposition{
					Predicate: p(`pai\`),
					Argument:  &PredicateArgument{Predicate: p("suq/")},
				},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&CoPPreposition{
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationPreposition{
						Predicate: p(`mai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
					Right: &PredicationPreposition{
						Predicate: p(`pai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
				},
				&CoPPreposition{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationPreposition{
						Predicate: p(`mai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
					Right: &PredicationPreposition{
						Predicate: p(`pai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
				},
				&CoPPreposition{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ri"},
					Left: &PredicationPreposition{
						Predicate: p(`mai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
					Right: &PredicationPreposition{
						Predicate: p(`pai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
				},
				&CoPPreposition{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationPreposition{
						Predicate: p(`de\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
					Right: &PredicationPreposition{
						Predicate: p(`pai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
				},
				&CoPPreposition{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationPreposition{
						Predicate: p(`mai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
					Right: &PredicationPreposition{
						Predicate: p(`de\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
				},
				&CoPPreposition{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Right: &PredicationPreposition{
						Predicate: p(`mai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
					Left: &PredicationPreposition{
						Predicate: p(`pai\`),
						Argument:  &PredicateArgument{Predicate: p("suq/")},
					},
				},
			},
		},
		{
			n: &PredicationContent{
				Predication: Predication{Predicate: p("hio?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&PredicationContent{
					Predication: Predication{Predicate: p("mai?")},
				},
			},
		},
		{
			n: &LUContent{
				LU:        Word{T: "lu^"},
				Statement: &Predication{Predicate: p("hio?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&LUContent{
					LU:        Word{T: "lu^lu-"},
					Statement: &Predication{Predicate: p("hio?")},
				},
				&LUContent{
					LU:        Word{T: "lu^"},
					Statement: &Predication{Predicate: p("mai?")},
				},
			},
		},
		{
			n: &CoPContent{
				TO0: &Word{T: "to"},
				TO1: &Word{T: "to"},
				RU:  Word{T: "ru"},
				Left: &PredicationContent{
					Predication: Predication{Predicate: p("mai?")},
				},
				Right: &PredicationContent{
					Predication: Predication{Predicate: p("pai?")},
				},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&CoPContent{
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationContent{
						Predication: Predication{Predicate: p("mai?")},
					},
					Right: &PredicationContent{
						Predication: Predication{Predicate: p("pai?")},
					},
				},
				&CoPContent{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationContent{
						Predication: Predication{Predicate: p("mai?")},
					},
					Right: &PredicationContent{
						Predication: Predication{Predicate: p("pai?")},
					},
				},
				&CoPContent{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ri"},
					Left: &PredicationContent{
						Predication: Predication{Predicate: p("mai?")},
					},
					Right: &PredicationContent{
						Predication: Predication{Predicate: p("pai?")},
					},
				},
				&CoPContent{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationContent{
						Predication: Predication{Predicate: p("de?")},
					},
					Right: &PredicationContent{
						Predication: Predication{Predicate: p("pai?")},
					},
				},
				&CoPContent{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationContent{
						Predication: Predication{Predicate: p("mai?")},
					},
					Right: &PredicationContent{
						Predication: Predication{Predicate: p("de?")},
					},
				},
				&CoPContent{
					TO0: &Word{T: "to"},
					TO1: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Right: &PredicationContent{
						Predication: Predication{Predicate: p("mai?")},
					},
					Left: &PredicationContent{
						Predication: Predication{Predicate: p("pai?")},
					},
				},
			},
		},
		{
			n: &Parenthetical{
				KI:        Word{T: "ki"},
				KIO:       Word{T: "kio"},
				Discourse: Discourse{(*Interjection)(&Word{T: "m"})},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Parenthetical{
					KI:        Word{T: "kiki"},
					KIO:       Word{T: "kio"},
					Discourse: Discourse{(*Interjection)(&Word{T: "m"})},
				},
				&Parenthetical{
					KI:        Word{T: "ki"},
					KIO:       Word{T: "kiokio"},
					Discourse: Discourse{(*Interjection)(&Word{T: "m"})},
				},
				&Parenthetical{
					KI:        Word{T: "ki"},
					KIO:       Word{T: "kio"},
					Discourse: Discourse{(*Interjection)(&Word{T: "ha"})},
				},
			},
		},
		{
			n: &Incidental{
				JU:        Word{T: "ju"},
				Statement: &Predication{Predicate: p("hio?")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Incidental{
					JU:        Word{T: "juju"},
					Statement: &Predication{Predicate: p("hio?")},
				},
				&Incidental{
					JU:        Word{T: "ju"},
					Statement: &Predication{Predicate: p("mai?")},
				},
			},
		},
		{
			n: &Vocative{
				HU:       Word{T: "hu"},
				Argument: &PredicateArgument{Predicate: p("hoaq/")},
			},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Vocative{
					HU:       Word{T: "huhu"},
					Argument: &PredicateArgument{Predicate: p("hoaq/")},
				},
				&Vocative{
					HU:       Word{T: "hu"},
					Argument: &PredicateArgument{Predicate: p("loq/")},
				},
			},
		},
		{
			n: (*Interjection)(&Word{T: "ha"}),
			ne: []Node{
				(*Space)(&Word{T: " "}),
				(*Interjection)(&Word{T: "hio"}),
			},
		},
		{
			n: (*Space)(&Word{T: " "}),
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				(*Space)(&Word{T: "\n"}),
			},
		},
		{
			n: &Word{T: "jai?"},
			ne: []Node{
				(*Interjection)(&Word{T: "m"}),
				&Word{T: "mai?"},
			},
		},
	}
	for _, test := range tests {
		as := toString(test.n)
		if !Equal(test.n, test.n) {
			t.Errorf("Equal(%q, %q)=false, want true", as, as)
		}
		for _, ne := range test.ne {
			if Equal(test.n, ne) {
				bs := toString(ne)
				t.Errorf("Equal(%q, %q)=true, want false", as, bs)
			}
			if Equal(ne, test.n) {
				bs := toString(ne)
				t.Errorf("Equal(%q, %q)=true, want false", as, bs)
			}
		}
	}
}

func p(t string) *WordPredicate { return &WordPredicate{T: t} }

func toString(node Node) string {
	var s strings.Builder
	Visit(node, FuncVisitor(func(n Node) {
		if w, ok := n.(*Word); ok {
			s.WriteString(w.T)
		}
	}))
	return s.String()
}
