package tone

import (
	"strings"
	"unicode/utf8"
)

var (
	// Diacritic is a mapping from tones and letters to the diacritic form of the letter.
	Diacritic = map[rune]map[string]string{
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

	toneNameRunes = `-/V?^\~`

	tone = func() map[rune]rune {
		m := make(map[rune]rune)
		for t, rs := range Diacritic {
			for _, str := range rs {
				r, _ := utf8.DecodeRuneInString(str)
				m[r] = t
			}
		}
		return m
	}()

	ascii = func() map[rune]rune {
		m := make(map[rune]rune)
		for _, rs := range Diacritic {
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

	vowels = "aeiıouAEIOU" + func() string {
		var vs string
		for _, rs := range Diacritic {
			for _, d := range rs {
				vs += d
			}
		}
		return vs
	}()

	// None is the tone marker indicating no tone.
	None rune // 0
)

// ToASCII converts a single word string to standard ASCII form,
// with a tone marker following each syllable.
func ToASCII(txt string) string {
	var s strings.Builder
	var t rune
	for _, r := range txt {
		if r == 'ı' {
			s.WriteRune('i')
			continue
		}
		if t != None {
			if !strings.ContainsRune(vowels, r) && r != 'q' {
				s.WriteRune(t)
				t = None
			}
			s.WriteRune(r)
		} else {
			if t = tone[r]; t != None {
				s.WriteRune(ascii[r])
				continue
			}
			s.WriteRune(r)
		}
	}
	if t != None {
		s.WriteRune(t)
	}
	return s.String()
}

// Strip returns the ASCII representation of a single word
// with capitalization preserved, but all tone markers stripped.
func Strip(text string) string {
	var s strings.Builder
	for _, r := range ToASCII(text) {
		if !strings.ContainsRune(toneNameRunes, r) {
			s.WriteRune(r)
		}
	}
	return s.String()
}

// WithTone returns the string with the tone of the first syllable
// replaced with the given tone.
func WithTone(str string, t rune) string {
	i := strings.IndexAny(str, vowels)
	if i < 0 {
		return str
	}
	r, w := utf8.DecodeRuneInString(str[i:])
	var R string
	if a, ok := ascii[r]; ok {
		R = string([]rune{a})
	} else {
		R = string([]rune{r})
	}
	if t != None {
		R = Diacritic[t][R]
	}
	if R == "i" {
		R = "ı"
	}
	return str[:i] + R + str[i+w:]
}

// Tone returns the tone of first syllable of the string.
func Tone(s string) rune {
	ascii := ToASCII(s)
	i := strings.IndexAny(ascii, toneNameRunes)
	if i < 0 {
		return None
	}
	return rune(ascii[i])
}
