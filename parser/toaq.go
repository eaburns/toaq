// FreeSans has all the glyphs used in this file.
// For Acme users:
// Font /mnt/font/FreeSans/11a/font

package parser

import "github.com/eaburns/peggy/peg"

const (
	_full_text                                         int = 0
	_text                                              int = 1
	_discourse                                         int = 2
	_sentence                                          int = 3
	_sentence_1                                        int = 4
	_sentence_2                                        int = 5
	_sentence_3                                        int = 6
	_fragment                                          int = 7
	_coP_sentence                                      int = 8
	_co_bar_sentence                                   int = 9
	_forethought_coP_sentence                          int = 10
	_forethought_coP_sentence_1                        int = 11
	_forethought_co_bar_sentence                       int = 12
	_statement                                         int = 13
	_statement_1                                       int = 14
	_statement_2                                       int = 15
	_statement_3                                       int = 16
	_prenex                                            int = 17
	_predication                                       int = 18
	_coP_statement                                     int = 19
	_co_bar_statement                                  int = 20
	_forethought_coP_statement                         int = 21
	_forethought_coP_statement_1                       int = 22
	_forethought_co_bar_statement                      int = 23
	_predicate                                         int = 24
	_predicate_1                                       int = 25
	_predicate_2                                       int = 26
	_predicate_3                                       int = 27
	_predicate_4                                       int = 28
	_serial_predicate                                  int = 29
	_serial_predicate_2                                int = 30
	_coP_pred                                          int = 31
	_co_bar_pred                                       int = 32
	_forethought_coP_pred                              int = 33
	_forethought_coP_pred_1                            int = 34
	_forethought_co_bar_pred                           int = 35
	_LU_predicate                                      int = 36
	_LU_predicate_tone                                 int = 37
	_MI_predicate                                      int = 38
	_MI_predicate_1                                    int = 39
	_MI_predicate_tone                                 int = 40
	_PO_predicate                                      int = 41
	_PO_predicate_1                                    int = 42
	_PO_predicate_tone                                 int = 43
	_quotation_predicate                               int = 44
	_MO_predicate                                      int = 45
	_MO_predicate_tone                                 int = 46
	_terms                                             int = 47
	_terms_2                                           int = 48
	_term                                              int = 49
	_argument                                          int = 50
	_arg_1                                             int = 51
	_arg_2                                             int = 52
	_arg_3                                             int = 53
	_arg_4                                             int = 54
	_arg_5                                             int = 55
	_arg_6                                             int = 56
	_arg_7                                             int = 57
	_serial_argument                                   int = 58
	_coP_arg                                           int = 59
	_co_bar_arg                                        int = 60
	_forethought_coP_arg                               int = 61
	_forethought_coP_arg_1                             int = 62
	_forethought_co_bar_arg                            int = 63
	_coP_pred_arg                                      int = 64
	_forethought_coP_pred_arg                          int = 65
	_forethought_coP_pred_arg_1                        int = 66
	_LU_arg                                            int = 67
	_LU_arg_tone                                       int = 68
	_MI_arg                                            int = 69
	_MI_arg_1                                          int = 70
	_MI_arg_tone                                       int = 71
	_PO_arg                                            int = 72
	_PO_arg_1                                          int = 73
	_PO_arg_tone                                       int = 74
	_quotation_argument                                int = 75
	_MO_argument                                       int = 76
	_MO_argument_tone                                  int = 77
	_relative_clause                                   int = 78
	_relative_clause_1                                 int = 79
	_relative_clause_2                                 int = 80
	_relative_clause_3                                 int = 81
	_relative_predication                              int = 82
	_coP_rel_statement                                 int = 83
	_coP_rel                                           int = 84
	_co_bar_rel                                        int = 85
	_forethought_coP_rel                               int = 86
	_forethought_coP_rel_1                             int = 87
	_forethought_co_bar_rel                            int = 88
	_relative_predicate                                int = 89
	_relative_predicate_1                              int = 90
	_relative_predicate_2                              int = 91
	_relative_predicate_3                              int = 92
	_serial_relative_predicate                         int = 93
	_coP_pred_relative_predicate                       int = 94
	_forethought_coP_pred_relative_predicate           int = 95
	_forethought_coP_pred_relative_predicate_1         int = 96
	_LU_relative                                       int = 97
	_LU_relative_tone                                  int = 98
	_MI_relative_predicate                             int = 99
	_MI_relative_predicate_1                           int = 100
	_MI_relative_predicate_tone                        int = 101
	_PO_relative_predicate                             int = 102
	_PO_relative_predicate_1                           int = 103
	_PO_relative_predicate_tone                        int = 104
	_quotation_relative_predicate                      int = 105
	_MO_relative_predicate                             int = 106
	_MO_relative_predicate_tone                        int = 107
	_termset                                           int = 108
	_termset_II                                        int = 109
	_forethought_coP_term_II                           int = 110
	_forethought_co_bar_term_II                        int = 111
	_termset_III                                       int = 112
	_forethought_coP_term_III                          int = 113
	_forethought_co_bar_term_III                       int = 114
	_termset_IV                                        int = 115
	_forethought_coP_term_IV                           int = 116
	_forethought_co_bar_term_IV                        int = 117
	_termset_V                                         int = 118
	_forethought_coP_term_V                            int = 119
	_forethought_co_bar_term_V                         int = 120
	_forethought_connective                            int = 121
	_terms_II                                          int = 122
	_terms_III                                         int = 123
	_terms_IV                                          int = 124
	_terms_V                                           int = 125
	_adverb                                            int = 126
	_adverb_1                                          int = 127
	_adverb_2                                          int = 128
	_adverb_3                                          int = 129
	_adverb_4                                          int = 130
	_coP_adverb                                        int = 131
	_co_bar_adverb                                     int = 132
	_forethought_coP_adverb                            int = 133
	_forethought_coP_adverb_1                          int = 134
	_forethought_co_bar_adverb                         int = 135
	_serial_adverb                                     int = 136
	_coP_pred_adverb                                   int = 137
	_forethought_coP_pred_adverb                       int = 138
	_forethought_coP_pred_adverb_1                     int = 139
	_LU_adverb                                         int = 140
	_LU_adverb_tone                                    int = 141
	_MI_adverb                                         int = 142
	_MI_adverb_1                                       int = 143
	_MI_adverb_tone                                    int = 144
	_PO_adverb                                         int = 145
	_PO_adverb_1                                       int = 146
	_PO_adverb_tone                                    int = 147
	_quotation_adverb                                  int = 148
	_MO_adverb                                         int = 149
	_MO_adverb_tone                                    int = 150
	_prepositional_phrase                              int = 151
	_prepositional_phrase_1                            int = 152
	_prepositional_phrase_2                            int = 153
	_coP_prepositional_phrase                          int = 154
	_co_bar_prepositional_phrase                       int = 155
	_forethought_coP_prepositional_phrase              int = 156
	_forethought_coP_prepositional_phrase_1            int = 157
	_forethought_co_bar_prepositional_phrase           int = 158
	_preposition                                       int = 159
	_preposition_1                                     int = 160
	_preposition_2                                     int = 161
	_preposition_3                                     int = 162
	_preposition_4                                     int = 163
	_coP_preposition                                   int = 164
	_co_bar_preposition                                int = 165
	_forethought_coP_preposition                       int = 166
	_forethought_coP_preposition_1                     int = 167
	_forethought_co_bar_preposition                    int = 168
	_serial_preposition                                int = 169
	_coP_pred_preposition                              int = 170
	_forethought_coP_pred_preposition                  int = 171
	_forethought_coP_pred_preposition_1                int = 172
	_LU_preposition                                    int = 173
	_LU_preposition_tone                               int = 174
	_MI_preposition                                    int = 175
	_MI_preposition_1                                  int = 176
	_MI_preposition_tone                               int = 177
	_PO_preposition                                    int = 178
	_PO_preposition_1                                  int = 179
	_PO_preposition_tone                               int = 180
	_quotation_preposition                             int = 181
	_MO_preposition                                    int = 182
	_MO_preposition_tone                               int = 183
	_content_clause                                    int = 184
	_content_clause_1                                  int = 185
	_content_predication                               int = 186
	_coP_content_statement                             int = 187
	_content_predicate                                 int = 188
	_content_predicate_1                               int = 189
	_content_predicate_2                               int = 190
	_content_predicate_3                               int = 191
	_serial_content_predicate                          int = 192
	_coP_pred_content_predicate                        int = 193
	_forethought_coP_pred_content_predicate            int = 194
	_forethought_coP_pred_content_predicate_1          int = 195
	_LU_content                                        int = 196
	_LU_content_tone                                   int = 197
	_MI_content_predicate                              int = 198
	_MI_content_predicate_1                            int = 199
	_MI_content_tone                                   int = 200
	_PO_content_predicate                              int = 201
	_PO_content_predicate_1                            int = 202
	_PO_content_tone                                   int = 203
	_quotation_content_predicate                       int = 204
	_MO_content_predicate                              int = 205
	_MO_content_predicate_tone                         int = 206
	_freemod                                           int = 207
	_parenthetical                                     int = 208
	_parenthetical_1                                   int = 209
	_incidental                                        int = 210
	_vocative                                          int = 211
	_prefix                                            int = 212
	_focus                                             int = 213
	_end_quote                                         int = 214
	_end_predicatizer                                  int = 215
	_end_statement                                     int = 216
	_sentence_prefix                                   int = 217
	_end_prenex                                        int = 218
	_start_incidental                                  int = 219
	_start_parenthetical                               int = 220
	_end_parenthetical                                 int = 221
	_vocative_marker                                   int = 222
	_linking_word                                      int = 223
	_connective                                        int = 224
	_illocutionary                                     int = 225
	_quantifier                                        int = 226
	_interjection                                      int = 227
	_forethought_marker                                int = 228
	_function_word                                     int = 229
	_BI                                                int = 230
	_DA                                                int = 231
	_GA                                                int = 232
	_GO                                                int = 233
	_HA                                                int = 234
	_HU                                                int = 235
	_JU                                                int = 236
	_KU                                                int = 237
	_KI                                                int = 238
	_KIO                                               int = 239
	_KEO                                               int = 240
	_LU                                                int = 241
	_MU                                                int = 242
	_MI                                                int = 243
	_MO                                                int = 244
	_NA                                                int = 245
	_PO                                                int = 246
	_RA                                                int = 247
	_TU                                                int = 248
	_TO                                                int = 249
	_TEO                                               int = 250
	_syllable__compound_desinence__compound_tone       int = 251
	_syllable__arg_desinence__arg_tone                 int = 252
	_syllable__relative_desinence__relative_tone       int = 253
	_syllable__verb_desinence__verb_tone               int = 254
	_syllable__content_desinence__content_tone         int = 255
	_syllable__preposition_desinence__preposition_tone int = 256
	_syllable__adverb_desinence__adverb_tone           int = 257
	_neutral_syllable                                  int = 258
	_compound_syllable                                 int = 259
	_arg_syllable                                      int = 260
	_relative_syllable                                 int = 261
	_verb_syllable                                     int = 262
	_content_syllable                                  int = 263
	_preposition_syllable                              int = 264
	_adverb_syllable                                   int = 265
	_boundary                                          int = 266
	_initial                                           int = 267
	_desinence__a__u__i__o__e                          int = 268
	_desinence__ā__ū__ī__ō__ē                          int = 269
	_desinence__á__ú__í__ó__é                          int = 270
	_desinence__ǎ__ǔ__ǐ__ǒ__ě                          int = 271
	_desinence__ả__ủ__ỉ__ỏ__ẻ                          int = 272
	_desinence__â__û__î__ô__ê                          int = 273
	_desinence__à__ù__ì__ò__è                          int = 274
	_desinence__ã__ũ__ĩ__õ__ẽ                          int = 275
	_neutral_desinence                                 int = 276
	_compound_desinence                                int = 277
	_arg_desinence                                     int = 278
	_relative_desinence                                int = 279
	_verb_desinence                                    int = 280
	_content_desinence                                 int = 281
	_preposition_desinence                             int = 282
	_adverb_desinence                                  int = 283
	_tone                                              int = 284
	_A                                                 int = 285
	_U                                                 int = 286
	_I                                                 int = 287
	_O                                                 int = 288
	_E                                                 int = 289
	_ā                                                 int = 290
	_ū                                                 int = 291
	_ī                                                 int = 292
	_ō                                                 int = 293
	_ē                                                 int = 294
	_macron_combiner                                   int = 295
	_compound_tone                                     int = 296
	_á                                                 int = 297
	_ú                                                 int = 298
	_í                                                 int = 299
	_ó                                                 int = 300
	_é                                                 int = 301
	_acute_combiner                                    int = 302
	_arg_tone                                          int = 303
	_ǎ                                                 int = 304
	_ǔ                                                 int = 305
	_ǐ                                                 int = 306
	_ǒ                                                 int = 307
	_ě                                                 int = 308
	_caron_combiner                                    int = 309
	_breve_combiner                                    int = 310
	_relative_tone                                     int = 311
	_ả                                                 int = 312
	_ủ                                                 int = 313
	_ỉ                                                 int = 314
	_ỏ                                                 int = 315
	_ẻ                                                 int = 316
	_hook_combiner                                     int = 317
	_verb_tone                                         int = 318
	_â                                                 int = 319
	_û                                                 int = 320
	_î                                                 int = 321
	_ô                                                 int = 322
	_ê                                                 int = 323
	_circumflex_combiner                               int = 324
	_content_tone                                      int = 325
	_à                                                 int = 326
	_ù                                                 int = 327
	_ì                                                 int = 328
	_ò                                                 int = 329
	_è                                                 int = 330
	_grave_combiner                                    int = 331
	_preposition_tone                                  int = 332
	_ã                                                 int = 333
	_ũ                                                 int = 334
	_ĩ                                                 int = 335
	_õ                                                 int = 336
	_ẽ                                                 int = 337
	_tilde_combiner                                    int = 338
	_adverb_tone                                       int = 339
	_a                                                 int = 340
	_b                                                 int = 341
	_c                                                 int = 342
	_d                                                 int = 343
	_e                                                 int = 344
	_f                                                 int = 345
	_g                                                 int = 346
	_h                                                 int = 347
	_i                                                 int = 348
	_j                                                 int = 349
	_k                                                 int = 350
	_l                                                 int = 351
	_m                                                 int = 352
	_n                                                 int = 353
	_o                                                 int = 354
	_p                                                 int = 355
	_q                                                 int = 356
	_r                                                 int = 357
	_s                                                 int = 358
	_t                                                 int = 359
	_u                                                 int = 360
	_w                                                 int = 361
	_y                                                 int = 362
	_spaces                                            int = 363
	_EOF                                               int = 364

	_N int = 365
)

type _Parser struct {
	text     string
	deltaPos [][_N]int32
	deltaErr [][_N]int32
	node     map[_key]*peg.Node
	fail     map[_key]*peg.Fail
	act      map[_key]interface{}
	lastFail int
	data     interface{}
}

type _key struct {
	start int
	rule  int
}

func _NewParser(text string) *_Parser {
	return &_Parser{
		text:     text,
		deltaPos: make([][_N]int32, len(text)+1),
		deltaErr: make([][_N]int32, len(text)+1),
		node:     make(map[_key]*peg.Node),
		fail:     make(map[_key]*peg.Fail),
		act:      make(map[_key]interface{}),
	}
}

func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _memoize(parser *_Parser, rule, start, pos, perr int) (int, int) {
	parser.lastFail = perr
	derr := perr - start
	parser.deltaErr[start][rule] = int32(derr + 1)
	if pos >= 0 {
		dpos := pos - start
		parser.deltaPos[start][rule] = int32(dpos + 1)
		return dpos, derr
	}
	parser.deltaPos[start][rule] = -1
	return -1, derr
}

func _memo(parser *_Parser, rule, start int) (int, int, bool) {
	dp := parser.deltaPos[start][rule]
	if dp == 0 {
		return 0, 0, false
	}
	if dp > 0 {
		dp--
	}
	de := parser.deltaErr[start][rule] - 1
	return int(dp), int(de), true
}

func _failMemo(parser *_Parser, rule, start, errPos int) (int, *peg.Fail) {
	if start > parser.lastFail {
		return -1, &peg.Fail{}
	}
	dp := parser.deltaPos[start][rule]
	de := parser.deltaErr[start][rule]
	if start+int(de-1) < errPos {
		if dp > 0 {
			return start + int(dp-1), &peg.Fail{}
		}
		return -1, &peg.Fail{}
	}
	f := parser.fail[_key{start: start, rule: rule}]
	if dp < 0 && f != nil {
		return -1, f
	}
	if dp > 0 && f != nil {
		return start + int(dp-1), f
	}
	return start, nil
}

func _accept(parser *_Parser, f func(*_Parser, int) (int, int), pos, perr *int) bool {
	dp, de := f(parser, *pos)
	*perr = _max(*perr, *pos+de)
	if dp < 0 {
		return false
	}
	*pos += dp
	return true
}

func _node(parser *_Parser, f func(*_Parser, int) (int, *peg.Node), node *peg.Node, pos *int) bool {
	p, kid := f(parser, *pos)
	if kid == nil {
		return false
	}
	node.Kids = append(node.Kids, kid)
	*pos = p
	return true
}

func _fail(parser *_Parser, f func(*_Parser, int, int) (int, *peg.Fail), errPos int, node *peg.Fail, pos *int) bool {
	p, kid := f(parser, *pos, errPos)
	if kid.Want != "" || len(kid.Kids) > 0 {
		node.Kids = append(node.Kids, kid)
	}
	if p < 0 {
		return false
	}
	*pos = p
	return true
}

func _next(parser *_Parser, pos int) (rune, int) {
	r, w := peg.DecodeRuneInString(parser.text[pos:])
	return r, w
}

func _sub(parser *_Parser, start, end int, kids []*peg.Node) *peg.Node {
	node := &peg.Node{
		Text: parser.text[start:end],
		Kids: make([]*peg.Node, len(kids)),
	}
	copy(node.Kids, kids)
	return node
}

func _leaf(parser *_Parser, start, end int) *peg.Node {
	return &peg.Node{Text: parser.text[start:end]}
}

// A no-op function to mark a variable as used.
func use(interface{}) {}

func _full_textAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _full_text, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// text EOF
	// text
	if !_accept(parser, _textAccepts, &pos, &perr) {
		goto fail
	}
	// EOF
	if !_accept(parser, _EOFAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _full_text, start, pos, perr)
fail:
	return _memoize(parser, _full_text, start, -1, perr)
}

func _full_textNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_full_text]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _full_text}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "full_text"}
	// text EOF
	// text
	if !_node(parser, _textNode, node, &pos) {
		goto fail
	}
	// EOF
	if !_node(parser, _EOFNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _full_textFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _full_text, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "full_text",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _full_text}
	// text EOF
	// text
	if !_fail(parser, _textFail, errPos, failure, &pos) {
		goto fail
	}
	// EOF
	if !_fail(parser, _EOFFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _full_textAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_full_text]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _full_text}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// text EOF
	{
		var node0 string
		// text
		if p, n := _textAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// EOF
		if p, n := _EOFAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _textAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _text, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// spaces? freemod? spaces? discourse? spaces? EOF?
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// freemod?
	{
		pos6 := pos
		// freemod
		if !_accept(parser, _freemodAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// spaces?
	{
		pos10 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok12
	fail11:
		pos = pos10
	ok12:
	}
	// discourse?
	{
		pos14 := pos
		// discourse
		if !_accept(parser, _discourseAccepts, &pos, &perr) {
			goto fail15
		}
		goto ok16
	fail15:
		pos = pos14
	ok16:
	}
	// spaces?
	{
		pos18 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail19
		}
		goto ok20
	fail19:
		pos = pos18
	ok20:
	}
	// EOF?
	{
		pos22 := pos
		// EOF
		if !_accept(parser, _EOFAccepts, &pos, &perr) {
			goto fail23
		}
		goto ok24
	fail23:
		pos = pos22
	ok24:
	}
	return _memoize(parser, _text, start, pos, perr)
}

func _textNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_text]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _text}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "text"}
	// spaces? freemod? spaces? discourse? spaces? EOF?
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// freemod?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// freemod
		if !_node(parser, _freemodNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// spaces?
	{
		nkids9 := len(node.Kids)
		pos10 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail11
		}
		goto ok12
	fail11:
		node.Kids = node.Kids[:nkids9]
		pos = pos10
	ok12:
	}
	// discourse?
	{
		nkids13 := len(node.Kids)
		pos14 := pos
		// discourse
		if !_node(parser, _discourseNode, node, &pos) {
			goto fail15
		}
		goto ok16
	fail15:
		node.Kids = node.Kids[:nkids13]
		pos = pos14
	ok16:
	}
	// spaces?
	{
		nkids17 := len(node.Kids)
		pos18 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail19
		}
		goto ok20
	fail19:
		node.Kids = node.Kids[:nkids17]
		pos = pos18
	ok20:
	}
	// EOF?
	{
		nkids21 := len(node.Kids)
		pos22 := pos
		// EOF
		if !_node(parser, _EOFNode, node, &pos) {
			goto fail23
		}
		goto ok24
	fail23:
		node.Kids = node.Kids[:nkids21]
		pos = pos22
	ok24:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
}

func _textFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _text, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "text",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _text}
	// spaces? freemod? spaces? discourse? spaces? EOF?
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// freemod?
	{
		pos6 := pos
		// freemod
		if !_fail(parser, _freemodFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// spaces?
	{
		pos10 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok12
	fail11:
		pos = pos10
	ok12:
	}
	// discourse?
	{
		pos14 := pos
		// discourse
		if !_fail(parser, _discourseFail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok16
	fail15:
		pos = pos14
	ok16:
	}
	// spaces?
	{
		pos18 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail19
		}
		goto ok20
	fail19:
		pos = pos18
	ok20:
	}
	// EOF?
	{
		pos22 := pos
		// EOF
		if !_fail(parser, _EOFFail, errPos, failure, &pos) {
			goto fail23
		}
		goto ok24
	fail23:
		pos = pos22
	ok24:
	}
	parser.fail[key] = failure
	return pos, failure
}

func _textAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_text]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _text}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// spaces? freemod? spaces? discourse? spaces? EOF?
	{
		var node0 string
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// freemod?
		{
			pos6 := pos
			// freemod
			if p, n := _freemodAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos10 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok12
		fail11:
			node0 = ""
			pos = pos10
		ok12:
		}
		node, node0 = node+node0, ""
		// discourse?
		{
			pos14 := pos
			// discourse
			if p, n := _discourseAction(parser, pos); n == nil {
				goto fail15
			} else {
				node0 = *n
				pos = p
			}
			goto ok16
		fail15:
			node0 = ""
			pos = pos14
		ok16:
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos18 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail19
			} else {
				node0 = *n
				pos = p
			}
			goto ok20
		fail19:
			node0 = ""
			pos = pos18
		ok20:
		}
		node, node0 = node+node0, ""
		// EOF?
		{
			pos22 := pos
			// EOF
			if p, n := _EOFAction(parser, pos); n == nil {
				goto fail23
			} else {
				node0 = *n
				pos = p
			}
			goto ok24
		fail23:
			node0 = ""
			pos = pos22
		ok24:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
}

func _discourseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _discourse, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (sentence/fragment)+
	// (sentence/fragment)
	// sentence/fragment
	{
		pos7 := pos
		// sentence
		if !_accept(parser, _sentenceAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok4
	fail8:
		pos = pos7
		// fragment
		if !_accept(parser, _fragmentAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok4
	fail9:
		pos = pos7
		goto fail
	ok4:
	}
	for {
		pos1 := pos
		// (sentence/fragment)
		// sentence/fragment
		{
			pos13 := pos
			// sentence
			if !_accept(parser, _sentenceAccepts, &pos, &perr) {
				goto fail14
			}
			goto ok10
		fail14:
			pos = pos13
			// fragment
			if !_accept(parser, _fragmentAccepts, &pos, &perr) {
				goto fail15
			}
			goto ok10
		fail15:
			pos = pos13
			goto fail3
		ok10:
		}
		continue
	fail3:
		pos = pos1
		break
	}
	return _memoize(parser, _discourse, start, pos, perr)
fail:
	return _memoize(parser, _discourse, start, -1, perr)
}

func _discourseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_discourse]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _discourse}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "discourse"}
	// (sentence/fragment)+
	// (sentence/fragment)
	{
		nkids4 := len(node.Kids)
		pos05 := pos
		// sentence/fragment
		{
			pos9 := pos
			nkids7 := len(node.Kids)
			// sentence
			if !_node(parser, _sentenceNode, node, &pos) {
				goto fail10
			}
			goto ok6
		fail10:
			node.Kids = node.Kids[:nkids7]
			pos = pos9
			// fragment
			if !_node(parser, _fragmentNode, node, &pos) {
				goto fail11
			}
			goto ok6
		fail11:
			node.Kids = node.Kids[:nkids7]
			pos = pos9
			goto fail
		ok6:
		}
		sub := _sub(parser, pos05, pos, node.Kids[nkids4:])
		node.Kids = append(node.Kids[:nkids4], sub)
	}
	for {
		nkids0 := len(node.Kids)
		pos1 := pos
		// (sentence/fragment)
		{
			nkids12 := len(node.Kids)
			pos013 := pos
			// sentence/fragment
			{
				pos17 := pos
				nkids15 := len(node.Kids)
				// sentence
				if !_node(parser, _sentenceNode, node, &pos) {
					goto fail18
				}
				goto ok14
			fail18:
				node.Kids = node.Kids[:nkids15]
				pos = pos17
				// fragment
				if !_node(parser, _fragmentNode, node, &pos) {
					goto fail19
				}
				goto ok14
			fail19:
				node.Kids = node.Kids[:nkids15]
				pos = pos17
				goto fail3
			ok14:
			}
			sub := _sub(parser, pos013, pos, node.Kids[nkids12:])
			node.Kids = append(node.Kids[:nkids12], sub)
		}
		continue
	fail3:
		node.Kids = node.Kids[:nkids0]
		pos = pos1
		break
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _discourseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _discourse, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "discourse",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _discourse}
	// (sentence/fragment)+
	// (sentence/fragment)
	// sentence/fragment
	{
		pos7 := pos
		// sentence
		if !_fail(parser, _sentenceFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok4
	fail8:
		pos = pos7
		// fragment
		if !_fail(parser, _fragmentFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok4
	fail9:
		pos = pos7
		goto fail
	ok4:
	}
	for {
		pos1 := pos
		// (sentence/fragment)
		// sentence/fragment
		{
			pos13 := pos
			// sentence
			if !_fail(parser, _sentenceFail, errPos, failure, &pos) {
				goto fail14
			}
			goto ok10
		fail14:
			pos = pos13
			// fragment
			if !_fail(parser, _fragmentFail, errPos, failure, &pos) {
				goto fail15
			}
			goto ok10
		fail15:
			pos = pos13
			goto fail3
		ok10:
		}
		continue
	fail3:
		pos = pos1
		break
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _discourseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_discourse]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _discourse}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (sentence/fragment)+
	{
		var node2 string
		// (sentence/fragment)
		// sentence/fragment
		{
			pos7 := pos
			var node6 string
			// sentence
			if p, n := _sentenceAction(parser, pos); n == nil {
				goto fail8
			} else {
				node2 = *n
				pos = p
			}
			goto ok4
		fail8:
			node2 = node6
			pos = pos7
			// fragment
			if p, n := _fragmentAction(parser, pos); n == nil {
				goto fail9
			} else {
				node2 = *n
				pos = p
			}
			goto ok4
		fail9:
			node2 = node6
			pos = pos7
			goto fail
		ok4:
		}
		node += node2
	}
	for {
		pos1 := pos
		var node2 string
		// (sentence/fragment)
		// sentence/fragment
		{
			pos13 := pos
			var node12 string
			// sentence
			if p, n := _sentenceAction(parser, pos); n == nil {
				goto fail14
			} else {
				node2 = *n
				pos = p
			}
			goto ok10
		fail14:
			node2 = node12
			pos = pos13
			// fragment
			if p, n := _fragmentAction(parser, pos); n == nil {
				goto fail15
			} else {
				node2 = *n
				pos = p
			}
			goto ok10
		fail15:
			node2 = node12
			pos = pos13
			goto fail3
		ok10:
		}
		node += node2
		continue
	fail3:
		pos = pos1
		break
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _sentence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_sentence/sentence_1
	{
		pos3 := pos
		// coP_sentence
		if !_accept(parser, _coP_sentenceAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// sentence_1
		if !_accept(parser, _sentence_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _sentence, start, pos, perr)
fail:
	return _memoize(parser, _sentence, start, -1, perr)
}

func _sentenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "sentence"}
	// coP_sentence/sentence_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_sentence
		if !_node(parser, _coP_sentenceNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// sentence_1
		if !_node(parser, _sentence_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _sentence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence}
	// coP_sentence/sentence_1
	{
		pos3 := pos
		// coP_sentence
		if !_fail(parser, _coP_sentenceFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// sentence_1
		if !_fail(parser, _sentence_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_sentence/sentence_1
	{
		pos3 := pos
		var node2 string
		// coP_sentence
		if p, n := _coP_sentenceAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// sentence_1
		if p, n := _sentence_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentence_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _sentence_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_sentence/sentence_2
	{
		pos3 := pos
		// forethought_coP_sentence
		if !_accept(parser, _forethought_coP_sentenceAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// sentence_2
		if !_accept(parser, _sentence_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _sentence_1, start, pos, perr)
fail:
	return _memoize(parser, _sentence_1, start, -1, perr)
}

func _sentence_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_sentence_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "sentence_1"}
	// forethought_coP_sentence/sentence_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_sentence
		if !_node(parser, _forethought_coP_sentenceNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// sentence_2
		if !_node(parser, _sentence_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _sentence_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_1}
	// forethought_coP_sentence/sentence_2
	{
		pos3 := pos
		// forethought_coP_sentence
		if !_fail(parser, _forethought_coP_sentenceFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// sentence_2
		if !_fail(parser, _sentence_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_sentence_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_sentence/sentence_2
	{
		pos3 := pos
		var node2 string
		// forethought_coP_sentence
		if p, n := _forethought_coP_sentenceAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// sentence_2
		if p, n := _sentence_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentence_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _sentence_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// spaces? sentence_prefix? sentence_3
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// sentence_prefix?
	{
		pos6 := pos
		// sentence_prefix
		if !_accept(parser, _sentence_prefixAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// sentence_3
	if !_accept(parser, _sentence_3Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _sentence_2, start, pos, perr)
fail:
	return _memoize(parser, _sentence_2, start, -1, perr)
}

func _sentence_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_sentence_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "sentence_2"}
	// spaces? sentence_prefix? sentence_3
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// sentence_prefix?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// sentence_prefix
		if !_node(parser, _sentence_prefixNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// sentence_3
	if !_node(parser, _sentence_3Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _sentence_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_2}
	// spaces? sentence_prefix? sentence_3
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// sentence_prefix?
	{
		pos6 := pos
		// sentence_prefix
		if !_fail(parser, _sentence_prefixFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// sentence_3
	if !_fail(parser, _sentence_3Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_sentence_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// spaces? sentence_prefix? sentence_3
	{
		var node0 string
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// sentence_prefix?
		{
			pos6 := pos
			// sentence_prefix
			if p, n := _sentence_prefixAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// sentence_3
		if p, n := _sentence_3Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentence_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _sentence_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// spaces? statement spaces? illocutionary?
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// illocutionary?
	{
		pos10 := pos
		// illocutionary
		if !_accept(parser, _illocutionaryAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok12
	fail11:
		pos = pos10
	ok12:
	}
	return _memoize(parser, _sentence_3, start, pos, perr)
fail:
	return _memoize(parser, _sentence_3, start, -1, perr)
}

func _sentence_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_sentence_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "sentence_3"}
	// spaces? statement spaces? illocutionary?
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// illocutionary?
	{
		nkids9 := len(node.Kids)
		pos10 := pos
		// illocutionary
		if !_node(parser, _illocutionaryNode, node, &pos) {
			goto fail11
		}
		goto ok12
	fail11:
		node.Kids = node.Kids[:nkids9]
		pos = pos10
	ok12:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _sentence_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_3}
	// spaces? statement spaces? illocutionary?
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// illocutionary?
	{
		pos10 := pos
		// illocutionary
		if !_fail(parser, _illocutionaryFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok12
	fail11:
		pos = pos10
	ok12:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_sentence_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// spaces? statement spaces? illocutionary?
	{
		var node0 string
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos6 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// illocutionary?
		{
			pos10 := pos
			// illocutionary
			if p, n := _illocutionaryAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok12
		fail11:
			node0 = ""
			pos = pos10
		ok12:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _fragmentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _fragment, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// spaces? coP_rel/spaces? relative_clause/spaces? prenex/spaces? terms
	{
		pos3 := pos
		// spaces? coP_rel
		// spaces?
		{
			pos7 := pos
			// spaces
			if !_accept(parser, _spacesAccepts, &pos, &perr) {
				goto fail8
			}
			goto ok9
		fail8:
			pos = pos7
		ok9:
		}
		// coP_rel
		if !_accept(parser, _coP_relAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// spaces? relative_clause
		// spaces?
		{
			pos13 := pos
			// spaces
			if !_accept(parser, _spacesAccepts, &pos, &perr) {
				goto fail14
			}
			goto ok15
		fail14:
			pos = pos13
		ok15:
		}
		// relative_clause
		if !_accept(parser, _relative_clauseAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// spaces? prenex
		// spaces?
		{
			pos19 := pos
			// spaces
			if !_accept(parser, _spacesAccepts, &pos, &perr) {
				goto fail20
			}
			goto ok21
		fail20:
			pos = pos19
		ok21:
		}
		// prenex
		if !_accept(parser, _prenexAccepts, &pos, &perr) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		// spaces? terms
		// spaces?
		{
			pos25 := pos
			// spaces
			if !_accept(parser, _spacesAccepts, &pos, &perr) {
				goto fail26
			}
			goto ok27
		fail26:
			pos = pos25
		ok27:
		}
		// terms
		if !_accept(parser, _termsAccepts, &pos, &perr) {
			goto fail22
		}
		goto ok0
	fail22:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _fragment, start, pos, perr)
fail:
	return _memoize(parser, _fragment, start, -1, perr)
}

func _fragmentNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_fragment]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _fragment}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "fragment"}
	// spaces? coP_rel/spaces? relative_clause/spaces? prenex/spaces? terms
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// spaces? coP_rel
		// spaces?
		{
			nkids6 := len(node.Kids)
			pos7 := pos
			// spaces
			if !_node(parser, _spacesNode, node, &pos) {
				goto fail8
			}
			goto ok9
		fail8:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
		ok9:
		}
		// coP_rel
		if !_node(parser, _coP_relNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// spaces? relative_clause
		// spaces?
		{
			nkids12 := len(node.Kids)
			pos13 := pos
			// spaces
			if !_node(parser, _spacesNode, node, &pos) {
				goto fail14
			}
			goto ok15
		fail14:
			node.Kids = node.Kids[:nkids12]
			pos = pos13
		ok15:
		}
		// relative_clause
		if !_node(parser, _relative_clauseNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// spaces? prenex
		// spaces?
		{
			nkids18 := len(node.Kids)
			pos19 := pos
			// spaces
			if !_node(parser, _spacesNode, node, &pos) {
				goto fail20
			}
			goto ok21
		fail20:
			node.Kids = node.Kids[:nkids18]
			pos = pos19
		ok21:
		}
		// prenex
		if !_node(parser, _prenexNode, node, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// spaces? terms
		// spaces?
		{
			nkids24 := len(node.Kids)
			pos25 := pos
			// spaces
			if !_node(parser, _spacesNode, node, &pos) {
				goto fail26
			}
			goto ok27
		fail26:
			node.Kids = node.Kids[:nkids24]
			pos = pos25
		ok27:
		}
		// terms
		if !_node(parser, _termsNode, node, &pos) {
			goto fail22
		}
		goto ok0
	fail22:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _fragmentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _fragment, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "fragment",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _fragment}
	// spaces? coP_rel/spaces? relative_clause/spaces? prenex/spaces? terms
	{
		pos3 := pos
		// spaces? coP_rel
		// spaces?
		{
			pos7 := pos
			// spaces
			if !_fail(parser, _spacesFail, errPos, failure, &pos) {
				goto fail8
			}
			goto ok9
		fail8:
			pos = pos7
		ok9:
		}
		// coP_rel
		if !_fail(parser, _coP_relFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// spaces? relative_clause
		// spaces?
		{
			pos13 := pos
			// spaces
			if !_fail(parser, _spacesFail, errPos, failure, &pos) {
				goto fail14
			}
			goto ok15
		fail14:
			pos = pos13
		ok15:
		}
		// relative_clause
		if !_fail(parser, _relative_clauseFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// spaces? prenex
		// spaces?
		{
			pos19 := pos
			// spaces
			if !_fail(parser, _spacesFail, errPos, failure, &pos) {
				goto fail20
			}
			goto ok21
		fail20:
			pos = pos19
		ok21:
		}
		// prenex
		if !_fail(parser, _prenexFail, errPos, failure, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		// spaces? terms
		// spaces?
		{
			pos25 := pos
			// spaces
			if !_fail(parser, _spacesFail, errPos, failure, &pos) {
				goto fail26
			}
			goto ok27
		fail26:
			pos = pos25
		ok27:
		}
		// terms
		if !_fail(parser, _termsFail, errPos, failure, &pos) {
			goto fail22
		}
		goto ok0
	fail22:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _fragmentAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_fragment]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _fragment}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// spaces? coP_rel/spaces? relative_clause/spaces? prenex/spaces? terms
	{
		pos3 := pos
		var node2 string
		// spaces? coP_rel
		{
			var node5 string
			// spaces?
			{
				pos7 := pos
				// spaces
				if p, n := _spacesAction(parser, pos); n == nil {
					goto fail8
				} else {
					node5 = *n
					pos = p
				}
				goto ok9
			fail8:
				node5 = ""
				pos = pos7
			ok9:
			}
			node, node5 = node+node5, ""
			// coP_rel
			if p, n := _coP_relAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// spaces? relative_clause
		{
			var node11 string
			// spaces?
			{
				pos13 := pos
				// spaces
				if p, n := _spacesAction(parser, pos); n == nil {
					goto fail14
				} else {
					node11 = *n
					pos = p
				}
				goto ok15
			fail14:
				node11 = ""
				pos = pos13
			ok15:
			}
			node, node11 = node+node11, ""
			// relative_clause
			if p, n := _relative_clauseAction(parser, pos); n == nil {
				goto fail10
			} else {
				node11 = *n
				pos = p
			}
			node, node11 = node+node11, ""
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// spaces? prenex
		{
			var node17 string
			// spaces?
			{
				pos19 := pos
				// spaces
				if p, n := _spacesAction(parser, pos); n == nil {
					goto fail20
				} else {
					node17 = *n
					pos = p
				}
				goto ok21
			fail20:
				node17 = ""
				pos = pos19
			ok21:
			}
			node, node17 = node+node17, ""
			// prenex
			if p, n := _prenexAction(parser, pos); n == nil {
				goto fail16
			} else {
				node17 = *n
				pos = p
			}
			node, node17 = node+node17, ""
		}
		goto ok0
	fail16:
		node = node2
		pos = pos3
		// spaces? terms
		{
			var node23 string
			// spaces?
			{
				pos25 := pos
				// spaces
				if p, n := _spacesAction(parser, pos); n == nil {
					goto fail26
				} else {
					node23 = *n
					pos = p
				}
				goto ok27
			fail26:
				node23 = ""
				pos = pos25
			ok27:
			}
			node, node23 = node+node23, ""
			// terms
			if p, n := _termsAction(parser, pos); n == nil {
				goto fail22
			} else {
				node23 = *n
				pos = p
			}
			node, node23 = node+node23, ""
		}
		goto ok0
	fail22:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_sentenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_sentence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// sentence_1 spaces? co_bar_sentence
	// sentence_1
	if !_accept(parser, _sentence_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_sentence
	if !_accept(parser, _co_bar_sentenceAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_sentence, start, pos, perr)
fail:
	return _memoize(parser, _coP_sentence, start, -1, perr)
}

func _coP_sentenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_sentence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_sentence"}
	// sentence_1 spaces? co_bar_sentence
	// sentence_1
	if !_node(parser, _sentence_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_sentence
	if !_node(parser, _co_bar_sentenceNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_sentenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_sentence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_sentence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_sentence}
	// sentence_1 spaces? co_bar_sentence
	// sentence_1
	if !_fail(parser, _sentence_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_sentence
	if !_fail(parser, _co_bar_sentenceFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_sentenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_sentence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// sentence_1 spaces? co_bar_sentence
	{
		var node0 string
		// sentence_1
		if p, n := _sentence_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_sentence
		if p, n := _co_bar_sentenceAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_sentenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_sentence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? sentence_1
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// sentence_1
	if !_accept(parser, _sentence_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_sentence, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_sentence, start, -1, perr)
}

func _co_bar_sentenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_sentence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_sentence"}
	// connective spaces? sentence_1
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// sentence_1
	if !_node(parser, _sentence_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_sentenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_sentence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_sentence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_sentence}
	// connective spaces? sentence_1
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// sentence_1
	if !_fail(parser, _sentence_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_sentenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_sentence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? sentence_1
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// sentence_1
		if p, n := _sentence_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_sentenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_sentence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_sentence_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_sentence_1
	if !_accept(parser, _forethought_coP_sentence_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_sentence, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_sentence, start, -1, perr)
}

func _forethought_coP_sentenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_sentence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_sentence"}
	// forethought_connective spaces? forethought_coP_sentence_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_sentence_1
	if !_node(parser, _forethought_coP_sentence_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_sentenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_sentence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_sentence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_sentence}
	// forethought_connective spaces? forethought_coP_sentence_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_sentence_1
	if !_fail(parser, _forethought_coP_sentence_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_sentenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_sentence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_sentence_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_sentence_1
		if p, n := _forethought_coP_sentence_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_sentence_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_sentence_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// sentence spaces? forethought_co_bar_sentence
	// sentence
	if !_accept(parser, _sentenceAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_sentence
	if !_accept(parser, _forethought_co_bar_sentenceAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_sentence_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_sentence_1, start, -1, perr)
}

func _forethought_coP_sentence_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_sentence_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_sentence_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_sentence_1"}
	// sentence spaces? forethought_co_bar_sentence
	// sentence
	if !_node(parser, _sentenceNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_sentence
	if !_node(parser, _forethought_co_bar_sentenceNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_sentence_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_sentence_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_sentence_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_sentence_1}
	// sentence spaces? forethought_co_bar_sentence
	// sentence
	if !_fail(parser, _sentenceFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_sentence
	if !_fail(parser, _forethought_co_bar_sentenceFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_sentence_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_sentence_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_sentence_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// sentence spaces? forethought_co_bar_sentence
	{
		var node0 string
		// sentence
		if p, n := _sentenceAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_sentence
		if p, n := _forethought_co_bar_sentenceAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_sentenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_sentence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? sentence
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// sentence
	if !_accept(parser, _sentenceAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_sentence, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_sentence, start, -1, perr)
}

func _forethought_co_bar_sentenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_sentence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_sentence"}
	// forethought_marker spaces? sentence
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// sentence
	if !_node(parser, _sentenceNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_sentenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_sentence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_sentence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_sentence}
	// forethought_marker spaces? sentence
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// sentence
	if !_fail(parser, _sentenceFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_sentenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_sentence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? sentence
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// sentence
		if p, n := _sentenceAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_statement/statement_1
	{
		pos3 := pos
		// coP_statement
		if !_accept(parser, _coP_statementAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// statement_1
		if !_accept(parser, _statement_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _statement, start, pos, perr)
fail:
	return _memoize(parser, _statement, start, -1, perr)
}

func _statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "statement"}
	// coP_statement/statement_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_statement
		if !_node(parser, _coP_statementNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// statement_1
		if !_node(parser, _statement_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement}
	// coP_statement/statement_1
	{
		pos3 := pos
		// coP_statement
		if !_fail(parser, _coP_statementFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// statement_1
		if !_fail(parser, _statement_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_statement/statement_1
	{
		pos3 := pos
		var node2 string
		// coP_statement
		if p, n := _coP_statementAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// statement_1
		if p, n := _statement_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _statement_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _statement_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_statement/statement_2
	{
		pos3 := pos
		// forethought_coP_statement
		if !_accept(parser, _forethought_coP_statementAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// statement_2
		if !_accept(parser, _statement_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _statement_1, start, pos, perr)
fail:
	return _memoize(parser, _statement_1, start, -1, perr)
}

func _statement_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_statement_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "statement_1"}
	// forethought_coP_statement/statement_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_statement
		if !_node(parser, _forethought_coP_statementNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// statement_2
		if !_node(parser, _statement_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _statement_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _statement_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement_1}
	// forethought_coP_statement/statement_2
	{
		pos3 := pos
		// forethought_coP_statement
		if !_fail(parser, _forethought_coP_statementFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// statement_2
		if !_fail(parser, _statement_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _statement_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_statement_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_statement/statement_2
	{
		pos3 := pos
		var node2 string
		// forethought_coP_statement
		if p, n := _forethought_coP_statementAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// statement_2
		if p, n := _statement_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _statement_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _statement_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// statement_3 end_statement?
	// statement_3
	if !_accept(parser, _statement_3Accepts, &pos, &perr) {
		goto fail
	}
	// end_statement?
	{
		pos2 := pos
		// end_statement
		if !_accept(parser, _end_statementAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	return _memoize(parser, _statement_2, start, pos, perr)
fail:
	return _memoize(parser, _statement_2, start, -1, perr)
}

func _statement_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_statement_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "statement_2"}
	// statement_3 end_statement?
	// statement_3
	if !_node(parser, _statement_3Node, node, &pos) {
		goto fail
	}
	// end_statement?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// end_statement
		if !_node(parser, _end_statementNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _statement_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _statement_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement_2}
	// statement_3 end_statement?
	// statement_3
	if !_fail(parser, _statement_3Fail, errPos, failure, &pos) {
		goto fail
	}
	// end_statement?
	{
		pos2 := pos
		// end_statement
		if !_fail(parser, _end_statementFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _statement_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_statement_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// statement_3 end_statement?
	{
		var node0 string
		// statement_3
		if p, n := _statement_3Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// end_statement?
		{
			pos2 := pos
			// end_statement
			if p, n := _end_statementAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _statement_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _statement_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// prenex? spaces? predication
	// prenex?
	{
		pos2 := pos
		// prenex
		if !_accept(parser, _prenexAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// predication
	if !_accept(parser, _predicationAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _statement_3, start, pos, perr)
fail:
	return _memoize(parser, _statement_3, start, -1, perr)
}

func _statement_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_statement_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "statement_3"}
	// prenex? spaces? predication
	// prenex?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// prenex
		if !_node(parser, _prenexNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// spaces?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// predication
	if !_node(parser, _predicationNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _statement_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _statement_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement_3}
	// prenex? spaces? predication
	// prenex?
	{
		pos2 := pos
		// prenex
		if !_fail(parser, _prenexFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// predication
	if !_fail(parser, _predicationFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _statement_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_statement_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// prenex? spaces? predication
	{
		var node0 string
		// prenex?
		{
			pos2 := pos
			// prenex
			if p, n := _prenexAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos6 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// predication
		if p, n := _predicationAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prenexAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _prenex, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// terms spaces? end_prenex
	// terms
	if !_accept(parser, _termsAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_prenex
	if !_accept(parser, _end_prenexAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _prenex, start, pos, perr)
fail:
	return _memoize(parser, _prenex, start, -1, perr)
}

func _prenexNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_prenex]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prenex}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "prenex"}
	// terms spaces? end_prenex
	// terms
	if !_node(parser, _termsNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_prenex
	if !_node(parser, _end_prenexNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prenexFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _prenex, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prenex",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prenex}
	// terms spaces? end_prenex
	// terms
	if !_fail(parser, _termsFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_prenex
	if !_fail(parser, _end_prenexFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prenexAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_prenex]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prenex}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// terms spaces? end_prenex
	{
		var node0 string
		// terms
		if p, n := _termsAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_prenex
		if p, n := _end_prenexAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicationAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _predication, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// predicate spaces? terms?
	// predicate
	if !_accept(parser, _predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms?
	{
		pos6 := pos
		// terms
		if !_accept(parser, _termsAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _predication, start, pos, perr)
fail:
	return _memoize(parser, _predication, start, -1, perr)
}

func _predicationNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predication}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "predication"}
	// predicate spaces? terms?
	// predicate
	if !_node(parser, _predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// terms
		if !_node(parser, _termsNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicationFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _predication, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predication",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predication}
	// predicate spaces? terms?
	// predicate
	if !_fail(parser, _predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms?
	{
		pos6 := pos
		// terms
		if !_fail(parser, _termsFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicationAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predication}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// predicate spaces? terms?
	{
		var node0 string
		// predicate
		if p, n := _predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms?
		{
			pos6 := pos
			// terms
			if p, n := _termsAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// statement_1 spaces? co_bar_statement
	// statement_1
	if !_accept(parser, _statement_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_accept(parser, _co_bar_statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_statement, start, pos, perr)
fail:
	return _memoize(parser, _coP_statement, start, -1, perr)
}

func _coP_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_statement"}
	// statement_1 spaces? co_bar_statement
	// statement_1
	if !_node(parser, _statement_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_node(parser, _co_bar_statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_statement}
	// statement_1 spaces? co_bar_statement
	// statement_1
	if !_fail(parser, _statement_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_fail(parser, _co_bar_statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// statement_1 spaces? co_bar_statement
	{
		var node0 string
		// statement_1
		if p, n := _statement_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_statement
		if p, n := _co_bar_statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? statement_1
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement_1
	if !_accept(parser, _statement_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_statement, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_statement, start, -1, perr)
}

func _co_bar_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_statement"}
	// connective spaces? statement_1
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement_1
	if !_node(parser, _statement_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_statement}
	// connective spaces? statement_1
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement_1
	if !_fail(parser, _statement_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? statement_1
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement_1
		if p, n := _statement_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_statement_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_statement_1
	if !_accept(parser, _forethought_coP_statement_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_statement, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_statement, start, -1, perr)
}

func _forethought_coP_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_statement"}
	// forethought_connective spaces? forethought_coP_statement_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_statement_1
	if !_node(parser, _forethought_coP_statement_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_statement}
	// forethought_connective spaces? forethought_coP_statement_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_statement_1
	if !_fail(parser, _forethought_coP_statement_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_statement_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_statement_1
		if p, n := _forethought_coP_statement_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_statement_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_statement_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// statement spaces? forethought_co_bar_statement
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_statement
	if !_accept(parser, _forethought_co_bar_statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_statement_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_statement_1, start, -1, perr)
}

func _forethought_coP_statement_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_statement_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_statement_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_statement_1"}
	// statement spaces? forethought_co_bar_statement
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_statement
	if !_node(parser, _forethought_co_bar_statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_statement_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_statement_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_statement_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_statement_1}
	// statement spaces? forethought_co_bar_statement
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_statement
	if !_fail(parser, _forethought_co_bar_statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_statement_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_statement_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_statement_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// statement spaces? forethought_co_bar_statement
	{
		var node0 string
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_statement
		if p, n := _forethought_co_bar_statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? statement
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_statement, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_statement, start, -1, perr)
}

func _forethought_co_bar_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_statement"}
	// forethought_marker spaces? statement
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_statement}
	// forethought_marker spaces? statement
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? statement
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial_predicate
	if !_accept(parser, _serial_predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _predicate, start, pos, perr)
fail:
	return _memoize(parser, _predicate, start, -1, perr)
}

func _predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "predicate"}
	// serial_predicate
	if !_node(parser, _serial_predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predicate}
	// serial_predicate
	if !_fail(parser, _serial_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// serial_predicate
	if p, n := _serial_predicateAction(parser, pos); n == nil {
		goto fail
	} else {
		node = *n
		pos = p
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_pred/predicate_2
	{
		pos3 := pos
		// coP_pred
		if !_accept(parser, _coP_predAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// predicate_2
		if !_accept(parser, _predicate_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _predicate_1, start, -1, perr)
}

func _predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "predicate_1"}
	// coP_pred/predicate_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_pred
		if !_node(parser, _coP_predNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// predicate_2
		if !_node(parser, _predicate_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predicate_1}
	// coP_pred/predicate_2
	{
		pos3 := pos
		// coP_pred
		if !_fail(parser, _coP_predFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// predicate_2
		if !_fail(parser, _predicate_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_pred/predicate_2
	{
		pos3 := pos
		var node2 string
		// coP_pred
		if p, n := _coP_predAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// predicate_2
		if p, n := _predicate_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicate_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _predicate_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_pred/LU_predicate/MI_predicate/PO_predicate/quotation_predicate/prefix spaces? predicate_2/predicate_3
	{
		pos3 := pos
		// forethought_coP_pred
		if !_accept(parser, _forethought_coP_predAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_predicate
		if !_accept(parser, _LU_predicateAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MI_predicate
		if !_accept(parser, _MI_predicateAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// PO_predicate
		if !_accept(parser, _PO_predicateAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// quotation_predicate
		if !_accept(parser, _quotation_predicateAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// prefix spaces? predicate_2
		// prefix
		if !_accept(parser, _prefixAccepts, &pos, &perr) {
			goto fail9
		}
		// spaces?
		{
			pos12 := pos
			// spaces
			if !_accept(parser, _spacesAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok14
		fail13:
			pos = pos12
		ok14:
		}
		// predicate_2
		if !_accept(parser, _predicate_2Accepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// predicate_3
		if !_accept(parser, _predicate_3Accepts, &pos, &perr) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _predicate_2, start, pos, perr)
fail:
	return _memoize(parser, _predicate_2, start, -1, perr)
}

func _predicate_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "predicate_2"}
	// forethought_coP_pred/LU_predicate/MI_predicate/PO_predicate/quotation_predicate/prefix spaces? predicate_2/predicate_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_pred
		if !_node(parser, _forethought_coP_predNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_predicate
		if !_node(parser, _LU_predicateNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_predicate
		if !_node(parser, _MI_predicateNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_predicate
		if !_node(parser, _PO_predicateNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// quotation_predicate
		if !_node(parser, _quotation_predicateNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// prefix spaces? predicate_2
		// prefix
		if !_node(parser, _prefixNode, node, &pos) {
			goto fail9
		}
		// spaces?
		{
			nkids11 := len(node.Kids)
			pos12 := pos
			// spaces
			if !_node(parser, _spacesNode, node, &pos) {
				goto fail13
			}
			goto ok14
		fail13:
			node.Kids = node.Kids[:nkids11]
			pos = pos12
		ok14:
		}
		// predicate_2
		if !_node(parser, _predicate_2Node, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// predicate_3
		if !_node(parser, _predicate_3Node, node, &pos) {
			goto fail15
		}
		goto ok0
	fail15:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicate_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _predicate_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predicate_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predicate_2}
	// forethought_coP_pred/LU_predicate/MI_predicate/PO_predicate/quotation_predicate/prefix spaces? predicate_2/predicate_3
	{
		pos3 := pos
		// forethought_coP_pred
		if !_fail(parser, _forethought_coP_predFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_predicate
		if !_fail(parser, _LU_predicateFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MI_predicate
		if !_fail(parser, _MI_predicateFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// PO_predicate
		if !_fail(parser, _PO_predicateFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// quotation_predicate
		if !_fail(parser, _quotation_predicateFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// prefix spaces? predicate_2
		// prefix
		if !_fail(parser, _prefixFail, errPos, failure, &pos) {
			goto fail9
		}
		// spaces?
		{
			pos12 := pos
			// spaces
			if !_fail(parser, _spacesFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok14
		fail13:
			pos = pos12
		ok14:
		}
		// predicate_2
		if !_fail(parser, _predicate_2Fail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// predicate_3
		if !_fail(parser, _predicate_3Fail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicate_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_pred/LU_predicate/MI_predicate/PO_predicate/quotation_predicate/prefix spaces? predicate_2/predicate_3
	{
		pos3 := pos
		var node2 string
		// forethought_coP_pred
		if p, n := _forethought_coP_predAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// LU_predicate
		if p, n := _LU_predicateAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// MI_predicate
		if p, n := _MI_predicateAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// PO_predicate
		if p, n := _PO_predicateAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// quotation_predicate
		if p, n := _quotation_predicateAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// prefix spaces? predicate_2
		{
			var node10 string
			// prefix
			if p, n := _prefixAction(parser, pos); n == nil {
				goto fail9
			} else {
				node10 = *n
				pos = p
			}
			node, node10 = node+node10, ""
			// spaces?
			{
				pos12 := pos
				// spaces
				if p, n := _spacesAction(parser, pos); n == nil {
					goto fail13
				} else {
					node10 = *n
					pos = p
				}
				goto ok14
			fail13:
				node10 = ""
				pos = pos12
			ok14:
			}
			node, node10 = node+node10, ""
			// predicate_2
			if p, n := _predicate_2Action(parser, pos); n == nil {
				goto fail9
			} else {
				node10 = *n
				pos = p
			}
			node, node10 = node+node10, ""
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// predicate_3
		if p, n := _predicate_3Action(parser, pos); n == nil {
			goto fail15
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail15:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicate_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _predicate_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// predicate_4 spaces? freemod?
	// predicate_4
	if !_accept(parser, _predicate_4Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// freemod?
	{
		pos6 := pos
		// freemod
		if !_accept(parser, _freemodAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _predicate_3, start, pos, perr)
fail:
	return _memoize(parser, _predicate_3, start, -1, perr)
}

func _predicate_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_predicate_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "predicate_3"}
	// predicate_4 spaces? freemod?
	// predicate_4
	if !_node(parser, _predicate_4Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// freemod?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// freemod
		if !_node(parser, _freemodNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicate_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _predicate_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predicate_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predicate_3}
	// predicate_4 spaces? freemod?
	// predicate_4
	if !_fail(parser, _predicate_4Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// freemod?
	{
		pos6 := pos
		// freemod
		if !_fail(parser, _freemodFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicate_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_predicate_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// predicate_4 spaces? freemod?
	{
		var node0 string
		// predicate_4
		if p, n := _predicate_4Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// freemod?
		{
			pos6 := pos
			// freemod
			if p, n := _freemodAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicate_4Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _predicate_4, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// verb_syllable compound_syllable+/!function_word verb_syllable
	{
		pos3 := pos
		// verb_syllable compound_syllable+
		// verb_syllable
		if !_accept(parser, _verb_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word verb_syllable
		// !function_word
		{
			pos13 := pos
			perr15 := perr
			// function_word
			if !_accept(parser, _function_wordAccepts, &pos, &perr) {
				goto ok12
			}
			pos = pos13
			perr = _max(perr15, pos)
			goto fail10
		ok12:
			pos = pos13
			perr = perr15
		}
		// verb_syllable
		if !_accept(parser, _verb_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _predicate_4, start, pos, perr)
fail:
	return _memoize(parser, _predicate_4, start, -1, perr)
}

func _predicate_4Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_predicate_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_4}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "predicate_4"}
	// verb_syllable compound_syllable+/!function_word verb_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// verb_syllable compound_syllable+
		// verb_syllable
		if !_node(parser, _verb_syllableNode, node, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_node(parser, _compound_syllableNode, node, &pos) {
			goto fail4
		}
		for {
			nkids6 := len(node.Kids)
			pos7 := pos
			// compound_syllable
			if !_node(parser, _compound_syllableNode, node, &pos) {
				goto fail9
			}
			continue
		fail9:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
			break
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// !function_word verb_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(node.Kids)
			// function_word
			if !_node(parser, _function_wordNode, node, &pos) {
				goto ok12
			}
			pos = pos13
			node.Kids = node.Kids[:nkids14]
			goto fail10
		ok12:
			pos = pos13
			node.Kids = node.Kids[:nkids14]
		}
		// verb_syllable
		if !_node(parser, _verb_syllableNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicate_4Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _predicate_4, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predicate_4",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predicate_4}
	// verb_syllable compound_syllable+/!function_word verb_syllable
	{
		pos3 := pos
		// verb_syllable compound_syllable+
		// verb_syllable
		if !_fail(parser, _verb_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word verb_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(failure.Kids)
			// function_word
			if !_fail(parser, _function_wordFail, errPos, failure, &pos) {
				goto ok12
			}
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "!function_word",
				})
			}
			goto fail10
		ok12:
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
		}
		// verb_syllable
		if !_fail(parser, _verb_syllableFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicate_4Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_predicate_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_4}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// verb_syllable compound_syllable+/!function_word verb_syllable
	{
		pos3 := pos
		var node2 string
		// verb_syllable compound_syllable+
		{
			var node5 string
			// verb_syllable
			if p, n := _verb_syllableAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
			// compound_syllable+
			{
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail4
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
			}
			for {
				pos7 := pos
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail9
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
				continue
			fail9:
				pos = pos7
				break
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// !function_word verb_syllable
		{
			var node11 string
			// !function_word
			{
				pos13 := pos
				// function_word
				if p, n := _function_wordAction(parser, pos); n == nil {
					goto ok12
				} else {
					pos = p
				}
				pos = pos13
				goto fail10
			ok12:
				pos = pos13
				node11 = ""
			}
			node, node11 = node+node11, ""
			// verb_syllable
			if p, n := _verb_syllableAction(parser, pos); n == nil {
				goto fail10
			} else {
				node11 = *n
				pos = p
			}
			node, node11 = node+node11, ""
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _serial_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _serial_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial_predicate_2/predicate_1
	{
		pos3 := pos
		// serial_predicate_2
		if !_accept(parser, _serial_predicate_2Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// predicate_1
		if !_accept(parser, _predicate_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _serial_predicate, start, pos, perr)
fail:
	return _memoize(parser, _serial_predicate, start, -1, perr)
}

func _serial_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_serial_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "serial_predicate"}
	// serial_predicate_2/predicate_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial_predicate_2
		if !_node(parser, _serial_predicate_2Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// predicate_1
		if !_node(parser, _predicate_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _serial_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _serial_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "serial_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _serial_predicate}
	// serial_predicate_2/predicate_1
	{
		pos3 := pos
		// serial_predicate_2
		if !_fail(parser, _serial_predicate_2Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// predicate_1
		if !_fail(parser, _predicate_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _serial_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_serial_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// serial_predicate_2/predicate_1
	{
		pos3 := pos
		var node2 string
		// serial_predicate_2
		if p, n := _serial_predicate_2Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// predicate_1
		if p, n := _predicate_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _serial_predicate_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _serial_predicate_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// predicate_1 spaces? serial_predicate
	// predicate_1
	if !_accept(parser, _predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_accept(parser, _serial_predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _serial_predicate_2, start, pos, perr)
fail:
	return _memoize(parser, _serial_predicate_2, start, -1, perr)
}

func _serial_predicate_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_serial_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_predicate_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "serial_predicate_2"}
	// predicate_1 spaces? serial_predicate
	// predicate_1
	if !_node(parser, _predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_node(parser, _serial_predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _serial_predicate_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _serial_predicate_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "serial_predicate_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _serial_predicate_2}
	// predicate_1 spaces? serial_predicate
	// predicate_1
	if !_fail(parser, _predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_fail(parser, _serial_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _serial_predicate_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_serial_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_predicate_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// predicate_1 spaces? serial_predicate
	{
		var node0 string
		// predicate_1
		if p, n := _predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// serial_predicate
		if p, n := _serial_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_predAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_pred, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// predicate_2 spaces? co_bar_pred
	// predicate_2
	if !_accept(parser, _predicate_2Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_accept(parser, _co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_pred, start, pos, perr)
fail:
	return _memoize(parser, _coP_pred, start, -1, perr)
}

func _coP_predNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_pred"}
	// predicate_2 spaces? co_bar_pred
	// predicate_2
	if !_node(parser, _predicate_2Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_node(parser, _co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_predFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_pred, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_pred",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_pred}
	// predicate_2 spaces? co_bar_pred
	// predicate_2
	if !_fail(parser, _predicate_2Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_fail(parser, _co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_predAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// predicate_2 spaces? co_bar_pred
	{
		var node0 string
		// predicate_2
		if p, n := _predicate_2Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_pred
		if p, n := _co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_predAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_pred, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? predicate_1
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// predicate_1
	if !_accept(parser, _predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_pred, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_pred, start, -1, perr)
}

func _co_bar_predNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_pred}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_pred"}
	// connective spaces? predicate_1
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// predicate_1
	if !_node(parser, _predicate_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_predFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_pred, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_pred",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_pred}
	// connective spaces? predicate_1
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// predicate_1
	if !_fail(parser, _predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_predAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_pred}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? predicate_1
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// predicate_1
		if p, n := _predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_predAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_pred_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_1
	if !_accept(parser, _forethought_coP_pred_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred, start, -1, perr)
}

func _forethought_coP_predNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred"}
	// forethought_connective spaces? forethought_coP_pred_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_1
	if !_node(parser, _forethought_coP_pred_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_predFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred}
	// forethought_connective spaces? forethought_coP_pred_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_1
	if !_fail(parser, _forethought_coP_pred_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_predAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_pred_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_pred_1
		if p, n := _forethought_coP_pred_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// predicate spaces? forethought_co_bar_pred
	// predicate
	if !_accept(parser, _predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_accept(parser, _forethought_co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_1, start, -1, perr)
}

func _forethought_coP_pred_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_1"}
	// predicate spaces? forethought_co_bar_pred
	// predicate
	if !_node(parser, _predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_node(parser, _forethought_co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_1}
	// predicate spaces? forethought_co_bar_pred
	// predicate
	if !_fail(parser, _predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_fail(parser, _forethought_co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// predicate spaces? forethought_co_bar_pred
	{
		var node0 string
		// predicate
		if p, n := _predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_pred
		if p, n := _forethought_co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_predAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_pred, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? predicate
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// predicate
	if !_accept(parser, _predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_pred, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_pred, start, -1, perr)
}

func _forethought_co_bar_predNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_pred}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_pred"}
	// forethought_marker spaces? predicate
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// predicate
	if !_node(parser, _predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_predFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_pred, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_pred",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_pred}
	// forethought_marker spaces? predicate
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// predicate
	if !_fail(parser, _predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_predAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_pred]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_pred}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? predicate
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// predicate
		if p, n := _predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// LU_predicate_tone spaces? statement
	// LU_predicate_tone
	if !_accept(parser, _LU_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_predicate, start, pos, perr)
fail:
	return _memoize(parser, _LU_predicate, start, -1, perr)
}

func _LU_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_predicate"}
	// LU_predicate_tone spaces? statement
	// LU_predicate_tone
	if !_node(parser, _LU_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_predicate}
	// LU_predicate_tone spaces? statement
	// LU_predicate_tone
	if !_fail(parser, _LU_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// LU_predicate_tone spaces? statement
	{
		var node0 string
		// LU_predicate_tone
		if p, n := _LU_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &LU verb_syllable
	// &LU
	{
		pos2 := pos
		perr4 := perr
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// verb_syllable
	if !_accept(parser, _verb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _LU_predicate_tone, start, -1, perr)
}

func _LU_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_predicate_tone"}
	// &LU verb_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// LU
		if !_node(parser, _LUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// verb_syllable
	if !_node(parser, _verb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_predicate_tone}
	// &LU verb_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&LU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// verb_syllable
	if !_fail(parser, _verb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &LU verb_syllable
	{
		var node0 string
		// &LU
		{
			pos2 := pos
			// LU
			if p, n := _LUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// verb_syllable
		if p, n := _verb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_predicate_1 spaces? end_predicatizer?
	// MI_predicate_1
	if !_accept(parser, _MI_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _MI_predicate, start, pos, perr)
fail:
	return _memoize(parser, _MI_predicate, start, -1, perr)
}

func _MI_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_predicate"}
	// MI_predicate_1 spaces? end_predicatizer?
	// MI_predicate_1
	if !_node(parser, _MI_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_predicate}
	// MI_predicate_1 spaces? end_predicatizer?
	// MI_predicate_1
	if !_fail(parser, _MI_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_predicate_1 spaces? end_predicatizer?
	{
		var node0 string
		// MI_predicate_1
		if p, n := _MI_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_predicate_tone
	if !_accept(parser, _MI_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_accept(parser, _predicateAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_accept(parser, _adverbAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	return _memoize(parser, _MI_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _MI_predicate_1, start, -1, perr)
}

func _MI_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_predicate_1"}
	// MI_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_predicate_tone
	if !_node(parser, _MI_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// predicate/argument/adverb/prepositional_phrase
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// predicate
			if !_node(parser, _predicateNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail14
			}
			goto ok7
		fail14:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			goto fail
		ok7:
		}
		sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
		node.Kids = append(node.Kids[:nkids5], sub)
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_predicate_1}
	// MI_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_predicate_tone
	if !_fail(parser, _MI_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_fail(parser, _predicateFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_fail(parser, _adverbFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	{
		var node0 string
		// MI_predicate_tone
		if p, n := _MI_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (predicate/argument/adverb/prepositional_phrase)
		// predicate/argument/adverb/prepositional_phrase
		{
			pos8 := pos
			var node7 string
			// predicate
			if p, n := _predicateAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// adverb
			if p, n := _adverbAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// prepositional_phrase
			if p, n := _prepositional_phraseAction(parser, pos); n == nil {
				goto fail12
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail12:
			node0 = node7
			pos = pos8
			goto fail
		ok5:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MI verb_syllable
	// &MI
	{
		pos2 := pos
		perr4 := perr
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// verb_syllable
	if !_accept(parser, _verb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MI_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _MI_predicate_tone, start, -1, perr)
}

func _MI_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_predicate_tone"}
	// &MI verb_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MI
		if !_node(parser, _MINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// verb_syllable
	if !_node(parser, _verb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_predicate_tone}
	// &MI verb_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// verb_syllable
	if !_fail(parser, _verb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MI verb_syllable
	{
		var node0 string
		// &MI
		{
			pos2 := pos
			// MI
			if p, n := _MIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// verb_syllable
		if p, n := _verb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_predicate_1 spaces? end_predicatizer?
	// PO_predicate_1
	if !_accept(parser, _PO_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _PO_predicate, start, pos, perr)
fail:
	return _memoize(parser, _PO_predicate, start, -1, perr)
}

func _PO_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_predicate"}
	// PO_predicate_1 spaces? end_predicatizer?
	// PO_predicate_1
	if !_node(parser, _PO_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_predicate}
	// PO_predicate_1 spaces? end_predicatizer?
	// PO_predicate_1
	if !_fail(parser, _PO_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_predicate_1 spaces? end_predicatizer?
	{
		var node0 string
		// PO_predicate_1
		if p, n := _PO_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_predicate_tone spaces? argument
	// PO_predicate_tone
	if !_accept(parser, _PO_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _PO_predicate_1, start, -1, perr)
}

func _PO_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_predicate_1"}
	// PO_predicate_tone spaces? argument
	// PO_predicate_tone
	if !_node(parser, _PO_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_predicate_1}
	// PO_predicate_tone spaces? argument
	// PO_predicate_tone
	if !_fail(parser, _PO_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_predicate_tone spaces? argument
	{
		var node0 string
		// PO_predicate_tone
		if p, n := _PO_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &PO verb_syllable
	// &PO
	{
		pos2 := pos
		perr4 := perr
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// verb_syllable
	if !_accept(parser, _verb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _PO_predicate_tone, start, -1, perr)
}

func _PO_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_predicate_tone"}
	// &PO verb_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// verb_syllable
	if !_node(parser, _verb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_predicate_tone}
	// &PO verb_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&PO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// verb_syllable
	if !_fail(parser, _verb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &PO verb_syllable
	{
		var node0 string
		// &PO
		{
			pos2 := pos
			// PO
			if p, n := _POAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// verb_syllable
		if p, n := _verb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quotation_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _quotation_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_predicate spaces? end_quote
	// MO_predicate
	if !_accept(parser, _MO_predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_accept(parser, _end_quoteAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _quotation_predicate, start, pos, perr)
fail:
	return _memoize(parser, _quotation_predicate, start, -1, perr)
}

func _quotation_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_quotation_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quotation_predicate"}
	// MO_predicate spaces? end_quote
	// MO_predicate
	if !_node(parser, _MO_predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_quote
	if !_node(parser, _end_quoteNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quotation_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _quotation_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quotation_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quotation_predicate}
	// MO_predicate spaces? end_quote
	// MO_predicate
	if !_fail(parser, _MO_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_fail(parser, _end_quoteFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quotation_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_quotation_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_predicate spaces? end_quote
	{
		var node0 string
		// MO_predicate
		if p, n := _MO_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_quote
		if p, n := _end_quoteAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_predicate_tone spaces? discourse
	// MO_predicate_tone
	if !_accept(parser, _MO_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_accept(parser, _discourseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_predicate, start, pos, perr)
fail:
	return _memoize(parser, _MO_predicate, start, -1, perr)
}

func _MO_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_predicate"}
	// MO_predicate_tone spaces? discourse
	// MO_predicate_tone
	if !_node(parser, _MO_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// discourse
	if !_node(parser, _discourseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_predicate}
	// MO_predicate_tone spaces? discourse
	// MO_predicate_tone
	if !_fail(parser, _MO_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_fail(parser, _discourseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_predicate_tone spaces? discourse
	{
		var node0 string
		// MO_predicate_tone
		if p, n := _MO_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// discourse
		if p, n := _discourseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MO verb_syllable
	// &MO
	{
		pos2 := pos
		perr4 := perr
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// verb_syllable
	if !_accept(parser, _verb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _MO_predicate_tone, start, -1, perr)
}

func _MO_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_predicate_tone"}
	// &MO verb_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MO
		if !_node(parser, _MONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// verb_syllable
	if !_node(parser, _verb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_predicate_tone}
	// &MO verb_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// verb_syllable
	if !_fail(parser, _verb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MO verb_syllable
	{
		var node0 string
		// &MO
		{
			pos2 := pos
			// MO
			if p, n := _MOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// verb_syllable
		if p, n := _verb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termsAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _terms, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// terms_2/term
	{
		pos3 := pos
		// terms_2
		if !_accept(parser, _terms_2Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// term
		if !_accept(parser, _termAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _terms, start, pos, perr)
fail:
	return _memoize(parser, _terms, start, -1, perr)
}

func _termsNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_terms]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "terms"}
	// terms_2/term
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// terms_2
		if !_node(parser, _terms_2Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// term
		if !_node(parser, _termNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _termsFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _terms, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms}
	// terms_2/term
	{
		pos3 := pos
		// terms_2
		if !_fail(parser, _terms_2Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// term
		if !_fail(parser, _termFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _termsAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_terms]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// terms_2/term
	{
		pos3 := pos
		var node2 string
		// terms_2
		if p, n := _terms_2Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// term
		if p, n := _termAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _terms_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// term spaces? terms
	// term
	if !_accept(parser, _termAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms
	if !_accept(parser, _termsAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _terms_2, start, pos, perr)
fail:
	return _memoize(parser, _terms_2, start, -1, perr)
}

func _terms_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_terms_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "terms_2"}
	// term spaces? terms
	// term
	if !_node(parser, _termNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms
	if !_node(parser, _termsNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _terms_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_2}
	// term spaces? terms
	// term
	if !_fail(parser, _termFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms
	if !_fail(parser, _termsFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_terms_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// term spaces? terms
	{
		var node0 string
		// term
		if p, n := _termAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms
		if p, n := _termsAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _term, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (linking_word? spaces? argument)/adverb/termset/prepositional_phrase
	{
		pos3 := pos
		// (linking_word? spaces? argument)
		// linking_word? spaces? argument
		// linking_word?
		{
			pos7 := pos
			// linking_word
			if !_accept(parser, _linking_wordAccepts, &pos, &perr) {
				goto fail8
			}
			goto ok9
		fail8:
			pos = pos7
		ok9:
		}
		// spaces?
		{
			pos11 := pos
			// spaces
			if !_accept(parser, _spacesAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok13
		fail12:
			pos = pos11
		ok13:
		}
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb
		if !_accept(parser, _adverbAccepts, &pos, &perr) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		// termset
		if !_accept(parser, _termsetAccepts, &pos, &perr) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		// prepositional_phrase
		if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _term, start, pos, perr)
fail:
	return _memoize(parser, _term, start, -1, perr)
}

func _termNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_term]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _term}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "term"}
	// (linking_word? spaces? argument)/adverb/termset/prepositional_phrase
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// (linking_word? spaces? argument)
		{
			nkids5 := len(node.Kids)
			pos06 := pos
			// linking_word? spaces? argument
			// linking_word?
			{
				nkids8 := len(node.Kids)
				pos9 := pos
				// linking_word
				if !_node(parser, _linking_wordNode, node, &pos) {
					goto fail10
				}
				goto ok11
			fail10:
				node.Kids = node.Kids[:nkids8]
				pos = pos9
			ok11:
			}
			// spaces?
			{
				nkids12 := len(node.Kids)
				pos13 := pos
				// spaces
				if !_node(parser, _spacesNode, node, &pos) {
					goto fail14
				}
				goto ok15
			fail14:
				node.Kids = node.Kids[:nkids12]
				pos = pos13
			ok15:
			}
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail4
			}
			sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
			node.Kids = append(node.Kids[:nkids5], sub)
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb
		if !_node(parser, _adverbNode, node, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// termset
		if !_node(parser, _termsetNode, node, &pos) {
			goto fail17
		}
		goto ok0
	fail17:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// prepositional_phrase
		if !_node(parser, _prepositional_phraseNode, node, &pos) {
			goto fail18
		}
		goto ok0
	fail18:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _termFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _term, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "term",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _term}
	// (linking_word? spaces? argument)/adverb/termset/prepositional_phrase
	{
		pos3 := pos
		// (linking_word? spaces? argument)
		// linking_word? spaces? argument
		// linking_word?
		{
			pos7 := pos
			// linking_word
			if !_fail(parser, _linking_wordFail, errPos, failure, &pos) {
				goto fail8
			}
			goto ok9
		fail8:
			pos = pos7
		ok9:
		}
		// spaces?
		{
			pos11 := pos
			// spaces
			if !_fail(parser, _spacesFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok13
		fail12:
			pos = pos11
		ok13:
		}
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb
		if !_fail(parser, _adverbFail, errPos, failure, &pos) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		// termset
		if !_fail(parser, _termsetFail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		// prepositional_phrase
		if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _termAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_term]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _term}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (linking_word? spaces? argument)/adverb/termset/prepositional_phrase
	{
		pos3 := pos
		var node2 string
		// (linking_word? spaces? argument)
		// linking_word? spaces? argument
		{
			var node5 string
			// linking_word?
			{
				pos7 := pos
				// linking_word
				if p, n := _linking_wordAction(parser, pos); n == nil {
					goto fail8
				} else {
					node5 = *n
					pos = p
				}
				goto ok9
			fail8:
				node5 = ""
				pos = pos7
			ok9:
			}
			node, node5 = node+node5, ""
			// spaces?
			{
				pos11 := pos
				// spaces
				if p, n := _spacesAction(parser, pos); n == nil {
					goto fail12
				} else {
					node5 = *n
					pos = p
				}
				goto ok13
			fail12:
				node5 = ""
				pos = pos11
			ok13:
			}
			node, node5 = node+node5, ""
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// adverb
		if p, n := _adverbAction(parser, pos); n == nil {
			goto fail14
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail14:
		node = node2
		pos = pos3
		// termset
		if p, n := _termsetAction(parser, pos); n == nil {
			goto fail15
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail15:
		node = node2
		pos = pos3
		// prepositional_phrase
		if p, n := _prepositional_phraseAction(parser, pos); n == nil {
			goto fail16
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail16:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _argumentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _argument, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_arg/arg_1
	{
		pos3 := pos
		// coP_arg
		if !_accept(parser, _coP_argAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_1
		if !_accept(parser, _arg_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _argument, start, pos, perr)
fail:
	return _memoize(parser, _argument, start, -1, perr)
}

func _argumentNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _argument}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "argument"}
	// coP_arg/arg_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_arg
		if !_node(parser, _coP_argNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_1
		if !_node(parser, _arg_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _argumentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _argument, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "argument",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _argument}
	// coP_arg/arg_1
	{
		pos3 := pos
		// coP_arg
		if !_fail(parser, _coP_argFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_1
		if !_fail(parser, _arg_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _argumentAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _argument}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_arg/arg_1
	{
		pos3 := pos
		var node2 string
		// coP_arg
		if p, n := _coP_argAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// arg_1
		if p, n := _arg_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_arg/arg_2
	{
		pos3 := pos
		// forethought_coP_arg
		if !_accept(parser, _forethought_coP_argAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_2
		if !_accept(parser, _arg_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _arg_1, start, pos, perr)
fail:
	return _memoize(parser, _arg_1, start, -1, perr)
}

func _arg_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_1"}
	// forethought_coP_arg/arg_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_arg
		if !_node(parser, _forethought_coP_argNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_2
		if !_node(parser, _arg_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_1}
	// forethought_coP_arg/arg_2
	{
		pos3 := pos
		// forethought_coP_arg
		if !_fail(parser, _forethought_coP_argFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_2
		if !_fail(parser, _arg_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_arg/arg_2
	{
		pos3 := pos
		var node2 string
		// forethought_coP_arg
		if p, n := _forethought_coP_argAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// arg_2
		if p, n := _arg_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// focus? spaces? arg_3
	// focus?
	{
		pos2 := pos
		// focus
		if !_accept(parser, _focusAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// arg_3
	if !_accept(parser, _arg_3Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _arg_2, start, pos, perr)
fail:
	return _memoize(parser, _arg_2, start, -1, perr)
}

func _arg_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_2"}
	// focus? spaces? arg_3
	// focus?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// focus
		if !_node(parser, _focusNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// spaces?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// arg_3
	if !_node(parser, _arg_3Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_2}
	// focus? spaces? arg_3
	// focus?
	{
		pos2 := pos
		// focus
		if !_fail(parser, _focusFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// arg_3
	if !_fail(parser, _arg_3Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// focus? spaces? arg_3
	{
		var node0 string
		// focus?
		{
			pos2 := pos
			// focus
			if p, n := _focusAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos6 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// arg_3
		if p, n := _arg_3Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// quantifier? spaces? arg_4
	// quantifier?
	{
		pos2 := pos
		// quantifier
		if !_accept(parser, _quantifierAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// arg_4
	if !_accept(parser, _arg_4Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _arg_3, start, pos, perr)
fail:
	return _memoize(parser, _arg_3, start, -1, perr)
}

func _arg_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_3"}
	// quantifier? spaces? arg_4
	// quantifier?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// quantifier
		if !_node(parser, _quantifierNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// spaces?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// arg_4
	if !_node(parser, _arg_4Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_3}
	// quantifier? spaces? arg_4
	// quantifier?
	{
		pos2 := pos
		// quantifier
		if !_fail(parser, _quantifierFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// arg_4
	if !_fail(parser, _arg_4Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// quantifier? spaces? arg_4
	{
		var node0 string
		// quantifier?
		{
			pos2 := pos
			// quantifier
			if p, n := _quantifierAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos6 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// arg_4
		if p, n := _arg_4Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_4Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_4, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// arg_5 spaces? (relative_clause/freemod)?
	// arg_5
	if !_accept(parser, _arg_5Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (relative_clause/freemod)?
	{
		pos6 := pos
		// (relative_clause/freemod)
		// relative_clause/freemod
		{
			pos11 := pos
			// relative_clause
			if !_accept(parser, _relative_clauseAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// freemod
			if !_accept(parser, _freemodAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail7
		ok8:
		}
		goto ok14
	fail7:
		pos = pos6
	ok14:
	}
	return _memoize(parser, _arg_4, start, pos, perr)
fail:
	return _memoize(parser, _arg_4, start, -1, perr)
}

func _arg_4Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_4}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_4"}
	// arg_5 spaces? (relative_clause/freemod)?
	// arg_5
	if !_node(parser, _arg_5Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (relative_clause/freemod)?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// (relative_clause/freemod)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// relative_clause/freemod
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// relative_clause
				if !_node(parser, _relative_clauseNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// freemod
				if !_node(parser, _freemodNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail7
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
		}
		goto ok16
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok16:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_4Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_4, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_4",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_4}
	// arg_5 spaces? (relative_clause/freemod)?
	// arg_5
	if !_fail(parser, _arg_5Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (relative_clause/freemod)?
	{
		pos6 := pos
		// (relative_clause/freemod)
		// relative_clause/freemod
		{
			pos11 := pos
			// relative_clause
			if !_fail(parser, _relative_clauseFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// freemod
			if !_fail(parser, _freemodFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail7
		ok8:
		}
		goto ok14
	fail7:
		pos = pos6
	ok14:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_4Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_4}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// arg_5 spaces? (relative_clause/freemod)?
	{
		var node0 string
		// arg_5
		if p, n := _arg_5Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (relative_clause/freemod)?
		{
			pos6 := pos
			// (relative_clause/freemod)
			// relative_clause/freemod
			{
				pos11 := pos
				var node10 string
				// relative_clause
				if p, n := _relative_clauseAction(parser, pos); n == nil {
					goto fail12
				} else {
					node0 = *n
					pos = p
				}
				goto ok8
			fail12:
				node0 = node10
				pos = pos11
				// freemod
				if p, n := _freemodAction(parser, pos); n == nil {
					goto fail13
				} else {
					node0 = *n
					pos = p
				}
				goto ok8
			fail13:
				node0 = node10
				pos = pos11
				goto fail7
			ok8:
			}
			goto ok14
		fail7:
			node0 = ""
			pos = pos6
		ok14:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_5Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_5, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial_argument/arg_6
	{
		pos3 := pos
		// serial_argument
		if !_accept(parser, _serial_argumentAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_6
		if !_accept(parser, _arg_6Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _arg_5, start, pos, perr)
fail:
	return _memoize(parser, _arg_5, start, -1, perr)
}

func _arg_5Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_5]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_5}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_5"}
	// serial_argument/arg_6
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial_argument
		if !_node(parser, _serial_argumentNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_6
		if !_node(parser, _arg_6Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_5Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_5, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_5",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_5}
	// serial_argument/arg_6
	{
		pos3 := pos
		// serial_argument
		if !_fail(parser, _serial_argumentFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_6
		if !_fail(parser, _arg_6Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_5Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_5]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_5}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// serial_argument/arg_6
	{
		pos3 := pos
		var node2 string
		// serial_argument
		if p, n := _serial_argumentAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// arg_6
		if p, n := _arg_6Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_6Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_6, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_clause/coP_pred_arg/forethought_coP_pred_arg/LU_arg/MI_arg/PO_arg/quotation_argument/arg_7
	{
		pos3 := pos
		// content_clause
		if !_accept(parser, _content_clauseAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// coP_pred_arg
		if !_accept(parser, _coP_pred_argAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// forethought_coP_pred_arg
		if !_accept(parser, _forethought_coP_pred_argAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// LU_arg
		if !_accept(parser, _LU_argAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// MI_arg
		if !_accept(parser, _MI_argAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// PO_arg
		if !_accept(parser, _PO_argAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// quotation_argument
		if !_accept(parser, _quotation_argumentAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// arg_7
		if !_accept(parser, _arg_7Accepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _arg_6, start, pos, perr)
fail:
	return _memoize(parser, _arg_6, start, -1, perr)
}

func _arg_6Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_6]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_6}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_6"}
	// content_clause/coP_pred_arg/forethought_coP_pred_arg/LU_arg/MI_arg/PO_arg/quotation_argument/arg_7
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// content_clause
		if !_node(parser, _content_clauseNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// coP_pred_arg
		if !_node(parser, _coP_pred_argNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_coP_pred_arg
		if !_node(parser, _forethought_coP_pred_argNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_arg
		if !_node(parser, _LU_argNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_arg
		if !_node(parser, _MI_argNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_arg
		if !_node(parser, _PO_argNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// quotation_argument
		if !_node(parser, _quotation_argumentNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_7
		if !_node(parser, _arg_7Node, node, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_6Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_6, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_6",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_6}
	// content_clause/coP_pred_arg/forethought_coP_pred_arg/LU_arg/MI_arg/PO_arg/quotation_argument/arg_7
	{
		pos3 := pos
		// content_clause
		if !_fail(parser, _content_clauseFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// coP_pred_arg
		if !_fail(parser, _coP_pred_argFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// forethought_coP_pred_arg
		if !_fail(parser, _forethought_coP_pred_argFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// LU_arg
		if !_fail(parser, _LU_argFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// MI_arg
		if !_fail(parser, _MI_argFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// PO_arg
		if !_fail(parser, _PO_argFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// quotation_argument
		if !_fail(parser, _quotation_argumentFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// arg_7
		if !_fail(parser, _arg_7Fail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_6Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_6]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_6}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_clause/coP_pred_arg/forethought_coP_pred_arg/LU_arg/MI_arg/PO_arg/quotation_argument/arg_7
	{
		pos3 := pos
		var node2 string
		// content_clause
		if p, n := _content_clauseAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// coP_pred_arg
		if p, n := _coP_pred_argAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// forethought_coP_pred_arg
		if p, n := _forethought_coP_pred_argAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// LU_arg
		if p, n := _LU_argAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// MI_arg
		if p, n := _MI_argAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// PO_arg
		if p, n := _PO_argAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// quotation_argument
		if p, n := _quotation_argumentAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// arg_7
		if p, n := _arg_7Action(parser, pos); n == nil {
			goto fail11
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail11:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_7Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_7, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// arg_syllable compound_syllable+/!function_word arg_syllable
	{
		pos3 := pos
		// arg_syllable compound_syllable+
		// arg_syllable
		if !_accept(parser, _arg_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word arg_syllable
		// !function_word
		{
			pos13 := pos
			perr15 := perr
			// function_word
			if !_accept(parser, _function_wordAccepts, &pos, &perr) {
				goto ok12
			}
			pos = pos13
			perr = _max(perr15, pos)
			goto fail10
		ok12:
			pos = pos13
			perr = perr15
		}
		// arg_syllable
		if !_accept(parser, _arg_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _arg_7, start, pos, perr)
fail:
	return _memoize(parser, _arg_7, start, -1, perr)
}

func _arg_7Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_7]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_7}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_7"}
	// arg_syllable compound_syllable+/!function_word arg_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// arg_syllable compound_syllable+
		// arg_syllable
		if !_node(parser, _arg_syllableNode, node, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_node(parser, _compound_syllableNode, node, &pos) {
			goto fail4
		}
		for {
			nkids6 := len(node.Kids)
			pos7 := pos
			// compound_syllable
			if !_node(parser, _compound_syllableNode, node, &pos) {
				goto fail9
			}
			continue
		fail9:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
			break
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// !function_word arg_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(node.Kids)
			// function_word
			if !_node(parser, _function_wordNode, node, &pos) {
				goto ok12
			}
			pos = pos13
			node.Kids = node.Kids[:nkids14]
			goto fail10
		ok12:
			pos = pos13
			node.Kids = node.Kids[:nkids14]
		}
		// arg_syllable
		if !_node(parser, _arg_syllableNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_7Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_7, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_7",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_7}
	// arg_syllable compound_syllable+/!function_word arg_syllable
	{
		pos3 := pos
		// arg_syllable compound_syllable+
		// arg_syllable
		if !_fail(parser, _arg_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word arg_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(failure.Kids)
			// function_word
			if !_fail(parser, _function_wordFail, errPos, failure, &pos) {
				goto ok12
			}
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "!function_word",
				})
			}
			goto fail10
		ok12:
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
		}
		// arg_syllable
		if !_fail(parser, _arg_syllableFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_7Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_7]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_7}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// arg_syllable compound_syllable+/!function_word arg_syllable
	{
		pos3 := pos
		var node2 string
		// arg_syllable compound_syllable+
		{
			var node5 string
			// arg_syllable
			if p, n := _arg_syllableAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
			// compound_syllable+
			{
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail4
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
			}
			for {
				pos7 := pos
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail9
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
				continue
			fail9:
				pos = pos7
				break
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// !function_word arg_syllable
		{
			var node11 string
			// !function_word
			{
				pos13 := pos
				// function_word
				if p, n := _function_wordAction(parser, pos); n == nil {
					goto ok12
				} else {
					pos = p
				}
				pos = pos13
				goto fail10
			ok12:
				pos = pos13
				node11 = ""
			}
			node, node11 = node+node11, ""
			// arg_syllable
			if p, n := _arg_syllableAction(parser, pos); n == nil {
				goto fail10
			} else {
				node11 = *n
				pos = p
			}
			node, node11 = node+node11, ""
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _serial_argumentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _serial_argument, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// arg_6 spaces? serial_predicate
	// arg_6
	if !_accept(parser, _arg_6Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_accept(parser, _serial_predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _serial_argument, start, pos, perr)
fail:
	return _memoize(parser, _serial_argument, start, -1, perr)
}

func _serial_argumentNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_serial_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_argument}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "serial_argument"}
	// arg_6 spaces? serial_predicate
	// arg_6
	if !_node(parser, _arg_6Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_node(parser, _serial_predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _serial_argumentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _serial_argument, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "serial_argument",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _serial_argument}
	// arg_6 spaces? serial_predicate
	// arg_6
	if !_fail(parser, _arg_6Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_fail(parser, _serial_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _serial_argumentAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_serial_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_argument}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// arg_6 spaces? serial_predicate
	{
		var node0 string
		// arg_6
		if p, n := _arg_6Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// serial_predicate
		if p, n := _serial_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// arg_1 spaces? co_bar_arg
	// arg_1
	if !_accept(parser, _arg_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_arg
	if !_accept(parser, _co_bar_argAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_arg, start, pos, perr)
fail:
	return _memoize(parser, _coP_arg, start, -1, perr)
}

func _coP_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_arg"}
	// arg_1 spaces? co_bar_arg
	// arg_1
	if !_node(parser, _arg_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_arg
	if !_node(parser, _co_bar_argNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_arg}
	// arg_1 spaces? co_bar_arg
	// arg_1
	if !_fail(parser, _arg_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_arg
	if !_fail(parser, _co_bar_argFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// arg_1 spaces? co_bar_arg
	{
		var node0 string
		// arg_1
		if p, n := _arg_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_arg
		if p, n := _co_bar_argAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? argument
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_arg, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_arg, start, -1, perr)
}

func _co_bar_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_arg"}
	// connective spaces? argument
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_arg}
	// connective spaces? argument
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? argument
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_arg_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_arg_1
	if !_accept(parser, _forethought_coP_arg_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_arg, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_arg, start, -1, perr)
}

func _forethought_coP_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_arg"}
	// forethought_connective spaces? forethought_coP_arg_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_arg_1
	if !_node(parser, _forethought_coP_arg_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_arg}
	// forethought_connective spaces? forethought_coP_arg_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_arg_1
	if !_fail(parser, _forethought_coP_arg_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_arg_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_arg_1
		if p, n := _forethought_coP_arg_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_arg_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_arg_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// argument spaces? forethought_co_bar_arg
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_arg
	if !_accept(parser, _forethought_co_bar_argAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_arg_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_arg_1, start, -1, perr)
}

func _forethought_coP_arg_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_arg_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_arg_1"}
	// argument spaces? forethought_co_bar_arg
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_arg
	if !_node(parser, _forethought_co_bar_argNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_arg_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_arg_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_arg_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_arg_1}
	// argument spaces? forethought_co_bar_arg
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_arg
	if !_fail(parser, _forethought_co_bar_argFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_arg_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_arg_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// argument spaces? forethought_co_bar_arg
	{
		var node0 string
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_arg
		if p, n := _forethought_co_bar_argAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? argument
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_arg, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_arg, start, -1, perr)
}

func _forethought_co_bar_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_arg"}
	// forethought_marker spaces? argument
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_arg}
	// forethought_marker spaces? argument
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? argument
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_pred_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_pred_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// arg_7 spaces? co_bar_pred
	// arg_7
	if !_accept(parser, _arg_7Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_accept(parser, _co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_pred_arg, start, pos, perr)
fail:
	return _memoize(parser, _coP_pred_arg, start, -1, perr)
}

func _coP_pred_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_pred_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_pred_arg"}
	// arg_7 spaces? co_bar_pred
	// arg_7
	if !_node(parser, _arg_7Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_node(parser, _co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_pred_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_pred_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_pred_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_pred_arg}
	// arg_7 spaces? co_bar_pred
	// arg_7
	if !_fail(parser, _arg_7Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_fail(parser, _co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_pred_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_pred_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// arg_7 spaces? co_bar_pred
	{
		var node0 string
		// arg_7
		if p, n := _arg_7Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_pred
		if p, n := _co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_pred_arg_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_arg_1
	if !_accept(parser, _forethought_coP_pred_arg_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_arg, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_arg, start, -1, perr)
}

func _forethought_coP_pred_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_arg"}
	// forethought_connective spaces? forethought_coP_pred_arg_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_arg_1
	if !_node(parser, _forethought_coP_pred_arg_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_arg}
	// forethought_connective spaces? forethought_coP_pred_arg_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_arg_1
	if !_fail(parser, _forethought_coP_pred_arg_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_pred_arg_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_pred_arg_1
		if p, n := _forethought_coP_pred_arg_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_arg_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_arg_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// argument spaces? forethought_co_bar_pred
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_accept(parser, _forethought_co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_arg_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_arg_1, start, -1, perr)
}

func _forethought_coP_pred_arg_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_arg_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_arg_1"}
	// argument spaces? forethought_co_bar_pred
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_node(parser, _forethought_co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_arg_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_arg_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_arg_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_arg_1}
	// argument spaces? forethought_co_bar_pred
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_fail(parser, _forethought_co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_arg_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_arg_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// argument spaces? forethought_co_bar_pred
	{
		var node0 string
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_pred
		if p, n := _forethought_co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// LU_arg_tone spaces? statement
	// LU_arg_tone
	if !_accept(parser, _LU_arg_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_arg, start, pos, perr)
fail:
	return _memoize(parser, _LU_arg, start, -1, perr)
}

func _LU_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_arg"}
	// LU_arg_tone spaces? statement
	// LU_arg_tone
	if !_node(parser, _LU_arg_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_arg}
	// LU_arg_tone spaces? statement
	// LU_arg_tone
	if !_fail(parser, _LU_arg_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// LU_arg_tone spaces? statement
	{
		var node0 string
		// LU_arg_tone
		if p, n := _LU_arg_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_arg_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_arg_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &LU arg_syllable
	// &LU
	{
		pos2 := pos
		perr4 := perr
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// arg_syllable
	if !_accept(parser, _arg_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_arg_tone, start, pos, perr)
fail:
	return _memoize(parser, _LU_arg_tone, start, -1, perr)
}

func _LU_arg_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_arg_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_arg_tone"}
	// &LU arg_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// LU
		if !_node(parser, _LUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// arg_syllable
	if !_node(parser, _arg_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_arg_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_arg_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_arg_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_arg_tone}
	// &LU arg_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&LU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// arg_syllable
	if !_fail(parser, _arg_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_arg_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_arg_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &LU arg_syllable
	{
		var node0 string
		// &LU
		{
			pos2 := pos
			// LU
			if p, n := _LUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// arg_syllable
		if p, n := _arg_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_arg_1 spaces? end_predicatizer?
	// MI_arg_1
	if !_accept(parser, _MI_arg_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _MI_arg, start, pos, perr)
fail:
	return _memoize(parser, _MI_arg, start, -1, perr)
}

func _MI_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_arg"}
	// MI_arg_1 spaces? end_predicatizer?
	// MI_arg_1
	if !_node(parser, _MI_arg_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_arg}
	// MI_arg_1 spaces? end_predicatizer?
	// MI_arg_1
	if !_fail(parser, _MI_arg_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_arg_1 spaces? end_predicatizer?
	{
		var node0 string
		// MI_arg_1
		if p, n := _MI_arg_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_arg_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_arg_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_arg_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_arg_tone
	if !_accept(parser, _MI_arg_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_accept(parser, _predicateAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_accept(parser, _adverbAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	return _memoize(parser, _MI_arg_1, start, pos, perr)
fail:
	return _memoize(parser, _MI_arg_1, start, -1, perr)
}

func _MI_arg_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_arg_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_arg_1"}
	// MI_arg_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_arg_tone
	if !_node(parser, _MI_arg_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// predicate/argument/adverb/prepositional_phrase
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// predicate
			if !_node(parser, _predicateNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail14
			}
			goto ok7
		fail14:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			goto fail
		ok7:
		}
		sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
		node.Kids = append(node.Kids[:nkids5], sub)
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_arg_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_arg_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_arg_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_arg_1}
	// MI_arg_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_arg_tone
	if !_fail(parser, _MI_arg_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_fail(parser, _predicateFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_fail(parser, _adverbFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_arg_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_arg_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_arg_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	{
		var node0 string
		// MI_arg_tone
		if p, n := _MI_arg_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (predicate/argument/adverb/prepositional_phrase)
		// predicate/argument/adverb/prepositional_phrase
		{
			pos8 := pos
			var node7 string
			// predicate
			if p, n := _predicateAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// adverb
			if p, n := _adverbAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// prepositional_phrase
			if p, n := _prepositional_phraseAction(parser, pos); n == nil {
				goto fail12
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail12:
			node0 = node7
			pos = pos8
			goto fail
		ok5:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_arg_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_arg_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MI arg_syllable
	// &MI
	{
		pos2 := pos
		perr4 := perr
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// arg_syllable
	if !_accept(parser, _arg_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MI_arg_tone, start, pos, perr)
fail:
	return _memoize(parser, _MI_arg_tone, start, -1, perr)
}

func _MI_arg_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_arg_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_arg_tone"}
	// &MI arg_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MI
		if !_node(parser, _MINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// arg_syllable
	if !_node(parser, _arg_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_arg_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_arg_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_arg_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_arg_tone}
	// &MI arg_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// arg_syllable
	if !_fail(parser, _arg_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_arg_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_arg_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MI arg_syllable
	{
		var node0 string
		// &MI
		{
			pos2 := pos
			// MI
			if p, n := _MIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// arg_syllable
		if p, n := _arg_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_argAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_arg, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_arg_1 spaces? end_predicatizer?
	// PO_arg_1
	if !_accept(parser, _PO_arg_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _PO_arg, start, pos, perr)
fail:
	return _memoize(parser, _PO_arg, start, -1, perr)
}

func _PO_argNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_arg}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_arg"}
	// PO_arg_1 spaces? end_predicatizer?
	// PO_arg_1
	if !_node(parser, _PO_arg_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_argFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_arg, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_arg",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_arg}
	// PO_arg_1 spaces? end_predicatizer?
	// PO_arg_1
	if !_fail(parser, _PO_arg_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_argAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_arg]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_arg}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_arg_1 spaces? end_predicatizer?
	{
		var node0 string
		// PO_arg_1
		if p, n := _PO_arg_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_arg_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_arg_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_arg_tone spaces? argument
	// PO_arg_tone
	if !_accept(parser, _PO_arg_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_arg_1, start, pos, perr)
fail:
	return _memoize(parser, _PO_arg_1, start, -1, perr)
}

func _PO_arg_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_arg_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_arg_1"}
	// PO_arg_tone spaces? argument
	// PO_arg_tone
	if !_node(parser, _PO_arg_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_arg_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_arg_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_arg_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_arg_1}
	// PO_arg_tone spaces? argument
	// PO_arg_tone
	if !_fail(parser, _PO_arg_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_arg_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_arg_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_arg_tone spaces? argument
	{
		var node0 string
		// PO_arg_tone
		if p, n := _PO_arg_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_arg_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_arg_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &PO arg_syllable
	// &PO
	{
		pos2 := pos
		perr4 := perr
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// arg_syllable
	if !_accept(parser, _arg_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_arg_tone, start, pos, perr)
fail:
	return _memoize(parser, _PO_arg_tone, start, -1, perr)
}

func _PO_arg_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_arg_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_arg_tone"}
	// &PO arg_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// arg_syllable
	if !_node(parser, _arg_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_arg_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_arg_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_arg_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_arg_tone}
	// &PO arg_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&PO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// arg_syllable
	if !_fail(parser, _arg_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_arg_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_arg_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &PO arg_syllable
	{
		var node0 string
		// &PO
		{
			pos2 := pos
			// PO
			if p, n := _POAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// arg_syllable
		if p, n := _arg_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quotation_argumentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _quotation_argument, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_argument spaces? end_quote
	// MO_argument
	if !_accept(parser, _MO_argumentAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_accept(parser, _end_quoteAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _quotation_argument, start, pos, perr)
fail:
	return _memoize(parser, _quotation_argument, start, -1, perr)
}

func _quotation_argumentNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_quotation_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_argument}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quotation_argument"}
	// MO_argument spaces? end_quote
	// MO_argument
	if !_node(parser, _MO_argumentNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_quote
	if !_node(parser, _end_quoteNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quotation_argumentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _quotation_argument, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quotation_argument",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quotation_argument}
	// MO_argument spaces? end_quote
	// MO_argument
	if !_fail(parser, _MO_argumentFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_fail(parser, _end_quoteFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quotation_argumentAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_quotation_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_argument}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_argument spaces? end_quote
	{
		var node0 string
		// MO_argument
		if p, n := _MO_argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_quote
		if p, n := _end_quoteAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_argumentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_argument, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_argument_tone spaces? discourse
	// MO_argument_tone
	if !_accept(parser, _MO_argument_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_accept(parser, _discourseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_argument, start, pos, perr)
fail:
	return _memoize(parser, _MO_argument, start, -1, perr)
}

func _MO_argumentNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_argument}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_argument"}
	// MO_argument_tone spaces? discourse
	// MO_argument_tone
	if !_node(parser, _MO_argument_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// discourse
	if !_node(parser, _discourseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_argumentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_argument, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_argument",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_argument}
	// MO_argument_tone spaces? discourse
	// MO_argument_tone
	if !_fail(parser, _MO_argument_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_fail(parser, _discourseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_argumentAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_argument}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_argument_tone spaces? discourse
	{
		var node0 string
		// MO_argument_tone
		if p, n := _MO_argument_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// discourse
		if p, n := _discourseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_argument_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_argument_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MO arg_syllable
	// &MO
	{
		pos2 := pos
		perr4 := perr
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// arg_syllable
	if !_accept(parser, _arg_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_argument_tone, start, pos, perr)
fail:
	return _memoize(parser, _MO_argument_tone, start, -1, perr)
}

func _MO_argument_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_argument_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_argument_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_argument_tone"}
	// &MO arg_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MO
		if !_node(parser, _MONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// arg_syllable
	if !_node(parser, _arg_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_argument_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_argument_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_argument_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_argument_tone}
	// &MO arg_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// arg_syllable
	if !_fail(parser, _arg_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_argument_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_argument_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_argument_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MO arg_syllable
	{
		var node0 string
		// &MO
		{
			pos2 := pos
			// MO
			if p, n := _MOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// arg_syllable
		if p, n := _arg_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_clauseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_clause, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_rel/LU_relative/relative_clause_1
	{
		pos3 := pos
		// coP_rel
		if !_accept(parser, _coP_relAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_relative
		if !_accept(parser, _LU_relativeAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// relative_clause_1
		if !_accept(parser, _relative_clause_1Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_clause, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause, start, -1, perr)
}

func _relative_clauseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_clause]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_clause"}
	// coP_rel/LU_relative/relative_clause_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_rel
		if !_node(parser, _coP_relNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_relative
		if !_node(parser, _LU_relativeNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_clause_1
		if !_node(parser, _relative_clause_1Node, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_clauseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_clause, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause}
	// coP_rel/LU_relative/relative_clause_1
	{
		pos3 := pos
		// coP_rel
		if !_fail(parser, _coP_relFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_relative
		if !_fail(parser, _LU_relativeFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// relative_clause_1
		if !_fail(parser, _relative_clause_1Fail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_clauseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_clause]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_rel/LU_relative/relative_clause_1
	{
		pos3 := pos
		var node2 string
		// coP_rel
		if p, n := _coP_relAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// LU_relative
		if p, n := _LU_relativeAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// relative_clause_1
		if p, n := _relative_clause_1Action(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_clause_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_clause_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_rel/relative_clause_2
	{
		pos3 := pos
		// forethought_coP_rel
		if !_accept(parser, _forethought_coP_relAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_clause_2
		if !_accept(parser, _relative_clause_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_clause_1, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause_1, start, -1, perr)
}

func _relative_clause_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_clause_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_clause_1"}
	// forethought_coP_rel/relative_clause_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_rel
		if !_node(parser, _forethought_coP_relNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_clause_2
		if !_node(parser, _relative_clause_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_clause_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_clause_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause_1}
	// forethought_coP_rel/relative_clause_2
	{
		pos3 := pos
		// forethought_coP_rel
		if !_fail(parser, _forethought_coP_relFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_clause_2
		if !_fail(parser, _relative_clause_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_clause_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_clause_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_rel/relative_clause_2
	{
		pos3 := pos
		var node2 string
		// forethought_coP_rel
		if p, n := _forethought_coP_relAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// relative_clause_2
		if p, n := _relative_clause_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_clause_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_clause_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_rel_statement/relative_clause_3
	{
		pos3 := pos
		// coP_rel_statement
		if !_accept(parser, _coP_rel_statementAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_clause_3
		if !_accept(parser, _relative_clause_3Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_clause_2, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause_2, start, -1, perr)
}

func _relative_clause_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_clause_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_clause_2"}
	// coP_rel_statement/relative_clause_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_rel_statement
		if !_node(parser, _coP_rel_statementNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_clause_3
		if !_node(parser, _relative_clause_3Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_clause_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_clause_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause_2}
	// coP_rel_statement/relative_clause_3
	{
		pos3 := pos
		// coP_rel_statement
		if !_fail(parser, _coP_rel_statementFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_clause_3
		if !_fail(parser, _relative_clause_3Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_clause_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_clause_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_rel_statement/relative_clause_3
	{
		pos3 := pos
		var node2 string
		// coP_rel_statement
		if p, n := _coP_rel_statementAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// relative_clause_3
		if p, n := _relative_clause_3Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_clause_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_clause_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_predication spaces? end_statement?
	// relative_predication
	if !_accept(parser, _relative_predicationAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_statement?
	{
		pos6 := pos
		// end_statement
		if !_accept(parser, _end_statementAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _relative_clause_3, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause_3, start, -1, perr)
}

func _relative_clause_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_clause_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_clause_3"}
	// relative_predication spaces? end_statement?
	// relative_predication
	if !_node(parser, _relative_predicationNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_statement?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_statement
		if !_node(parser, _end_statementNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_clause_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_clause_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause_3}
	// relative_predication spaces? end_statement?
	// relative_predication
	if !_fail(parser, _relative_predicationFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_statement?
	{
		pos6 := pos
		// end_statement
		if !_fail(parser, _end_statementFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_clause_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_clause_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_predication spaces? end_statement?
	{
		var node0 string
		// relative_predication
		if p, n := _relative_predicationAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_statement?
		{
			pos6 := pos
			// end_statement
			if p, n := _end_statementAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_predicationAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_predication, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_predicate spaces? terms?
	// relative_predicate
	if !_accept(parser, _relative_predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms?
	{
		pos6 := pos
		// terms
		if !_accept(parser, _termsAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _relative_predication, start, pos, perr)
fail:
	return _memoize(parser, _relative_predication, start, -1, perr)
}

func _relative_predicationNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predication}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_predication"}
	// relative_predicate spaces? terms?
	// relative_predicate
	if !_node(parser, _relative_predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// terms
		if !_node(parser, _termsNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_predicationFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_predication, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_predication",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_predication}
	// relative_predicate spaces? terms?
	// relative_predicate
	if !_fail(parser, _relative_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms?
	{
		pos6 := pos
		// terms
		if !_fail(parser, _termsFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_predicationAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predication}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_predicate spaces? terms?
	{
		var node0 string
		// relative_predicate
		if p, n := _relative_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms?
		{
			pos6 := pos
			// terms
			if p, n := _termsAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_rel_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_rel_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_clause_3 spaces? co_bar_statement
	// relative_clause_3
	if !_accept(parser, _relative_clause_3Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_accept(parser, _co_bar_statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_rel_statement, start, pos, perr)
fail:
	return _memoize(parser, _coP_rel_statement, start, -1, perr)
}

func _coP_rel_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_rel_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_rel_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_rel_statement"}
	// relative_clause_3 spaces? co_bar_statement
	// relative_clause_3
	if !_node(parser, _relative_clause_3Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_node(parser, _co_bar_statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_rel_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_rel_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_rel_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_rel_statement}
	// relative_clause_3 spaces? co_bar_statement
	// relative_clause_3
	if !_fail(parser, _relative_clause_3Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_fail(parser, _co_bar_statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_rel_statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_rel_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_rel_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_clause_3 spaces? co_bar_statement
	{
		var node0 string
		// relative_clause_3
		if p, n := _relative_clause_3Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_statement
		if p, n := _co_bar_statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_relAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_rel, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_clause_1 spaces? co_bar_rel
	// relative_clause_1
	if !_accept(parser, _relative_clause_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_rel
	if !_accept(parser, _co_bar_relAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_rel, start, pos, perr)
fail:
	return _memoize(parser, _coP_rel, start, -1, perr)
}

func _coP_relNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_rel}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_rel"}
	// relative_clause_1 spaces? co_bar_rel
	// relative_clause_1
	if !_node(parser, _relative_clause_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_rel
	if !_node(parser, _co_bar_relNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_relFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_rel, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_rel",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_rel}
	// relative_clause_1 spaces? co_bar_rel
	// relative_clause_1
	if !_fail(parser, _relative_clause_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_rel
	if !_fail(parser, _co_bar_relFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_relAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_rel}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_clause_1 spaces? co_bar_rel
	{
		var node0 string
		// relative_clause_1
		if p, n := _relative_clause_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_rel
		if p, n := _co_bar_relAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_relAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_rel, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? relative_clause_1
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// relative_clause_1
	if !_accept(parser, _relative_clause_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_rel, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_rel, start, -1, perr)
}

func _co_bar_relNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_rel}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_rel"}
	// connective spaces? relative_clause_1
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// relative_clause_1
	if !_node(parser, _relative_clause_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_relFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_rel, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_rel",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_rel}
	// connective spaces? relative_clause_1
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// relative_clause_1
	if !_fail(parser, _relative_clause_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_relAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_rel}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? relative_clause_1
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// relative_clause_1
		if p, n := _relative_clause_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_relAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_rel, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_rel_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_rel_1
	if !_accept(parser, _forethought_coP_rel_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_rel, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_rel, start, -1, perr)
}

func _forethought_coP_relNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_rel}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_rel"}
	// forethought_connective spaces? forethought_coP_rel_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_rel_1
	if !_node(parser, _forethought_coP_rel_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_relFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_rel, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_rel",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_rel}
	// forethought_connective spaces? forethought_coP_rel_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_rel_1
	if !_fail(parser, _forethought_coP_rel_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_relAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_rel}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_rel_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_rel_1
		if p, n := _forethought_coP_rel_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_rel_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_rel_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_clause spaces? forethought_co_bar_rel
	// relative_clause
	if !_accept(parser, _relative_clauseAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_rel
	if !_accept(parser, _forethought_co_bar_relAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_rel_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_rel_1, start, -1, perr)
}

func _forethought_coP_rel_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_rel_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_rel_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_rel_1"}
	// relative_clause spaces? forethought_co_bar_rel
	// relative_clause
	if !_node(parser, _relative_clauseNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_rel
	if !_node(parser, _forethought_co_bar_relNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_rel_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_rel_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_rel_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_rel_1}
	// relative_clause spaces? forethought_co_bar_rel
	// relative_clause
	if !_fail(parser, _relative_clauseFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_rel
	if !_fail(parser, _forethought_co_bar_relFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_rel_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_rel_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_rel_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_clause spaces? forethought_co_bar_rel
	{
		var node0 string
		// relative_clause
		if p, n := _relative_clauseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_rel
		if p, n := _forethought_co_bar_relAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_relAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_rel, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? relative_clause
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// relative_clause
	if !_accept(parser, _relative_clauseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_rel, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_rel, start, -1, perr)
}

func _forethought_co_bar_relNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_rel}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_rel"}
	// forethought_marker spaces? relative_clause
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// relative_clause
	if !_node(parser, _relative_clauseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_relFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_rel, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_rel",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_rel}
	// forethought_marker spaces? relative_clause
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// relative_clause
	if !_fail(parser, _relative_clauseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_relAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_rel]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_rel}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? relative_clause
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// relative_clause
		if p, n := _relative_clauseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial_relative_predicate/relative_predicate_1
	{
		pos3 := pos
		// serial_relative_predicate
		if !_accept(parser, _serial_relative_predicateAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_predicate_1
		if !_accept(parser, _relative_predicate_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _relative_predicate, start, -1, perr)
}

func _relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_predicate"}
	// serial_relative_predicate/relative_predicate_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial_relative_predicate
		if !_node(parser, _serial_relative_predicateNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_predicate_1
		if !_node(parser, _relative_predicate_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_predicate}
	// serial_relative_predicate/relative_predicate_1
	{
		pos3 := pos
		// serial_relative_predicate
		if !_fail(parser, _serial_relative_predicateFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_predicate_1
		if !_fail(parser, _relative_predicate_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// serial_relative_predicate/relative_predicate_1
	{
		pos3 := pos
		var node2 string
		// serial_relative_predicate
		if p, n := _serial_relative_predicateAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// relative_predicate_1
		if p, n := _relative_predicate_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_pred_relative_predicate/forethought_coP_pred_relative_predicate/relative_predicate_2
	{
		pos3 := pos
		// coP_pred_relative_predicate
		if !_accept(parser, _coP_pred_relative_predicateAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_relative_predicate
		if !_accept(parser, _forethought_coP_pred_relative_predicateAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// relative_predicate_2
		if !_accept(parser, _relative_predicate_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _relative_predicate_1, start, -1, perr)
}

func _relative_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_predicate_1"}
	// coP_pred_relative_predicate/forethought_coP_pred_relative_predicate/relative_predicate_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_pred_relative_predicate
		if !_node(parser, _coP_pred_relative_predicateNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_coP_pred_relative_predicate
		if !_node(parser, _forethought_coP_pred_relative_predicateNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_predicate_2
		if !_node(parser, _relative_predicate_2Node, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_predicate_1}
	// coP_pred_relative_predicate/forethought_coP_pred_relative_predicate/relative_predicate_2
	{
		pos3 := pos
		// coP_pred_relative_predicate
		if !_fail(parser, _coP_pred_relative_predicateFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_relative_predicate
		if !_fail(parser, _forethought_coP_pred_relative_predicateFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// relative_predicate_2
		if !_fail(parser, _relative_predicate_2Fail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_pred_relative_predicate/forethought_coP_pred_relative_predicate/relative_predicate_2
	{
		pos3 := pos
		var node2 string
		// coP_pred_relative_predicate
		if p, n := _coP_pred_relative_predicateAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_coP_pred_relative_predicate
		if p, n := _forethought_coP_pred_relative_predicateAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// relative_predicate_2
		if p, n := _relative_predicate_2Action(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_predicate_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_predicate_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_relative_predicate/PO_relative_predicate/quotation_relative_predicate/relative_predicate_3
	{
		pos3 := pos
		// MI_relative_predicate
		if !_accept(parser, _MI_relative_predicateAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_relative_predicate
		if !_accept(parser, _PO_relative_predicateAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// quotation_relative_predicate
		if !_accept(parser, _quotation_relative_predicateAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// relative_predicate_3
		if !_accept(parser, _relative_predicate_3Accepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_predicate_2, start, pos, perr)
fail:
	return _memoize(parser, _relative_predicate_2, start, -1, perr)
}

func _relative_predicate_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_predicate_2"}
	// MI_relative_predicate/PO_relative_predicate/quotation_relative_predicate/relative_predicate_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// MI_relative_predicate
		if !_node(parser, _MI_relative_predicateNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_relative_predicate
		if !_node(parser, _PO_relative_predicateNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// quotation_relative_predicate
		if !_node(parser, _quotation_relative_predicateNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_predicate_3
		if !_node(parser, _relative_predicate_3Node, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_predicate_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_predicate_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_predicate_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_predicate_2}
	// MI_relative_predicate/PO_relative_predicate/quotation_relative_predicate/relative_predicate_3
	{
		pos3 := pos
		// MI_relative_predicate
		if !_fail(parser, _MI_relative_predicateFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_relative_predicate
		if !_fail(parser, _PO_relative_predicateFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// quotation_relative_predicate
		if !_fail(parser, _quotation_relative_predicateFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// relative_predicate_3
		if !_fail(parser, _relative_predicate_3Fail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_predicate_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_relative_predicate/PO_relative_predicate/quotation_relative_predicate/relative_predicate_3
	{
		pos3 := pos
		var node2 string
		// MI_relative_predicate
		if p, n := _MI_relative_predicateAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// PO_relative_predicate
		if p, n := _PO_relative_predicateAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// quotation_relative_predicate
		if p, n := _quotation_relative_predicateAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// relative_predicate_3
		if p, n := _relative_predicate_3Action(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_predicate_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_predicate_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_syllable compound_syllable+/!function_word relative_syllable
	{
		pos3 := pos
		// relative_syllable compound_syllable+
		// relative_syllable
		if !_accept(parser, _relative_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word relative_syllable
		// !function_word
		{
			pos13 := pos
			perr15 := perr
			// function_word
			if !_accept(parser, _function_wordAccepts, &pos, &perr) {
				goto ok12
			}
			pos = pos13
			perr = _max(perr15, pos)
			goto fail10
		ok12:
			pos = pos13
			perr = perr15
		}
		// relative_syllable
		if !_accept(parser, _relative_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_predicate_3, start, pos, perr)
fail:
	return _memoize(parser, _relative_predicate_3, start, -1, perr)
}

func _relative_predicate_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_predicate_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_predicate_3"}
	// relative_syllable compound_syllable+/!function_word relative_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// relative_syllable compound_syllable+
		// relative_syllable
		if !_node(parser, _relative_syllableNode, node, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_node(parser, _compound_syllableNode, node, &pos) {
			goto fail4
		}
		for {
			nkids6 := len(node.Kids)
			pos7 := pos
			// compound_syllable
			if !_node(parser, _compound_syllableNode, node, &pos) {
				goto fail9
			}
			continue
		fail9:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
			break
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// !function_word relative_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(node.Kids)
			// function_word
			if !_node(parser, _function_wordNode, node, &pos) {
				goto ok12
			}
			pos = pos13
			node.Kids = node.Kids[:nkids14]
			goto fail10
		ok12:
			pos = pos13
			node.Kids = node.Kids[:nkids14]
		}
		// relative_syllable
		if !_node(parser, _relative_syllableNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_predicate_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_predicate_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_predicate_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_predicate_3}
	// relative_syllable compound_syllable+/!function_word relative_syllable
	{
		pos3 := pos
		// relative_syllable compound_syllable+
		// relative_syllable
		if !_fail(parser, _relative_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word relative_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(failure.Kids)
			// function_word
			if !_fail(parser, _function_wordFail, errPos, failure, &pos) {
				goto ok12
			}
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "!function_word",
				})
			}
			goto fail10
		ok12:
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
		}
		// relative_syllable
		if !_fail(parser, _relative_syllableFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_predicate_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_predicate_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_syllable compound_syllable+/!function_word relative_syllable
	{
		pos3 := pos
		var node2 string
		// relative_syllable compound_syllable+
		{
			var node5 string
			// relative_syllable
			if p, n := _relative_syllableAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
			// compound_syllable+
			{
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail4
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
			}
			for {
				pos7 := pos
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail9
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
				continue
			fail9:
				pos = pos7
				break
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// !function_word relative_syllable
		{
			var node11 string
			// !function_word
			{
				pos13 := pos
				// function_word
				if p, n := _function_wordAction(parser, pos); n == nil {
					goto ok12
				} else {
					pos = p
				}
				pos = pos13
				goto fail10
			ok12:
				pos = pos13
				node11 = ""
			}
			node, node11 = node+node11, ""
			// relative_syllable
			if p, n := _relative_syllableAction(parser, pos); n == nil {
				goto fail10
			} else {
				node11 = *n
				pos = p
			}
			node, node11 = node+node11, ""
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _serial_relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _serial_relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_predicate_1 spaces? serial_predicate
	// relative_predicate_1
	if !_accept(parser, _relative_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_accept(parser, _serial_predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _serial_relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _serial_relative_predicate, start, -1, perr)
}

func _serial_relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_serial_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "serial_relative_predicate"}
	// relative_predicate_1 spaces? serial_predicate
	// relative_predicate_1
	if !_node(parser, _relative_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_node(parser, _serial_predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _serial_relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _serial_relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "serial_relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _serial_relative_predicate}
	// relative_predicate_1 spaces? serial_predicate
	// relative_predicate_1
	if !_fail(parser, _relative_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_fail(parser, _serial_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _serial_relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_serial_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_predicate_1 spaces? serial_predicate
	{
		var node0 string
		// relative_predicate_1
		if p, n := _relative_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// serial_predicate
		if p, n := _serial_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_pred_relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_pred_relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_predicate_2 spaces? co_bar_pred
	// relative_predicate_2
	if !_accept(parser, _relative_predicate_2Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_accept(parser, _co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_pred_relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _coP_pred_relative_predicate, start, -1, perr)
}

func _coP_pred_relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_pred_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_pred_relative_predicate"}
	// relative_predicate_2 spaces? co_bar_pred
	// relative_predicate_2
	if !_node(parser, _relative_predicate_2Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_node(parser, _co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_pred_relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_pred_relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_pred_relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_pred_relative_predicate}
	// relative_predicate_2 spaces? co_bar_pred
	// relative_predicate_2
	if !_fail(parser, _relative_predicate_2Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_fail(parser, _co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_pred_relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_pred_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_predicate_2 spaces? co_bar_pred
	{
		var node0 string
		// relative_predicate_2
		if p, n := _relative_predicate_2Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_pred
		if p, n := _co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_pred_relative_predicate_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_relative_predicate_1
	if !_accept(parser, _forethought_coP_pred_relative_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_relative_predicate, start, -1, perr)
}

func _forethought_coP_pred_relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_relative_predicate"}
	// forethought_connective spaces? forethought_coP_pred_relative_predicate_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_relative_predicate_1
	if !_node(parser, _forethought_coP_pred_relative_predicate_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_relative_predicate}
	// forethought_connective spaces? forethought_coP_pred_relative_predicate_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_relative_predicate_1
	if !_fail(parser, _forethought_coP_pred_relative_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_pred_relative_predicate_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_pred_relative_predicate_1
		if p, n := _forethought_coP_pred_relative_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_relative_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_relative_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// relative_predicate spaces? forethought_co_bar_pred
	// relative_predicate
	if !_accept(parser, _relative_predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_accept(parser, _forethought_co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_relative_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_relative_predicate_1, start, -1, perr)
}

func _forethought_coP_pred_relative_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_relative_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_relative_predicate_1"}
	// relative_predicate spaces? forethought_co_bar_pred
	// relative_predicate
	if !_node(parser, _relative_predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_node(parser, _forethought_co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_relative_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_relative_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_relative_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_relative_predicate_1}
	// relative_predicate spaces? forethought_co_bar_pred
	// relative_predicate
	if !_fail(parser, _relative_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_fail(parser, _forethought_co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_relative_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_relative_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// relative_predicate spaces? forethought_co_bar_pred
	{
		var node0 string
		// relative_predicate
		if p, n := _relative_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_pred
		if p, n := _forethought_co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_relativeAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_relative, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// LU_relative_tone spaces? statement
	// LU_relative_tone
	if !_accept(parser, _LU_relative_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_relative, start, pos, perr)
fail:
	return _memoize(parser, _LU_relative, start, -1, perr)
}

func _LU_relativeNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_relative]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_relative}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_relative"}
	// LU_relative_tone spaces? statement
	// LU_relative_tone
	if !_node(parser, _LU_relative_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_relativeFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_relative, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_relative",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_relative}
	// LU_relative_tone spaces? statement
	// LU_relative_tone
	if !_fail(parser, _LU_relative_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_relativeAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_relative]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_relative}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// LU_relative_tone spaces? statement
	{
		var node0 string
		// LU_relative_tone
		if p, n := _LU_relative_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_relative_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_relative_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &LU relative_syllable
	// &LU
	{
		pos2 := pos
		perr4 := perr
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// relative_syllable
	if !_accept(parser, _relative_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_relative_tone, start, pos, perr)
fail:
	return _memoize(parser, _LU_relative_tone, start, -1, perr)
}

func _LU_relative_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_relative_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_relative_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_relative_tone"}
	// &LU relative_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// LU
		if !_node(parser, _LUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// relative_syllable
	if !_node(parser, _relative_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_relative_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_relative_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_relative_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_relative_tone}
	// &LU relative_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&LU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// relative_syllable
	if !_fail(parser, _relative_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_relative_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_relative_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_relative_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &LU relative_syllable
	{
		var node0 string
		// &LU
		{
			pos2 := pos
			// LU
			if p, n := _LUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// relative_syllable
		if p, n := _relative_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_relative_predicate_1 spaces? end_predicatizer?
	// MI_relative_predicate_1
	if !_accept(parser, _MI_relative_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _MI_relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _MI_relative_predicate, start, -1, perr)
}

func _MI_relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_relative_predicate"}
	// MI_relative_predicate_1 spaces? end_predicatizer?
	// MI_relative_predicate_1
	if !_node(parser, _MI_relative_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_relative_predicate}
	// MI_relative_predicate_1 spaces? end_predicatizer?
	// MI_relative_predicate_1
	if !_fail(parser, _MI_relative_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_relative_predicate_1 spaces? end_predicatizer?
	{
		var node0 string
		// MI_relative_predicate_1
		if p, n := _MI_relative_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_relative_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_relative_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_relative_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_relative_predicate_tone
	if !_accept(parser, _MI_relative_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_accept(parser, _predicateAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_accept(parser, _adverbAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	return _memoize(parser, _MI_relative_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _MI_relative_predicate_1, start, -1, perr)
}

func _MI_relative_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_relative_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_relative_predicate_1"}
	// MI_relative_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_relative_predicate_tone
	if !_node(parser, _MI_relative_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// predicate/argument/adverb/prepositional_phrase
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// predicate
			if !_node(parser, _predicateNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail14
			}
			goto ok7
		fail14:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			goto fail
		ok7:
		}
		sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
		node.Kids = append(node.Kids[:nkids5], sub)
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_relative_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_relative_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_relative_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_relative_predicate_1}
	// MI_relative_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_relative_predicate_tone
	if !_fail(parser, _MI_relative_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_fail(parser, _predicateFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_fail(parser, _adverbFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_relative_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_relative_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_relative_predicate_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	{
		var node0 string
		// MI_relative_predicate_tone
		if p, n := _MI_relative_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (predicate/argument/adverb/prepositional_phrase)
		// predicate/argument/adverb/prepositional_phrase
		{
			pos8 := pos
			var node7 string
			// predicate
			if p, n := _predicateAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// adverb
			if p, n := _adverbAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// prepositional_phrase
			if p, n := _prepositional_phraseAction(parser, pos); n == nil {
				goto fail12
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail12:
			node0 = node7
			pos = pos8
			goto fail
		ok5:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_relative_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_relative_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MI relative_syllable
	// &MI
	{
		pos2 := pos
		perr4 := perr
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// relative_syllable
	if !_accept(parser, _relative_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MI_relative_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _MI_relative_predicate_tone, start, -1, perr)
}

func _MI_relative_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_relative_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_relative_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_relative_predicate_tone"}
	// &MI relative_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MI
		if !_node(parser, _MINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// relative_syllable
	if !_node(parser, _relative_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_relative_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_relative_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_relative_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_relative_predicate_tone}
	// &MI relative_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// relative_syllable
	if !_fail(parser, _relative_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_relative_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_relative_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_relative_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MI relative_syllable
	{
		var node0 string
		// &MI
		{
			pos2 := pos
			// MI
			if p, n := _MIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// relative_syllable
		if p, n := _relative_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_relative_predicate_1 spaces? end_predicatizer?
	// PO_relative_predicate_1
	if !_accept(parser, _PO_relative_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _PO_relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _PO_relative_predicate, start, -1, perr)
}

func _PO_relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_relative_predicate"}
	// PO_relative_predicate_1 spaces? end_predicatizer?
	// PO_relative_predicate_1
	if !_node(parser, _PO_relative_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_relative_predicate}
	// PO_relative_predicate_1 spaces? end_predicatizer?
	// PO_relative_predicate_1
	if !_fail(parser, _PO_relative_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_relative_predicate_1 spaces? end_predicatizer?
	{
		var node0 string
		// PO_relative_predicate_1
		if p, n := _PO_relative_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_relative_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_relative_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_relative_predicate_tone spaces? argument
	// PO_relative_predicate_tone
	if !_accept(parser, _PO_relative_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_relative_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _PO_relative_predicate_1, start, -1, perr)
}

func _PO_relative_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_relative_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_relative_predicate_1"}
	// PO_relative_predicate_tone spaces? argument
	// PO_relative_predicate_tone
	if !_node(parser, _PO_relative_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_relative_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_relative_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_relative_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_relative_predicate_1}
	// PO_relative_predicate_tone spaces? argument
	// PO_relative_predicate_tone
	if !_fail(parser, _PO_relative_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_relative_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_relative_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_relative_predicate_tone spaces? argument
	{
		var node0 string
		// PO_relative_predicate_tone
		if p, n := _PO_relative_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_relative_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_relative_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &PO relative_syllable
	// &PO
	{
		pos2 := pos
		perr4 := perr
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// relative_syllable
	if !_accept(parser, _relative_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_relative_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _PO_relative_predicate_tone, start, -1, perr)
}

func _PO_relative_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_relative_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_relative_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_relative_predicate_tone"}
	// &PO relative_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// relative_syllable
	if !_node(parser, _relative_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_relative_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_relative_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_relative_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_relative_predicate_tone}
	// &PO relative_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&PO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// relative_syllable
	if !_fail(parser, _relative_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_relative_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_relative_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_relative_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &PO relative_syllable
	{
		var node0 string
		// &PO
		{
			pos2 := pos
			// PO
			if p, n := _POAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// relative_syllable
		if p, n := _relative_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quotation_relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _quotation_relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_relative_predicate spaces? end_quote
	// MO_relative_predicate
	if !_accept(parser, _MO_relative_predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_accept(parser, _end_quoteAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _quotation_relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _quotation_relative_predicate, start, -1, perr)
}

func _quotation_relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_quotation_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quotation_relative_predicate"}
	// MO_relative_predicate spaces? end_quote
	// MO_relative_predicate
	if !_node(parser, _MO_relative_predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_quote
	if !_node(parser, _end_quoteNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quotation_relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _quotation_relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quotation_relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quotation_relative_predicate}
	// MO_relative_predicate spaces? end_quote
	// MO_relative_predicate
	if !_fail(parser, _MO_relative_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_fail(parser, _end_quoteFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quotation_relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_quotation_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_relative_predicate spaces? end_quote
	{
		var node0 string
		// MO_relative_predicate
		if p, n := _MO_relative_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_quote
		if p, n := _end_quoteAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_relative_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_relative_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_relative_predicate_tone spaces? discourse
	// MO_relative_predicate_tone
	if !_accept(parser, _MO_relative_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_accept(parser, _discourseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_relative_predicate, start, pos, perr)
fail:
	return _memoize(parser, _MO_relative_predicate, start, -1, perr)
}

func _MO_relative_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_relative_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_relative_predicate"}
	// MO_relative_predicate_tone spaces? discourse
	// MO_relative_predicate_tone
	if !_node(parser, _MO_relative_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// discourse
	if !_node(parser, _discourseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_relative_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_relative_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_relative_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_relative_predicate}
	// MO_relative_predicate_tone spaces? discourse
	// MO_relative_predicate_tone
	if !_fail(parser, _MO_relative_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_fail(parser, _discourseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_relative_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_relative_predicate_tone spaces? discourse
	{
		var node0 string
		// MO_relative_predicate_tone
		if p, n := _MO_relative_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// discourse
		if p, n := _discourseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_relative_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_relative_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MO relative_syllable
	// &MO
	{
		pos2 := pos
		perr4 := perr
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// relative_syllable
	if !_accept(parser, _relative_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_relative_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _MO_relative_predicate_tone, start, -1, perr)
}

func _MO_relative_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_relative_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_relative_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_relative_predicate_tone"}
	// &MO relative_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MO
		if !_node(parser, _MONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// relative_syllable
	if !_node(parser, _relative_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_relative_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_relative_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_relative_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_relative_predicate_tone}
	// &MO relative_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// relative_syllable
	if !_fail(parser, _relative_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_relative_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_relative_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_relative_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MO relative_syllable
	{
		var node0 string
		// &MO
		{
			pos2 := pos
			// MO
			if p, n := _MOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// relative_syllable
		if p, n := _relative_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termsetAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _termset, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// termset_V/termset_IV/termset_III/termset_II
	{
		pos3 := pos
		// termset_V
		if !_accept(parser, _termset_VAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// termset_IV
		if !_accept(parser, _termset_IVAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// termset_III
		if !_accept(parser, _termset_IIIAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// termset_II
		if !_accept(parser, _termset_IIAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _termset, start, pos, perr)
fail:
	return _memoize(parser, _termset, start, -1, perr)
}

func _termsetNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_termset]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "termset"}
	// termset_V/termset_IV/termset_III/termset_II
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// termset_V
		if !_node(parser, _termset_VNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// termset_IV
		if !_node(parser, _termset_IVNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// termset_III
		if !_node(parser, _termset_IIINode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// termset_II
		if !_node(parser, _termset_IINode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _termsetFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _termset, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "termset",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _termset}
	// termset_V/termset_IV/termset_III/termset_II
	{
		pos3 := pos
		// termset_V
		if !_fail(parser, _termset_VFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// termset_IV
		if !_fail(parser, _termset_IVFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// termset_III
		if !_fail(parser, _termset_IIIFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// termset_II
		if !_fail(parser, _termset_IIFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _termsetAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_termset]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// termset_V/termset_IV/termset_III/termset_II
	{
		pos3 := pos
		var node2 string
		// termset_V
		if p, n := _termset_VAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// termset_IV
		if p, n := _termset_IVAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// termset_III
		if p, n := _termset_IIIAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// termset_II
		if p, n := _termset_IIAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termset_IIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _termset_II, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_term_II
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_II
	if !_accept(parser, _forethought_coP_term_IIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _termset_II, start, pos, perr)
fail:
	return _memoize(parser, _termset_II, start, -1, perr)
}

func _termset_IINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_termset_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_II}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "termset_II"}
	// forethought_connective spaces? forethought_coP_term_II
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_term_II
	if !_node(parser, _forethought_coP_term_IINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _termset_IIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _termset_II, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "termset_II",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _termset_II}
	// forethought_connective spaces? forethought_coP_term_II
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_II
	if !_fail(parser, _forethought_coP_term_IIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _termset_IIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_termset_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_II}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_term_II
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_term_II
		if p, n := _forethought_coP_term_IIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_term_IIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_term_II, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// terms_II spaces? forethought_co_bar_term_II
	// terms_II
	if !_accept(parser, _terms_IIAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_II
	if !_accept(parser, _forethought_co_bar_term_IIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_term_II, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_term_II, start, -1, perr)
}

func _forethought_coP_term_IINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_term_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_II}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_term_II"}
	// terms_II spaces? forethought_co_bar_term_II
	// terms_II
	if !_node(parser, _terms_IINode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_II
	if !_node(parser, _forethought_co_bar_term_IINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_term_IIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_term_II, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_term_II",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_term_II}
	// terms_II spaces? forethought_co_bar_term_II
	// terms_II
	if !_fail(parser, _terms_IIFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_II
	if !_fail(parser, _forethought_co_bar_term_IIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_term_IIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_term_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_II}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// terms_II spaces? forethought_co_bar_term_II
	{
		var node0 string
		// terms_II
		if p, n := _terms_IIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_term_II
		if p, n := _forethought_co_bar_term_IIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_term_IIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_term_II, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? terms_II
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_II
	if !_accept(parser, _terms_IIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_term_II, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_term_II, start, -1, perr)
}

func _forethought_co_bar_term_IINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_II}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_term_II"}
	// forethought_marker spaces? terms_II
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms_II
	if !_node(parser, _terms_IINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_term_IIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_term_II, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_term_II",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_term_II}
	// forethought_marker spaces? terms_II
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_II
	if !_fail(parser, _terms_IIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_term_IIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_II}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? terms_II
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms_II
		if p, n := _terms_IIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termset_IIIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _termset_III, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_term_III
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_III
	if !_accept(parser, _forethought_coP_term_IIIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _termset_III, start, pos, perr)
fail:
	return _memoize(parser, _termset_III, start, -1, perr)
}

func _termset_IIINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_termset_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_III}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "termset_III"}
	// forethought_connective spaces? forethought_coP_term_III
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_term_III
	if !_node(parser, _forethought_coP_term_IIINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _termset_IIIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _termset_III, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "termset_III",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _termset_III}
	// forethought_connective spaces? forethought_coP_term_III
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_III
	if !_fail(parser, _forethought_coP_term_IIIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _termset_IIIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_termset_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_III}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_term_III
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_term_III
		if p, n := _forethought_coP_term_IIIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_term_IIIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_term_III, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// terms_III spaces? forethought_co_bar_term_III
	// terms_III
	if !_accept(parser, _terms_IIIAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_III
	if !_accept(parser, _forethought_co_bar_term_IIIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_term_III, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_term_III, start, -1, perr)
}

func _forethought_coP_term_IIINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_term_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_III}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_term_III"}
	// terms_III spaces? forethought_co_bar_term_III
	// terms_III
	if !_node(parser, _terms_IIINode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_III
	if !_node(parser, _forethought_co_bar_term_IIINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_term_IIIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_term_III, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_term_III",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_term_III}
	// terms_III spaces? forethought_co_bar_term_III
	// terms_III
	if !_fail(parser, _terms_IIIFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_III
	if !_fail(parser, _forethought_co_bar_term_IIIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_term_IIIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_term_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_III}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// terms_III spaces? forethought_co_bar_term_III
	{
		var node0 string
		// terms_III
		if p, n := _terms_IIIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_term_III
		if p, n := _forethought_co_bar_term_IIIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_term_IIIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_term_III, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? terms_III
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_III
	if !_accept(parser, _terms_IIIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_term_III, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_term_III, start, -1, perr)
}

func _forethought_co_bar_term_IIINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_III}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_term_III"}
	// forethought_marker spaces? terms_III
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms_III
	if !_node(parser, _terms_IIINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_term_IIIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_term_III, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_term_III",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_term_III}
	// forethought_marker spaces? terms_III
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_III
	if !_fail(parser, _terms_IIIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_term_IIIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_III}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? terms_III
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms_III
		if p, n := _terms_IIIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termset_IVAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _termset_IV, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_term_IV
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_IV
	if !_accept(parser, _forethought_coP_term_IVAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _termset_IV, start, pos, perr)
fail:
	return _memoize(parser, _termset_IV, start, -1, perr)
}

func _termset_IVNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_termset_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_IV}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "termset_IV"}
	// forethought_connective spaces? forethought_coP_term_IV
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_term_IV
	if !_node(parser, _forethought_coP_term_IVNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _termset_IVFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _termset_IV, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "termset_IV",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _termset_IV}
	// forethought_connective spaces? forethought_coP_term_IV
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_IV
	if !_fail(parser, _forethought_coP_term_IVFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _termset_IVAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_termset_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_IV}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_term_IV
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_term_IV
		if p, n := _forethought_coP_term_IVAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_term_IVAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_term_IV, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// terms_IV spaces? forethought_co_bar_term_IV
	// terms_IV
	if !_accept(parser, _terms_IVAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_IV
	if !_accept(parser, _forethought_co_bar_term_IVAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_term_IV, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_term_IV, start, -1, perr)
}

func _forethought_coP_term_IVNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_term_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_IV}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_term_IV"}
	// terms_IV spaces? forethought_co_bar_term_IV
	// terms_IV
	if !_node(parser, _terms_IVNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_IV
	if !_node(parser, _forethought_co_bar_term_IVNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_term_IVFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_term_IV, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_term_IV",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_term_IV}
	// terms_IV spaces? forethought_co_bar_term_IV
	// terms_IV
	if !_fail(parser, _terms_IVFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_IV
	if !_fail(parser, _forethought_co_bar_term_IVFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_term_IVAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_term_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_IV}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// terms_IV spaces? forethought_co_bar_term_IV
	{
		var node0 string
		// terms_IV
		if p, n := _terms_IVAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_term_IV
		if p, n := _forethought_co_bar_term_IVAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_term_IVAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_term_IV, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? terms_IV
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_IV
	if !_accept(parser, _terms_IVAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_term_IV, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_term_IV, start, -1, perr)
}

func _forethought_co_bar_term_IVNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_IV}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_term_IV"}
	// forethought_marker spaces? terms_IV
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms_IV
	if !_node(parser, _terms_IVNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_term_IVFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_term_IV, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_term_IV",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_term_IV}
	// forethought_marker spaces? terms_IV
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_IV
	if !_fail(parser, _terms_IVFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_term_IVAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_IV}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? terms_IV
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms_IV
		if p, n := _terms_IVAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termset_VAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _termset_V, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_term_V
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_V
	if !_accept(parser, _forethought_coP_term_VAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _termset_V, start, pos, perr)
fail:
	return _memoize(parser, _termset_V, start, -1, perr)
}

func _termset_VNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_termset_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_V}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "termset_V"}
	// forethought_connective spaces? forethought_coP_term_V
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_term_V
	if !_node(parser, _forethought_coP_term_VNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _termset_VFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _termset_V, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "termset_V",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _termset_V}
	// forethought_connective spaces? forethought_coP_term_V
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_term_V
	if !_fail(parser, _forethought_coP_term_VFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _termset_VAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_termset_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset_V}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_term_V
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_term_V
		if p, n := _forethought_coP_term_VAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_term_VAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_term_V, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// terms_V spaces? forethought_co_bar_term_V
	// terms_V
	if !_accept(parser, _terms_VAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_V
	if !_accept(parser, _forethought_co_bar_term_VAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_term_V, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_term_V, start, -1, perr)
}

func _forethought_coP_term_VNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_term_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_V}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_term_V"}
	// terms_V spaces? forethought_co_bar_term_V
	// terms_V
	if !_node(parser, _terms_VNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_V
	if !_node(parser, _forethought_co_bar_term_VNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_term_VFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_term_V, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_term_V",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_term_V}
	// terms_V spaces? forethought_co_bar_term_V
	// terms_V
	if !_fail(parser, _terms_VFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_term_V
	if !_fail(parser, _forethought_co_bar_term_VFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_term_VAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_term_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_term_V}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// terms_V spaces? forethought_co_bar_term_V
	{
		var node0 string
		// terms_V
		if p, n := _terms_VAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_term_V
		if p, n := _forethought_co_bar_term_VAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_term_VAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_term_V, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? terms_V
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_V
	if !_accept(parser, _terms_VAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_term_V, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_term_V, start, -1, perr)
}

func _forethought_co_bar_term_VNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_V}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_term_V"}
	// forethought_marker spaces? terms_V
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms_V
	if !_node(parser, _terms_VNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_term_VFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_term_V, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_term_V",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_term_V}
	// forethought_marker spaces? terms_V
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_V
	if !_fail(parser, _terms_VFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_term_VAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_term_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_term_V}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? terms_V
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms_V
		if p, n := _terms_VAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_connectiveAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_connective, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// spaces? forethought_marker spaces? connective
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_connective, start, pos, perr)
fail:
	return _memoize(parser, _forethought_connective, start, -1, perr)
}

func _forethought_connectiveNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_connective]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_connective}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_connective"}
	// spaces? forethought_marker spaces? connective
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_connectiveFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_connective, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_connective",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_connective}
	// spaces? forethought_marker spaces? connective
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos6 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_connectiveAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_connective]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_connective}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// spaces? forethought_marker spaces? connective
	{
		var node0 string
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos6 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_IIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _terms_II, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// term spaces? term
	// term
	if !_accept(parser, _termAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// term
	if !_accept(parser, _termAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _terms_II, start, pos, perr)
fail:
	return _memoize(parser, _terms_II, start, -1, perr)
}

func _terms_IINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_terms_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_II}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "terms_II"}
	// term spaces? term
	// term
	if !_node(parser, _termNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// term
	if !_node(parser, _termNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_IIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _terms_II, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_II",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_II}
	// term spaces? term
	// term
	if !_fail(parser, _termFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// term
	if !_fail(parser, _termFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_IIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_terms_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_II}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// term spaces? term
	{
		var node0 string
		// term
		if p, n := _termAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// term
		if p, n := _termAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_IIIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _terms_III, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// term spaces? terms_II
	// term
	if !_accept(parser, _termAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_II
	if !_accept(parser, _terms_IIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _terms_III, start, pos, perr)
fail:
	return _memoize(parser, _terms_III, start, -1, perr)
}

func _terms_IIINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_terms_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_III}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "terms_III"}
	// term spaces? terms_II
	// term
	if !_node(parser, _termNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms_II
	if !_node(parser, _terms_IINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_IIIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _terms_III, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_III",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_III}
	// term spaces? terms_II
	// term
	if !_fail(parser, _termFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_II
	if !_fail(parser, _terms_IIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_IIIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_terms_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_III}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// term spaces? terms_II
	{
		var node0 string
		// term
		if p, n := _termAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms_II
		if p, n := _terms_IIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_IVAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _terms_IV, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// term spaces? terms_III
	// term
	if !_accept(parser, _termAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_III
	if !_accept(parser, _terms_IIIAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _terms_IV, start, pos, perr)
fail:
	return _memoize(parser, _terms_IV, start, -1, perr)
}

func _terms_IVNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_terms_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_IV}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "terms_IV"}
	// term spaces? terms_III
	// term
	if !_node(parser, _termNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms_III
	if !_node(parser, _terms_IIINode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_IVFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _terms_IV, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_IV",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_IV}
	// term spaces? terms_III
	// term
	if !_fail(parser, _termFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_III
	if !_fail(parser, _terms_IIIFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_IVAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_terms_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_IV}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// term spaces? terms_III
	{
		var node0 string
		// term
		if p, n := _termAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms_III
		if p, n := _terms_IIIAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_VAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _terms_V, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// term spaces? terms_IV
	// term
	if !_accept(parser, _termAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_IV
	if !_accept(parser, _terms_IVAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _terms_V, start, pos, perr)
fail:
	return _memoize(parser, _terms_V, start, -1, perr)
}

func _terms_VNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_terms_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_V}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "terms_V"}
	// term spaces? terms_IV
	// term
	if !_node(parser, _termNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms_IV
	if !_node(parser, _terms_IVNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_VFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _terms_V, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_V",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_V}
	// term spaces? terms_IV
	// term
	if !_fail(parser, _termFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms_IV
	if !_fail(parser, _terms_IVFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_VAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_terms_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_V}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// term spaces? terms_IV
	{
		var node0 string
		// term
		if p, n := _termAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms_IV
		if p, n := _terms_IVAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_adverb/adverb_1
	{
		pos3 := pos
		// coP_adverb
		if !_accept(parser, _coP_adverbAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_1
		if !_accept(parser, _adverb_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _adverb, start, pos, perr)
fail:
	return _memoize(parser, _adverb, start, -1, perr)
}

func _adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb"}
	// coP_adverb/adverb_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_adverb
		if !_node(parser, _coP_adverbNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_1
		if !_node(parser, _adverb_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb}
	// coP_adverb/adverb_1
	{
		pos3 := pos
		// coP_adverb
		if !_fail(parser, _coP_adverbFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_1
		if !_fail(parser, _adverb_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_adverb/adverb_1
	{
		pos3 := pos
		var node2 string
		// coP_adverb
		if p, n := _coP_adverbAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// adverb_1
		if p, n := _adverb_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _adverb_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_adverb/adverb_2
	{
		pos3 := pos
		// forethought_coP_adverb
		if !_accept(parser, _forethought_coP_adverbAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_2
		if !_accept(parser, _adverb_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _adverb_1, start, pos, perr)
fail:
	return _memoize(parser, _adverb_1, start, -1, perr)
}

func _adverb_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb_1"}
	// forethought_coP_adverb/adverb_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_adverb
		if !_node(parser, _forethought_coP_adverbNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_2
		if !_node(parser, _adverb_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_1}
	// forethought_coP_adverb/adverb_2
	{
		pos3 := pos
		// forethought_coP_adverb
		if !_fail(parser, _forethought_coP_adverbFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_2
		if !_fail(parser, _adverb_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_adverb/adverb_2
	{
		pos3 := pos
		var node2 string
		// forethought_coP_adverb
		if p, n := _forethought_coP_adverbAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// adverb_2
		if p, n := _adverb_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _adverb_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial_adverb/adverb_3
	{
		pos3 := pos
		// serial_adverb
		if !_accept(parser, _serial_adverbAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_3
		if !_accept(parser, _adverb_3Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _adverb_2, start, pos, perr)
fail:
	return _memoize(parser, _adverb_2, start, -1, perr)
}

func _adverb_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb_2"}
	// serial_adverb/adverb_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial_adverb
		if !_node(parser, _serial_adverbNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_3
		if !_node(parser, _adverb_3Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_2}
	// serial_adverb/adverb_3
	{
		pos3 := pos
		// serial_adverb
		if !_fail(parser, _serial_adverbFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_3
		if !_fail(parser, _adverb_3Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// serial_adverb/adverb_3
	{
		pos3 := pos
		var node2 string
		// serial_adverb
		if p, n := _serial_adverbAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// adverb_3
		if p, n := _adverb_3Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _adverb_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_pred_adverb/forethought_coP_pred_adverb/LU_adverb/MI_adverb/PO_adverb/quotation_adverb/adverb_4
	{
		pos3 := pos
		// coP_pred_adverb
		if !_accept(parser, _coP_pred_adverbAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_adverb
		if !_accept(parser, _forethought_coP_pred_adverbAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_adverb
		if !_accept(parser, _LU_adverbAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_adverb
		if !_accept(parser, _MI_adverbAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_adverb
		if !_accept(parser, _PO_adverbAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// quotation_adverb
		if !_accept(parser, _quotation_adverbAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// adverb_4
		if !_accept(parser, _adverb_4Accepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _adverb_3, start, pos, perr)
fail:
	return _memoize(parser, _adverb_3, start, -1, perr)
}

func _adverb_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb_3"}
	// coP_pred_adverb/forethought_coP_pred_adverb/LU_adverb/MI_adverb/PO_adverb/quotation_adverb/adverb_4
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_pred_adverb
		if !_node(parser, _coP_pred_adverbNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_coP_pred_adverb
		if !_node(parser, _forethought_coP_pred_adverbNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_adverb
		if !_node(parser, _LU_adverbNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_adverb
		if !_node(parser, _MI_adverbNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_adverb
		if !_node(parser, _PO_adverbNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// quotation_adverb
		if !_node(parser, _quotation_adverbNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_4
		if !_node(parser, _adverb_4Node, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_3}
	// coP_pred_adverb/forethought_coP_pred_adverb/LU_adverb/MI_adverb/PO_adverb/quotation_adverb/adverb_4
	{
		pos3 := pos
		// coP_pred_adverb
		if !_fail(parser, _coP_pred_adverbFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_adverb
		if !_fail(parser, _forethought_coP_pred_adverbFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_adverb
		if !_fail(parser, _LU_adverbFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_adverb
		if !_fail(parser, _MI_adverbFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_adverb
		if !_fail(parser, _PO_adverbFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// quotation_adverb
		if !_fail(parser, _quotation_adverbFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// adverb_4
		if !_fail(parser, _adverb_4Fail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_pred_adverb/forethought_coP_pred_adverb/LU_adverb/MI_adverb/PO_adverb/quotation_adverb/adverb_4
	{
		pos3 := pos
		var node2 string
		// coP_pred_adverb
		if p, n := _coP_pred_adverbAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_coP_pred_adverb
		if p, n := _forethought_coP_pred_adverbAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// LU_adverb
		if p, n := _LU_adverbAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// MI_adverb
		if p, n := _MI_adverbAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// PO_adverb
		if p, n := _PO_adverbAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// quotation_adverb
		if p, n := _quotation_adverbAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// adverb_4
		if p, n := _adverb_4Action(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _adverb_4Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb_4, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// adverb_syllable compound_syllable+/!function_word adverb_syllable
	{
		pos3 := pos
		// adverb_syllable compound_syllable+
		// adverb_syllable
		if !_accept(parser, _adverb_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word adverb_syllable
		// !function_word
		{
			pos13 := pos
			perr15 := perr
			// function_word
			if !_accept(parser, _function_wordAccepts, &pos, &perr) {
				goto ok12
			}
			pos = pos13
			perr = _max(perr15, pos)
			goto fail10
		ok12:
			pos = pos13
			perr = perr15
		}
		// adverb_syllable
		if !_accept(parser, _adverb_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _adverb_4, start, pos, perr)
fail:
	return _memoize(parser, _adverb_4, start, -1, perr)
}

func _adverb_4Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_4}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb_4"}
	// adverb_syllable compound_syllable+/!function_word adverb_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// adverb_syllable compound_syllable+
		// adverb_syllable
		if !_node(parser, _adverb_syllableNode, node, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_node(parser, _compound_syllableNode, node, &pos) {
			goto fail4
		}
		for {
			nkids6 := len(node.Kids)
			pos7 := pos
			// compound_syllable
			if !_node(parser, _compound_syllableNode, node, &pos) {
				goto fail9
			}
			continue
		fail9:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
			break
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// !function_word adverb_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(node.Kids)
			// function_word
			if !_node(parser, _function_wordNode, node, &pos) {
				goto ok12
			}
			pos = pos13
			node.Kids = node.Kids[:nkids14]
			goto fail10
		ok12:
			pos = pos13
			node.Kids = node.Kids[:nkids14]
		}
		// adverb_syllable
		if !_node(parser, _adverb_syllableNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_4Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb_4, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_4",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_4}
	// adverb_syllable compound_syllable+/!function_word adverb_syllable
	{
		pos3 := pos
		// adverb_syllable compound_syllable+
		// adverb_syllable
		if !_fail(parser, _adverb_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word adverb_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(failure.Kids)
			// function_word
			if !_fail(parser, _function_wordFail, errPos, failure, &pos) {
				goto ok12
			}
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "!function_word",
				})
			}
			goto fail10
		ok12:
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
		}
		// adverb_syllable
		if !_fail(parser, _adverb_syllableFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_4Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_4}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// adverb_syllable compound_syllable+/!function_word adverb_syllable
	{
		pos3 := pos
		var node2 string
		// adverb_syllable compound_syllable+
		{
			var node5 string
			// adverb_syllable
			if p, n := _adverb_syllableAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
			// compound_syllable+
			{
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail4
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
			}
			for {
				pos7 := pos
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail9
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
				continue
			fail9:
				pos = pos7
				break
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// !function_word adverb_syllable
		{
			var node11 string
			// !function_word
			{
				pos13 := pos
				// function_word
				if p, n := _function_wordAction(parser, pos); n == nil {
					goto ok12
				} else {
					pos = p
				}
				pos = pos13
				goto fail10
			ok12:
				pos = pos13
				node11 = ""
			}
			node, node11 = node+node11, ""
			// adverb_syllable
			if p, n := _adverb_syllableAction(parser, pos); n == nil {
				goto fail10
			} else {
				node11 = *n
				pos = p
			}
			node, node11 = node+node11, ""
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// adverb_1 spaces? co_bar_adverb
	// adverb_1
	if !_accept(parser, _adverb_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_adverb
	if !_accept(parser, _co_bar_adverbAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_adverb, start, pos, perr)
fail:
	return _memoize(parser, _coP_adverb, start, -1, perr)
}

func _coP_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_adverb"}
	// adverb_1 spaces? co_bar_adverb
	// adverb_1
	if !_node(parser, _adverb_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_adverb
	if !_node(parser, _co_bar_adverbNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_adverb}
	// adverb_1 spaces? co_bar_adverb
	// adverb_1
	if !_fail(parser, _adverb_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_adverb
	if !_fail(parser, _co_bar_adverbFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// adverb_1 spaces? co_bar_adverb
	{
		var node0 string
		// adverb_1
		if p, n := _adverb_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_adverb
		if p, n := _co_bar_adverbAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? adverb
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// adverb
	if !_accept(parser, _adverbAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_adverb, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_adverb, start, -1, perr)
}

func _co_bar_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_adverb"}
	// connective spaces? adverb
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// adverb
	if !_node(parser, _adverbNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_adverb}
	// connective spaces? adverb
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// adverb
	if !_fail(parser, _adverbFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? adverb
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// adverb
		if p, n := _adverbAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_adverb_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_adverb_1
	if !_accept(parser, _forethought_coP_adverb_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_adverb, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_adverb, start, -1, perr)
}

func _forethought_coP_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_adverb"}
	// forethought_connective spaces? forethought_coP_adverb_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_adverb_1
	if !_node(parser, _forethought_coP_adverb_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_adverb}
	// forethought_connective spaces? forethought_coP_adverb_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_adverb_1
	if !_fail(parser, _forethought_coP_adverb_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_adverb_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_adverb_1
		if p, n := _forethought_coP_adverb_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_adverb_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_adverb_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// adverb spaces? forethought_co_bar_adverb
	// adverb
	if !_accept(parser, _adverbAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_adverb
	if !_accept(parser, _forethought_co_bar_adverbAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_adverb_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_adverb_1, start, -1, perr)
}

func _forethought_coP_adverb_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_adverb_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_adverb_1"}
	// adverb spaces? forethought_co_bar_adverb
	// adverb
	if !_node(parser, _adverbNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_adverb
	if !_node(parser, _forethought_co_bar_adverbNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_adverb_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_adverb_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_adverb_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_adverb_1}
	// adverb spaces? forethought_co_bar_adverb
	// adverb
	if !_fail(parser, _adverbFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_adverb
	if !_fail(parser, _forethought_co_bar_adverbFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_adverb_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_adverb_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// adverb spaces? forethought_co_bar_adverb
	{
		var node0 string
		// adverb
		if p, n := _adverbAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_adverb
		if p, n := _forethought_co_bar_adverbAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? adverb
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// adverb
	if !_accept(parser, _adverbAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_adverb, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_adverb, start, -1, perr)
}

func _forethought_co_bar_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_adverb"}
	// forethought_marker spaces? adverb
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// adverb
	if !_node(parser, _adverbNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_adverb}
	// forethought_marker spaces? adverb
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// adverb
	if !_fail(parser, _adverbFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? adverb
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// adverb
		if p, n := _adverbAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _serial_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _serial_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// adverb_3 spaces? serial_predicate
	// adverb_3
	if !_accept(parser, _adverb_3Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_accept(parser, _serial_predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _serial_adverb, start, pos, perr)
fail:
	return _memoize(parser, _serial_adverb, start, -1, perr)
}

func _serial_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_serial_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "serial_adverb"}
	// adverb_3 spaces? serial_predicate
	// adverb_3
	if !_node(parser, _adverb_3Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_node(parser, _serial_predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _serial_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _serial_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "serial_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _serial_adverb}
	// adverb_3 spaces? serial_predicate
	// adverb_3
	if !_fail(parser, _adverb_3Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_fail(parser, _serial_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _serial_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_serial_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// adverb_3 spaces? serial_predicate
	{
		var node0 string
		// adverb_3
		if p, n := _adverb_3Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// serial_predicate
		if p, n := _serial_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_pred_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_pred_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// adverb_4 spaces? co_bar_pred
	// adverb_4
	if !_accept(parser, _adverb_4Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_accept(parser, _co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_pred_adverb, start, pos, perr)
fail:
	return _memoize(parser, _coP_pred_adverb, start, -1, perr)
}

func _coP_pred_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_pred_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_pred_adverb"}
	// adverb_4 spaces? co_bar_pred
	// adverb_4
	if !_node(parser, _adverb_4Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_node(parser, _co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_pred_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_pred_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_pred_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_pred_adverb}
	// adverb_4 spaces? co_bar_pred
	// adverb_4
	if !_fail(parser, _adverb_4Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_fail(parser, _co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_pred_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_pred_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// adverb_4 spaces? co_bar_pred
	{
		var node0 string
		// adverb_4
		if p, n := _adverb_4Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_pred
		if p, n := _co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_pred_adverb_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_adverb_1
	if !_accept(parser, _forethought_coP_pred_adverb_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_adverb, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_adverb, start, -1, perr)
}

func _forethought_coP_pred_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_adverb"}
	// forethought_connective spaces? forethought_coP_pred_adverb_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_adverb_1
	if !_node(parser, _forethought_coP_pred_adverb_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_adverb}
	// forethought_connective spaces? forethought_coP_pred_adverb_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_adverb_1
	if !_fail(parser, _forethought_coP_pred_adverb_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_pred_adverb_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_pred_adverb_1
		if p, n := _forethought_coP_pred_adverb_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_adverb_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_adverb_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// adverb spaces? forethought_co_bar_pred
	// adverb
	if !_accept(parser, _adverbAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_accept(parser, _forethought_co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_adverb_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_adverb_1, start, -1, perr)
}

func _forethought_coP_pred_adverb_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_adverb_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_adverb_1"}
	// adverb spaces? forethought_co_bar_pred
	// adverb
	if !_node(parser, _adverbNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_node(parser, _forethought_co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_adverb_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_adverb_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_adverb_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_adverb_1}
	// adverb spaces? forethought_co_bar_pred
	// adverb
	if !_fail(parser, _adverbFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_fail(parser, _forethought_co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_adverb_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_adverb_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// adverb spaces? forethought_co_bar_pred
	{
		var node0 string
		// adverb
		if p, n := _adverbAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_pred
		if p, n := _forethought_co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// LU_adverb_tone spaces? statement
	// LU_adverb_tone
	if !_accept(parser, _LU_adverb_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_adverb, start, pos, perr)
fail:
	return _memoize(parser, _LU_adverb, start, -1, perr)
}

func _LU_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_adverb"}
	// LU_adverb_tone spaces? statement
	// LU_adverb_tone
	if !_node(parser, _LU_adverb_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_adverb}
	// LU_adverb_tone spaces? statement
	// LU_adverb_tone
	if !_fail(parser, _LU_adverb_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// LU_adverb_tone spaces? statement
	{
		var node0 string
		// LU_adverb_tone
		if p, n := _LU_adverb_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_adverb_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_adverb_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &LU adverb_syllable
	// &LU
	{
		pos2 := pos
		perr4 := perr
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// adverb_syllable
	if !_accept(parser, _adverb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_adverb_tone, start, pos, perr)
fail:
	return _memoize(parser, _LU_adverb_tone, start, -1, perr)
}

func _LU_adverb_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_adverb_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_adverb_tone"}
	// &LU adverb_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// LU
		if !_node(parser, _LUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// adverb_syllable
	if !_node(parser, _adverb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_adverb_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_adverb_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_adverb_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_adverb_tone}
	// &LU adverb_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&LU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// adverb_syllable
	if !_fail(parser, _adverb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_adverb_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_adverb_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &LU adverb_syllable
	{
		var node0 string
		// &LU
		{
			pos2 := pos
			// LU
			if p, n := _LUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// adverb_syllable
		if p, n := _adverb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_adverb_1 spaces? end_predicatizer?
	// MI_adverb_1
	if !_accept(parser, _MI_adverb_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _MI_adverb, start, pos, perr)
fail:
	return _memoize(parser, _MI_adverb, start, -1, perr)
}

func _MI_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_adverb"}
	// MI_adverb_1 spaces? end_predicatizer?
	// MI_adverb_1
	if !_node(parser, _MI_adverb_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_adverb}
	// MI_adverb_1 spaces? end_predicatizer?
	// MI_adverb_1
	if !_fail(parser, _MI_adverb_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_adverb_1 spaces? end_predicatizer?
	{
		var node0 string
		// MI_adverb_1
		if p, n := _MI_adverb_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_adverb_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_adverb_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_adverb_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_adverb_tone
	if !_accept(parser, _MI_adverb_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_accept(parser, _predicateAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_accept(parser, _adverbAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	return _memoize(parser, _MI_adverb_1, start, pos, perr)
fail:
	return _memoize(parser, _MI_adverb_1, start, -1, perr)
}

func _MI_adverb_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_adverb_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_adverb_1"}
	// MI_adverb_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_adverb_tone
	if !_node(parser, _MI_adverb_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// predicate/argument/adverb/prepositional_phrase
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// predicate
			if !_node(parser, _predicateNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail14
			}
			goto ok7
		fail14:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			goto fail
		ok7:
		}
		sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
		node.Kids = append(node.Kids[:nkids5], sub)
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_adverb_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_adverb_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_adverb_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_adverb_1}
	// MI_adverb_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_adverb_tone
	if !_fail(parser, _MI_adverb_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_fail(parser, _predicateFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_fail(parser, _adverbFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_adverb_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_adverb_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_adverb_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	{
		var node0 string
		// MI_adverb_tone
		if p, n := _MI_adverb_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (predicate/argument/adverb/prepositional_phrase)
		// predicate/argument/adverb/prepositional_phrase
		{
			pos8 := pos
			var node7 string
			// predicate
			if p, n := _predicateAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// adverb
			if p, n := _adverbAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// prepositional_phrase
			if p, n := _prepositional_phraseAction(parser, pos); n == nil {
				goto fail12
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail12:
			node0 = node7
			pos = pos8
			goto fail
		ok5:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_adverb_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_adverb_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MI adverb_syllable
	// &MI
	{
		pos2 := pos
		perr4 := perr
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// adverb_syllable
	if !_accept(parser, _adverb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MI_adverb_tone, start, pos, perr)
fail:
	return _memoize(parser, _MI_adverb_tone, start, -1, perr)
}

func _MI_adverb_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_adverb_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_adverb_tone"}
	// &MI adverb_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MI
		if !_node(parser, _MINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// adverb_syllable
	if !_node(parser, _adverb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_adverb_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_adverb_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_adverb_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_adverb_tone}
	// &MI adverb_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// adverb_syllable
	if !_fail(parser, _adverb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_adverb_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_adverb_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MI adverb_syllable
	{
		var node0 string
		// &MI
		{
			pos2 := pos
			// MI
			if p, n := _MIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// adverb_syllable
		if p, n := _adverb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_adverb_1 spaces? end_predicatizer?
	// PO_adverb_1
	if !_accept(parser, _PO_adverb_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _PO_adverb, start, pos, perr)
fail:
	return _memoize(parser, _PO_adverb, start, -1, perr)
}

func _PO_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_adverb"}
	// PO_adverb_1 spaces? end_predicatizer?
	// PO_adverb_1
	if !_node(parser, _PO_adverb_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_adverb}
	// PO_adverb_1 spaces? end_predicatizer?
	// PO_adverb_1
	if !_fail(parser, _PO_adverb_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_adverb_1 spaces? end_predicatizer?
	{
		var node0 string
		// PO_adverb_1
		if p, n := _PO_adverb_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_adverb_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_adverb_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_adverb_tone spaces? argument
	// PO_adverb_tone
	if !_accept(parser, _PO_adverb_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_adverb_1, start, pos, perr)
fail:
	return _memoize(parser, _PO_adverb_1, start, -1, perr)
}

func _PO_adverb_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_adverb_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_adverb_1"}
	// PO_adverb_tone spaces? argument
	// PO_adverb_tone
	if !_node(parser, _PO_adverb_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_adverb_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_adverb_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_adverb_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_adverb_1}
	// PO_adverb_tone spaces? argument
	// PO_adverb_tone
	if !_fail(parser, _PO_adverb_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_adverb_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_adverb_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_adverb_tone spaces? argument
	{
		var node0 string
		// PO_adverb_tone
		if p, n := _PO_adverb_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_adverb_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_adverb_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &PO adverb_syllable
	// &PO
	{
		pos2 := pos
		perr4 := perr
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// adverb_syllable
	if !_accept(parser, _adverb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_adverb_tone, start, pos, perr)
fail:
	return _memoize(parser, _PO_adverb_tone, start, -1, perr)
}

func _PO_adverb_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_adverb_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_adverb_tone"}
	// &PO adverb_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// adverb_syllable
	if !_node(parser, _adverb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_adverb_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_adverb_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_adverb_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_adverb_tone}
	// &PO adverb_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&PO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// adverb_syllable
	if !_fail(parser, _adverb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_adverb_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_adverb_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &PO adverb_syllable
	{
		var node0 string
		// &PO
		{
			pos2 := pos
			// PO
			if p, n := _POAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// adverb_syllable
		if p, n := _adverb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quotation_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _quotation_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_adverb spaces? end_quote
	// MO_adverb
	if !_accept(parser, _MO_adverbAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_accept(parser, _end_quoteAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _quotation_adverb, start, pos, perr)
fail:
	return _memoize(parser, _quotation_adverb, start, -1, perr)
}

func _quotation_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_quotation_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quotation_adverb"}
	// MO_adverb spaces? end_quote
	// MO_adverb
	if !_node(parser, _MO_adverbNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_quote
	if !_node(parser, _end_quoteNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quotation_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _quotation_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quotation_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quotation_adverb}
	// MO_adverb spaces? end_quote
	// MO_adverb
	if !_fail(parser, _MO_adverbFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_fail(parser, _end_quoteFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quotation_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_quotation_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_adverb spaces? end_quote
	{
		var node0 string
		// MO_adverb
		if p, n := _MO_adverbAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_quote
		if p, n := _end_quoteAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_adverb_tone spaces? discourse
	// MO_adverb_tone
	if !_accept(parser, _MO_adverb_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_accept(parser, _discourseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_adverb, start, pos, perr)
fail:
	return _memoize(parser, _MO_adverb, start, -1, perr)
}

func _MO_adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_adverb}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_adverb"}
	// MO_adverb_tone spaces? discourse
	// MO_adverb_tone
	if !_node(parser, _MO_adverb_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// discourse
	if !_node(parser, _discourseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_adverb}
	// MO_adverb_tone spaces? discourse
	// MO_adverb_tone
	if !_fail(parser, _MO_adverb_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_fail(parser, _discourseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_adverbAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_adverb_tone spaces? discourse
	{
		var node0 string
		// MO_adverb_tone
		if p, n := _MO_adverb_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// discourse
		if p, n := _discourseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_adverb_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_adverb_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MO adverb_syllable
	// &MO
	{
		pos2 := pos
		perr4 := perr
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// adverb_syllable
	if !_accept(parser, _adverb_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_adverb_tone, start, pos, perr)
fail:
	return _memoize(parser, _MO_adverb_tone, start, -1, perr)
}

func _MO_adverb_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_adverb_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_adverb_tone"}
	// &MO adverb_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MO
		if !_node(parser, _MONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// adverb_syllable
	if !_node(parser, _adverb_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_adverb_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_adverb_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_adverb_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_adverb_tone}
	// &MO adverb_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// adverb_syllable
	if !_fail(parser, _adverb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_adverb_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_adverb_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MO adverb_syllable
	{
		var node0 string
		// &MO
		{
			pos2 := pos
			// MO
			if p, n := _MOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// adverb_syllable
		if p, n := _adverb_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prepositional_phraseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _prepositional_phrase, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_prepositional_phrase/prepositional_phrase_1
	{
		pos3 := pos
		// coP_prepositional_phrase
		if !_accept(parser, _coP_prepositional_phraseAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_1
		if !_accept(parser, _prepositional_phrase_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _prepositional_phrase, start, pos, perr)
fail:
	return _memoize(parser, _prepositional_phrase, start, -1, perr)
}

func _prepositional_phraseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "prepositional_phrase"}
	// coP_prepositional_phrase/prepositional_phrase_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_prepositional_phrase
		if !_node(parser, _coP_prepositional_phraseNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// prepositional_phrase_1
		if !_node(parser, _prepositional_phrase_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prepositional_phraseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _prepositional_phrase, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prepositional_phrase",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prepositional_phrase}
	// coP_prepositional_phrase/prepositional_phrase_1
	{
		pos3 := pos
		// coP_prepositional_phrase
		if !_fail(parser, _coP_prepositional_phraseFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_1
		if !_fail(parser, _prepositional_phrase_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prepositional_phraseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_prepositional_phrase/prepositional_phrase_1
	{
		pos3 := pos
		var node2 string
		// coP_prepositional_phrase
		if p, n := _coP_prepositional_phraseAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// prepositional_phrase_1
		if p, n := _prepositional_phrase_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prepositional_phrase_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _prepositional_phrase_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_prepositional_phrase/prepositional_phrase_2
	{
		pos3 := pos
		// forethought_coP_prepositional_phrase
		if !_accept(parser, _forethought_coP_prepositional_phraseAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_2
		if !_accept(parser, _prepositional_phrase_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _prepositional_phrase_1, start, pos, perr)
fail:
	return _memoize(parser, _prepositional_phrase_1, start, -1, perr)
}

func _prepositional_phrase_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_prepositional_phrase_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "prepositional_phrase_1"}
	// forethought_coP_prepositional_phrase/prepositional_phrase_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_prepositional_phrase
		if !_node(parser, _forethought_coP_prepositional_phraseNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// prepositional_phrase_2
		if !_node(parser, _prepositional_phrase_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prepositional_phrase_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _prepositional_phrase_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prepositional_phrase_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prepositional_phrase_1}
	// forethought_coP_prepositional_phrase/prepositional_phrase_2
	{
		pos3 := pos
		// forethought_coP_prepositional_phrase
		if !_fail(parser, _forethought_coP_prepositional_phraseFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_2
		if !_fail(parser, _prepositional_phrase_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prepositional_phrase_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_prepositional_phrase_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_prepositional_phrase/prepositional_phrase_2
	{
		pos3 := pos
		var node2 string
		// forethought_coP_prepositional_phrase
		if p, n := _forethought_coP_prepositional_phraseAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// prepositional_phrase_2
		if p, n := _prepositional_phrase_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prepositional_phrase_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _prepositional_phrase_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// preposition spaces? argument
	// preposition
	if !_accept(parser, _prepositionAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _prepositional_phrase_2, start, pos, perr)
fail:
	return _memoize(parser, _prepositional_phrase_2, start, -1, perr)
}

func _prepositional_phrase_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_prepositional_phrase_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "prepositional_phrase_2"}
	// preposition spaces? argument
	// preposition
	if !_node(parser, _prepositionNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prepositional_phrase_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _prepositional_phrase_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prepositional_phrase_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prepositional_phrase_2}
	// preposition spaces? argument
	// preposition
	if !_fail(parser, _prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prepositional_phrase_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_prepositional_phrase_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// preposition spaces? argument
	{
		var node0 string
		// preposition
		if p, n := _prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_prepositional_phraseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_prepositional_phrase, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// prepositional_phrase_1 spaces? co_bar_prepositional_phrase
	// prepositional_phrase_1
	if !_accept(parser, _prepositional_phrase_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_prepositional_phrase
	if !_accept(parser, _co_bar_prepositional_phraseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_prepositional_phrase, start, pos, perr)
fail:
	return _memoize(parser, _coP_prepositional_phrase, start, -1, perr)
}

func _coP_prepositional_phraseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_prepositional_phrase}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_prepositional_phrase"}
	// prepositional_phrase_1 spaces? co_bar_prepositional_phrase
	// prepositional_phrase_1
	if !_node(parser, _prepositional_phrase_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_prepositional_phrase
	if !_node(parser, _co_bar_prepositional_phraseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_prepositional_phraseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_prepositional_phrase, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_prepositional_phrase",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_prepositional_phrase}
	// prepositional_phrase_1 spaces? co_bar_prepositional_phrase
	// prepositional_phrase_1
	if !_fail(parser, _prepositional_phrase_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_prepositional_phrase
	if !_fail(parser, _co_bar_prepositional_phraseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_prepositional_phraseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_prepositional_phrase}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// prepositional_phrase_1 spaces? co_bar_prepositional_phrase
	{
		var node0 string
		// prepositional_phrase_1
		if p, n := _prepositional_phrase_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_prepositional_phrase
		if p, n := _co_bar_prepositional_phraseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_prepositional_phraseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_prepositional_phrase, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? prepositional_phrase_1
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// prepositional_phrase_1
	if !_accept(parser, _prepositional_phrase_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_prepositional_phrase, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_prepositional_phrase, start, -1, perr)
}

func _co_bar_prepositional_phraseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_prepositional_phrase}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_prepositional_phrase"}
	// connective spaces? prepositional_phrase_1
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// prepositional_phrase_1
	if !_node(parser, _prepositional_phrase_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_prepositional_phraseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_prepositional_phrase, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_prepositional_phrase",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_prepositional_phrase}
	// connective spaces? prepositional_phrase_1
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// prepositional_phrase_1
	if !_fail(parser, _prepositional_phrase_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_prepositional_phraseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_prepositional_phrase}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? prepositional_phrase_1
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// prepositional_phrase_1
		if p, n := _prepositional_phrase_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_prepositional_phraseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_prepositional_phrase, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_prepositional_phrase_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_prepositional_phrase_1
	if !_accept(parser, _forethought_coP_prepositional_phrase_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_prepositional_phrase, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_prepositional_phrase, start, -1, perr)
}

func _forethought_coP_prepositional_phraseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_prepositional_phrase}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_prepositional_phrase"}
	// forethought_connective spaces? forethought_coP_prepositional_phrase_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_prepositional_phrase_1
	if !_node(parser, _forethought_coP_prepositional_phrase_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_prepositional_phraseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_prepositional_phrase, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_prepositional_phrase",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_prepositional_phrase}
	// forethought_connective spaces? forethought_coP_prepositional_phrase_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_prepositional_phrase_1
	if !_fail(parser, _forethought_coP_prepositional_phrase_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_prepositional_phraseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_prepositional_phrase}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_prepositional_phrase_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_prepositional_phrase_1
		if p, n := _forethought_coP_prepositional_phrase_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_prepositional_phrase_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_prepositional_phrase_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// prepositional_phrase spaces? forethought_co_bar_prepositional_phrase
	// prepositional_phrase
	if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_prepositional_phrase
	if !_accept(parser, _forethought_co_bar_prepositional_phraseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_prepositional_phrase_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_prepositional_phrase_1, start, -1, perr)
}

func _forethought_coP_prepositional_phrase_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_prepositional_phrase_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_prepositional_phrase_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_prepositional_phrase_1"}
	// prepositional_phrase spaces? forethought_co_bar_prepositional_phrase
	// prepositional_phrase
	if !_node(parser, _prepositional_phraseNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_prepositional_phrase
	if !_node(parser, _forethought_co_bar_prepositional_phraseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_prepositional_phrase_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_prepositional_phrase_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_prepositional_phrase_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_prepositional_phrase_1}
	// prepositional_phrase spaces? forethought_co_bar_prepositional_phrase
	// prepositional_phrase
	if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_prepositional_phrase
	if !_fail(parser, _forethought_co_bar_prepositional_phraseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_prepositional_phrase_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_prepositional_phrase_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_prepositional_phrase_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// prepositional_phrase spaces? forethought_co_bar_prepositional_phrase
	{
		var node0 string
		// prepositional_phrase
		if p, n := _prepositional_phraseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_prepositional_phrase
		if p, n := _forethought_co_bar_prepositional_phraseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_prepositional_phraseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_prepositional_phrase, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? prepositional_phrase
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// prepositional_phrase
	if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_prepositional_phrase, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_prepositional_phrase, start, -1, perr)
}

func _forethought_co_bar_prepositional_phraseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_prepositional_phrase}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_prepositional_phrase"}
	// forethought_marker spaces? prepositional_phrase
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// prepositional_phrase
	if !_node(parser, _prepositional_phraseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_prepositional_phraseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_prepositional_phrase, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_prepositional_phrase",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_prepositional_phrase}
	// forethought_marker spaces? prepositional_phrase
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// prepositional_phrase
	if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_prepositional_phraseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_prepositional_phrase}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? prepositional_phrase
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// prepositional_phrase
		if p, n := _prepositional_phraseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_preposition/preposition_1
	{
		pos3 := pos
		// coP_preposition
		if !_accept(parser, _coP_prepositionAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_1
		if !_accept(parser, _preposition_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _preposition, start, pos, perr)
fail:
	return _memoize(parser, _preposition, start, -1, perr)
}

func _prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition"}
	// coP_preposition/preposition_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_preposition
		if !_node(parser, _coP_prepositionNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_1
		if !_node(parser, _preposition_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition}
	// coP_preposition/preposition_1
	{
		pos3 := pos
		// coP_preposition
		if !_fail(parser, _coP_prepositionFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_1
		if !_fail(parser, _preposition_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_preposition/preposition_1
	{
		pos3 := pos
		var node2 string
		// coP_preposition
		if p, n := _coP_prepositionAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// preposition_1
		if p, n := _preposition_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _preposition_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_coP_preposition/preposition_2
	{
		pos3 := pos
		// forethought_coP_preposition
		if !_accept(parser, _forethought_coP_prepositionAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_2
		if !_accept(parser, _preposition_2Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _preposition_1, start, pos, perr)
fail:
	return _memoize(parser, _preposition_1, start, -1, perr)
}

func _preposition_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition_1"}
	// forethought_coP_preposition/preposition_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_coP_preposition
		if !_node(parser, _forethought_coP_prepositionNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_2
		if !_node(parser, _preposition_2Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _preposition_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_1}
	// forethought_coP_preposition/preposition_2
	{
		pos3 := pos
		// forethought_coP_preposition
		if !_fail(parser, _forethought_coP_prepositionFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_2
		if !_fail(parser, _preposition_2Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_coP_preposition/preposition_2
	{
		pos3 := pos
		var node2 string
		// forethought_coP_preposition
		if p, n := _forethought_coP_prepositionAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// preposition_2
		if p, n := _preposition_2Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _preposition_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial_preposition/preposition_3
	{
		pos3 := pos
		// serial_preposition
		if !_accept(parser, _serial_prepositionAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_3
		if !_accept(parser, _preposition_3Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _preposition_2, start, pos, perr)
fail:
	return _memoize(parser, _preposition_2, start, -1, perr)
}

func _preposition_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition_2"}
	// serial_preposition/preposition_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial_preposition
		if !_node(parser, _serial_prepositionNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_3
		if !_node(parser, _preposition_3Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _preposition_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_2}
	// serial_preposition/preposition_3
	{
		pos3 := pos
		// serial_preposition
		if !_fail(parser, _serial_prepositionFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_3
		if !_fail(parser, _preposition_3Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// serial_preposition/preposition_3
	{
		pos3 := pos
		var node2 string
		// serial_preposition
		if p, n := _serial_prepositionAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// preposition_3
		if p, n := _preposition_3Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _preposition_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_pred_preposition/forethought_coP_pred_preposition/LU_preposition/MI_preposition/PO_preposition/quotation_preposition/preposition_4
	{
		pos3 := pos
		// coP_pred_preposition
		if !_accept(parser, _coP_pred_prepositionAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_preposition
		if !_accept(parser, _forethought_coP_pred_prepositionAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_preposition
		if !_accept(parser, _LU_prepositionAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_preposition
		if !_accept(parser, _MI_prepositionAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_preposition
		if !_accept(parser, _PO_prepositionAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// quotation_preposition
		if !_accept(parser, _quotation_prepositionAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// preposition_4
		if !_accept(parser, _preposition_4Accepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _preposition_3, start, pos, perr)
fail:
	return _memoize(parser, _preposition_3, start, -1, perr)
}

func _preposition_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition_3"}
	// coP_pred_preposition/forethought_coP_pred_preposition/LU_preposition/MI_preposition/PO_preposition/quotation_preposition/preposition_4
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_pred_preposition
		if !_node(parser, _coP_pred_prepositionNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_coP_pred_preposition
		if !_node(parser, _forethought_coP_pred_prepositionNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_preposition
		if !_node(parser, _LU_prepositionNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_preposition
		if !_node(parser, _MI_prepositionNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_preposition
		if !_node(parser, _PO_prepositionNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// quotation_preposition
		if !_node(parser, _quotation_prepositionNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_4
		if !_node(parser, _preposition_4Node, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _preposition_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_3}
	// coP_pred_preposition/forethought_coP_pred_preposition/LU_preposition/MI_preposition/PO_preposition/quotation_preposition/preposition_4
	{
		pos3 := pos
		// coP_pred_preposition
		if !_fail(parser, _coP_pred_prepositionFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_preposition
		if !_fail(parser, _forethought_coP_pred_prepositionFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_preposition
		if !_fail(parser, _LU_prepositionFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_preposition
		if !_fail(parser, _MI_prepositionFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_preposition
		if !_fail(parser, _PO_prepositionFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// quotation_preposition
		if !_fail(parser, _quotation_prepositionFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// preposition_4
		if !_fail(parser, _preposition_4Fail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_pred_preposition/forethought_coP_pred_preposition/LU_preposition/MI_preposition/PO_preposition/quotation_preposition/preposition_4
	{
		pos3 := pos
		var node2 string
		// coP_pred_preposition
		if p, n := _coP_pred_prepositionAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_coP_pred_preposition
		if p, n := _forethought_coP_pred_prepositionAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// LU_preposition
		if p, n := _LU_prepositionAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// MI_preposition
		if p, n := _MI_prepositionAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// PO_preposition
		if p, n := _PO_prepositionAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// quotation_preposition
		if p, n := _quotation_prepositionAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// preposition_4
		if p, n := _preposition_4Action(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _preposition_4Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_4, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// preposition_syllable compound_syllable+/!function_word preposition_syllable
	{
		pos3 := pos
		// preposition_syllable compound_syllable+
		// preposition_syllable
		if !_accept(parser, _preposition_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word preposition_syllable
		// !function_word
		{
			pos13 := pos
			perr15 := perr
			// function_word
			if !_accept(parser, _function_wordAccepts, &pos, &perr) {
				goto ok12
			}
			pos = pos13
			perr = _max(perr15, pos)
			goto fail10
		ok12:
			pos = pos13
			perr = perr15
		}
		// preposition_syllable
		if !_accept(parser, _preposition_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _preposition_4, start, pos, perr)
fail:
	return _memoize(parser, _preposition_4, start, -1, perr)
}

func _preposition_4Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_4}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition_4"}
	// preposition_syllable compound_syllable+/!function_word preposition_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// preposition_syllable compound_syllable+
		// preposition_syllable
		if !_node(parser, _preposition_syllableNode, node, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_node(parser, _compound_syllableNode, node, &pos) {
			goto fail4
		}
		for {
			nkids6 := len(node.Kids)
			pos7 := pos
			// compound_syllable
			if !_node(parser, _compound_syllableNode, node, &pos) {
				goto fail9
			}
			continue
		fail9:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
			break
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// !function_word preposition_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(node.Kids)
			// function_word
			if !_node(parser, _function_wordNode, node, &pos) {
				goto ok12
			}
			pos = pos13
			node.Kids = node.Kids[:nkids14]
			goto fail10
		ok12:
			pos = pos13
			node.Kids = node.Kids[:nkids14]
		}
		// preposition_syllable
		if !_node(parser, _preposition_syllableNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _preposition_4Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition_4, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_4",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_4}
	// preposition_syllable compound_syllable+/!function_word preposition_syllable
	{
		pos3 := pos
		// preposition_syllable compound_syllable+
		// preposition_syllable
		if !_fail(parser, _preposition_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word preposition_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(failure.Kids)
			// function_word
			if !_fail(parser, _function_wordFail, errPos, failure, &pos) {
				goto ok12
			}
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "!function_word",
				})
			}
			goto fail10
		ok12:
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
		}
		// preposition_syllable
		if !_fail(parser, _preposition_syllableFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_4Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_4}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// preposition_syllable compound_syllable+/!function_word preposition_syllable
	{
		pos3 := pos
		var node2 string
		// preposition_syllable compound_syllable+
		{
			var node5 string
			// preposition_syllable
			if p, n := _preposition_syllableAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
			// compound_syllable+
			{
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail4
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
			}
			for {
				pos7 := pos
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail9
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
				continue
			fail9:
				pos = pos7
				break
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// !function_word preposition_syllable
		{
			var node11 string
			// !function_word
			{
				pos13 := pos
				// function_word
				if p, n := _function_wordAction(parser, pos); n == nil {
					goto ok12
				} else {
					pos = p
				}
				pos = pos13
				goto fail10
			ok12:
				pos = pos13
				node11 = ""
			}
			node, node11 = node+node11, ""
			// preposition_syllable
			if p, n := _preposition_syllableAction(parser, pos); n == nil {
				goto fail10
			} else {
				node11 = *n
				pos = p
			}
			node, node11 = node+node11, ""
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// preposition_1 spaces? co_bar_preposition
	// preposition_1
	if !_accept(parser, _preposition_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_preposition
	if !_accept(parser, _co_bar_prepositionAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_preposition, start, pos, perr)
fail:
	return _memoize(parser, _coP_preposition, start, -1, perr)
}

func _coP_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_preposition"}
	// preposition_1 spaces? co_bar_preposition
	// preposition_1
	if !_node(parser, _preposition_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_preposition
	if !_node(parser, _co_bar_prepositionNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_preposition}
	// preposition_1 spaces? co_bar_preposition
	// preposition_1
	if !_fail(parser, _preposition_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_preposition
	if !_fail(parser, _co_bar_prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// preposition_1 spaces? co_bar_preposition
	{
		var node0 string
		// preposition_1
		if p, n := _preposition_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_preposition
		if p, n := _co_bar_prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _co_bar_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _co_bar_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// connective spaces? preposition
	// connective
	if !_accept(parser, _connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// preposition
	if !_accept(parser, _prepositionAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _co_bar_preposition, start, pos, perr)
fail:
	return _memoize(parser, _co_bar_preposition, start, -1, perr)
}

func _co_bar_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_co_bar_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "co_bar_preposition"}
	// connective spaces? preposition
	// connective
	if !_node(parser, _connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// preposition
	if !_node(parser, _prepositionNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _co_bar_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _co_bar_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "co_bar_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _co_bar_preposition}
	// connective spaces? preposition
	// connective
	if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// preposition
	if !_fail(parser, _prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _co_bar_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_co_bar_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _co_bar_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// connective spaces? preposition
	{
		var node0 string
		// connective
		if p, n := _connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// preposition
		if p, n := _prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_preposition_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_preposition_1
	if !_accept(parser, _forethought_coP_preposition_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_preposition, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_preposition, start, -1, perr)
}

func _forethought_coP_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_preposition"}
	// forethought_connective spaces? forethought_coP_preposition_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_preposition_1
	if !_node(parser, _forethought_coP_preposition_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_preposition}
	// forethought_connective spaces? forethought_coP_preposition_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_preposition_1
	if !_fail(parser, _forethought_coP_preposition_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_preposition_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_preposition_1
		if p, n := _forethought_coP_preposition_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_preposition_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_preposition_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// preposition spaces? forethought_co_bar_preposition
	// preposition
	if !_accept(parser, _prepositionAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_preposition
	if !_accept(parser, _forethought_co_bar_prepositionAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_preposition_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_preposition_1, start, -1, perr)
}

func _forethought_coP_preposition_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_preposition_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_preposition_1"}
	// preposition spaces? forethought_co_bar_preposition
	// preposition
	if !_node(parser, _prepositionNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_preposition
	if !_node(parser, _forethought_co_bar_prepositionNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_preposition_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_preposition_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_preposition_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_preposition_1}
	// preposition spaces? forethought_co_bar_preposition
	// preposition
	if !_fail(parser, _prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_preposition
	if !_fail(parser, _forethought_co_bar_prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_preposition_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_preposition_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// preposition spaces? forethought_co_bar_preposition
	{
		var node0 string
		// preposition
		if p, n := _prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_preposition
		if p, n := _forethought_co_bar_prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_co_bar_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_co_bar_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_marker spaces? preposition
	// forethought_marker
	if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// preposition
	if !_accept(parser, _prepositionAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_co_bar_preposition, start, pos, perr)
fail:
	return _memoize(parser, _forethought_co_bar_preposition, start, -1, perr)
}

func _forethought_co_bar_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_co_bar_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_co_bar_preposition"}
	// forethought_marker spaces? preposition
	// forethought_marker
	if !_node(parser, _forethought_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// preposition
	if !_node(parser, _prepositionNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_co_bar_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_co_bar_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_co_bar_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_co_bar_preposition}
	// forethought_marker spaces? preposition
	// forethought_marker
	if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// preposition
	if !_fail(parser, _prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_co_bar_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_co_bar_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_co_bar_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_marker spaces? preposition
	{
		var node0 string
		// forethought_marker
		if p, n := _forethought_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// preposition
		if p, n := _prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _serial_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _serial_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// preposition_3 spaces? serial_predicate
	// preposition_3
	if !_accept(parser, _preposition_3Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_accept(parser, _serial_predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _serial_preposition, start, pos, perr)
fail:
	return _memoize(parser, _serial_preposition, start, -1, perr)
}

func _serial_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_serial_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "serial_preposition"}
	// preposition_3 spaces? serial_predicate
	// preposition_3
	if !_node(parser, _preposition_3Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_node(parser, _serial_predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _serial_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _serial_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "serial_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _serial_preposition}
	// preposition_3 spaces? serial_predicate
	// preposition_3
	if !_fail(parser, _preposition_3Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_fail(parser, _serial_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _serial_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_serial_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// preposition_3 spaces? serial_predicate
	{
		var node0 string
		// preposition_3
		if p, n := _preposition_3Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// serial_predicate
		if p, n := _serial_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_pred_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_pred_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// preposition_4 spaces? co_bar_pred
	// preposition_4
	if !_accept(parser, _preposition_4Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_accept(parser, _co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_pred_preposition, start, pos, perr)
fail:
	return _memoize(parser, _coP_pred_preposition, start, -1, perr)
}

func _coP_pred_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_pred_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_pred_preposition"}
	// preposition_4 spaces? co_bar_pred
	// preposition_4
	if !_node(parser, _preposition_4Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_node(parser, _co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_pred_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_pred_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_pred_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_pred_preposition}
	// preposition_4 spaces? co_bar_pred
	// preposition_4
	if !_fail(parser, _preposition_4Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_fail(parser, _co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_pred_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_pred_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// preposition_4 spaces? co_bar_pred
	{
		var node0 string
		// preposition_4
		if p, n := _preposition_4Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_pred
		if p, n := _co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_pred_preposition_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_preposition_1
	if !_accept(parser, _forethought_coP_pred_preposition_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_preposition, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_preposition, start, -1, perr)
}

func _forethought_coP_pred_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_preposition"}
	// forethought_connective spaces? forethought_coP_pred_preposition_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_preposition_1
	if !_node(parser, _forethought_coP_pred_preposition_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_preposition}
	// forethought_connective spaces? forethought_coP_pred_preposition_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_preposition_1
	if !_fail(parser, _forethought_coP_pred_preposition_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_pred_preposition_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_pred_preposition_1
		if p, n := _forethought_coP_pred_preposition_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_preposition_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_preposition_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// preposition spaces? forethought_co_bar_pred
	// preposition
	if !_accept(parser, _prepositionAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_accept(parser, _forethought_co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_preposition_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_preposition_1, start, -1, perr)
}

func _forethought_coP_pred_preposition_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_preposition_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_preposition_1"}
	// preposition spaces? forethought_co_bar_pred
	// preposition
	if !_node(parser, _prepositionNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_node(parser, _forethought_co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_preposition_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_preposition_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_preposition_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_preposition_1}
	// preposition spaces? forethought_co_bar_pred
	// preposition
	if !_fail(parser, _prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_fail(parser, _forethought_co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_preposition_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_preposition_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// preposition spaces? forethought_co_bar_pred
	{
		var node0 string
		// preposition
		if p, n := _prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_pred
		if p, n := _forethought_co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// LU_preposition_tone spaces? statement
	// LU_preposition_tone
	if !_accept(parser, _LU_preposition_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_preposition, start, pos, perr)
fail:
	return _memoize(parser, _LU_preposition, start, -1, perr)
}

func _LU_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_preposition"}
	// LU_preposition_tone spaces? statement
	// LU_preposition_tone
	if !_node(parser, _LU_preposition_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_preposition}
	// LU_preposition_tone spaces? statement
	// LU_preposition_tone
	if !_fail(parser, _LU_preposition_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// LU_preposition_tone spaces? statement
	{
		var node0 string
		// LU_preposition_tone
		if p, n := _LU_preposition_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_preposition_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_preposition_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &LU preposition_syllable
	// &LU
	{
		pos2 := pos
		perr4 := perr
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// preposition_syllable
	if !_accept(parser, _preposition_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_preposition_tone, start, pos, perr)
fail:
	return _memoize(parser, _LU_preposition_tone, start, -1, perr)
}

func _LU_preposition_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_preposition_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_preposition_tone"}
	// &LU preposition_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// LU
		if !_node(parser, _LUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// preposition_syllable
	if !_node(parser, _preposition_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_preposition_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_preposition_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_preposition_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_preposition_tone}
	// &LU preposition_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&LU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// preposition_syllable
	if !_fail(parser, _preposition_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_preposition_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_preposition_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &LU preposition_syllable
	{
		var node0 string
		// &LU
		{
			pos2 := pos
			// LU
			if p, n := _LUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// preposition_syllable
		if p, n := _preposition_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_preposition_1 spaces? end_predicatizer?
	// MI_preposition_1
	if !_accept(parser, _MI_preposition_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _MI_preposition, start, pos, perr)
fail:
	return _memoize(parser, _MI_preposition, start, -1, perr)
}

func _MI_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_preposition"}
	// MI_preposition_1 spaces? end_predicatizer?
	// MI_preposition_1
	if !_node(parser, _MI_preposition_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_preposition}
	// MI_preposition_1 spaces? end_predicatizer?
	// MI_preposition_1
	if !_fail(parser, _MI_preposition_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_preposition_1 spaces? end_predicatizer?
	{
		var node0 string
		// MI_preposition_1
		if p, n := _MI_preposition_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_preposition_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_preposition_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_preposition_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_preposition_tone
	if !_accept(parser, _MI_preposition_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_accept(parser, _predicateAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_accept(parser, _adverbAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	return _memoize(parser, _MI_preposition_1, start, pos, perr)
fail:
	return _memoize(parser, _MI_preposition_1, start, -1, perr)
}

func _MI_preposition_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_preposition_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_preposition_1"}
	// MI_preposition_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_preposition_tone
	if !_node(parser, _MI_preposition_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// predicate/argument/adverb/prepositional_phrase
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// predicate
			if !_node(parser, _predicateNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail14
			}
			goto ok7
		fail14:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			goto fail
		ok7:
		}
		sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
		node.Kids = append(node.Kids[:nkids5], sub)
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_preposition_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_preposition_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_preposition_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_preposition_1}
	// MI_preposition_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_preposition_tone
	if !_fail(parser, _MI_preposition_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_fail(parser, _predicateFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_fail(parser, _adverbFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_preposition_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_preposition_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_preposition_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	{
		var node0 string
		// MI_preposition_tone
		if p, n := _MI_preposition_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (predicate/argument/adverb/prepositional_phrase)
		// predicate/argument/adverb/prepositional_phrase
		{
			pos8 := pos
			var node7 string
			// predicate
			if p, n := _predicateAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// adverb
			if p, n := _adverbAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// prepositional_phrase
			if p, n := _prepositional_phraseAction(parser, pos); n == nil {
				goto fail12
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail12:
			node0 = node7
			pos = pos8
			goto fail
		ok5:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_preposition_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_preposition_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MI preposition_syllable
	// &MI
	{
		pos2 := pos
		perr4 := perr
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// preposition_syllable
	if !_accept(parser, _preposition_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MI_preposition_tone, start, pos, perr)
fail:
	return _memoize(parser, _MI_preposition_tone, start, -1, perr)
}

func _MI_preposition_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_preposition_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_preposition_tone"}
	// &MI preposition_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MI
		if !_node(parser, _MINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// preposition_syllable
	if !_node(parser, _preposition_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_preposition_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_preposition_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_preposition_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_preposition_tone}
	// &MI preposition_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// preposition_syllable
	if !_fail(parser, _preposition_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_preposition_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_preposition_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MI preposition_syllable
	{
		var node0 string
		// &MI
		{
			pos2 := pos
			// MI
			if p, n := _MIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// preposition_syllable
		if p, n := _preposition_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_preposition_1 spaces? end_predicatizer?
	// PO_preposition_1
	if !_accept(parser, _PO_preposition_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _PO_preposition, start, pos, perr)
fail:
	return _memoize(parser, _PO_preposition, start, -1, perr)
}

func _PO_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_preposition"}
	// PO_preposition_1 spaces? end_predicatizer?
	// PO_preposition_1
	if !_node(parser, _PO_preposition_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_preposition}
	// PO_preposition_1 spaces? end_predicatizer?
	// PO_preposition_1
	if !_fail(parser, _PO_preposition_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_preposition_1 spaces? end_predicatizer?
	{
		var node0 string
		// PO_preposition_1
		if p, n := _PO_preposition_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_preposition_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_preposition_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_preposition_tone spaces? argument
	// PO_preposition_tone
	if !_accept(parser, _PO_preposition_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_preposition_1, start, pos, perr)
fail:
	return _memoize(parser, _PO_preposition_1, start, -1, perr)
}

func _PO_preposition_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_preposition_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_preposition_1"}
	// PO_preposition_tone spaces? argument
	// PO_preposition_tone
	if !_node(parser, _PO_preposition_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_preposition_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_preposition_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_preposition_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_preposition_1}
	// PO_preposition_tone spaces? argument
	// PO_preposition_tone
	if !_fail(parser, _PO_preposition_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_preposition_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_preposition_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_preposition_tone spaces? argument
	{
		var node0 string
		// PO_preposition_tone
		if p, n := _PO_preposition_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_preposition_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_preposition_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &PO preposition_syllable
	// &PO
	{
		pos2 := pos
		perr4 := perr
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// preposition_syllable
	if !_accept(parser, _preposition_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_preposition_tone, start, pos, perr)
fail:
	return _memoize(parser, _PO_preposition_tone, start, -1, perr)
}

func _PO_preposition_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_preposition_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_preposition_tone"}
	// &PO preposition_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// preposition_syllable
	if !_node(parser, _preposition_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_preposition_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_preposition_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_preposition_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_preposition_tone}
	// &PO preposition_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&PO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// preposition_syllable
	if !_fail(parser, _preposition_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_preposition_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_preposition_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &PO preposition_syllable
	{
		var node0 string
		// &PO
		{
			pos2 := pos
			// PO
			if p, n := _POAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// preposition_syllable
		if p, n := _preposition_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quotation_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _quotation_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_preposition spaces? end_quote
	// MO_preposition
	if !_accept(parser, _MO_prepositionAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_accept(parser, _end_quoteAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _quotation_preposition, start, pos, perr)
fail:
	return _memoize(parser, _quotation_preposition, start, -1, perr)
}

func _quotation_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_quotation_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quotation_preposition"}
	// MO_preposition spaces? end_quote
	// MO_preposition
	if !_node(parser, _MO_prepositionNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_quote
	if !_node(parser, _end_quoteNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quotation_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _quotation_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quotation_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quotation_preposition}
	// MO_preposition spaces? end_quote
	// MO_preposition
	if !_fail(parser, _MO_prepositionFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_fail(parser, _end_quoteFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quotation_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_quotation_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_preposition spaces? end_quote
	{
		var node0 string
		// MO_preposition
		if p, n := _MO_prepositionAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_quote
		if p, n := _end_quoteAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_preposition_tone spaces? discourse
	// MO_preposition_tone
	if !_accept(parser, _MO_preposition_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_accept(parser, _discourseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_preposition, start, pos, perr)
fail:
	return _memoize(parser, _MO_preposition, start, -1, perr)
}

func _MO_prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_preposition}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_preposition"}
	// MO_preposition_tone spaces? discourse
	// MO_preposition_tone
	if !_node(parser, _MO_preposition_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// discourse
	if !_node(parser, _discourseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_preposition}
	// MO_preposition_tone spaces? discourse
	// MO_preposition_tone
	if !_fail(parser, _MO_preposition_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_fail(parser, _discourseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_prepositionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_preposition_tone spaces? discourse
	{
		var node0 string
		// MO_preposition_tone
		if p, n := _MO_preposition_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// discourse
		if p, n := _discourseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_preposition_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_preposition_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MO preposition_syllable
	// &MO
	{
		pos2 := pos
		perr4 := perr
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// preposition_syllable
	if !_accept(parser, _preposition_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_preposition_tone, start, pos, perr)
fail:
	return _memoize(parser, _MO_preposition_tone, start, -1, perr)
}

func _MO_preposition_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_preposition_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_preposition_tone"}
	// &MO preposition_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MO
		if !_node(parser, _MONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// preposition_syllable
	if !_node(parser, _preposition_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_preposition_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_preposition_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_preposition_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_preposition_tone}
	// &MO preposition_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// preposition_syllable
	if !_fail(parser, _preposition_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_preposition_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_preposition_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MO preposition_syllable
	{
		var node0 string
		// &MO
		{
			pos2 := pos
			// MO
			if p, n := _MOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// preposition_syllable
		if p, n := _preposition_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_clauseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_clause, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_content_statement/content_clause_1
	{
		pos3 := pos
		// coP_content_statement
		if !_accept(parser, _coP_content_statementAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// content_clause_1
		if !_accept(parser, _content_clause_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_clause, start, pos, perr)
fail:
	return _memoize(parser, _content_clause, start, -1, perr)
}

func _content_clauseNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_clause]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_clause}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_clause"}
	// coP_content_statement/content_clause_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_content_statement
		if !_node(parser, _coP_content_statementNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// content_clause_1
		if !_node(parser, _content_clause_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_clauseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_clause, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_clause",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_clause}
	// coP_content_statement/content_clause_1
	{
		pos3 := pos
		// coP_content_statement
		if !_fail(parser, _coP_content_statementFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// content_clause_1
		if !_fail(parser, _content_clause_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_clauseAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_clause]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_clause}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_content_statement/content_clause_1
	{
		pos3 := pos
		var node2 string
		// coP_content_statement
		if p, n := _coP_content_statementAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// content_clause_1
		if p, n := _content_clause_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_clause_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_clause_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_predication spaces? end_statement?/LU_content
	{
		pos3 := pos
		// content_predication spaces? end_statement?
		// content_predication
		if !_accept(parser, _content_predicationAccepts, &pos, &perr) {
			goto fail4
		}
		// spaces?
		{
			pos7 := pos
			// spaces
			if !_accept(parser, _spacesAccepts, &pos, &perr) {
				goto fail8
			}
			goto ok9
		fail8:
			pos = pos7
		ok9:
		}
		// end_statement?
		{
			pos11 := pos
			// end_statement
			if !_accept(parser, _end_statementAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok13
		fail12:
			pos = pos11
		ok13:
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_content
		if !_accept(parser, _LU_contentAccepts, &pos, &perr) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_clause_1, start, pos, perr)
fail:
	return _memoize(parser, _content_clause_1, start, -1, perr)
}

func _content_clause_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_clause_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_clause_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_clause_1"}
	// content_predication spaces? end_statement?/LU_content
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// content_predication spaces? end_statement?
		// content_predication
		if !_node(parser, _content_predicationNode, node, &pos) {
			goto fail4
		}
		// spaces?
		{
			nkids6 := len(node.Kids)
			pos7 := pos
			// spaces
			if !_node(parser, _spacesNode, node, &pos) {
				goto fail8
			}
			goto ok9
		fail8:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
		ok9:
		}
		// end_statement?
		{
			nkids10 := len(node.Kids)
			pos11 := pos
			// end_statement
			if !_node(parser, _end_statementNode, node, &pos) {
				goto fail12
			}
			goto ok13
		fail12:
			node.Kids = node.Kids[:nkids10]
			pos = pos11
		ok13:
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_content
		if !_node(parser, _LU_contentNode, node, &pos) {
			goto fail14
		}
		goto ok0
	fail14:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_clause_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_clause_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_clause_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_clause_1}
	// content_predication spaces? end_statement?/LU_content
	{
		pos3 := pos
		// content_predication spaces? end_statement?
		// content_predication
		if !_fail(parser, _content_predicationFail, errPos, failure, &pos) {
			goto fail4
		}
		// spaces?
		{
			pos7 := pos
			// spaces
			if !_fail(parser, _spacesFail, errPos, failure, &pos) {
				goto fail8
			}
			goto ok9
		fail8:
			pos = pos7
		ok9:
		}
		// end_statement?
		{
			pos11 := pos
			// end_statement
			if !_fail(parser, _end_statementFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok13
		fail12:
			pos = pos11
		ok13:
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_content
		if !_fail(parser, _LU_contentFail, errPos, failure, &pos) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_clause_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_clause_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_clause_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_predication spaces? end_statement?/LU_content
	{
		pos3 := pos
		var node2 string
		// content_predication spaces? end_statement?
		{
			var node5 string
			// content_predication
			if p, n := _content_predicationAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
			// spaces?
			{
				pos7 := pos
				// spaces
				if p, n := _spacesAction(parser, pos); n == nil {
					goto fail8
				} else {
					node5 = *n
					pos = p
				}
				goto ok9
			fail8:
				node5 = ""
				pos = pos7
			ok9:
			}
			node, node5 = node+node5, ""
			// end_statement?
			{
				pos11 := pos
				// end_statement
				if p, n := _end_statementAction(parser, pos); n == nil {
					goto fail12
				} else {
					node5 = *n
					pos = p
				}
				goto ok13
			fail12:
				node5 = ""
				pos = pos11
			ok13:
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// LU_content
		if p, n := _LU_contentAction(parser, pos); n == nil {
			goto fail14
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail14:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_predicationAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_predication, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_predicate spaces? terms?
	// content_predicate
	if !_accept(parser, _content_predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms?
	{
		pos6 := pos
		// terms
		if !_accept(parser, _termsAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _content_predication, start, pos, perr)
fail:
	return _memoize(parser, _content_predication, start, -1, perr)
}

func _content_predicationNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predication}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_predication"}
	// content_predicate spaces? terms?
	// content_predicate
	if !_node(parser, _content_predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// terms?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// terms
		if !_node(parser, _termsNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_predicationFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_predication, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_predication",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_predication}
	// content_predicate spaces? terms?
	// content_predicate
	if !_fail(parser, _content_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// terms?
	{
		pos6 := pos
		// terms
		if !_fail(parser, _termsFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_predicationAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predication}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_predicate spaces? terms?
	{
		var node0 string
		// content_predicate
		if p, n := _content_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// terms?
		{
			pos6 := pos
			// terms
			if p, n := _termsAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_content_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_content_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_clause_1 spaces? co_bar_statement
	// content_clause_1
	if !_accept(parser, _content_clause_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_accept(parser, _co_bar_statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_content_statement, start, pos, perr)
fail:
	return _memoize(parser, _coP_content_statement, start, -1, perr)
}

func _coP_content_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_content_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_content_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_content_statement"}
	// content_clause_1 spaces? co_bar_statement
	// content_clause_1
	if !_node(parser, _content_clause_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_node(parser, _co_bar_statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_content_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_content_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_content_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_content_statement}
	// content_clause_1 spaces? co_bar_statement
	// content_clause_1
	if !_fail(parser, _content_clause_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_statement
	if !_fail(parser, _co_bar_statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_content_statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_content_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_content_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_clause_1 spaces? co_bar_statement
	{
		var node0 string
		// content_clause_1
		if p, n := _content_clause_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_statement
		if p, n := _co_bar_statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial_content_predicate/content_predicate_1
	{
		pos3 := pos
		// serial_content_predicate
		if !_accept(parser, _serial_content_predicateAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// content_predicate_1
		if !_accept(parser, _content_predicate_1Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _content_predicate, start, -1, perr)
}

func _content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_predicate"}
	// serial_content_predicate/content_predicate_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial_content_predicate
		if !_node(parser, _serial_content_predicateNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// content_predicate_1
		if !_node(parser, _content_predicate_1Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_predicate}
	// serial_content_predicate/content_predicate_1
	{
		pos3 := pos
		// serial_content_predicate
		if !_fail(parser, _serial_content_predicateFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// content_predicate_1
		if !_fail(parser, _content_predicate_1Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// serial_content_predicate/content_predicate_1
	{
		pos3 := pos
		var node2 string
		// serial_content_predicate
		if p, n := _serial_content_predicateAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// content_predicate_1
		if p, n := _content_predicate_1Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// coP_pred_content_predicate/forethought_coP_pred_content_predicate/content_predicate_2
	{
		pos3 := pos
		// coP_pred_content_predicate
		if !_accept(parser, _coP_pred_content_predicateAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_content_predicate
		if !_accept(parser, _forethought_coP_pred_content_predicateAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// content_predicate_2
		if !_accept(parser, _content_predicate_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _content_predicate_1, start, -1, perr)
}

func _content_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_predicate_1"}
	// coP_pred_content_predicate/forethought_coP_pred_content_predicate/content_predicate_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// coP_pred_content_predicate
		if !_node(parser, _coP_pred_content_predicateNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_coP_pred_content_predicate
		if !_node(parser, _forethought_coP_pred_content_predicateNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// content_predicate_2
		if !_node(parser, _content_predicate_2Node, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_predicate_1}
	// coP_pred_content_predicate/forethought_coP_pred_content_predicate/content_predicate_2
	{
		pos3 := pos
		// coP_pred_content_predicate
		if !_fail(parser, _coP_pred_content_predicateFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_coP_pred_content_predicate
		if !_fail(parser, _forethought_coP_pred_content_predicateFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// content_predicate_2
		if !_fail(parser, _content_predicate_2Fail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// coP_pred_content_predicate/forethought_coP_pred_content_predicate/content_predicate_2
	{
		pos3 := pos
		var node2 string
		// coP_pred_content_predicate
		if p, n := _coP_pred_content_predicateAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_coP_pred_content_predicate
		if p, n := _forethought_coP_pred_content_predicateAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// content_predicate_2
		if p, n := _content_predicate_2Action(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_predicate_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_predicate_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_content_predicate/PO_content_predicate/quotation_content_predicate/content_predicate_3
	{
		pos3 := pos
		// MI_content_predicate
		if !_accept(parser, _MI_content_predicateAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_content_predicate
		if !_accept(parser, _PO_content_predicateAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// quotation_content_predicate
		if !_accept(parser, _quotation_content_predicateAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// content_predicate_3
		if !_accept(parser, _content_predicate_3Accepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_predicate_2, start, pos, perr)
fail:
	return _memoize(parser, _content_predicate_2, start, -1, perr)
}

func _content_predicate_2Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_2}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_predicate_2"}
	// MI_content_predicate/PO_content_predicate/quotation_content_predicate/content_predicate_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// MI_content_predicate
		if !_node(parser, _MI_content_predicateNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_content_predicate
		if !_node(parser, _PO_content_predicateNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// quotation_content_predicate
		if !_node(parser, _quotation_content_predicateNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// content_predicate_3
		if !_node(parser, _content_predicate_3Node, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_predicate_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_predicate_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_predicate_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_predicate_2}
	// MI_content_predicate/PO_content_predicate/quotation_content_predicate/content_predicate_3
	{
		pos3 := pos
		// MI_content_predicate
		if !_fail(parser, _MI_content_predicateFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_content_predicate
		if !_fail(parser, _PO_content_predicateFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// quotation_content_predicate
		if !_fail(parser, _quotation_content_predicateFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// content_predicate_3
		if !_fail(parser, _content_predicate_3Fail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_predicate_2Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_2}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_content_predicate/PO_content_predicate/quotation_content_predicate/content_predicate_3
	{
		pos3 := pos
		var node2 string
		// MI_content_predicate
		if p, n := _MI_content_predicateAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// PO_content_predicate
		if p, n := _PO_content_predicateAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// quotation_content_predicate
		if p, n := _quotation_content_predicateAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// content_predicate_3
		if p, n := _content_predicate_3Action(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_predicate_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_predicate_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_syllable compound_syllable+/!function_word content_syllable
	{
		pos3 := pos
		// content_syllable compound_syllable+
		// content_syllable
		if !_accept(parser, _content_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_accept(parser, _compound_syllableAccepts, &pos, &perr) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word content_syllable
		// !function_word
		{
			pos13 := pos
			perr15 := perr
			// function_word
			if !_accept(parser, _function_wordAccepts, &pos, &perr) {
				goto ok12
			}
			pos = pos13
			perr = _max(perr15, pos)
			goto fail10
		ok12:
			pos = pos13
			perr = perr15
		}
		// content_syllable
		if !_accept(parser, _content_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_predicate_3, start, pos, perr)
fail:
	return _memoize(parser, _content_predicate_3, start, -1, perr)
}

func _content_predicate_3Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_predicate_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_3}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_predicate_3"}
	// content_syllable compound_syllable+/!function_word content_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// content_syllable compound_syllable+
		// content_syllable
		if !_node(parser, _content_syllableNode, node, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_node(parser, _compound_syllableNode, node, &pos) {
			goto fail4
		}
		for {
			nkids6 := len(node.Kids)
			pos7 := pos
			// compound_syllable
			if !_node(parser, _compound_syllableNode, node, &pos) {
				goto fail9
			}
			continue
		fail9:
			node.Kids = node.Kids[:nkids6]
			pos = pos7
			break
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// !function_word content_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(node.Kids)
			// function_word
			if !_node(parser, _function_wordNode, node, &pos) {
				goto ok12
			}
			pos = pos13
			node.Kids = node.Kids[:nkids14]
			goto fail10
		ok12:
			pos = pos13
			node.Kids = node.Kids[:nkids14]
		}
		// content_syllable
		if !_node(parser, _content_syllableNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_predicate_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_predicate_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_predicate_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_predicate_3}
	// content_syllable compound_syllable+/!function_word content_syllable
	{
		pos3 := pos
		// content_syllable compound_syllable+
		// content_syllable
		if !_fail(parser, _content_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		// compound_syllable+
		// compound_syllable
		if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		for {
			pos7 := pos
			// compound_syllable
			if !_fail(parser, _compound_syllableFail, errPos, failure, &pos) {
				goto fail9
			}
			continue
		fail9:
			pos = pos7
			break
		}
		goto ok0
	fail4:
		pos = pos3
		// !function_word content_syllable
		// !function_word
		{
			pos13 := pos
			nkids14 := len(failure.Kids)
			// function_word
			if !_fail(parser, _function_wordFail, errPos, failure, &pos) {
				goto ok12
			}
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "!function_word",
				})
			}
			goto fail10
		ok12:
			pos = pos13
			failure.Kids = failure.Kids[:nkids14]
		}
		// content_syllable
		if !_fail(parser, _content_syllableFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_predicate_3Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_predicate_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_3}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_syllable compound_syllable+/!function_word content_syllable
	{
		pos3 := pos
		var node2 string
		// content_syllable compound_syllable+
		{
			var node5 string
			// content_syllable
			if p, n := _content_syllableAction(parser, pos); n == nil {
				goto fail4
			} else {
				node5 = *n
				pos = p
			}
			node, node5 = node+node5, ""
			// compound_syllable+
			{
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail4
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
			}
			for {
				pos7 := pos
				var node8 string
				// compound_syllable
				if p, n := _compound_syllableAction(parser, pos); n == nil {
					goto fail9
				} else {
					node8 = *n
					pos = p
				}
				node5 += node8
				continue
			fail9:
				pos = pos7
				break
			}
			node, node5 = node+node5, ""
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// !function_word content_syllable
		{
			var node11 string
			// !function_word
			{
				pos13 := pos
				// function_word
				if p, n := _function_wordAction(parser, pos); n == nil {
					goto ok12
				} else {
					pos = p
				}
				pos = pos13
				goto fail10
			ok12:
				pos = pos13
				node11 = ""
			}
			node, node11 = node+node11, ""
			// content_syllable
			if p, n := _content_syllableAction(parser, pos); n == nil {
				goto fail10
			} else {
				node11 = *n
				pos = p
			}
			node, node11 = node+node11, ""
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _serial_content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _serial_content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_predicate_1 spaces? serial_predicate
	// content_predicate_1
	if !_accept(parser, _content_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_accept(parser, _serial_predicateAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _serial_content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _serial_content_predicate, start, -1, perr)
}

func _serial_content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_serial_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "serial_content_predicate"}
	// content_predicate_1 spaces? serial_predicate
	// content_predicate_1
	if !_node(parser, _content_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_node(parser, _serial_predicateNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _serial_content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _serial_content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "serial_content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _serial_content_predicate}
	// content_predicate_1 spaces? serial_predicate
	// content_predicate_1
	if !_fail(parser, _content_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// serial_predicate
	if !_fail(parser, _serial_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _serial_content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_serial_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_predicate_1 spaces? serial_predicate
	{
		var node0 string
		// content_predicate_1
		if p, n := _content_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// serial_predicate
		if p, n := _serial_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _coP_pred_content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _coP_pred_content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_predicate_2 spaces? co_bar_pred
	// content_predicate_2
	if !_accept(parser, _content_predicate_2Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_accept(parser, _co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _coP_pred_content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _coP_pred_content_predicate, start, -1, perr)
}

func _coP_pred_content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_coP_pred_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "coP_pred_content_predicate"}
	// content_predicate_2 spaces? co_bar_pred
	// content_predicate_2
	if !_node(parser, _content_predicate_2Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_node(parser, _co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _coP_pred_content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _coP_pred_content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "coP_pred_content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _coP_pred_content_predicate}
	// content_predicate_2 spaces? co_bar_pred
	// content_predicate_2
	if !_fail(parser, _content_predicate_2Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// co_bar_pred
	if !_fail(parser, _co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _coP_pred_content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_coP_pred_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _coP_pred_content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_predicate_2 spaces? co_bar_pred
	{
		var node0 string
		// content_predicate_2
		if p, n := _content_predicate_2Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// co_bar_pred
		if p, n := _co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_connective spaces? forethought_coP_pred_content_predicate_1
	// forethought_connective
	if !_accept(parser, _forethought_connectiveAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_content_predicate_1
	if !_accept(parser, _forethought_coP_pred_content_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_content_predicate, start, -1, perr)
}

func _forethought_coP_pred_content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_content_predicate"}
	// forethought_connective spaces? forethought_coP_pred_content_predicate_1
	// forethought_connective
	if !_node(parser, _forethought_connectiveNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_content_predicate_1
	if !_node(parser, _forethought_coP_pred_content_predicate_1Node, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_content_predicate}
	// forethought_connective spaces? forethought_coP_pred_content_predicate_1
	// forethought_connective
	if !_fail(parser, _forethought_connectiveFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_coP_pred_content_predicate_1
	if !_fail(parser, _forethought_coP_pred_content_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// forethought_connective spaces? forethought_coP_pred_content_predicate_1
	{
		var node0 string
		// forethought_connective
		if p, n := _forethought_connectiveAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_coP_pred_content_predicate_1
		if p, n := _forethought_coP_pred_content_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_coP_pred_content_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_coP_pred_content_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// content_predicate spaces? forethought_co_bar_pred
	// content_predicate
	if !_accept(parser, _content_predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_accept(parser, _forethought_co_bar_predAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_coP_pred_content_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _forethought_coP_pred_content_predicate_1, start, -1, perr)
}

func _forethought_coP_pred_content_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_coP_pred_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_content_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_coP_pred_content_predicate_1"}
	// content_predicate spaces? forethought_co_bar_pred
	// content_predicate
	if !_node(parser, _content_predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_node(parser, _forethought_co_bar_predNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_coP_pred_content_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_coP_pred_content_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_coP_pred_content_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_coP_pred_content_predicate_1}
	// content_predicate spaces? forethought_co_bar_pred
	// content_predicate
	if !_fail(parser, _content_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// forethought_co_bar_pred
	if !_fail(parser, _forethought_co_bar_predFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_coP_pred_content_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_coP_pred_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_coP_pred_content_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// content_predicate spaces? forethought_co_bar_pred
	{
		var node0 string
		// content_predicate
		if p, n := _content_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// forethought_co_bar_pred
		if p, n := _forethought_co_bar_predAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_contentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_content, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// LU_content_tone statement
	// LU_content_tone
	if !_accept(parser, _LU_content_toneAccepts, &pos, &perr) {
		goto fail
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_content, start, pos, perr)
fail:
	return _memoize(parser, _LU_content, start, -1, perr)
}

func _LU_contentNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_content]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_content}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_content"}
	// LU_content_tone statement
	// LU_content_tone
	if !_node(parser, _LU_content_toneNode, node, &pos) {
		goto fail
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_contentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_content, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_content",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_content}
	// LU_content_tone statement
	// LU_content_tone
	if !_fail(parser, _LU_content_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_contentAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_content]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_content}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// LU_content_tone statement
	{
		var node0 string
		// LU_content_tone
		if p, n := _LU_content_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LU_content_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU_content_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &LU content_syllable
	// &LU
	{
		pos2 := pos
		perr4 := perr
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// content_syllable
	if !_accept(parser, _content_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _LU_content_tone, start, pos, perr)
fail:
	return _memoize(parser, _LU_content_tone, start, -1, perr)
}

func _LU_content_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_content_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU_content_tone"}
	// &LU content_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// LU
		if !_node(parser, _LUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// content_syllable
	if !_node(parser, _content_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LU_content_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU_content_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU_content_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU_content_tone}
	// &LU content_syllable
	// &LU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&LU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// content_syllable
	if !_fail(parser, _content_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LU_content_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU_content_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &LU content_syllable
	{
		var node0 string
		// &LU
		{
			pos2 := pos
			// LU
			if p, n := _LUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// content_syllable
		if p, n := _content_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_content_predicate_1 spaces? end_predicatizer?
	// MI_content_predicate_1
	if !_accept(parser, _MI_content_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _MI_content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _MI_content_predicate, start, -1, perr)
}

func _MI_content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_content_predicate"}
	// MI_content_predicate_1 spaces? end_predicatizer?
	// MI_content_predicate_1
	if !_node(parser, _MI_content_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_content_predicate}
	// MI_content_predicate_1 spaces? end_predicatizer?
	// MI_content_predicate_1
	if !_fail(parser, _MI_content_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_content_predicate_1 spaces? end_predicatizer?
	{
		var node0 string
		// MI_content_predicate_1
		if p, n := _MI_content_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_content_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_content_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MI_content_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_content_tone
	if !_accept(parser, _MI_content_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_accept(parser, _predicateAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_accept(parser, _adverbAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	return _memoize(parser, _MI_content_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _MI_content_predicate_1, start, -1, perr)
}

func _MI_content_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_content_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_content_predicate_1"}
	// MI_content_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_content_tone
	if !_node(parser, _MI_content_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// predicate/argument/adverb/prepositional_phrase
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// predicate
			if !_node(parser, _predicateNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail14
			}
			goto ok7
		fail14:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			goto fail
		ok7:
		}
		sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
		node.Kids = append(node.Kids[:nkids5], sub)
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_content_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_content_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_content_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_content_predicate_1}
	// MI_content_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	// MI_content_tone
	if !_fail(parser, _MI_content_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (predicate/argument/adverb/prepositional_phrase)
	// predicate/argument/adverb/prepositional_phrase
	{
		pos8 := pos
		// predicate
		if !_fail(parser, _predicateFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// adverb
		if !_fail(parser, _adverbFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// prepositional_phrase
		if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_content_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_content_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MI_content_tone spaces? (predicate/argument/adverb/prepositional_phrase)
	{
		var node0 string
		// MI_content_tone
		if p, n := _MI_content_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (predicate/argument/adverb/prepositional_phrase)
		// predicate/argument/adverb/prepositional_phrase
		{
			pos8 := pos
			var node7 string
			// predicate
			if p, n := _predicateAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// adverb
			if p, n := _adverbAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// prepositional_phrase
			if p, n := _prepositional_phraseAction(parser, pos); n == nil {
				goto fail12
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail12:
			node0 = node7
			pos = pos8
			goto fail
		ok5:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_content_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI_content_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MI content_syllable
	// &MI
	{
		pos2 := pos
		perr4 := perr
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// content_syllable
	if !_accept(parser, _content_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MI_content_tone, start, pos, perr)
fail:
	return _memoize(parser, _MI_content_tone, start, -1, perr)
}

func _MI_content_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_content_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_content_tone"}
	// &MI content_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MI
		if !_node(parser, _MINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// content_syllable
	if !_node(parser, _content_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MI_content_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI_content_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_content_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_content_tone}
	// &MI content_syllable
	// &MI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// content_syllable
	if !_fail(parser, _content_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MI_content_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_content_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MI content_syllable
	{
		var node0 string
		// &MI
		{
			pos2 := pos
			// MI
			if p, n := _MIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// content_syllable
		if p, n := _content_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_content_predicate_1 spaces? end_predicatizer?
	// PO_content_predicate_1
	if !_accept(parser, _PO_content_predicate_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_accept(parser, _end_predicatizerAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	return _memoize(parser, _PO_content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _PO_content_predicate, start, -1, perr)
}

func _PO_content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_content_predicate"}
	// PO_content_predicate_1 spaces? end_predicatizer?
	// PO_content_predicate_1
	if !_node(parser, _PO_content_predicate_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// end_predicatizer
		if !_node(parser, _end_predicatizerNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_content_predicate}
	// PO_content_predicate_1 spaces? end_predicatizer?
	// PO_content_predicate_1
	if !_fail(parser, _PO_content_predicate_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_predicatizer?
	{
		pos6 := pos
		// end_predicatizer
		if !_fail(parser, _end_predicatizerFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_content_predicate_1 spaces? end_predicatizer?
	{
		var node0 string
		// PO_content_predicate_1
		if p, n := _PO_content_predicate_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_predicatizer?
		{
			pos6 := pos
			// end_predicatizer
			if p, n := _end_predicatizerAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_content_predicate_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_content_predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// PO_content_tone spaces? argument
	// PO_content_tone
	if !_accept(parser, _PO_content_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_content_predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _PO_content_predicate_1, start, -1, perr)
}

func _PO_content_predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_content_predicate_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_content_predicate_1"}
	// PO_content_tone spaces? argument
	// PO_content_tone
	if !_node(parser, _PO_content_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_content_predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_content_predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_content_predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_content_predicate_1}
	// PO_content_tone spaces? argument
	// PO_content_tone
	if !_fail(parser, _PO_content_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_content_predicate_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_content_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// PO_content_tone spaces? argument
	{
		var node0 string
		// PO_content_tone
		if p, n := _PO_content_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _PO_content_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO_content_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &PO content_syllable
	// &PO
	{
		pos2 := pos
		perr4 := perr
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// content_syllable
	if !_accept(parser, _content_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _PO_content_tone, start, pos, perr)
fail:
	return _memoize(parser, _PO_content_tone, start, -1, perr)
}

func _PO_content_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_content_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO_content_tone"}
	// &PO content_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// content_syllable
	if !_node(parser, _content_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _PO_content_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO_content_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO_content_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO_content_tone}
	// &PO content_syllable
	// &PO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&PO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// content_syllable
	if !_fail(parser, _content_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _PO_content_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO_content_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &PO content_syllable
	{
		var node0 string
		// &PO
		{
			pos2 := pos
			// PO
			if p, n := _POAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// content_syllable
		if p, n := _content_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quotation_content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _quotation_content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_content_predicate spaces? end_quote
	// MO_content_predicate
	if !_accept(parser, _MO_content_predicateAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_accept(parser, _end_quoteAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _quotation_content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _quotation_content_predicate, start, -1, perr)
}

func _quotation_content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_quotation_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quotation_content_predicate"}
	// MO_content_predicate spaces? end_quote
	// MO_content_predicate
	if !_node(parser, _MO_content_predicateNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_quote
	if !_node(parser, _end_quoteNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quotation_content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _quotation_content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quotation_content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quotation_content_predicate}
	// MO_content_predicate spaces? end_quote
	// MO_content_predicate
	if !_fail(parser, _MO_content_predicateFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_quote
	if !_fail(parser, _end_quoteFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quotation_content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_quotation_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quotation_content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_content_predicate spaces? end_quote
	{
		var node0 string
		// MO_content_predicate
		if p, n := _MO_content_predicateAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_quote
		if p, n := _end_quoteAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_content_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_content_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// MO_content_predicate_tone spaces? discourse
	// MO_content_predicate_tone
	if !_accept(parser, _MO_content_predicate_toneAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_accept(parser, _discourseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_content_predicate, start, pos, perr)
fail:
	return _memoize(parser, _MO_content_predicate, start, -1, perr)
}

func _MO_content_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_content_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_content_predicate"}
	// MO_content_predicate_tone spaces? discourse
	// MO_content_predicate_tone
	if !_node(parser, _MO_content_predicate_toneNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// discourse
	if !_node(parser, _discourseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_content_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_content_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_content_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_content_predicate}
	// MO_content_predicate_tone spaces? discourse
	// MO_content_predicate_tone
	if !_fail(parser, _MO_content_predicate_toneFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_fail(parser, _discourseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_content_predicateAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// MO_content_predicate_tone spaces? discourse
	{
		var node0 string
		// MO_content_predicate_tone
		if p, n := _MO_content_predicate_toneAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// discourse
		if p, n := _discourseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MO_content_predicate_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO_content_predicate_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MO content_syllable
	// &MO
	{
		pos2 := pos
		perr4 := perr
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// content_syllable
	if !_accept(parser, _content_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _MO_content_predicate_tone, start, pos, perr)
fail:
	return _memoize(parser, _MO_content_predicate_tone, start, -1, perr)
}

func _MO_content_predicate_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO_content_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_content_predicate_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO_content_predicate_tone"}
	// &MO content_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MO
		if !_node(parser, _MONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// content_syllable
	if !_node(parser, _content_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MO_content_predicate_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO_content_predicate_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO_content_predicate_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO_content_predicate_tone}
	// &MO content_syllable
	// &MO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// content_syllable
	if !_fail(parser, _content_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MO_content_predicate_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO_content_predicate_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO_content_predicate_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MO content_syllable
	{
		var node0 string
		// &MO
		{
			pos2 := pos
			// MO
			if p, n := _MOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// content_syllable
		if p, n := _content_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _freemodAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _freemod, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// spaces? (interjection/parenthetical/incidental/vocative) spaces? freemod?
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (interjection/parenthetical/incidental/vocative)
	// interjection/parenthetical/incidental/vocative
	{
		pos8 := pos
		// interjection
		if !_accept(parser, _interjectionAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// parenthetical
		if !_accept(parser, _parentheticalAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// incidental
		if !_accept(parser, _incidentalAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// vocative
		if !_accept(parser, _vocativeAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	// spaces?
	{
		pos14 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail15
		}
		goto ok16
	fail15:
		pos = pos14
	ok16:
	}
	// freemod?
	{
		pos18 := pos
		// freemod
		if !_accept(parser, _freemodAccepts, &pos, &perr) {
			goto fail19
		}
		goto ok20
	fail19:
		pos = pos18
	ok20:
	}
	return _memoize(parser, _freemod, start, pos, perr)
fail:
	return _memoize(parser, _freemod, start, -1, perr)
}

func _freemodNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_freemod]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _freemod}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "freemod"}
	// spaces? (interjection/parenthetical/incidental/vocative) spaces? freemod?
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// (interjection/parenthetical/incidental/vocative)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// interjection/parenthetical/incidental/vocative
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// interjection
			if !_node(parser, _interjectionNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// parenthetical
			if !_node(parser, _parentheticalNode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// incidental
			if !_node(parser, _incidentalNode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// vocative
			if !_node(parser, _vocativeNode, node, &pos) {
				goto fail14
			}
			goto ok7
		fail14:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			goto fail
		ok7:
		}
		sub := _sub(parser, pos06, pos, node.Kids[nkids5:])
		node.Kids = append(node.Kids[:nkids5], sub)
	}
	// spaces?
	{
		nkids15 := len(node.Kids)
		pos16 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail17
		}
		goto ok18
	fail17:
		node.Kids = node.Kids[:nkids15]
		pos = pos16
	ok18:
	}
	// freemod?
	{
		nkids19 := len(node.Kids)
		pos20 := pos
		// freemod
		if !_node(parser, _freemodNode, node, &pos) {
			goto fail21
		}
		goto ok22
	fail21:
		node.Kids = node.Kids[:nkids19]
		pos = pos20
	ok22:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _freemodFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _freemod, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "freemod",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _freemod}
	// spaces? (interjection/parenthetical/incidental/vocative) spaces? freemod?
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// (interjection/parenthetical/incidental/vocative)
	// interjection/parenthetical/incidental/vocative
	{
		pos8 := pos
		// interjection
		if !_fail(parser, _interjectionFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// parenthetical
		if !_fail(parser, _parentheticalFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// incidental
		if !_fail(parser, _incidentalFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// vocative
		if !_fail(parser, _vocativeFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	// spaces?
	{
		pos14 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok16
	fail15:
		pos = pos14
	ok16:
	}
	// freemod?
	{
		pos18 := pos
		// freemod
		if !_fail(parser, _freemodFail, errPos, failure, &pos) {
			goto fail19
		}
		goto ok20
	fail19:
		pos = pos18
	ok20:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _freemodAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_freemod]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _freemod}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// spaces? (interjection/parenthetical/incidental/vocative) spaces? freemod?
	{
		var node0 string
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// (interjection/parenthetical/incidental/vocative)
		// interjection/parenthetical/incidental/vocative
		{
			pos8 := pos
			var node7 string
			// interjection
			if p, n := _interjectionAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// parenthetical
			if p, n := _parentheticalAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// incidental
			if p, n := _incidentalAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// vocative
			if p, n := _vocativeAction(parser, pos); n == nil {
				goto fail12
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail12:
			node0 = node7
			pos = pos8
			goto fail
		ok5:
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos14 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail15
			} else {
				node0 = *n
				pos = p
			}
			goto ok16
		fail15:
			node0 = ""
			pos = pos14
		ok16:
		}
		node, node0 = node+node0, ""
		// freemod?
		{
			pos18 := pos
			// freemod
			if p, n := _freemodAction(parser, pos); n == nil {
				goto fail19
			} else {
				node0 = *n
				pos = p
			}
			goto ok20
		fail19:
			node0 = ""
			pos = pos18
		ok20:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _parentheticalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _parenthetical, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// parenthetical_1 spaces? end_parenthetical
	// parenthetical_1
	if !_accept(parser, _parenthetical_1Accepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_parenthetical
	if !_accept(parser, _end_parentheticalAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _parenthetical, start, pos, perr)
fail:
	return _memoize(parser, _parenthetical, start, -1, perr)
}

func _parentheticalNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _parenthetical}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "parenthetical"}
	// parenthetical_1 spaces? end_parenthetical
	// parenthetical_1
	if !_node(parser, _parenthetical_1Node, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// end_parenthetical
	if !_node(parser, _end_parentheticalNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _parentheticalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _parenthetical, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "parenthetical",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _parenthetical}
	// parenthetical_1 spaces? end_parenthetical
	// parenthetical_1
	if !_fail(parser, _parenthetical_1Fail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// end_parenthetical
	if !_fail(parser, _end_parentheticalFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _parentheticalAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _parenthetical}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// parenthetical_1 spaces? end_parenthetical
	{
		var node0 string
		// parenthetical_1
		if p, n := _parenthetical_1Action(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// end_parenthetical
		if p, n := _end_parentheticalAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _parenthetical_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _parenthetical_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// start_parenthetical spaces? discourse
	// start_parenthetical
	if !_accept(parser, _start_parentheticalAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_accept(parser, _discourseAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _parenthetical_1, start, pos, perr)
fail:
	return _memoize(parser, _parenthetical_1, start, -1, perr)
}

func _parenthetical_1Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_parenthetical_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _parenthetical_1}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "parenthetical_1"}
	// start_parenthetical spaces? discourse
	// start_parenthetical
	if !_node(parser, _start_parentheticalNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// discourse
	if !_node(parser, _discourseNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _parenthetical_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _parenthetical_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "parenthetical_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _parenthetical_1}
	// start_parenthetical spaces? discourse
	// start_parenthetical
	if !_fail(parser, _start_parentheticalFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// discourse
	if !_fail(parser, _discourseFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _parenthetical_1Action(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_parenthetical_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _parenthetical_1}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// start_parenthetical spaces? discourse
	{
		var node0 string
		// start_parenthetical
		if p, n := _start_parentheticalAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// discourse
		if p, n := _discourseAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _incidentalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _incidental, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// spaces? start_incidental freemod? statement
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// start_incidental
	if !_accept(parser, _start_incidentalAccepts, &pos, &perr) {
		goto fail
	}
	// freemod?
	{
		pos6 := pos
		// freemod
		if !_accept(parser, _freemodAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// statement
	if !_accept(parser, _statementAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _incidental, start, pos, perr)
fail:
	return _memoize(parser, _incidental, start, -1, perr)
}

func _incidentalNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_incidental]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _incidental}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "incidental"}
	// spaces? start_incidental freemod? statement
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// start_incidental
	if !_node(parser, _start_incidentalNode, node, &pos) {
		goto fail
	}
	// freemod?
	{
		nkids5 := len(node.Kids)
		pos6 := pos
		// freemod
		if !_node(parser, _freemodNode, node, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		node.Kids = node.Kids[:nkids5]
		pos = pos6
	ok8:
	}
	// statement
	if !_node(parser, _statementNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _incidentalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _incidental, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "incidental",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _incidental}
	// spaces? start_incidental freemod? statement
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// start_incidental
	if !_fail(parser, _start_incidentalFail, errPos, failure, &pos) {
		goto fail
	}
	// freemod?
	{
		pos6 := pos
		// freemod
		if !_fail(parser, _freemodFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok8
	fail7:
		pos = pos6
	ok8:
	}
	// statement
	if !_fail(parser, _statementFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _incidentalAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_incidental]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _incidental}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// spaces? start_incidental freemod? statement
	{
		var node0 string
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// start_incidental
		if p, n := _start_incidentalAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// freemod?
		{
			pos6 := pos
			// freemod
			if p, n := _freemodAction(parser, pos); n == nil {
				goto fail7
			} else {
				node0 = *n
				pos = p
			}
			goto ok8
		fail7:
			node0 = ""
			pos = pos6
		ok8:
		}
		node, node0 = node+node0, ""
		// statement
		if p, n := _statementAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _vocativeAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _vocative, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// vocative_marker spaces? argument
	// vocative_marker
	if !_accept(parser, _vocative_markerAccepts, &pos, &perr) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_accept(parser, _spacesAccepts, &pos, &perr) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_accept(parser, _argumentAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _vocative, start, pos, perr)
fail:
	return _memoize(parser, _vocative, start, -1, perr)
}

func _vocativeNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_vocative]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _vocative}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "vocative"}
	// vocative_marker spaces? argument
	// vocative_marker
	if !_node(parser, _vocative_markerNode, node, &pos) {
		goto fail
	}
	// spaces?
	{
		nkids1 := len(node.Kids)
		pos2 := pos
		// spaces
		if !_node(parser, _spacesNode, node, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		node.Kids = node.Kids[:nkids1]
		pos = pos2
	ok4:
	}
	// argument
	if !_node(parser, _argumentNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _vocativeFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _vocative, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "vocative",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _vocative}
	// vocative_marker spaces? argument
	// vocative_marker
	if !_fail(parser, _vocative_markerFail, errPos, failure, &pos) {
		goto fail
	}
	// spaces?
	{
		pos2 := pos
		// spaces
		if !_fail(parser, _spacesFail, errPos, failure, &pos) {
			goto fail3
		}
		goto ok4
	fail3:
		pos = pos2
	ok4:
	}
	// argument
	if !_fail(parser, _argumentFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _vocativeAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_vocative]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _vocative}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// vocative_marker spaces? argument
	{
		var node0 string
		// vocative_marker
		if p, n := _vocative_markerAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// spaces?
		{
			pos2 := pos
			// spaces
			if p, n := _spacesAction(parser, pos); n == nil {
				goto fail3
			} else {
				node0 = *n
				pos = p
			}
			goto ok4
		fail3:
			node0 = ""
			pos = pos2
		ok4:
		}
		node, node0 = node+node0, ""
		// argument
		if p, n := _argumentAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prefixAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _prefix, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &MU neutral_syllable freemod?
	// &MU
	{
		pos2 := pos
		perr4 := perr
		// MU
		if !_accept(parser, _MUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	// freemod?
	{
		pos7 := pos
		// freemod
		if !_accept(parser, _freemodAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok9
	fail8:
		pos = pos7
	ok9:
	}
	return _memoize(parser, _prefix, start, pos, perr)
fail:
	return _memoize(parser, _prefix, start, -1, perr)
}

func _prefixNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_prefix]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prefix}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "prefix"}
	// &MU neutral_syllable freemod?
	// &MU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MU
		if !_node(parser, _MUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	// freemod?
	{
		nkids6 := len(node.Kids)
		pos7 := pos
		// freemod
		if !_node(parser, _freemodNode, node, &pos) {
			goto fail8
		}
		goto ok9
	fail8:
		node.Kids = node.Kids[:nkids6]
		pos = pos7
	ok9:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prefixFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _prefix, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prefix",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prefix}
	// &MU neutral_syllable freemod?
	// &MU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MU
		if !_fail(parser, _MUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	// freemod?
	{
		pos7 := pos
		// freemod
		if !_fail(parser, _freemodFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok9
	fail8:
		pos = pos7
	ok9:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prefixAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_prefix]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prefix}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &MU neutral_syllable freemod?
	{
		var node0 string
		// &MU
		{
			pos2 := pos
			// MU
			if p, n := _MUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// freemod?
		{
			pos7 := pos
			// freemod
			if p, n := _freemodAction(parser, pos); n == nil {
				goto fail8
			} else {
				node0 = *n
				pos = p
			}
			goto ok9
		fail8:
			node0 = ""
			pos = pos7
		ok9:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _focusAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _focus, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &KU neutral_syllable
	// &KU
	{
		pos2 := pos
		perr4 := perr
		// KU
		if !_accept(parser, _KUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _focus, start, pos, perr)
fail:
	return _memoize(parser, _focus, start, -1, perr)
}

func _focusNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_focus]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _focus}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "focus"}
	// &KU neutral_syllable
	// &KU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// KU
		if !_node(parser, _KUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _focusFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _focus, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "focus",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _focus}
	// &KU neutral_syllable
	// &KU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// KU
		if !_fail(parser, _KUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&KU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _focusAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_focus]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _focus}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &KU neutral_syllable
	{
		var node0 string
		// &KU
		{
			pos2 := pos
			// KU
			if p, n := _KUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_quoteAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _end_quote, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &TEO neutral_syllable
	// &TEO
	{
		pos2 := pos
		perr4 := perr
		// TEO
		if !_accept(parser, _TEOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _end_quote, start, pos, perr)
fail:
	return _memoize(parser, _end_quote, start, -1, perr)
}

func _end_quoteNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_end_quote]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_quote}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "end_quote"}
	// &TEO neutral_syllable
	// &TEO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// TEO
		if !_node(parser, _TEONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_quoteFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _end_quote, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_quote",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_quote}
	// &TEO neutral_syllable
	// &TEO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// TEO
		if !_fail(parser, _TEOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&TEO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_quoteAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_end_quote]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_quote}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &TEO neutral_syllable
	{
		var node0 string
		// &TEO
		{
			pos2 := pos
			// TEO
			if p, n := _TEOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_predicatizerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _end_predicatizer, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &GA neutral_syllable
	// &GA
	{
		pos2 := pos
		perr4 := perr
		// GA
		if !_accept(parser, _GAAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _end_predicatizer, start, pos, perr)
fail:
	return _memoize(parser, _end_predicatizer, start, -1, perr)
}

func _end_predicatizerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_end_predicatizer]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_predicatizer}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "end_predicatizer"}
	// &GA neutral_syllable
	// &GA
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// GA
		if !_node(parser, _GANode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_predicatizerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _end_predicatizer, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_predicatizer",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_predicatizer}
	// &GA neutral_syllable
	// &GA
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// GA
		if !_fail(parser, _GAFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&GA",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_predicatizerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_end_predicatizer]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_predicatizer}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &GA neutral_syllable
	{
		var node0 string
		// &GA
		{
			pos2 := pos
			// GA
			if p, n := _GAAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _end_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &NA neutral_syllable freemod?
	// &NA
	{
		pos2 := pos
		perr4 := perr
		// NA
		if !_accept(parser, _NAAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	// freemod?
	{
		pos7 := pos
		// freemod
		if !_accept(parser, _freemodAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok9
	fail8:
		pos = pos7
	ok9:
	}
	return _memoize(parser, _end_statement, start, pos, perr)
fail:
	return _memoize(parser, _end_statement, start, -1, perr)
}

func _end_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_end_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "end_statement"}
	// &NA neutral_syllable freemod?
	// &NA
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// NA
		if !_node(parser, _NANode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	// freemod?
	{
		nkids6 := len(node.Kids)
		pos7 := pos
		// freemod
		if !_node(parser, _freemodNode, node, &pos) {
			goto fail8
		}
		goto ok9
	fail8:
		node.Kids = node.Kids[:nkids6]
		pos = pos7
	ok9:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _end_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_statement}
	// &NA neutral_syllable freemod?
	// &NA
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// NA
		if !_fail(parser, _NAFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&NA",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	// freemod?
	{
		pos7 := pos
		// freemod
		if !_fail(parser, _freemodFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok9
	fail8:
		pos = pos7
	ok9:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_statementAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_end_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &NA neutral_syllable freemod?
	{
		var node0 string
		// &NA
		{
			pos2 := pos
			// NA
			if p, n := _NAAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// freemod?
		{
			pos7 := pos
			// freemod
			if p, n := _freemodAction(parser, pos); n == nil {
				goto fail8
			} else {
				node0 = *n
				pos = p
			}
			goto ok9
		fail8:
			node0 = ""
			pos = pos7
		ok9:
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentence_prefixAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _sentence_prefix, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &KEO neutral_syllable
	// &KEO
	{
		pos2 := pos
		perr4 := perr
		// KEO
		if !_accept(parser, _KEOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _sentence_prefix, start, pos, perr)
fail:
	return _memoize(parser, _sentence_prefix, start, -1, perr)
}

func _sentence_prefixNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_sentence_prefix]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_prefix}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "sentence_prefix"}
	// &KEO neutral_syllable
	// &KEO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// KEO
		if !_node(parser, _KEONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_prefixFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _sentence_prefix, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_prefix",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_prefix}
	// &KEO neutral_syllable
	// &KEO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// KEO
		if !_fail(parser, _KEOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&KEO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_prefixAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_sentence_prefix]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_prefix}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &KEO neutral_syllable
	{
		var node0 string
		// &KEO
		{
			pos2 := pos
			// KEO
			if p, n := _KEOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_prenexAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _end_prenex, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &BI neutral_syllable
	// &BI
	{
		pos2 := pos
		perr4 := perr
		// BI
		if !_accept(parser, _BIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _end_prenex, start, pos, perr)
fail:
	return _memoize(parser, _end_prenex, start, -1, perr)
}

func _end_prenexNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_end_prenex]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_prenex}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "end_prenex"}
	// &BI neutral_syllable
	// &BI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// BI
		if !_node(parser, _BINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_prenexFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _end_prenex, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_prenex",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_prenex}
	// &BI neutral_syllable
	// &BI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// BI
		if !_fail(parser, _BIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&BI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_prenexAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_end_prenex]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_prenex}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &BI neutral_syllable
	{
		var node0 string
		// &BI
		{
			pos2 := pos
			// BI
			if p, n := _BIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _start_incidentalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _start_incidental, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &JU neutral_syllable
	// &JU
	{
		pos2 := pos
		perr4 := perr
		// JU
		if !_accept(parser, _JUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _start_incidental, start, pos, perr)
fail:
	return _memoize(parser, _start_incidental, start, -1, perr)
}

func _start_incidentalNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_start_incidental]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _start_incidental}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "start_incidental"}
	// &JU neutral_syllable
	// &JU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// JU
		if !_node(parser, _JUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _start_incidentalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _start_incidental, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "start_incidental",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _start_incidental}
	// &JU neutral_syllable
	// &JU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// JU
		if !_fail(parser, _JUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&JU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _start_incidentalAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_start_incidental]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _start_incidental}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &JU neutral_syllable
	{
		var node0 string
		// &JU
		{
			pos2 := pos
			// JU
			if p, n := _JUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _start_parentheticalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _start_parenthetical, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &KIO neutral_syllable
	// &KIO
	{
		pos2 := pos
		perr4 := perr
		// KIO
		if !_accept(parser, _KIOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _start_parenthetical, start, pos, perr)
fail:
	return _memoize(parser, _start_parenthetical, start, -1, perr)
}

func _start_parentheticalNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_start_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _start_parenthetical}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "start_parenthetical"}
	// &KIO neutral_syllable
	// &KIO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// KIO
		if !_node(parser, _KIONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _start_parentheticalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _start_parenthetical, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "start_parenthetical",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _start_parenthetical}
	// &KIO neutral_syllable
	// &KIO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// KIO
		if !_fail(parser, _KIOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&KIO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _start_parentheticalAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_start_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _start_parenthetical}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &KIO neutral_syllable
	{
		var node0 string
		// &KIO
		{
			pos2 := pos
			// KIO
			if p, n := _KIOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_parentheticalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _end_parenthetical, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &KI neutral_syllable
	// &KI
	{
		pos2 := pos
		perr4 := perr
		// KI
		if !_accept(parser, _KIAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _end_parenthetical, start, pos, perr)
fail:
	return _memoize(parser, _end_parenthetical, start, -1, perr)
}

func _end_parentheticalNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_end_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_parenthetical}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "end_parenthetical"}
	// &KI neutral_syllable
	// &KI
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// KI
		if !_node(parser, _KINode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_parentheticalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _end_parenthetical, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_parenthetical",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_parenthetical}
	// &KI neutral_syllable
	// &KI
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// KI
		if !_fail(parser, _KIFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&KI",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_parentheticalAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_end_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_parenthetical}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &KI neutral_syllable
	{
		var node0 string
		// &KI
		{
			pos2 := pos
			// KI
			if p, n := _KIAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _vocative_markerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _vocative_marker, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &HU neutral_syllable
	// &HU
	{
		pos2 := pos
		perr4 := perr
		// HU
		if !_accept(parser, _HUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _vocative_marker, start, pos, perr)
fail:
	return _memoize(parser, _vocative_marker, start, -1, perr)
}

func _vocative_markerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_vocative_marker]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _vocative_marker}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "vocative_marker"}
	// &HU neutral_syllable
	// &HU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// HU
		if !_node(parser, _HUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _vocative_markerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _vocative_marker, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "vocative_marker",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _vocative_marker}
	// &HU neutral_syllable
	// &HU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// HU
		if !_fail(parser, _HUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&HU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _vocative_markerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_vocative_marker]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _vocative_marker}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &HU neutral_syllable
	{
		var node0 string
		// &HU
		{
			pos2 := pos
			// HU
			if p, n := _HUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _linking_wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _linking_word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &GO neutral_syllable
	// &GO
	{
		pos2 := pos
		perr4 := perr
		// GO
		if !_accept(parser, _GOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _linking_word, start, pos, perr)
fail:
	return _memoize(parser, _linking_word, start, -1, perr)
}

func _linking_wordNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_linking_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _linking_word}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "linking_word"}
	// &GO neutral_syllable
	// &GO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// GO
		if !_node(parser, _GONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _linking_wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _linking_word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "linking_word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _linking_word}
	// &GO neutral_syllable
	// &GO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// GO
		if !_fail(parser, _GOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&GO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _linking_wordAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_linking_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _linking_word}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &GO neutral_syllable
	{
		var node0 string
		// &GO
		{
			pos2 := pos
			// GO
			if p, n := _GOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _connectiveAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _connective, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &RA neutral_syllable
	// &RA
	{
		pos2 := pos
		perr4 := perr
		// RA
		if !_accept(parser, _RAAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _connective, start, pos, perr)
fail:
	return _memoize(parser, _connective, start, -1, perr)
}

func _connectiveNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_connective]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _connective}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "connective"}
	// &RA neutral_syllable
	// &RA
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// RA
		if !_node(parser, _RANode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _connectiveFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _connective, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "connective",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _connective}
	// &RA neutral_syllable
	// &RA
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// RA
		if !_fail(parser, _RAFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&RA",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _connectiveAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_connective]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _connective}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &RA neutral_syllable
	{
		var node0 string
		// &RA
		{
			pos2 := pos
			// RA
			if p, n := _RAAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _illocutionaryAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _illocutionary, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &DA neutral_syllable
	// &DA
	{
		pos2 := pos
		perr4 := perr
		// DA
		if !_accept(parser, _DAAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _illocutionary, start, pos, perr)
fail:
	return _memoize(parser, _illocutionary, start, -1, perr)
}

func _illocutionaryNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_illocutionary]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _illocutionary}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "illocutionary"}
	// &DA neutral_syllable
	// &DA
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// DA
		if !_node(parser, _DANode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _illocutionaryFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _illocutionary, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "illocutionary",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _illocutionary}
	// &DA neutral_syllable
	// &DA
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// DA
		if !_fail(parser, _DAFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&DA",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _illocutionaryAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_illocutionary]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _illocutionary}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &DA neutral_syllable
	{
		var node0 string
		// &DA
		{
			pos2 := pos
			// DA
			if p, n := _DAAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quantifierAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _quantifier, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &TU neutral_syllable
	// &TU
	{
		pos2 := pos
		perr4 := perr
		// TU
		if !_accept(parser, _TUAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _quantifier, start, pos, perr)
fail:
	return _memoize(parser, _quantifier, start, -1, perr)
}

func _quantifierNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_quantifier]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quantifier}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quantifier"}
	// &TU neutral_syllable
	// &TU
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// TU
		if !_node(parser, _TUNode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quantifierFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _quantifier, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quantifier",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quantifier}
	// &TU neutral_syllable
	// &TU
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// TU
		if !_fail(parser, _TUFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&TU",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quantifierAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_quantifier]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quantifier}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &TU neutral_syllable
	{
		var node0 string
		// &TU
		{
			pos2 := pos
			// TU
			if p, n := _TUAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _interjectionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _interjection, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &HA neutral_syllable
	// &HA
	{
		pos2 := pos
		perr4 := perr
		// HA
		if !_accept(parser, _HAAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _interjection, start, pos, perr)
fail:
	return _memoize(parser, _interjection, start, -1, perr)
}

func _interjectionNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_interjection]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _interjection}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "interjection"}
	// &HA neutral_syllable
	// &HA
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// HA
		if !_node(parser, _HANode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _interjectionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _interjection, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "interjection",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _interjection}
	// &HA neutral_syllable
	// &HA
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// HA
		if !_fail(parser, _HAFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&HA",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _interjectionAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_interjection]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _interjection}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &HA neutral_syllable
	{
		var node0 string
		// &HA
		{
			pos2 := pos
			// HA
			if p, n := _HAAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_markerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _forethought_marker, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &TO neutral_syllable
	// &TO
	{
		pos2 := pos
		perr4 := perr
		// TO
		if !_accept(parser, _TOAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _forethought_marker, start, pos, perr)
fail:
	return _memoize(parser, _forethought_marker, start, -1, perr)
}

func _forethought_markerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_forethought_marker]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_marker}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "forethought_marker"}
	// &TO neutral_syllable
	// &TO
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// TO
		if !_node(parser, _TONode, node, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// neutral_syllable
	if !_node(parser, _neutral_syllableNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_markerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _forethought_marker, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_marker",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_marker}
	// &TO neutral_syllable
	// &TO
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// TO
		if !_fail(parser, _TOFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&TO",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// neutral_syllable
	if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_markerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_forethought_marker]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_marker}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// &TO neutral_syllable
	{
		var node0 string
		// &TO
		{
			pos2 := pos
			// TO
			if p, n := _TOAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// neutral_syllable
		if p, n := _neutral_syllableAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _function_wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _function_word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/NA/PO/RA/TU/TO/TEO
	{
		pos3 := pos
		// BI
		if !_accept(parser, _BIAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// DA
		if !_accept(parser, _DAAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// GA
		if !_accept(parser, _GAAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// GO
		if !_accept(parser, _GOAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// HA
		if !_accept(parser, _HAAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// HU
		if !_accept(parser, _HUAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// JU
		if !_accept(parser, _JUAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// KU
		if !_accept(parser, _KUAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		// KI
		if !_accept(parser, _KIAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok0
	fail12:
		pos = pos3
		// KIO
		if !_accept(parser, _KIOAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok0
	fail13:
		pos = pos3
		// KEO
		if !_accept(parser, _KEOAccepts, &pos, &perr) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		// MU
		if !_accept(parser, _MUAccepts, &pos, &perr) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail17
		}
		goto ok0
	fail17:
		pos = pos3
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail18
		}
		goto ok0
	fail18:
		pos = pos3
		// NA
		if !_accept(parser, _NAAccepts, &pos, &perr) {
			goto fail19
		}
		goto ok0
	fail19:
		pos = pos3
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail20
		}
		goto ok0
	fail20:
		pos = pos3
		// RA
		if !_accept(parser, _RAAccepts, &pos, &perr) {
			goto fail21
		}
		goto ok0
	fail21:
		pos = pos3
		// TU
		if !_accept(parser, _TUAccepts, &pos, &perr) {
			goto fail22
		}
		goto ok0
	fail22:
		pos = pos3
		// TO
		if !_accept(parser, _TOAccepts, &pos, &perr) {
			goto fail23
		}
		goto ok0
	fail23:
		pos = pos3
		// TEO
		if !_accept(parser, _TEOAccepts, &pos, &perr) {
			goto fail24
		}
		goto ok0
	fail24:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _function_word, start, pos, perr)
fail:
	return _memoize(parser, _function_word, start, -1, perr)
}

func _function_wordNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _function_word}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "function_word"}
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/NA/PO/RA/TU/TO/TEO
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// BI
		if !_node(parser, _BINode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// DA
		if !_node(parser, _DANode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// GA
		if !_node(parser, _GANode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// GO
		if !_node(parser, _GONode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// HA
		if !_node(parser, _HANode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// HU
		if !_node(parser, _HUNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// JU
		if !_node(parser, _JUNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// KU
		if !_node(parser, _KUNode, node, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// KI
		if !_node(parser, _KINode, node, &pos) {
			goto fail12
		}
		goto ok0
	fail12:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// KIO
		if !_node(parser, _KIONode, node, &pos) {
			goto fail13
		}
		goto ok0
	fail13:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// KEO
		if !_node(parser, _KEONode, node, &pos) {
			goto fail14
		}
		goto ok0
	fail14:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU
		if !_node(parser, _LUNode, node, &pos) {
			goto fail15
		}
		goto ok0
	fail15:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MU
		if !_node(parser, _MUNode, node, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI
		if !_node(parser, _MINode, node, &pos) {
			goto fail17
		}
		goto ok0
	fail17:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MO
		if !_node(parser, _MONode, node, &pos) {
			goto fail18
		}
		goto ok0
	fail18:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// NA
		if !_node(parser, _NANode, node, &pos) {
			goto fail19
		}
		goto ok0
	fail19:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail20
		}
		goto ok0
	fail20:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// RA
		if !_node(parser, _RANode, node, &pos) {
			goto fail21
		}
		goto ok0
	fail21:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// TU
		if !_node(parser, _TUNode, node, &pos) {
			goto fail22
		}
		goto ok0
	fail22:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// TO
		if !_node(parser, _TONode, node, &pos) {
			goto fail23
		}
		goto ok0
	fail23:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// TEO
		if !_node(parser, _TEONode, node, &pos) {
			goto fail24
		}
		goto ok0
	fail24:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		goto fail
	ok0:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _function_wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _function_word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "function_word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _function_word}
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/NA/PO/RA/TU/TO/TEO
	{
		pos3 := pos
		// BI
		if !_fail(parser, _BIFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// DA
		if !_fail(parser, _DAFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// GA
		if !_fail(parser, _GAFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// GO
		if !_fail(parser, _GOFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// HA
		if !_fail(parser, _HAFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// HU
		if !_fail(parser, _HUFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// JU
		if !_fail(parser, _JUFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// KU
		if !_fail(parser, _KUFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		// KI
		if !_fail(parser, _KIFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok0
	fail12:
		pos = pos3
		// KIO
		if !_fail(parser, _KIOFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok0
	fail13:
		pos = pos3
		// KEO
		if !_fail(parser, _KEOFail, errPos, failure, &pos) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		// MU
		if !_fail(parser, _MUFail, errPos, failure, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail17
		}
		goto ok0
	fail17:
		pos = pos3
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail18
		}
		goto ok0
	fail18:
		pos = pos3
		// NA
		if !_fail(parser, _NAFail, errPos, failure, &pos) {
			goto fail19
		}
		goto ok0
	fail19:
		pos = pos3
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail20
		}
		goto ok0
	fail20:
		pos = pos3
		// RA
		if !_fail(parser, _RAFail, errPos, failure, &pos) {
			goto fail21
		}
		goto ok0
	fail21:
		pos = pos3
		// TU
		if !_fail(parser, _TUFail, errPos, failure, &pos) {
			goto fail22
		}
		goto ok0
	fail22:
		pos = pos3
		// TO
		if !_fail(parser, _TOFail, errPos, failure, &pos) {
			goto fail23
		}
		goto ok0
	fail23:
		pos = pos3
		// TEO
		if !_fail(parser, _TEOFail, errPos, failure, &pos) {
			goto fail24
		}
		goto ok0
	fail24:
		pos = pos3
		goto fail
	ok0:
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _function_wordAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _function_word}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/NA/PO/RA/TU/TO/TEO
	{
		pos3 := pos
		var node2 string
		// BI
		if p, n := _BIAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// DA
		if p, n := _DAAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// GA
		if p, n := _GAAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// GO
		if p, n := _GOAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// HA
		if p, n := _HAAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// HU
		if p, n := _HUAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// JU
		if p, n := _JUAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// KU
		if p, n := _KUAction(parser, pos); n == nil {
			goto fail11
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail11:
		node = node2
		pos = pos3
		// KI
		if p, n := _KIAction(parser, pos); n == nil {
			goto fail12
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail12:
		node = node2
		pos = pos3
		// KIO
		if p, n := _KIOAction(parser, pos); n == nil {
			goto fail13
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail13:
		node = node2
		pos = pos3
		// KEO
		if p, n := _KEOAction(parser, pos); n == nil {
			goto fail14
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail14:
		node = node2
		pos = pos3
		// LU
		if p, n := _LUAction(parser, pos); n == nil {
			goto fail15
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail15:
		node = node2
		pos = pos3
		// MU
		if p, n := _MUAction(parser, pos); n == nil {
			goto fail16
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail16:
		node = node2
		pos = pos3
		// MI
		if p, n := _MIAction(parser, pos); n == nil {
			goto fail17
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail17:
		node = node2
		pos = pos3
		// MO
		if p, n := _MOAction(parser, pos); n == nil {
			goto fail18
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail18:
		node = node2
		pos = pos3
		// NA
		if p, n := _NAAction(parser, pos); n == nil {
			goto fail19
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail19:
		node = node2
		pos = pos3
		// PO
		if p, n := _POAction(parser, pos); n == nil {
			goto fail20
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail20:
		node = node2
		pos = pos3
		// RA
		if p, n := _RAAction(parser, pos); n == nil {
			goto fail21
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail21:
		node = node2
		pos = pos3
		// TU
		if p, n := _TUAction(parser, pos); n == nil {
			goto fail22
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail22:
		node = node2
		pos = pos3
		// TO
		if p, n := _TOAction(parser, pos); n == nil {
			goto fail23
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail23:
		node = node2
		pos = pos3
		// TEO
		if p, n := _TEOAction(parser, pos); n == nil {
			goto fail24
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail24:
		node = node2
		pos = pos3
		goto fail
	ok0:
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _BIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _BI, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (b I/p A) &(tone? boundary)
	// (b I/p A)
	// b I/p A
	{
		pos4 := pos
		// b I
		// b
		if !_accept(parser, _bAccepts, &pos, &perr) {
			goto fail5
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// p A
		// p
		if !_accept(parser, _pAccepts, &pos, &perr) {
			goto fail7
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos10 := pos
		perr12 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos16 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail17
			}
			goto ok18
		fail17:
			pos = pos16
		ok18:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok9
	fail13:
		pos = pos10
		perr = _max(perr12, pos)
		goto fail
	ok9:
		pos = pos10
		perr = perr12
	}
	return _memoize(parser, _BI, start, pos, perr)
fail:
	return _memoize(parser, _BI, start, -1, perr)
}

func _BINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_BI]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _BI}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "BI"}
	// (b I/p A) &(tone? boundary)
	// (b I/p A)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// b I/p A
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// b I
			// b
			if !_node(parser, _bNode, node, &pos) {
				goto fail7
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// p A
			// p
			if !_node(parser, _pNode, node, &pos) {
				goto fail9
			}
			// A
			if !_node(parser, _ANode, node, &pos) {
				goto fail9
			}
			goto ok3
		fail9:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			goto fail
		ok3:
		}
		sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
		node.Kids = append(node.Kids[:nkids1], sub)
	}
	// &(tone? boundary)
	{
		pos12 := pos
		nkids13 := len(node.Kids)
		// (tone? boundary)
		{
			nkids16 := len(node.Kids)
			pos017 := pos
			// tone? boundary
			// tone?
			{
				nkids19 := len(node.Kids)
				pos20 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail21
				}
				goto ok22
			fail21:
				node.Kids = node.Kids[:nkids19]
				pos = pos20
			ok22:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail15
			}
			sub := _sub(parser, pos017, pos, node.Kids[nkids16:])
			node.Kids = append(node.Kids[:nkids16], sub)
		}
		goto ok11
	fail15:
		pos = pos12
		goto fail
	ok11:
		pos = pos12
		node.Kids = node.Kids[:nkids13]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _BIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _BI, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "BI",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _BI}
	// (b I/p A) &(tone? boundary)
	// (b I/p A)
	// b I/p A
	{
		pos4 := pos
		// b I
		// b
		if !_fail(parser, _bFail, errPos, failure, &pos) {
			goto fail5
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// p A
		// p
		if !_fail(parser, _pFail, errPos, failure, &pos) {
			goto fail7
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos10 := pos
		nkids11 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos16 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail17
			}
			goto ok18
		fail17:
			pos = pos16
		ok18:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok9
	fail13:
		pos = pos10
		failure.Kids = failure.Kids[:nkids11]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok9:
		pos = pos10
		failure.Kids = failure.Kids[:nkids11]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _BIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_BI]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _BI}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (b I/p A) &(tone? boundary)
	{
		var node0 string
		// (b I/p A)
		// b I/p A
		{
			pos4 := pos
			var node3 string
			// b I
			{
				var node6 string
				// b
				if p, n := _bAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// I
				if p, n := _IAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
			}
			goto ok1
		fail5:
			node0 = node3
			pos = pos4
			// p A
			{
				var node8 string
				// p
				if p, n := _pAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// A
				if p, n := _AAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
			}
			goto ok1
		fail7:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos10 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos16 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail17
				} else {
					pos = p
				}
				goto ok18
			fail17:
				pos = pos16
			ok18:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail13
			} else {
				pos = p
			}
			goto ok9
		fail13:
			pos = pos10
			goto fail
		ok9:
			pos = pos10
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _DAAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _DA, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (d A/m O q/k A/s O/b A) &(tone? boundary)
	// (d A/m O q/k A/s O/b A)
	// d A/m O q/k A/s O/b A
	{
		pos4 := pos
		// d A
		// d
		if !_accept(parser, _dAccepts, &pos, &perr) {
			goto fail5
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// m O q
		// m
		if !_accept(parser, _mAccepts, &pos, &perr) {
			goto fail7
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
			goto fail7
		}
		// q
		if !_accept(parser, _qAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// k A
		// k
		if !_accept(parser, _kAccepts, &pos, &perr) {
			goto fail9
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// s O
		// s
		if !_accept(parser, _sAccepts, &pos, &perr) {
			goto fail11
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// b A
		// b
		if !_accept(parser, _bAccepts, &pos, &perr) {
			goto fail13
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos16 := pos
		perr18 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos22 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail23
			}
			goto ok24
		fail23:
			pos = pos22
		ok24:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail19
		}
		goto ok15
	fail19:
		pos = pos16
		perr = _max(perr18, pos)
		goto fail
	ok15:
		pos = pos16
		perr = perr18
	}
	return _memoize(parser, _DA, start, pos, perr)
fail:
	return _memoize(parser, _DA, start, -1, perr)
}

func _DANode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_DA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _DA}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "DA"}
	// (d A/m O q/k A/s O/b A) &(tone? boundary)
	// (d A/m O q/k A/s O/b A)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// d A/m O q/k A/s O/b A
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// d A
			// d
			if !_node(parser, _dNode, node, &pos) {
				goto fail7
			}
			// A
			if !_node(parser, _ANode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// m O q
			// m
			if !_node(parser, _mNode, node, &pos) {
				goto fail9
			}
			// O
			if !_node(parser, _ONode, node, &pos) {
				goto fail9
			}
			// q
			if !_node(parser, _qNode, node, &pos) {
				goto fail9
			}
			goto ok3
		fail9:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// k A
			// k
			if !_node(parser, _kNode, node, &pos) {
				goto fail11
			}
			// A
			if !_node(parser, _ANode, node, &pos) {
				goto fail11
			}
			goto ok3
		fail11:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// s O
			// s
			if !_node(parser, _sNode, node, &pos) {
				goto fail13
			}
			// O
			if !_node(parser, _ONode, node, &pos) {
				goto fail13
			}
			goto ok3
		fail13:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// b A
			// b
			if !_node(parser, _bNode, node, &pos) {
				goto fail15
			}
			// A
			if !_node(parser, _ANode, node, &pos) {
				goto fail15
			}
			goto ok3
		fail15:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			goto fail
		ok3:
		}
		sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
		node.Kids = append(node.Kids[:nkids1], sub)
	}
	// &(tone? boundary)
	{
		pos18 := pos
		nkids19 := len(node.Kids)
		// (tone? boundary)
		{
			nkids22 := len(node.Kids)
			pos023 := pos
			// tone? boundary
			// tone?
			{
				nkids25 := len(node.Kids)
				pos26 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail27
				}
				goto ok28
			fail27:
				node.Kids = node.Kids[:nkids25]
				pos = pos26
			ok28:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail21
			}
			sub := _sub(parser, pos023, pos, node.Kids[nkids22:])
			node.Kids = append(node.Kids[:nkids22], sub)
		}
		goto ok17
	fail21:
		pos = pos18
		goto fail
	ok17:
		pos = pos18
		node.Kids = node.Kids[:nkids19]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _DAFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _DA, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "DA",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _DA}
	// (d A/m O q/k A/s O/b A) &(tone? boundary)
	// (d A/m O q/k A/s O/b A)
	// d A/m O q/k A/s O/b A
	{
		pos4 := pos
		// d A
		// d
		if !_fail(parser, _dFail, errPos, failure, &pos) {
			goto fail5
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// m O q
		// m
		if !_fail(parser, _mFail, errPos, failure, &pos) {
			goto fail7
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
			goto fail7
		}
		// q
		if !_fail(parser, _qFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// k A
		// k
		if !_fail(parser, _kFail, errPos, failure, &pos) {
			goto fail9
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// s O
		// s
		if !_fail(parser, _sFail, errPos, failure, &pos) {
			goto fail11
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// b A
		// b
		if !_fail(parser, _bFail, errPos, failure, &pos) {
			goto fail13
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos16 := pos
		nkids17 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos22 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail23
			}
			goto ok24
		fail23:
			pos = pos22
		ok24:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail19
		}
		goto ok15
	fail19:
		pos = pos16
		failure.Kids = failure.Kids[:nkids17]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok15:
		pos = pos16
		failure.Kids = failure.Kids[:nkids17]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _DAAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_DA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _DA}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (d A/m O q/k A/s O/b A) &(tone? boundary)
	{
		var node0 string
		// (d A/m O q/k A/s O/b A)
		// d A/m O q/k A/s O/b A
		{
			pos4 := pos
			var node3 string
			// d A
			{
				var node6 string
				// d
				if p, n := _dAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// A
				if p, n := _AAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
			}
			goto ok1
		fail5:
			node0 = node3
			pos = pos4
			// m O q
			{
				var node8 string
				// m
				if p, n := _mAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// O
				if p, n := _OAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// q
				if p, n := _qAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
			}
			goto ok1
		fail7:
			node0 = node3
			pos = pos4
			// k A
			{
				var node10 string
				// k
				if p, n := _kAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
				// A
				if p, n := _AAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
			}
			goto ok1
		fail9:
			node0 = node3
			pos = pos4
			// s O
			{
				var node12 string
				// s
				if p, n := _sAction(parser, pos); n == nil {
					goto fail11
				} else {
					node12 = *n
					pos = p
				}
				node0, node12 = node0+node12, ""
				// O
				if p, n := _OAction(parser, pos); n == nil {
					goto fail11
				} else {
					node12 = *n
					pos = p
				}
				node0, node12 = node0+node12, ""
			}
			goto ok1
		fail11:
			node0 = node3
			pos = pos4
			// b A
			{
				var node14 string
				// b
				if p, n := _bAction(parser, pos); n == nil {
					goto fail13
				} else {
					node14 = *n
					pos = p
				}
				node0, node14 = node0+node14, ""
				// A
				if p, n := _AAction(parser, pos); n == nil {
					goto fail13
				} else {
					node14 = *n
					pos = p
				}
				node0, node14 = node0+node14, ""
			}
			goto ok1
		fail13:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos16 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos22 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail23
				} else {
					pos = p
				}
				goto ok24
			fail23:
				pos = pos22
			ok24:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail19
			} else {
				pos = p
			}
			goto ok15
		fail19:
			pos = pos16
			goto fail
		ok15:
			pos = pos16
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _GAAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _GA, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// g A &(tone? boundary)
	// g
	if !_accept(parser, _gAccepts, &pos, &perr) {
		goto fail
	}
	// A
	if !_accept(parser, _AAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _GA, start, pos, perr)
fail:
	return _memoize(parser, _GA, start, -1, perr)
}

func _GANode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_GA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _GA}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "GA"}
	// g A &(tone? boundary)
	// g
	if !_node(parser, _gNode, node, &pos) {
		goto fail
	}
	// A
	if !_node(parser, _ANode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _GAFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _GA, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "GA",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _GA}
	// g A &(tone? boundary)
	// g
	if !_fail(parser, _gFail, errPos, failure, &pos) {
		goto fail
	}
	// A
	if !_fail(parser, _AFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _GAAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_GA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _GA}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// g A &(tone? boundary)
	{
		var node0 string
		// g
		if p, n := _gAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// A
		if p, n := _AAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _GOAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _GO, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (g O/f I/c U/k E) &(tone? boundary)
	// (g O/f I/c U/k E)
	// g O/f I/c U/k E
	{
		pos4 := pos
		// g O
		// g
		if !_accept(parser, _gAccepts, &pos, &perr) {
			goto fail5
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// f I
		// f
		if !_accept(parser, _fAccepts, &pos, &perr) {
			goto fail7
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// c U
		// c
		if !_accept(parser, _cAccepts, &pos, &perr) {
			goto fail9
		}
		// U
		if !_accept(parser, _UAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// k E
		// k
		if !_accept(parser, _kAccepts, &pos, &perr) {
			goto fail11
		}
		// E
		if !_accept(parser, _EAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos14 := pos
		perr16 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos20 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail21
			}
			goto ok22
		fail21:
			pos = pos20
		ok22:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail17
		}
		goto ok13
	fail17:
		pos = pos14
		perr = _max(perr16, pos)
		goto fail
	ok13:
		pos = pos14
		perr = perr16
	}
	return _memoize(parser, _GO, start, pos, perr)
fail:
	return _memoize(parser, _GO, start, -1, perr)
}

func _GONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_GO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _GO}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "GO"}
	// (g O/f I/c U/k E) &(tone? boundary)
	// (g O/f I/c U/k E)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// g O/f I/c U/k E
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// g O
			// g
			if !_node(parser, _gNode, node, &pos) {
				goto fail7
			}
			// O
			if !_node(parser, _ONode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// f I
			// f
			if !_node(parser, _fNode, node, &pos) {
				goto fail9
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail9
			}
			goto ok3
		fail9:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// c U
			// c
			if !_node(parser, _cNode, node, &pos) {
				goto fail11
			}
			// U
			if !_node(parser, _UNode, node, &pos) {
				goto fail11
			}
			goto ok3
		fail11:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// k E
			// k
			if !_node(parser, _kNode, node, &pos) {
				goto fail13
			}
			// E
			if !_node(parser, _ENode, node, &pos) {
				goto fail13
			}
			goto ok3
		fail13:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			goto fail
		ok3:
		}
		sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
		node.Kids = append(node.Kids[:nkids1], sub)
	}
	// &(tone? boundary)
	{
		pos16 := pos
		nkids17 := len(node.Kids)
		// (tone? boundary)
		{
			nkids20 := len(node.Kids)
			pos021 := pos
			// tone? boundary
			// tone?
			{
				nkids23 := len(node.Kids)
				pos24 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail25
				}
				goto ok26
			fail25:
				node.Kids = node.Kids[:nkids23]
				pos = pos24
			ok26:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail19
			}
			sub := _sub(parser, pos021, pos, node.Kids[nkids20:])
			node.Kids = append(node.Kids[:nkids20], sub)
		}
		goto ok15
	fail19:
		pos = pos16
		goto fail
	ok15:
		pos = pos16
		node.Kids = node.Kids[:nkids17]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _GOFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _GO, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "GO",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _GO}
	// (g O/f I/c U/k E) &(tone? boundary)
	// (g O/f I/c U/k E)
	// g O/f I/c U/k E
	{
		pos4 := pos
		// g O
		// g
		if !_fail(parser, _gFail, errPos, failure, &pos) {
			goto fail5
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// f I
		// f
		if !_fail(parser, _fFail, errPos, failure, &pos) {
			goto fail7
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// c U
		// c
		if !_fail(parser, _cFail, errPos, failure, &pos) {
			goto fail9
		}
		// U
		if !_fail(parser, _UFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// k E
		// k
		if !_fail(parser, _kFail, errPos, failure, &pos) {
			goto fail11
		}
		// E
		if !_fail(parser, _EFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos14 := pos
		nkids15 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos20 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail21
			}
			goto ok22
		fail21:
			pos = pos20
		ok22:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail17
		}
		goto ok13
	fail17:
		pos = pos14
		failure.Kids = failure.Kids[:nkids15]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok13:
		pos = pos14
		failure.Kids = failure.Kids[:nkids15]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _GOAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_GO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _GO}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (g O/f I/c U/k E) &(tone? boundary)
	{
		var node0 string
		// (g O/f I/c U/k E)
		// g O/f I/c U/k E
		{
			pos4 := pos
			var node3 string
			// g O
			{
				var node6 string
				// g
				if p, n := _gAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// O
				if p, n := _OAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
			}
			goto ok1
		fail5:
			node0 = node3
			pos = pos4
			// f I
			{
				var node8 string
				// f
				if p, n := _fAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// I
				if p, n := _IAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
			}
			goto ok1
		fail7:
			node0 = node3
			pos = pos4
			// c U
			{
				var node10 string
				// c
				if p, n := _cAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
				// U
				if p, n := _UAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
			}
			goto ok1
		fail9:
			node0 = node3
			pos = pos4
			// k E
			{
				var node12 string
				// k
				if p, n := _kAction(parser, pos); n == nil {
					goto fail11
				} else {
					node12 = *n
					pos = p
				}
				node0, node12 = node0+node12, ""
				// E
				if p, n := _EAction(parser, pos); n == nil {
					goto fail11
				} else {
					node12 = *n
					pos = p
				}
				node0, node12 = node0+node12, ""
			}
			goto ok1
		fail11:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos14 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos20 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail21
				} else {
					pos = p
				}
				goto ok22
			fail21:
				pos = pos20
			ok22:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail17
			} else {
				pos = p
			}
			goto ok13
		fail17:
			pos = pos14
			goto fail
		ok13:
			pos = pos14
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _HAAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _HA, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (h A/h U e/h I a/m) &(tone? boundary)
	// (h A/h U e/h I a/m)
	// h A/h U e/h I a/m
	{
		pos4 := pos
		// h A
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail5
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// h U e
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail7
		}
		// U
		if !_accept(parser, _UAccepts, &pos, &perr) {
			goto fail7
		}
		// e
		if !_accept(parser, _eAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// h I a
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail9
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail9
		}
		// a
		if !_accept(parser, _aAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// m
		if !_accept(parser, _mAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos13 := pos
		perr15 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos19 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail20
			}
			goto ok21
		fail20:
			pos = pos19
		ok21:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail16
		}
		goto ok12
	fail16:
		pos = pos13
		perr = _max(perr15, pos)
		goto fail
	ok12:
		pos = pos13
		perr = perr15
	}
	return _memoize(parser, _HA, start, pos, perr)
fail:
	return _memoize(parser, _HA, start, -1, perr)
}

func _HANode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_HA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _HA}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "HA"}
	// (h A/h U e/h I a/m) &(tone? boundary)
	// (h A/h U e/h I a/m)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// h A/h U e/h I a/m
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// h A
			// h
			if !_node(parser, _hNode, node, &pos) {
				goto fail7
			}
			// A
			if !_node(parser, _ANode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// h U e
			// h
			if !_node(parser, _hNode, node, &pos) {
				goto fail9
			}
			// U
			if !_node(parser, _UNode, node, &pos) {
				goto fail9
			}
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail9
			}
			goto ok3
		fail9:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// h I a
			// h
			if !_node(parser, _hNode, node, &pos) {
				goto fail11
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail11
			}
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail11
			}
			goto ok3
		fail11:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// m
			if !_node(parser, _mNode, node, &pos) {
				goto fail13
			}
			goto ok3
		fail13:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			goto fail
		ok3:
		}
		sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
		node.Kids = append(node.Kids[:nkids1], sub)
	}
	// &(tone? boundary)
	{
		pos15 := pos
		nkids16 := len(node.Kids)
		// (tone? boundary)
		{
			nkids19 := len(node.Kids)
			pos020 := pos
			// tone? boundary
			// tone?
			{
				nkids22 := len(node.Kids)
				pos23 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail24
				}
				goto ok25
			fail24:
				node.Kids = node.Kids[:nkids22]
				pos = pos23
			ok25:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail18
			}
			sub := _sub(parser, pos020, pos, node.Kids[nkids19:])
			node.Kids = append(node.Kids[:nkids19], sub)
		}
		goto ok14
	fail18:
		pos = pos15
		goto fail
	ok14:
		pos = pos15
		node.Kids = node.Kids[:nkids16]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _HAFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _HA, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "HA",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _HA}
	// (h A/h U e/h I a/m) &(tone? boundary)
	// (h A/h U e/h I a/m)
	// h A/h U e/h I a/m
	{
		pos4 := pos
		// h A
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail5
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// h U e
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail7
		}
		// U
		if !_fail(parser, _UFail, errPos, failure, &pos) {
			goto fail7
		}
		// e
		if !_fail(parser, _eFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// h I a
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail9
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail9
		}
		// a
		if !_fail(parser, _aFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// m
		if !_fail(parser, _mFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos13 := pos
		nkids14 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos19 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail20
			}
			goto ok21
		fail20:
			pos = pos19
		ok21:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail16
		}
		goto ok12
	fail16:
		pos = pos13
		failure.Kids = failure.Kids[:nkids14]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok12:
		pos = pos13
		failure.Kids = failure.Kids[:nkids14]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _HAAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_HA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _HA}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (h A/h U e/h I a/m) &(tone? boundary)
	{
		var node0 string
		// (h A/h U e/h I a/m)
		// h A/h U e/h I a/m
		{
			pos4 := pos
			var node3 string
			// h A
			{
				var node6 string
				// h
				if p, n := _hAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// A
				if p, n := _AAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
			}
			goto ok1
		fail5:
			node0 = node3
			pos = pos4
			// h U e
			{
				var node8 string
				// h
				if p, n := _hAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// U
				if p, n := _UAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
			}
			goto ok1
		fail7:
			node0 = node3
			pos = pos4
			// h I a
			{
				var node10 string
				// h
				if p, n := _hAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
				// I
				if p, n := _IAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
			}
			goto ok1
		fail9:
			node0 = node3
			pos = pos4
			// m
			if p, n := _mAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok1
		fail11:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos13 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos19 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail20
				} else {
					pos = p
				}
				goto ok21
			fail20:
				pos = pos19
			ok21:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail16
			} else {
				pos = p
			}
			goto ok12
		fail16:
			pos = pos13
			goto fail
		ok12:
			pos = pos13
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _HUAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _HU, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// h U &(tone? boundary)
	// h
	if !_accept(parser, _hAccepts, &pos, &perr) {
		goto fail
	}
	// U
	if !_accept(parser, _UAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _HU, start, pos, perr)
fail:
	return _memoize(parser, _HU, start, -1, perr)
}

func _HUNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_HU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _HU}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "HU"}
	// h U &(tone? boundary)
	// h
	if !_node(parser, _hNode, node, &pos) {
		goto fail
	}
	// U
	if !_node(parser, _UNode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _HUFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _HU, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "HU",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _HU}
	// h U &(tone? boundary)
	// h
	if !_fail(parser, _hFail, errPos, failure, &pos) {
		goto fail
	}
	// U
	if !_fail(parser, _UFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _HUAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_HU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _HU}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// h U &(tone? boundary)
	{
		var node0 string
		// h
		if p, n := _hAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// U
		if p, n := _UAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _JUAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _JU, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// j U &(tone? boundary)
	// j
	if !_accept(parser, _jAccepts, &pos, &perr) {
		goto fail
	}
	// U
	if !_accept(parser, _UAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _JU, start, pos, perr)
fail:
	return _memoize(parser, _JU, start, -1, perr)
}

func _JUNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_JU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _JU}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "JU"}
	// j U &(tone? boundary)
	// j
	if !_node(parser, _jNode, node, &pos) {
		goto fail
	}
	// U
	if !_node(parser, _UNode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _JUFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _JU, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "JU",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _JU}
	// j U &(tone? boundary)
	// j
	if !_fail(parser, _jFail, errPos, failure, &pos) {
		goto fail
	}
	// U
	if !_fail(parser, _UFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _JUAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_JU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _JU}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// j U &(tone? boundary)
	{
		var node0 string
		// j
		if p, n := _jAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// U
		if p, n := _UAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _KUAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _KU, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// k U &(tone? boundary)
	// k
	if !_accept(parser, _kAccepts, &pos, &perr) {
		goto fail
	}
	// U
	if !_accept(parser, _UAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _KU, start, pos, perr)
fail:
	return _memoize(parser, _KU, start, -1, perr)
}

func _KUNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_KU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KU}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "KU"}
	// k U &(tone? boundary)
	// k
	if !_node(parser, _kNode, node, &pos) {
		goto fail
	}
	// U
	if !_node(parser, _UNode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _KUFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _KU, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "KU",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _KU}
	// k U &(tone? boundary)
	// k
	if !_fail(parser, _kFail, errPos, failure, &pos) {
		goto fail
	}
	// U
	if !_fail(parser, _UFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _KUAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_KU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KU}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// k U &(tone? boundary)
	{
		var node0 string
		// k
		if p, n := _kAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// U
		if p, n := _UAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _KIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _KI, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// k I &(tone? boundary)
	// k
	if !_accept(parser, _kAccepts, &pos, &perr) {
		goto fail
	}
	// I
	if !_accept(parser, _IAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _KI, start, pos, perr)
fail:
	return _memoize(parser, _KI, start, -1, perr)
}

func _KINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_KI]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KI}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "KI"}
	// k I &(tone? boundary)
	// k
	if !_node(parser, _kNode, node, &pos) {
		goto fail
	}
	// I
	if !_node(parser, _INode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _KIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _KI, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "KI",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _KI}
	// k I &(tone? boundary)
	// k
	if !_fail(parser, _kFail, errPos, failure, &pos) {
		goto fail
	}
	// I
	if !_fail(parser, _IFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _KIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_KI]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KI}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// k I &(tone? boundary)
	{
		var node0 string
		// k
		if p, n := _kAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// I
		if p, n := _IAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _KIOAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _KIO, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// k I o &(tone? boundary)
	// k
	if !_accept(parser, _kAccepts, &pos, &perr) {
		goto fail
	}
	// I
	if !_accept(parser, _IAccepts, &pos, &perr) {
		goto fail
	}
	// o
	if !_accept(parser, _oAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _KIO, start, pos, perr)
fail:
	return _memoize(parser, _KIO, start, -1, perr)
}

func _KIONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_KIO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KIO}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "KIO"}
	// k I o &(tone? boundary)
	// k
	if !_node(parser, _kNode, node, &pos) {
		goto fail
	}
	// I
	if !_node(parser, _INode, node, &pos) {
		goto fail
	}
	// o
	if !_node(parser, _oNode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _KIOFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _KIO, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "KIO",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _KIO}
	// k I o &(tone? boundary)
	// k
	if !_fail(parser, _kFail, errPos, failure, &pos) {
		goto fail
	}
	// I
	if !_fail(parser, _IFail, errPos, failure, &pos) {
		goto fail
	}
	// o
	if !_fail(parser, _oFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _KIOAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_KIO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KIO}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// k I o &(tone? boundary)
	{
		var node0 string
		// k
		if p, n := _kAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// I
		if p, n := _IAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// o
		if p, n := _oAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _KEOAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _KEO, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (k E o/j E) &(tone? boundary)
	// (k E o/j E)
	// k E o/j E
	{
		pos4 := pos
		// k E o
		// k
		if !_accept(parser, _kAccepts, &pos, &perr) {
			goto fail5
		}
		// E
		if !_accept(parser, _EAccepts, &pos, &perr) {
			goto fail5
		}
		// o
		if !_accept(parser, _oAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// j E
		// j
		if !_accept(parser, _jAccepts, &pos, &perr) {
			goto fail7
		}
		// E
		if !_accept(parser, _EAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos10 := pos
		perr12 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos16 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail17
			}
			goto ok18
		fail17:
			pos = pos16
		ok18:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok9
	fail13:
		pos = pos10
		perr = _max(perr12, pos)
		goto fail
	ok9:
		pos = pos10
		perr = perr12
	}
	return _memoize(parser, _KEO, start, pos, perr)
fail:
	return _memoize(parser, _KEO, start, -1, perr)
}

func _KEONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_KEO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KEO}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "KEO"}
	// (k E o/j E) &(tone? boundary)
	// (k E o/j E)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// k E o/j E
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// k E o
			// k
			if !_node(parser, _kNode, node, &pos) {
				goto fail7
			}
			// E
			if !_node(parser, _ENode, node, &pos) {
				goto fail7
			}
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// j E
			// j
			if !_node(parser, _jNode, node, &pos) {
				goto fail9
			}
			// E
			if !_node(parser, _ENode, node, &pos) {
				goto fail9
			}
			goto ok3
		fail9:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			goto fail
		ok3:
		}
		sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
		node.Kids = append(node.Kids[:nkids1], sub)
	}
	// &(tone? boundary)
	{
		pos12 := pos
		nkids13 := len(node.Kids)
		// (tone? boundary)
		{
			nkids16 := len(node.Kids)
			pos017 := pos
			// tone? boundary
			// tone?
			{
				nkids19 := len(node.Kids)
				pos20 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail21
				}
				goto ok22
			fail21:
				node.Kids = node.Kids[:nkids19]
				pos = pos20
			ok22:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail15
			}
			sub := _sub(parser, pos017, pos, node.Kids[nkids16:])
			node.Kids = append(node.Kids[:nkids16], sub)
		}
		goto ok11
	fail15:
		pos = pos12
		goto fail
	ok11:
		pos = pos12
		node.Kids = node.Kids[:nkids13]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _KEOFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _KEO, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "KEO",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _KEO}
	// (k E o/j E) &(tone? boundary)
	// (k E o/j E)
	// k E o/j E
	{
		pos4 := pos
		// k E o
		// k
		if !_fail(parser, _kFail, errPos, failure, &pos) {
			goto fail5
		}
		// E
		if !_fail(parser, _EFail, errPos, failure, &pos) {
			goto fail5
		}
		// o
		if !_fail(parser, _oFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// j E
		// j
		if !_fail(parser, _jFail, errPos, failure, &pos) {
			goto fail7
		}
		// E
		if !_fail(parser, _EFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos10 := pos
		nkids11 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos16 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail17
			}
			goto ok18
		fail17:
			pos = pos16
		ok18:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok9
	fail13:
		pos = pos10
		failure.Kids = failure.Kids[:nkids11]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok9:
		pos = pos10
		failure.Kids = failure.Kids[:nkids11]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _KEOAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_KEO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _KEO}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (k E o/j E) &(tone? boundary)
	{
		var node0 string
		// (k E o/j E)
		// k E o/j E
		{
			pos4 := pos
			var node3 string
			// k E o
			{
				var node6 string
				// k
				if p, n := _kAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// E
				if p, n := _EAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
			}
			goto ok1
		fail5:
			node0 = node3
			pos = pos4
			// j E
			{
				var node8 string
				// j
				if p, n := _jAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// E
				if p, n := _EAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
			}
			goto ok1
		fail7:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos10 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos16 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail17
				} else {
					pos = p
				}
				goto ok18
			fail17:
				pos = pos16
			ok18:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail13
			} else {
				pos = p
			}
			goto ok9
		fail13:
			pos = pos10
			goto fail
		ok9:
			pos = pos10
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _LUAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _LU, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (l U/l I/l O/m A/t I o) &(tone? boundary)
	// (l U/l I/l O/m A/t I o)
	// l U/l I/l O/m A/t I o
	{
		pos4 := pos
		// l U
		// l
		if !_accept(parser, _lAccepts, &pos, &perr) {
			goto fail5
		}
		// U
		if !_accept(parser, _UAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// l I
		// l
		if !_accept(parser, _lAccepts, &pos, &perr) {
			goto fail7
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// l O
		// l
		if !_accept(parser, _lAccepts, &pos, &perr) {
			goto fail9
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// m A
		// m
		if !_accept(parser, _mAccepts, &pos, &perr) {
			goto fail11
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// t I o
		// t
		if !_accept(parser, _tAccepts, &pos, &perr) {
			goto fail13
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail13
		}
		// o
		if !_accept(parser, _oAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos16 := pos
		perr18 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos22 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail23
			}
			goto ok24
		fail23:
			pos = pos22
		ok24:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail19
		}
		goto ok15
	fail19:
		pos = pos16
		perr = _max(perr18, pos)
		goto fail
	ok15:
		pos = pos16
		perr = perr18
	}
	return _memoize(parser, _LU, start, pos, perr)
fail:
	return _memoize(parser, _LU, start, -1, perr)
}

func _LUNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_LU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "LU"}
	// (l U/l I/l O/m A/t I o) &(tone? boundary)
	// (l U/l I/l O/m A/t I o)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// l U/l I/l O/m A/t I o
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// l U
			// l
			if !_node(parser, _lNode, node, &pos) {
				goto fail7
			}
			// U
			if !_node(parser, _UNode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// l I
			// l
			if !_node(parser, _lNode, node, &pos) {
				goto fail9
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail9
			}
			goto ok3
		fail9:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// l O
			// l
			if !_node(parser, _lNode, node, &pos) {
				goto fail11
			}
			// O
			if !_node(parser, _ONode, node, &pos) {
				goto fail11
			}
			goto ok3
		fail11:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// m A
			// m
			if !_node(parser, _mNode, node, &pos) {
				goto fail13
			}
			// A
			if !_node(parser, _ANode, node, &pos) {
				goto fail13
			}
			goto ok3
		fail13:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// t I o
			// t
			if !_node(parser, _tNode, node, &pos) {
				goto fail15
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail15
			}
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail15
			}
			goto ok3
		fail15:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			goto fail
		ok3:
		}
		sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
		node.Kids = append(node.Kids[:nkids1], sub)
	}
	// &(tone? boundary)
	{
		pos18 := pos
		nkids19 := len(node.Kids)
		// (tone? boundary)
		{
			nkids22 := len(node.Kids)
			pos023 := pos
			// tone? boundary
			// tone?
			{
				nkids25 := len(node.Kids)
				pos26 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail27
				}
				goto ok28
			fail27:
				node.Kids = node.Kids[:nkids25]
				pos = pos26
			ok28:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail21
			}
			sub := _sub(parser, pos023, pos, node.Kids[nkids22:])
			node.Kids = append(node.Kids[:nkids22], sub)
		}
		goto ok17
	fail21:
		pos = pos18
		goto fail
	ok17:
		pos = pos18
		node.Kids = node.Kids[:nkids19]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _LUFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _LU, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "LU",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _LU}
	// (l U/l I/l O/m A/t I o) &(tone? boundary)
	// (l U/l I/l O/m A/t I o)
	// l U/l I/l O/m A/t I o
	{
		pos4 := pos
		// l U
		// l
		if !_fail(parser, _lFail, errPos, failure, &pos) {
			goto fail5
		}
		// U
		if !_fail(parser, _UFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// l I
		// l
		if !_fail(parser, _lFail, errPos, failure, &pos) {
			goto fail7
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// l O
		// l
		if !_fail(parser, _lFail, errPos, failure, &pos) {
			goto fail9
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// m A
		// m
		if !_fail(parser, _mFail, errPos, failure, &pos) {
			goto fail11
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// t I o
		// t
		if !_fail(parser, _tFail, errPos, failure, &pos) {
			goto fail13
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail13
		}
		// o
		if !_fail(parser, _oFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos16 := pos
		nkids17 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos22 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail23
			}
			goto ok24
		fail23:
			pos = pos22
		ok24:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail19
		}
		goto ok15
	fail19:
		pos = pos16
		failure.Kids = failure.Kids[:nkids17]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok15:
		pos = pos16
		failure.Kids = failure.Kids[:nkids17]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _LUAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_LU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _LU}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (l U/l I/l O/m A/t I o) &(tone? boundary)
	{
		var node0 string
		// (l U/l I/l O/m A/t I o)
		// l U/l I/l O/m A/t I o
		{
			pos4 := pos
			var node3 string
			// l U
			{
				var node6 string
				// l
				if p, n := _lAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// U
				if p, n := _UAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
			}
			goto ok1
		fail5:
			node0 = node3
			pos = pos4
			// l I
			{
				var node8 string
				// l
				if p, n := _lAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// I
				if p, n := _IAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
			}
			goto ok1
		fail7:
			node0 = node3
			pos = pos4
			// l O
			{
				var node10 string
				// l
				if p, n := _lAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
				// O
				if p, n := _OAction(parser, pos); n == nil {
					goto fail9
				} else {
					node10 = *n
					pos = p
				}
				node0, node10 = node0+node10, ""
			}
			goto ok1
		fail9:
			node0 = node3
			pos = pos4
			// m A
			{
				var node12 string
				// m
				if p, n := _mAction(parser, pos); n == nil {
					goto fail11
				} else {
					node12 = *n
					pos = p
				}
				node0, node12 = node0+node12, ""
				// A
				if p, n := _AAction(parser, pos); n == nil {
					goto fail11
				} else {
					node12 = *n
					pos = p
				}
				node0, node12 = node0+node12, ""
			}
			goto ok1
		fail11:
			node0 = node3
			pos = pos4
			// t I o
			{
				var node14 string
				// t
				if p, n := _tAction(parser, pos); n == nil {
					goto fail13
				} else {
					node14 = *n
					pos = p
				}
				node0, node14 = node0+node14, ""
				// I
				if p, n := _IAction(parser, pos); n == nil {
					goto fail13
				} else {
					node14 = *n
					pos = p
				}
				node0, node14 = node0+node14, ""
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail13
				} else {
					node14 = *n
					pos = p
				}
				node0, node14 = node0+node14, ""
			}
			goto ok1
		fail13:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos16 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos22 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail23
				} else {
					pos = p
				}
				goto ok24
			fail23:
				pos = pos22
			ok24:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail19
			} else {
				pos = p
			}
			goto ok15
		fail19:
			pos = pos16
			goto fail
		ok15:
			pos = pos16
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MUAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MU, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// m U &(tone? boundary)
	// m
	if !_accept(parser, _mAccepts, &pos, &perr) {
		goto fail
	}
	// U
	if !_accept(parser, _UAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _MU, start, pos, perr)
fail:
	return _memoize(parser, _MU, start, -1, perr)
}

func _MUNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MU}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MU"}
	// m U &(tone? boundary)
	// m
	if !_node(parser, _mNode, node, &pos) {
		goto fail
	}
	// U
	if !_node(parser, _UNode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MUFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MU, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MU",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MU}
	// m U &(tone? boundary)
	// m
	if !_fail(parser, _mFail, errPos, failure, &pos) {
		goto fail
	}
	// U
	if !_fail(parser, _UFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MUAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MU}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// m U &(tone? boundary)
	{
		var node0 string
		// m
		if p, n := _mAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// U
		if p, n := _UAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MI, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (m I/s h U) &(tone? boundary)
	// (m I/s h U)
	// m I/s h U
	{
		pos4 := pos
		// m I
		// m
		if !_accept(parser, _mAccepts, &pos, &perr) {
			goto fail5
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// s h U
		// s
		if !_accept(parser, _sAccepts, &pos, &perr) {
			goto fail7
		}
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail7
		}
		// U
		if !_accept(parser, _UAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos10 := pos
		perr12 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos16 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail17
			}
			goto ok18
		fail17:
			pos = pos16
		ok18:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok9
	fail13:
		pos = pos10
		perr = _max(perr12, pos)
		goto fail
	ok9:
		pos = pos10
		perr = perr12
	}
	return _memoize(parser, _MI, start, pos, perr)
fail:
	return _memoize(parser, _MI, start, -1, perr)
}

func _MINode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MI]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI"}
	// (m I/s h U) &(tone? boundary)
	// (m I/s h U)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// m I/s h U
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// m I
			// m
			if !_node(parser, _mNode, node, &pos) {
				goto fail7
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// s h U
			// s
			if !_node(parser, _sNode, node, &pos) {
				goto fail9
			}
			// h
			if !_node(parser, _hNode, node, &pos) {
				goto fail9
			}
			// U
			if !_node(parser, _UNode, node, &pos) {
				goto fail9
			}
			goto ok3
		fail9:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			goto fail
		ok3:
		}
		sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
		node.Kids = append(node.Kids[:nkids1], sub)
	}
	// &(tone? boundary)
	{
		pos12 := pos
		nkids13 := len(node.Kids)
		// (tone? boundary)
		{
			nkids16 := len(node.Kids)
			pos017 := pos
			// tone? boundary
			// tone?
			{
				nkids19 := len(node.Kids)
				pos20 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail21
				}
				goto ok22
			fail21:
				node.Kids = node.Kids[:nkids19]
				pos = pos20
			ok22:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail15
			}
			sub := _sub(parser, pos017, pos, node.Kids[nkids16:])
			node.Kids = append(node.Kids[:nkids16], sub)
		}
		goto ok11
	fail15:
		pos = pos12
		goto fail
	ok11:
		pos = pos12
		node.Kids = node.Kids[:nkids13]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MI, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI}
	// (m I/s h U) &(tone? boundary)
	// (m I/s h U)
	// m I/s h U
	{
		pos4 := pos
		// m I
		// m
		if !_fail(parser, _mFail, errPos, failure, &pos) {
			goto fail5
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// s h U
		// s
		if !_fail(parser, _sFail, errPos, failure, &pos) {
			goto fail7
		}
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail7
		}
		// U
		if !_fail(parser, _UFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos10 := pos
		nkids11 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos16 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail17
			}
			goto ok18
		fail17:
			pos = pos16
		ok18:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok9
	fail13:
		pos = pos10
		failure.Kids = failure.Kids[:nkids11]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok9:
		pos = pos10
		failure.Kids = failure.Kids[:nkids11]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MIAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MI]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (m I/s h U) &(tone? boundary)
	{
		var node0 string
		// (m I/s h U)
		// m I/s h U
		{
			pos4 := pos
			var node3 string
			// m I
			{
				var node6 string
				// m
				if p, n := _mAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
				// I
				if p, n := _IAction(parser, pos); n == nil {
					goto fail5
				} else {
					node6 = *n
					pos = p
				}
				node0, node6 = node0+node6, ""
			}
			goto ok1
		fail5:
			node0 = node3
			pos = pos4
			// s h U
			{
				var node8 string
				// s
				if p, n := _sAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// h
				if p, n := _hAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
				// U
				if p, n := _UAction(parser, pos); n == nil {
					goto fail7
				} else {
					node8 = *n
					pos = p
				}
				node0, node8 = node0+node8, ""
			}
			goto ok1
		fail7:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos10 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos16 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail17
				} else {
					pos = p
				}
				goto ok18
			fail17:
				pos = pos16
			ok18:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail13
			} else {
				pos = p
			}
			goto ok9
		fail13:
			pos = pos10
			goto fail
		ok9:
			pos = pos10
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MOAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MO, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// m O &(tone? boundary)
	// m
	if !_accept(parser, _mAccepts, &pos, &perr) {
		goto fail
	}
	// O
	if !_accept(parser, _OAccepts, &pos, &perr) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		perr4 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	return _memoize(parser, _MO, start, pos, perr)
fail:
	return _memoize(parser, _MO, start, -1, perr)
}

func _MONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MO"}
	// m O &(tone? boundary)
	// m
	if !_node(parser, _mNode, node, &pos) {
		goto fail
	}
	// O
	if !_node(parser, _ONode, node, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (tone? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// tone? boundary
			// tone?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail11
				}
				goto ok12
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok12:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail5
			}
			sub := _sub(parser, pos07, pos, node.Kids[nkids6:])
			node.Kids = append(node.Kids[:nkids6], sub)
		}
		goto ok1
	fail5:
		pos = pos2
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _MOFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MO, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MO",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MO}
	// m O &(tone? boundary)
	// m
	if !_fail(parser, _mFail, errPos, failure, &pos) {
		goto fail
	}
	// O
	if !_fail(parser, _OFail, errPos, failure, &pos) {
		goto fail
	}
	// &(tone? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos8 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _MOAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MO}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// m O &(tone? boundary)
	{
		var node0 string
		// m
		if p, n := _mAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// O
		if p, n := _OAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos2 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos8 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail9
				} else {
					pos = p
				}
				goto ok10
			fail9:
				pos = pos8
			ok10:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
	}
	return pos, &node
}