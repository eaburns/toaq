package ast

import "testing"

func TestBraces(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "",
			want: "",
		},
		{
			in:   "hio?",
			want: "hỉo",
		},
		{
			in:   "  hio?  ",
			want: "hỉo",
		},
		{
			in:   "hio? ji/ roi ho/ suq/",
			want: "(hỉo {[jí roı hó] súq})",
		},
		{
			in:   "rai? da ru rai? da",
			want: "({rảı da} ru {rảı da})",
		},
		{
			in:   "ji/ suq/ pa hio? ka",
			want: "({[<jí súq> pa] hỉo} ka)",
		},
		{
			in:   "rai? na ru rai? na",
			want: "({rảı na} ru {rảı na})",
		},
		{
			in:   "mu kui?",
			want: "(mu kủı)",
		},
		{
			in:   "po?ji- pai?",
			want: "(pỏjī pảı)",
		},
		{
			in:   "mi/ Hoaq?",
			want: "(mí Hỏaq)",
		},
		{
			in:   "po/ ji/",
			want: "(pó jí)",
		},
		{
			in:   "mo/ hu mi/ Hoe?mai- teo",
			want: "(mó {hu [mí Hỏemāı]} teo)",
		},
		{
			in:   "lu/ cho? ji/ hoa/",
			want: "(lú {chỏ [jí hóa]})",
		},
		{
			in:   "ji/ ru jai?",
			want: "(jí ru jảı)",
		},
		{
			in:   "go suq/ ji/ pa hio?",
			want: "({[<go súq> jí] pa} hỉo)",
		},
		{
			in:   "pai? to ru ji/ suq/ to suq/ ji/",
			want: "(pảı {to ru [jí súq] to [súq jí]})",
		},
		{
			in:   "jai/ toaqVpoq-",
			want: "(jáı tǒaqpōq)",
		},
		{
			in:   "jai/ luV cho? ji/ hoa/",
			want: "(jáı {lǔ [chỏ <jí hóa>]})",
		},
		{
			in:   "jai/ poqV ru toaqVpoq-",
			want: "(jáı {pǒq ru tǒaqpōq})",
		},
		{
			in:   "meo? ji/ bu~",
			want: "(mẻo {jí bũ})",
		},
		{
			in:   "meo? ji/ bu~ ru jai~",
			want: "(mẻo {jí [bũ ru jãı]})",
		},
		{
			in:   "ji? ti\\ ni/",
			want: "(jỉ {tì ní})",
		},
		{
			in:   "ji? ti\\ ni/ ru rao\\ ni/daq-",
			want: "(jỉ {[tì ní] ru [rào nídāq]})",
		},
		{
			in:   "cho? ji/ pai^ ji/ suq/",
			want: "(chỏ {jí [pâı <jí súq>]})",
		},
		{
			in:   "cho? ji/ lu^ ji/ suq/ pa pai?",
			want: "(chỏ {jí [lû <({jí súq} pa) pảı>]})",
		},
		{
			in:   "jai^ na ru bu? meo?",
			want: "({jâı na} ru {bủ mẻo})",
		},
		{
			in:   "kio hio? ka ki",
			want: "(kıo {hỉo ka} kı)",
		},
		{
			in:   "hahahaha hia",
			want: "(ha {ha [ha <ha hıa>]})",
		},
		{
			in:   "ji? ju de? ho/",
			want: "(jỉ {ju [dẻ hó]})",
		},
		{
			in:   " ",
			want: "",
		},
	}
	for _, test := range tests {
		p := NewParser(test.in)
		text, err := p.Text()
		if err != nil {
			t.Errorf("parse(%q)=%v, want nil", test.in, err)
			continue
		}
		if b := BracesString(text); b != test.want {
			t.Errorf("BracesString(%q)=%s, want %s", test.in, b, test.want)
		}
	}
}
