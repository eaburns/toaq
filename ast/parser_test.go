package ast

import (
	"strings"
	"testing"

	"github.com/eaburns/pretty"
	"github.com/google/go-cmp/cmp"
)

func TestDiscourse(t *testing.T) {
	tests := []parserTest{
		{
			name: "empty",
			text: "",
			want: &Text{},
		},
		{
			name: "empty",
			text: "",
			want: &Text{},
		},
		{
			name: "leading space",
			text: " ",
			want: &Text{Leading: &Space{T: " "}},
		},
		{
			name: "sentence",
			text: "hio?",
			want: &Text{
				Discourse: Discourse{
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
						},
					},
				},
			},
		},
		{
			name: "fragment",
			text: "ji/ bi",
			want: &Text{
				Discourse: Discourse{
					&Prenex{
						Terms: Terms{
							&PredicateArgument{
								Predicate: &WordPredicate{T: "ji/"},
							},
						},
						BI: Word{T: "bi"},
					},
				},
			},
		},
		{
			name: "multiple items",
			text: "hio? ka jai? da ji/",
			want: &Text{
				Discourse: Discourse{
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
						},
						DA: &Word{T: "ka"},
					},
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "jai?"},
						},
						DA: &Word{T: "da"},
					},
					Terms{
						&PredicateArgument{
							Predicate: &WordPredicate{T: "ji/"},
						},
					},
				},
			},
		},
		{
			name: "error",
			text: "  ☹q",
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func discourse(n Node) Node { return &Text{Discourse: Discourse{n}} }

func TestSentence(t *testing.T) {
	tests := []parserTest{
		{
			name: "statement sentence",
			text: "je hio? ka",
			want: discourse(
				&StatementSentence{
					JE: &Word{T: "je"},
					Statement: &Predication{
						Predicate: &WordPredicate{T: "hio?"},
					},
					DA: &Word{T: "ka"},
				},
			),
		},
		{
			name: "sentence afterthought cop",
			text: "hio? ka ru jai? da ru meo?",
			want: discourse(
				&CoPSentence{
					Left: &StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
						},
						DA: &Word{T: "ka"},
					},
					RU: Word{T: "ru"},
					Right: &CoPSentence{
						Left: &StatementSentence{
							Statement: &Predication{
								Predicate: &WordPredicate{T: "jai?"},
							},
							DA: &Word{T: "da"},
						},
						RU: Word{T: "ru"},
						Right: &StatementSentence{
							Statement: &Predication{
								Predicate: &WordPredicate{T: "meo?"},
							},
						},
					},
				},
			),
		},
		{
			name: "sentence forethought cop",
			text: "to ru hio? ka to jai?",
			want: discourse(
				&CoPSentence{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
						},
						DA: &Word{T: "ka"},
					},
					TO1: &Word{T: "to"},
					Right: &StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "jai?"},
						},
					},
				},
			),
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func sentence(n Statement) Node {
	return discourse(&StatementSentence{Statement: n})
}

func TestStatement(t *testing.T) {
	tests := []parserTest{
		{
			name: "predication",
			text: "hio? ji/ suq/ na",
			want: sentence(
				&Predication{
					Predicate: &WordPredicate{T: "hio?"},
					Terms: Terms{
						&PredicateArgument{
							Predicate: &WordPredicate{T: "ji/"},
						},
						&PredicateArgument{
							Predicate: &WordPredicate{T: "suq/"},
						},
					},
					NA: &Word{T: "na"},
				},
			),
		},
		{
			name: "prenex statement",
			text: "sa do/ bi jai? do/ na ru meo? do/",
			want: sentence(
				&PrenexStatement{
					Prenex: Prenex{
						Terms: Terms{
							&PredicateArgument{
								Quantifier: &Word{T: "sa"},
								Predicate:  &WordPredicate{T: "do/"},
							},
						},
						BI: Word{T: "bi"},
					},
					Statement: &CoPStatement{
						Left: &Predication{
							Predicate: &WordPredicate{T: "jai?"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "do/"},
								},
							},
							NA: &Word{T: "na"},
						},
						RU: Word{T: "ru"},
						Right: &Predication{
							Predicate: &WordPredicate{T: "meo?"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "do/"},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "statement afterthought cop",
			text: "hio? na ru jai?",
			want: sentence(
				&CoPStatement{
					Left: &Predication{
						Predicate: &WordPredicate{T: "hio?"},
						NA:        &Word{T: "na"},
					},
					RU: Word{T: "ru"},
					Right: &Predication{
						Predicate: &WordPredicate{T: "jai?"},
					},
				},
			),
		},
		{
			name: "statement afterthought cop",
			text: "hio? na ru jai? na ru meo?",
			want: sentence(
				&CoPStatement{
					Left: &Predication{
						Predicate: &WordPredicate{T: "hio?"},
						NA:        &Word{T: "na"},
					},
					RU: Word{T: "ru"},
					Right: &CoPStatement{
						Left: &Predication{
							Predicate: &WordPredicate{T: "jai?"},
							NA:        &Word{T: "na"},
						},
						RU: Word{T: "ru"},
						Right: &Predication{
							Predicate: &WordPredicate{T: "meo?"},
						},
					},
				},
			),
		},
		{
			name: "statement forethought cop",
			text: "to ru hio? na to jai?",
			want: sentence(
				&CoPStatement{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &Predication{
						Predicate: &WordPredicate{T: "hio?"},
						NA:        &Word{T: "na"},
					},
					TO1: &Word{T: "to"},
					Right: &Predication{
						Predicate: &WordPredicate{T: "jai?"},
					},
				},
			),
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func statement(n Predicate, ts ...Term) Node {
	return sentence(&Predication{Predicate: n, Terms: Terms(ts)})
}

func TestPredicate(t *testing.T) {
	tests := []parserTest{
		{
			name: "word predicate",
			text: "hio?",
			want: statement(
				&WordPredicate{T: "hio?"},
			),
		},
		{
			name: "prefixed predicate",
			text: "mu hio?",
			want: statement(
				&PrefixedPredicate{
					MU:        Word{T: "mu"},
					Predicate: &WordPredicate{T: "hio?"},
				},
			),
		},
		{
			name: "serial predicate",
			text: "gi? gai? mai?",
			want: statement(
				&SerialPredicate{
					Left: &WordPredicate{T: "gi?"},
					Right: &SerialPredicate{
						Left:  &WordPredicate{T: "gai?"},
						Right: &WordPredicate{T: "mai?"},
					},
				},
			),
		},
		{
			name: "mi predicate",
			text: "mi? toaq? ga",
			want: statement(
				&MIPredicate{
					MI:     Word{T: "mi?"},
					Phrase: &WordPredicate{T: "toaq?"},
					GA:     &Word{T: "ga"},
				},
			),
		},
		{
			name: "po predicate",
			text: "po? toaq/ ga",
			want: statement(
				&POPredicate{
					PO: Word{T: "po?"},
					Argument: &PredicateArgument{
						Predicate: &WordPredicate{T: "toaq/"},
					},
					GA: &Word{T: "ga"},
				},
			),
		},
		{
			name: "jei predicate",
			text: "jei? toaq/ ga",
			want: statement(
				&POPredicate{
					PO: Word{T: "jei?"},
					Argument: &PredicateArgument{
						Predicate: &WordPredicate{T: "toaq/"},
					},
					GA: &Word{T: "ga"},
				},
			),
		},
		{
			name: "mo predicate",
			text: "mo? hio? teo",
			want: statement(
				&MOPredicate{
					MO: Word{T: "mo?"},
					Discourse: Discourse{
						&StatementSentence{
							Statement: &Predication{
								Predicate: &WordPredicate{T: "hio?"},
							},
						},
					},
					TEO: Word{T: "teo"},
				},
			),
		},
		{
			name: "lu predicate",
			text: "lu? mai? na",
			want: statement(
				&LUPredicate{
					LU: Word{T: "lu?"},
					Statement: &Predication{
						Predicate: &WordPredicate{T: "mai?"},
						NA:        &Word{T: "na"},
					},
				},
			),
		},
		{
			name: "predicate afterthought cop",
			text: "hio? ru jai?",
			want: statement(
				&CoPPredicate{
					Left:  &WordPredicate{T: "hio?"},
					RU:    Word{T: "ru"},
					Right: &WordPredicate{T: "jai?"},
				},
			),
		},
		{
			name: "predicate forethought cop",
			text: "to ru hio? to jai?",
			want: statement(
				&CoPPredicate{
					TO0:   &Word{T: "to"},
					RU:    Word{T: "ru"},
					Left:  &WordPredicate{T: "hio?"},
					TO1:   &Word{T: "to"},
					Right: &WordPredicate{T: "jai?"},
				},
			),
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestTerms(t *testing.T) {
	tests := []parserTest{
		{
			name: "linked argument",
			text: "hio? go ji/",
			want: statement(
				&WordPredicate{T: "hio?"},
				&LinkedTerm{
					GO: Word{T: "go"},
					Argument: &PredicateArgument{
						Predicate: &WordPredicate{T: "ji/"},
					},
				},
			),
		},
		{
			name: "termset",
			text: "hio? to ru ji/ suq/ to suq/ ji/",
			want: statement(
				&WordPredicate{T: "hio?"},
				&TermSet{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: Terms{
						&PredicateArgument{
							Predicate: &WordPredicate{T: "ji/"},
						},
						&PredicateArgument{
							Predicate: &WordPredicate{T: "suq/"},
						},
					},
					TO1: &Word{T: "to"},
					Right: Terms{
						&PredicateArgument{
							Predicate: &WordPredicate{T: "suq/"},
						},
						&PredicateArgument{
							Predicate: &WordPredicate{T: "ji/"},
						},
					},
				},
			),
		},
		{
			name: "unbalanced termset",
			text: "hio? to ru ji/ suq/ mai~ to suq/ ji/☹",
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestArgument(t *testing.T) {
	tests := []parserTest{
		{
			name: "predicate argument",
			text: "hio? ku sa ji/ maiV suq/",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateArgument{
					Focus:      &Word{T: "ku"},
					Quantifier: &Word{T: "sa"},
					Predicate:  &WordPredicate{T: "ji/"},
					Relative: &PredicationRelative{
						Predication: Predication{
							Predicate: &WordPredicate{T: "maiV"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "suq/"},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "complex predicate argument",
			text: "hio? ji/ ru suq?",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateArgument{
					Predicate: &CoPPredicate{
						Left:  &WordPredicate{T: "ji/"},
						RU:    Word{T: "ru"},
						Right: &WordPredicate{T: "suq?"},
					},
				},
			),
		},
		{
			name: "argument afterthought cop",
			text: "hio? ji/ ru suq/",
			want: statement(
				&WordPredicate{T: "hio?"},
				&CoPArgument{
					Left: &PredicateArgument{
						Predicate: &WordPredicate{T: "ji/"},
					},
					RU: Word{T: "ru"},
					Right: &PredicateArgument{
						Predicate: &WordPredicate{T: "suq/"},
					},
				},
			),
		},
		{
			name: "argument forethought cop",
			text: "hio? to ru ji/ to suq/",
			want: statement(
				&WordPredicate{T: "hio?"},
				&CoPArgument{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicateArgument{
						Predicate: &WordPredicate{T: "ji/"},
					},
					TO1: &Word{T: "to"},
					Right: &PredicateArgument{
						Predicate: &WordPredicate{T: "suq/"},
					},
				},
			),
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestRelative(t *testing.T) {
	tests := []parserTest{
		{
			name: "predication relative",
			text: "hio? ji/ maiV suq/",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateArgument{
					Predicate: &WordPredicate{T: "ji/"},
					Relative: &PredicationRelative{
						Predication: Predication{
							Predicate: &WordPredicate{T: "maiV"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "suq/"},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "complex predication relative",
			text: "hio? ji/ maiV ru pai?",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateArgument{
					Predicate: &WordPredicate{T: "ji/"},
					Relative: &PredicationRelative{
						Predication: Predication{
							Predicate: &CoPPredicate{
								Left:  &WordPredicate{T: "maiV"},
								RU:    Word{T: "ru"},
								Right: &WordPredicate{T: "pai?"},
							},
						},
					},
				},
			),
		},
		{
			name: "lu relative",
			text: "hio? ji/ luV mai? suq/ na",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateArgument{
					Predicate: &WordPredicate{T: "ji/"},
					Relative: &LURelative{
						LU: Word{T: "luV"},
						Statement: &Predication{
							Predicate: &WordPredicate{T: "mai?"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "suq/"},
								},
							},
							NA: &Word{T: "na"},
						},
					},
				},
			),
		},
		{
			name: "relative afterthought cop",
			text: "hio? ji/ maiV ru paiV",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateArgument{
					Predicate: &WordPredicate{T: "ji/"},
					Relative: &CoPRelative{
						Left: &PredicationRelative{
							Predication: Predication{
								Predicate: &WordPredicate{T: "maiV"},
							},
						},
						RU: Word{T: "ru"},
						Right: &PredicationRelative{
							Predication: Predication{
								Predicate: &WordPredicate{T: "paiV"},
							},
						},
					},
				},
			),
		},
		{
			name: "relative forethought cop",
			text: "hio? ji/ to ru maiV to paiV",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateArgument{
					Predicate: &WordPredicate{T: "ji/"},
					Relative: &CoPRelative{
						TO0: &Word{T: "to"},
						RU:  Word{T: "ru"},
						Left: &PredicationRelative{
							Predication: Predication{
								Predicate: &WordPredicate{T: "maiV"},
							},
						},
						TO1: &Word{T: "to"},
						Right: &PredicationRelative{
							Predication: Predication{
								Predicate: &WordPredicate{T: "paiV"},
							},
						},
					},
				},
			),
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestAdverb(t *testing.T) {
	tests := []parserTest{
		{
			name: "predicate adverb",
			text: "hio? jai~",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateAdverb{
					Predicate: &WordPredicate{T: "jai~"},
				},
			),
		},
		{
			name: "complex predicate adverb",
			text: "hio? jai~ ru meo?",
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicateAdverb{
					Predicate: &CoPPredicate{
						Left:  &WordPredicate{T: "jai~"},
						RU:    Word{T: "ru"},
						Right: &WordPredicate{T: "meo?"},
					},
				},
			),
		},
		{
			name: "adverb afterthought cop",
			text: "hio? jai~ ru meo~",
			want: statement(
				&WordPredicate{T: "hio?"},
				&CoPAdverb{
					Left: &PredicateAdverb{
						Predicate: &WordPredicate{T: "jai~"},
					},
					RU: Word{T: "ru"},
					Right: &PredicateAdverb{
						Predicate: &WordPredicate{T: "meo~"},
					},
				},
			),
		},
		{
			name: "adverb forethought cop",
			text: "hio? to ru jai~ to meo~",
			want: statement(
				&WordPredicate{T: "hio?"},
				&CoPAdverb{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicateAdverb{
						Predicate: &WordPredicate{T: "jai~"},
					},
					TO1: &Word{T: "to"},
					Right: &PredicateAdverb{
						Predicate: &WordPredicate{T: "meo~"},
					},
				},
			),
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestPreposition(t *testing.T) {
	tests := []parserTest{
		{
			name: "predication preposition",
			text: `hio? ti\ ni/`,
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicationPreposition{
					Predicate: &WordPredicate{T: `ti\`},
					Argument: &PredicateArgument{
						Predicate: &WordPredicate{T: "ni/"},
					},
				},
			),
		},
		{
			name: "complex predication verb-tone preposition",
			text: `hio? mai\ ru pai? ni/`,
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicationPreposition{
					Predicate: &CoPPredicate{
						Left:  &WordPredicate{T: `mai\`},
						RU:    Word{T: "ru"},
						Right: &WordPredicate{T: "pai?"},
					},
					Argument: &PredicateArgument{
						Predicate: &WordPredicate{T: "ni/"},
					},
				},
			),
		},
		{
			name: "complex predication prep-tone preposition",
			text: `hio? mai\ ru pai\ ni/`,
			want: statement(
				&WordPredicate{T: "hio?"},
				&PredicationPreposition{
					Predicate: &CoPPredicate{
						Left:  &WordPredicate{T: `mai\`},
						RU:    Word{T: "ru"},
						Right: &WordPredicate{T: `pai\`},
					},
					Argument: &PredicateArgument{
						Predicate: &WordPredicate{T: "ni/"},
					},
				},
			),
		},
		{
			name: "preposition afterthought cop",
			text: `hio? ti\ ni/ ru kui\ ji/`,
			want: statement(
				&WordPredicate{T: "hio?"},
				&CoPPreposition{
					Left: &PredicationPreposition{
						Predicate: &WordPredicate{T: `ti\`},
						Argument: &PredicateArgument{
							Predicate: &WordPredicate{T: "ni/"},
						},
					},
					RU: Word{T: "ru"},
					Right: &PredicationPreposition{
						Predicate: &WordPredicate{T: `kui\`},
						Argument: &PredicateArgument{
							Predicate: &WordPredicate{T: "ji/"},
						},
					},
				},
			),
		},
		{
			name: "adverb forethought cop",
			text: `hio? to ru ti\ ni/ to kui\ ji/`,
			want: statement(
				&WordPredicate{T: "hio?"},
				&CoPPreposition{
					TO0: &Word{T: "to"},
					RU:  Word{T: "ru"},
					Left: &PredicationPreposition{
						Predicate: &WordPredicate{T: `ti\`},
						Argument: &PredicateArgument{
							Predicate: &WordPredicate{T: "ni/"},
						},
					},
					TO1: &Word{T: "to"},
					Right: &PredicationPreposition{
						Predicate: &WordPredicate{T: `kui\`},
						Argument: &PredicateArgument{
							Predicate: &WordPredicate{T: "ji/"},
						},
					},
				},
			),
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestContent(t *testing.T) {
	tests := []parserTest{
		{
			name: "predication content",
			text: "kui? jai^ ji/",
			want: statement(
				&WordPredicate{T: "kui?"},
				&PredicateArgument{
					Predicate: &PredicationContent{
						Predication: Predication{
							Predicate: &WordPredicate{T: "jai^"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "ji/"},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "complex predication content",
			text: "kui? jai^ ru pai? ji/",
			want: statement(
				&WordPredicate{T: "kui?"},
				&PredicateArgument{
					Predicate: &PredicationContent{
						Predication: Predication{
							Predicate: &CoPPredicate{
								Left:  &WordPredicate{T: "jai^"},
								RU:    Word{T: "ru"},
								Right: &WordPredicate{T: "pai?"},
							},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "ji/"},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "lu content",
			text: "kui? lu^ jai? ji/ na",
			want: statement(
				&WordPredicate{T: "kui?"},
				&PredicateArgument{
					Predicate: &LUContent{
						LU: Word{T: "lu^"},
						Statement: &Predication{
							Predicate: &WordPredicate{T: "jai?"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "ji/"},
								},
							},
							NA: &Word{T: "na"},
						},
					},
				},
			),
		},
		{
			name: "content afterthought cop",
			text: "kui? jai^ ji/ na ru meo? suq/",
			want: statement(
				&WordPredicate{T: "kui?"},
				&PredicateArgument{
					Predicate: &CoPContent{
						Left: &PredicationContent{
							Predication: Predication{
								Predicate: &WordPredicate{T: "jai^"},
								Terms: Terms{
									&PredicateArgument{
										Predicate: &WordPredicate{T: "ji/"},
									},
								},
								NA: &Word{T: "na"},
							},
						},
						RU: Word{T: "ru"},
						Right: &Predication{
							Predicate: &WordPredicate{T: "meo?"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "suq/"},
								},
							},
						},
					},
				},
			),
		},
		// There is no content forethought CoP.
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestMod(t *testing.T) {
	tests := []parserTest{
		{
			name: "interjection",
			text: "hia",
			want: &Text{
				Discourse: Discourse{
					&Interjection{T: "hia"},
				},
			},
		},
		{
			name: "parenthetical",
			text: "kio hio? ki",
			want: &Text{

				Discourse: Discourse{
					&Parenthetical{
						KIO: Word{T: "kio"},
						Discourse: Discourse{
							&StatementSentence{
								Statement: &Predication{
									Predicate: &WordPredicate{T: "hio?"},
								},
							},
						},
						KI: Word{T: "ki"},
					},
				},
			},
		},
		{
			name: "incidental",
			text: "ju hio? na",
			want: &Text{
				Discourse: Discourse{
					&Incidental{
						JU: Word{T: "ju"},
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
							NA:        &Word{T: "na"},
						},
					},
				},
			},
		},
		{
			name: "vocative",
			text: "hu mi/ hoaq? ga",
			want: &Text{
				Discourse: Discourse{
					&Vocative{
						HU: Word{T: "hu"},
						Argument: &PredicateArgument{
							Predicate: &MIPredicate{
								MI:     Word{T: "mi/"},
								Phrase: &WordPredicate{T: "hoaq?"},
								GA:     &Word{T: "ga"},
							},
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

func TestWhitespace(t *testing.T) {
	tests := []parserTest{
		{
			name: "spaces, tabs, and newlines",
			text: "	\nhio?    ji/",
			want: &Text{
				Leading: &Space{T: "	\n"},
				Discourse: Discourse{
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "ji/"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "various punctuation",
			text: ".—(hio?    )!!ji/«]",
			want: &Text{
				Leading: &Space{T: ".—("},
				Discourse: Discourse{
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
							Terms: Terms{
								&PredicateArgument{
									Predicate: &WordPredicate{T: "ji/"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "allow ? after moq",
			text: "hio? moq??? jai?",
			want: &Text{
				Discourse: Discourse{
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "hio?"},
						},
						DA: &Word{T: "moq"},
					},
					&StatementSentence{
						Statement: &Predication{
							Predicate: &WordPredicate{T: "jai?"},
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		test.run(t)
	}
}

type parserTest struct {
	name string
	text string
	want Node
}

func (test parserTest) run(t *testing.T) {
	t.Run(test.name, func(t *testing.T) {
		if strings.ContainsRune(test.text, '☹') {
			test.runErrTest(t)
			return
		}
		test.runParseTest(t)
	})
}

func (test parserTest) runParseTest(t *testing.T) {
	text, err := NewParser(test.text).Text()
	if err != nil {
		t.Errorf("%q, got error %v, want %q", test.text, err, test.want)
		return
	}
	normalize(text)
	if !Equal(test.want, text) {
		t.Log(cmp.Diff(test.want, text))
		t.Errorf("parse(%q)=%s\nwant\n%s", test.text, pretty.String(text), pretty.String(test.want))
	}
}

func (test parserTest) runErrTest(t *testing.T) {
	want := strings.IndexRune(test.text, '☹')
	input := strings.Replace(test.text, "☹", "", 1)
	_, err := NewParser(test.text).Text()
	if err == nil {
		t.Errorf("%q, got error %v, want error at %d", input, err, want)
		return
	}
	if got := err.ByteOffset(); got != want {
		t.Errorf("%q, got error %v at %d, want error at %d", input, err, got, want)
	}
}

// normalize sets the S and E of all words to 0,
// converts text from diacritic form to ASCII,
// and removes all space nodes following a word
// (all but a space at the beginning of Text.Leading).
func normalize(node Node) {
	Visit(node, FuncVisitor(func(n Node) {
		if w, ok := n.(*Word); ok {
			w.S, w.E = 0, 0
			ToASCII(w)
			for {
				s, ok := w.M.(*Space)
				if !ok {
					break
				}
				w.M = s.M
			}
		}
	}))
}
