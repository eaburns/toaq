package parser

import "testing"

func TestToASCII(t *testing.T) {
	tests := []struct {
		text, want string
	}{
		{"", ""},
		{"hıa", "hia"},
		{"jí", "ji/"},
		{"pójījī", "po/ji-ji-"},
		{"rảı", "rai?"},
	}
	for _, test := range tests {
		if got := toASCII(test.text); got != test.want {
			t.Errorf("toASCII(%q)=%q, want %q", test.text, got, test.want)
		}
	}
}
