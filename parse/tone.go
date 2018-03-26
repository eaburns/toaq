package parse

import "unicode/utf8"

var diacritic = map[rune]map[string]string{
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

// addTone swaps the first letter of des with its diacritic variant given the tone rune.
func addTone(des, tone string) string {
	t, _ := utf8.DecodeRuneInString(tone)
	_, w := utf8.DecodeRuneInString(des)
	return diacritic[t][des[:w]] + des[w:]
}
