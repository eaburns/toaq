package ast_test

import (
	"reflect"
	"testing"

	"github.com/eaburns/pretty"
	"github.com/eaburns/toaq/ast"
	"github.com/eaburns/toaq/parser"
)

func TestVisit(t *testing.T) {
	type typeCount struct {
		typ   reflect.Type
		count int
	}
	tests := []struct {
		name   string
		text   string
		counts []typeCount
	}{
		{
			name: "empty",
			text: "",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.Word{}),
					count: 0,
				},
			},
		},
		{
			name: "words",
			text: " hio? ji/ suq/ ka ha ",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.Space{}),
					count: 6,
				},
				{
					typ:   reflect.TypeOf(&ast.WordPredicate{}),
					count: 3,
				},
				{
					typ:   reflect.TypeOf(&ast.Interjection{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.Word{}),
					count: 11,
				},
			},
		},
		{
			name: "sentences",
			text: "pai? da ru mai? da",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.CoPSentence{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.StatementSentence{}),
					count: 2,
				},
			},
		},
		{
			name: "fragments",
			text: "ji/ bi raiV na rai~ rai/",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.Prenex{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.PredicationRelative{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.PredicateAdverb{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.PredicateArgument{}),
					count: 2,
				},
			},
		},
		{
			name: "statements",
			text: "jai/ bi mai? na ru pai? na",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.PrenexStatement{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.Predication{}),
					count: 2,
				},
				{
					typ:   reflect.TypeOf(&ast.CoPStatement{}),
					count: 1,
				},
			},
		},
		{
			name: "predicates",
			text: "mu rai? ru mai? pai? mi? ji/ po? ji/ mo? ji/ teo lu? rai? na",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.PrefixedPredicate{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.SerialPredicate{}),
					count: 5,
				},
				{
					typ:   reflect.TypeOf(&ast.WordPredicate{}),
					count: 7,
				},
				{
					typ:   reflect.TypeOf(&ast.MIPredicate{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.POPredicate{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.MOPredicate{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.LUPredicate{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.CoPPredicate{}),
					count: 1,
				},
			},
		},
		{
			name: "terms",
			text: "to ru go ji/ fi suq/ to cu ho/ ke maq/",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.LinkedTerm{}),
					count: 4,
				},
				{
					typ:   reflect.TypeOf(&ast.TermSet{}),
					count: 1,
				},
			},
		},
		{
			name: "arguments",
			text: "ku sia rai/ maiV ru sia rai/ paiV",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.PredicateArgument{}),
					count: 2,
				},
				{
					typ:   reflect.TypeOf(&ast.CoPArgument{}),
					count: 1,
				},
			},
		},
		{
			name: "relative clauses",
			text: "rai/ to ru maiV to luV pai?",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.PredicationRelative{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.LURelative{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.CoPRelative{}),
					count: 1,
				},
			},
		},
		{
			name: "adverbs",
			text: "rai~ ru rai~",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.PredicateAdverb{}),
					count: 2,
				},
				{
					typ:   reflect.TypeOf(&ast.CoPAdverb{}),
					count: 1,
				},
			},
		},
		{
			name: "prepositions",
			text: `rai\ ji/ ru rai\ suq/`,
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.PredicationPreposition{}),
					count: 2,
				},
				{
					typ:   reflect.TypeOf(&ast.CoPPreposition{}),
					count: 1,
				},
			},
		},
		{
			name: "content clauses",
			text: "mai^ na ru jai? na lu^ pai? na",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.PredicationContent{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.LUContent{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.CoPContent{}),
					count: 1,
				},
			},
		},
		{
			name: "freemod",
			text: "kio hio? ki ju mai? na hu mi/ Hoaq? ha",
			counts: []typeCount{
				{
					typ:   reflect.TypeOf(&ast.Parenthetical{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.Incidental{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.Vocative{}),
					count: 1,
				},
				{
					typ:   reflect.TypeOf(&ast.Interjection{}),
					count: 1,
				},
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			txt, err := parser.New(test.text).Text()
			if err != nil {
				t.Fatalf("parse(%q)=%v, want nil", test.text, err)
			}
			ok := true
			for _, tc := range test.counts {
				var got int
				ast.Visit(txt, func(n ast.Node) bool {
					if v := reflect.ValueOf(n); v.Type() == tc.typ {
						got++
					}
					return true
				})
				if got != tc.count {
					ok = false
					t.Errorf("%s: got: %d\nwant: %d", tc.typ, got, tc.count)
				}
			}
			if !ok {
				t.Logf("%s\n", pretty.String(txt))
			}
		})
	}
}
