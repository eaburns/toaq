package tone

import (
	"testing"
)

func TestToASCII(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"", ""},
		{"hıa", "hia"},
		{"jí", "ji/"},
		{"pójījī", "po/ji-ji-"},
		{"rảı", "rai?"},
		{"rảI", "raI?"},
	}
	for _, test := range tests {
		if got := ToASCII(test.in); got != test.want {
			t.Errorf("ToASCII(%q)=%q, want %q", test.in, got, test.want)
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
