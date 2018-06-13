package logic

import (
	"reflect"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/eaburns/toaq/ast"
)

func TestEmpty(t *testing.T) { (&interpTest{in: "", want: ""}).run(t) }

func TestFragment(t *testing.T) {
	// TODO: ju and ki/kio fragments should be treated as separate sentences.
}

func TestSentence(t *testing.T) { /* TODO */ }

func TestPredication(t *testing.T) {
	tests := []interpTest{
		{
			name: "0-ary predication",
			in:   "gi?",
			want: "gi().",
		},
		{
			name: "1-ary predication",
			in:   "gi? hoq/",
			want: "[ιX : hoq(X)] gi(X).",
		},
		{
			name: "2-ary predication",
			in:   "mai? ji/ suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] mai(J, S).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestPrenexStatement(t *testing.T) {
	tests := []interpTest{
		{
			name: "simple bi with 1 args",
			in:   "ji/ bi nuo?kuai-",
			want: "[ιJ : ji(J)] nuokuai().",
		},
		{
			name: "simple bi with multiple args",
			in:   "haq/ ji/ bi noq?gi-",
			want: "[ιH : haq(H)] [ιJ : ji(J)] noqgi().",
		},
		{
			name: "simple pa with 1 args",
			in:   "ji/ pa nuo?kuai-",
			want: "[ιJ : ji(J)] nuokuai(J).",
		},
		{
			name: "simple pa with multiple args",
			in:   "haq/ ji/ pa noq?gi-",
			want: "[ιH : haq(H)] [ιJ : ji(J)] noqgi(H, J).",
		},
		{
			name: "pa with both prenex args and predication args",
			in:   "ji/ pa mai? suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] mai(J, S).",
		},
		{
			name: "args distribute over predicate CoP",
			in:   "haq/ ji/ pa noq?gi- ru noq?si-",
			want: "[ιH : haq(H)] [ιJ : ji(J)] noqgi(H, J) ∧ noqsi(H, J).",
		},
		{
			name: "quantified arg",
			in:   `sa do/ bi gi? do/`,
			want: "∃X gi(X).",
		},
		{
			name: "arg CoP",
			in:   "chuq/guao- ru huai/chuo- ji/ pa mu choq?",
			// TODO: all else equal, ι-bound terms should appear in lexical order.
			want: "[ιC : chuqguao(C)] [ιJ : ji(J)] [ιH : huaichuo(H)] choq(J, C) ∧ choq(J, H).",
		},
		{
			name: "arg CoP roi",
			in:   `ji/ roi suq/ pa deoq?`,
			want: "[ιJ : ji(J)] [ιS : suq(S)] deoq(⌊J, S⌉).",
		},
		{
			name: "termset",
			in:   `to ru ji/ suq/ to suq/ ji/ pa deoq?`,
			want: "[ιJ0 : ji(J0)] [ιS0 : suq(S0)] [ιS1 : suq(S1)] [ιJ1 : ji(J1)] deoq(J0, S0) ∧ deoq(S1, J1).",
		},
		{
			name: "linked term",
			in:   `go ji/ pa choaq?`,
			want: "[ιJ : ji(J)] choaq(·, J).",
		},
		{
			name: "linked roi term",
			in:   `go ji/ roi suq/ pa choaq?`,
			want: "[ιJ : ji(J)] [ιS : suq(S)] choaq(·, ⌊J, S⌉).",
		},
		{
			name: "preposition",
			in:   `ne\ pui/poq-chao- bi fa? ji/`,
			want: "[ιJ : ji(J)] [ιP : puipoqchao(P)] ne({fa(J)}, P).",
		},
		{
			name: "adverb",
			in:   "ni~chaq- bi fa? ji/",
			want: "[ιJ : ji(J)] nichaq({fa(J)}).",
		},
		{
			name: "nested pa-prenex",
			in:   "ji/ pa to ru suq/ pa mai? to de?",
			want: "[ιJ : ji(J)] [ιS : suq(S)] mai(J, S) ∧ de(J).",
		},
		// TODO: Should nested prenexes work with after-through connected statemets?
		{
			name: "bi-prenex nested in a pa-prenex",
			in:   "ji/ pa to ru suq/ bi mai? to de?",
			// TODO: Is it a bug that this does not give mai(J)?
			want: "[ιJ : ji(J)] [ιS : suq(S)] mai() ∧ de(J).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestCoPStatement(t *testing.T) {
	tests := []interpTest{
		{
			name: "predication",
			in:   "hie? ji/ chea/ na ra deo? suq/ ",
			want: "[ιJ : ji(J)] [ιC : chea(C)] [ιS : suq(S)] hie(J, C) ∨ deo(S).",
		},
		{
			name: "predication forethrought",
			in:   "to ra hie? ji/ chea/ to deo? suq/ ",
			want: "[ιJ : ji(J)] [ιC : chea(C)] [ιS : suq(S)] hie(J, C) ∨ deo(S).",
		},
		{
			name: "prenex",
			in:   "ji/ chea/ pa hie? na ra suq/ pa bo?doa-",
			want: "[ιJ : ji(J)] [ιC : chea(C)] [ιS : suq(S)] hie(J, C) ∨ bodoa(J, C, S).",
		},
		{
			name: "prenex forethought",
			in:   "to ra ji/ chea/ pa hie? to suq/ pa deo?",
			want: "[ιJ : ji(J)] [ιC : chea(C)] [ιS : suq(S)] hie(J, C) ∨ deo(S).",
		},
		{
			name: "prenex and predication",
			in:   "to ra ji/ chea/ pa hie? to deo? suq/",
			want: "[ιJ : ji(J)] [ιC : chea(C)] [ιS : suq(S)] hie(J, C) ∨ deo(S).",
		},
		{
			name: "nested",
			in:   "to ru de? ji/ to gi? na ra hui?",
			want: "[ιJ : ji(J)] de(J) ∧ (gi() ∨ hui()).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestWordPredicate(t *testing.T) {
	tests := []interpTest{
		{
			name: "simple",
			in:   "hio?ji-suq-ka-",
			want: "hiojisuqka().",
		},
		{
			name: "bound variable",
			in:   "sa do/ bi do? ji/",
			want: "[ιJ : ji(J)] ∃X X(J).",
		},
		{
			name: "unbound variable",
			in:   "do? ji/",
			want: "[ιJ : ji(J)] U(J).",
		},
		{
			name: "bound hoa",
			in:   "lu? hoa? ji/",
			want: "[ιJ : ji(J)] {λX X(J)}().",
		},
		{
			name: "unbound hoa",
			in:   "hoa? ji/",
			want: "[ιJ : ji(J)] U(J).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestQuantifiedPredicate(t *testing.T) {
	t.Skip()
	panic("quantified predicates are unimplemented")
}

func TestMUPredicate(t *testing.T) {
	tests := []interpTest{
		{
			name: "simple",
			in:   "mu pai? suq/ ji/",
			want: "[ιS : suq(S)] [ιJ : ji(J)] pai(J, S).",
		},
		{
			name: "makes nil arg",
			in:   "mu kui? de^",
			want: "kui(·, {de()}).",
		},
		{
			name: "with linking terms",
			in:   "mu pai? go ji/ suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] pai(J, S).",
		},
		{
			name: "nested",
			in:   "mu mu mu mu pai? ji/ suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] pai(J, S).",
		},
		{
			name: "not simplified away when with a serial",
			in:   "suq? mu pai? ho/",
			want: "[ιH : ho(H)] suq∘(mu pai)(H).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestSerialPredicate(t *testing.T) {
	tests := []interpTest{
		{
			name: "simple",
			in:   "gu? poq?",
			want: "gu∘poq().",
		},
		{
			name: "nested",
			in:   "gu? poq? de?",
			want: "gu∘poq∘de().",
		},
		{
			name: "mu",
			in:   "gi? mu kui?",
			want: "gi∘(mu kui)().",
		},
		{
			name: "mi",
			in:   "ji? mi? Hoaq?",
			want: "ji∘{λX chua(‹Hỏaq›, X)}().",
		},
		{
			name: "po",
			in:   "ji? po? poq/",
			want: "[ιP : poq(P)] ji∘{λX raq(X, P)}().",
		},
		{
			name: "lu",
			in:   "ji? lu? mai? ji/ hoa/",
			want: "[ιJ : ji(J)] ji∘{λX mai(J, X)}().",
		},
		{
			name: "CoP pred",
			in:   "ji? mi? Hoaq? ga ru de?",
			want: "ji∘{λX chua(‹Hỏaq›, X)}() ∧ ji∘de().",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestMIPredicate(t *testing.T) {
	tests := []interpTest{
		{
			name: "mi pred",
			in:   "mi? Hoaq?",
			want: "{λX chua(‹Hỏaq›, X)}().",
		},
		{
			name: "shu pred",
			in:   "shu? toa?",
			want: "{λX X=‹tỏa›}().",
		},
		{
			name: "mi arg",
			in:   "rai? mi/ Hoaq?",
			want: "[ιX : chua(‹Hỏaq›, X)] rai(X).",
		},
		{
			name: "shu arg",
			in:   "rai? shu/ toa?",
			want: "rai(‹tỏa›).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestPOPredicate(t *testing.T) {
	tests := []interpTest{
		{
			name: "po pred",
			in:   "po? ji/",
			want: "[ιJ : ji(J)] {λX raq(X, J)}().",
		},
		{
			name: "jei pred",
			in:   "jei? ji/",
			want: "[ιJ : ji(J)] {λX X=J}().",
		},
		{
			name: "po arg",
			in:   "rai? po/ ji/",
			want: "[ιJ : ji(J)] [ιR : raq(R, J)] rai(R).",
		},
		{
			name: "jei arg",
			in:   "rai? jei/ ji/",
			// TODO: Substitue variables bound by an =.
			want: "[ιJ : ji(J)] [ιX : X=J] rai(X).",
		},
		{
			name: "quantifier",
			in:   "rai? po/ sa poq/",
			want: "[ιX : [∃P : poq(P)] raq(X, P)] rai(X).",
		},
		{
			name: "roi",
			in:   "rai? po/ to roi ji/ to suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] [ιX : raq(X, ⌊J, S⌉)] rai(X).",
		},
		{
			name: "CoP",
			in:   "rai? po/ ji/ ru suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] [ιX : raq(X, J) ∧ raq(X, S)] rai(X).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func TestMOPredicate(t *testing.T) {
	tests := []interpTest{
		{
			name: "pred",
			in:   "mo? Hio? ji/ suq/ ka teo",
			want: "{λX X=«Hỉo jí súq ka»}().",
		},
		{
			name: "arg",
			in:   "rai? mo/ Hio? ji/ suq/ ka teo",
			want: "rai(«Hỉo jí súq ka»).",
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

// TODO: lu, li, lo, ma, tio
func TestLUPredicate(t *testing.T) { /* TODO */ }

func TestCoPPredicate(t *testing.T) { /* TODO */ }

func TestCoPArgument(t *testing.T) {
	tests := []interpTest{
		{
			name: "ru",
			in:   "de? ji/ ru suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] de(J) ∧ de(S).",
		},
		{
			name: "ra",
			in:   "de? ji/ ra suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] de(J) ∨ de(S).",
		},
		{
			name: "ri",
			in:   "de? ji/ ri suq/",
			want: "[ιJ : ji(J)] [ιS : suq(S)] de(J) ? de(S).",
		},
		// roi is handled separately
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

// TODO: test relative clause in each of the *Argument tests.

// TODO: test all quantifiers.
func TestWordPredicateArgument(t *testing.T) { /* TODO */ }

// TODO: lu, li, lo, ma, tio
func TestLUPredicateArgument(t *testing.T) { /* TODO */ }

// TODO: mi, shu
func TestMIPredicateArgument(t *testing.T) { /* TODO */ }

// TODO: po, jei — jei currently named pe
func TestPOPredicateArgument(t *testing.T) { /* TODO */ }

func TestMOPredicateArgument(t *testing.T) { /* TODO */ }

func TestVariableArgument(t *testing.T) { /* TODO */ }

func TestHoaArgument(t *testing.T) { /* TODO */ }

func TestROIArgument(t *testing.T) { /* TODO */ }

func TestContentArgument(t *testing.T) { /* TODO */ }

func TestTermSet(t *testing.T) { /* TODO */ }

func TestLinkedTerm(t *testing.T) { /* TODO */ }

func TestCoPAdverb(t *testing.T) { /* TODO */ }

func TestPredicationAdverb(t *testing.T) { /* TODO */ }

func TestCoPPreposition(t *testing.T) { /* TODO */ }

func TestPredicationPreposition(t *testing.T) { /* TODO */ }

// TestTermScope tests normal, left-scoping terms.
func TestTermScope(t *testing.T) { /* TODO */ }

// TestHIScope tests that hi-bound terms have subordinate-top scope.
func TestHIScope(t *testing.T) { /* TODO */ }

// TestLambdaScope tests that λ-bound terms have lambda-top scope.
func TestLambdaScope(t *testing.T) { /* TODO */ }

// TestIotaScope tests that ι-bound terms have top scope.
func TestIotaScope(t *testing.T) { /* TODO */ }

type interpTest struct {
	name string
	in   string
	want string
}

func (test *interpTest) run(t *testing.T) {
	p := ast.NewParser(test.in)
	text, err := p.Text()
	if err != nil {
		t.Fatalf("parse(%s) failed: %v", test.in, err)
	}
	got := String(Interpret(text))
	gotTokens := tokenize(got)
	wantTokens := tokenize(test.want)
	if len(gotTokens) != len(wantTokens) {
		t.Fatalf("interp(%s)=%v, want %v (%d tokens != %d tokens)",
			test.in, got, test.want, len(gotTokens), len(wantTokens))
	}
	vars := make(map[string]string)
	for i, g := range gotTokens {
		w := wantTokens[i]
		if r, _ := utf8.DecodeRuneInString(w); !unicode.IsUpper(r) {
			if g != w {
				t.Errorf("interp(%s)=%v, want %v; token[%d]=%v, want %v",
					test.in, got, test.want, i, g, w)
			}
			continue
		}

		// Expect a variable.
		if r, _ := utf8.DecodeRuneInString(g); !unicode.IsUpper(r) {
			t.Errorf("interp(%s)=%v, want %v; token[%d]=%v, want %v (var)",
				test.in, got, test.want, i, g, w)
			continue
		}
		if v, ok := vars[w]; !ok {
			vars[w] = g
		} else if v != g {
			t.Errorf("interp(%s)=%v, want %v; token[%d]=%v, want %v (w=%v)",
				test.in, got, test.want, i, g, v, w)
		}
	}
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{in: "", want: nil},
		{in: "abc", want: []string{"abc"}},
		{in: "XYZ", want: []string{"XYZ"}},
		{in: "XYZ123", want: []string{"XYZ123"}},
		{in: "123", want: []string{"1", "2", "3"}}, // no upper-case prefix
		{in: "∀∃()", want: []string{"∀", "∃", "(", ")"}},
		{
			in: "[ιJ0: ji(J0)] [∀S: suq(S)] mai(J0, S)",
			want: []string{
				"[", "ι", "J0", ":", "ji", "(", "J0", ")", "]",
				"[", "∀", "S", ":", "suq", "(", "S", ")", "]",
				"mai", "(", "J0", ",", "S", ")",
			},
		},
	}
	for _, test := range tests {
		got := tokenize(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("tokenize(%s)=%#v, want %#v", test.in, got, test.want)
		}
	}
}

// tokenize returns the tokens of a string.
// Tokens are:
// 	- strings of lower-case letter runes.
// 	- strings of upper-case letters followed by zero or more digits.
// 	- any other non-whitespace rune
func tokenize(str string) []string {
	var tokens []string
	for len(str) > 0 {
		switch r, w := utf8.DecodeRuneInString(str); {
		case unicode.IsSpace(r):
			str = str[w:]
		case unicode.IsUpper(r):
			n := tokenLen(str, unicode.IsUpper, unicode.IsDigit)
			tokens, str = append(tokens, str[:n]), str[n:]
		case unicode.IsLower(r):
			n := tokenLen(str, unicode.IsLower)
			tokens, str = append(tokens, str[:n]), str[n:]
		default:
			tokens = append(tokens, str[:w])
			str = str[w:]
		}
	}
	return tokens
}

func tokenLen(str string, fs ...func(rune) bool) int {
	var i int
	r, w := utf8.DecodeRuneInString(str)
	for _, f := range fs {
		for f(r) {
			i += w
			r, w = utf8.DecodeRuneInString(str[i:])
		}
	}
	return i
}
