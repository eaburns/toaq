package tone

import (
	"testing"

	"github.com/eaburns/pretty"
	. "github.com/eaburns/toaq/ast"
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

func TestWithTone(t *testing.T) {
	tests := []struct {
		word string
		tone rune
		want string
	}{
		{"rảı", '-', "rāı"},
		{"rảı", '/', "ráı"},
		{"rảı", 'V', "rǎı"},
		{"ràı", '?', "rảı"},
		{"rảı", '^', "râı"},
		{"rảı", '\\', "ràı"},
		{"rảı", '~', "rãı"},
		{"rảırāı", '-', "rāırāı"},
		{"rảırāı", '/', "ráırāı"},
		{"rảırāı", 'V', "rǎırāı"},
		{"ràırāı", '?', "rảırāı"},
		{"rảırāı", '^', "râırāı"},
		{"rảırāı", '\\', "ràırāı"},
		{"rảırāı", '~', "rãırāı"},
		{"rẢırāı", '~', "rÃırāı"},
		{"rẢı", None, "rAı"},
		{"Hỉo", None, "Hıo"},
		{"rkTk", 'V', "rkTk"},
		{"do", '/', "dó"},
	}
	for _, test := range tests {
		if got := WithTone(test.word, test.tone); got != test.want {
			t.Errorf("SetTone(%q, %c)=%q, want %q", test.word, test.tone, got, test.want)
		}
	}
}
