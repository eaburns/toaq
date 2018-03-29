package parser

import (
	"strings"
	"unicode/utf8"
)

var (
	diacritic = map[rune]map[string]string{
		'-': {
			"a": "ā",
			"A": "Ā",
			"u": "ū",
			"U": "Ū",
			"i": "ī",
			"ı": "ī",
			"I": "Ī",
			"o": "ō",
			"O": "Ō",
			"e": "ē",
			"E": "Ē",
		},
		'/': {
			"a": "á",
			"A": "Á",
			"u": "ú",
			"U": "Ú",
			"i": "í",
			"ı": "í",
			"I": "Í",
			"o": "ó",
			"O": "Ó",
			"e": "é",
			"E": "É",
		},
		'V': {
			"a": "ǎ",
			"A": "Ǎ",
			"u": "ǔ",
			"U": "Ǔ",
			"i": "ǐ",
			"ı": "ǐ",
			"I": "Ǐ",
			"o": "ǒ",
			"O": "Ǒ",
			"e": "ě",
			"E": "Ě",
		},
		'?': {
			"a": "ả",
			"A": "Ả",
			"u": "ủ",
			"U": "Ủ",
			"i": "ỉ",
			"ı": "ỉ",
			"I": "Ỉ",
			"o": "ỏ",
			"O": "Ỏ",
			"e": "ẻ",
			"E": "Ẻ",
		},
		'^': {
			"a": "â",
			"A": "Â",
			"u": "û",
			"U": "Û",
			"i": "î",
			"ı": "î",
			"I": "Î",
			"o": "ô",
			"O": "Ô",
			"e": "ê",
			"E": "Ê",
		},
		'\\': {
			"a": "à",
			"A": "À",
			"u": "ù",
			"U": "Ù",
			"i": "ì",
			"ı": "ì",
			"I": "Ì",
			"o": "ò",
			"O": "Ò",
			"e": "è",
			"E": "È",
		},
		'~': {
			"a": "ã",
			"A": "Ã",
			"u": "ũ",
			"U": "Ũ",
			"i": "ĩ",
			"ı": "ĩ",
			"I": "Ĩ",
			"o": "õ",
			"O": "Õ",
			"e": "ẽ",
			"E": "Ẽ",
		},
	}

	tone = func() map[rune]rune {
		m := make(map[rune]rune)
		for t, rs := range diacritic {
			for _, str := range rs {
				r, _ := utf8.DecodeRuneInString(str)
				m[r] = t
			}
		}
		return m
	}()

	ascii = func() map[rune]rune {
		m := make(map[rune]rune)
		for _, rs := range diacritic {
			for ascii, unicode := range rs {
				aR, _ := utf8.DecodeRuneInString(ascii)
				if aR == 'ı' {
					continue
				}
				uR, _ := utf8.DecodeRuneInString(unicode)
				m[uR] = aR
			}
		}
		return m
	}()

	vowels = "aeiou"

	noTone rune // 0
)

// addTone swaps the first letter of des with its diacritic variant given the tone rune.
func addTone(des, tone string) string {
	t, _ := utf8.DecodeRuneInString(tone)
	_, w := utf8.DecodeRuneInString(des)
	return diacritic[t][des[:w]] + des[w:]
}

func toASCII(txt string) string {
	var s strings.Builder
	var t rune
	for _, r := range txt {
		if r == 'ı' {
			s.WriteRune('i')
			continue
		}
		if t != noTone {
			if !strings.ContainsRune(vowels, r) && r != 'q' {
				s.WriteRune(t)
				t = noTone
			}
			s.WriteRune(r)
		} else {
			if t = tone[r]; t != noTone {
				s.WriteRune(ascii[r])
				continue
			}
			s.WriteRune(r)
		}
	}
	if t != noTone {
		s.WriteRune(t)
	}
	return s.String()
}
