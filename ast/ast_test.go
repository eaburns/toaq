package ast

import (
	"testing"

	"github.com/eaburns/pretty"
)

func TestToASCII(t *testing.T) {
	tests := []struct {
		n, want Node
	}{
		{&Text{}, &Text{}},
		{&Word{T: "hıa"}, &Word{T: "hia"}},
		{&Word{T: "jí"}, &Word{T: "ji/"}},
		{&Word{T: "pójījī"}, &Word{T: "po/ji-ji-"}},
		{&Word{T: "rảı"}, &Word{T: "rai?"}},
		{&Word{T: "rảI"}, &Word{T: "raI?"}},
		{
			&Text{
				Discourse: Discourse{
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "rảı"},
						},
					},
				},
			},
			&Text{
				Discourse: Discourse{
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "rai?"},
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		nstr := pretty.String(test.n)
		ToASCII(test.n)
		if !Equal(test.n, test.want) {
			t.Errorf("ToASCII(%q)=%q, want %q", nstr, pretty.String(test.n), pretty.String(test.want))
		}
	}
}
