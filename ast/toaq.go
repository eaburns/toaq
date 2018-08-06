package ast

import (
	"unicode"

	"github.com/eaburns/peggy/peg"
	"github.com/eaburns/toaq/tone"
)

const (
	_full_text                                                        int = 0
	_text                                                             int = 1
	_discourse                                                        int = 2
	_fragment                                                         int = 3
	_sentence                                                         int = 4
	_sentence_1                                                       int = 5
	_sentence_2                                                       int = 6
	_sentence_prefix_spaces                                           int = 7
	_sentence_3                                                       int = 8
	_statement                                                        int = 9
	_statement_1                                                      int = 10
	_statement_2                                                      int = 11
	_statement_3                                                      int = 12
	_predication                                                      int = 13
	_prenex_spaces                                                    int = 14
	_prenex                                                           int = 15
	_predicate                                                        int = 16
	_predicate_1                                                      int = 17
	_predicate_2                                                      int = 18
	_prefixed_predicate                                               int = 19
	_predicate_3                                                      int = 20
	_serial_predicate                                                 int = 21
	_terms                                                            int = 22
	_terms_2                                                          int = 23
	_term                                                             int = 24
	_linked_argument                                                  int = 25
	_linking_word_spaces                                              int = 26
	_termset                                                          int = 27
	_terms_II                                                         int = 28
	_terms_III                                                        int = 29
	_terms_IV                                                         int = 30
	_terms_V                                                          int = 31
	_argument                                                         int = 32
	_arg_1                                                            int = 33
	_arg_2                                                            int = 34
	_focus_spaces                                                     int = 35
	_arg_3                                                            int = 36
	_quantifier_spaces                                                int = 37
	_arg_4                                                            int = 38
	_arg_5                                                            int = 39
	_arg_6                                                            int = 40
	_arg_7                                                            int = 41
	_relative_clause                                                  int = 42
	_relative_clause_1                                                int = 43
	_relative_clause_2                                                int = 44
	_relative_clause_3                                                int = 45
	_relative_predication                                             int = 46
	_relative_predicate                                               int = 47
	_relative_predicate_1                                             int = 48
	_relative_predicate_2                                             int = 49
	_adverb                                                           int = 50
	_adverb_1                                                         int = 51
	_adverb_2                                                         int = 52
	_adverb_3                                                         int = 53
	_adverb_4                                                         int = 54
	_prepositional_phrase                                             int = 55
	_prepositional_phrase_1                                           int = 56
	_prepositional_phrase_2                                           int = 57
	_preposition                                                      int = 58
	_preposition_1                                                    int = 59
	_preposition_2                                                    int = 60
	_preposition_3                                                    int = 61
	_preposition_4                                                    int = 62
	_content_clause                                                   int = 63
	_content_clause_1                                                 int = 64
	_content_statement                                                int = 65
	_content_predication                                              int = 66
	_content_predicate                                                int = 67
	_content_predicate_1                                              int = 68
	_content_predicate_2                                              int = 69
	_forethought_connective                                           int = 70
	_MI_complement                                                    int = 71
	__                                                                int = 72
	_space_or_freemod                                                 int = 73
	_space_word                                                       int = 74
	_freemod                                                          int = 75
	_parenthetical                                                    int = 76
	_parenthetical_1                                                  int = 77
	_incidental                                                       int = 78
	_vocative                                                         int = 79
	_word                                                             int = 80
	_non_function_word                                                int = 81
	_neutral_function_word                                            int = 82
	_toned_function_word                                              int = 83
	_initial_syllable                                                 int = 84
	_prefix                                                           int = 85
	_focus                                                            int = 86
	_end_quote                                                        int = 87
	_end_predicatizer                                                 int = 88
	_end_statement                                                    int = 89
	_sentence_prefix                                                  int = 90
	_end_prenex                                                       int = 91
	_start_incidental                                                 int = 92
	_start_parenthetical                                              int = 93
	_end_parenthetical                                                int = 94
	_vocative_marker                                                  int = 95
	_linking_word                                                     int = 96
	_connective                                                       int = 97
	_illocutionary                                                    int = 98
	_moq                                                              int = 99
	_quantifier                                                       int = 100
	_interjection                                                     int = 101
	_forethought_marker                                               int = 102
	_function_word                                                    int = 103
	_BI                                                               int = 104
	_DA                                                               int = 105
	_MOQ                                                              int = 106
	_GA                                                               int = 107
	_GO                                                               int = 108
	_HA                                                               int = 109
	_HU                                                               int = 110
	_JU                                                               int = 111
	_KU                                                               int = 112
	_KI                                                               int = 113
	_KIO                                                              int = 114
	_KEO                                                              int = 115
	_LU                                                               int = 116
	_MU                                                               int = 117
	_MI                                                               int = 118
	_MO                                                               int = 119
	_NA                                                               int = 120
	_PO                                                               int = 121
	_RA                                                               int = 122
	_TU                                                               int = 123
	_TO                                                               int = 124
	_TEO                                                              int = 125
	_neutral_syllable                                                 int = 126
	_moq_syllable                                                     int = 127
	_compound_syllable                                                int = 128
	_arg_syllable                                                     int = 129
	_relative_syllable                                                int = 130
	_verb_syllable                                                    int = 131
	_content_syllable                                                 int = 132
	_preposition_syllable                                             int = 133
	_adverb_syllable                                                  int = 134
	_boundary                                                         int = 135
	_initial                                                          int = 136
	_neutral_desinence                                                int = 137
	_compound_desinence                                               int = 138
	_arg_desinence                                                    int = 139
	_relative_desinence                                               int = 140
	_verb_desinence                                                   int = 141
	_content_desinence                                                int = 142
	_preposition_desinence                                            int = 143
	_adverb_desinence                                                 int = 144
	_tone                                                             int = 145
	_A                                                                int = 146
	_U                                                                int = 147
	_I                                                                int = 148
	_O                                                                int = 149
	_E                                                                int = 150
	_ā                                                                int = 151
	_ū                                                                int = 152
	_ī                                                                int = 153
	_ō                                                                int = 154
	_ē                                                                int = 155
	_macron_combiner                                                  int = 156
	_compound_tone                                                    int = 157
	_á                                                                int = 158
	_ú                                                                int = 159
	_í                                                                int = 160
	_ó                                                                int = 161
	_é                                                                int = 162
	_acute_combiner                                                   int = 163
	_arg_tone                                                         int = 164
	_ǎ                                                                int = 165
	_ǔ                                                                int = 166
	_ǐ                                                                int = 167
	_ǒ                                                                int = 168
	_ě                                                                int = 169
	_caron_combiner                                                   int = 170
	_breve_combiner                                                   int = 171
	_relative_tone                                                    int = 172
	_ả                                                                int = 173
	_ủ                                                                int = 174
	_ỉ                                                                int = 175
	_ỏ                                                                int = 176
	_ẻ                                                                int = 177
	_hook_combiner                                                    int = 178
	_verb_tone                                                        int = 179
	_â                                                                int = 180
	_û                                                                int = 181
	_î                                                                int = 182
	_ô                                                                int = 183
	_ê                                                                int = 184
	_circumflex_combiner                                              int = 185
	_content_tone                                                     int = 186
	_à                                                                int = 187
	_ù                                                                int = 188
	_ì                                                                int = 189
	_ò                                                                int = 190
	_è                                                                int = 191
	_grave_combiner                                                   int = 192
	_preposition_tone                                                 int = 193
	_ã                                                                int = 194
	_ũ                                                                int = 195
	_ĩ                                                                int = 196
	_õ                                                                int = 197
	_ẽ                                                                int = 198
	_tilde_combiner                                                   int = 199
	_adverb_tone                                                      int = 200
	_a                                                                int = 201
	_b                                                                int = 202
	_c                                                                int = 203
	_d                                                                int = 204
	_e                                                                int = 205
	_f                                                                int = 206
	_g                                                                int = 207
	_h                                                                int = 208
	_i                                                                int = 209
	_j                                                                int = 210
	_k                                                                int = 211
	_l                                                                int = 212
	_m                                                                int = 213
	_n                                                                int = 214
	_o                                                                int = 215
	_p                                                                int = 216
	_q                                                                int = 217
	_r                                                                int = 218
	_s                                                                int = 219
	_t                                                                int = 220
	_u                                                                int = 221
	_w                                                                int = 222
	_y                                                                int = 223
	_space                                                            int = 224
	_punctuation                                                      int = 225
	_questions                                                        int = 226
	_EOF                                                              int = 227
	_afterthought_cop__relative_clause_1__relative_clause_1           int = 228
	_afterthought_cop__sentence_1__sentence                           int = 229
	_forethought_cop__sentence__sentence                              int = 230
	_afterthought_cop__statement_2__statement                         int = 231
	_forethought_cop__statement__statement                            int = 232
	_afterthought_cop__predicate_2__predicate_1                       int = 233
	_forethought_cop_pred__predicate                                  int = 234
	_LU_predicate__verb_syllable                                      int = 235
	_MI_phrase__verb_syllable                                         int = 236
	_PO_phrase__verb_syllable                                         int = 237
	_MO_phrase__verb_syllable                                         int = 238
	_predicate_word__verb_syllable                                    int = 239
	_serial__predicate_1                                              int = 240
	_forethought_cop__terms_V__terms_V                                int = 241
	_forethought_cop__terms_IV__terms_IV                              int = 242
	_forethought_cop__terms_III__terms_III                            int = 243
	_forethought_cop__terms_II__terms_II                              int = 244
	_afterthought_cop__arg_1__argument                                int = 245
	_forethought_cop__argument__argument                              int = 246
	_serial__arg_6                                                    int = 247
	_afterthought_cop_pred__arg_7                                     int = 248
	_forethought_cop_pred__arg_5                                      int = 249
	_LU_predicate__arg_syllable                                       int = 250
	_MI_phrase__arg_syllable                                          int = 251
	_PO_phrase__arg_syllable                                          int = 252
	_MO_phrase__arg_syllable                                          int = 253
	_predicate_word__arg_syllable                                     int = 254
	_LU_phrase__relative_syllable                                     int = 255
	_forethought_cop__relative_clause__relative_clause                int = 256
	_afterthought_cop__relative_clause_3__statement                   int = 257
	_serial__relative_predicate_1                                     int = 258
	_afterthought_cop_pred__relative_predicate_2                      int = 259
	_forethought_cop_pred__relative_predicate                         int = 260
	_MI_phrase__relative_syllable                                     int = 261
	_PO_phrase__relative_syllable                                     int = 262
	_MO_phrase__relative_syllable                                     int = 263
	_predicate_word__relative_syllable                                int = 264
	_afterthought_cop__adverb_1__adverb                               int = 265
	_forethought_cop__adverb__adverb                                  int = 266
	_serial__adverb_3                                                 int = 267
	_afterthought_cop_pred__adverb_4                                  int = 268
	_forethought_cop_pred__adverb_3                                   int = 269
	_LU_predicate__adverb_syllable                                    int = 270
	_MI_phrase__adverb_syllable                                       int = 271
	_PO_phrase__adverb_syllable                                       int = 272
	_MO_phrase__adverb_syllable                                       int = 273
	_predicate_word__adverb_syllable                                  int = 274
	_afterthought_cop__prepositional_phrase_1__prepositional_phrase_1 int = 275
	_forethought_cop__prepositional_phrase__prepositional_phrase      int = 276
	_afterthought_cop__preposition_1__preposition                     int = 277
	_forethought_cop__preposition__preposition                        int = 278
	_serial__preposition_3                                            int = 279
	_afterthought_cop_pred__preposition_4                             int = 280
	_forethought_cop_pred__preposition                                int = 281
	_LU_predicate__preposition_syllable                               int = 282
	_MI_phrase__preposition_syllable                                  int = 283
	_PO_phrase__preposition_syllable                                  int = 284
	_MO_phrase__preposition_syllable                                  int = 285
	_predicate_word__preposition_syllable                             int = 286
	_afterthought_cop__content_clause_1__statement                    int = 287
	_LU_phrase__content_syllable                                      int = 288
	_serial__content_predicate_1                                      int = 289
	_afterthought_cop_pred__content_predicate_2                       int = 290
	_forethought_cop_pred__content_predicate_2                        int = 291
	_MI_phrase__content_syllable                                      int = 292
	_PO_phrase__content_syllable                                      int = 293
	_MO_phrase__content_syllable                                      int = 294
	_predicate_word__content_syllable                                 int = 295
	_syllable__compound_desinence__compound_tone                      int = 296
	_syllable__arg_desinence__arg_tone                                int = 297
	_syllable__relative_desinence__relative_tone                      int = 298
	_syllable__verb_desinence__verb_tone                              int = 299
	_syllable__content_desinence__content_tone                        int = 300
	_syllable__preposition_desinence__preposition_tone                int = 301
	_syllable__adverb_desinence__adverb_tone                          int = 302
	_desinence__a__u__i__o__e                                         int = 303
	_desinence__ā__ū__ī__ō__ē                                         int = 304
	_desinence__á__ú__í__ó__é                                         int = 305
	_desinence__ǎ__ǔ__ǐ__ǒ__ě                                         int = 306
	_desinence__ả__ủ__ỉ__ỏ__ẻ                                         int = 307
	_desinence__â__û__î__ô__ê                                         int = 308
	_desinence__à__ù__ì__ò__è                                         int = 309
	_desinence__ã__ũ__ĩ__õ__ẽ                                         int = 310
	_cop_bar__relative_clause_1                                       int = 311
	_cop_bar__sentence                                                int = 312
	_forethought_cop_1__sentence__sentence                            int = 313
	_cop_bar__statement                                               int = 314
	_forethought_cop_1__statement__statement                          int = 315
	_cop_bar__predicate_1                                             int = 316
	_forethought_cop__predicate__predicate                            int = 317
	_LU_phrase__verb_syllable                                         int = 318
	_MI_phrase_1__verb_syllable                                       int = 319
	_PO_phrase_1__verb_syllable                                       int = 320
	_MO_phrase_1__verb_syllable                                       int = 321
	_forethought_cop_1__terms_V__terms_V                              int = 322
	_forethought_cop_1__terms_IV__terms_IV                            int = 323
	_forethought_cop_1__terms_III__terms_III                          int = 324
	_forethought_cop_1__terms_II__terms_II                            int = 325
	_cop_bar__argument                                                int = 326
	_forethought_cop_1__argument__argument                            int = 327
	_afterthought_cop__arg_7__predicate                               int = 328
	_forethought_cop__arg_5__predicate                                int = 329
	_LU_phrase__arg_syllable                                          int = 330
	_MI_phrase_1__arg_syllable                                        int = 331
	_PO_phrase_1__arg_syllable                                        int = 332
	_MO_phrase_1__arg_syllable                                        int = 333
	_LU_word__relative_syllable                                       int = 334
	_forethought_cop_1__relative_clause__relative_clause              int = 335
	_afterthought_cop__relative_predicate_2__predicate                int = 336
	_forethought_cop__relative_predicate__predicate                   int = 337
	_MI_phrase_1__relative_syllable                                   int = 338
	_PO_phrase_1__relative_syllable                                   int = 339
	_MO_phrase_1__relative_syllable                                   int = 340
	_cop_bar__adverb                                                  int = 341
	_forethought_cop_1__adverb__adverb                                int = 342
	_afterthought_cop__adverb_4__predicate                            int = 343
	_forethought_cop__adverb_3__predicate                             int = 344
	_LU_phrase__adverb_syllable                                       int = 345
	_MI_phrase_1__adverb_syllable                                     int = 346
	_PO_phrase_1__adverb_syllable                                     int = 347
	_MO_phrase_1__adverb_syllable                                     int = 348
	_cop_bar__prepositional_phrase_1                                  int = 349
	_forethought_cop_1__prepositional_phrase__prepositional_phrase    int = 350
	_cop_bar__preposition                                             int = 351
	_forethought_cop_1__preposition__preposition                      int = 352
	_afterthought_cop__preposition_4__predicate                       int = 353
	_forethought_cop__preposition__predicate                          int = 354
	_LU_phrase__preposition_syllable                                  int = 355
	_MI_phrase_1__preposition_syllable                                int = 356
	_PO_phrase_1__preposition_syllable                                int = 357
	_MO_phrase_1__preposition_syllable                                int = 358
	_LU_word__content_syllable                                        int = 359
	_afterthought_cop__content_predicate_2__predicate                 int = 360
	_forethought_cop__content_predicate_2__predicate                  int = 361
	_MI_phrase_1__content_syllable                                    int = 362
	_PO_phrase_1__content_syllable                                    int = 363
	_MO_phrase_1__content_syllable                                    int = 364
	_forethought_cop_bar__sentence                                    int = 365
	_forethought_cop_bar__statement                                   int = 366
	_forethought_cop_1__predicate__predicate                          int = 367
	_LU_word__verb_syllable                                           int = 368
	_MI_word__verb_syllable                                           int = 369
	_PO_word__verb_syllable                                           int = 370
	_MO_word__verb_syllable                                           int = 371
	_forethought_cop_bar__terms_V                                     int = 372
	_forethought_cop_bar__terms_IV                                    int = 373
	_forethought_cop_bar__terms_III                                   int = 374
	_forethought_cop_bar__terms_II                                    int = 375
	_forethought_cop_bar__argument                                    int = 376
	_cop_bar__predicate                                               int = 377
	_forethought_cop_1__arg_5__predicate                              int = 378
	_LU_word__arg_syllable                                            int = 379
	_MI_word__arg_syllable                                            int = 380
	_PO_word__arg_syllable                                            int = 381
	_MO_word__arg_syllable                                            int = 382
	_forethought_cop_bar__relative_clause                             int = 383
	_forethought_cop_1__relative_predicate__predicate                 int = 384
	_MI_word__relative_syllable                                       int = 385
	_PO_word__relative_syllable                                       int = 386
	_MO_word__relative_syllable                                       int = 387
	_forethought_cop_bar__adverb                                      int = 388
	_forethought_cop_1__adverb_3__predicate                           int = 389
	_LU_word__adverb_syllable                                         int = 390
	_MI_word__adverb_syllable                                         int = 391
	_PO_word__adverb_syllable                                         int = 392
	_MO_word__adverb_syllable                                         int = 393
	_forethought_cop_bar__prepositional_phrase                        int = 394
	_forethought_cop_bar__preposition                                 int = 395
	_forethought_cop_1__preposition__predicate                        int = 396
	_LU_word__preposition_syllable                                    int = 397
	_MI_word__preposition_syllable                                    int = 398
	_PO_word__preposition_syllable                                    int = 399
	_MO_word__preposition_syllable                                    int = 400
	_forethought_cop_1__content_predicate_2__predicate                int = 401
	_MI_word__content_syllable                                        int = 402
	_PO_word__content_syllable                                        int = 403
	_MO_word__content_syllable                                        int = 404
	_forethought_cop_bar__predicate                                   int = 405

	_N int = 406
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
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _full_text, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// t:text EOF
	// t:text
	{
		pos1 := pos
		// text
		if !_accept(parser, _textAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
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
	var labels [1]string
	use(labels)
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
	// action
	// t:text EOF
	// t:text
	{
		pos1 := pos
		// text
		if !_node(parser, _textNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
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
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _full_text, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "full_text",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _full_text}
	// action
	// t:text EOF
	// t:text
	{
		pos1 := pos
		// text
		if !_fail(parser, _textFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
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

func _full_textAction(parser *_Parser, start int) (int, *Text) {
	var labels [1]string
	use(labels)
	var label0 Text
	dp := parser.deltaPos[start][_full_text]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _full_text}
	n := parser.act[key]
	if n != nil {
		n := n.(Text)
		return start + int(dp-1), &n
	}
	var node Text
	pos := start
	// action
	{
		start0 := pos
		// t:text EOF
		// t:text
		{
			pos2 := pos
			// text
			if p, n := _textAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// EOF
		if p, n := _EOFAction(parser, pos); n == nil {
			goto fail
		} else {
			pos = p
		}
		node = func(
			start, end int, t Text) Text {
			return Text(t)
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _textAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _text, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// s:space_word? d:discourse? EOF?
	// s:space_word?
	{
		pos1 := pos
		// space_word?
		{
			pos3 := pos
			// space_word
			if !_accept(parser, _space_wordAccepts, &pos, &perr) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// d:discourse?
	{
		pos6 := pos
		// discourse?
		{
			pos8 := pos
			// discourse
			if !_accept(parser, _discourseAccepts, &pos, &perr) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		labels[1] = parser.text[pos6:pos]
	}
	// EOF?
	{
		pos12 := pos
		// EOF
		if !_accept(parser, _EOFAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok14
	fail13:
		pos = pos12
	ok14:
	}
	return _memoize(parser, _text, start, pos, perr)
}

func _textNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// action
	// s:space_word? d:discourse? EOF?
	// s:space_word?
	{
		pos1 := pos
		// space_word?
		{
			nkids2 := len(node.Kids)
			pos3 := pos
			// space_word
			if !_node(parser, _space_wordNode, node, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			node.Kids = node.Kids[:nkids2]
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// d:discourse?
	{
		pos6 := pos
		// discourse?
		{
			nkids7 := len(node.Kids)
			pos8 := pos
			// discourse
			if !_node(parser, _discourseNode, node, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			node.Kids = node.Kids[:nkids7]
			pos = pos8
		ok10:
		}
		labels[1] = parser.text[pos6:pos]
	}
	// EOF?
	{
		nkids11 := len(node.Kids)
		pos12 := pos
		// EOF
		if !_node(parser, _EOFNode, node, &pos) {
			goto fail13
		}
		goto ok14
	fail13:
		node.Kids = node.Kids[:nkids11]
		pos = pos12
	ok14:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
}

func _textFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _text, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "text",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _text}
	// action
	// s:space_word? d:discourse? EOF?
	// s:space_word?
	{
		pos1 := pos
		// space_word?
		{
			pos3 := pos
			// space_word
			if !_fail(parser, _space_wordFail, errPos, failure, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// d:discourse?
	{
		pos6 := pos
		// discourse?
		{
			pos8 := pos
			// discourse
			if !_fail(parser, _discourseFail, errPos, failure, &pos) {
				goto fail9
			}
			goto ok10
		fail9:
			pos = pos8
		ok10:
		}
		labels[1] = parser.text[pos6:pos]
	}
	// EOF?
	{
		pos12 := pos
		// EOF
		if !_fail(parser, _EOFFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok14
	fail13:
		pos = pos12
	ok14:
	}
	parser.fail[key] = failure
	return pos, failure
}

func _textAction(parser *_Parser, start int) (int, *Text) {
	var labels [2]string
	use(labels)
	var label0 *Mod
	var label1 *[]Node
	dp := parser.deltaPos[start][_text]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _text}
	n := parser.act[key]
	if n != nil {
		n := n.(Text)
		return start + int(dp-1), &n
	}
	var node Text
	pos := start
	// action
	{
		start0 := pos
		// s:space_word? d:discourse? EOF?
		// s:space_word?
		{
			pos2 := pos
			// space_word?
			{
				pos4 := pos
				label0 = new(Mod)
				// space_word
				if p, n := _space_wordAction(parser, pos); n == nil {
					goto fail5
				} else {
					*label0 = *n
					pos = p
				}
				goto ok6
			fail5:
				label0 = nil
				pos = pos4
			ok6:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// d:discourse?
		{
			pos7 := pos
			// discourse?
			{
				pos9 := pos
				label1 = new([]Node)
				// discourse
				if p, n := _discourseAction(parser, pos); n == nil {
					goto fail10
				} else {
					*label1 = *n
					pos = p
				}
				goto ok11
			fail10:
				label1 = nil
				pos = pos9
			ok11:
			}
			labels[1] = parser.text[pos7:pos]
		}
		// EOF?
		{
			pos13 := pos
			// EOF
			if p, n := _EOFAction(parser, pos); n == nil {
				goto fail14
			} else {
				pos = p
			}
			goto ok15
		fail14:
			pos = pos13
		ok15:
		}
		node = func(
			start, end int, d *[]Node, s *Mod) Text {
			return Text(text(s, d))
		}(
			start0, pos, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
}

func _discourseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [4]string
	use(labels)
	if dp, de, ok := _memo(parser, _discourse, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})+
	// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
	// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
	{
		pos7 := pos
		// action
		// x0:sentence s0:_
		// x0:sentence
		{
			pos10 := pos
			// sentence
			if !_accept(parser, _sentenceAccepts, &pos, &perr) {
				goto fail8
			}
			labels[0] = parser.text[pos10:pos]
		}
		// s0:_
		{
			pos11 := pos
			// _
			if !_accept(parser, __Accepts, &pos, &perr) {
				goto fail8
			}
			labels[1] = parser.text[pos11:pos]
		}
		goto ok4
	fail8:
		pos = pos7
		// action
		// x1:fragment s1:_
		// x1:fragment
		{
			pos14 := pos
			// fragment
			if !_accept(parser, _fragmentAccepts, &pos, &perr) {
				goto fail12
			}
			labels[2] = parser.text[pos14:pos]
		}
		// s1:_
		{
			pos15 := pos
			// _
			if !_accept(parser, __Accepts, &pos, &perr) {
				goto fail12
			}
			labels[3] = parser.text[pos15:pos]
		}
		goto ok4
	fail12:
		pos = pos7
		goto fail
	ok4:
	}
	for {
		pos1 := pos
		// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
		// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
		{
			pos19 := pos
			// action
			// x0:sentence s0:_
			// x0:sentence
			{
				pos22 := pos
				// sentence
				if !_accept(parser, _sentenceAccepts, &pos, &perr) {
					goto fail20
				}
				labels[0] = parser.text[pos22:pos]
			}
			// s0:_
			{
				pos23 := pos
				// _
				if !_accept(parser, __Accepts, &pos, &perr) {
					goto fail20
				}
				labels[1] = parser.text[pos23:pos]
			}
			goto ok16
		fail20:
			pos = pos19
			// action
			// x1:fragment s1:_
			// x1:fragment
			{
				pos26 := pos
				// fragment
				if !_accept(parser, _fragmentAccepts, &pos, &perr) {
					goto fail24
				}
				labels[2] = parser.text[pos26:pos]
			}
			// s1:_
			{
				pos27 := pos
				// _
				if !_accept(parser, __Accepts, &pos, &perr) {
					goto fail24
				}
				labels[3] = parser.text[pos27:pos]
			}
			goto ok16
		fail24:
			pos = pos19
			goto fail3
		ok16:
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
	var labels [4]string
	use(labels)
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
	// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})+
	// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
	{
		nkids4 := len(node.Kids)
		pos05 := pos
		// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
		{
			pos9 := pos
			nkids7 := len(node.Kids)
			// action
			// x0:sentence s0:_
			// x0:sentence
			{
				pos12 := pos
				// sentence
				if !_node(parser, _sentenceNode, node, &pos) {
					goto fail10
				}
				labels[0] = parser.text[pos12:pos]
			}
			// s0:_
			{
				pos13 := pos
				// _
				if !_node(parser, __Node, node, &pos) {
					goto fail10
				}
				labels[1] = parser.text[pos13:pos]
			}
			goto ok6
		fail10:
			node.Kids = node.Kids[:nkids7]
			pos = pos9
			// action
			// x1:fragment s1:_
			// x1:fragment
			{
				pos16 := pos
				// fragment
				if !_node(parser, _fragmentNode, node, &pos) {
					goto fail14
				}
				labels[2] = parser.text[pos16:pos]
			}
			// s1:_
			{
				pos17 := pos
				// _
				if !_node(parser, __Node, node, &pos) {
					goto fail14
				}
				labels[3] = parser.text[pos17:pos]
			}
			goto ok6
		fail14:
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
		// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
		{
			nkids18 := len(node.Kids)
			pos019 := pos
			// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
			{
				pos23 := pos
				nkids21 := len(node.Kids)
				// action
				// x0:sentence s0:_
				// x0:sentence
				{
					pos26 := pos
					// sentence
					if !_node(parser, _sentenceNode, node, &pos) {
						goto fail24
					}
					labels[0] = parser.text[pos26:pos]
				}
				// s0:_
				{
					pos27 := pos
					// _
					if !_node(parser, __Node, node, &pos) {
						goto fail24
					}
					labels[1] = parser.text[pos27:pos]
				}
				goto ok20
			fail24:
				node.Kids = node.Kids[:nkids21]
				pos = pos23
				// action
				// x1:fragment s1:_
				// x1:fragment
				{
					pos30 := pos
					// fragment
					if !_node(parser, _fragmentNode, node, &pos) {
						goto fail28
					}
					labels[2] = parser.text[pos30:pos]
				}
				// s1:_
				{
					pos31 := pos
					// _
					if !_node(parser, __Node, node, &pos) {
						goto fail28
					}
					labels[3] = parser.text[pos31:pos]
				}
				goto ok20
			fail28:
				node.Kids = node.Kids[:nkids21]
				pos = pos23
				goto fail3
			ok20:
			}
			sub := _sub(parser, pos019, pos, node.Kids[nkids18:])
			node.Kids = append(node.Kids[:nkids18], sub)
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
	var labels [4]string
	use(labels)
	pos, failure := _failMemo(parser, _discourse, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "discourse",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _discourse}
	// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})+
	// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
	// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
	{
		pos7 := pos
		// action
		// x0:sentence s0:_
		// x0:sentence
		{
			pos10 := pos
			// sentence
			if !_fail(parser, _sentenceFail, errPos, failure, &pos) {
				goto fail8
			}
			labels[0] = parser.text[pos10:pos]
		}
		// s0:_
		{
			pos11 := pos
			// _
			if !_fail(parser, __Fail, errPos, failure, &pos) {
				goto fail8
			}
			labels[1] = parser.text[pos11:pos]
		}
		goto ok4
	fail8:
		pos = pos7
		// action
		// x1:fragment s1:_
		// x1:fragment
		{
			pos14 := pos
			// fragment
			if !_fail(parser, _fragmentFail, errPos, failure, &pos) {
				goto fail12
			}
			labels[2] = parser.text[pos14:pos]
		}
		// s1:_
		{
			pos15 := pos
			// _
			if !_fail(parser, __Fail, errPos, failure, &pos) {
				goto fail12
			}
			labels[3] = parser.text[pos15:pos]
		}
		goto ok4
	fail12:
		pos = pos7
		goto fail
	ok4:
	}
	for {
		pos1 := pos
		// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
		// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
		{
			pos19 := pos
			// action
			// x0:sentence s0:_
			// x0:sentence
			{
				pos22 := pos
				// sentence
				if !_fail(parser, _sentenceFail, errPos, failure, &pos) {
					goto fail20
				}
				labels[0] = parser.text[pos22:pos]
			}
			// s0:_
			{
				pos23 := pos
				// _
				if !_fail(parser, __Fail, errPos, failure, &pos) {
					goto fail20
				}
				labels[1] = parser.text[pos23:pos]
			}
			goto ok16
		fail20:
			pos = pos19
			// action
			// x1:fragment s1:_
			// x1:fragment
			{
				pos26 := pos
				// fragment
				if !_fail(parser, _fragmentFail, errPos, failure, &pos) {
					goto fail24
				}
				labels[2] = parser.text[pos26:pos]
			}
			// s1:_
			{
				pos27 := pos
				// _
				if !_fail(parser, __Fail, errPos, failure, &pos) {
					goto fail24
				}
				labels[3] = parser.text[pos27:pos]
			}
			goto ok16
		fail24:
			pos = pos19
			goto fail3
		ok16:
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

func _discourseAction(parser *_Parser, start int) (int, *[]Node) {
	var labels [4]string
	use(labels)
	var label0 Sentence
	var label1 *Mod
	var label2 Fragment
	var label3 *Mod
	dp := parser.deltaPos[start][_discourse]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _discourse}
	n := parser.act[key]
	if n != nil {
		n := n.([]Node)
		return start + int(dp-1), &n
	}
	var node []Node
	pos := start
	// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})+
	{
		var node2 Node
		// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
		// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
		{
			pos7 := pos
			var node6 Node
			// action
			{
				start9 := pos
				// x0:sentence s0:_
				// x0:sentence
				{
					pos11 := pos
					// sentence
					if p, n := _sentenceAction(parser, pos); n == nil {
						goto fail8
					} else {
						label0 = *n
						pos = p
					}
					labels[0] = parser.text[pos11:pos]
				}
				// s0:_
				{
					pos12 := pos
					// _
					if p, n := __Action(parser, pos); n == nil {
						goto fail8
					} else {
						label1 = *n
						pos = p
					}
					labels[1] = parser.text[pos12:pos]
				}
				node2 = func(
					start, end int, s0 *Mod, x0 Sentence) Node {
					return Node(x0.modSentence(s0))
				}(
					start9, pos, label1, label0)
			}
			goto ok4
		fail8:
			node2 = node6
			pos = pos7
			// action
			{
				start14 := pos
				// x1:fragment s1:_
				// x1:fragment
				{
					pos16 := pos
					// fragment
					if p, n := _fragmentAction(parser, pos); n == nil {
						goto fail13
					} else {
						label2 = *n
						pos = p
					}
					labels[2] = parser.text[pos16:pos]
				}
				// s1:_
				{
					pos17 := pos
					// _
					if p, n := __Action(parser, pos); n == nil {
						goto fail13
					} else {
						label3 = *n
						pos = p
					}
					labels[3] = parser.text[pos17:pos]
				}
				node2 = func(
					start, end int, s0 *Mod, s1 *Mod, x0 Sentence, x1 Fragment) Node {
					return Node(x1.modFragment(s1))
				}(
					start14, pos, label1, label3, label0, label2)
			}
			goto ok4
		fail13:
			node2 = node6
			pos = pos7
			goto fail
		ok4:
		}
		node = append(node, node2)
	}
	for {
		pos1 := pos
		var node2 Node
		// (x0:sentence s0:_ {…}/x1:fragment s1:_ {…})
		// x0:sentence s0:_ {…}/x1:fragment s1:_ {…}
		{
			pos21 := pos
			var node20 Node
			// action
			{
				start23 := pos
				// x0:sentence s0:_
				// x0:sentence
				{
					pos25 := pos
					// sentence
					if p, n := _sentenceAction(parser, pos); n == nil {
						goto fail22
					} else {
						label0 = *n
						pos = p
					}
					labels[0] = parser.text[pos25:pos]
				}
				// s0:_
				{
					pos26 := pos
					// _
					if p, n := __Action(parser, pos); n == nil {
						goto fail22
					} else {
						label1 = *n
						pos = p
					}
					labels[1] = parser.text[pos26:pos]
				}
				node2 = func(
					start, end int, s0 *Mod, x0 Sentence) Node {
					return Node(x0.modSentence(s0))
				}(
					start23, pos, label1, label0)
			}
			goto ok18
		fail22:
			node2 = node20
			pos = pos21
			// action
			{
				start28 := pos
				// x1:fragment s1:_
				// x1:fragment
				{
					pos30 := pos
					// fragment
					if p, n := _fragmentAction(parser, pos); n == nil {
						goto fail27
					} else {
						label2 = *n
						pos = p
					}
					labels[2] = parser.text[pos30:pos]
				}
				// s1:_
				{
					pos31 := pos
					// _
					if p, n := __Action(parser, pos); n == nil {
						goto fail27
					} else {
						label3 = *n
						pos = p
					}
					labels[3] = parser.text[pos31:pos]
				}
				node2 = func(
					start, end int, s0 *Mod, s1 *Mod, x0 Sentence, x1 Fragment) Node {
					return Node(x1.modFragment(s1))
				}(
					start28, pos, label1, label3, label0, label2)
			}
			goto ok18
		fail27:
			node2 = node20
			pos = pos21
			goto fail3
		ok18:
		}
		node = append(node, node2)
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

func _fragmentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [7]string
	use(labels)
	if dp, de, ok := _memo(parser, _fragment, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}) s:_
	// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
	{
		pos1 := pos
		// (x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
		// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}
		{
			pos5 := pos
			// action
			// x0:afterthought_cop<relative_clause_1, relative_clause_1>
			{
				pos7 := pos
				// afterthought_cop<relative_clause_1, relative_clause_1>
				if !_accept(parser, _afterthought_cop__relative_clause_1__relative_clause_1Accepts, &pos, &perr) {
					goto fail6
				}
				labels[0] = parser.text[pos7:pos]
			}
			goto ok2
		fail6:
			pos = pos5
			// action
			// x1:relative_clause
			{
				pos9 := pos
				// relative_clause
				if !_accept(parser, _relative_clauseAccepts, &pos, &perr) {
					goto fail8
				}
				labels[1] = parser.text[pos9:pos]
			}
			goto ok2
		fail8:
			pos = pos5
			// action
			// x2:prenex
			{
				pos11 := pos
				// prenex
				if !_accept(parser, _prenexAccepts, &pos, &perr) {
					goto fail10
				}
				labels[2] = parser.text[pos11:pos]
			}
			goto ok2
		fail10:
			pos = pos5
			// action
			// x3:terms
			{
				pos13 := pos
				// terms
				if !_accept(parser, _termsAccepts, &pos, &perr) {
					goto fail12
				}
				labels[3] = parser.text[pos13:pos]
			}
			goto ok2
		fail12:
			pos = pos5
			// action
			// x4:freemod
			{
				pos15 := pos
				// freemod
				if !_accept(parser, _freemodAccepts, &pos, &perr) {
					goto fail14
				}
				labels[4] = parser.text[pos15:pos]
			}
			goto ok2
		fail14:
			pos = pos5
			goto fail
		ok2:
		}
		labels[5] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos16 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[6] = parser.text[pos16:pos]
	}
	return _memoize(parser, _fragment, start, pos, perr)
fail:
	return _memoize(parser, _fragment, start, -1, perr)
}

func _fragmentNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [7]string
	use(labels)
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
	// action
	// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}) s:_
	// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
	{
		pos1 := pos
		// (x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
		{
			nkids2 := len(node.Kids)
			pos03 := pos
			// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}
			{
				pos7 := pos
				nkids5 := len(node.Kids)
				// action
				// x0:afterthought_cop<relative_clause_1, relative_clause_1>
				{
					pos9 := pos
					// afterthought_cop<relative_clause_1, relative_clause_1>
					if !_node(parser, _afterthought_cop__relative_clause_1__relative_clause_1Node, node, &pos) {
						goto fail8
					}
					labels[0] = parser.text[pos9:pos]
				}
				goto ok4
			fail8:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				// action
				// x1:relative_clause
				{
					pos11 := pos
					// relative_clause
					if !_node(parser, _relative_clauseNode, node, &pos) {
						goto fail10
					}
					labels[1] = parser.text[pos11:pos]
				}
				goto ok4
			fail10:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				// action
				// x2:prenex
				{
					pos13 := pos
					// prenex
					if !_node(parser, _prenexNode, node, &pos) {
						goto fail12
					}
					labels[2] = parser.text[pos13:pos]
				}
				goto ok4
			fail12:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				// action
				// x3:terms
				{
					pos15 := pos
					// terms
					if !_node(parser, _termsNode, node, &pos) {
						goto fail14
					}
					labels[3] = parser.text[pos15:pos]
				}
				goto ok4
			fail14:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				// action
				// x4:freemod
				{
					pos17 := pos
					// freemod
					if !_node(parser, _freemodNode, node, &pos) {
						goto fail16
					}
					labels[4] = parser.text[pos17:pos]
				}
				goto ok4
			fail16:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				goto fail
			ok4:
			}
			sub := _sub(parser, pos03, pos, node.Kids[nkids2:])
			node.Kids = append(node.Kids[:nkids2], sub)
		}
		labels[5] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos18 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[6] = parser.text[pos18:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _fragmentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [7]string
	use(labels)
	pos, failure := _failMemo(parser, _fragment, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "fragment",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _fragment}
	// action
	// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}) s:_
	// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
	{
		pos1 := pos
		// (x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
		// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}
		{
			pos5 := pos
			// action
			// x0:afterthought_cop<relative_clause_1, relative_clause_1>
			{
				pos7 := pos
				// afterthought_cop<relative_clause_1, relative_clause_1>
				if !_fail(parser, _afterthought_cop__relative_clause_1__relative_clause_1Fail, errPos, failure, &pos) {
					goto fail6
				}
				labels[0] = parser.text[pos7:pos]
			}
			goto ok2
		fail6:
			pos = pos5
			// action
			// x1:relative_clause
			{
				pos9 := pos
				// relative_clause
				if !_fail(parser, _relative_clauseFail, errPos, failure, &pos) {
					goto fail8
				}
				labels[1] = parser.text[pos9:pos]
			}
			goto ok2
		fail8:
			pos = pos5
			// action
			// x2:prenex
			{
				pos11 := pos
				// prenex
				if !_fail(parser, _prenexFail, errPos, failure, &pos) {
					goto fail10
				}
				labels[2] = parser.text[pos11:pos]
			}
			goto ok2
		fail10:
			pos = pos5
			// action
			// x3:terms
			{
				pos13 := pos
				// terms
				if !_fail(parser, _termsFail, errPos, failure, &pos) {
					goto fail12
				}
				labels[3] = parser.text[pos13:pos]
			}
			goto ok2
		fail12:
			pos = pos5
			// action
			// x4:freemod
			{
				pos15 := pos
				// freemod
				if !_fail(parser, _freemodFail, errPos, failure, &pos) {
					goto fail14
				}
				labels[4] = parser.text[pos15:pos]
			}
			goto ok2
		fail14:
			pos = pos5
			goto fail
		ok2:
		}
		labels[5] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos16 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[6] = parser.text[pos16:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _fragmentAction(parser *_Parser, start int) (int, *Fragment) {
	var labels [7]string
	use(labels)
	var label3 Terms
	var label4 Mod
	var label5 Fragment
	var label6 *Mod
	var label0 (*CoP)
	var label1 Relative
	var label2 (*Prenex)
	dp := parser.deltaPos[start][_fragment]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _fragment}
	n := parser.act[key]
	if n != nil {
		n := n.(Fragment)
		return start + int(dp-1), &n
	}
	var node Fragment
	pos := start
	// action
	{
		start0 := pos
		// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}) s:_
		// f:(x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
		{
			pos2 := pos
			// (x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…})
			// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:relative_clause {…}/x2:prenex {…}/x3:terms {…}/x4:freemod {…}
			{
				pos6 := pos
				var node5 Fragment
				// action
				{
					start8 := pos
					// x0:afterthought_cop<relative_clause_1, relative_clause_1>
					{
						pos9 := pos
						// afterthought_cop<relative_clause_1, relative_clause_1>
						if p, n := _afterthought_cop__relative_clause_1__relative_clause_1Action(parser, pos); n == nil {
							goto fail7
						} else {
							label0 = *n
							pos = p
						}
						labels[0] = parser.text[pos9:pos]
					}
					label5 = func(
						start, end int, x0 *CoP) Fragment {
						return Fragment((*CoPRelative)(x0))
					}(
						start8, pos, label0)
				}
				goto ok3
			fail7:
				label5 = node5
				pos = pos6
				// action
				{
					start11 := pos
					// x1:relative_clause
					{
						pos12 := pos
						// relative_clause
						if p, n := _relative_clauseAction(parser, pos); n == nil {
							goto fail10
						} else {
							label1 = *n
							pos = p
						}
						labels[1] = parser.text[pos12:pos]
					}
					label5 = func(
						start, end int, x0 *CoP, x1 Relative) Fragment {
						return Fragment(x1)
					}(
						start11, pos, label0, label1)
				}
				goto ok3
			fail10:
				label5 = node5
				pos = pos6
				// action
				{
					start14 := pos
					// x2:prenex
					{
						pos15 := pos
						// prenex
						if p, n := _prenexAction(parser, pos); n == nil {
							goto fail13
						} else {
							label2 = *n
							pos = p
						}
						labels[2] = parser.text[pos15:pos]
					}
					label5 = func(
						start, end int, x0 *CoP, x1 Relative, x2 *Prenex) Fragment {
						return Fragment(x2)
					}(
						start14, pos, label0, label1, label2)
				}
				goto ok3
			fail13:
				label5 = node5
				pos = pos6
				// action
				{
					start17 := pos
					// x3:terms
					{
						pos18 := pos
						// terms
						if p, n := _termsAction(parser, pos); n == nil {
							goto fail16
						} else {
							label3 = *n
							pos = p
						}
						labels[3] = parser.text[pos18:pos]
					}
					label5 = func(
						start, end int, x0 *CoP, x1 Relative, x2 *Prenex, x3 Terms) Fragment {
						return Fragment(x3)
					}(
						start17, pos, label0, label1, label2, label3)
				}
				goto ok3
			fail16:
				label5 = node5
				pos = pos6
				// action
				{
					start20 := pos
					// x4:freemod
					{
						pos21 := pos
						// freemod
						if p, n := _freemodAction(parser, pos); n == nil {
							goto fail19
						} else {
							label4 = *n
							pos = p
						}
						labels[4] = parser.text[pos21:pos]
					}
					label5 = func(
						start, end int, x0 *CoP, x1 Relative, x2 *Prenex, x3 Terms, x4 Mod) Fragment {
						return Fragment(x4)
					}(
						start20, pos, label0, label1, label2, label3, label4)
				}
				goto ok3
			fail19:
				label5 = node5
				pos = pos6
				goto fail
			ok3:
			}
			labels[5] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos22 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label6 = *n
				pos = p
			}
			labels[6] = parser.text[pos22:pos]
		}
		node = func(
			start, end int, f Fragment, s *Mod, x0 *CoP, x1 Relative, x2 *Prenex, x3 Terms, x4 Mod) Fragment {
			return Fragment(f.modFragment(s))
		}(
			start0, pos, label5, label6, label0, label1, label2, label3, label4)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _sentence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:afterthought_cop<sentence_1, sentence> {…}/sentence_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<sentence_1, sentence>
		{
			pos5 := pos
			// afterthought_cop<sentence_1, sentence>
			if !_accept(parser, _afterthought_cop__sentence_1__sentenceAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// sentence_1
		if !_accept(parser, _sentence_1Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _sentence, start, pos, perr)
fail:
	return _memoize(parser, _sentence, start, -1, perr)
}

func _sentenceNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:afterthought_cop<sentence_1, sentence> {…}/sentence_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:afterthought_cop<sentence_1, sentence>
		{
			pos5 := pos
			// afterthought_cop<sentence_1, sentence>
			if !_node(parser, _afterthought_cop__sentence_1__sentenceNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// sentence_1
		if !_node(parser, _sentence_1Node, node, &pos) {
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

func _sentenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _sentence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence}
	// x:afterthought_cop<sentence_1, sentence> {…}/sentence_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<sentence_1, sentence>
		{
			pos5 := pos
			// afterthought_cop<sentence_1, sentence>
			if !_fail(parser, _afterthought_cop__sentence_1__sentenceFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// sentence_1
		if !_fail(parser, _sentence_1Fail, errPos, failure, &pos) {
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

func _sentenceAction(parser *_Parser, start int) (int, *Sentence) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_sentence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence}
	n := parser.act[key]
	if n != nil {
		n := n.(Sentence)
		return start + int(dp-1), &n
	}
	var node Sentence
	pos := start
	// x:afterthought_cop<sentence_1, sentence> {…}/sentence_1
	{
		pos3 := pos
		var node2 Sentence
		// action
		{
			start5 := pos
			// x:afterthought_cop<sentence_1, sentence>
			{
				pos6 := pos
				// afterthought_cop<sentence_1, sentence>
				if p, n := _afterthought_cop__sentence_1__sentenceAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Sentence {
				return Sentence((*CoPSentence)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// sentence_1
		if p, n := _sentence_1Action(parser, pos); n == nil {
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

func _sentence_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _sentence_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// sentence_2/x:forethought_cop<sentence, sentence> {…}
	{
		pos3 := pos
		// sentence_2
		if !_accept(parser, _sentence_2Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x:forethought_cop<sentence, sentence>
		{
			pos6 := pos
			// forethought_cop<sentence, sentence>
			if !_accept(parser, _forethought_cop__sentence__sentenceAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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
	var labels [1]string
	use(labels)
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
	// sentence_2/x:forethought_cop<sentence, sentence> {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// sentence_2
		if !_node(parser, _sentence_2Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x:forethought_cop<sentence, sentence>
		{
			pos6 := pos
			// forethought_cop<sentence, sentence>
			if !_node(parser, _forethought_cop__sentence__sentenceNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _sentence_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_1}
	// sentence_2/x:forethought_cop<sentence, sentence> {…}
	{
		pos3 := pos
		// sentence_2
		if !_fail(parser, _sentence_2Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x:forethought_cop<sentence, sentence>
		{
			pos6 := pos
			// forethought_cop<sentence, sentence>
			if !_fail(parser, _forethought_cop__sentence__sentenceFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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

func _sentence_1Action(parser *_Parser, start int) (int, *Sentence) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_sentence_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Sentence)
		return start + int(dp-1), &n
	}
	var node Sentence
	pos := start
	// sentence_2/x:forethought_cop<sentence, sentence> {…}
	{
		pos3 := pos
		var node2 Sentence
		// sentence_2
		if p, n := _sentence_2Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// x:forethought_cop<sentence, sentence>
			{
				pos7 := pos
				// forethought_cop<sentence, sentence>
				if p, n := _forethought_cop__sentence__sentenceAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos7:pos]
			}
			node = func(
				start, end int, x *CoP) Sentence {
				return Sentence((*CoPSentence)(x))
			}(
				start6, pos, label0)
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
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _sentence_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// je:sentence_prefix_spaces? se:sentence_3
	// je:sentence_prefix_spaces?
	{
		pos1 := pos
		// sentence_prefix_spaces?
		{
			pos3 := pos
			// sentence_prefix_spaces
			if !_accept(parser, _sentence_prefix_spacesAccepts, &pos, &perr) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// se:sentence_3
	{
		pos6 := pos
		// sentence_3
		if !_accept(parser, _sentence_3Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	return _memoize(parser, _sentence_2, start, pos, perr)
fail:
	return _memoize(parser, _sentence_2, start, -1, perr)
}

func _sentence_2Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// action
	// je:sentence_prefix_spaces? se:sentence_3
	// je:sentence_prefix_spaces?
	{
		pos1 := pos
		// sentence_prefix_spaces?
		{
			nkids2 := len(node.Kids)
			pos3 := pos
			// sentence_prefix_spaces
			if !_node(parser, _sentence_prefix_spacesNode, node, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			node.Kids = node.Kids[:nkids2]
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// se:sentence_3
	{
		pos6 := pos
		// sentence_3
		if !_node(parser, _sentence_3Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _sentence_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_2}
	// action
	// je:sentence_prefix_spaces? se:sentence_3
	// je:sentence_prefix_spaces?
	{
		pos1 := pos
		// sentence_prefix_spaces?
		{
			pos3 := pos
			// sentence_prefix_spaces
			if !_fail(parser, _sentence_prefix_spacesFail, errPos, failure, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// se:sentence_3
	{
		pos6 := pos
		// sentence_3
		if !_fail(parser, _sentence_3Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_2Action(parser *_Parser, start int) (int, *Sentence) {
	var labels [2]string
	use(labels)
	var label0 *Word
	var label1 *StatementSentence
	dp := parser.deltaPos[start][_sentence_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Sentence)
		return start + int(dp-1), &n
	}
	var node Sentence
	pos := start
	// action
	{
		start0 := pos
		// je:sentence_prefix_spaces? se:sentence_3
		// je:sentence_prefix_spaces?
		{
			pos2 := pos
			// sentence_prefix_spaces?
			{
				pos4 := pos
				label0 = new(Word)
				// sentence_prefix_spaces
				if p, n := _sentence_prefix_spacesAction(parser, pos); n == nil {
					goto fail5
				} else {
					*label0 = *n
					pos = p
				}
				goto ok6
			fail5:
				label0 = nil
				pos = pos4
			ok6:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// se:sentence_3
		{
			pos7 := pos
			// sentence_3
			if p, n := _sentence_3Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, je *Word, se *StatementSentence) Sentence {
			se.JE = je
			return Sentence(se)
		}(
			start0, pos, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentence_prefix_spacesAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _sentence_prefix_spaces, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// je:sentence_prefix s:_
	// je:sentence_prefix
	{
		pos1 := pos
		// sentence_prefix
		if !_accept(parser, _sentence_prefixAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	return _memoize(parser, _sentence_prefix_spaces, start, pos, perr)
fail:
	return _memoize(parser, _sentence_prefix_spaces, start, -1, perr)
}

func _sentence_prefix_spacesNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_sentence_prefix_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_prefix_spaces}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "sentence_prefix_spaces"}
	// action
	// je:sentence_prefix s:_
	// je:sentence_prefix
	{
		pos1 := pos
		// sentence_prefix
		if !_node(parser, _sentence_prefixNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_prefix_spacesFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _sentence_prefix_spaces, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_prefix_spaces",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_prefix_spaces}
	// action
	// je:sentence_prefix s:_
	// je:sentence_prefix
	{
		pos1 := pos
		// sentence_prefix
		if !_fail(parser, _sentence_prefixFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_prefix_spacesAction(parser *_Parser, start int) (int, *Word) {
	var labels [2]string
	use(labels)
	var label0 Word
	var label1 *Mod
	dp := parser.deltaPos[start][_sentence_prefix_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_prefix_spaces}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// je:sentence_prefix s:_
		// je:sentence_prefix
		{
			pos2 := pos
			// sentence_prefix
			if p, n := _sentence_prefixAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		node = func(
			start, end int, je Word, s *Mod) Word {
			return Word(*je.mod(s))
		}(
			start0, pos, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentence_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _sentence_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// st:statement s:_ da:illocutionary?
	// st:statement
	{
		pos1 := pos
		// statement
		if !_accept(parser, _statementAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// da:illocutionary?
	{
		pos3 := pos
		// illocutionary?
		{
			pos5 := pos
			// illocutionary
			if !_accept(parser, _illocutionaryAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _sentence_3, start, pos, perr)
fail:
	return _memoize(parser, _sentence_3, start, -1, perr)
}

func _sentence_3Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// st:statement s:_ da:illocutionary?
	// st:statement
	{
		pos1 := pos
		// statement
		if !_node(parser, _statementNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// da:illocutionary?
	{
		pos3 := pos
		// illocutionary?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// illocutionary
			if !_node(parser, _illocutionaryNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _sentence_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_3}
	// action
	// st:statement s:_ da:illocutionary?
	// st:statement
	{
		pos1 := pos
		// statement
		if !_fail(parser, _statementFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// da:illocutionary?
	{
		pos3 := pos
		// illocutionary?
		{
			pos5 := pos
			// illocutionary
			if !_fail(parser, _illocutionaryFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_3Action(parser *_Parser, start int) (int, **StatementSentence) {
	var labels [3]string
	use(labels)
	var label0 Statement
	var label1 *Mod
	var label2 *Word
	dp := parser.deltaPos[start][_sentence_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_3}
	n := parser.act[key]
	if n != nil {
		n := n.(*StatementSentence)
		return start + int(dp-1), &n
	}
	var node *StatementSentence
	pos := start
	// action
	{
		start0 := pos
		// st:statement s:_ da:illocutionary?
		// st:statement
		{
			pos2 := pos
			// statement
			if p, n := _statementAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// da:illocutionary?
		{
			pos4 := pos
			// illocutionary?
			{
				pos6 := pos
				label2 = new(Word)
				// illocutionary
				if p, n := _illocutionaryAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, da *Word, s *Mod, st Statement) *StatementSentence {
			return &StatementSentence{Statement: st.modStatement(s), DA: da}
		}(
			start0, pos, label2, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:prenex_spaces? s:statement_1
	// p:prenex_spaces?
	{
		pos1 := pos
		// prenex_spaces?
		{
			pos3 := pos
			// prenex_spaces
			if !_accept(parser, _prenex_spacesAccepts, &pos, &perr) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:statement_1
	{
		pos6 := pos
		// statement_1
		if !_accept(parser, _statement_1Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	return _memoize(parser, _statement, start, pos, perr)
fail:
	return _memoize(parser, _statement, start, -1, perr)
}

func _statementNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// action
	// p:prenex_spaces? s:statement_1
	// p:prenex_spaces?
	{
		pos1 := pos
		// prenex_spaces?
		{
			nkids2 := len(node.Kids)
			pos3 := pos
			// prenex_spaces
			if !_node(parser, _prenex_spacesNode, node, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			node.Kids = node.Kids[:nkids2]
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:statement_1
	{
		pos6 := pos
		// statement_1
		if !_node(parser, _statement_1Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement}
	// action
	// p:prenex_spaces? s:statement_1
	// p:prenex_spaces?
	{
		pos1 := pos
		// prenex_spaces?
		{
			pos3 := pos
			// prenex_spaces
			if !_fail(parser, _prenex_spacesFail, errPos, failure, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:statement_1
	{
		pos6 := pos
		// statement_1
		if !_fail(parser, _statement_1Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _statementAction(parser *_Parser, start int) (int, *Statement) {
	var labels [2]string
	use(labels)
	var label0 *Prenex
	var label1 Statement
	dp := parser.deltaPos[start][_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement}
	n := parser.act[key]
	if n != nil {
		n := n.(Statement)
		return start + int(dp-1), &n
	}
	var node Statement
	pos := start
	// action
	{
		start0 := pos
		// p:prenex_spaces? s:statement_1
		// p:prenex_spaces?
		{
			pos2 := pos
			// prenex_spaces?
			{
				pos4 := pos
				label0 = new(Prenex)
				// prenex_spaces
				if p, n := _prenex_spacesAction(parser, pos); n == nil {
					goto fail5
				} else {
					*label0 = *n
					pos = p
				}
				goto ok6
			fail5:
				label0 = nil
				pos = pos4
			ok6:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:statement_1
		{
			pos7 := pos
			// statement_1
			if p, n := _statement_1Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, p *Prenex, s Statement) Statement {
			return Statement(prenexStmt(p, s))
		}(
			start0, pos, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _statement_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _statement_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// n:afterthought_cop<statement_2, statement> {…}/statement_2
	{
		pos3 := pos
		// action
		// n:afterthought_cop<statement_2, statement>
		{
			pos5 := pos
			// afterthought_cop<statement_2, statement>
			if !_accept(parser, _afterthought_cop__statement_2__statementAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// statement_2
		if !_accept(parser, _statement_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _statement_1, start, pos, perr)
fail:
	return _memoize(parser, _statement_1, start, -1, perr)
}

func _statement_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// n:afterthought_cop<statement_2, statement> {…}/statement_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// n:afterthought_cop<statement_2, statement>
		{
			pos5 := pos
			// afterthought_cop<statement_2, statement>
			if !_node(parser, _afterthought_cop__statement_2__statementNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// statement_2
		if !_node(parser, _statement_2Node, node, &pos) {
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

func _statement_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _statement_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement_1}
	// n:afterthought_cop<statement_2, statement> {…}/statement_2
	{
		pos3 := pos
		// action
		// n:afterthought_cop<statement_2, statement>
		{
			pos5 := pos
			// afterthought_cop<statement_2, statement>
			if !_fail(parser, _afterthought_cop__statement_2__statementFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// statement_2
		if !_fail(parser, _statement_2Fail, errPos, failure, &pos) {
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

func _statement_1Action(parser *_Parser, start int) (int, *Statement) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_statement_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Statement)
		return start + int(dp-1), &n
	}
	var node Statement
	pos := start
	// n:afterthought_cop<statement_2, statement> {…}/statement_2
	{
		pos3 := pos
		var node2 Statement
		// action
		{
			start5 := pos
			// n:afterthought_cop<statement_2, statement>
			{
				pos6 := pos
				// afterthought_cop<statement_2, statement>
				if p, n := _afterthought_cop__statement_2__statementAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, n *CoP) Statement {
				return Statement((*CoPStatement)(n))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// statement_2
		if p, n := _statement_2Action(parser, pos); n == nil {
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

func _statement_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _statement_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// statement_3/n:forethought_cop<statement, statement> {…}
	{
		pos3 := pos
		// statement_3
		if !_accept(parser, _statement_3Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// n:forethought_cop<statement, statement>
		{
			pos6 := pos
			// forethought_cop<statement, statement>
			if !_accept(parser, _forethought_cop__statement__statementAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _statement_2, start, pos, perr)
fail:
	return _memoize(parser, _statement_2, start, -1, perr)
}

func _statement_2Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// statement_3/n:forethought_cop<statement, statement> {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// statement_3
		if !_node(parser, _statement_3Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// n:forethought_cop<statement, statement>
		{
			pos6 := pos
			// forethought_cop<statement, statement>
			if !_node(parser, _forethought_cop__statement__statementNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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

func _statement_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _statement_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement_2}
	// statement_3/n:forethought_cop<statement, statement> {…}
	{
		pos3 := pos
		// statement_3
		if !_fail(parser, _statement_3Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// n:forethought_cop<statement, statement>
		{
			pos6 := pos
			// forethought_cop<statement, statement>
			if !_fail(parser, _forethought_cop__statement__statementFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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

func _statement_2Action(parser *_Parser, start int) (int, *Statement) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_statement_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Statement)
		return start + int(dp-1), &n
	}
	var node Statement
	pos := start
	// statement_3/n:forethought_cop<statement, statement> {…}
	{
		pos3 := pos
		var node2 Statement
		// statement_3
		if p, n := _statement_3Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// n:forethought_cop<statement, statement>
			{
				pos7 := pos
				// forethought_cop<statement, statement>
				if p, n := _forethought_cop__statement__statementAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos7:pos]
			}
			node = func(
				start, end int, n *CoP) Statement {
				return Statement((*CoPStatement)(n))
			}(
				start6, pos, label0)
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

func _statement_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _statement_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:predication s:_ na:end_statement?
	// p:predication
	{
		pos1 := pos
		// predication
		if !_accept(parser, _predicationAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			pos5 := pos
			// end_statement
			if !_accept(parser, _end_statementAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _statement_3, start, pos, perr)
fail:
	return _memoize(parser, _statement_3, start, -1, perr)
}

func _statement_3Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// p:predication s:_ na:end_statement?
	// p:predication
	{
		pos1 := pos
		// predication
		if !_node(parser, _predicationNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// end_statement
			if !_node(parser, _end_statementNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _statement_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _statement_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "statement_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _statement_3}
	// action
	// p:predication s:_ na:end_statement?
	// p:predication
	{
		pos1 := pos
		// predication
		if !_fail(parser, _predicationFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			pos5 := pos
			// end_statement
			if !_fail(parser, _end_statementFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _statement_3Action(parser *_Parser, start int) (int, *Statement) {
	var labels [3]string
	use(labels)
	var label0 Predication
	var label1 *Mod
	var label2 *Word
	dp := parser.deltaPos[start][_statement_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _statement_3}
	n := parser.act[key]
	if n != nil {
		n := n.(Statement)
		return start + int(dp-1), &n
	}
	var node Statement
	pos := start
	// action
	{
		start0 := pos
		// p:predication s:_ na:end_statement?
		// p:predication
		{
			pos2 := pos
			// predication
			if p, n := _predicationAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// na:end_statement?
		{
			pos4 := pos
			// end_statement?
			{
				pos6 := pos
				label2 = new(Word)
				// end_statement
				if p, n := _end_statementAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, na *Word, p Predication, s *Mod) Statement {
			return Statement(endPredication(p, s, na))
		}(
			start0, pos, label2, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicationAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _predication, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:predicate s:_ ts:terms?
	// p:predicate
	{
		pos1 := pos
		// predicate
		if !_accept(parser, _predicateAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			pos5 := pos
			// terms
			if !_accept(parser, _termsAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _predication, start, pos, perr)
fail:
	return _memoize(parser, _predication, start, -1, perr)
}

func _predicationNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// p:predicate s:_ ts:terms?
	// p:predicate
	{
		pos1 := pos
		// predicate
		if !_node(parser, _predicateNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// terms
			if !_node(parser, _termsNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicationFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _predication, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predication",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predication}
	// action
	// p:predicate s:_ ts:terms?
	// p:predicate
	{
		pos1 := pos
		// predicate
		if !_fail(parser, _predicateFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			pos5 := pos
			// terms
			if !_fail(parser, _termsFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicationAction(parser *_Parser, start int) (int, *Predication) {
	var labels [3]string
	use(labels)
	var label0 Predicate
	var label1 *Mod
	var label2 *Terms
	dp := parser.deltaPos[start][_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predication}
	n := parser.act[key]
	if n != nil {
		n := n.(Predication)
		return start + int(dp-1), &n
	}
	var node Predication
	pos := start
	// action
	{
		start0 := pos
		// p:predicate s:_ ts:terms?
		// p:predicate
		{
			pos2 := pos
			// predicate
			if p, n := _predicateAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// ts:terms?
		{
			pos4 := pos
			// terms?
			{
				pos6 := pos
				label2 = new(Terms)
				// terms
				if p, n := _termsAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, p Predicate, s *Mod, ts *Terms) Predication {
			return Predication(predication(p, s, ts))
		}(
			start0, pos, label0, label1, label2)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prenex_spacesAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _prenex_spaces, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:prenex s:_
	// p:prenex
	{
		pos1 := pos
		// prenex
		if !_accept(parser, _prenexAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	return _memoize(parser, _prenex_spaces, start, pos, perr)
fail:
	return _memoize(parser, _prenex_spaces, start, -1, perr)
}

func _prenex_spacesNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_prenex_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prenex_spaces}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "prenex_spaces"}
	// action
	// p:prenex s:_
	// p:prenex
	{
		pos1 := pos
		// prenex
		if !_node(parser, _prenexNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prenex_spacesFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _prenex_spaces, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prenex_spaces",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prenex_spaces}
	// action
	// p:prenex s:_
	// p:prenex
	{
		pos1 := pos
		// prenex
		if !_fail(parser, _prenexFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prenex_spacesAction(parser *_Parser, start int) (int, *Prenex) {
	var labels [2]string
	use(labels)
	var label0 (*Prenex)
	var label1 *Mod
	dp := parser.deltaPos[start][_prenex_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prenex_spaces}
	n := parser.act[key]
	if n != nil {
		n := n.(Prenex)
		return start + int(dp-1), &n
	}
	var node Prenex
	pos := start
	// action
	{
		start0 := pos
		// p:prenex s:_
		// p:prenex
		{
			pos2 := pos
			// prenex
			if p, n := _prenexAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		node = func(
			start, end int, p *Prenex, s *Mod) Prenex {
			return Prenex(*p.mod(s))
		}(
			start0, pos, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prenexAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _prenex, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// ts:terms s:_ bi:end_prenex
	// ts:terms
	{
		pos1 := pos
		// terms
		if !_accept(parser, _termsAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// bi:end_prenex
	{
		pos3 := pos
		// end_prenex
		if !_accept(parser, _end_prenexAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _prenex, start, pos, perr)
fail:
	return _memoize(parser, _prenex, start, -1, perr)
}

func _prenexNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// ts:terms s:_ bi:end_prenex
	// ts:terms
	{
		pos1 := pos
		// terms
		if !_node(parser, _termsNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// bi:end_prenex
	{
		pos3 := pos
		// end_prenex
		if !_node(parser, _end_prenexNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prenexFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _prenex, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prenex",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prenex}
	// action
	// ts:terms s:_ bi:end_prenex
	// ts:terms
	{
		pos1 := pos
		// terms
		if !_fail(parser, _termsFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// bi:end_prenex
	{
		pos3 := pos
		// end_prenex
		if !_fail(parser, _end_prenexFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prenexAction(parser *_Parser, start int) (int, *(*Prenex)) {
	var labels [3]string
	use(labels)
	var label1 *Mod
	var label2 Word
	var label0 Terms
	dp := parser.deltaPos[start][_prenex]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prenex}
	n := parser.act[key]
	if n != nil {
		n := n.((*Prenex))
		return start + int(dp-1), &n
	}
	var node (*Prenex)
	pos := start
	// action
	{
		start0 := pos
		// ts:terms s:_ bi:end_prenex
		// ts:terms
		{
			pos2 := pos
			// terms
			if p, n := _termsAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// bi:end_prenex
		{
			pos4 := pos
			// end_prenex
			if p, n := _end_prenexAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, bi Word, s *Mod, ts Terms) *Prenex {
			return (*Prenex)(prenex(ts, s, bi))
		}(
			start0, pos, label2, label1, label0)
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

func _predicateAction(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
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
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _predicate_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:afterthought_cop<predicate_2, predicate_1> {…}/predicate_2
	{
		pos3 := pos
		// action
		// x:afterthought_cop<predicate_2, predicate_1>
		{
			pos5 := pos
			// afterthought_cop<predicate_2, predicate_1>
			if !_accept(parser, _afterthought_cop__predicate_2__predicate_1Accepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// predicate_2
		if !_accept(parser, _predicate_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _predicate_1, start, pos, perr)
fail:
	return _memoize(parser, _predicate_1, start, -1, perr)
}

func _predicate_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:afterthought_cop<predicate_2, predicate_1> {…}/predicate_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:afterthought_cop<predicate_2, predicate_1>
		{
			pos5 := pos
			// afterthought_cop<predicate_2, predicate_1>
			if !_node(parser, _afterthought_cop__predicate_2__predicate_1Node, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// predicate_2
		if !_node(parser, _predicate_2Node, node, &pos) {
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

func _predicate_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _predicate_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predicate_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predicate_1}
	// x:afterthought_cop<predicate_2, predicate_1> {…}/predicate_2
	{
		pos3 := pos
		// action
		// x:afterthought_cop<predicate_2, predicate_1>
		{
			pos5 := pos
			// afterthought_cop<predicate_2, predicate_1>
			if !_fail(parser, _afterthought_cop__predicate_2__predicate_1Fail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// predicate_2
		if !_fail(parser, _predicate_2Fail, errPos, failure, &pos) {
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

func _predicate_1Action(parser *_Parser, start int) (int, *Predicate) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// x:afterthought_cop<predicate_2, predicate_1> {…}/predicate_2
	{
		pos3 := pos
		var node2 Predicate
		// action
		{
			start5 := pos
			// x:afterthought_cop<predicate_2, predicate_1>
			{
				pos6 := pos
				// afterthought_cop<predicate_2, predicate_1>
				if p, n := _afterthought_cop__predicate_2__predicate_1Action(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Predicate {
				return Predicate((*CoPPredicate)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// predicate_2
		if p, n := _predicate_2Action(parser, pos); n == nil {
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

func _predicate_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _predicate_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// forethought_cop_pred<predicate>/LU_predicate<verb_syllable>/MI_phrase<verb_syllable>/PO_phrase<verb_syllable>/MO_phrase<verb_syllable>/prefixed_predicate/predicate_3
	{
		pos3 := pos
		// forethought_cop_pred<predicate>
		if !_accept(parser, _forethought_cop_pred__predicateAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_predicate<verb_syllable>
		if !_accept(parser, _LU_predicate__verb_syllableAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MI_phrase<verb_syllable>
		if !_accept(parser, _MI_phrase__verb_syllableAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// PO_phrase<verb_syllable>
		if !_accept(parser, _PO_phrase__verb_syllableAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// MO_phrase<verb_syllable>
		if !_accept(parser, _MO_phrase__verb_syllableAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// prefixed_predicate
		if !_accept(parser, _prefixed_predicateAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// predicate_3
		if !_accept(parser, _predicate_3Accepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
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
	// forethought_cop_pred<predicate>/LU_predicate<verb_syllable>/MI_phrase<verb_syllable>/PO_phrase<verb_syllable>/MO_phrase<verb_syllable>/prefixed_predicate/predicate_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// forethought_cop_pred<predicate>
		if !_node(parser, _forethought_cop_pred__predicateNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_predicate<verb_syllable>
		if !_node(parser, _LU_predicate__verb_syllableNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_phrase<verb_syllable>
		if !_node(parser, _MI_phrase__verb_syllableNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_phrase<verb_syllable>
		if !_node(parser, _PO_phrase__verb_syllableNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MO_phrase<verb_syllable>
		if !_node(parser, _MO_phrase__verb_syllableNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// prefixed_predicate
		if !_node(parser, _prefixed_predicateNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// predicate_3
		if !_node(parser, _predicate_3Node, node, &pos) {
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
	// forethought_cop_pred<predicate>/LU_predicate<verb_syllable>/MI_phrase<verb_syllable>/PO_phrase<verb_syllable>/MO_phrase<verb_syllable>/prefixed_predicate/predicate_3
	{
		pos3 := pos
		// forethought_cop_pred<predicate>
		if !_fail(parser, _forethought_cop_pred__predicateFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// LU_predicate<verb_syllable>
		if !_fail(parser, _LU_predicate__verb_syllableFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MI_phrase<verb_syllable>
		if !_fail(parser, _MI_phrase__verb_syllableFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// PO_phrase<verb_syllable>
		if !_fail(parser, _PO_phrase__verb_syllableFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// MO_phrase<verb_syllable>
		if !_fail(parser, _MO_phrase__verb_syllableFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// prefixed_predicate
		if !_fail(parser, _prefixed_predicateFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// predicate_3
		if !_fail(parser, _predicate_3Fail, errPos, failure, &pos) {
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

func _predicate_2Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// forethought_cop_pred<predicate>/LU_predicate<verb_syllable>/MI_phrase<verb_syllable>/PO_phrase<verb_syllable>/MO_phrase<verb_syllable>/prefixed_predicate/predicate_3
	{
		pos3 := pos
		var node2 Predicate
		// forethought_cop_pred<predicate>
		if p, n := _forethought_cop_pred__predicateAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// LU_predicate<verb_syllable>
		if p, n := _LU_predicate__verb_syllableAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// MI_phrase<verb_syllable>
		if p, n := _MI_phrase__verb_syllableAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// PO_phrase<verb_syllable>
		if p, n := _PO_phrase__verb_syllableAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// MO_phrase<verb_syllable>
		if p, n := _MO_phrase__verb_syllableAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// prefixed_predicate
		if p, n := _prefixed_predicateAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// predicate_3
		if p, n := _predicate_3Action(parser, pos); n == nil {
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

func _prefixed_predicateAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _prefixed_predicate, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// x:prefix s:_ p:predicate_2
	// x:prefix
	{
		pos1 := pos
		// prefix
		if !_accept(parser, _prefixAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// p:predicate_2
	{
		pos3 := pos
		// predicate_2
		if !_accept(parser, _predicate_2Accepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _prefixed_predicate, start, pos, perr)
fail:
	return _memoize(parser, _prefixed_predicate, start, -1, perr)
}

func _prefixed_predicateNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
	dp := parser.deltaPos[start][_prefixed_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prefixed_predicate}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "prefixed_predicate"}
	// action
	// x:prefix s:_ p:predicate_2
	// x:prefix
	{
		pos1 := pos
		// prefix
		if !_node(parser, _prefixNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// p:predicate_2
	{
		pos3 := pos
		// predicate_2
		if !_node(parser, _predicate_2Node, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prefixed_predicateFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _prefixed_predicate, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prefixed_predicate",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prefixed_predicate}
	// action
	// x:prefix s:_ p:predicate_2
	// x:prefix
	{
		pos1 := pos
		// prefix
		if !_fail(parser, _prefixFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// p:predicate_2
	{
		pos3 := pos
		// predicate_2
		if !_fail(parser, _predicate_2Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prefixed_predicateAction(parser *_Parser, start int) (int, *Predicate) {
	var labels [3]string
	use(labels)
	var label0 Word
	var label1 *Mod
	var label2 Predicate
	dp := parser.deltaPos[start][_prefixed_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prefixed_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// action
	{
		start0 := pos
		// x:prefix s:_ p:predicate_2
		// x:prefix
		{
			pos2 := pos
			// prefix
			if p, n := _prefixAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// p:predicate_2
		{
			pos4 := pos
			// predicate_2
			if p, n := _predicate_2Action(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, p Predicate, s *Mod, x Word) Predicate {
			return Predicate(prefixed(x, s, p))
		}(
			start0, pos, label2, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _predicate_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _predicate_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:predicate_word<verb_syllable> s:_
	// p:predicate_word<verb_syllable>
	{
		pos1 := pos
		// predicate_word<verb_syllable>
		if !_accept(parser, _predicate_word__verb_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	return _memoize(parser, _predicate_3, start, pos, perr)
fail:
	return _memoize(parser, _predicate_3, start, -1, perr)
}

func _predicate_3Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// action
	// p:predicate_word<verb_syllable> s:_
	// p:predicate_word<verb_syllable>
	{
		pos1 := pos
		// predicate_word<verb_syllable>
		if !_node(parser, _predicate_word__verb_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _predicate_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _predicate_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "predicate_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _predicate_3}
	// action
	// p:predicate_word<verb_syllable> s:_
	// p:predicate_word<verb_syllable>
	{
		pos1 := pos
		// predicate_word<verb_syllable>
		if !_fail(parser, _predicate_word__verb_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _predicate_3Action(parser *_Parser, start int) (int, *Predicate) {
	var labels [2]string
	use(labels)
	var label1 *Mod
	var label0 Predicate
	dp := parser.deltaPos[start][_predicate_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _predicate_3}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// action
	{
		start0 := pos
		// p:predicate_word<verb_syllable> s:_
		// p:predicate_word<verb_syllable>
		{
			pos2 := pos
			// predicate_word<verb_syllable>
			if p, n := _predicate_word__verb_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		node = func(
			start, end int, p Predicate, s *Mod) Predicate {
			return Predicate(p.modPredicate(s))
		}(
			start0, pos, label0, label1)
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
	// serial<predicate_1>/predicate_1
	{
		pos3 := pos
		// serial<predicate_1>
		if !_accept(parser, _serial__predicate_1Accepts, &pos, &perr) {
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
	// serial<predicate_1>/predicate_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial<predicate_1>
		if !_node(parser, _serial__predicate_1Node, node, &pos) {
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
	// serial<predicate_1>/predicate_1
	{
		pos3 := pos
		// serial<predicate_1>
		if !_fail(parser, _serial__predicate_1Fail, errPos, failure, &pos) {
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

func _serial_predicateAction(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_serial_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _serial_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// serial<predicate_1>/predicate_1
	{
		pos3 := pos
		var node2 Predicate
		// serial<predicate_1>
		if p, n := _serial__predicate_1Action(parser, pos); n == nil {
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

func _termsAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _terms, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// terms_2/t:term {…}
	{
		pos3 := pos
		// terms_2
		if !_accept(parser, _terms_2Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// t:term
		{
			pos6 := pos
			// term
			if !_accept(parser, _termAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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
	var labels [1]string
	use(labels)
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
	// terms_2/t:term {…}
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
		// action
		// t:term
		{
			pos6 := pos
			// term
			if !_node(parser, _termNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _terms, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms}
	// terms_2/t:term {…}
	{
		pos3 := pos
		// terms_2
		if !_fail(parser, _terms_2Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// t:term
		{
			pos6 := pos
			// term
			if !_fail(parser, _termFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos6:pos]
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

func _termsAction(parser *_Parser, start int) (int, *Terms) {
	var labels [1]string
	use(labels)
	var label0 Term
	dp := parser.deltaPos[start][_terms]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms}
	n := parser.act[key]
	if n != nil {
		n := n.(Terms)
		return start + int(dp-1), &n
	}
	var node Terms
	pos := start
	// terms_2/t:term {…}
	{
		pos3 := pos
		var node2 Terms
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
		// action
		{
			start6 := pos
			// t:term
			{
				pos7 := pos
				// term
				if p, n := _termAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos7:pos]
			}
			node = func(
				start, end int, t Term) Terms {
				return Terms{t}
			}(
				start6, pos, label0)
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
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _terms_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// t:term s:_ ts:terms
	// t:term
	{
		pos1 := pos
		// term
		if !_accept(parser, _termAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms
	{
		pos3 := pos
		// terms
		if !_accept(parser, _termsAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _terms_2, start, pos, perr)
fail:
	return _memoize(parser, _terms_2, start, -1, perr)
}

func _terms_2Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// t:term s:_ ts:terms
	// t:term
	{
		pos1 := pos
		// term
		if !_node(parser, _termNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms
	{
		pos3 := pos
		// terms
		if !_node(parser, _termsNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _terms_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_2}
	// action
	// t:term s:_ ts:terms
	// t:term
	{
		pos1 := pos
		// term
		if !_fail(parser, _termFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms
	{
		pos3 := pos
		// terms
		if !_fail(parser, _termsFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_2Action(parser *_Parser, start int) (int, *Terms) {
	var labels [3]string
	use(labels)
	var label0 Term
	var label1 *Mod
	var label2 Terms
	dp := parser.deltaPos[start][_terms_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Terms)
		return start + int(dp-1), &n
	}
	var node Terms
	pos := start
	// action
	{
		start0 := pos
		// t:term s:_ ts:terms
		// t:term
		{
			pos2 := pos
			// term
			if p, n := _termAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// ts:terms
		{
			pos4 := pos
			// terms
			if p, n := _termsAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, s *Mod, t Term, ts Terms) Terms {
			return Terms(ts.Prepend(t.modTerm(s)))
		}(
			start0, pos, label1, label0, label2)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [4]string
	use(labels)
	if dp, de, ok := _memo(parser, _term, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:linked_argument {…}/x1:adverb {…}/x2:termset {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		// action
		// x0:linked_argument
		{
			pos5 := pos
			// linked_argument
			if !_accept(parser, _linked_argumentAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:adverb
		{
			pos7 := pos
			// adverb
			if !_accept(parser, _adverbAccepts, &pos, &perr) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:termset
		{
			pos9 := pos
			// termset
			if !_accept(parser, _termsetAccepts, &pos, &perr) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:prepositional_phrase
		{
			pos11 := pos
			// prepositional_phrase
			if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _term, start, pos, perr)
fail:
	return _memoize(parser, _term, start, -1, perr)
}

func _termNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [4]string
	use(labels)
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
	// x0:linked_argument {…}/x1:adverb {…}/x2:termset {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:linked_argument
		{
			pos5 := pos
			// linked_argument
			if !_node(parser, _linked_argumentNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x1:adverb
		{
			pos7 := pos
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x2:termset
		{
			pos9 := pos
			// termset
			if !_node(parser, _termsetNode, node, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x3:prepositional_phrase
		{
			pos11 := pos
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
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

func _termFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [4]string
	use(labels)
	pos, failure := _failMemo(parser, _term, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "term",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _term}
	// x0:linked_argument {…}/x1:adverb {…}/x2:termset {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		// action
		// x0:linked_argument
		{
			pos5 := pos
			// linked_argument
			if !_fail(parser, _linked_argumentFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:adverb
		{
			pos7 := pos
			// adverb
			if !_fail(parser, _adverbFail, errPos, failure, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:termset
		{
			pos9 := pos
			// termset
			if !_fail(parser, _termsetFail, errPos, failure, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:prepositional_phrase
		{
			pos11 := pos
			// prepositional_phrase
			if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
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

func _termAction(parser *_Parser, start int) (int, *Term) {
	var labels [4]string
	use(labels)
	var label0 Term
	var label1 Adverb
	var label2 (*TermSet)
	var label3 Preposition
	dp := parser.deltaPos[start][_term]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _term}
	n := parser.act[key]
	if n != nil {
		n := n.(Term)
		return start + int(dp-1), &n
	}
	var node Term
	pos := start
	// x0:linked_argument {…}/x1:adverb {…}/x2:termset {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		var node2 Term
		// action
		{
			start5 := pos
			// x0:linked_argument
			{
				pos6 := pos
				// linked_argument
				if p, n := _linked_argumentAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 Term) Term {
				return Term(x0)
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start8 := pos
			// x1:adverb
			{
				pos9 := pos
				// adverb
				if p, n := _adverbAction(parser, pos); n == nil {
					goto fail7
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos9:pos]
			}
			node = func(
				start, end int, x0 Term, x1 Adverb) Term {
				return Term(x1)
			}(
				start8, pos, label0, label1)
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// action
		{
			start11 := pos
			// x2:termset
			{
				pos12 := pos
				// termset
				if p, n := _termsetAction(parser, pos); n == nil {
					goto fail10
				} else {
					label2 = *n
					pos = p
				}
				labels[2] = parser.text[pos12:pos]
			}
			node = func(
				start, end int, x0 Term, x1 Adverb, x2 *TermSet) Term {
				return Term(x2)
			}(
				start11, pos, label0, label1, label2)
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// action
		{
			start14 := pos
			// x3:prepositional_phrase
			{
				pos15 := pos
				// prepositional_phrase
				if p, n := _prepositional_phraseAction(parser, pos); n == nil {
					goto fail13
				} else {
					label3 = *n
					pos = p
				}
				labels[3] = parser.text[pos15:pos]
			}
			node = func(
				start, end int, x0 Term, x1 Adverb, x2 *TermSet, x3 Preposition) Term {
				return Term(x3)
			}(
				start14, pos, label0, label1, label2, label3)
		}
		goto ok0
	fail13:
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

func _linked_argumentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _linked_argument, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// g:linking_word_spaces? a:argument
	// g:linking_word_spaces?
	{
		pos1 := pos
		// linking_word_spaces?
		{
			pos3 := pos
			// linking_word_spaces
			if !_accept(parser, _linking_word_spacesAccepts, &pos, &perr) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:argument
	{
		pos6 := pos
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	return _memoize(parser, _linked_argument, start, pos, perr)
fail:
	return _memoize(parser, _linked_argument, start, -1, perr)
}

func _linked_argumentNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_linked_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _linked_argument}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "linked_argument"}
	// action
	// g:linking_word_spaces? a:argument
	// g:linking_word_spaces?
	{
		pos1 := pos
		// linking_word_spaces?
		{
			nkids2 := len(node.Kids)
			pos3 := pos
			// linking_word_spaces
			if !_node(parser, _linking_word_spacesNode, node, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			node.Kids = node.Kids[:nkids2]
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:argument
	{
		pos6 := pos
		// argument
		if !_node(parser, _argumentNode, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _linked_argumentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _linked_argument, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "linked_argument",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _linked_argument}
	// action
	// g:linking_word_spaces? a:argument
	// g:linking_word_spaces?
	{
		pos1 := pos
		// linking_word_spaces?
		{
			pos3 := pos
			// linking_word_spaces
			if !_fail(parser, _linking_word_spacesFail, errPos, failure, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:argument
	{
		pos6 := pos
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _linked_argumentAction(parser *_Parser, start int) (int, *Term) {
	var labels [2]string
	use(labels)
	var label0 *(*Word)
	var label1 Argument
	dp := parser.deltaPos[start][_linked_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _linked_argument}
	n := parser.act[key]
	if n != nil {
		n := n.(Term)
		return start + int(dp-1), &n
	}
	var node Term
	pos := start
	// action
	{
		start0 := pos
		// g:linking_word_spaces? a:argument
		// g:linking_word_spaces?
		{
			pos2 := pos
			// linking_word_spaces?
			{
				pos4 := pos
				label0 = new((*Word))
				// linking_word_spaces
				if p, n := _linking_word_spacesAction(parser, pos); n == nil {
					goto fail5
				} else {
					*label0 = *n
					pos = p
				}
				goto ok6
			fail5:
				label0 = nil
				pos = pos4
			ok6:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// a:argument
		{
			pos7 := pos
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, a Argument, g *(*Word)) Term {
			return Term(linkedTerm(g, a))
		}(
			start0, pos, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _linking_word_spacesAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _linking_word_spaces, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// w:linking_word s:_
	// w:linking_word
	{
		pos1 := pos
		// linking_word
		if !_accept(parser, _linking_wordAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	return _memoize(parser, _linking_word_spaces, start, pos, perr)
fail:
	return _memoize(parser, _linking_word_spaces, start, -1, perr)
}

func _linking_word_spacesNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_linking_word_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _linking_word_spaces}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "linking_word_spaces"}
	// action
	// w:linking_word s:_
	// w:linking_word
	{
		pos1 := pos
		// linking_word
		if !_node(parser, _linking_wordNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _linking_word_spacesFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _linking_word_spaces, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "linking_word_spaces",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _linking_word_spaces}
	// action
	// w:linking_word s:_
	// w:linking_word
	{
		pos1 := pos
		// linking_word
		if !_fail(parser, _linking_wordFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _linking_word_spacesAction(parser *_Parser, start int) (int, *(*Word)) {
	var labels [2]string
	use(labels)
	var label0 Word
	var label1 *Mod
	dp := parser.deltaPos[start][_linking_word_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _linking_word_spaces}
	n := parser.act[key]
	if n != nil {
		n := n.((*Word))
		return start + int(dp-1), &n
	}
	var node (*Word)
	pos := start
	// action
	{
		start0 := pos
		// w:linking_word s:_
		// w:linking_word
		{
			pos2 := pos
			// linking_word
			if p, n := _linking_wordAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		node = func(
			start, end int, s *Mod, w Word) *Word {
			return (*Word)(w.mod(s))
		}(
			start0, pos, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _termsetAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [5]string
	use(labels)
	if dp, de, ok := _memo(parser, _termset, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:forethought_cop<terms_V, terms_V> {…}/x1:forethought_cop<terms_IV, terms_IV> {…}/x2:forethought_cop<terms_III, terms_III> {…}/x3:forethought_cop<terms_II, terms_II> {…}/x4:forethought_cop<terms_V, terms_V> {…}
	{
		pos3 := pos
		// action
		// x0:forethought_cop<terms_V, terms_V>
		{
			pos5 := pos
			// forethought_cop<terms_V, terms_V>
			if !_accept(parser, _forethought_cop__terms_V__terms_VAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:forethought_cop<terms_IV, terms_IV>
		{
			pos7 := pos
			// forethought_cop<terms_IV, terms_IV>
			if !_accept(parser, _forethought_cop__terms_IV__terms_IVAccepts, &pos, &perr) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:forethought_cop<terms_III, terms_III>
		{
			pos9 := pos
			// forethought_cop<terms_III, terms_III>
			if !_accept(parser, _forethought_cop__terms_III__terms_IIIAccepts, &pos, &perr) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:forethought_cop<terms_II, terms_II>
		{
			pos11 := pos
			// forethought_cop<terms_II, terms_II>
			if !_accept(parser, _forethought_cop__terms_II__terms_IIAccepts, &pos, &perr) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
		}
		goto ok0
	fail10:
		pos = pos3
		// action
		// x4:forethought_cop<terms_V, terms_V>
		{
			pos13 := pos
			// forethought_cop<terms_V, terms_V>
			if !_accept(parser, _forethought_cop__terms_V__terms_VAccepts, &pos, &perr) {
				goto fail12
			}
			labels[4] = parser.text[pos13:pos]
		}
		goto ok0
	fail12:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _termset, start, pos, perr)
fail:
	return _memoize(parser, _termset, start, -1, perr)
}

func _termsetNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [5]string
	use(labels)
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
	// x0:forethought_cop<terms_V, terms_V> {…}/x1:forethought_cop<terms_IV, terms_IV> {…}/x2:forethought_cop<terms_III, terms_III> {…}/x3:forethought_cop<terms_II, terms_II> {…}/x4:forethought_cop<terms_V, terms_V> {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:forethought_cop<terms_V, terms_V>
		{
			pos5 := pos
			// forethought_cop<terms_V, terms_V>
			if !_node(parser, _forethought_cop__terms_V__terms_VNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x1:forethought_cop<terms_IV, terms_IV>
		{
			pos7 := pos
			// forethought_cop<terms_IV, terms_IV>
			if !_node(parser, _forethought_cop__terms_IV__terms_IVNode, node, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x2:forethought_cop<terms_III, terms_III>
		{
			pos9 := pos
			// forethought_cop<terms_III, terms_III>
			if !_node(parser, _forethought_cop__terms_III__terms_IIINode, node, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x3:forethought_cop<terms_II, terms_II>
		{
			pos11 := pos
			// forethought_cop<terms_II, terms_II>
			if !_node(parser, _forethought_cop__terms_II__terms_IINode, node, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x4:forethought_cop<terms_V, terms_V>
		{
			pos13 := pos
			// forethought_cop<terms_V, terms_V>
			if !_node(parser, _forethought_cop__terms_V__terms_VNode, node, &pos) {
				goto fail12
			}
			labels[4] = parser.text[pos13:pos]
		}
		goto ok0
	fail12:
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
	var labels [5]string
	use(labels)
	pos, failure := _failMemo(parser, _termset, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "termset",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _termset}
	// x0:forethought_cop<terms_V, terms_V> {…}/x1:forethought_cop<terms_IV, terms_IV> {…}/x2:forethought_cop<terms_III, terms_III> {…}/x3:forethought_cop<terms_II, terms_II> {…}/x4:forethought_cop<terms_V, terms_V> {…}
	{
		pos3 := pos
		// action
		// x0:forethought_cop<terms_V, terms_V>
		{
			pos5 := pos
			// forethought_cop<terms_V, terms_V>
			if !_fail(parser, _forethought_cop__terms_V__terms_VFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:forethought_cop<terms_IV, terms_IV>
		{
			pos7 := pos
			// forethought_cop<terms_IV, terms_IV>
			if !_fail(parser, _forethought_cop__terms_IV__terms_IVFail, errPos, failure, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:forethought_cop<terms_III, terms_III>
		{
			pos9 := pos
			// forethought_cop<terms_III, terms_III>
			if !_fail(parser, _forethought_cop__terms_III__terms_IIIFail, errPos, failure, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:forethought_cop<terms_II, terms_II>
		{
			pos11 := pos
			// forethought_cop<terms_II, terms_II>
			if !_fail(parser, _forethought_cop__terms_II__terms_IIFail, errPos, failure, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
		}
		goto ok0
	fail10:
		pos = pos3
		// action
		// x4:forethought_cop<terms_V, terms_V>
		{
			pos13 := pos
			// forethought_cop<terms_V, terms_V>
			if !_fail(parser, _forethought_cop__terms_V__terms_VFail, errPos, failure, &pos) {
				goto fail12
			}
			labels[4] = parser.text[pos13:pos]
		}
		goto ok0
	fail12:
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

func _termsetAction(parser *_Parser, start int) (int, *(*TermSet)) {
	var labels [5]string
	use(labels)
	var label4 (*CoP)
	var label0 (*CoP)
	var label1 (*CoP)
	var label2 (*CoP)
	var label3 (*CoP)
	dp := parser.deltaPos[start][_termset]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _termset}
	n := parser.act[key]
	if n != nil {
		n := n.((*TermSet))
		return start + int(dp-1), &n
	}
	var node (*TermSet)
	pos := start
	// x0:forethought_cop<terms_V, terms_V> {…}/x1:forethought_cop<terms_IV, terms_IV> {…}/x2:forethought_cop<terms_III, terms_III> {…}/x3:forethought_cop<terms_II, terms_II> {…}/x4:forethought_cop<terms_V, terms_V> {…}
	{
		pos3 := pos
		var node2 (*TermSet)
		// action
		{
			start5 := pos
			// x0:forethought_cop<terms_V, terms_V>
			{
				pos6 := pos
				// forethought_cop<terms_V, terms_V>
				if p, n := _forethought_cop__terms_V__terms_VAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 *CoP) *TermSet {
				return (*TermSet)(x0)
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start8 := pos
			// x1:forethought_cop<terms_IV, terms_IV>
			{
				pos9 := pos
				// forethought_cop<terms_IV, terms_IV>
				if p, n := _forethought_cop__terms_IV__terms_IVAction(parser, pos); n == nil {
					goto fail7
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos9:pos]
			}
			node = func(
				start, end int, x0 *CoP, x1 *CoP) *TermSet {
				return (*TermSet)(x1)
			}(
				start8, pos, label0, label1)
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// action
		{
			start11 := pos
			// x2:forethought_cop<terms_III, terms_III>
			{
				pos12 := pos
				// forethought_cop<terms_III, terms_III>
				if p, n := _forethought_cop__terms_III__terms_IIIAction(parser, pos); n == nil {
					goto fail10
				} else {
					label2 = *n
					pos = p
				}
				labels[2] = parser.text[pos12:pos]
			}
			node = func(
				start, end int, x0 *CoP, x1 *CoP, x2 *CoP) *TermSet {
				return (*TermSet)(x2)
			}(
				start11, pos, label0, label1, label2)
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// action
		{
			start14 := pos
			// x3:forethought_cop<terms_II, terms_II>
			{
				pos15 := pos
				// forethought_cop<terms_II, terms_II>
				if p, n := _forethought_cop__terms_II__terms_IIAction(parser, pos); n == nil {
					goto fail13
				} else {
					label3 = *n
					pos = p
				}
				labels[3] = parser.text[pos15:pos]
			}
			node = func(
				start, end int, x0 *CoP, x1 *CoP, x2 *CoP, x3 *CoP) *TermSet {
				return (*TermSet)(x3)
			}(
				start14, pos, label0, label1, label2, label3)
		}
		goto ok0
	fail13:
		node = node2
		pos = pos3
		// action
		{
			start17 := pos
			// x4:forethought_cop<terms_V, terms_V>
			{
				pos18 := pos
				// forethought_cop<terms_V, terms_V>
				if p, n := _forethought_cop__terms_V__terms_VAction(parser, pos); n == nil {
					goto fail16
				} else {
					label4 = *n
					pos = p
				}
				labels[4] = parser.text[pos18:pos]
			}
			node = func(
				start, end int, x0 *CoP, x1 *CoP, x2 *CoP, x3 *CoP, x4 *CoP) *TermSet {
				return (*TermSet)(x4)
			}(
				start17, pos, label0, label1, label2, label3, label4)
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

func _terms_IIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _terms_II, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// l:term s:_ r:term
	// l:term
	{
		pos1 := pos
		// term
		if !_accept(parser, _termAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:term
	{
		pos3 := pos
		// term
		if !_accept(parser, _termAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _terms_II, start, pos, perr)
fail:
	return _memoize(parser, _terms_II, start, -1, perr)
}

func _terms_IINode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// l:term s:_ r:term
	// l:term
	{
		pos1 := pos
		// term
		if !_node(parser, _termNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:term
	{
		pos3 := pos
		// term
		if !_node(parser, _termNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_IIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _terms_II, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_II",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_II}
	// action
	// l:term s:_ r:term
	// l:term
	{
		pos1 := pos
		// term
		if !_fail(parser, _termFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:term
	{
		pos3 := pos
		// term
		if !_fail(parser, _termFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_IIAction(parser *_Parser, start int) (int, *Terms) {
	var labels [3]string
	use(labels)
	var label1 *Mod
	var label2 Term
	var label0 Term
	dp := parser.deltaPos[start][_terms_II]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_II}
	n := parser.act[key]
	if n != nil {
		n := n.(Terms)
		return start + int(dp-1), &n
	}
	var node Terms
	pos := start
	// action
	{
		start0 := pos
		// l:term s:_ r:term
		// l:term
		{
			pos2 := pos
			// term
			if p, n := _termAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// r:term
		{
			pos4 := pos
			// term
			if p, n := _termAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, l Term, r Term, s *Mod) Terms {
			return Terms{l.modTerm(s), r}
		}(
			start0, pos, label0, label2, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_IIIAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _terms_III, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// l:term s:_ r:terms_II
	// l:term
	{
		pos1 := pos
		// term
		if !_accept(parser, _termAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_II
	{
		pos3 := pos
		// terms_II
		if !_accept(parser, _terms_IIAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _terms_III, start, pos, perr)
fail:
	return _memoize(parser, _terms_III, start, -1, perr)
}

func _terms_IIINode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// l:term s:_ r:terms_II
	// l:term
	{
		pos1 := pos
		// term
		if !_node(parser, _termNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_II
	{
		pos3 := pos
		// terms_II
		if !_node(parser, _terms_IINode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_IIIFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _terms_III, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_III",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_III}
	// action
	// l:term s:_ r:terms_II
	// l:term
	{
		pos1 := pos
		// term
		if !_fail(parser, _termFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_II
	{
		pos3 := pos
		// terms_II
		if !_fail(parser, _terms_IIFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_IIIAction(parser *_Parser, start int) (int, *Terms) {
	var labels [3]string
	use(labels)
	var label0 Term
	var label1 *Mod
	var label2 Terms
	dp := parser.deltaPos[start][_terms_III]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_III}
	n := parser.act[key]
	if n != nil {
		n := n.(Terms)
		return start + int(dp-1), &n
	}
	var node Terms
	pos := start
	// action
	{
		start0 := pos
		// l:term s:_ r:terms_II
		// l:term
		{
			pos2 := pos
			// term
			if p, n := _termAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// r:terms_II
		{
			pos4 := pos
			// terms_II
			if p, n := _terms_IIAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, l Term, r Terms, s *Mod) Terms {
			return Terms(r.Prepend(l.modTerm(s)))
		}(
			start0, pos, label0, label2, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_IVAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _terms_IV, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// l:term s:_ r:terms_III
	// l:term
	{
		pos1 := pos
		// term
		if !_accept(parser, _termAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_III
	{
		pos3 := pos
		// terms_III
		if !_accept(parser, _terms_IIIAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _terms_IV, start, pos, perr)
fail:
	return _memoize(parser, _terms_IV, start, -1, perr)
}

func _terms_IVNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// l:term s:_ r:terms_III
	// l:term
	{
		pos1 := pos
		// term
		if !_node(parser, _termNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_III
	{
		pos3 := pos
		// terms_III
		if !_node(parser, _terms_IIINode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_IVFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _terms_IV, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_IV",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_IV}
	// action
	// l:term s:_ r:terms_III
	// l:term
	{
		pos1 := pos
		// term
		if !_fail(parser, _termFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_III
	{
		pos3 := pos
		// terms_III
		if !_fail(parser, _terms_IIIFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_IVAction(parser *_Parser, start int) (int, *Terms) {
	var labels [3]string
	use(labels)
	var label0 Term
	var label1 *Mod
	var label2 Terms
	dp := parser.deltaPos[start][_terms_IV]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_IV}
	n := parser.act[key]
	if n != nil {
		n := n.(Terms)
		return start + int(dp-1), &n
	}
	var node Terms
	pos := start
	// action
	{
		start0 := pos
		// l:term s:_ r:terms_III
		// l:term
		{
			pos2 := pos
			// term
			if p, n := _termAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// r:terms_III
		{
			pos4 := pos
			// terms_III
			if p, n := _terms_IIIAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, l Term, r Terms, s *Mod) Terms {
			return Terms(r.Prepend(l.modTerm(s)))
		}(
			start0, pos, label0, label2, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _terms_VAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _terms_V, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// l:term s:_ r:terms_IV
	// l:term
	{
		pos1 := pos
		// term
		if !_accept(parser, _termAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_IV
	{
		pos3 := pos
		// terms_IV
		if !_accept(parser, _terms_IVAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _terms_V, start, pos, perr)
fail:
	return _memoize(parser, _terms_V, start, -1, perr)
}

func _terms_VNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// l:term s:_ r:terms_IV
	// l:term
	{
		pos1 := pos
		// term
		if !_node(parser, _termNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_IV
	{
		pos3 := pos
		// terms_IV
		if !_node(parser, _terms_IVNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _terms_VFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _terms_V, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "terms_V",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _terms_V}
	// action
	// l:term s:_ r:terms_IV
	// l:term
	{
		pos1 := pos
		// term
		if !_fail(parser, _termFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// r:terms_IV
	{
		pos3 := pos
		// terms_IV
		if !_fail(parser, _terms_IVFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _terms_VAction(parser *_Parser, start int) (int, *Terms) {
	var labels [3]string
	use(labels)
	var label2 Terms
	var label0 Term
	var label1 *Mod
	dp := parser.deltaPos[start][_terms_V]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _terms_V}
	n := parser.act[key]
	if n != nil {
		n := n.(Terms)
		return start + int(dp-1), &n
	}
	var node Terms
	pos := start
	// action
	{
		start0 := pos
		// l:term s:_ r:terms_IV
		// l:term
		{
			pos2 := pos
			// term
			if p, n := _termAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// r:terms_IV
		{
			pos4 := pos
			// terms_IV
			if p, n := _terms_IVAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, l Term, r Terms, s *Mod) Terms {
			return Terms(r.Prepend(l.modTerm(s)))
		}(
			start0, pos, label0, label2, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _argumentAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _argument, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:afterthought_cop<arg_1, argument> {…}/arg_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<arg_1, argument>
		{
			pos5 := pos
			// afterthought_cop<arg_1, argument>
			if !_accept(parser, _afterthought_cop__arg_1__argumentAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_1
		if !_accept(parser, _arg_1Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _argument, start, pos, perr)
fail:
	return _memoize(parser, _argument, start, -1, perr)
}

func _argumentNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:afterthought_cop<arg_1, argument> {…}/arg_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:afterthought_cop<arg_1, argument>
		{
			pos5 := pos
			// afterthought_cop<arg_1, argument>
			if !_node(parser, _afterthought_cop__arg_1__argumentNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_1
		if !_node(parser, _arg_1Node, node, &pos) {
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

func _argumentFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _argument, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "argument",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _argument}
	// x:afterthought_cop<arg_1, argument> {…}/arg_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<arg_1, argument>
		{
			pos5 := pos
			// afterthought_cop<arg_1, argument>
			if !_fail(parser, _afterthought_cop__arg_1__argumentFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_1
		if !_fail(parser, _arg_1Fail, errPos, failure, &pos) {
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

func _argumentAction(parser *_Parser, start int) (int, *Argument) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_argument]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _argument}
	n := parser.act[key]
	if n != nil {
		n := n.(Argument)
		return start + int(dp-1), &n
	}
	var node Argument
	pos := start
	// x:afterthought_cop<arg_1, argument> {…}/arg_1
	{
		pos3 := pos
		var node2 Argument
		// action
		{
			start5 := pos
			// x:afterthought_cop<arg_1, argument>
			{
				pos6 := pos
				// afterthought_cop<arg_1, argument>
				if p, n := _afterthought_cop__arg_1__argumentAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Argument {
				return Argument((*CoPArgument)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// arg_1
		if p, n := _arg_1Action(parser, pos); n == nil {
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

func _arg_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _arg_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:forethought_cop<argument, argument> {…}/arg_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<argument, argument>
		{
			pos5 := pos
			// forethought_cop<argument, argument>
			if !_accept(parser, _forethought_cop__argument__argumentAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_2
		if !_accept(parser, _arg_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _arg_1, start, pos, perr)
fail:
	return _memoize(parser, _arg_1, start, -1, perr)
}

func _arg_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:forethought_cop<argument, argument> {…}/arg_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:forethought_cop<argument, argument>
		{
			pos5 := pos
			// forethought_cop<argument, argument>
			if !_node(parser, _forethought_cop__argument__argumentNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_2
		if !_node(parser, _arg_2Node, node, &pos) {
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

func _arg_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _arg_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_1}
	// x:forethought_cop<argument, argument> {…}/arg_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<argument, argument>
		{
			pos5 := pos
			// forethought_cop<argument, argument>
			if !_fail(parser, _forethought_cop__argument__argumentFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_2
		if !_fail(parser, _arg_2Fail, errPos, failure, &pos) {
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

func _arg_1Action(parser *_Parser, start int) (int, *Argument) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_arg_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Argument)
		return start + int(dp-1), &n
	}
	var node Argument
	pos := start
	// x:forethought_cop<argument, argument> {…}/arg_2
	{
		pos3 := pos
		var node2 Argument
		// action
		{
			start5 := pos
			// x:forethought_cop<argument, argument>
			{
				pos6 := pos
				// forethought_cop<argument, argument>
				if p, n := _forethought_cop__argument__argumentAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Argument {
				return Argument((*CoPArgument)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// arg_2
		if p, n := _arg_2Action(parser, pos); n == nil {
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

func _arg_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _arg_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// ku:focus_spaces? a:arg_3
	// ku:focus_spaces?
	{
		pos1 := pos
		// focus_spaces?
		{
			pos3 := pos
			// focus_spaces
			if !_accept(parser, _focus_spacesAccepts, &pos, &perr) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:arg_3
	{
		pos6 := pos
		// arg_3
		if !_accept(parser, _arg_3Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	return _memoize(parser, _arg_2, start, pos, perr)
fail:
	return _memoize(parser, _arg_2, start, -1, perr)
}

func _arg_2Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// action
	// ku:focus_spaces? a:arg_3
	// ku:focus_spaces?
	{
		pos1 := pos
		// focus_spaces?
		{
			nkids2 := len(node.Kids)
			pos3 := pos
			// focus_spaces
			if !_node(parser, _focus_spacesNode, node, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			node.Kids = node.Kids[:nkids2]
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:arg_3
	{
		pos6 := pos
		// arg_3
		if !_node(parser, _arg_3Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _arg_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_2}
	// action
	// ku:focus_spaces? a:arg_3
	// ku:focus_spaces?
	{
		pos1 := pos
		// focus_spaces?
		{
			pos3 := pos
			// focus_spaces
			if !_fail(parser, _focus_spacesFail, errPos, failure, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:arg_3
	{
		pos6 := pos
		// arg_3
		if !_fail(parser, _arg_3Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_2Action(parser *_Parser, start int) (int, *Argument) {
	var labels [2]string
	use(labels)
	var label0 *Word
	var label1 PredicateArgument
	dp := parser.deltaPos[start][_arg_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Argument)
		return start + int(dp-1), &n
	}
	var node Argument
	pos := start
	// action
	{
		start0 := pos
		// ku:focus_spaces? a:arg_3
		// ku:focus_spaces?
		{
			pos2 := pos
			// focus_spaces?
			{
				pos4 := pos
				label0 = new(Word)
				// focus_spaces
				if p, n := _focus_spacesAction(parser, pos); n == nil {
					goto fail5
				} else {
					*label0 = *n
					pos = p
				}
				goto ok6
			fail5:
				label0 = nil
				pos = pos4
			ok6:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// a:arg_3
		{
			pos7 := pos
			// arg_3
			if p, n := _arg_3Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, a PredicateArgument, ku *Word) Argument {
			return Argument(focusedArgument(ku, a))
		}(
			start0, pos, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _focus_spacesAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _focus_spaces, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// ku:focus s:_
	// ku:focus
	{
		pos1 := pos
		// focus
		if !_accept(parser, _focusAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	return _memoize(parser, _focus_spaces, start, pos, perr)
fail:
	return _memoize(parser, _focus_spaces, start, -1, perr)
}

func _focus_spacesNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_focus_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _focus_spaces}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "focus_spaces"}
	// action
	// ku:focus s:_
	// ku:focus
	{
		pos1 := pos
		// focus
		if !_node(parser, _focusNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _focus_spacesFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _focus_spaces, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "focus_spaces",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _focus_spaces}
	// action
	// ku:focus s:_
	// ku:focus
	{
		pos1 := pos
		// focus
		if !_fail(parser, _focusFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _focus_spacesAction(parser *_Parser, start int) (int, *Word) {
	var labels [2]string
	use(labels)
	var label0 Word
	var label1 *Mod
	dp := parser.deltaPos[start][_focus_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _focus_spaces}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// ku:focus s:_
		// ku:focus
		{
			pos2 := pos
			// focus
			if p, n := _focusAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		node = func(
			start, end int, ku Word, s *Mod) Word {
			return Word(*ku.mod(s))
		}(
			start0, pos, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _arg_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// q:quantifier_spaces? a:arg_4
	// q:quantifier_spaces?
	{
		pos1 := pos
		// quantifier_spaces?
		{
			pos3 := pos
			// quantifier_spaces
			if !_accept(parser, _quantifier_spacesAccepts, &pos, &perr) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:arg_4
	{
		pos6 := pos
		// arg_4
		if !_accept(parser, _arg_4Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	return _memoize(parser, _arg_3, start, pos, perr)
fail:
	return _memoize(parser, _arg_3, start, -1, perr)
}

func _arg_3Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// action
	// q:quantifier_spaces? a:arg_4
	// q:quantifier_spaces?
	{
		pos1 := pos
		// quantifier_spaces?
		{
			nkids2 := len(node.Kids)
			pos3 := pos
			// quantifier_spaces
			if !_node(parser, _quantifier_spacesNode, node, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			node.Kids = node.Kids[:nkids2]
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:arg_4
	{
		pos6 := pos
		// arg_4
		if !_node(parser, _arg_4Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _arg_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_3}
	// action
	// q:quantifier_spaces? a:arg_4
	// q:quantifier_spaces?
	{
		pos1 := pos
		// quantifier_spaces?
		{
			pos3 := pos
			// quantifier_spaces
			if !_fail(parser, _quantifier_spacesFail, errPos, failure, &pos) {
				goto fail4
			}
			goto ok5
		fail4:
			pos = pos3
		ok5:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// a:arg_4
	{
		pos6 := pos
		// arg_4
		if !_fail(parser, _arg_4Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_3Action(parser *_Parser, start int) (int, *PredicateArgument) {
	var labels [2]string
	use(labels)
	var label0 *Word
	var label1 PredicateArgument
	dp := parser.deltaPos[start][_arg_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_3}
	n := parser.act[key]
	if n != nil {
		n := n.(PredicateArgument)
		return start + int(dp-1), &n
	}
	var node PredicateArgument
	pos := start
	// action
	{
		start0 := pos
		// q:quantifier_spaces? a:arg_4
		// q:quantifier_spaces?
		{
			pos2 := pos
			// quantifier_spaces?
			{
				pos4 := pos
				label0 = new(Word)
				// quantifier_spaces
				if p, n := _quantifier_spacesAction(parser, pos); n == nil {
					goto fail5
				} else {
					*label0 = *n
					pos = p
				}
				goto ok6
			fail5:
				label0 = nil
				pos = pos4
			ok6:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// a:arg_4
		{
			pos7 := pos
			// arg_4
			if p, n := _arg_4Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, a PredicateArgument, q *Word) PredicateArgument {
			return PredicateArgument(quantifiedArgument(q, a))
		}(
			start0, pos, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quantifier_spacesAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _quantifier_spaces, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// q:quantifier s:_
	// q:quantifier
	{
		pos1 := pos
		// quantifier
		if !_accept(parser, _quantifierAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	return _memoize(parser, _quantifier_spaces, start, pos, perr)
fail:
	return _memoize(parser, _quantifier_spaces, start, -1, perr)
}

func _quantifier_spacesNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_quantifier_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quantifier_spaces}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "quantifier_spaces"}
	// action
	// q:quantifier s:_
	// q:quantifier
	{
		pos1 := pos
		// quantifier
		if !_node(parser, _quantifierNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quantifier_spacesFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _quantifier_spaces, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quantifier_spaces",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quantifier_spaces}
	// action
	// q:quantifier s:_
	// q:quantifier
	{
		pos1 := pos
		// quantifier
		if !_fail(parser, _quantifierFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quantifier_spacesAction(parser *_Parser, start int) (int, *Word) {
	var labels [2]string
	use(labels)
	var label0 Word
	var label1 *Mod
	dp := parser.deltaPos[start][_quantifier_spaces]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quantifier_spaces}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// q:quantifier s:_
		// q:quantifier
		{
			pos2 := pos
			// quantifier
			if p, n := _quantifierAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		node = func(
			start, end int, q Word, s *Mod) Word {
			return Word(*q.mod(s))
		}(
			start0, pos, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_4Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _arg_4, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// a:arg_5 s:_ rc:relative_clause?
	// a:arg_5
	{
		pos1 := pos
		// arg_5
		if !_accept(parser, _arg_5Accepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// rc:relative_clause?
	{
		pos3 := pos
		// relative_clause?
		{
			pos5 := pos
			// relative_clause
			if !_accept(parser, _relative_clauseAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _arg_4, start, pos, perr)
fail:
	return _memoize(parser, _arg_4, start, -1, perr)
}

func _arg_4Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// a:arg_5 s:_ rc:relative_clause?
	// a:arg_5
	{
		pos1 := pos
		// arg_5
		if !_node(parser, _arg_5Node, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// rc:relative_clause?
	{
		pos3 := pos
		// relative_clause?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// relative_clause
			if !_node(parser, _relative_clauseNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_4Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _arg_4, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_4",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_4}
	// action
	// a:arg_5 s:_ rc:relative_clause?
	// a:arg_5
	{
		pos1 := pos
		// arg_5
		if !_fail(parser, _arg_5Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// rc:relative_clause?
	{
		pos3 := pos
		// relative_clause?
		{
			pos5 := pos
			// relative_clause
			if !_fail(parser, _relative_clauseFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_4Action(parser *_Parser, start int) (int, *PredicateArgument) {
	var labels [3]string
	use(labels)
	var label0 *PredicateArgument
	var label1 *Mod
	var label2 *Relative
	dp := parser.deltaPos[start][_arg_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_4}
	n := parser.act[key]
	if n != nil {
		n := n.(PredicateArgument)
		return start + int(dp-1), &n
	}
	var node PredicateArgument
	pos := start
	// action
	{
		start0 := pos
		// a:arg_5 s:_ rc:relative_clause?
		// a:arg_5
		{
			pos2 := pos
			// arg_5
			if p, n := _arg_5Action(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// rc:relative_clause?
		{
			pos4 := pos
			// relative_clause?
			{
				pos6 := pos
				label2 = new(Relative)
				// relative_clause
				if p, n := _relative_clauseAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, a *PredicateArgument, rc *Relative, s *Mod) PredicateArgument {
			return PredicateArgument(argumentRelative(*a, s, rc))
		}(
			start0, pos, label0, label2, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_5Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _arg_5, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:(serial<arg_6>/arg_6)
	{
		pos0 := pos
		// (serial<arg_6>/arg_6)
		// serial<arg_6>/arg_6
		{
			pos4 := pos
			// serial<arg_6>
			if !_accept(parser, _serial__arg_6Accepts, &pos, &perr) {
				goto fail5
			}
			goto ok1
		fail5:
			pos = pos4
			// arg_6
			if !_accept(parser, _arg_6Accepts, &pos, &perr) {
				goto fail6
			}
			goto ok1
		fail6:
			pos = pos4
			goto fail
		ok1:
		}
		labels[0] = parser.text[pos0:pos]
	}
	return _memoize(parser, _arg_5, start, pos, perr)
fail:
	return _memoize(parser, _arg_5, start, -1, perr)
}

func _arg_5Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// p:(serial<arg_6>/arg_6)
	{
		pos0 := pos
		// (serial<arg_6>/arg_6)
		{
			nkids1 := len(node.Kids)
			pos02 := pos
			// serial<arg_6>/arg_6
			{
				pos6 := pos
				nkids4 := len(node.Kids)
				// serial<arg_6>
				if !_node(parser, _serial__arg_6Node, node, &pos) {
					goto fail7
				}
				goto ok3
			fail7:
				node.Kids = node.Kids[:nkids4]
				pos = pos6
				// arg_6
				if !_node(parser, _arg_6Node, node, &pos) {
					goto fail8
				}
				goto ok3
			fail8:
				node.Kids = node.Kids[:nkids4]
				pos = pos6
				goto fail
			ok3:
			}
			sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
			node.Kids = append(node.Kids[:nkids1], sub)
		}
		labels[0] = parser.text[pos0:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_5Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _arg_5, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_5",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_5}
	// action
	// p:(serial<arg_6>/arg_6)
	{
		pos0 := pos
		// (serial<arg_6>/arg_6)
		// serial<arg_6>/arg_6
		{
			pos4 := pos
			// serial<arg_6>
			if !_fail(parser, _serial__arg_6Fail, errPos, failure, &pos) {
				goto fail5
			}
			goto ok1
		fail5:
			pos = pos4
			// arg_6
			if !_fail(parser, _arg_6Fail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok1
		fail6:
			pos = pos4
			goto fail
		ok1:
		}
		labels[0] = parser.text[pos0:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_5Action(parser *_Parser, start int) (int, **PredicateArgument) {
	var labels [1]string
	use(labels)
	var label0 Predicate
	dp := parser.deltaPos[start][_arg_5]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_5}
	n := parser.act[key]
	if n != nil {
		n := n.(*PredicateArgument)
		return start + int(dp-1), &n
	}
	var node *PredicateArgument
	pos := start
	// action
	{
		start0 := pos
		// p:(serial<arg_6>/arg_6)
		{
			pos1 := pos
			// (serial<arg_6>/arg_6)
			// serial<arg_6>/arg_6
			{
				pos5 := pos
				var node4 Predicate
				// serial<arg_6>
				if p, n := _serial__arg_6Action(parser, pos); n == nil {
					goto fail6
				} else {
					label0 = *n
					pos = p
				}
				goto ok2
			fail6:
				label0 = node4
				pos = pos5
				// arg_6
				if p, n := _arg_6Action(parser, pos); n == nil {
					goto fail7
				} else {
					label0 = *n
					pos = p
				}
				goto ok2
			fail7:
				label0 = node4
				pos = pos5
				goto fail
			ok2:
			}
			labels[0] = parser.text[pos1:pos]
		}
		node = func(
			start, end int, p Predicate) *PredicateArgument {
			return &PredicateArgument{Predicate: p}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_6Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _arg_6, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:content_clause {…}/afterthought_cop_pred<arg_7>/forethought_cop_pred<arg_5>/LU_predicate<arg_syllable>/MI_phrase<arg_syllable>/PO_phrase<arg_syllable>/MO_phrase<arg_syllable>/arg_7
	{
		pos3 := pos
		// action
		// x0:content_clause
		{
			pos5 := pos
			// content_clause
			if !_accept(parser, _content_clauseAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// afterthought_cop_pred<arg_7>
		if !_accept(parser, _afterthought_cop_pred__arg_7Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// forethought_cop_pred<arg_5>
		if !_accept(parser, _forethought_cop_pred__arg_5Accepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// LU_predicate<arg_syllable>
		if !_accept(parser, _LU_predicate__arg_syllableAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// MI_phrase<arg_syllable>
		if !_accept(parser, _MI_phrase__arg_syllableAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// PO_phrase<arg_syllable>
		if !_accept(parser, _PO_phrase__arg_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// MO_phrase<arg_syllable>
		if !_accept(parser, _MO_phrase__arg_syllableAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		// arg_7
		if !_accept(parser, _arg_7Accepts, &pos, &perr) {
			goto fail12
		}
		goto ok0
	fail12:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _arg_6, start, pos, perr)
fail:
	return _memoize(parser, _arg_6, start, -1, perr)
}

func _arg_6Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x0:content_clause {…}/afterthought_cop_pred<arg_7>/forethought_cop_pred<arg_5>/LU_predicate<arg_syllable>/MI_phrase<arg_syllable>/PO_phrase<arg_syllable>/MO_phrase<arg_syllable>/arg_7
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:content_clause
		{
			pos5 := pos
			// content_clause
			if !_node(parser, _content_clauseNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// afterthought_cop_pred<arg_7>
		if !_node(parser, _afterthought_cop_pred__arg_7Node, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_cop_pred<arg_5>
		if !_node(parser, _forethought_cop_pred__arg_5Node, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_predicate<arg_syllable>
		if !_node(parser, _LU_predicate__arg_syllableNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_phrase<arg_syllable>
		if !_node(parser, _MI_phrase__arg_syllableNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_phrase<arg_syllable>
		if !_node(parser, _PO_phrase__arg_syllableNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MO_phrase<arg_syllable>
		if !_node(parser, _MO_phrase__arg_syllableNode, node, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_7
		if !_node(parser, _arg_7Node, node, &pos) {
			goto fail12
		}
		goto ok0
	fail12:
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
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _arg_6, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_6",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_6}
	// x0:content_clause {…}/afterthought_cop_pred<arg_7>/forethought_cop_pred<arg_5>/LU_predicate<arg_syllable>/MI_phrase<arg_syllable>/PO_phrase<arg_syllable>/MO_phrase<arg_syllable>/arg_7
	{
		pos3 := pos
		// action
		// x0:content_clause
		{
			pos5 := pos
			// content_clause
			if !_fail(parser, _content_clauseFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// afterthought_cop_pred<arg_7>
		if !_fail(parser, _afterthought_cop_pred__arg_7Fail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// forethought_cop_pred<arg_5>
		if !_fail(parser, _forethought_cop_pred__arg_5Fail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// LU_predicate<arg_syllable>
		if !_fail(parser, _LU_predicate__arg_syllableFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// MI_phrase<arg_syllable>
		if !_fail(parser, _MI_phrase__arg_syllableFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// PO_phrase<arg_syllable>
		if !_fail(parser, _PO_phrase__arg_syllableFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// MO_phrase<arg_syllable>
		if !_fail(parser, _MO_phrase__arg_syllableFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		// arg_7
		if !_fail(parser, _arg_7Fail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok0
	fail12:
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

func _arg_6Action(parser *_Parser, start int) (int, *Predicate) {
	var labels [1]string
	use(labels)
	var label0 Content
	dp := parser.deltaPos[start][_arg_6]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_6}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// x0:content_clause {…}/afterthought_cop_pred<arg_7>/forethought_cop_pred<arg_5>/LU_predicate<arg_syllable>/MI_phrase<arg_syllable>/PO_phrase<arg_syllable>/MO_phrase<arg_syllable>/arg_7
	{
		pos3 := pos
		var node2 Predicate
		// action
		{
			start5 := pos
			// x0:content_clause
			{
				pos6 := pos
				// content_clause
				if p, n := _content_clauseAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 Content) Predicate {
				return Predicate(x0)
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// afterthought_cop_pred<arg_7>
		if p, n := _afterthought_cop_pred__arg_7Action(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// forethought_cop_pred<arg_5>
		if p, n := _forethought_cop_pred__arg_5Action(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// LU_predicate<arg_syllable>
		if p, n := _LU_predicate__arg_syllableAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// MI_phrase<arg_syllable>
		if p, n := _MI_phrase__arg_syllableAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// PO_phrase<arg_syllable>
		if p, n := _PO_phrase__arg_syllableAction(parser, pos); n == nil {
			goto fail11
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail11:
		node = node2
		pos = pos3
		// MO_phrase<arg_syllable>
		if p, n := _MO_phrase__arg_syllableAction(parser, pos); n == nil {
			goto fail12
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail12:
		node = node2
		pos = pos3
		// arg_7
		if p, n := _arg_7Action(parser, pos); n == nil {
			goto fail13
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail13:
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
	// predicate_word<arg_syllable>
	if !_accept(parser, _predicate_word__arg_syllableAccepts, &pos, &perr) {
		goto fail
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
	// predicate_word<arg_syllable>
	if !_node(parser, _predicate_word__arg_syllableNode, node, &pos) {
		goto fail
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
	// predicate_word<arg_syllable>
	if !_fail(parser, _predicate_word__arg_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_7Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_arg_7]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_7}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// predicate_word<arg_syllable>
	if p, n := _predicate_word__arg_syllableAction(parser, pos); n == nil {
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

func _relative_clauseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _relative_clause, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:LU_phrase<relative_syllable> {…}/relative_clause_1
	{
		pos3 := pos
		// action
		// x0:afterthought_cop<relative_clause_1, relative_clause_1>
		{
			pos5 := pos
			// afterthought_cop<relative_clause_1, relative_clause_1>
			if !_accept(parser, _afterthought_cop__relative_clause_1__relative_clause_1Accepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:LU_phrase<relative_syllable>
		{
			pos7 := pos
			// LU_phrase<relative_syllable>
			if !_accept(parser, _LU_phrase__relative_syllableAccepts, &pos, &perr) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// relative_clause_1
		if !_accept(parser, _relative_clause_1Accepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_clause, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause, start, -1, perr)
}

func _relative_clauseNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:LU_phrase<relative_syllable> {…}/relative_clause_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:afterthought_cop<relative_clause_1, relative_clause_1>
		{
			pos5 := pos
			// afterthought_cop<relative_clause_1, relative_clause_1>
			if !_node(parser, _afterthought_cop__relative_clause_1__relative_clause_1Node, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x1:LU_phrase<relative_syllable>
		{
			pos7 := pos
			// LU_phrase<relative_syllable>
			if !_node(parser, _LU_phrase__relative_syllableNode, node, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_clause_1
		if !_node(parser, _relative_clause_1Node, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
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
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _relative_clause, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause}
	// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:LU_phrase<relative_syllable> {…}/relative_clause_1
	{
		pos3 := pos
		// action
		// x0:afterthought_cop<relative_clause_1, relative_clause_1>
		{
			pos5 := pos
			// afterthought_cop<relative_clause_1, relative_clause_1>
			if !_fail(parser, _afterthought_cop__relative_clause_1__relative_clause_1Fail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:LU_phrase<relative_syllable>
		{
			pos7 := pos
			// LU_phrase<relative_syllable>
			if !_fail(parser, _LU_phrase__relative_syllableFail, errPos, failure, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// relative_clause_1
		if !_fail(parser, _relative_clause_1Fail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
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

func _relative_clauseAction(parser *_Parser, start int) (int, *Relative) {
	var labels [2]string
	use(labels)
	var label0 (*CoP)
	var label1 *LUPhrase
	dp := parser.deltaPos[start][_relative_clause]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause}
	n := parser.act[key]
	if n != nil {
		n := n.(Relative)
		return start + int(dp-1), &n
	}
	var node Relative
	pos := start
	// x0:afterthought_cop<relative_clause_1, relative_clause_1> {…}/x1:LU_phrase<relative_syllable> {…}/relative_clause_1
	{
		pos3 := pos
		var node2 Relative
		// action
		{
			start5 := pos
			// x0:afterthought_cop<relative_clause_1, relative_clause_1>
			{
				pos6 := pos
				// afterthought_cop<relative_clause_1, relative_clause_1>
				if p, n := _afterthought_cop__relative_clause_1__relative_clause_1Action(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 *CoP) Relative {
				return Relative((*CoPRelative)(x0))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start8 := pos
			// x1:LU_phrase<relative_syllable>
			{
				pos9 := pos
				// LU_phrase<relative_syllable>
				if p, n := _LU_phrase__relative_syllableAction(parser, pos); n == nil {
					goto fail7
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos9:pos]
			}
			node = func(
				start, end int, x0 *CoP, x1 *LUPhrase) Relative {
				return Relative((*LURelative)(x1))
			}(
				start8, pos, label0, label1)
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// relative_clause_1
		if p, n := _relative_clause_1Action(parser, pos); n == nil {
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

func _relative_clause_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _relative_clause_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:forethought_cop<relative_clause, relative_clause> {…}/relative_clause_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<relative_clause, relative_clause>
		{
			pos5 := pos
			// forethought_cop<relative_clause, relative_clause>
			if !_accept(parser, _forethought_cop__relative_clause__relative_clauseAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_clause_2
		if !_accept(parser, _relative_clause_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_clause_1, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause_1, start, -1, perr)
}

func _relative_clause_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:forethought_cop<relative_clause, relative_clause> {…}/relative_clause_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:forethought_cop<relative_clause, relative_clause>
		{
			pos5 := pos
			// forethought_cop<relative_clause, relative_clause>
			if !_node(parser, _forethought_cop__relative_clause__relative_clauseNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_clause_2
		if !_node(parser, _relative_clause_2Node, node, &pos) {
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

func _relative_clause_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _relative_clause_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause_1}
	// x:forethought_cop<relative_clause, relative_clause> {…}/relative_clause_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<relative_clause, relative_clause>
		{
			pos5 := pos
			// forethought_cop<relative_clause, relative_clause>
			if !_fail(parser, _forethought_cop__relative_clause__relative_clauseFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_clause_2
		if !_fail(parser, _relative_clause_2Fail, errPos, failure, &pos) {
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

func _relative_clause_1Action(parser *_Parser, start int) (int, *Relative) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_relative_clause_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Relative)
		return start + int(dp-1), &n
	}
	var node Relative
	pos := start
	// x:forethought_cop<relative_clause, relative_clause> {…}/relative_clause_2
	{
		pos3 := pos
		var node2 Relative
		// action
		{
			start5 := pos
			// x:forethought_cop<relative_clause, relative_clause>
			{
				pos6 := pos
				// forethought_cop<relative_clause, relative_clause>
				if p, n := _forethought_cop__relative_clause__relative_clauseAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Relative {
				return Relative((*CoPRelative)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// relative_clause_2
		if p, n := _relative_clause_2Action(parser, pos); n == nil {
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

func _relative_clause_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _relative_clause_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:afterthought_cop<relative_clause_3, statement> {…}/x1:relative_clause_3 {…}
	{
		pos3 := pos
		// action
		// x0:afterthought_cop<relative_clause_3, statement>
		{
			pos5 := pos
			// afterthought_cop<relative_clause_3, statement>
			if !_accept(parser, _afterthought_cop__relative_clause_3__statementAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:relative_clause_3
		{
			pos7 := pos
			// relative_clause_3
			if !_accept(parser, _relative_clause_3Accepts, &pos, &perr) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _relative_clause_2, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause_2, start, -1, perr)
}

func _relative_clause_2Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// x0:afterthought_cop<relative_clause_3, statement> {…}/x1:relative_clause_3 {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:afterthought_cop<relative_clause_3, statement>
		{
			pos5 := pos
			// afterthought_cop<relative_clause_3, statement>
			if !_node(parser, _afterthought_cop__relative_clause_3__statementNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x1:relative_clause_3
		{
			pos7 := pos
			// relative_clause_3
			if !_node(parser, _relative_clause_3Node, node, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
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

func _relative_clause_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _relative_clause_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause_2}
	// x0:afterthought_cop<relative_clause_3, statement> {…}/x1:relative_clause_3 {…}
	{
		pos3 := pos
		// action
		// x0:afterthought_cop<relative_clause_3, statement>
		{
			pos5 := pos
			// afterthought_cop<relative_clause_3, statement>
			if !_fail(parser, _afterthought_cop__relative_clause_3__statementFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:relative_clause_3
		{
			pos7 := pos
			// relative_clause_3
			if !_fail(parser, _relative_clause_3Fail, errPos, failure, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
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

func _relative_clause_2Action(parser *_Parser, start int) (int, *Relative) {
	var labels [2]string
	use(labels)
	var label0 (*CoP)
	var label1 *PredicationRelative
	dp := parser.deltaPos[start][_relative_clause_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Relative)
		return start + int(dp-1), &n
	}
	var node Relative
	pos := start
	// x0:afterthought_cop<relative_clause_3, statement> {…}/x1:relative_clause_3 {…}
	{
		pos3 := pos
		var node2 Relative
		// action
		{
			start5 := pos
			// x0:afterthought_cop<relative_clause_3, statement>
			{
				pos6 := pos
				// afterthought_cop<relative_clause_3, statement>
				if p, n := _afterthought_cop__relative_clause_3__statementAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 *CoP) Relative {
				return Relative((*CoPRelative)(x0))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start8 := pos
			// x1:relative_clause_3
			{
				pos9 := pos
				// relative_clause_3
				if p, n := _relative_clause_3Action(parser, pos); n == nil {
					goto fail7
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos9:pos]
			}
			node = func(
				start, end int, x0 *CoP, x1 *PredicationRelative) Relative {
				return Relative(x1)
			}(
				start8, pos, label0, label1)
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

func _relative_clause_3Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _relative_clause_3, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:relative_predication s:_ na:end_statement?
	// p:relative_predication
	{
		pos1 := pos
		// relative_predication
		if !_accept(parser, _relative_predicationAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			pos5 := pos
			// end_statement
			if !_accept(parser, _end_statementAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _relative_clause_3, start, pos, perr)
fail:
	return _memoize(parser, _relative_clause_3, start, -1, perr)
}

func _relative_clause_3Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// p:relative_predication s:_ na:end_statement?
	// p:relative_predication
	{
		pos1 := pos
		// relative_predication
		if !_node(parser, _relative_predicationNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// end_statement
			if !_node(parser, _end_statementNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_clause_3Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _relative_clause_3, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_clause_3",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_clause_3}
	// action
	// p:relative_predication s:_ na:end_statement?
	// p:relative_predication
	{
		pos1 := pos
		// relative_predication
		if !_fail(parser, _relative_predicationFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			pos5 := pos
			// end_statement
			if !_fail(parser, _end_statementFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_clause_3Action(parser *_Parser, start int) (int, **PredicationRelative) {
	var labels [3]string
	use(labels)
	var label0 Predication
	var label1 *Mod
	var label2 *Word
	dp := parser.deltaPos[start][_relative_clause_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_clause_3}
	n := parser.act[key]
	if n != nil {
		n := n.(*PredicationRelative)
		return start + int(dp-1), &n
	}
	var node *PredicationRelative
	pos := start
	// action
	{
		start0 := pos
		// p:relative_predication s:_ na:end_statement?
		// p:relative_predication
		{
			pos2 := pos
			// relative_predication
			if p, n := _relative_predicationAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// na:end_statement?
		{
			pos4 := pos
			// end_statement?
			{
				pos6 := pos
				label2 = new(Word)
				// end_statement
				if p, n := _end_statementAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, na *Word, p Predication, s *Mod) *PredicationRelative {
			return &PredicationRelative{Predication: *endPredication(p, s, na)}
		}(
			start0, pos, label2, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_predicationAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _relative_predication, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:relative_predicate s:_ ts:terms?
	// p:relative_predicate
	{
		pos1 := pos
		// relative_predicate
		if !_accept(parser, _relative_predicateAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			pos5 := pos
			// terms
			if !_accept(parser, _termsAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _relative_predication, start, pos, perr)
fail:
	return _memoize(parser, _relative_predication, start, -1, perr)
}

func _relative_predicationNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// p:relative_predicate s:_ ts:terms?
	// p:relative_predicate
	{
		pos1 := pos
		// relative_predicate
		if !_node(parser, _relative_predicateNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// terms
			if !_node(parser, _termsNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_predicationFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _relative_predication, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_predication",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_predication}
	// action
	// p:relative_predicate s:_ ts:terms?
	// p:relative_predicate
	{
		pos1 := pos
		// relative_predicate
		if !_fail(parser, _relative_predicateFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			pos5 := pos
			// terms
			if !_fail(parser, _termsFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_predicationAction(parser *_Parser, start int) (int, *Predication) {
	var labels [3]string
	use(labels)
	var label2 *Terms
	var label0 Predicate
	var label1 *Mod
	dp := parser.deltaPos[start][_relative_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predication}
	n := parser.act[key]
	if n != nil {
		n := n.(Predication)
		return start + int(dp-1), &n
	}
	var node Predication
	pos := start
	// action
	{
		start0 := pos
		// p:relative_predicate s:_ ts:terms?
		// p:relative_predicate
		{
			pos2 := pos
			// relative_predicate
			if p, n := _relative_predicateAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// ts:terms?
		{
			pos4 := pos
			// terms?
			{
				pos6 := pos
				label2 = new(Terms)
				// terms
				if p, n := _termsAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, p Predicate, s *Mod, ts *Terms) Predication {
			return Predication(predication(p, s, ts))
		}(
			start0, pos, label0, label1, label2)
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
	// serial<relative_predicate_1>/relative_predicate_1
	{
		pos3 := pos
		// serial<relative_predicate_1>
		if !_accept(parser, _serial__relative_predicate_1Accepts, &pos, &perr) {
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
	// serial<relative_predicate_1>/relative_predicate_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial<relative_predicate_1>
		if !_node(parser, _serial__relative_predicate_1Node, node, &pos) {
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
	// serial<relative_predicate_1>/relative_predicate_1
	{
		pos3 := pos
		// serial<relative_predicate_1>
		if !_fail(parser, _serial__relative_predicate_1Fail, errPos, failure, &pos) {
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

func _relative_predicateAction(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_relative_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// serial<relative_predicate_1>/relative_predicate_1
	{
		pos3 := pos
		var node2 Predicate
		// serial<relative_predicate_1>
		if p, n := _serial__relative_predicate_1Action(parser, pos); n == nil {
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
	// afterthought_cop_pred<relative_predicate_2>/forethought_cop_pred<relative_predicate>/relative_predicate_2
	{
		pos3 := pos
		// afterthought_cop_pred<relative_predicate_2>
		if !_accept(parser, _afterthought_cop_pred__relative_predicate_2Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<relative_predicate>
		if !_accept(parser, _forethought_cop_pred__relative_predicateAccepts, &pos, &perr) {
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
	// afterthought_cop_pred<relative_predicate_2>/forethought_cop_pred<relative_predicate>/relative_predicate_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// afterthought_cop_pred<relative_predicate_2>
		if !_node(parser, _afterthought_cop_pred__relative_predicate_2Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_cop_pred<relative_predicate>
		if !_node(parser, _forethought_cop_pred__relative_predicateNode, node, &pos) {
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
	// afterthought_cop_pred<relative_predicate_2>/forethought_cop_pred<relative_predicate>/relative_predicate_2
	{
		pos3 := pos
		// afterthought_cop_pred<relative_predicate_2>
		if !_fail(parser, _afterthought_cop_pred__relative_predicate_2Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<relative_predicate>
		if !_fail(parser, _forethought_cop_pred__relative_predicateFail, errPos, failure, &pos) {
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

func _relative_predicate_1Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_relative_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// afterthought_cop_pred<relative_predicate_2>/forethought_cop_pred<relative_predicate>/relative_predicate_2
	{
		pos3 := pos
		var node2 Predicate
		// afterthought_cop_pred<relative_predicate_2>
		if p, n := _afterthought_cop_pred__relative_predicate_2Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_cop_pred<relative_predicate>
		if p, n := _forethought_cop_pred__relative_predicateAction(parser, pos); n == nil {
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
	// MI_phrase<relative_syllable>/PO_phrase<relative_syllable>/MO_phrase<relative_syllable>/predicate_word<relative_syllable>
	{
		pos3 := pos
		// MI_phrase<relative_syllable>
		if !_accept(parser, _MI_phrase__relative_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_phrase<relative_syllable>
		if !_accept(parser, _PO_phrase__relative_syllableAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MO_phrase<relative_syllable>
		if !_accept(parser, _MO_phrase__relative_syllableAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// predicate_word<relative_syllable>
		if !_accept(parser, _predicate_word__relative_syllableAccepts, &pos, &perr) {
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
	// MI_phrase<relative_syllable>/PO_phrase<relative_syllable>/MO_phrase<relative_syllable>/predicate_word<relative_syllable>
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// MI_phrase<relative_syllable>
		if !_node(parser, _MI_phrase__relative_syllableNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_phrase<relative_syllable>
		if !_node(parser, _PO_phrase__relative_syllableNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MO_phrase<relative_syllable>
		if !_node(parser, _MO_phrase__relative_syllableNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// predicate_word<relative_syllable>
		if !_node(parser, _predicate_word__relative_syllableNode, node, &pos) {
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
	// MI_phrase<relative_syllable>/PO_phrase<relative_syllable>/MO_phrase<relative_syllable>/predicate_word<relative_syllable>
	{
		pos3 := pos
		// MI_phrase<relative_syllable>
		if !_fail(parser, _MI_phrase__relative_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_phrase<relative_syllable>
		if !_fail(parser, _PO_phrase__relative_syllableFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MO_phrase<relative_syllable>
		if !_fail(parser, _MO_phrase__relative_syllableFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// predicate_word<relative_syllable>
		if !_fail(parser, _predicate_word__relative_syllableFail, errPos, failure, &pos) {
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

func _relative_predicate_2Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_relative_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_predicate_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// MI_phrase<relative_syllable>/PO_phrase<relative_syllable>/MO_phrase<relative_syllable>/predicate_word<relative_syllable>
	{
		pos3 := pos
		var node2 Predicate
		// MI_phrase<relative_syllable>
		if p, n := _MI_phrase__relative_syllableAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// PO_phrase<relative_syllable>
		if p, n := _PO_phrase__relative_syllableAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// MO_phrase<relative_syllable>
		if p, n := _MO_phrase__relative_syllableAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// predicate_word<relative_syllable>
		if p, n := _predicate_word__relative_syllableAction(parser, pos); n == nil {
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

func _adverbAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _adverb, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:afterthought_cop<adverb_1, adverb> {…}/adverb_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<adverb_1, adverb>
		{
			pos5 := pos
			// afterthought_cop<adverb_1, adverb>
			if !_accept(parser, _afterthought_cop__adverb_1__adverbAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_1
		if !_accept(parser, _adverb_1Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _adverb, start, pos, perr)
fail:
	return _memoize(parser, _adverb, start, -1, perr)
}

func _adverbNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:afterthought_cop<adverb_1, adverb> {…}/adverb_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:afterthought_cop<adverb_1, adverb>
		{
			pos5 := pos
			// afterthought_cop<adverb_1, adverb>
			if !_node(parser, _afterthought_cop__adverb_1__adverbNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_1
		if !_node(parser, _adverb_1Node, node, &pos) {
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

func _adverbFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _adverb, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb}
	// x:afterthought_cop<adverb_1, adverb> {…}/adverb_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<adverb_1, adverb>
		{
			pos5 := pos
			// afterthought_cop<adverb_1, adverb>
			if !_fail(parser, _afterthought_cop__adverb_1__adverbFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_1
		if !_fail(parser, _adverb_1Fail, errPos, failure, &pos) {
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

func _adverbAction(parser *_Parser, start int) (int, *Adverb) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_adverb]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb}
	n := parser.act[key]
	if n != nil {
		n := n.(Adverb)
		return start + int(dp-1), &n
	}
	var node Adverb
	pos := start
	// x:afterthought_cop<adverb_1, adverb> {…}/adverb_1
	{
		pos3 := pos
		var node2 Adverb
		// action
		{
			start5 := pos
			// x:afterthought_cop<adverb_1, adverb>
			{
				pos6 := pos
				// afterthought_cop<adverb_1, adverb>
				if p, n := _afterthought_cop__adverb_1__adverbAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Adverb {
				return Adverb((*CoPAdverb)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// adverb_1
		if p, n := _adverb_1Action(parser, pos); n == nil {
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

func _adverb_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _adverb_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:forethought_cop<adverb, adverb> {…}/adverb_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<adverb, adverb>
		{
			pos5 := pos
			// forethought_cop<adverb, adverb>
			if !_accept(parser, _forethought_cop__adverb__adverbAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_2
		if !_accept(parser, _adverb_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _adverb_1, start, pos, perr)
fail:
	return _memoize(parser, _adverb_1, start, -1, perr)
}

func _adverb_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:forethought_cop<adverb, adverb> {…}/adverb_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:forethought_cop<adverb, adverb>
		{
			pos5 := pos
			// forethought_cop<adverb, adverb>
			if !_node(parser, _forethought_cop__adverb__adverbNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_2
		if !_node(parser, _adverb_2Node, node, &pos) {
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

func _adverb_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _adverb_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_1}
	// x:forethought_cop<adverb, adverb> {…}/adverb_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<adverb, adverb>
		{
			pos5 := pos
			// forethought_cop<adverb, adverb>
			if !_fail(parser, _forethought_cop__adverb__adverbFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// adverb_2
		if !_fail(parser, _adverb_2Fail, errPos, failure, &pos) {
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

func _adverb_1Action(parser *_Parser, start int) (int, *Adverb) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_adverb_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Adverb)
		return start + int(dp-1), &n
	}
	var node Adverb
	pos := start
	// x:forethought_cop<adverb, adverb> {…}/adverb_2
	{
		pos3 := pos
		var node2 Adverb
		// action
		{
			start5 := pos
			// x:forethought_cop<adverb, adverb>
			{
				pos6 := pos
				// forethought_cop<adverb, adverb>
				if p, n := _forethought_cop__adverb__adverbAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Adverb {
				return Adverb((*CoPAdverb)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// adverb_2
		if p, n := _adverb_2Action(parser, pos); n == nil {
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

func _adverb_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _adverb_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:(serial<adverb_3>/adverb_3)
	{
		pos0 := pos
		// (serial<adverb_3>/adverb_3)
		// serial<adverb_3>/adverb_3
		{
			pos4 := pos
			// serial<adverb_3>
			if !_accept(parser, _serial__adverb_3Accepts, &pos, &perr) {
				goto fail5
			}
			goto ok1
		fail5:
			pos = pos4
			// adverb_3
			if !_accept(parser, _adverb_3Accepts, &pos, &perr) {
				goto fail6
			}
			goto ok1
		fail6:
			pos = pos4
			goto fail
		ok1:
		}
		labels[0] = parser.text[pos0:pos]
	}
	return _memoize(parser, _adverb_2, start, pos, perr)
fail:
	return _memoize(parser, _adverb_2, start, -1, perr)
}

func _adverb_2Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// p:(serial<adverb_3>/adverb_3)
	{
		pos0 := pos
		// (serial<adverb_3>/adverb_3)
		{
			nkids1 := len(node.Kids)
			pos02 := pos
			// serial<adverb_3>/adverb_3
			{
				pos6 := pos
				nkids4 := len(node.Kids)
				// serial<adverb_3>
				if !_node(parser, _serial__adverb_3Node, node, &pos) {
					goto fail7
				}
				goto ok3
			fail7:
				node.Kids = node.Kids[:nkids4]
				pos = pos6
				// adverb_3
				if !_node(parser, _adverb_3Node, node, &pos) {
					goto fail8
				}
				goto ok3
			fail8:
				node.Kids = node.Kids[:nkids4]
				pos = pos6
				goto fail
			ok3:
			}
			sub := _sub(parser, pos02, pos, node.Kids[nkids1:])
			node.Kids = append(node.Kids[:nkids1], sub)
		}
		labels[0] = parser.text[pos0:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _adverb_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_2}
	// action
	// p:(serial<adverb_3>/adverb_3)
	{
		pos0 := pos
		// (serial<adverb_3>/adverb_3)
		// serial<adverb_3>/adverb_3
		{
			pos4 := pos
			// serial<adverb_3>
			if !_fail(parser, _serial__adverb_3Fail, errPos, failure, &pos) {
				goto fail5
			}
			goto ok1
		fail5:
			pos = pos4
			// adverb_3
			if !_fail(parser, _adverb_3Fail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok1
		fail6:
			pos = pos4
			goto fail
		ok1:
		}
		labels[0] = parser.text[pos0:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_2Action(parser *_Parser, start int) (int, *Adverb) {
	var labels [1]string
	use(labels)
	var label0 Predicate
	dp := parser.deltaPos[start][_adverb_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Adverb)
		return start + int(dp-1), &n
	}
	var node Adverb
	pos := start
	// action
	{
		start0 := pos
		// p:(serial<adverb_3>/adverb_3)
		{
			pos1 := pos
			// (serial<adverb_3>/adverb_3)
			// serial<adverb_3>/adverb_3
			{
				pos5 := pos
				var node4 Predicate
				// serial<adverb_3>
				if p, n := _serial__adverb_3Action(parser, pos); n == nil {
					goto fail6
				} else {
					label0 = *n
					pos = p
				}
				goto ok2
			fail6:
				label0 = node4
				pos = pos5
				// adverb_3
				if p, n := _adverb_3Action(parser, pos); n == nil {
					goto fail7
				} else {
					label0 = *n
					pos = p
				}
				goto ok2
			fail7:
				label0 = node4
				pos = pos5
				goto fail
			ok2:
			}
			labels[0] = parser.text[pos1:pos]
		}
		node = func(
			start, end int, p Predicate) Adverb {
			return Adverb(&PredicateAdverb{Predicate: p})
		}(
			start0, pos, label0)
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
	// afterthought_cop_pred<adverb_4>/forethought_cop_pred<adverb_3>/LU_predicate<adverb_syllable>/MI_phrase<adverb_syllable>/PO_phrase<adverb_syllable>/MO_phrase<adverb_syllable>/adverb_4
	{
		pos3 := pos
		// afterthought_cop_pred<adverb_4>
		if !_accept(parser, _afterthought_cop_pred__adverb_4Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<adverb_3>
		if !_accept(parser, _forethought_cop_pred__adverb_3Accepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_predicate<adverb_syllable>
		if !_accept(parser, _LU_predicate__adverb_syllableAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_phrase<adverb_syllable>
		if !_accept(parser, _MI_phrase__adverb_syllableAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_phrase<adverb_syllable>
		if !_accept(parser, _PO_phrase__adverb_syllableAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// MO_phrase<adverb_syllable>
		if !_accept(parser, _MO_phrase__adverb_syllableAccepts, &pos, &perr) {
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
	// afterthought_cop_pred<adverb_4>/forethought_cop_pred<adverb_3>/LU_predicate<adverb_syllable>/MI_phrase<adverb_syllable>/PO_phrase<adverb_syllable>/MO_phrase<adverb_syllable>/adverb_4
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// afterthought_cop_pred<adverb_4>
		if !_node(parser, _afterthought_cop_pred__adverb_4Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_cop_pred<adverb_3>
		if !_node(parser, _forethought_cop_pred__adverb_3Node, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_predicate<adverb_syllable>
		if !_node(parser, _LU_predicate__adverb_syllableNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_phrase<adverb_syllable>
		if !_node(parser, _MI_phrase__adverb_syllableNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_phrase<adverb_syllable>
		if !_node(parser, _PO_phrase__adverb_syllableNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MO_phrase<adverb_syllable>
		if !_node(parser, _MO_phrase__adverb_syllableNode, node, &pos) {
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
	// afterthought_cop_pred<adverb_4>/forethought_cop_pred<adverb_3>/LU_predicate<adverb_syllable>/MI_phrase<adverb_syllable>/PO_phrase<adverb_syllable>/MO_phrase<adverb_syllable>/adverb_4
	{
		pos3 := pos
		// afterthought_cop_pred<adverb_4>
		if !_fail(parser, _afterthought_cop_pred__adverb_4Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<adverb_3>
		if !_fail(parser, _forethought_cop_pred__adverb_3Fail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_predicate<adverb_syllable>
		if !_fail(parser, _LU_predicate__adverb_syllableFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_phrase<adverb_syllable>
		if !_fail(parser, _MI_phrase__adverb_syllableFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_phrase<adverb_syllable>
		if !_fail(parser, _PO_phrase__adverb_syllableFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// MO_phrase<adverb_syllable>
		if !_fail(parser, _MO_phrase__adverb_syllableFail, errPos, failure, &pos) {
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

func _adverb_3Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_adverb_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_3}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// afterthought_cop_pred<adverb_4>/forethought_cop_pred<adverb_3>/LU_predicate<adverb_syllable>/MI_phrase<adverb_syllable>/PO_phrase<adverb_syllable>/MO_phrase<adverb_syllable>/adverb_4
	{
		pos3 := pos
		var node2 Predicate
		// afterthought_cop_pred<adverb_4>
		if p, n := _afterthought_cop_pred__adverb_4Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_cop_pred<adverb_3>
		if p, n := _forethought_cop_pred__adverb_3Action(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// LU_predicate<adverb_syllable>
		if p, n := _LU_predicate__adverb_syllableAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// MI_phrase<adverb_syllable>
		if p, n := _MI_phrase__adverb_syllableAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// PO_phrase<adverb_syllable>
		if p, n := _PO_phrase__adverb_syllableAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// MO_phrase<adverb_syllable>
		if p, n := _MO_phrase__adverb_syllableAction(parser, pos); n == nil {
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
	// predicate_word<adverb_syllable>
	if !_accept(parser, _predicate_word__adverb_syllableAccepts, &pos, &perr) {
		goto fail
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
	// predicate_word<adverb_syllable>
	if !_node(parser, _predicate_word__adverb_syllableNode, node, &pos) {
		goto fail
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
	// predicate_word<adverb_syllable>
	if !_fail(parser, _predicate_word__adverb_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_4Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_adverb_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_4}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// predicate_word<adverb_syllable>
	if p, n := _predicate_word__adverb_syllableAction(parser, pos); n == nil {
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

func _prepositional_phraseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _prepositional_phrase, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1> {…}/prepositional_phrase_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
		{
			pos5 := pos
			// afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
			if !_accept(parser, _afterthought_cop__prepositional_phrase_1__prepositional_phrase_1Accepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_1
		if !_accept(parser, _prepositional_phrase_1Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _prepositional_phrase, start, pos, perr)
fail:
	return _memoize(parser, _prepositional_phrase, start, -1, perr)
}

func _prepositional_phraseNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1> {…}/prepositional_phrase_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
		{
			pos5 := pos
			// afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
			if !_node(parser, _afterthought_cop__prepositional_phrase_1__prepositional_phrase_1Node, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// prepositional_phrase_1
		if !_node(parser, _prepositional_phrase_1Node, node, &pos) {
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

func _prepositional_phraseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _prepositional_phrase, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prepositional_phrase",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prepositional_phrase}
	// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1> {…}/prepositional_phrase_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
		{
			pos5 := pos
			// afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
			if !_fail(parser, _afterthought_cop__prepositional_phrase_1__prepositional_phrase_1Fail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_1
		if !_fail(parser, _prepositional_phrase_1Fail, errPos, failure, &pos) {
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

func _prepositional_phraseAction(parser *_Parser, start int) (int, *Preposition) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_prepositional_phrase]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase}
	n := parser.act[key]
	if n != nil {
		n := n.(Preposition)
		return start + int(dp-1), &n
	}
	var node Preposition
	pos := start
	// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1> {…}/prepositional_phrase_1
	{
		pos3 := pos
		var node2 Preposition
		// action
		{
			start5 := pos
			// x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
			{
				pos6 := pos
				// afterthought_cop<prepositional_phrase_1, prepositional_phrase_1>
				if p, n := _afterthought_cop__prepositional_phrase_1__prepositional_phrase_1Action(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Preposition {
				return Preposition((*CoPPreposition)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// prepositional_phrase_1
		if p, n := _prepositional_phrase_1Action(parser, pos); n == nil {
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

func _prepositional_phrase_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _prepositional_phrase_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:forethought_cop<prepositional_phrase, prepositional_phrase> {…}/prepositional_phrase_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<prepositional_phrase, prepositional_phrase>
		{
			pos5 := pos
			// forethought_cop<prepositional_phrase, prepositional_phrase>
			if !_accept(parser, _forethought_cop__prepositional_phrase__prepositional_phraseAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_2
		if !_accept(parser, _prepositional_phrase_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _prepositional_phrase_1, start, pos, perr)
fail:
	return _memoize(parser, _prepositional_phrase_1, start, -1, perr)
}

func _prepositional_phrase_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:forethought_cop<prepositional_phrase, prepositional_phrase> {…}/prepositional_phrase_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:forethought_cop<prepositional_phrase, prepositional_phrase>
		{
			pos5 := pos
			// forethought_cop<prepositional_phrase, prepositional_phrase>
			if !_node(parser, _forethought_cop__prepositional_phrase__prepositional_phraseNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// prepositional_phrase_2
		if !_node(parser, _prepositional_phrase_2Node, node, &pos) {
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

func _prepositional_phrase_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _prepositional_phrase_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prepositional_phrase_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prepositional_phrase_1}
	// x:forethought_cop<prepositional_phrase, prepositional_phrase> {…}/prepositional_phrase_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<prepositional_phrase, prepositional_phrase>
		{
			pos5 := pos
			// forethought_cop<prepositional_phrase, prepositional_phrase>
			if !_fail(parser, _forethought_cop__prepositional_phrase__prepositional_phraseFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// prepositional_phrase_2
		if !_fail(parser, _prepositional_phrase_2Fail, errPos, failure, &pos) {
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

func _prepositional_phrase_1Action(parser *_Parser, start int) (int, *Preposition) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_prepositional_phrase_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Preposition)
		return start + int(dp-1), &n
	}
	var node Preposition
	pos := start
	// x:forethought_cop<prepositional_phrase, prepositional_phrase> {…}/prepositional_phrase_2
	{
		pos3 := pos
		var node2 Preposition
		// action
		{
			start5 := pos
			// x:forethought_cop<prepositional_phrase, prepositional_phrase>
			{
				pos6 := pos
				// forethought_cop<prepositional_phrase, prepositional_phrase>
				if p, n := _forethought_cop__prepositional_phrase__prepositional_phraseAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Preposition {
				return Preposition((*CoPPreposition)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// prepositional_phrase_2
		if p, n := _prepositional_phrase_2Action(parser, pos); n == nil {
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

func _prepositional_phrase_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _prepositional_phrase_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:preposition s:_ a:argument
	// p:preposition
	{
		pos1 := pos
		// preposition
		if !_accept(parser, _prepositionAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// a:argument
	{
		pos3 := pos
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _prepositional_phrase_2, start, pos, perr)
fail:
	return _memoize(parser, _prepositional_phrase_2, start, -1, perr)
}

func _prepositional_phrase_2Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// p:preposition s:_ a:argument
	// p:preposition
	{
		pos1 := pos
		// preposition
		if !_node(parser, _prepositionNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// a:argument
	{
		pos3 := pos
		// argument
		if !_node(parser, _argumentNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prepositional_phrase_2Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _prepositional_phrase_2, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prepositional_phrase_2",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prepositional_phrase_2}
	// action
	// p:preposition s:_ a:argument
	// p:preposition
	{
		pos1 := pos
		// preposition
		if !_fail(parser, _prepositionFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// a:argument
	{
		pos3 := pos
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prepositional_phrase_2Action(parser *_Parser, start int) (int, *Preposition) {
	var labels [3]string
	use(labels)
	var label0 Predicate
	var label1 *Mod
	var label2 Argument
	dp := parser.deltaPos[start][_prepositional_phrase_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prepositional_phrase_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Preposition)
		return start + int(dp-1), &n
	}
	var node Preposition
	pos := start
	// action
	{
		start0 := pos
		// p:preposition s:_ a:argument
		// p:preposition
		{
			pos2 := pos
			// preposition
			if p, n := _prepositionAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// a:argument
		{
			pos4 := pos
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, a Argument, p Predicate, s *Mod) Preposition {
			return Preposition(prepPhrase(p, s, a))
		}(
			start0, pos, label2, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _prepositionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _preposition, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:afterthought_cop<preposition_1, preposition> {…}/preposition_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<preposition_1, preposition>
		{
			pos5 := pos
			// afterthought_cop<preposition_1, preposition>
			if !_accept(parser, _afterthought_cop__preposition_1__prepositionAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_1
		if !_accept(parser, _preposition_1Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _preposition, start, pos, perr)
fail:
	return _memoize(parser, _preposition, start, -1, perr)
}

func _prepositionNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:afterthought_cop<preposition_1, preposition> {…}/preposition_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:afterthought_cop<preposition_1, preposition>
		{
			pos5 := pos
			// afterthought_cop<preposition_1, preposition>
			if !_node(parser, _afterthought_cop__preposition_1__prepositionNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_1
		if !_node(parser, _preposition_1Node, node, &pos) {
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

func _prepositionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _preposition, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition}
	// x:afterthought_cop<preposition_1, preposition> {…}/preposition_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<preposition_1, preposition>
		{
			pos5 := pos
			// afterthought_cop<preposition_1, preposition>
			if !_fail(parser, _afterthought_cop__preposition_1__prepositionFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_1
		if !_fail(parser, _preposition_1Fail, errPos, failure, &pos) {
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

func _prepositionAction(parser *_Parser, start int) (int, *Predicate) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_preposition]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// x:afterthought_cop<preposition_1, preposition> {…}/preposition_1
	{
		pos3 := pos
		var node2 Predicate
		// action
		{
			start5 := pos
			// x:afterthought_cop<preposition_1, preposition>
			{
				pos6 := pos
				// afterthought_cop<preposition_1, preposition>
				if p, n := _afterthought_cop__preposition_1__prepositionAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Predicate {
				return Predicate((*CoPPredicate)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// preposition_1
		if p, n := _preposition_1Action(parser, pos); n == nil {
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

func _preposition_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _preposition_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:forethought_cop<preposition, preposition> {…}/preposition_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<preposition, preposition>
		{
			pos5 := pos
			// forethought_cop<preposition, preposition>
			if !_accept(parser, _forethought_cop__preposition__prepositionAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_2
		if !_accept(parser, _preposition_2Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _preposition_1, start, pos, perr)
fail:
	return _memoize(parser, _preposition_1, start, -1, perr)
}

func _preposition_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:forethought_cop<preposition, preposition> {…}/preposition_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:forethought_cop<preposition, preposition>
		{
			pos5 := pos
			// forethought_cop<preposition, preposition>
			if !_node(parser, _forethought_cop__preposition__prepositionNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_2
		if !_node(parser, _preposition_2Node, node, &pos) {
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

func _preposition_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _preposition_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_1}
	// x:forethought_cop<preposition, preposition> {…}/preposition_2
	{
		pos3 := pos
		// action
		// x:forethought_cop<preposition, preposition>
		{
			pos5 := pos
			// forethought_cop<preposition, preposition>
			if !_fail(parser, _forethought_cop__preposition__prepositionFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// preposition_2
		if !_fail(parser, _preposition_2Fail, errPos, failure, &pos) {
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

func _preposition_1Action(parser *_Parser, start int) (int, *Predicate) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_preposition_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// x:forethought_cop<preposition, preposition> {…}/preposition_2
	{
		pos3 := pos
		var node2 Predicate
		// action
		{
			start5 := pos
			// x:forethought_cop<preposition, preposition>
			{
				pos6 := pos
				// forethought_cop<preposition, preposition>
				if p, n := _forethought_cop__preposition__prepositionAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Predicate {
				return Predicate((*CoPPredicate)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// preposition_2
		if p, n := _preposition_2Action(parser, pos); n == nil {
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

func _preposition_2Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_2, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// serial<preposition_3>/preposition_3
	{
		pos3 := pos
		// serial<preposition_3>
		if !_accept(parser, _serial__preposition_3Accepts, &pos, &perr) {
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
	// serial<preposition_3>/preposition_3
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial<preposition_3>
		if !_node(parser, _serial__preposition_3Node, node, &pos) {
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
	// serial<preposition_3>/preposition_3
	{
		pos3 := pos
		// serial<preposition_3>
		if !_fail(parser, _serial__preposition_3Fail, errPos, failure, &pos) {
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

func _preposition_2Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_preposition_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// serial<preposition_3>/preposition_3
	{
		pos3 := pos
		var node2 Predicate
		// serial<preposition_3>
		if p, n := _serial__preposition_3Action(parser, pos); n == nil {
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
	// afterthought_cop_pred<preposition_4>/forethought_cop_pred<preposition>/LU_predicate<preposition_syllable>/MI_phrase<preposition_syllable>/PO_phrase<preposition_syllable>/MO_phrase<preposition_syllable>/preposition_4
	{
		pos3 := pos
		// afterthought_cop_pred<preposition_4>
		if !_accept(parser, _afterthought_cop_pred__preposition_4Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<preposition>
		if !_accept(parser, _forethought_cop_pred__prepositionAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_predicate<preposition_syllable>
		if !_accept(parser, _LU_predicate__preposition_syllableAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_phrase<preposition_syllable>
		if !_accept(parser, _MI_phrase__preposition_syllableAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_phrase<preposition_syllable>
		if !_accept(parser, _PO_phrase__preposition_syllableAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// MO_phrase<preposition_syllable>
		if !_accept(parser, _MO_phrase__preposition_syllableAccepts, &pos, &perr) {
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
	// afterthought_cop_pred<preposition_4>/forethought_cop_pred<preposition>/LU_predicate<preposition_syllable>/MI_phrase<preposition_syllable>/PO_phrase<preposition_syllable>/MO_phrase<preposition_syllable>/preposition_4
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// afterthought_cop_pred<preposition_4>
		if !_node(parser, _afterthought_cop_pred__preposition_4Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_cop_pred<preposition>
		if !_node(parser, _forethought_cop_pred__prepositionNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// LU_predicate<preposition_syllable>
		if !_node(parser, _LU_predicate__preposition_syllableNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MI_phrase<preposition_syllable>
		if !_node(parser, _MI_phrase__preposition_syllableNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_phrase<preposition_syllable>
		if !_node(parser, _PO_phrase__preposition_syllableNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MO_phrase<preposition_syllable>
		if !_node(parser, _MO_phrase__preposition_syllableNode, node, &pos) {
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
	// afterthought_cop_pred<preposition_4>/forethought_cop_pred<preposition>/LU_predicate<preposition_syllable>/MI_phrase<preposition_syllable>/PO_phrase<preposition_syllable>/MO_phrase<preposition_syllable>/preposition_4
	{
		pos3 := pos
		// afterthought_cop_pred<preposition_4>
		if !_fail(parser, _afterthought_cop_pred__preposition_4Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<preposition>
		if !_fail(parser, _forethought_cop_pred__prepositionFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// LU_predicate<preposition_syllable>
		if !_fail(parser, _LU_predicate__preposition_syllableFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// MI_phrase<preposition_syllable>
		if !_fail(parser, _MI_phrase__preposition_syllableFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// PO_phrase<preposition_syllable>
		if !_fail(parser, _PO_phrase__preposition_syllableFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// MO_phrase<preposition_syllable>
		if !_fail(parser, _MO_phrase__preposition_syllableFail, errPos, failure, &pos) {
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

func _preposition_3Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_preposition_3]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_3}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// afterthought_cop_pred<preposition_4>/forethought_cop_pred<preposition>/LU_predicate<preposition_syllable>/MI_phrase<preposition_syllable>/PO_phrase<preposition_syllable>/MO_phrase<preposition_syllable>/preposition_4
	{
		pos3 := pos
		var node2 Predicate
		// afterthought_cop_pred<preposition_4>
		if p, n := _afterthought_cop_pred__preposition_4Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_cop_pred<preposition>
		if p, n := _forethought_cop_pred__prepositionAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// LU_predicate<preposition_syllable>
		if p, n := _LU_predicate__preposition_syllableAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// MI_phrase<preposition_syllable>
		if p, n := _MI_phrase__preposition_syllableAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// PO_phrase<preposition_syllable>
		if p, n := _PO_phrase__preposition_syllableAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// MO_phrase<preposition_syllable>
		if p, n := _MO_phrase__preposition_syllableAction(parser, pos); n == nil {
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
	// predicate_word<preposition_syllable>
	if !_accept(parser, _predicate_word__preposition_syllableAccepts, &pos, &perr) {
		goto fail
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
	// predicate_word<preposition_syllable>
	if !_node(parser, _predicate_word__preposition_syllableNode, node, &pos) {
		goto fail
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
	// predicate_word<preposition_syllable>
	if !_fail(parser, _predicate_word__preposition_syllableFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_4Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_preposition_4]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_4}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// predicate_word<preposition_syllable>
	if p, n := _predicate_word__preposition_syllableAction(parser, pos); n == nil {
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

func _content_clauseAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _content_clause, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x:afterthought_cop<content_clause_1, statement> {…}/content_clause_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<content_clause_1, statement>
		{
			pos5 := pos
			// afterthought_cop<content_clause_1, statement>
			if !_accept(parser, _afterthought_cop__content_clause_1__statementAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// content_clause_1
		if !_accept(parser, _content_clause_1Accepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_clause, start, pos, perr)
fail:
	return _memoize(parser, _content_clause, start, -1, perr)
}

func _content_clauseNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// x:afterthought_cop<content_clause_1, statement> {…}/content_clause_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x:afterthought_cop<content_clause_1, statement>
		{
			pos5 := pos
			// afterthought_cop<content_clause_1, statement>
			if !_node(parser, _afterthought_cop__content_clause_1__statementNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// content_clause_1
		if !_node(parser, _content_clause_1Node, node, &pos) {
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

func _content_clauseFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _content_clause, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_clause",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_clause}
	// x:afterthought_cop<content_clause_1, statement> {…}/content_clause_1
	{
		pos3 := pos
		// action
		// x:afterthought_cop<content_clause_1, statement>
		{
			pos5 := pos
			// afterthought_cop<content_clause_1, statement>
			if !_fail(parser, _afterthought_cop__content_clause_1__statementFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// content_clause_1
		if !_fail(parser, _content_clause_1Fail, errPos, failure, &pos) {
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

func _content_clauseAction(parser *_Parser, start int) (int, *Content) {
	var labels [1]string
	use(labels)
	var label0 (*CoP)
	dp := parser.deltaPos[start][_content_clause]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_clause}
	n := parser.act[key]
	if n != nil {
		n := n.(Content)
		return start + int(dp-1), &n
	}
	var node Content
	pos := start
	// x:afterthought_cop<content_clause_1, statement> {…}/content_clause_1
	{
		pos3 := pos
		var node2 Content
		// action
		{
			start5 := pos
			// x:afterthought_cop<content_clause_1, statement>
			{
				pos6 := pos
				// afterthought_cop<content_clause_1, statement>
				if p, n := _afterthought_cop__content_clause_1__statementAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x *CoP) Content {
				return Content((*CoPContent)(x))
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// content_clause_1
		if p, n := _content_clause_1Action(parser, pos); n == nil {
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

func _content_clause_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _content_clause_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:content_statement {…}/x1:LU_phrase<content_syllable> {…}
	{
		pos3 := pos
		// action
		// x0:content_statement
		{
			pos5 := pos
			// content_statement
			if !_accept(parser, _content_statementAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:LU_phrase<content_syllable>
		{
			pos7 := pos
			// LU_phrase<content_syllable>
			if !_accept(parser, _LU_phrase__content_syllableAccepts, &pos, &perr) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _content_clause_1, start, pos, perr)
fail:
	return _memoize(parser, _content_clause_1, start, -1, perr)
}

func _content_clause_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
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
	// x0:content_statement {…}/x1:LU_phrase<content_syllable> {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:content_statement
		{
			pos5 := pos
			// content_statement
			if !_node(parser, _content_statementNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x1:LU_phrase<content_syllable>
		{
			pos7 := pos
			// LU_phrase<content_syllable>
			if !_node(parser, _LU_phrase__content_syllableNode, node, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
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

func _content_clause_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _content_clause_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_clause_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_clause_1}
	// x0:content_statement {…}/x1:LU_phrase<content_syllable> {…}
	{
		pos3 := pos
		// action
		// x0:content_statement
		{
			pos5 := pos
			// content_statement
			if !_fail(parser, _content_statementFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:LU_phrase<content_syllable>
		{
			pos7 := pos
			// LU_phrase<content_syllable>
			if !_fail(parser, _LU_phrase__content_syllableFail, errPos, failure, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
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

func _content_clause_1Action(parser *_Parser, start int) (int, *Content) {
	var labels [2]string
	use(labels)
	var label0 *PredicationContent
	var label1 *LUPhrase
	dp := parser.deltaPos[start][_content_clause_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_clause_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Content)
		return start + int(dp-1), &n
	}
	var node Content
	pos := start
	// x0:content_statement {…}/x1:LU_phrase<content_syllable> {…}
	{
		pos3 := pos
		var node2 Content
		// action
		{
			start5 := pos
			// x0:content_statement
			{
				pos6 := pos
				// content_statement
				if p, n := _content_statementAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 *PredicationContent) Content {
				return Content(x0)
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start8 := pos
			// x1:LU_phrase<content_syllable>
			{
				pos9 := pos
				// LU_phrase<content_syllable>
				if p, n := _LU_phrase__content_syllableAction(parser, pos); n == nil {
					goto fail7
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos9:pos]
			}
			node = func(
				start, end int, x0 *PredicationContent, x1 *LUPhrase) Content {
				return Content((*LUContent)(x1))
			}(
				start8, pos, label0, label1)
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

func _content_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _content_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:content_predication s:_ na:end_statement?
	// p:content_predication
	{
		pos1 := pos
		// content_predication
		if !_accept(parser, _content_predicationAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			pos5 := pos
			// end_statement
			if !_accept(parser, _end_statementAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _content_statement, start, pos, perr)
fail:
	return _memoize(parser, _content_statement, start, -1, perr)
}

func _content_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
	dp := parser.deltaPos[start][_content_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_statement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_statement"}
	// action
	// p:content_predication s:_ na:end_statement?
	// p:content_predication
	{
		pos1 := pos
		// content_predication
		if !_node(parser, _content_predicationNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// end_statement
			if !_node(parser, _end_statementNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _content_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_statement}
	// action
	// p:content_predication s:_ na:end_statement?
	// p:content_predication
	{
		pos1 := pos
		// content_predication
		if !_fail(parser, _content_predicationFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// na:end_statement?
	{
		pos3 := pos
		// end_statement?
		{
			pos5 := pos
			// end_statement
			if !_fail(parser, _end_statementFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_statementAction(parser *_Parser, start int) (int, **PredicationContent) {
	var labels [3]string
	use(labels)
	var label0 Predication
	var label1 *Mod
	var label2 *Word
	dp := parser.deltaPos[start][_content_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(*PredicationContent)
		return start + int(dp-1), &n
	}
	var node *PredicationContent
	pos := start
	// action
	{
		start0 := pos
		// p:content_predication s:_ na:end_statement?
		// p:content_predication
		{
			pos2 := pos
			// content_predication
			if p, n := _content_predicationAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// na:end_statement?
		{
			pos4 := pos
			// end_statement?
			{
				pos6 := pos
				label2 = new(Word)
				// end_statement
				if p, n := _end_statementAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, na *Word, p Predication, s *Mod) *PredicationContent {
			return &PredicationContent{Predication: *endPredication(p, s, na)}
		}(
			start0, pos, label2, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_predicationAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _content_predication, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:content_predicate s:_ ts:terms?
	// p:content_predicate
	{
		pos1 := pos
		// content_predicate
		if !_accept(parser, _content_predicateAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			pos5 := pos
			// terms
			if !_accept(parser, _termsAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _content_predication, start, pos, perr)
fail:
	return _memoize(parser, _content_predication, start, -1, perr)
}

func _content_predicationNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// p:content_predicate s:_ ts:terms?
	// p:content_predicate
	{
		pos1 := pos
		// content_predicate
		if !_node(parser, _content_predicateNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			nkids4 := len(node.Kids)
			pos5 := pos
			// terms
			if !_node(parser, _termsNode, node, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			node.Kids = node.Kids[:nkids4]
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_predicationFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _content_predication, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_predication",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_predication}
	// action
	// p:content_predicate s:_ ts:terms?
	// p:content_predicate
	{
		pos1 := pos
		// content_predicate
		if !_fail(parser, _content_predicateFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ts:terms?
	{
		pos3 := pos
		// terms?
		{
			pos5 := pos
			// terms
			if !_fail(parser, _termsFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok7
		fail6:
			pos = pos5
		ok7:
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_predicationAction(parser *_Parser, start int) (int, *Predication) {
	var labels [3]string
	use(labels)
	var label0 Predicate
	var label1 *Mod
	var label2 *Terms
	dp := parser.deltaPos[start][_content_predication]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predication}
	n := parser.act[key]
	if n != nil {
		n := n.(Predication)
		return start + int(dp-1), &n
	}
	var node Predication
	pos := start
	// action
	{
		start0 := pos
		// p:content_predicate s:_ ts:terms?
		// p:content_predicate
		{
			pos2 := pos
			// content_predicate
			if p, n := _content_predicateAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// ts:terms?
		{
			pos4 := pos
			// terms?
			{
				pos6 := pos
				label2 = new(Terms)
				// terms
				if p, n := _termsAction(parser, pos); n == nil {
					goto fail7
				} else {
					*label2 = *n
					pos = p
				}
				goto ok8
			fail7:
				label2 = nil
				pos = pos6
			ok8:
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, p Predicate, s *Mod, ts *Terms) Predication {
			return Predication(predication(p, s, ts))
		}(
			start0, pos, label0, label1, label2)
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
	// serial<content_predicate_1>/content_predicate_1
	{
		pos3 := pos
		// serial<content_predicate_1>
		if !_accept(parser, _serial__content_predicate_1Accepts, &pos, &perr) {
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
	// serial<content_predicate_1>/content_predicate_1
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// serial<content_predicate_1>
		if !_node(parser, _serial__content_predicate_1Node, node, &pos) {
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
	// serial<content_predicate_1>/content_predicate_1
	{
		pos3 := pos
		// serial<content_predicate_1>
		if !_fail(parser, _serial__content_predicate_1Fail, errPos, failure, &pos) {
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

func _content_predicateAction(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_content_predicate]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// serial<content_predicate_1>/content_predicate_1
	{
		pos3 := pos
		var node2 Predicate
		// serial<content_predicate_1>
		if p, n := _serial__content_predicate_1Action(parser, pos); n == nil {
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
	// afterthought_cop_pred<content_predicate_2>/forethought_cop_pred<content_predicate_2>/content_predicate_2
	{
		pos3 := pos
		// afterthought_cop_pred<content_predicate_2>
		if !_accept(parser, _afterthought_cop_pred__content_predicate_2Accepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<content_predicate_2>
		if !_accept(parser, _forethought_cop_pred__content_predicate_2Accepts, &pos, &perr) {
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
	// afterthought_cop_pred<content_predicate_2>/forethought_cop_pred<content_predicate_2>/content_predicate_2
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// afterthought_cop_pred<content_predicate_2>
		if !_node(parser, _afterthought_cop_pred__content_predicate_2Node, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// forethought_cop_pred<content_predicate_2>
		if !_node(parser, _forethought_cop_pred__content_predicate_2Node, node, &pos) {
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
	// afterthought_cop_pred<content_predicate_2>/forethought_cop_pred<content_predicate_2>/content_predicate_2
	{
		pos3 := pos
		// afterthought_cop_pred<content_predicate_2>
		if !_fail(parser, _afterthought_cop_pred__content_predicate_2Fail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// forethought_cop_pred<content_predicate_2>
		if !_fail(parser, _forethought_cop_pred__content_predicate_2Fail, errPos, failure, &pos) {
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

func _content_predicate_1Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_content_predicate_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// afterthought_cop_pred<content_predicate_2>/forethought_cop_pred<content_predicate_2>/content_predicate_2
	{
		pos3 := pos
		var node2 Predicate
		// afterthought_cop_pred<content_predicate_2>
		if p, n := _afterthought_cop_pred__content_predicate_2Action(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// forethought_cop_pred<content_predicate_2>
		if p, n := _forethought_cop_pred__content_predicate_2Action(parser, pos); n == nil {
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
	// MI_phrase<content_syllable>/PO_phrase<content_syllable>/MO_phrase<content_syllable>/predicate_word<content_syllable>
	{
		pos3 := pos
		// MI_phrase<content_syllable>
		if !_accept(parser, _MI_phrase__content_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_phrase<content_syllable>
		if !_accept(parser, _PO_phrase__content_syllableAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MO_phrase<content_syllable>
		if !_accept(parser, _MO_phrase__content_syllableAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// predicate_word<content_syllable>
		if !_accept(parser, _predicate_word__content_syllableAccepts, &pos, &perr) {
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
	// MI_phrase<content_syllable>/PO_phrase<content_syllable>/MO_phrase<content_syllable>/predicate_word<content_syllable>
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// MI_phrase<content_syllable>
		if !_node(parser, _MI_phrase__content_syllableNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO_phrase<content_syllable>
		if !_node(parser, _PO_phrase__content_syllableNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// MO_phrase<content_syllable>
		if !_node(parser, _MO_phrase__content_syllableNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// predicate_word<content_syllable>
		if !_node(parser, _predicate_word__content_syllableNode, node, &pos) {
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
	// MI_phrase<content_syllable>/PO_phrase<content_syllable>/MO_phrase<content_syllable>/predicate_word<content_syllable>
	{
		pos3 := pos
		// MI_phrase<content_syllable>
		if !_fail(parser, _MI_phrase__content_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// PO_phrase<content_syllable>
		if !_fail(parser, _PO_phrase__content_syllableFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// MO_phrase<content_syllable>
		if !_fail(parser, _MO_phrase__content_syllableFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// predicate_word<content_syllable>
		if !_fail(parser, _predicate_word__content_syllableFail, errPos, failure, &pos) {
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

func _content_predicate_2Action(parser *_Parser, start int) (int, *Predicate) {
	dp := parser.deltaPos[start][_content_predicate_2]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_predicate_2}
	n := parser.act[key]
	if n != nil {
		n := n.(Predicate)
		return start + int(dp-1), &n
	}
	var node Predicate
	pos := start
	// MI_phrase<content_syllable>/PO_phrase<content_syllable>/MO_phrase<content_syllable>/predicate_word<content_syllable>
	{
		pos3 := pos
		var node2 Predicate
		// MI_phrase<content_syllable>
		if p, n := _MI_phrase__content_syllableAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// PO_phrase<content_syllable>
		if p, n := _PO_phrase__content_syllableAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// MO_phrase<content_syllable>
		if p, n := _MO_phrase__content_syllableAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// predicate_word<content_syllable>
		if p, n := _predicate_word__content_syllableAction(parser, pos); n == nil {
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

func _forethought_connectiveAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _forethought_connective, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// to:forethought_marker s:_ ru:connective
	// to:forethought_marker
	{
		pos1 := pos
		// forethought_marker
		if !_accept(parser, _forethought_markerAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ru:connective
	{
		pos3 := pos
		// connective
		if !_accept(parser, _connectiveAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _forethought_connective, start, pos, perr)
fail:
	return _memoize(parser, _forethought_connective, start, -1, perr)
}

func _forethought_connectiveNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// to:forethought_marker s:_ ru:connective
	// to:forethought_marker
	{
		pos1 := pos
		// forethought_marker
		if !_node(parser, _forethought_markerNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ru:connective
	{
		pos3 := pos
		// connective
		if !_node(parser, _connectiveNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_connectiveFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _forethought_connective, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_connective",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_connective}
	// action
	// to:forethought_marker s:_ ru:connective
	// to:forethought_marker
	{
		pos1 := pos
		// forethought_marker
		if !_fail(parser, _forethought_markerFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ru:connective
	{
		pos3 := pos
		// connective
		if !_fail(parser, _connectiveFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_connectiveAction(parser *_Parser, start int) (int, *CoP) {
	var labels [3]string
	use(labels)
	var label0 Word
	var label1 *Mod
	var label2 Word
	dp := parser.deltaPos[start][_forethought_connective]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_connective}
	n := parser.act[key]
	if n != nil {
		n := n.(CoP)
		return start + int(dp-1), &n
	}
	var node CoP
	pos := start
	// action
	{
		start0 := pos
		// to:forethought_marker s:_ ru:connective
		// to:forethought_marker
		{
			pos2 := pos
			// forethought_marker
			if p, n := _forethought_markerAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// ru:connective
		{
			pos4 := pos
			// connective
			if p, n := _connectiveAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, ru Word, s *Mod, to Word) CoP {
			return CoP(connector(to, s, ru))
		}(
			start0, pos, label2, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _MI_complementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [4]string
	use(labels)
	if dp, de, ok := _memo(parser, _MI_complement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:predicate {…}/x1:argument {…}/x2:adverb {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		// action
		// x0:predicate
		{
			pos5 := pos
			// predicate
			if !_accept(parser, _predicateAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:argument
		{
			pos7 := pos
			// argument
			if !_accept(parser, _argumentAccepts, &pos, &perr) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:adverb
		{
			pos9 := pos
			// adverb
			if !_accept(parser, _adverbAccepts, &pos, &perr) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:prepositional_phrase
		{
			pos11 := pos
			// prepositional_phrase
			if !_accept(parser, _prepositional_phraseAccepts, &pos, &perr) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _MI_complement, start, pos, perr)
fail:
	return _memoize(parser, _MI_complement, start, -1, perr)
}

func _MI_complementNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [4]string
	use(labels)
	dp := parser.deltaPos[start][_MI_complement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_complement}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MI_complement"}
	// x0:predicate {…}/x1:argument {…}/x2:adverb {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:predicate
		{
			pos5 := pos
			// predicate
			if !_node(parser, _predicateNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x1:argument
		{
			pos7 := pos
			// argument
			if !_node(parser, _argumentNode, node, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x2:adverb
		{
			pos9 := pos
			// adverb
			if !_node(parser, _adverbNode, node, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x3:prepositional_phrase
		{
			pos11 := pos
			// prepositional_phrase
			if !_node(parser, _prepositional_phraseNode, node, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
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

func _MI_complementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [4]string
	use(labels)
	pos, failure := _failMemo(parser, _MI_complement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MI_complement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MI_complement}
	// x0:predicate {…}/x1:argument {…}/x2:adverb {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		// action
		// x0:predicate
		{
			pos5 := pos
			// predicate
			if !_fail(parser, _predicateFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:argument
		{
			pos7 := pos
			// argument
			if !_fail(parser, _argumentFail, errPos, failure, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:adverb
		{
			pos9 := pos
			// adverb
			if !_fail(parser, _adverbFail, errPos, failure, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:prepositional_phrase
		{
			pos11 := pos
			// prepositional_phrase
			if !_fail(parser, _prepositional_phraseFail, errPos, failure, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
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

func _MI_complementAction(parser *_Parser, start int) (int, *Phrase) {
	var labels [4]string
	use(labels)
	var label0 Predicate
	var label1 Argument
	var label2 Adverb
	var label3 Preposition
	dp := parser.deltaPos[start][_MI_complement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MI_complement}
	n := parser.act[key]
	if n != nil {
		n := n.(Phrase)
		return start + int(dp-1), &n
	}
	var node Phrase
	pos := start
	// x0:predicate {…}/x1:argument {…}/x2:adverb {…}/x3:prepositional_phrase {…}
	{
		pos3 := pos
		var node2 Phrase
		// action
		{
			start5 := pos
			// x0:predicate
			{
				pos6 := pos
				// predicate
				if p, n := _predicateAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 Predicate) Phrase {
				return Phrase(x0)
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start8 := pos
			// x1:argument
			{
				pos9 := pos
				// argument
				if p, n := _argumentAction(parser, pos); n == nil {
					goto fail7
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos9:pos]
			}
			node = func(
				start, end int, x0 Predicate, x1 Argument) Phrase {
				return Phrase(x1)
			}(
				start8, pos, label0, label1)
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// action
		{
			start11 := pos
			// x2:adverb
			{
				pos12 := pos
				// adverb
				if p, n := _adverbAction(parser, pos); n == nil {
					goto fail10
				} else {
					label2 = *n
					pos = p
				}
				labels[2] = parser.text[pos12:pos]
			}
			node = func(
				start, end int, x0 Predicate, x1 Argument, x2 Adverb) Phrase {
				return Phrase(x2)
			}(
				start11, pos, label0, label1, label2)
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// action
		{
			start14 := pos
			// x3:prepositional_phrase
			{
				pos15 := pos
				// prepositional_phrase
				if p, n := _prepositional_phraseAction(parser, pos); n == nil {
					goto fail13
				} else {
					label3 = *n
					pos = p
				}
				labels[3] = parser.text[pos15:pos]
			}
			node = func(
				start, end int, x0 Predicate, x1 Argument, x2 Adverb, x3 Preposition) Phrase {
				return Phrase(x3)
			}(
				start14, pos, label0, label1, label2, label3)
		}
		goto ok0
	fail13:
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

func __Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, __, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// space_or_freemod?
	{
		pos1 := pos
		// space_or_freemod
		if !_accept(parser, _space_or_freemodAccepts, &pos, &perr) {
			goto fail2
		}
		goto ok3
	fail2:
		pos = pos1
	ok3:
	}
	return _memoize(parser, __, start, pos, perr)
}

func __Node(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][__]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: __}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "_"}
	// space_or_freemod?
	{
		nkids0 := len(node.Kids)
		pos1 := pos
		// space_or_freemod
		if !_node(parser, _space_or_freemodNode, node, &pos) {
			goto fail2
		}
		goto ok3
	fail2:
		node.Kids = node.Kids[:nkids0]
		pos = pos1
	ok3:
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
}

func __Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, __, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "_",
		Pos:  int(start),
	}
	key := _key{start: start, rule: __}
	// space_or_freemod?
	{
		pos1 := pos
		// space_or_freemod
		if !_fail(parser, _space_or_freemodFail, errPos, failure, &pos) {
			goto fail2
		}
		goto ok3
	fail2:
		pos = pos1
	ok3:
	}
	parser.fail[key] = failure
	return pos, failure
}

func __Action(parser *_Parser, start int) (int, **Mod) {
	dp := parser.deltaPos[start][__]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: __}
	n := parser.act[key]
	if n != nil {
		n := n.(*Mod)
		return start + int(dp-1), &n
	}
	var node *Mod
	pos := start
	// space_or_freemod?
	{
		pos1 := pos
		node = new(Mod)
		// space_or_freemod
		if p, n := _space_or_freemodAction(parser, pos); n == nil {
			goto fail2
		} else {
			*node = *n
			pos = p
		}
		goto ok3
	fail2:
		node = nil
		pos = pos1
	ok3:
	}
	parser.act[key] = node
	return pos, &node
}

func _space_or_freemodAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _space_or_freemod, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// m:(freemod/space_word) ms:space_or_freemod?
	// m:(freemod/space_word)
	{
		pos1 := pos
		// (freemod/space_word)
		// freemod/space_word
		{
			pos5 := pos
			// freemod
			if !_accept(parser, _freemodAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok2
		fail6:
			pos = pos5
			// space_word
			if !_accept(parser, _space_wordAccepts, &pos, &perr) {
				goto fail7
			}
			goto ok2
		fail7:
			pos = pos5
			goto fail
		ok2:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// ms:space_or_freemod?
	{
		pos8 := pos
		// space_or_freemod?
		{
			pos10 := pos
			// space_or_freemod
			if !_accept(parser, _space_or_freemodAccepts, &pos, &perr) {
				goto fail11
			}
			goto ok12
		fail11:
			pos = pos10
		ok12:
		}
		labels[1] = parser.text[pos8:pos]
	}
	return _memoize(parser, _space_or_freemod, start, pos, perr)
fail:
	return _memoize(parser, _space_or_freemod, start, -1, perr)
}

func _space_or_freemodNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_space_or_freemod]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _space_or_freemod}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "space_or_freemod"}
	// action
	// m:(freemod/space_word) ms:space_or_freemod?
	// m:(freemod/space_word)
	{
		pos1 := pos
		// (freemod/space_word)
		{
			nkids2 := len(node.Kids)
			pos03 := pos
			// freemod/space_word
			{
				pos7 := pos
				nkids5 := len(node.Kids)
				// freemod
				if !_node(parser, _freemodNode, node, &pos) {
					goto fail8
				}
				goto ok4
			fail8:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				// space_word
				if !_node(parser, _space_wordNode, node, &pos) {
					goto fail9
				}
				goto ok4
			fail9:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				goto fail
			ok4:
			}
			sub := _sub(parser, pos03, pos, node.Kids[nkids2:])
			node.Kids = append(node.Kids[:nkids2], sub)
		}
		labels[0] = parser.text[pos1:pos]
	}
	// ms:space_or_freemod?
	{
		pos10 := pos
		// space_or_freemod?
		{
			nkids11 := len(node.Kids)
			pos12 := pos
			// space_or_freemod
			if !_node(parser, _space_or_freemodNode, node, &pos) {
				goto fail13
			}
			goto ok14
		fail13:
			node.Kids = node.Kids[:nkids11]
			pos = pos12
		ok14:
		}
		labels[1] = parser.text[pos10:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _space_or_freemodFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _space_or_freemod, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "space_or_freemod",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _space_or_freemod}
	// action
	// m:(freemod/space_word) ms:space_or_freemod?
	// m:(freemod/space_word)
	{
		pos1 := pos
		// (freemod/space_word)
		// freemod/space_word
		{
			pos5 := pos
			// freemod
			if !_fail(parser, _freemodFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok2
		fail6:
			pos = pos5
			// space_word
			if !_fail(parser, _space_wordFail, errPos, failure, &pos) {
				goto fail7
			}
			goto ok2
		fail7:
			pos = pos5
			goto fail
		ok2:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// ms:space_or_freemod?
	{
		pos8 := pos
		// space_or_freemod?
		{
			pos10 := pos
			// space_or_freemod
			if !_fail(parser, _space_or_freemodFail, errPos, failure, &pos) {
				goto fail11
			}
			goto ok12
		fail11:
			pos = pos10
		ok12:
		}
		labels[1] = parser.text[pos8:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _space_or_freemodAction(parser *_Parser, start int) (int, *Mod) {
	var labels [2]string
	use(labels)
	var label0 Mod
	var label1 *Mod
	dp := parser.deltaPos[start][_space_or_freemod]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _space_or_freemod}
	n := parser.act[key]
	if n != nil {
		n := n.(Mod)
		return start + int(dp-1), &n
	}
	var node Mod
	pos := start
	// action
	{
		start0 := pos
		// m:(freemod/space_word) ms:space_or_freemod?
		// m:(freemod/space_word)
		{
			pos2 := pos
			// (freemod/space_word)
			// freemod/space_word
			{
				pos6 := pos
				var node5 Mod
				// freemod
				if p, n := _freemodAction(parser, pos); n == nil {
					goto fail7
				} else {
					label0 = *n
					pos = p
				}
				goto ok3
			fail7:
				label0 = node5
				pos = pos6
				// space_word
				if p, n := _space_wordAction(parser, pos); n == nil {
					goto fail8
				} else {
					label0 = *n
					pos = p
				}
				goto ok3
			fail8:
				label0 = node5
				pos = pos6
				goto fail
			ok3:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// ms:space_or_freemod?
		{
			pos9 := pos
			// space_or_freemod?
			{
				pos11 := pos
				label1 = new(Mod)
				// space_or_freemod
				if p, n := _space_or_freemodAction(parser, pos); n == nil {
					goto fail12
				} else {
					*label1 = *n
					pos = p
				}
				goto ok13
			fail12:
				label1 = nil
				pos = pos11
			ok13:
			}
			labels[1] = parser.text[pos9:pos]
		}
		node = func(
			start, end int, m Mod, ms *Mod) Mod {
			return Mod(m.mod(ms))
		}(
			start0, pos, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _space_wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _space_word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// s:space+
	{
		pos0 := pos
		// space+
		// space
		if !_accept(parser, _spaceAccepts, &pos, &perr) {
			goto fail
		}
		for {
			pos2 := pos
			// space
			if !_accept(parser, _spaceAccepts, &pos, &perr) {
				goto fail4
			}
			continue
		fail4:
			pos = pos2
			break
		}
		labels[0] = parser.text[pos0:pos]
	}
	return _memoize(parser, _space_word, start, pos, perr)
fail:
	return _memoize(parser, _space_word, start, -1, perr)
}

func _space_wordNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_space_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _space_word}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "space_word"}
	// action
	// s:space+
	{
		pos0 := pos
		// space+
		// space
		if !_node(parser, _spaceNode, node, &pos) {
			goto fail
		}
		for {
			nkids1 := len(node.Kids)
			pos2 := pos
			// space
			if !_node(parser, _spaceNode, node, &pos) {
				goto fail4
			}
			continue
		fail4:
			node.Kids = node.Kids[:nkids1]
			pos = pos2
			break
		}
		labels[0] = parser.text[pos0:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _space_wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _space_word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "space_word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _space_word}
	// action
	// s:space+
	{
		pos0 := pos
		// space+
		// space
		if !_fail(parser, _spaceFail, errPos, failure, &pos) {
			goto fail
		}
		for {
			pos2 := pos
			// space
			if !_fail(parser, _spaceFail, errPos, failure, &pos) {
				goto fail4
			}
			continue
		fail4:
			pos = pos2
			break
		}
		labels[0] = parser.text[pos0:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _space_wordAction(parser *_Parser, start int) (int, *Mod) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_space_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _space_word}
	n := parser.act[key]
	if n != nil {
		n := n.(Mod)
		return start + int(dp-1), &n
	}
	var node Mod
	pos := start
	// action
	{
		start0 := pos
		// s:space+
		{
			pos1 := pos
			// space+
			{
				var node4 string
				// space
				if p, n := _spaceAction(parser, pos); n == nil {
					goto fail
				} else {
					node4 = *n
					pos = p
				}
				label0 += node4
			}
			for {
				pos3 := pos
				var node4 string
				// space
				if p, n := _spaceAction(parser, pos); n == nil {
					goto fail5
				} else {
					node4 = *n
					pos = p
				}
				label0 += node4
				continue
			fail5:
				pos = pos3
				break
			}
			labels[0] = parser.text[pos1:pos]
		}
		node = func(
			start, end int, s string) Mod {
			return Mod(&Space{S: start, E: end, T: s})
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _freemodAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [4]string
	use(labels)
	if dp, de, ok := _memo(parser, _freemod, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// x0:interjection {…}/x1:parenthetical {…}/x2:incidental {…}/x3:vocative {…}
	{
		pos3 := pos
		// action
		// x0:interjection
		{
			pos5 := pos
			// interjection
			if !_accept(parser, _interjectionAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:parenthetical
		{
			pos7 := pos
			// parenthetical
			if !_accept(parser, _parentheticalAccepts, &pos, &perr) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:incidental
		{
			pos9 := pos
			// incidental
			if !_accept(parser, _incidentalAccepts, &pos, &perr) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:vocative
		{
			pos11 := pos
			// vocative
			if !_accept(parser, _vocativeAccepts, &pos, &perr) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _freemod, start, pos, perr)
fail:
	return _memoize(parser, _freemod, start, -1, perr)
}

func _freemodNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [4]string
	use(labels)
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
	// x0:interjection {…}/x1:parenthetical {…}/x2:incidental {…}/x3:vocative {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// x0:interjection
		{
			pos5 := pos
			// interjection
			if !_node(parser, _interjectionNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x1:parenthetical
		{
			pos7 := pos
			// parenthetical
			if !_node(parser, _parentheticalNode, node, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x2:incidental
		{
			pos9 := pos
			// incidental
			if !_node(parser, _incidentalNode, node, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// x3:vocative
		{
			pos11 := pos
			// vocative
			if !_node(parser, _vocativeNode, node, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
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

func _freemodFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [4]string
	use(labels)
	pos, failure := _failMemo(parser, _freemod, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "freemod",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _freemod}
	// x0:interjection {…}/x1:parenthetical {…}/x2:incidental {…}/x3:vocative {…}
	{
		pos3 := pos
		// action
		// x0:interjection
		{
			pos5 := pos
			// interjection
			if !_fail(parser, _interjectionFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos5:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// x1:parenthetical
		{
			pos7 := pos
			// parenthetical
			if !_fail(parser, _parentheticalFail, errPos, failure, &pos) {
				goto fail6
			}
			labels[1] = parser.text[pos7:pos]
		}
		goto ok0
	fail6:
		pos = pos3
		// action
		// x2:incidental
		{
			pos9 := pos
			// incidental
			if !_fail(parser, _incidentalFail, errPos, failure, &pos) {
				goto fail8
			}
			labels[2] = parser.text[pos9:pos]
		}
		goto ok0
	fail8:
		pos = pos3
		// action
		// x3:vocative
		{
			pos11 := pos
			// vocative
			if !_fail(parser, _vocativeFail, errPos, failure, &pos) {
				goto fail10
			}
			labels[3] = parser.text[pos11:pos]
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

func _freemodAction(parser *_Parser, start int) (int, *Mod) {
	var labels [4]string
	use(labels)
	var label1 (*Parenthetical)
	var label2 *Incidental
	var label3 *Vocative
	var label0 *Interjection
	dp := parser.deltaPos[start][_freemod]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _freemod}
	n := parser.act[key]
	if n != nil {
		n := n.(Mod)
		return start + int(dp-1), &n
	}
	var node Mod
	pos := start
	// x0:interjection {…}/x1:parenthetical {…}/x2:incidental {…}/x3:vocative {…}
	{
		pos3 := pos
		var node2 Mod
		// action
		{
			start5 := pos
			// x0:interjection
			{
				pos6 := pos
				// interjection
				if p, n := _interjectionAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos6:pos]
			}
			node = func(
				start, end int, x0 *Interjection) Mod {
				return Mod(x0)
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start8 := pos
			// x1:parenthetical
			{
				pos9 := pos
				// parenthetical
				if p, n := _parentheticalAction(parser, pos); n == nil {
					goto fail7
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos9:pos]
			}
			node = func(
				start, end int, x0 *Interjection, x1 *Parenthetical) Mod {
				return Mod(x1)
			}(
				start8, pos, label0, label1)
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// action
		{
			start11 := pos
			// x2:incidental
			{
				pos12 := pos
				// incidental
				if p, n := _incidentalAction(parser, pos); n == nil {
					goto fail10
				} else {
					label2 = *n
					pos = p
				}
				labels[2] = parser.text[pos12:pos]
			}
			node = func(
				start, end int, x0 *Interjection, x1 *Parenthetical, x2 *Incidental) Mod {
				return Mod(x2)
			}(
				start11, pos, label0, label1, label2)
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// action
		{
			start14 := pos
			// x3:vocative
			{
				pos15 := pos
				// vocative
				if p, n := _vocativeAction(parser, pos); n == nil {
					goto fail13
				} else {
					label3 = *n
					pos = p
				}
				labels[3] = parser.text[pos15:pos]
			}
			node = func(
				start, end int, x0 *Interjection, x1 *Parenthetical, x2 *Incidental, x3 *Vocative) Mod {
				return Mod(x3)
			}(
				start14, pos, label0, label1, label2, label3)
		}
		goto ok0
	fail13:
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

func _parentheticalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _parenthetical, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// p:parenthetical_1 s:_ ki:end_parenthetical
	// p:parenthetical_1
	{
		pos1 := pos
		// parenthetical_1
		if !_accept(parser, _parenthetical_1Accepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ki:end_parenthetical
	{
		pos3 := pos
		// end_parenthetical
		if !_accept(parser, _end_parentheticalAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _parenthetical, start, pos, perr)
fail:
	return _memoize(parser, _parenthetical, start, -1, perr)
}

func _parentheticalNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// p:parenthetical_1 s:_ ki:end_parenthetical
	// p:parenthetical_1
	{
		pos1 := pos
		// parenthetical_1
		if !_node(parser, _parenthetical_1Node, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ki:end_parenthetical
	{
		pos3 := pos
		// end_parenthetical
		if !_node(parser, _end_parentheticalNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _parentheticalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _parenthetical, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "parenthetical",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _parenthetical}
	// action
	// p:parenthetical_1 s:_ ki:end_parenthetical
	// p:parenthetical_1
	{
		pos1 := pos
		// parenthetical_1
		if !_fail(parser, _parenthetical_1Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// ki:end_parenthetical
	{
		pos3 := pos
		// end_parenthetical
		if !_fail(parser, _end_parentheticalFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _parentheticalAction(parser *_Parser, start int) (int, *(*Parenthetical)) {
	var labels [3]string
	use(labels)
	var label1 *Mod
	var label2 Word
	var label0 Parenthetical
	dp := parser.deltaPos[start][_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _parenthetical}
	n := parser.act[key]
	if n != nil {
		n := n.((*Parenthetical))
		return start + int(dp-1), &n
	}
	var node (*Parenthetical)
	pos := start
	// action
	{
		start0 := pos
		// p:parenthetical_1 s:_ ki:end_parenthetical
		// p:parenthetical_1
		{
			pos2 := pos
			// parenthetical_1
			if p, n := _parenthetical_1Action(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// ki:end_parenthetical
		{
			pos4 := pos
			// end_parenthetical
			if p, n := _end_parentheticalAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, ki Word, p Parenthetical, s *Mod) *Parenthetical {
			return (*Parenthetical)(endParen(p, s, ki))
		}(
			start0, pos, label2, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _parenthetical_1Accepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _parenthetical_1, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// kio:start_parenthetical s:_ d:discourse
	// kio:start_parenthetical
	{
		pos1 := pos
		// start_parenthetical
		if !_accept(parser, _start_parentheticalAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// d:discourse
	{
		pos3 := pos
		// discourse
		if !_accept(parser, _discourseAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _parenthetical_1, start, pos, perr)
fail:
	return _memoize(parser, _parenthetical_1, start, -1, perr)
}

func _parenthetical_1Node(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// kio:start_parenthetical s:_ d:discourse
	// kio:start_parenthetical
	{
		pos1 := pos
		// start_parenthetical
		if !_node(parser, _start_parentheticalNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// d:discourse
	{
		pos3 := pos
		// discourse
		if !_node(parser, _discourseNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _parenthetical_1Fail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _parenthetical_1, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "parenthetical_1",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _parenthetical_1}
	// action
	// kio:start_parenthetical s:_ d:discourse
	// kio:start_parenthetical
	{
		pos1 := pos
		// start_parenthetical
		if !_fail(parser, _start_parentheticalFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// d:discourse
	{
		pos3 := pos
		// discourse
		if !_fail(parser, _discourseFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _parenthetical_1Action(parser *_Parser, start int) (int, *Parenthetical) {
	var labels [3]string
	use(labels)
	var label0 Word
	var label1 *Mod
	var label2 []Node
	dp := parser.deltaPos[start][_parenthetical_1]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _parenthetical_1}
	n := parser.act[key]
	if n != nil {
		n := n.(Parenthetical)
		return start + int(dp-1), &n
	}
	var node Parenthetical
	pos := start
	// action
	{
		start0 := pos
		// kio:start_parenthetical s:_ d:discourse
		// kio:start_parenthetical
		{
			pos2 := pos
			// start_parenthetical
			if p, n := _start_parentheticalAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// d:discourse
		{
			pos4 := pos
			// discourse
			if p, n := _discourseAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, d []Node, kio Word, s *Mod) Parenthetical {
			return Parenthetical{KIO: *kio.mod(s), Discourse: d}
		}(
			start0, pos, label2, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _incidentalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _incidental, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// ju:start_incidental s:_ st:statement
	// ju:start_incidental
	{
		pos1 := pos
		// start_incidental
		if !_accept(parser, _start_incidentalAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// st:statement
	{
		pos3 := pos
		// statement
		if !_accept(parser, _statementAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _incidental, start, pos, perr)
fail:
	return _memoize(parser, _incidental, start, -1, perr)
}

func _incidentalNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// ju:start_incidental s:_ st:statement
	// ju:start_incidental
	{
		pos1 := pos
		// start_incidental
		if !_node(parser, _start_incidentalNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// st:statement
	{
		pos3 := pos
		// statement
		if !_node(parser, _statementNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _incidentalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _incidental, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "incidental",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _incidental}
	// action
	// ju:start_incidental s:_ st:statement
	// ju:start_incidental
	{
		pos1 := pos
		// start_incidental
		if !_fail(parser, _start_incidentalFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// st:statement
	{
		pos3 := pos
		// statement
		if !_fail(parser, _statementFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _incidentalAction(parser *_Parser, start int) (int, **Incidental) {
	var labels [3]string
	use(labels)
	var label0 Word
	var label1 *Mod
	var label2 Statement
	dp := parser.deltaPos[start][_incidental]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _incidental}
	n := parser.act[key]
	if n != nil {
		n := n.(*Incidental)
		return start + int(dp-1), &n
	}
	var node *Incidental
	pos := start
	// action
	{
		start0 := pos
		// ju:start_incidental s:_ st:statement
		// ju:start_incidental
		{
			pos2 := pos
			// start_incidental
			if p, n := _start_incidentalAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// st:statement
		{
			pos4 := pos
			// statement
			if p, n := _statementAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, ju Word, s *Mod, st Statement) *Incidental {
			return &Incidental{JU: *ju.mod(s), Statement: st}
		}(
			start0, pos, label0, label1, label2)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _vocativeAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _vocative, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// hu:vocative_marker s:_ a:argument
	// hu:vocative_marker
	{
		pos1 := pos
		// vocative_marker
		if !_accept(parser, _vocative_markerAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_accept(parser, __Accepts, &pos, &perr) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// a:argument
	{
		pos3 := pos
		// argument
		if !_accept(parser, _argumentAccepts, &pos, &perr) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	return _memoize(parser, _vocative, start, pos, perr)
fail:
	return _memoize(parser, _vocative, start, -1, perr)
}

func _vocativeNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// action
	// hu:vocative_marker s:_ a:argument
	// hu:vocative_marker
	{
		pos1 := pos
		// vocative_marker
		if !_node(parser, _vocative_markerNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_node(parser, __Node, node, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// a:argument
	{
		pos3 := pos
		// argument
		if !_node(parser, _argumentNode, node, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _vocativeFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _vocative, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "vocative",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _vocative}
	// action
	// hu:vocative_marker s:_ a:argument
	// hu:vocative_marker
	{
		pos1 := pos
		// vocative_marker
		if !_fail(parser, _vocative_markerFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:_
	{
		pos2 := pos
		// _
		if !_fail(parser, __Fail, errPos, failure, &pos) {
			goto fail
		}
		labels[1] = parser.text[pos2:pos]
	}
	// a:argument
	{
		pos3 := pos
		// argument
		if !_fail(parser, _argumentFail, errPos, failure, &pos) {
			goto fail
		}
		labels[2] = parser.text[pos3:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _vocativeAction(parser *_Parser, start int) (int, **Vocative) {
	var labels [3]string
	use(labels)
	var label0 Word
	var label1 *Mod
	var label2 Argument
	dp := parser.deltaPos[start][_vocative]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _vocative}
	n := parser.act[key]
	if n != nil {
		n := n.(*Vocative)
		return start + int(dp-1), &n
	}
	var node *Vocative
	pos := start
	// action
	{
		start0 := pos
		// hu:vocative_marker s:_ a:argument
		// hu:vocative_marker
		{
			pos2 := pos
			// vocative_marker
			if p, n := _vocative_markerAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:_
		{
			pos3 := pos
			// _
			if p, n := __Action(parser, pos); n == nil {
				goto fail
			} else {
				label1 = *n
				pos = p
			}
			labels[1] = parser.text[pos3:pos]
		}
		// a:argument
		{
			pos4 := pos
			// argument
			if p, n := _argumentAction(parser, pos); n == nil {
				goto fail
			} else {
				label2 = *n
				pos = p
			}
			labels[2] = parser.text[pos4:pos]
		}
		node = func(
			start, end int, a Argument, hu Word, s *Mod) *Vocative {
			return &Vocative{HU: *hu.mod(s), Argument: a}
		}(
			start0, pos, label2, label0, label1)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [2]string
	use(labels)
	if dp, de, ok := _memo(parser, _word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// w:(non_function_word/toned_function_word/neutral_function_word) s:space_word?
	// w:(non_function_word/toned_function_word/neutral_function_word)
	{
		pos1 := pos
		// (non_function_word/toned_function_word/neutral_function_word)
		// non_function_word/toned_function_word/neutral_function_word
		{
			pos5 := pos
			// non_function_word
			if !_accept(parser, _non_function_wordAccepts, &pos, &perr) {
				goto fail6
			}
			goto ok2
		fail6:
			pos = pos5
			// toned_function_word
			if !_accept(parser, _toned_function_wordAccepts, &pos, &perr) {
				goto fail7
			}
			goto ok2
		fail7:
			pos = pos5
			// neutral_function_word
			if !_accept(parser, _neutral_function_wordAccepts, &pos, &perr) {
				goto fail8
			}
			goto ok2
		fail8:
			pos = pos5
			goto fail
		ok2:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:space_word?
	{
		pos9 := pos
		// space_word?
		{
			pos11 := pos
			// space_word
			if !_accept(parser, _space_wordAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok13
		fail12:
			pos = pos11
		ok13:
		}
		labels[1] = parser.text[pos9:pos]
	}
	return _memoize(parser, _word, start, pos, perr)
fail:
	return _memoize(parser, _word, start, -1, perr)
}

func _wordNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [2]string
	use(labels)
	dp := parser.deltaPos[start][_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _word}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "word"}
	// action
	// w:(non_function_word/toned_function_word/neutral_function_word) s:space_word?
	// w:(non_function_word/toned_function_word/neutral_function_word)
	{
		pos1 := pos
		// (non_function_word/toned_function_word/neutral_function_word)
		{
			nkids2 := len(node.Kids)
			pos03 := pos
			// non_function_word/toned_function_word/neutral_function_word
			{
				pos7 := pos
				nkids5 := len(node.Kids)
				// non_function_word
				if !_node(parser, _non_function_wordNode, node, &pos) {
					goto fail8
				}
				goto ok4
			fail8:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				// toned_function_word
				if !_node(parser, _toned_function_wordNode, node, &pos) {
					goto fail9
				}
				goto ok4
			fail9:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				// neutral_function_word
				if !_node(parser, _neutral_function_wordNode, node, &pos) {
					goto fail10
				}
				goto ok4
			fail10:
				node.Kids = node.Kids[:nkids5]
				pos = pos7
				goto fail
			ok4:
			}
			sub := _sub(parser, pos03, pos, node.Kids[nkids2:])
			node.Kids = append(node.Kids[:nkids2], sub)
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:space_word?
	{
		pos11 := pos
		// space_word?
		{
			nkids12 := len(node.Kids)
			pos13 := pos
			// space_word
			if !_node(parser, _space_wordNode, node, &pos) {
				goto fail14
			}
			goto ok15
		fail14:
			node.Kids = node.Kids[:nkids12]
			pos = pos13
		ok15:
		}
		labels[1] = parser.text[pos11:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [2]string
	use(labels)
	pos, failure := _failMemo(parser, _word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _word}
	// action
	// w:(non_function_word/toned_function_word/neutral_function_word) s:space_word?
	// w:(non_function_word/toned_function_word/neutral_function_word)
	{
		pos1 := pos
		// (non_function_word/toned_function_word/neutral_function_word)
		// non_function_word/toned_function_word/neutral_function_word
		{
			pos5 := pos
			// non_function_word
			if !_fail(parser, _non_function_wordFail, errPos, failure, &pos) {
				goto fail6
			}
			goto ok2
		fail6:
			pos = pos5
			// toned_function_word
			if !_fail(parser, _toned_function_wordFail, errPos, failure, &pos) {
				goto fail7
			}
			goto ok2
		fail7:
			pos = pos5
			// neutral_function_word
			if !_fail(parser, _neutral_function_wordFail, errPos, failure, &pos) {
				goto fail8
			}
			goto ok2
		fail8:
			pos = pos5
			goto fail
		ok2:
		}
		labels[0] = parser.text[pos1:pos]
	}
	// s:space_word?
	{
		pos9 := pos
		// space_word?
		{
			pos11 := pos
			// space_word
			if !_fail(parser, _space_wordFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok13
		fail12:
			pos = pos11
		ok13:
		}
		labels[1] = parser.text[pos9:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _wordAction(parser *_Parser, start int) (int, *(*Word)) {
	var labels [2]string
	use(labels)
	var label1 *Mod
	var label0 string
	dp := parser.deltaPos[start][_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _word}
	n := parser.act[key]
	if n != nil {
		n := n.((*Word))
		return start + int(dp-1), &n
	}
	var node (*Word)
	pos := start
	// action
	{
		start0 := pos
		// w:(non_function_word/toned_function_word/neutral_function_word) s:space_word?
		// w:(non_function_word/toned_function_word/neutral_function_word)
		{
			pos2 := pos
			// (non_function_word/toned_function_word/neutral_function_word)
			// non_function_word/toned_function_word/neutral_function_word
			{
				pos6 := pos
				var node5 string
				// non_function_word
				if p, n := _non_function_wordAction(parser, pos); n == nil {
					goto fail7
				} else {
					label0 = *n
					pos = p
				}
				goto ok3
			fail7:
				label0 = node5
				pos = pos6
				// toned_function_word
				if p, n := _toned_function_wordAction(parser, pos); n == nil {
					goto fail8
				} else {
					label0 = *n
					pos = p
				}
				goto ok3
			fail8:
				label0 = node5
				pos = pos6
				// neutral_function_word
				if p, n := _neutral_function_wordAction(parser, pos); n == nil {
					goto fail9
				} else {
					label0 = *n
					pos = p
				}
				goto ok3
			fail9:
				label0 = node5
				pos = pos6
				goto fail
			ok3:
			}
			labels[0] = parser.text[pos2:pos]
		}
		// s:space_word?
		{
			pos10 := pos
			// space_word?
			{
				pos12 := pos
				label1 = new(Mod)
				// space_word
				if p, n := _space_wordAction(parser, pos); n == nil {
					goto fail13
				} else {
					*label1 = *n
					pos = p
				}
				goto ok14
			fail13:
				label1 = nil
				pos = pos12
			ok14:
			}
			labels[1] = parser.text[pos10:pos]
		}
		node = func(
			start, end int, s *Mod, w string) *Word {
			return (*Word)(Word{S: start, E: end, T: w}.mod(s))
		}(
			start0, pos, label1, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _non_function_wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _non_function_word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// initial_syllable compound_syllable+/!function_word initial_syllable
	{
		pos3 := pos
		// initial_syllable compound_syllable+
		// initial_syllable
		if !_accept(parser, _initial_syllableAccepts, &pos, &perr) {
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
		// !function_word initial_syllable
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
		// initial_syllable
		if !_accept(parser, _initial_syllableAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _non_function_word, start, pos, perr)
fail:
	return _memoize(parser, _non_function_word, start, -1, perr)
}

func _non_function_wordNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_non_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _non_function_word}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "non_function_word"}
	// initial_syllable compound_syllable+/!function_word initial_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// initial_syllable compound_syllable+
		// initial_syllable
		if !_node(parser, _initial_syllableNode, node, &pos) {
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
		// !function_word initial_syllable
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
		// initial_syllable
		if !_node(parser, _initial_syllableNode, node, &pos) {
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

func _non_function_wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _non_function_word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "non_function_word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _non_function_word}
	// initial_syllable compound_syllable+/!function_word initial_syllable
	{
		pos3 := pos
		// initial_syllable compound_syllable+
		// initial_syllable
		if !_fail(parser, _initial_syllableFail, errPos, failure, &pos) {
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
		// !function_word initial_syllable
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
		// initial_syllable
		if !_fail(parser, _initial_syllableFail, errPos, failure, &pos) {
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

func _non_function_wordAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_non_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _non_function_word}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// initial_syllable compound_syllable+/!function_word initial_syllable
	{
		pos3 := pos
		var node2 string
		// initial_syllable compound_syllable+
		{
			var node5 string
			// initial_syllable
			if p, n := _initial_syllableAction(parser, pos); n == nil {
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
		// !function_word initial_syllable
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
			// initial_syllable
			if p, n := _initial_syllableAction(parser, pos); n == nil {
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

func _neutral_function_wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _neutral_function_word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// !toned_function_word &function_word neutral_syllable
	// !toned_function_word
	{
		pos2 := pos
		perr4 := perr
		// toned_function_word
		if !_accept(parser, _toned_function_wordAccepts, &pos, &perr) {
			goto ok1
		}
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// &function_word
	{
		pos6 := pos
		perr8 := perr
		// function_word
		if !_accept(parser, _function_wordAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos6
		perr = _max(perr8, pos)
		goto fail
	ok5:
		pos = pos6
		perr = perr8
	}
	// neutral_syllable
	if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _neutral_function_word, start, pos, perr)
fail:
	return _memoize(parser, _neutral_function_word, start, -1, perr)
}

func _neutral_function_wordNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_neutral_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _neutral_function_word}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "neutral_function_word"}
	// !toned_function_word &function_word neutral_syllable
	// !toned_function_word
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// toned_function_word
		if !_node(parser, _toned_function_wordNode, node, &pos) {
			goto ok1
		}
		pos = pos2
		node.Kids = node.Kids[:nkids3]
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// &function_word
	{
		pos6 := pos
		nkids7 := len(node.Kids)
		// function_word
		if !_node(parser, _function_wordNode, node, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos6
		goto fail
	ok5:
		pos = pos6
		node.Kids = node.Kids[:nkids7]
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

func _neutral_function_wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _neutral_function_word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "neutral_function_word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _neutral_function_word}
	// !toned_function_word &function_word neutral_syllable
	// !toned_function_word
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// toned_function_word
		if !_fail(parser, _toned_function_wordFail, errPos, failure, &pos) {
			goto ok1
		}
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "!toned_function_word",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// &function_word
	{
		pos6 := pos
		nkids7 := len(failure.Kids)
		// function_word
		if !_fail(parser, _function_wordFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos6
		failure.Kids = failure.Kids[:nkids7]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&function_word",
			})
		}
		goto fail
	ok5:
		pos = pos6
		failure.Kids = failure.Kids[:nkids7]
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

func _neutral_function_wordAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_neutral_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _neutral_function_word}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// !toned_function_word &function_word neutral_syllable
	{
		var node0 string
		// !toned_function_word
		{
			pos2 := pos
			// toned_function_word
			if p, n := _toned_function_wordAction(parser, pos); n == nil {
				goto ok1
			} else {
				pos = p
			}
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// &function_word
		{
			pos6 := pos
			// function_word
			if p, n := _function_wordAction(parser, pos); n == nil {
				goto fail9
			} else {
				pos = p
			}
			goto ok5
		fail9:
			pos = pos6
			goto fail
		ok5:
			pos = pos6
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

func _toned_function_wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _toned_function_word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// !neutral_syllable (LU/MI/MO/PO)
	// !neutral_syllable
	{
		pos2 := pos
		perr4 := perr
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto ok1
		}
		pos = pos2
		perr = _max(perr4, pos)
		goto fail
	ok1:
		pos = pos2
		perr = perr4
	}
	// (LU/MI/MO/PO)
	// LU/MI/MO/PO
	{
		pos8 := pos
		// LU
		if !_accept(parser, _LUAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// MI
		if !_accept(parser, _MIAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// MO
		if !_accept(parser, _MOAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok5
	fail12:
		pos = pos8
		goto fail
	ok5:
	}
	return _memoize(parser, _toned_function_word, start, pos, perr)
fail:
	return _memoize(parser, _toned_function_word, start, -1, perr)
}

func _toned_function_wordNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_toned_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _toned_function_word}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "toned_function_word"}
	// !neutral_syllable (LU/MI/MO/PO)
	// !neutral_syllable
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto ok1
		}
		pos = pos2
		node.Kids = node.Kids[:nkids3]
		goto fail
	ok1:
		pos = pos2
		node.Kids = node.Kids[:nkids3]
	}
	// (LU/MI/MO/PO)
	{
		nkids5 := len(node.Kids)
		pos06 := pos
		// LU/MI/MO/PO
		{
			pos10 := pos
			nkids8 := len(node.Kids)
			// LU
			if !_node(parser, _LUNode, node, &pos) {
				goto fail11
			}
			goto ok7
		fail11:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// MI
			if !_node(parser, _MINode, node, &pos) {
				goto fail12
			}
			goto ok7
		fail12:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// MO
			if !_node(parser, _MONode, node, &pos) {
				goto fail13
			}
			goto ok7
		fail13:
			node.Kids = node.Kids[:nkids8]
			pos = pos10
			// PO
			if !_node(parser, _PONode, node, &pos) {
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

func _toned_function_wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _toned_function_word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "toned_function_word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _toned_function_word}
	// !neutral_syllable (LU/MI/MO/PO)
	// !neutral_syllable
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto ok1
		}
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "!neutral_syllable",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// (LU/MI/MO/PO)
	// LU/MI/MO/PO
	{
		pos8 := pos
		// LU
		if !_fail(parser, _LUFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok5
	fail9:
		pos = pos8
		// MI
		if !_fail(parser, _MIFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok5
	fail10:
		pos = pos8
		// MO
		if !_fail(parser, _MOFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok5
	fail11:
		pos = pos8
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
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

func _toned_function_wordAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_toned_function_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _toned_function_word}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// !neutral_syllable (LU/MI/MO/PO)
	{
		var node0 string
		// !neutral_syllable
		{
			pos2 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto ok1
			} else {
				pos = p
			}
			pos = pos2
			goto fail
		ok1:
			pos = pos2
			node0 = ""
		}
		node, node0 = node+node0, ""
		// (LU/MI/MO/PO)
		// LU/MI/MO/PO
		{
			pos8 := pos
			var node7 string
			// LU
			if p, n := _LUAction(parser, pos); n == nil {
				goto fail9
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail9:
			node0 = node7
			pos = pos8
			// MI
			if p, n := _MIAction(parser, pos); n == nil {
				goto fail10
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail10:
			node0 = node7
			pos = pos8
			// MO
			if p, n := _MOAction(parser, pos); n == nil {
				goto fail11
			} else {
				node0 = *n
				pos = p
			}
			goto ok5
		fail11:
			node0 = node7
			pos = pos8
			// PO
			if p, n := _POAction(parser, pos); n == nil {
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

func _initial_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _initial_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// arg_syllable/relative_syllable/verb_syllable/content_syllable/preposition_syllable/adverb_syllable
	{
		pos3 := pos
		// arg_syllable
		if !_accept(parser, _arg_syllableAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_syllable
		if !_accept(parser, _relative_syllableAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// verb_syllable
		if !_accept(parser, _verb_syllableAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// content_syllable
		if !_accept(parser, _content_syllableAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// preposition_syllable
		if !_accept(parser, _preposition_syllableAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// adverb_syllable
		if !_accept(parser, _adverb_syllableAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _initial_syllable, start, pos, perr)
fail:
	return _memoize(parser, _initial_syllable, start, -1, perr)
}

func _initial_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_initial_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _initial_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "initial_syllable"}
	// arg_syllable/relative_syllable/verb_syllable/content_syllable/preposition_syllable/adverb_syllable
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// arg_syllable
		if !_node(parser, _arg_syllableNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_syllable
		if !_node(parser, _relative_syllableNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// verb_syllable
		if !_node(parser, _verb_syllableNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// content_syllable
		if !_node(parser, _content_syllableNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_syllable
		if !_node(parser, _preposition_syllableNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_syllable
		if !_node(parser, _adverb_syllableNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
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

func _initial_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _initial_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "initial_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _initial_syllable}
	// arg_syllable/relative_syllable/verb_syllable/content_syllable/preposition_syllable/adverb_syllable
	{
		pos3 := pos
		// arg_syllable
		if !_fail(parser, _arg_syllableFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// relative_syllable
		if !_fail(parser, _relative_syllableFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// verb_syllable
		if !_fail(parser, _verb_syllableFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// content_syllable
		if !_fail(parser, _content_syllableFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// preposition_syllable
		if !_fail(parser, _preposition_syllableFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// adverb_syllable
		if !_fail(parser, _adverb_syllableFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
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

func _initial_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_initial_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _initial_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// arg_syllable/relative_syllable/verb_syllable/content_syllable/preposition_syllable/adverb_syllable
	{
		pos3 := pos
		var node2 string
		// arg_syllable
		if p, n := _arg_syllableAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// relative_syllable
		if p, n := _relative_syllableAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// verb_syllable
		if p, n := _verb_syllableAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// content_syllable
		if p, n := _content_syllableAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// preposition_syllable
		if p, n := _preposition_syllableAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// adverb_syllable
		if p, n := _adverb_syllableAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
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

func _prefixAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _prefix, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &MU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _prefix, start, pos, perr)
fail:
	return _memoize(parser, _prefix, start, -1, perr)
}

func _prefixNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &MU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _prefixFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _prefix, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "prefix",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _prefix}
	// action
	// &MU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _prefixAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_prefix]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _prefix}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &MU w:neutral_syllable
		// &MU
		{
			pos3 := pos
			// MU
			if p, n := _MUAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _focusAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _focus, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &KU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _focus, start, pos, perr)
fail:
	return _memoize(parser, _focus, start, -1, perr)
}

func _focusNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &KU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _focusFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _focus, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "focus",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _focus}
	// action
	// &KU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _focusAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_focus]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _focus}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &KU w:neutral_syllable
		// &KU
		{
			pos3 := pos
			// KU
			if p, n := _KUAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_quoteAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _end_quote, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &TEO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _end_quote, start, pos, perr)
fail:
	return _memoize(parser, _end_quote, start, -1, perr)
}

func _end_quoteNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &TEO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_quoteFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _end_quote, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_quote",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_quote}
	// action
	// &TEO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_quoteAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_end_quote]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_quote}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &TEO w:neutral_syllable
		// &TEO
		{
			pos3 := pos
			// TEO
			if p, n := _TEOAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_predicatizerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _end_predicatizer, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &GA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _end_predicatizer, start, pos, perr)
fail:
	return _memoize(parser, _end_predicatizer, start, -1, perr)
}

func _end_predicatizerNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &GA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_predicatizerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _end_predicatizer, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_predicatizer",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_predicatizer}
	// action
	// &GA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_predicatizerAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_end_predicatizer]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_predicatizer}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &GA w:neutral_syllable
		// &GA
		{
			pos3 := pos
			// GA
			if p, n := _GAAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_statementAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _end_statement, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &NA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _end_statement, start, pos, perr)
fail:
	return _memoize(parser, _end_statement, start, -1, perr)
}

func _end_statementNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &NA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_statementFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _end_statement, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_statement",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_statement}
	// action
	// &NA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_statementAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_end_statement]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_statement}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &NA w:neutral_syllable
		// &NA
		{
			pos3 := pos
			// NA
			if p, n := _NAAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _sentence_prefixAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _sentence_prefix, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &KEO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _sentence_prefix, start, pos, perr)
fail:
	return _memoize(parser, _sentence_prefix, start, -1, perr)
}

func _sentence_prefixNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &KEO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _sentence_prefixFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _sentence_prefix, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "sentence_prefix",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _sentence_prefix}
	// action
	// &KEO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _sentence_prefixAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_sentence_prefix]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _sentence_prefix}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &KEO w:neutral_syllable
		// &KEO
		{
			pos3 := pos
			// KEO
			if p, n := _KEOAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_prenexAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _end_prenex, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &BI w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _end_prenex, start, pos, perr)
fail:
	return _memoize(parser, _end_prenex, start, -1, perr)
}

func _end_prenexNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &BI w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_prenexFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _end_prenex, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_prenex",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_prenex}
	// action
	// &BI w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_prenexAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_end_prenex]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_prenex}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &BI w:neutral_syllable
		// &BI
		{
			pos3 := pos
			// BI
			if p, n := _BIAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _start_incidentalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _start_incidental, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &JU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _start_incidental, start, pos, perr)
fail:
	return _memoize(parser, _start_incidental, start, -1, perr)
}

func _start_incidentalNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &JU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _start_incidentalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _start_incidental, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "start_incidental",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _start_incidental}
	// action
	// &JU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _start_incidentalAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_start_incidental]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _start_incidental}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &JU w:neutral_syllable
		// &JU
		{
			pos3 := pos
			// JU
			if p, n := _JUAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _start_parentheticalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _start_parenthetical, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &KIO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _start_parenthetical, start, pos, perr)
fail:
	return _memoize(parser, _start_parenthetical, start, -1, perr)
}

func _start_parentheticalNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &KIO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _start_parentheticalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _start_parenthetical, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "start_parenthetical",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _start_parenthetical}
	// action
	// &KIO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _start_parentheticalAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_start_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _start_parenthetical}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &KIO w:neutral_syllable
		// &KIO
		{
			pos3 := pos
			// KIO
			if p, n := _KIOAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _end_parentheticalAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _end_parenthetical, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &KI w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _end_parenthetical, start, pos, perr)
fail:
	return _memoize(parser, _end_parenthetical, start, -1, perr)
}

func _end_parentheticalNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &KI w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _end_parentheticalFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _end_parenthetical, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "end_parenthetical",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _end_parenthetical}
	// action
	// &KI w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _end_parentheticalAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_end_parenthetical]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _end_parenthetical}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &KI w:neutral_syllable
		// &KI
		{
			pos3 := pos
			// KI
			if p, n := _KIAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _vocative_markerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _vocative_marker, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &HU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _vocative_marker, start, pos, perr)
fail:
	return _memoize(parser, _vocative_marker, start, -1, perr)
}

func _vocative_markerNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &HU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _vocative_markerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _vocative_marker, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "vocative_marker",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _vocative_marker}
	// action
	// &HU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _vocative_markerAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_vocative_marker]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _vocative_marker}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &HU w:neutral_syllable
		// &HU
		{
			pos3 := pos
			// HU
			if p, n := _HUAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _linking_wordAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _linking_word, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &GO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _linking_word, start, pos, perr)
fail:
	return _memoize(parser, _linking_word, start, -1, perr)
}

func _linking_wordNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &GO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _linking_wordFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _linking_word, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "linking_word",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _linking_word}
	// action
	// &GO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _linking_wordAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_linking_word]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _linking_word}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &GO w:neutral_syllable
		// &GO
		{
			pos3 := pos
			// GO
			if p, n := _GOAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _connectiveAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _connective, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &RA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _connective, start, pos, perr)
fail:
	return _memoize(parser, _connective, start, -1, perr)
}

func _connectiveNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &RA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _connectiveFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _connective, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "connective",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _connective}
	// action
	// &RA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _connectiveAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_connective]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _connective}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &RA w:neutral_syllable
		// &RA
		{
			pos3 := pos
			// RA
			if p, n := _RAAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _illocutionaryAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [3]string
	use(labels)
	if dp, de, ok := _memo(parser, _illocutionary, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// &DA w:neutral_syllable {…}/moq:moq s:(questions?) {…}
	{
		pos3 := pos
		// action
		// &DA w:neutral_syllable
		// &DA
		{
			pos7 := pos
			perr9 := perr
			// DA
			if !_accept(parser, _DAAccepts, &pos, &perr) {
				goto fail10
			}
			goto ok6
		fail10:
			pos = pos7
			perr = _max(perr9, pos)
			goto fail4
		ok6:
			pos = pos7
			perr = perr9
		}
		// w:neutral_syllable
		{
			pos11 := pos
			// neutral_syllable
			if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
				goto fail4
			}
			labels[0] = parser.text[pos11:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// moq:moq s:(questions?)
		// moq:moq
		{
			pos14 := pos
			// moq
			if !_accept(parser, _moqAccepts, &pos, &perr) {
				goto fail12
			}
			labels[1] = parser.text[pos14:pos]
		}
		// s:(questions?)
		{
			pos15 := pos
			// (questions?)
			// questions?
			{
				pos17 := pos
				// questions
				if !_accept(parser, _questionsAccepts, &pos, &perr) {
					goto fail18
				}
				goto ok19
			fail18:
				pos = pos17
			ok19:
			}
			labels[2] = parser.text[pos15:pos]
		}
		goto ok0
	fail12:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _illocutionary, start, pos, perr)
fail:
	return _memoize(parser, _illocutionary, start, -1, perr)
}

func _illocutionaryNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [3]string
	use(labels)
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
	// &DA w:neutral_syllable {…}/moq:moq s:(questions?) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// action
		// &DA w:neutral_syllable
		// &DA
		{
			pos7 := pos
			nkids8 := len(node.Kids)
			// DA
			if !_node(parser, _DANode, node, &pos) {
				goto fail10
			}
			goto ok6
		fail10:
			pos = pos7
			goto fail4
		ok6:
			pos = pos7
			node.Kids = node.Kids[:nkids8]
		}
		// w:neutral_syllable
		{
			pos11 := pos
			// neutral_syllable
			if !_node(parser, _neutral_syllableNode, node, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos11:pos]
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// moq:moq s:(questions?)
		// moq:moq
		{
			pos14 := pos
			// moq
			if !_node(parser, _moqNode, node, &pos) {
				goto fail12
			}
			labels[1] = parser.text[pos14:pos]
		}
		// s:(questions?)
		{
			pos15 := pos
			// (questions?)
			{
				nkids16 := len(node.Kids)
				pos017 := pos
				// questions?
				{
					nkids18 := len(node.Kids)
					pos19 := pos
					// questions
					if !_node(parser, _questionsNode, node, &pos) {
						goto fail20
					}
					goto ok21
				fail20:
					node.Kids = node.Kids[:nkids18]
					pos = pos19
				ok21:
				}
				sub := _sub(parser, pos017, pos, node.Kids[nkids16:])
				node.Kids = append(node.Kids[:nkids16], sub)
			}
			labels[2] = parser.text[pos15:pos]
		}
		goto ok0
	fail12:
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

func _illocutionaryFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [3]string
	use(labels)
	pos, failure := _failMemo(parser, _illocutionary, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "illocutionary",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _illocutionary}
	// &DA w:neutral_syllable {…}/moq:moq s:(questions?) {…}
	{
		pos3 := pos
		// action
		// &DA w:neutral_syllable
		// &DA
		{
			pos7 := pos
			nkids8 := len(failure.Kids)
			// DA
			if !_fail(parser, _DAFail, errPos, failure, &pos) {
				goto fail10
			}
			goto ok6
		fail10:
			pos = pos7
			failure.Kids = failure.Kids[:nkids8]
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "&DA",
				})
			}
			goto fail4
		ok6:
			pos = pos7
			failure.Kids = failure.Kids[:nkids8]
		}
		// w:neutral_syllable
		{
			pos11 := pos
			// neutral_syllable
			if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
				goto fail4
			}
			labels[0] = parser.text[pos11:pos]
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// moq:moq s:(questions?)
		// moq:moq
		{
			pos14 := pos
			// moq
			if !_fail(parser, _moqFail, errPos, failure, &pos) {
				goto fail12
			}
			labels[1] = parser.text[pos14:pos]
		}
		// s:(questions?)
		{
			pos15 := pos
			// (questions?)
			// questions?
			{
				pos17 := pos
				// questions
				if !_fail(parser, _questionsFail, errPos, failure, &pos) {
					goto fail18
				}
				goto ok19
			fail18:
				pos = pos17
			ok19:
			}
			labels[2] = parser.text[pos15:pos]
		}
		goto ok0
	fail12:
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

func _illocutionaryAction(parser *_Parser, start int) (int, *Word) {
	var labels [3]string
	use(labels)
	var label0 string
	var label1 Word
	var label2 *Mod
	dp := parser.deltaPos[start][_illocutionary]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _illocutionary}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// &DA w:neutral_syllable {…}/moq:moq s:(questions?) {…}
	{
		pos3 := pos
		var node2 Word
		// action
		{
			start5 := pos
			// &DA w:neutral_syllable
			// &DA
			{
				pos8 := pos
				// DA
				if p, n := _DAAction(parser, pos); n == nil {
					goto fail11
				} else {
					pos = p
				}
				goto ok7
			fail11:
				pos = pos8
				goto fail4
			ok7:
				pos = pos8
			}
			// w:neutral_syllable
			{
				pos12 := pos
				// neutral_syllable
				if p, n := _neutral_syllableAction(parser, pos); n == nil {
					goto fail4
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos12:pos]
			}
			node = func(
				start, end int, w string) Word {
				return Word{S: start, E: end, T: w}
			}(
				start5, pos, label0)
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start14 := pos
			// moq:moq s:(questions?)
			// moq:moq
			{
				pos16 := pos
				// moq
				if p, n := _moqAction(parser, pos); n == nil {
					goto fail13
				} else {
					label1 = *n
					pos = p
				}
				labels[1] = parser.text[pos16:pos]
			}
			// s:(questions?)
			{
				pos17 := pos
				// (questions?)
				// questions?
				{
					pos19 := pos
					label2 = new(Mod)
					// questions
					if p, n := _questionsAction(parser, pos); n == nil {
						goto fail20
					} else {
						*label2 = *n
						pos = p
					}
					goto ok21
				fail20:
					label2 = nil
					pos = pos19
				ok21:
				}
				labels[2] = parser.text[pos17:pos]
			}
			node = func(
				start, end int, moq Word, s *Mod, w string) Word {
				return Word(*moq.mod(s))
			}(
				start14, pos, label1, label2, label0)
		}
		goto ok0
	fail13:
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

func _moqAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _moq, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &MOQ w:moq_syllable
	// &MOQ
	{
		pos2 := pos
		perr4 := perr
		// MOQ
		if !_accept(parser, _MOQAccepts, &pos, &perr) {
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
	// w:moq_syllable
	{
		pos6 := pos
		// moq_syllable
		if !_accept(parser, _moq_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _moq, start, pos, perr)
fail:
	return _memoize(parser, _moq, start, -1, perr)
}

func _moqNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_moq]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _moq}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "moq"}
	// action
	// &MOQ w:moq_syllable
	// &MOQ
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// MOQ
		if !_node(parser, _MOQNode, node, &pos) {
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
	// w:moq_syllable
	{
		pos6 := pos
		// moq_syllable
		if !_node(parser, _moq_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _moqFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _moq, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "moq",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _moq}
	// action
	// &MOQ w:moq_syllable
	// &MOQ
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// MOQ
		if !_fail(parser, _MOQFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&MOQ",
			})
		}
		goto fail
	ok1:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
	}
	// w:moq_syllable
	{
		pos6 := pos
		// moq_syllable
		if !_fail(parser, _moq_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _moqAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_moq]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _moq}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &MOQ w:moq_syllable
		// &MOQ
		{
			pos3 := pos
			// MOQ
			if p, n := _MOQAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:moq_syllable
		{
			pos7 := pos
			// moq_syllable
			if p, n := _moq_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _quantifierAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _quantifier, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &TU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _quantifier, start, pos, perr)
fail:
	return _memoize(parser, _quantifier, start, -1, perr)
}

func _quantifierNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &TU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _quantifierFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _quantifier, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "quantifier",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _quantifier}
	// action
	// &TU w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _quantifierAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_quantifier]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _quantifier}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &TU w:neutral_syllable
		// &TU
		{
			pos3 := pos
			// TU
			if p, n := _TUAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _interjectionAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _interjection, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &HA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _interjection, start, pos, perr)
fail:
	return _memoize(parser, _interjection, start, -1, perr)
}

func _interjectionNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &HA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _interjectionFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _interjection, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "interjection",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _interjection}
	// action
	// &HA w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _interjectionAction(parser *_Parser, start int) (int, **Interjection) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_interjection]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _interjection}
	n := parser.act[key]
	if n != nil {
		n := n.(*Interjection)
		return start + int(dp-1), &n
	}
	var node *Interjection
	pos := start
	// action
	{
		start0 := pos
		// &HA w:neutral_syllable
		// &HA
		{
			pos3 := pos
			// HA
			if p, n := _HAAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) *Interjection {
			return &Interjection{S: start, E: end, T: w}
		}(
			start0, pos, label0)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _forethought_markerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _forethought_marker, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// &TO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_accept(parser, _neutral_syllableAccepts, &pos, &perr) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	return _memoize(parser, _forethought_marker, start, pos, perr)
fail:
	return _memoize(parser, _forethought_marker, start, -1, perr)
}

func _forethought_markerNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
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
	// action
	// &TO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_node(parser, _neutral_syllableNode, node, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _forethought_markerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _forethought_marker, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "forethought_marker",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _forethought_marker}
	// action
	// &TO w:neutral_syllable
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
	// w:neutral_syllable
	{
		pos6 := pos
		// neutral_syllable
		if !_fail(parser, _neutral_syllableFail, errPos, failure, &pos) {
			goto fail
		}
		labels[0] = parser.text[pos6:pos]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _forethought_markerAction(parser *_Parser, start int) (int, *Word) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_forethought_marker]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _forethought_marker}
	n := parser.act[key]
	if n != nil {
		n := n.(Word)
		return start + int(dp-1), &n
	}
	var node Word
	pos := start
	// action
	{
		start0 := pos
		// &TO w:neutral_syllable
		// &TO
		{
			pos3 := pos
			// TO
			if p, n := _TOAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok2
		fail6:
			pos = pos3
			goto fail
		ok2:
			pos = pos3
		}
		// w:neutral_syllable
		{
			pos7 := pos
			// neutral_syllable
			if p, n := _neutral_syllableAction(parser, pos); n == nil {
				goto fail
			} else {
				label0 = *n
				pos = p
			}
			labels[0] = parser.text[pos7:pos]
		}
		node = func(
			start, end int, w string) Word {
			return Word{S: start, E: end, T: w}
		}(
			start0, pos, label0)
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
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/MOQ/NA/PO/RA/TU/TO/TEO
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
		// MOQ
		if !_accept(parser, _MOQAccepts, &pos, &perr) {
			goto fail19
		}
		goto ok0
	fail19:
		pos = pos3
		// NA
		if !_accept(parser, _NAAccepts, &pos, &perr) {
			goto fail20
		}
		goto ok0
	fail20:
		pos = pos3
		// PO
		if !_accept(parser, _POAccepts, &pos, &perr) {
			goto fail21
		}
		goto ok0
	fail21:
		pos = pos3
		// RA
		if !_accept(parser, _RAAccepts, &pos, &perr) {
			goto fail22
		}
		goto ok0
	fail22:
		pos = pos3
		// TU
		if !_accept(parser, _TUAccepts, &pos, &perr) {
			goto fail23
		}
		goto ok0
	fail23:
		pos = pos3
		// TO
		if !_accept(parser, _TOAccepts, &pos, &perr) {
			goto fail24
		}
		goto ok0
	fail24:
		pos = pos3
		// TEO
		if !_accept(parser, _TEOAccepts, &pos, &perr) {
			goto fail25
		}
		goto ok0
	fail25:
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
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/MOQ/NA/PO/RA/TU/TO/TEO
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
		// MOQ
		if !_node(parser, _MOQNode, node, &pos) {
			goto fail19
		}
		goto ok0
	fail19:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// NA
		if !_node(parser, _NANode, node, &pos) {
			goto fail20
		}
		goto ok0
	fail20:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// PO
		if !_node(parser, _PONode, node, &pos) {
			goto fail21
		}
		goto ok0
	fail21:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// RA
		if !_node(parser, _RANode, node, &pos) {
			goto fail22
		}
		goto ok0
	fail22:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// TU
		if !_node(parser, _TUNode, node, &pos) {
			goto fail23
		}
		goto ok0
	fail23:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// TO
		if !_node(parser, _TONode, node, &pos) {
			goto fail24
		}
		goto ok0
	fail24:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// TEO
		if !_node(parser, _TEONode, node, &pos) {
			goto fail25
		}
		goto ok0
	fail25:
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
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/MOQ/NA/PO/RA/TU/TO/TEO
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
		// MOQ
		if !_fail(parser, _MOQFail, errPos, failure, &pos) {
			goto fail19
		}
		goto ok0
	fail19:
		pos = pos3
		// NA
		if !_fail(parser, _NAFail, errPos, failure, &pos) {
			goto fail20
		}
		goto ok0
	fail20:
		pos = pos3
		// PO
		if !_fail(parser, _POFail, errPos, failure, &pos) {
			goto fail21
		}
		goto ok0
	fail21:
		pos = pos3
		// RA
		if !_fail(parser, _RAFail, errPos, failure, &pos) {
			goto fail22
		}
		goto ok0
	fail22:
		pos = pos3
		// TU
		if !_fail(parser, _TUFail, errPos, failure, &pos) {
			goto fail23
		}
		goto ok0
	fail23:
		pos = pos3
		// TO
		if !_fail(parser, _TOFail, errPos, failure, &pos) {
			goto fail24
		}
		goto ok0
	fail24:
		pos = pos3
		// TEO
		if !_fail(parser, _TEOFail, errPos, failure, &pos) {
			goto fail25
		}
		goto ok0
	fail25:
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
	// BI/DA/GA/GO/HA/HU/JU/KU/KI/KIO/KEO/LU/MU/MI/MO/MOQ/NA/PO/RA/TU/TO/TEO
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
		// MOQ
		if p, n := _MOQAction(parser, pos); n == nil {
			goto fail19
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail19:
		node = node2
		pos = pos3
		// NA
		if p, n := _NAAction(parser, pos); n == nil {
			goto fail20
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail20:
		node = node2
		pos = pos3
		// PO
		if p, n := _POAction(parser, pos); n == nil {
			goto fail21
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail21:
		node = node2
		pos = pos3
		// RA
		if p, n := _RAAction(parser, pos); n == nil {
			goto fail22
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail22:
		node = node2
		pos = pos3
		// TU
		if p, n := _TUAction(parser, pos); n == nil {
			goto fail23
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail23:
		node = node2
		pos = pos3
		// TO
		if p, n := _TOAction(parser, pos); n == nil {
			goto fail24
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail24:
		node = node2
		pos = pos3
		// TEO
		if p, n := _TEOAction(parser, pos); n == nil {
			goto fail25
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail25:
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
	// (d A/k A/s O/b A) &(tone? boundary)
	// (d A/k A/s O/b A)
	// d A/k A/s O/b A
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
		// k A
		// k
		if !_accept(parser, _kAccepts, &pos, &perr) {
			goto fail7
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// s O
		// s
		if !_accept(parser, _sAccepts, &pos, &perr) {
			goto fail9
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// b A
		// b
		if !_accept(parser, _bAccepts, &pos, &perr) {
			goto fail11
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
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
	// (d A/k A/s O/b A) &(tone? boundary)
	// (d A/k A/s O/b A)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// d A/k A/s O/b A
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
			// k A
			// k
			if !_node(parser, _kNode, node, &pos) {
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
			// s O
			// s
			if !_node(parser, _sNode, node, &pos) {
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
			// b A
			// b
			if !_node(parser, _bNode, node, &pos) {
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
	// (d A/k A/s O/b A) &(tone? boundary)
	// (d A/k A/s O/b A)
	// d A/k A/s O/b A
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
		// k A
		// k
		if !_fail(parser, _kFail, errPos, failure, &pos) {
			goto fail7
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// s O
		// s
		if !_fail(parser, _sFail, errPos, failure, &pos) {
			goto fail9
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// b A
		// b
		if !_fail(parser, _bFail, errPos, failure, &pos) {
			goto fail11
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
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
	// (d A/k A/s O/b A) &(tone? boundary)
	{
		var node0 string
		// (d A/k A/s O/b A)
		// d A/k A/s O/b A
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
			// k A
			{
				var node8 string
				// k
				if p, n := _kAction(parser, pos); n == nil {
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
			// s O
			{
				var node10 string
				// s
				if p, n := _sAction(parser, pos); n == nil {
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
			// b A
			{
				var node12 string
				// b
				if p, n := _bAction(parser, pos); n == nil {
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

func _MOQAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _MOQ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// m O q &((questions/tone)? boundary)
	// m
	if !_accept(parser, _mAccepts, &pos, &perr) {
		goto fail
	}
	// O
	if !_accept(parser, _OAccepts, &pos, &perr) {
		goto fail
	}
	// q
	if !_accept(parser, _qAccepts, &pos, &perr) {
		goto fail
	}
	// &((questions/tone)? boundary)
	{
		pos2 := pos
		perr4 := perr
		// ((questions/tone)? boundary)
		// (questions/tone)? boundary
		// (questions/tone)?
		{
			pos8 := pos
			// (questions/tone)
			// questions/tone
			{
				pos13 := pos
				// questions
				if !_accept(parser, _questionsAccepts, &pos, &perr) {
					goto fail14
				}
				goto ok10
			fail14:
				pos = pos13
				// tone
				if !_accept(parser, _toneAccepts, &pos, &perr) {
					goto fail15
				}
				goto ok10
			fail15:
				pos = pos13
				goto fail9
			ok10:
			}
			goto ok16
		fail9:
			pos = pos8
		ok16:
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
	return _memoize(parser, _MOQ, start, pos, perr)
fail:
	return _memoize(parser, _MOQ, start, -1, perr)
}

func _MOQNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_MOQ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MOQ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "MOQ"}
	// m O q &((questions/tone)? boundary)
	// m
	if !_node(parser, _mNode, node, &pos) {
		goto fail
	}
	// O
	if !_node(parser, _ONode, node, &pos) {
		goto fail
	}
	// q
	if !_node(parser, _qNode, node, &pos) {
		goto fail
	}
	// &((questions/tone)? boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// ((questions/tone)? boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// (questions/tone)? boundary
			// (questions/tone)?
			{
				nkids9 := len(node.Kids)
				pos10 := pos
				// (questions/tone)
				{
					nkids12 := len(node.Kids)
					pos013 := pos
					// questions/tone
					{
						pos17 := pos
						nkids15 := len(node.Kids)
						// questions
						if !_node(parser, _questionsNode, node, &pos) {
							goto fail18
						}
						goto ok14
					fail18:
						node.Kids = node.Kids[:nkids15]
						pos = pos17
						// tone
						if !_node(parser, _toneNode, node, &pos) {
							goto fail19
						}
						goto ok14
					fail19:
						node.Kids = node.Kids[:nkids15]
						pos = pos17
						goto fail11
					ok14:
					}
					sub := _sub(parser, pos013, pos, node.Kids[nkids12:])
					node.Kids = append(node.Kids[:nkids12], sub)
				}
				goto ok20
			fail11:
				node.Kids = node.Kids[:nkids9]
				pos = pos10
			ok20:
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

func _MOQFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _MOQ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "MOQ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _MOQ}
	// m O q &((questions/tone)? boundary)
	// m
	if !_fail(parser, _mFail, errPos, failure, &pos) {
		goto fail
	}
	// O
	if !_fail(parser, _OFail, errPos, failure, &pos) {
		goto fail
	}
	// q
	if !_fail(parser, _qFail, errPos, failure, &pos) {
		goto fail
	}
	// &((questions/tone)? boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// ((questions/tone)? boundary)
		// (questions/tone)? boundary
		// (questions/tone)?
		{
			pos8 := pos
			// (questions/tone)
			// questions/tone
			{
				pos13 := pos
				// questions
				if !_fail(parser, _questionsFail, errPos, failure, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				pos = pos13
				// tone
				if !_fail(parser, _toneFail, errPos, failure, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				pos = pos13
				goto fail9
			ok10:
			}
			goto ok16
		fail9:
			pos = pos8
		ok16:
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
				Want: "&((questions/tone)? boundary)",
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

func _MOQAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_MOQ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _MOQ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// m O q &((questions/tone)? boundary)
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
		// q
		if p, n := _qAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &((questions/tone)? boundary)
		{
			pos2 := pos
			// ((questions/tone)? boundary)
			// (questions/tone)? boundary
			// (questions/tone)?
			{
				pos8 := pos
				// (questions/tone)
				// questions/tone
				{
					pos13 := pos
					// questions
					if p, n := _questionsAction(parser, pos); n == nil {
						goto fail14
					} else {
						pos = p
					}
					goto ok10
				fail14:
					pos = pos13
					// tone
					if p, n := _toneAction(parser, pos); n == nil {
						goto fail15
					} else {
						pos = p
					}
					goto ok10
				fail15:
					pos = pos13
					goto fail9
				ok10:
				}
				goto ok16
			fail9:
				pos = pos8
			ok16:
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
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _NAAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _NA, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// n A &(tone? boundary)
	// n
	if !_accept(parser, _nAccepts, &pos, &perr) {
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
	return _memoize(parser, _NA, start, pos, perr)
fail:
	return _memoize(parser, _NA, start, -1, perr)
}

func _NANode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_NA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _NA}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "NA"}
	// n A &(tone? boundary)
	// n
	if !_node(parser, _nNode, node, &pos) {
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

func _NAFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _NA, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "NA",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _NA}
	// n A &(tone? boundary)
	// n
	if !_fail(parser, _nFail, errPos, failure, &pos) {
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

func _NAAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_NA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _NA}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// n A &(tone? boundary)
	{
		var node0 string
		// n
		if p, n := _nAction(parser, pos); n == nil {
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

func _POAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _PO, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (j E i/p O) &(tone? boundary)
	// (j E i/p O)
	// j E i/p O
	{
		pos4 := pos
		// j E i
		// j
		if !_accept(parser, _jAccepts, &pos, &perr) {
			goto fail5
		}
		// E
		if !_accept(parser, _EAccepts, &pos, &perr) {
			goto fail5
		}
		// i
		if !_accept(parser, _iAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// p O
		// p
		if !_accept(parser, _pAccepts, &pos, &perr) {
			goto fail7
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
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
	return _memoize(parser, _PO, start, pos, perr)
fail:
	return _memoize(parser, _PO, start, -1, perr)
}

func _PONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_PO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "PO"}
	// (j E i/p O) &(tone? boundary)
	// (j E i/p O)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// j E i/p O
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// j E i
			// j
			if !_node(parser, _jNode, node, &pos) {
				goto fail7
			}
			// E
			if !_node(parser, _ENode, node, &pos) {
				goto fail7
			}
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail7
			}
			goto ok3
		fail7:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// p O
			// p
			if !_node(parser, _pNode, node, &pos) {
				goto fail9
			}
			// O
			if !_node(parser, _ONode, node, &pos) {
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

func _POFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _PO, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "PO",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _PO}
	// (j E i/p O) &(tone? boundary)
	// (j E i/p O)
	// j E i/p O
	{
		pos4 := pos
		// j E i
		// j
		if !_fail(parser, _jFail, errPos, failure, &pos) {
			goto fail5
		}
		// E
		if !_fail(parser, _EFail, errPos, failure, &pos) {
			goto fail5
		}
		// i
		if !_fail(parser, _iFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// p O
		// p
		if !_fail(parser, _pFail, errPos, failure, &pos) {
			goto fail7
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
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

func _POAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_PO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _PO}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (j E i/p O) &(tone? boundary)
	{
		var node0 string
		// (j E i/p O)
		// j E i/p O
		{
			pos4 := pos
			var node3 string
			// j E i
			{
				var node6 string
				// j
				if p, n := _jAction(parser, pos); n == nil {
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
				// i
				if p, n := _iAction(parser, pos); n == nil {
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
			// p O
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
				// O
				if p, n := _OAction(parser, pos); n == nil {
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

func _RAAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _RA, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (r A/r U/r I/r O i/r O/r E) &(tone? boundary)
	// (r A/r U/r I/r O i/r O/r E)
	// r A/r U/r I/r O i/r O/r E
	{
		pos4 := pos
		// r A
		// r
		if !_accept(parser, _rAccepts, &pos, &perr) {
			goto fail5
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// r U
		// r
		if !_accept(parser, _rAccepts, &pos, &perr) {
			goto fail7
		}
		// U
		if !_accept(parser, _UAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// r I
		// r
		if !_accept(parser, _rAccepts, &pos, &perr) {
			goto fail9
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// r O i
		// r
		if !_accept(parser, _rAccepts, &pos, &perr) {
			goto fail11
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
			goto fail11
		}
		// i
		if !_accept(parser, _iAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// r O
		// r
		if !_accept(parser, _rAccepts, &pos, &perr) {
			goto fail13
		}
		// O
		if !_accept(parser, _OAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		// r E
		// r
		if !_accept(parser, _rAccepts, &pos, &perr) {
			goto fail15
		}
		// E
		if !_accept(parser, _EAccepts, &pos, &perr) {
			goto fail15
		}
		goto ok1
	fail15:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos18 := pos
		perr20 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos24 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail25
			}
			goto ok26
		fail25:
			pos = pos24
		ok26:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail21
		}
		goto ok17
	fail21:
		pos = pos18
		perr = _max(perr20, pos)
		goto fail
	ok17:
		pos = pos18
		perr = perr20
	}
	return _memoize(parser, _RA, start, pos, perr)
fail:
	return _memoize(parser, _RA, start, -1, perr)
}

func _RANode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_RA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _RA}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "RA"}
	// (r A/r U/r I/r O i/r O/r E) &(tone? boundary)
	// (r A/r U/r I/r O i/r O/r E)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// r A/r U/r I/r O i/r O/r E
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// r A
			// r
			if !_node(parser, _rNode, node, &pos) {
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
			// r U
			// r
			if !_node(parser, _rNode, node, &pos) {
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
			// r I
			// r
			if !_node(parser, _rNode, node, &pos) {
				goto fail11
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail11
			}
			goto ok3
		fail11:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// r O i
			// r
			if !_node(parser, _rNode, node, &pos) {
				goto fail13
			}
			// O
			if !_node(parser, _ONode, node, &pos) {
				goto fail13
			}
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail13
			}
			goto ok3
		fail13:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// r O
			// r
			if !_node(parser, _rNode, node, &pos) {
				goto fail15
			}
			// O
			if !_node(parser, _ONode, node, &pos) {
				goto fail15
			}
			goto ok3
		fail15:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// r E
			// r
			if !_node(parser, _rNode, node, &pos) {
				goto fail17
			}
			// E
			if !_node(parser, _ENode, node, &pos) {
				goto fail17
			}
			goto ok3
		fail17:
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
		pos20 := pos
		nkids21 := len(node.Kids)
		// (tone? boundary)
		{
			nkids24 := len(node.Kids)
			pos025 := pos
			// tone? boundary
			// tone?
			{
				nkids27 := len(node.Kids)
				pos28 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail29
				}
				goto ok30
			fail29:
				node.Kids = node.Kids[:nkids27]
				pos = pos28
			ok30:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail23
			}
			sub := _sub(parser, pos025, pos, node.Kids[nkids24:])
			node.Kids = append(node.Kids[:nkids24], sub)
		}
		goto ok19
	fail23:
		pos = pos20
		goto fail
	ok19:
		pos = pos20
		node.Kids = node.Kids[:nkids21]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _RAFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _RA, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "RA",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _RA}
	// (r A/r U/r I/r O i/r O/r E) &(tone? boundary)
	// (r A/r U/r I/r O i/r O/r E)
	// r A/r U/r I/r O i/r O/r E
	{
		pos4 := pos
		// r A
		// r
		if !_fail(parser, _rFail, errPos, failure, &pos) {
			goto fail5
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// r U
		// r
		if !_fail(parser, _rFail, errPos, failure, &pos) {
			goto fail7
		}
		// U
		if !_fail(parser, _UFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// r I
		// r
		if !_fail(parser, _rFail, errPos, failure, &pos) {
			goto fail9
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok1
	fail9:
		pos = pos4
		// r O i
		// r
		if !_fail(parser, _rFail, errPos, failure, &pos) {
			goto fail11
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
			goto fail11
		}
		// i
		if !_fail(parser, _iFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// r O
		// r
		if !_fail(parser, _rFail, errPos, failure, &pos) {
			goto fail13
		}
		// O
		if !_fail(parser, _OFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		// r E
		// r
		if !_fail(parser, _rFail, errPos, failure, &pos) {
			goto fail15
		}
		// E
		if !_fail(parser, _EFail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok1
	fail15:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos18 := pos
		nkids19 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos24 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail25
			}
			goto ok26
		fail25:
			pos = pos24
		ok26:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail21
		}
		goto ok17
	fail21:
		pos = pos18
		failure.Kids = failure.Kids[:nkids19]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok17:
		pos = pos18
		failure.Kids = failure.Kids[:nkids19]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _RAAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_RA]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _RA}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (r A/r U/r I/r O i/r O/r E) &(tone? boundary)
	{
		var node0 string
		// (r A/r U/r I/r O i/r O/r E)
		// r A/r U/r I/r O i/r O/r E
		{
			pos4 := pos
			var node3 string
			// r A
			{
				var node6 string
				// r
				if p, n := _rAction(parser, pos); n == nil {
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
			// r U
			{
				var node8 string
				// r
				if p, n := _rAction(parser, pos); n == nil {
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
			// r I
			{
				var node10 string
				// r
				if p, n := _rAction(parser, pos); n == nil {
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
			}
			goto ok1
		fail9:
			node0 = node3
			pos = pos4
			// r O i
			{
				var node12 string
				// r
				if p, n := _rAction(parser, pos); n == nil {
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
				// i
				if p, n := _iAction(parser, pos); n == nil {
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
			// r O
			{
				var node14 string
				// r
				if p, n := _rAction(parser, pos); n == nil {
					goto fail13
				} else {
					node14 = *n
					pos = p
				}
				node0, node14 = node0+node14, ""
				// O
				if p, n := _OAction(parser, pos); n == nil {
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
			// r E
			{
				var node16 string
				// r
				if p, n := _rAction(parser, pos); n == nil {
					goto fail15
				} else {
					node16 = *n
					pos = p
				}
				node0, node16 = node0+node16, ""
				// E
				if p, n := _EAction(parser, pos); n == nil {
					goto fail15
				} else {
					node16 = *n
					pos = p
				}
				node0, node16 = node0+node16, ""
			}
			goto ok1
		fail15:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos18 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos24 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail25
				} else {
					pos = p
				}
				goto ok26
			fail25:
				pos = pos24
			ok26:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail21
			} else {
				pos = p
			}
			goto ok17
		fail21:
			pos = pos18
			goto fail
		ok17:
			pos = pos18
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _TUAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _TU, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// (t A/t U/s I a/s A/h I/j A) &(tone? boundary)
	// (t A/t U/s I a/s A/h I/j A)
	// t A/t U/s I a/s A/h I/j A
	{
		pos4 := pos
		// t A
		// t
		if !_accept(parser, _tAccepts, &pos, &perr) {
			goto fail5
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// t U
		// t
		if !_accept(parser, _tAccepts, &pos, &perr) {
			goto fail7
		}
		// U
		if !_accept(parser, _UAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// s I a
		// s
		if !_accept(parser, _sAccepts, &pos, &perr) {
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
		// s A
		// s
		if !_accept(parser, _sAccepts, &pos, &perr) {
			goto fail11
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// h I
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail13
		}
		// I
		if !_accept(parser, _IAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		// j A
		// j
		if !_accept(parser, _jAccepts, &pos, &perr) {
			goto fail15
		}
		// A
		if !_accept(parser, _AAccepts, &pos, &perr) {
			goto fail15
		}
		goto ok1
	fail15:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos18 := pos
		perr20 := perr
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos24 := pos
			// tone
			if !_accept(parser, _toneAccepts, &pos, &perr) {
				goto fail25
			}
			goto ok26
		fail25:
			pos = pos24
		ok26:
		}
		// boundary
		if !_accept(parser, _boundaryAccepts, &pos, &perr) {
			goto fail21
		}
		goto ok17
	fail21:
		pos = pos18
		perr = _max(perr20, pos)
		goto fail
	ok17:
		pos = pos18
		perr = perr20
	}
	return _memoize(parser, _TU, start, pos, perr)
fail:
	return _memoize(parser, _TU, start, -1, perr)
}

func _TUNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_TU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _TU}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "TU"}
	// (t A/t U/s I a/s A/h I/j A) &(tone? boundary)
	// (t A/t U/s I a/s A/h I/j A)
	{
		nkids1 := len(node.Kids)
		pos02 := pos
		// t A/t U/s I a/s A/h I/j A
		{
			pos6 := pos
			nkids4 := len(node.Kids)
			// t A
			// t
			if !_node(parser, _tNode, node, &pos) {
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
			// t U
			// t
			if !_node(parser, _tNode, node, &pos) {
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
			// s I a
			// s
			if !_node(parser, _sNode, node, &pos) {
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
			// s A
			// s
			if !_node(parser, _sNode, node, &pos) {
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
			// h I
			// h
			if !_node(parser, _hNode, node, &pos) {
				goto fail15
			}
			// I
			if !_node(parser, _INode, node, &pos) {
				goto fail15
			}
			goto ok3
		fail15:
			node.Kids = node.Kids[:nkids4]
			pos = pos6
			// j A
			// j
			if !_node(parser, _jNode, node, &pos) {
				goto fail17
			}
			// A
			if !_node(parser, _ANode, node, &pos) {
				goto fail17
			}
			goto ok3
		fail17:
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
		pos20 := pos
		nkids21 := len(node.Kids)
		// (tone? boundary)
		{
			nkids24 := len(node.Kids)
			pos025 := pos
			// tone? boundary
			// tone?
			{
				nkids27 := len(node.Kids)
				pos28 := pos
				// tone
				if !_node(parser, _toneNode, node, &pos) {
					goto fail29
				}
				goto ok30
			fail29:
				node.Kids = node.Kids[:nkids27]
				pos = pos28
			ok30:
			}
			// boundary
			if !_node(parser, _boundaryNode, node, &pos) {
				goto fail23
			}
			sub := _sub(parser, pos025, pos, node.Kids[nkids24:])
			node.Kids = append(node.Kids[:nkids24], sub)
		}
		goto ok19
	fail23:
		pos = pos20
		goto fail
	ok19:
		pos = pos20
		node.Kids = node.Kids[:nkids21]
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _TUFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _TU, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "TU",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _TU}
	// (t A/t U/s I a/s A/h I/j A) &(tone? boundary)
	// (t A/t U/s I a/s A/h I/j A)
	// t A/t U/s I a/s A/h I/j A
	{
		pos4 := pos
		// t A
		// t
		if !_fail(parser, _tFail, errPos, failure, &pos) {
			goto fail5
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok1
	fail5:
		pos = pos4
		// t U
		// t
		if !_fail(parser, _tFail, errPos, failure, &pos) {
			goto fail7
		}
		// U
		if !_fail(parser, _UFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok1
	fail7:
		pos = pos4
		// s I a
		// s
		if !_fail(parser, _sFail, errPos, failure, &pos) {
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
		// s A
		// s
		if !_fail(parser, _sFail, errPos, failure, &pos) {
			goto fail11
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok1
	fail11:
		pos = pos4
		// h I
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail13
		}
		// I
		if !_fail(parser, _IFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok1
	fail13:
		pos = pos4
		// j A
		// j
		if !_fail(parser, _jFail, errPos, failure, &pos) {
			goto fail15
		}
		// A
		if !_fail(parser, _AFail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok1
	fail15:
		pos = pos4
		goto fail
	ok1:
	}
	// &(tone? boundary)
	{
		pos18 := pos
		nkids19 := len(failure.Kids)
		// (tone? boundary)
		// tone? boundary
		// tone?
		{
			pos24 := pos
			// tone
			if !_fail(parser, _toneFail, errPos, failure, &pos) {
				goto fail25
			}
			goto ok26
		fail25:
			pos = pos24
		ok26:
		}
		// boundary
		if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
			goto fail21
		}
		goto ok17
	fail21:
		pos = pos18
		failure.Kids = failure.Kids[:nkids19]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(tone? boundary)",
			})
		}
		goto fail
	ok17:
		pos = pos18
		failure.Kids = failure.Kids[:nkids19]
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _TUAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_TU]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _TU}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// (t A/t U/s I a/s A/h I/j A) &(tone? boundary)
	{
		var node0 string
		// (t A/t U/s I a/s A/h I/j A)
		// t A/t U/s I a/s A/h I/j A
		{
			pos4 := pos
			var node3 string
			// t A
			{
				var node6 string
				// t
				if p, n := _tAction(parser, pos); n == nil {
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
			// t U
			{
				var node8 string
				// t
				if p, n := _tAction(parser, pos); n == nil {
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
			// s I a
			{
				var node10 string
				// s
				if p, n := _sAction(parser, pos); n == nil {
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
			// s A
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
			// h I
			{
				var node14 string
				// h
				if p, n := _hAction(parser, pos); n == nil {
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
			}
			goto ok1
		fail13:
			node0 = node3
			pos = pos4
			// j A
			{
				var node16 string
				// j
				if p, n := _jAction(parser, pos); n == nil {
					goto fail15
				} else {
					node16 = *n
					pos = p
				}
				node0, node16 = node0+node16, ""
				// A
				if p, n := _AAction(parser, pos); n == nil {
					goto fail15
				} else {
					node16 = *n
					pos = p
				}
				node0, node16 = node0+node16, ""
			}
			goto ok1
		fail15:
			node0 = node3
			pos = pos4
			goto fail
		ok1:
		}
		node, node0 = node+node0, ""
		// &(tone? boundary)
		{
			pos18 := pos
			// (tone? boundary)
			// tone? boundary
			// tone?
			{
				pos24 := pos
				// tone
				if p, n := _toneAction(parser, pos); n == nil {
					goto fail25
				} else {
					pos = p
				}
				goto ok26
			fail25:
				pos = pos24
			ok26:
			}
			// boundary
			if p, n := _boundaryAction(parser, pos); n == nil {
				goto fail21
			} else {
				pos = p
			}
			goto ok17
		fail21:
			pos = pos18
			goto fail
		ok17:
			pos = pos18
			node0 = ""
		}
		node, node0 = node+node0, ""
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _TOAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _TO, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// t O &boundary
	// t
	if !_accept(parser, _tAccepts, &pos, &perr) {
		goto fail
	}
	// O
	if !_accept(parser, _OAccepts, &pos, &perr) {
		goto fail
	}
	// &boundary
	{
		pos2 := pos
		perr4 := perr
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
	return _memoize(parser, _TO, start, pos, perr)
fail:
	return _memoize(parser, _TO, start, -1, perr)
}

func _TONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_TO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _TO}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "TO"}
	// t O &boundary
	// t
	if !_node(parser, _tNode, node, &pos) {
		goto fail
	}
	// O
	if !_node(parser, _ONode, node, &pos) {
		goto fail
	}
	// &boundary
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// boundary
		if !_node(parser, _boundaryNode, node, &pos) {
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
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _TOFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _TO, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "TO",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _TO}
	// t O &boundary
	// t
	if !_fail(parser, _tFail, errPos, failure, &pos) {
		goto fail
	}
	// O
	if !_fail(parser, _OFail, errPos, failure, &pos) {
		goto fail
	}
	// &boundary
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
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
				Want: "&boundary",
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

func _TOAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_TO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _TO}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// t O &boundary
	{
		var node0 string
		// t
		if p, n := _tAction(parser, pos); n == nil {
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
		// &boundary
		{
			pos2 := pos
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

func _TEOAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _TEO, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// t E o &(tone? boundary)
	// t
	if !_accept(parser, _tAccepts, &pos, &perr) {
		goto fail
	}
	// E
	if !_accept(parser, _EAccepts, &pos, &perr) {
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
	return _memoize(parser, _TEO, start, pos, perr)
fail:
	return _memoize(parser, _TEO, start, -1, perr)
}

func _TEONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_TEO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _TEO}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "TEO"}
	// t E o &(tone? boundary)
	// t
	if !_node(parser, _tNode, node, &pos) {
		goto fail
	}
	// E
	if !_node(parser, _ENode, node, &pos) {
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

func _TEOFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _TEO, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "TEO",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _TEO}
	// t E o &(tone? boundary)
	// t
	if !_fail(parser, _tFail, errPos, failure, &pos) {
		goto fail
	}
	// E
	if !_fail(parser, _EFail, errPos, failure, &pos) {
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

func _TEOAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_TEO]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _TEO}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// t E o &(tone? boundary)
	{
		var node0 string
		// t
		if p, n := _tAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// E
		if p, n := _EAction(parser, pos); n == nil {
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

func _neutral_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _neutral_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// initial neutral_desinence &boundary
	// initial
	if !_accept(parser, _initialAccepts, &pos, &perr) {
		goto fail
	}
	// neutral_desinence
	if !_accept(parser, _neutral_desinenceAccepts, &pos, &perr) {
		goto fail
	}
	// &boundary
	{
		pos2 := pos
		perr4 := perr
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
	return _memoize(parser, _neutral_syllable, start, pos, perr)
fail:
	return _memoize(parser, _neutral_syllable, start, -1, perr)
}

func _neutral_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_neutral_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _neutral_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "neutral_syllable"}
	// initial neutral_desinence &boundary
	// initial
	if !_node(parser, _initialNode, node, &pos) {
		goto fail
	}
	// neutral_desinence
	if !_node(parser, _neutral_desinenceNode, node, &pos) {
		goto fail
	}
	// &boundary
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// boundary
		if !_node(parser, _boundaryNode, node, &pos) {
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
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _neutral_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _neutral_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "neutral_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _neutral_syllable}
	// initial neutral_desinence &boundary
	// initial
	if !_fail(parser, _initialFail, errPos, failure, &pos) {
		goto fail
	}
	// neutral_desinence
	if !_fail(parser, _neutral_desinenceFail, errPos, failure, &pos) {
		goto fail
	}
	// &boundary
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
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
				Want: "&boundary",
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

func _neutral_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_neutral_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _neutral_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// initial neutral_desinence &boundary
	{
		var node0 string
		// initial
		if p, n := _initialAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// neutral_desinence
		if p, n := _neutral_desinenceAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &boundary
		{
			pos2 := pos
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

func _moq_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _moq_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// initial neutral_desinence &(questions boundary/boundary)
	// initial
	if !_accept(parser, _initialAccepts, &pos, &perr) {
		goto fail
	}
	// neutral_desinence
	if !_accept(parser, _neutral_desinenceAccepts, &pos, &perr) {
		goto fail
	}
	// &(questions boundary/boundary)
	{
		pos2 := pos
		perr4 := perr
		// (questions boundary/boundary)
		// questions boundary/boundary
		{
			pos9 := pos
			// questions boundary
			// questions
			if !_accept(parser, _questionsAccepts, &pos, &perr) {
				goto fail10
			}
			// boundary
			if !_accept(parser, _boundaryAccepts, &pos, &perr) {
				goto fail10
			}
			goto ok6
		fail10:
			pos = pos9
			// boundary
			if !_accept(parser, _boundaryAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok6
		fail12:
			pos = pos9
			goto fail5
		ok6:
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
	return _memoize(parser, _moq_syllable, start, pos, perr)
fail:
	return _memoize(parser, _moq_syllable, start, -1, perr)
}

func _moq_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_moq_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _moq_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "moq_syllable"}
	// initial neutral_desinence &(questions boundary/boundary)
	// initial
	if !_node(parser, _initialNode, node, &pos) {
		goto fail
	}
	// neutral_desinence
	if !_node(parser, _neutral_desinenceNode, node, &pos) {
		goto fail
	}
	// &(questions boundary/boundary)
	{
		pos2 := pos
		nkids3 := len(node.Kids)
		// (questions boundary/boundary)
		{
			nkids6 := len(node.Kids)
			pos07 := pos
			// questions boundary/boundary
			{
				pos11 := pos
				nkids9 := len(node.Kids)
				// questions boundary
				// questions
				if !_node(parser, _questionsNode, node, &pos) {
					goto fail12
				}
				// boundary
				if !_node(parser, _boundaryNode, node, &pos) {
					goto fail12
				}
				goto ok8
			fail12:
				node.Kids = node.Kids[:nkids9]
				pos = pos11
				// boundary
				if !_node(parser, _boundaryNode, node, &pos) {
					goto fail14
				}
				goto ok8
			fail14:
				node.Kids = node.Kids[:nkids9]
				pos = pos11
				goto fail5
			ok8:
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

func _moq_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _moq_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "moq_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _moq_syllable}
	// initial neutral_desinence &(questions boundary/boundary)
	// initial
	if !_fail(parser, _initialFail, errPos, failure, &pos) {
		goto fail
	}
	// neutral_desinence
	if !_fail(parser, _neutral_desinenceFail, errPos, failure, &pos) {
		goto fail
	}
	// &(questions boundary/boundary)
	{
		pos2 := pos
		nkids3 := len(failure.Kids)
		// (questions boundary/boundary)
		// questions boundary/boundary
		{
			pos9 := pos
			// questions boundary
			// questions
			if !_fail(parser, _questionsFail, errPos, failure, &pos) {
				goto fail10
			}
			// boundary
			if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
				goto fail10
			}
			goto ok6
		fail10:
			pos = pos9
			// boundary
			if !_fail(parser, _boundaryFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok6
		fail12:
			pos = pos9
			goto fail5
		ok6:
		}
		goto ok1
	fail5:
		pos = pos2
		failure.Kids = failure.Kids[:nkids3]
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "&(questions boundary/boundary)",
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

func _moq_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_moq_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _moq_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// initial neutral_desinence &(questions boundary/boundary)
	{
		var node0 string
		// initial
		if p, n := _initialAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// neutral_desinence
		if p, n := _neutral_desinenceAction(parser, pos); n == nil {
			goto fail
		} else {
			node0 = *n
			pos = p
		}
		node, node0 = node+node0, ""
		// &(questions boundary/boundary)
		{
			pos2 := pos
			// (questions boundary/boundary)
			// questions boundary/boundary
			{
				pos9 := pos
				// questions boundary
				// questions
				if p, n := _questionsAction(parser, pos); n == nil {
					goto fail10
				} else {
					pos = p
				}
				// boundary
				if p, n := _boundaryAction(parser, pos); n == nil {
					goto fail10
				} else {
					pos = p
				}
				goto ok6
			fail10:
				pos = pos9
				// boundary
				if p, n := _boundaryAction(parser, pos); n == nil {
					goto fail12
				} else {
					pos = p
				}
				goto ok6
			fail12:
				pos = pos9
				goto fail5
			ok6:
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

func _compound_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _compound_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// syllable<compound_desinence, compound_tone>
	if !_accept(parser, _syllable__compound_desinence__compound_toneAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _compound_syllable, start, pos, perr)
fail:
	return _memoize(parser, _compound_syllable, start, -1, perr)
}

func _compound_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_compound_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _compound_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "compound_syllable"}
	// syllable<compound_desinence, compound_tone>
	if !_node(parser, _syllable__compound_desinence__compound_toneNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _compound_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _compound_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "compound_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _compound_syllable}
	// syllable<compound_desinence, compound_tone>
	if !_fail(parser, _syllable__compound_desinence__compound_toneFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _compound_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_compound_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _compound_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// syllable<compound_desinence, compound_tone>
	if p, n := _syllable__compound_desinence__compound_toneAction(parser, pos); n == nil {
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

func _arg_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// syllable<arg_desinence, arg_tone>
	if !_accept(parser, _syllable__arg_desinence__arg_toneAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _arg_syllable, start, pos, perr)
fail:
	return _memoize(parser, _arg_syllable, start, -1, perr)
}

func _arg_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_syllable"}
	// syllable<arg_desinence, arg_tone>
	if !_node(parser, _syllable__arg_desinence__arg_toneNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_syllable}
	// syllable<arg_desinence, arg_tone>
	if !_fail(parser, _syllable__arg_desinence__arg_toneFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// syllable<arg_desinence, arg_tone>
	if p, n := _syllable__arg_desinence__arg_toneAction(parser, pos); n == nil {
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

func _relative_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// syllable<relative_desinence, relative_tone>
	if !_accept(parser, _syllable__relative_desinence__relative_toneAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _relative_syllable, start, pos, perr)
fail:
	return _memoize(parser, _relative_syllable, start, -1, perr)
}

func _relative_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_syllable"}
	// syllable<relative_desinence, relative_tone>
	if !_node(parser, _syllable__relative_desinence__relative_toneNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_syllable}
	// syllable<relative_desinence, relative_tone>
	if !_fail(parser, _syllable__relative_desinence__relative_toneFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// syllable<relative_desinence, relative_tone>
	if p, n := _syllable__relative_desinence__relative_toneAction(parser, pos); n == nil {
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

func _verb_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _verb_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// syllable<verb_desinence, verb_tone>
	if !_accept(parser, _syllable__verb_desinence__verb_toneAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _verb_syllable, start, pos, perr)
fail:
	return _memoize(parser, _verb_syllable, start, -1, perr)
}

func _verb_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_verb_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _verb_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "verb_syllable"}
	// syllable<verb_desinence, verb_tone>
	if !_node(parser, _syllable__verb_desinence__verb_toneNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _verb_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _verb_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "verb_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _verb_syllable}
	// syllable<verb_desinence, verb_tone>
	if !_fail(parser, _syllable__verb_desinence__verb_toneFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _verb_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_verb_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _verb_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// syllable<verb_desinence, verb_tone>
	if p, n := _syllable__verb_desinence__verb_toneAction(parser, pos); n == nil {
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

func _content_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// syllable<content_desinence, content_tone>
	if !_accept(parser, _syllable__content_desinence__content_toneAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _content_syllable, start, pos, perr)
fail:
	return _memoize(parser, _content_syllable, start, -1, perr)
}

func _content_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_syllable"}
	// syllable<content_desinence, content_tone>
	if !_node(parser, _syllable__content_desinence__content_toneNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_syllable}
	// syllable<content_desinence, content_tone>
	if !_fail(parser, _syllable__content_desinence__content_toneFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// syllable<content_desinence, content_tone>
	if p, n := _syllable__content_desinence__content_toneAction(parser, pos); n == nil {
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

func _preposition_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// syllable<preposition_desinence, preposition_tone>
	if !_accept(parser, _syllable__preposition_desinence__preposition_toneAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _preposition_syllable, start, pos, perr)
fail:
	return _memoize(parser, _preposition_syllable, start, -1, perr)
}

func _preposition_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition_syllable"}
	// syllable<preposition_desinence, preposition_tone>
	if !_node(parser, _syllable__preposition_desinence__preposition_toneNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _preposition_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_syllable}
	// syllable<preposition_desinence, preposition_tone>
	if !_fail(parser, _syllable__preposition_desinence__preposition_toneFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// syllable<preposition_desinence, preposition_tone>
	if p, n := _syllable__preposition_desinence__preposition_toneAction(parser, pos); n == nil {
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

func _adverb_syllableAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb_syllable, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// syllable<adverb_desinence, adverb_tone>
	if !_accept(parser, _syllable__adverb_desinence__adverb_toneAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _adverb_syllable, start, pos, perr)
fail:
	return _memoize(parser, _adverb_syllable, start, -1, perr)
}

func _adverb_syllableNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_syllable}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb_syllable"}
	// syllable<adverb_desinence, adverb_tone>
	if !_node(parser, _syllable__adverb_desinence__adverb_toneNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_syllableFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb_syllable, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_syllable",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_syllable}
	// syllable<adverb_desinence, adverb_tone>
	if !_fail(parser, _syllable__adverb_desinence__adverb_toneFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_syllableAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb_syllable]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_syllable}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// syllable<adverb_desinence, adverb_tone>
	if p, n := _syllable__adverb_desinence__adverb_toneAction(parser, pos); n == nil {
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

func _boundaryAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _boundary, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// action
	// (initial/space/EOF)
	// initial/space/EOF
	{
		pos3 := pos
		// initial
		if !_accept(parser, _initialAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// space
		if !_accept(parser, _spaceAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// EOF
		if !_accept(parser, _EOFAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _boundary, start, pos, perr)
fail:
	return _memoize(parser, _boundary, start, -1, perr)
}

func _boundaryNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_boundary]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _boundary}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "boundary"}
	// action
	// (initial/space/EOF)
	{
		nkids0 := len(node.Kids)
		pos01 := pos
		// initial/space/EOF
		{
			pos5 := pos
			nkids3 := len(node.Kids)
			// initial
			if !_node(parser, _initialNode, node, &pos) {
				goto fail6
			}
			goto ok2
		fail6:
			node.Kids = node.Kids[:nkids3]
			pos = pos5
			// space
			if !_node(parser, _spaceNode, node, &pos) {
				goto fail7
			}
			goto ok2
		fail7:
			node.Kids = node.Kids[:nkids3]
			pos = pos5
			// EOF
			if !_node(parser, _EOFNode, node, &pos) {
				goto fail8
			}
			goto ok2
		fail8:
			node.Kids = node.Kids[:nkids3]
			pos = pos5
			goto fail
		ok2:
		}
		sub := _sub(parser, pos01, pos, node.Kids[nkids0:])
		node.Kids = append(node.Kids[:nkids0], sub)
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _boundaryFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _boundary, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "boundary",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _boundary}
	// action
	// (initial/space/EOF)
	// initial/space/EOF
	{
		pos3 := pos
		// initial
		if !_fail(parser, _initialFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// space
		if !_fail(parser, _spaceFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// EOF
		if !_fail(parser, _EOFFail, errPos, failure, &pos) {
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

func _boundaryAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_boundary]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _boundary}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// action
	{
		start0 := pos
		// (initial/space/EOF)
		// initial/space/EOF
		{
			pos4 := pos
			// initial
			if p, n := _initialAction(parser, pos); n == nil {
				goto fail5
			} else {
				pos = p
			}
			goto ok1
		fail5:
			pos = pos4
			// space
			if p, n := _spaceAction(parser, pos); n == nil {
				goto fail6
			} else {
				pos = p
			}
			goto ok1
		fail6:
			pos = pos4
			// EOF
			if p, n := _EOFAction(parser, pos); n == nil {
				goto fail7
			} else {
				pos = p
			}
			goto ok1
		fail7:
			pos = pos4
			goto fail
		ok1:
		}
		node = func(
			start, end int) string {
			return ""
		}(
			start0, pos)
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _initialAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _initial, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// b/c h/c/d/f/g/h/j/k/l/m/n/p/r/s h/s/t
	{
		pos3 := pos
		// b
		if !_accept(parser, _bAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// c h
		// c
		if !_accept(parser, _cAccepts, &pos, &perr) {
			goto fail5
		}
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// c
		if !_accept(parser, _cAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// d
		if !_accept(parser, _dAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// f
		if !_accept(parser, _fAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// g
		if !_accept(parser, _gAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		// j
		if !_accept(parser, _jAccepts, &pos, &perr) {
			goto fail12
		}
		goto ok0
	fail12:
		pos = pos3
		// k
		if !_accept(parser, _kAccepts, &pos, &perr) {
			goto fail13
		}
		goto ok0
	fail13:
		pos = pos3
		// l
		if !_accept(parser, _lAccepts, &pos, &perr) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		// m
		if !_accept(parser, _mAccepts, &pos, &perr) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		// n
		if !_accept(parser, _nAccepts, &pos, &perr) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		// p
		if !_accept(parser, _pAccepts, &pos, &perr) {
			goto fail17
		}
		goto ok0
	fail17:
		pos = pos3
		// r
		if !_accept(parser, _rAccepts, &pos, &perr) {
			goto fail18
		}
		goto ok0
	fail18:
		pos = pos3
		// s h
		// s
		if !_accept(parser, _sAccepts, &pos, &perr) {
			goto fail19
		}
		// h
		if !_accept(parser, _hAccepts, &pos, &perr) {
			goto fail19
		}
		goto ok0
	fail19:
		pos = pos3
		// s
		if !_accept(parser, _sAccepts, &pos, &perr) {
			goto fail21
		}
		goto ok0
	fail21:
		pos = pos3
		// t
		if !_accept(parser, _tAccepts, &pos, &perr) {
			goto fail22
		}
		goto ok0
	fail22:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _initial, start, pos, perr)
fail:
	return _memoize(parser, _initial, start, -1, perr)
}

func _initialNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_initial]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _initial}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "initial"}
	// b/c h/c/d/f/g/h/j/k/l/m/n/p/r/s h/s/t
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// b
		if !_node(parser, _bNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// c h
		// c
		if !_node(parser, _cNode, node, &pos) {
			goto fail5
		}
		// h
		if !_node(parser, _hNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// c
		if !_node(parser, _cNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// d
		if !_node(parser, _dNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// f
		if !_node(parser, _fNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// g
		if !_node(parser, _gNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// h
		if !_node(parser, _hNode, node, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// j
		if !_node(parser, _jNode, node, &pos) {
			goto fail12
		}
		goto ok0
	fail12:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// k
		if !_node(parser, _kNode, node, &pos) {
			goto fail13
		}
		goto ok0
	fail13:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// l
		if !_node(parser, _lNode, node, &pos) {
			goto fail14
		}
		goto ok0
	fail14:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// m
		if !_node(parser, _mNode, node, &pos) {
			goto fail15
		}
		goto ok0
	fail15:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// n
		if !_node(parser, _nNode, node, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// p
		if !_node(parser, _pNode, node, &pos) {
			goto fail17
		}
		goto ok0
	fail17:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// r
		if !_node(parser, _rNode, node, &pos) {
			goto fail18
		}
		goto ok0
	fail18:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// s h
		// s
		if !_node(parser, _sNode, node, &pos) {
			goto fail19
		}
		// h
		if !_node(parser, _hNode, node, &pos) {
			goto fail19
		}
		goto ok0
	fail19:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// s
		if !_node(parser, _sNode, node, &pos) {
			goto fail21
		}
		goto ok0
	fail21:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// t
		if !_node(parser, _tNode, node, &pos) {
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

func _initialFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _initial, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "initial",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _initial}
	// b/c h/c/d/f/g/h/j/k/l/m/n/p/r/s h/s/t
	{
		pos3 := pos
		// b
		if !_fail(parser, _bFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// c h
		// c
		if !_fail(parser, _cFail, errPos, failure, &pos) {
			goto fail5
		}
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// c
		if !_fail(parser, _cFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// d
		if !_fail(parser, _dFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// f
		if !_fail(parser, _fFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// g
		if !_fail(parser, _gFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		// j
		if !_fail(parser, _jFail, errPos, failure, &pos) {
			goto fail12
		}
		goto ok0
	fail12:
		pos = pos3
		// k
		if !_fail(parser, _kFail, errPos, failure, &pos) {
			goto fail13
		}
		goto ok0
	fail13:
		pos = pos3
		// l
		if !_fail(parser, _lFail, errPos, failure, &pos) {
			goto fail14
		}
		goto ok0
	fail14:
		pos = pos3
		// m
		if !_fail(parser, _mFail, errPos, failure, &pos) {
			goto fail15
		}
		goto ok0
	fail15:
		pos = pos3
		// n
		if !_fail(parser, _nFail, errPos, failure, &pos) {
			goto fail16
		}
		goto ok0
	fail16:
		pos = pos3
		// p
		if !_fail(parser, _pFail, errPos, failure, &pos) {
			goto fail17
		}
		goto ok0
	fail17:
		pos = pos3
		// r
		if !_fail(parser, _rFail, errPos, failure, &pos) {
			goto fail18
		}
		goto ok0
	fail18:
		pos = pos3
		// s h
		// s
		if !_fail(parser, _sFail, errPos, failure, &pos) {
			goto fail19
		}
		// h
		if !_fail(parser, _hFail, errPos, failure, &pos) {
			goto fail19
		}
		goto ok0
	fail19:
		pos = pos3
		// s
		if !_fail(parser, _sFail, errPos, failure, &pos) {
			goto fail21
		}
		goto ok0
	fail21:
		pos = pos3
		// t
		if !_fail(parser, _tFail, errPos, failure, &pos) {
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

func _initialAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_initial]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _initial}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// b/c h/c/d/f/g/h/j/k/l/m/n/p/r/s h/s/t
	{
		pos3 := pos
		var node2 string
		// b
		if p, n := _bAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// c h
		{
			var node6 string
			// c
			if p, n := _cAction(parser, pos); n == nil {
				goto fail5
			} else {
				node6 = *n
				pos = p
			}
			node, node6 = node+node6, ""
			// h
			if p, n := _hAction(parser, pos); n == nil {
				goto fail5
			} else {
				node6 = *n
				pos = p
			}
			node, node6 = node+node6, ""
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// c
		if p, n := _cAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// d
		if p, n := _dAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// f
		if p, n := _fAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// g
		if p, n := _gAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// h
		if p, n := _hAction(parser, pos); n == nil {
			goto fail11
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail11:
		node = node2
		pos = pos3
		// j
		if p, n := _jAction(parser, pos); n == nil {
			goto fail12
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail12:
		node = node2
		pos = pos3
		// k
		if p, n := _kAction(parser, pos); n == nil {
			goto fail13
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail13:
		node = node2
		pos = pos3
		// l
		if p, n := _lAction(parser, pos); n == nil {
			goto fail14
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail14:
		node = node2
		pos = pos3
		// m
		if p, n := _mAction(parser, pos); n == nil {
			goto fail15
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail15:
		node = node2
		pos = pos3
		// n
		if p, n := _nAction(parser, pos); n == nil {
			goto fail16
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail16:
		node = node2
		pos = pos3
		// p
		if p, n := _pAction(parser, pos); n == nil {
			goto fail17
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail17:
		node = node2
		pos = pos3
		// r
		if p, n := _rAction(parser, pos); n == nil {
			goto fail18
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail18:
		node = node2
		pos = pos3
		// s h
		{
			var node20 string
			// s
			if p, n := _sAction(parser, pos); n == nil {
				goto fail19
			} else {
				node20 = *n
				pos = p
			}
			node, node20 = node+node20, ""
			// h
			if p, n := _hAction(parser, pos); n == nil {
				goto fail19
			} else {
				node20 = *n
				pos = p
			}
			node, node20 = node+node20, ""
		}
		goto ok0
	fail19:
		node = node2
		pos = pos3
		// s
		if p, n := _sAction(parser, pos); n == nil {
			goto fail21
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail21:
		node = node2
		pos = pos3
		// t
		if p, n := _tAction(parser, pos); n == nil {
			goto fail22
		} else {
			node = *n
			pos = p
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

func _neutral_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _neutral_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<a, u, i, o, e>
	if !_accept(parser, _desinence__a__u__i__o__eAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _neutral_desinence, start, pos, perr)
fail:
	return _memoize(parser, _neutral_desinence, start, -1, perr)
}

func _neutral_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_neutral_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _neutral_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "neutral_desinence"}
	// desinence<a, u, i, o, e>
	if !_node(parser, _desinence__a__u__i__o__eNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _neutral_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _neutral_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "neutral_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _neutral_desinence}
	// desinence<a, u, i, o, e>
	if !_fail(parser, _desinence__a__u__i__o__eFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _neutral_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_neutral_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _neutral_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<a, u, i, o, e>
	if p, n := _desinence__a__u__i__o__eAction(parser, pos); n == nil {
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

func _compound_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _compound_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<ā, ū, ī, ō, ē>
	if !_accept(parser, _desinence__ā__ū__ī__ō__ēAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _compound_desinence, start, pos, perr)
fail:
	return _memoize(parser, _compound_desinence, start, -1, perr)
}

func _compound_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_compound_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _compound_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "compound_desinence"}
	// desinence<ā, ū, ī, ō, ē>
	if !_node(parser, _desinence__ā__ū__ī__ō__ēNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _compound_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _compound_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "compound_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _compound_desinence}
	// desinence<ā, ū, ī, ō, ē>
	if !_fail(parser, _desinence__ā__ū__ī__ō__ēFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _compound_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_compound_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _compound_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<ā, ū, ī, ō, ē>
	if p, n := _desinence__ā__ū__ī__ō__ēAction(parser, pos); n == nil {
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

func _arg_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<á, ú, í, ó, é>
	if !_accept(parser, _desinence__á__ú__í__ó__éAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _arg_desinence, start, pos, perr)
fail:
	return _memoize(parser, _arg_desinence, start, -1, perr)
}

func _arg_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_desinence"}
	// desinence<á, ú, í, ó, é>
	if !_node(parser, _desinence__á__ú__í__ó__éNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_desinence}
	// desinence<á, ú, í, ó, é>
	if !_fail(parser, _desinence__á__ú__í__ó__éFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<á, ú, í, ó, é>
	if p, n := _desinence__á__ú__í__ó__éAction(parser, pos); n == nil {
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

func _relative_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<ǎ, ǔ, ǐ, ǒ, ě>
	if !_accept(parser, _desinence__ǎ__ǔ__ǐ__ǒ__ěAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _relative_desinence, start, pos, perr)
fail:
	return _memoize(parser, _relative_desinence, start, -1, perr)
}

func _relative_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_desinence"}
	// desinence<ǎ, ǔ, ǐ, ǒ, ě>
	if !_node(parser, _desinence__ǎ__ǔ__ǐ__ǒ__ěNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_desinence}
	// desinence<ǎ, ǔ, ǐ, ǒ, ě>
	if !_fail(parser, _desinence__ǎ__ǔ__ǐ__ǒ__ěFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<ǎ, ǔ, ǐ, ǒ, ě>
	if p, n := _desinence__ǎ__ǔ__ǐ__ǒ__ěAction(parser, pos); n == nil {
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

func _verb_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _verb_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<ả, ủ, ỉ, ỏ, ẻ>
	if !_accept(parser, _desinence__ả__ủ__ỉ__ỏ__ẻAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _verb_desinence, start, pos, perr)
fail:
	return _memoize(parser, _verb_desinence, start, -1, perr)
}

func _verb_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_verb_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _verb_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "verb_desinence"}
	// desinence<ả, ủ, ỉ, ỏ, ẻ>
	if !_node(parser, _desinence__ả__ủ__ỉ__ỏ__ẻNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _verb_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _verb_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "verb_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _verb_desinence}
	// desinence<ả, ủ, ỉ, ỏ, ẻ>
	if !_fail(parser, _desinence__ả__ủ__ỉ__ỏ__ẻFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _verb_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_verb_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _verb_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<ả, ủ, ỉ, ỏ, ẻ>
	if p, n := _desinence__ả__ủ__ỉ__ỏ__ẻAction(parser, pos); n == nil {
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

func _content_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<â, û, î, ô, ê>
	if !_accept(parser, _desinence__â__û__î__ô__êAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _content_desinence, start, pos, perr)
fail:
	return _memoize(parser, _content_desinence, start, -1, perr)
}

func _content_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_desinence"}
	// desinence<â, û, î, ô, ê>
	if !_node(parser, _desinence__â__û__î__ô__êNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_desinence}
	// desinence<â, û, î, ô, ê>
	if !_fail(parser, _desinence__â__û__î__ô__êFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<â, û, î, ô, ê>
	if p, n := _desinence__â__û__î__ô__êAction(parser, pos); n == nil {
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

func _preposition_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<à, ù, ì, ò, è>
	if !_accept(parser, _desinence__à__ù__ì__ò__èAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _preposition_desinence, start, pos, perr)
fail:
	return _memoize(parser, _preposition_desinence, start, -1, perr)
}

func _preposition_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition_desinence"}
	// desinence<à, ù, ì, ò, è>
	if !_node(parser, _desinence__à__ù__ì__ò__èNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _preposition_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_desinence}
	// desinence<à, ù, ì, ò, è>
	if !_fail(parser, _desinence__à__ù__ì__ò__èFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<à, ù, ì, ò, è>
	if p, n := _desinence__à__ù__ì__ò__èAction(parser, pos); n == nil {
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

func _adverb_desinenceAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb_desinence, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// desinence<ã, ũ, ĩ, õ, ẽ>
	if !_accept(parser, _desinence__ã__ũ__ĩ__õ__ẽAccepts, &pos, &perr) {
		goto fail
	}
	return _memoize(parser, _adverb_desinence, start, pos, perr)
fail:
	return _memoize(parser, _adverb_desinence, start, -1, perr)
}

func _adverb_desinenceNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_desinence}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb_desinence"}
	// desinence<ã, ũ, ĩ, õ, ẽ>
	if !_node(parser, _desinence__ã__ũ__ĩ__õ__ẽNode, node, &pos) {
		goto fail
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_desinenceFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb_desinence, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_desinence",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_desinence}
	// desinence<ã, ũ, ĩ, õ, ẽ>
	if !_fail(parser, _desinence__ã__ũ__ĩ__õ__ẽFail, errPos, failure, &pos) {
		goto fail
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_desinenceAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb_desinence]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_desinence}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// desinence<ã, ũ, ĩ, õ, ẽ>
	if p, n := _desinence__ã__ũ__ĩ__õ__ẽAction(parser, pos); n == nil {
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

func _toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// compound_tone/arg_tone/verb_tone/relative_tone/adverb_tone/preposition_tone/content_tone
	{
		pos3 := pos
		// compound_tone
		if !_accept(parser, _compound_toneAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_tone
		if !_accept(parser, _arg_toneAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// verb_tone
		if !_accept(parser, _verb_toneAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// relative_tone
		if !_accept(parser, _relative_toneAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// adverb_tone
		if !_accept(parser, _adverb_toneAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// preposition_tone
		if !_accept(parser, _preposition_toneAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// content_tone
		if !_accept(parser, _content_toneAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _tone, start, pos, perr)
fail:
	return _memoize(parser, _tone, start, -1, perr)
}

func _toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "tone"}
	// compound_tone/arg_tone/verb_tone/relative_tone/adverb_tone/preposition_tone/content_tone
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// compound_tone
		if !_node(parser, _compound_toneNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// arg_tone
		if !_node(parser, _arg_toneNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// verb_tone
		if !_node(parser, _verb_toneNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// relative_tone
		if !_node(parser, _relative_toneNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// adverb_tone
		if !_node(parser, _adverb_toneNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// preposition_tone
		if !_node(parser, _preposition_toneNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// content_tone
		if !_node(parser, _content_toneNode, node, &pos) {
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

func _toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _tone}
	// compound_tone/arg_tone/verb_tone/relative_tone/adverb_tone/preposition_tone/content_tone
	{
		pos3 := pos
		// compound_tone
		if !_fail(parser, _compound_toneFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// arg_tone
		if !_fail(parser, _arg_toneFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// verb_tone
		if !_fail(parser, _verb_toneFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// relative_tone
		if !_fail(parser, _relative_toneFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// adverb_tone
		if !_fail(parser, _adverb_toneFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// preposition_tone
		if !_fail(parser, _preposition_toneFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// content_tone
		if !_fail(parser, _content_toneFail, errPos, failure, &pos) {
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

func _toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// compound_tone/arg_tone/verb_tone/relative_tone/adverb_tone/preposition_tone/content_tone
	{
		pos3 := pos
		var node2 string
		// compound_tone
		if p, n := _compound_toneAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// arg_tone
		if p, n := _arg_toneAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// verb_tone
		if p, n := _verb_toneAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// relative_tone
		if p, n := _relative_toneAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// adverb_tone
		if p, n := _adverb_toneAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// preposition_tone
		if p, n := _preposition_toneAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// content_tone
		if p, n := _content_toneAction(parser, pos); n == nil {
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

func _AAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _A, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// ā/á/ǎ/ả/â/à/ã/a
	{
		pos3 := pos
		// ā
		if !_accept(parser, _āAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// á
		if !_accept(parser, _áAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǎ
		if !_accept(parser, _ǎAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ả
		if !_accept(parser, _ảAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// â
		if !_accept(parser, _âAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// à
		if !_accept(parser, _àAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ã
		if !_accept(parser, _ãAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// a
		if !_accept(parser, _aAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _A, start, pos, perr)
fail:
	return _memoize(parser, _A, start, -1, perr)
}

func _ANode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_A]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _A}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "A"}
	// ā/á/ǎ/ả/â/à/ã/a
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// ā
		if !_node(parser, _āNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// á
		if !_node(parser, _áNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ǎ
		if !_node(parser, _ǎNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ả
		if !_node(parser, _ảNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// â
		if !_node(parser, _âNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// à
		if !_node(parser, _àNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ã
		if !_node(parser, _ãNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// a
		if !_node(parser, _aNode, node, &pos) {
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

func _AFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _A, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "A",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _A}
	// ā/á/ǎ/ả/â/à/ã/a
	{
		pos3 := pos
		// ā
		if !_fail(parser, _āFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// á
		if !_fail(parser, _áFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǎ
		if !_fail(parser, _ǎFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ả
		if !_fail(parser, _ảFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// â
		if !_fail(parser, _âFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// à
		if !_fail(parser, _àFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ã
		if !_fail(parser, _ãFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// a
		if !_fail(parser, _aFail, errPos, failure, &pos) {
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

func _AAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_A]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _A}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// ā/á/ǎ/ả/â/à/ã/a
	{
		pos3 := pos
		var node2 string
		// ā
		if p, n := _āAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// á
		if p, n := _áAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// ǎ
		if p, n := _ǎAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// ả
		if p, n := _ảAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// â
		if p, n := _âAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// à
		if p, n := _àAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// ã
		if p, n := _ãAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// a
		if p, n := _aAction(parser, pos); n == nil {
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

func _UAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _U, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// ū/ú/ǔ/ủ/û/ù/ũ/u
	{
		pos3 := pos
		// ū
		if !_accept(parser, _ūAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// ú
		if !_accept(parser, _úAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǔ
		if !_accept(parser, _ǔAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ủ
		if !_accept(parser, _ủAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// û
		if !_accept(parser, _ûAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// ù
		if !_accept(parser, _ùAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ũ
		if !_accept(parser, _ũAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// u
		if !_accept(parser, _uAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _U, start, pos, perr)
fail:
	return _memoize(parser, _U, start, -1, perr)
}

func _UNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_U]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _U}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "U"}
	// ū/ú/ǔ/ủ/û/ù/ũ/u
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// ū
		if !_node(parser, _ūNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ú
		if !_node(parser, _úNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ǔ
		if !_node(parser, _ǔNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ủ
		if !_node(parser, _ủNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// û
		if !_node(parser, _ûNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ù
		if !_node(parser, _ùNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ũ
		if !_node(parser, _ũNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// u
		if !_node(parser, _uNode, node, &pos) {
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

func _UFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _U, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "U",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _U}
	// ū/ú/ǔ/ủ/û/ù/ũ/u
	{
		pos3 := pos
		// ū
		if !_fail(parser, _ūFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// ú
		if !_fail(parser, _úFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǔ
		if !_fail(parser, _ǔFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ủ
		if !_fail(parser, _ủFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// û
		if !_fail(parser, _ûFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// ù
		if !_fail(parser, _ùFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ũ
		if !_fail(parser, _ũFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// u
		if !_fail(parser, _uFail, errPos, failure, &pos) {
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

func _UAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_U]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _U}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// ū/ú/ǔ/ủ/û/ù/ũ/u
	{
		pos3 := pos
		var node2 string
		// ū
		if p, n := _ūAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// ú
		if p, n := _úAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// ǔ
		if p, n := _ǔAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// ủ
		if p, n := _ủAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// û
		if p, n := _ûAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// ù
		if p, n := _ùAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// ũ
		if p, n := _ũAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// u
		if p, n := _uAction(parser, pos); n == nil {
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

func _IAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _I, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// ī/í/ǐ/ỉ/î/ì/ĩ/i
	{
		pos3 := pos
		// ī
		if !_accept(parser, _īAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// í
		if !_accept(parser, _íAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǐ
		if !_accept(parser, _ǐAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ỉ
		if !_accept(parser, _ỉAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// î
		if !_accept(parser, _îAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// ì
		if !_accept(parser, _ìAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ĩ
		if !_accept(parser, _ĩAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// i
		if !_accept(parser, _iAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _I, start, pos, perr)
fail:
	return _memoize(parser, _I, start, -1, perr)
}

func _INode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_I]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _I}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "I"}
	// ī/í/ǐ/ỉ/î/ì/ĩ/i
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// ī
		if !_node(parser, _īNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// í
		if !_node(parser, _íNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ǐ
		if !_node(parser, _ǐNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ỉ
		if !_node(parser, _ỉNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// î
		if !_node(parser, _îNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ì
		if !_node(parser, _ìNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ĩ
		if !_node(parser, _ĩNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// i
		if !_node(parser, _iNode, node, &pos) {
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

func _IFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _I, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "I",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _I}
	// ī/í/ǐ/ỉ/î/ì/ĩ/i
	{
		pos3 := pos
		// ī
		if !_fail(parser, _īFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// í
		if !_fail(parser, _íFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǐ
		if !_fail(parser, _ǐFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ỉ
		if !_fail(parser, _ỉFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// î
		if !_fail(parser, _îFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// ì
		if !_fail(parser, _ìFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ĩ
		if !_fail(parser, _ĩFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// i
		if !_fail(parser, _iFail, errPos, failure, &pos) {
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

func _IAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_I]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _I}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// ī/í/ǐ/ỉ/î/ì/ĩ/i
	{
		pos3 := pos
		var node2 string
		// ī
		if p, n := _īAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// í
		if p, n := _íAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// ǐ
		if p, n := _ǐAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// ỉ
		if p, n := _ỉAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// î
		if p, n := _îAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// ì
		if p, n := _ìAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// ĩ
		if p, n := _ĩAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// i
		if p, n := _iAction(parser, pos); n == nil {
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

func _OAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _O, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// ō/ó/ǒ/ỏ/ô/ò/õ/o
	{
		pos3 := pos
		// ō
		if !_accept(parser, _ōAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// ó
		if !_accept(parser, _óAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǒ
		if !_accept(parser, _ǒAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ỏ
		if !_accept(parser, _ỏAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// ô
		if !_accept(parser, _ôAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// ò
		if !_accept(parser, _òAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// õ
		if !_accept(parser, _õAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// o
		if !_accept(parser, _oAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _O, start, pos, perr)
fail:
	return _memoize(parser, _O, start, -1, perr)
}

func _ONode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_O]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _O}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "O"}
	// ō/ó/ǒ/ỏ/ô/ò/õ/o
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// ō
		if !_node(parser, _ōNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ó
		if !_node(parser, _óNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ǒ
		if !_node(parser, _ǒNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ỏ
		if !_node(parser, _ỏNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ô
		if !_node(parser, _ôNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ò
		if !_node(parser, _òNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// õ
		if !_node(parser, _õNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// o
		if !_node(parser, _oNode, node, &pos) {
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

func _OFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _O, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "O",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _O}
	// ō/ó/ǒ/ỏ/ô/ò/õ/o
	{
		pos3 := pos
		// ō
		if !_fail(parser, _ōFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// ó
		if !_fail(parser, _óFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ǒ
		if !_fail(parser, _ǒFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ỏ
		if !_fail(parser, _ỏFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// ô
		if !_fail(parser, _ôFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// ò
		if !_fail(parser, _òFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// õ
		if !_fail(parser, _õFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// o
		if !_fail(parser, _oFail, errPos, failure, &pos) {
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

func _OAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_O]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _O}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// ō/ó/ǒ/ỏ/ô/ò/õ/o
	{
		pos3 := pos
		var node2 string
		// ō
		if p, n := _ōAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// ó
		if p, n := _óAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// ǒ
		if p, n := _ǒAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// ỏ
		if p, n := _ỏAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// ô
		if p, n := _ôAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// ò
		if p, n := _òAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// õ
		if p, n := _õAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// o
		if p, n := _oAction(parser, pos); n == nil {
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

func _EAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _E, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// ē/é/ě/ẻ/ê/è/ẽ/e
	{
		pos3 := pos
		// ē
		if !_accept(parser, _ēAccepts, &pos, &perr) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// é
		if !_accept(parser, _éAccepts, &pos, &perr) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ě
		if !_accept(parser, _ěAccepts, &pos, &perr) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ẻ
		if !_accept(parser, _ẻAccepts, &pos, &perr) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// ê
		if !_accept(parser, _êAccepts, &pos, &perr) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// è
		if !_accept(parser, _èAccepts, &pos, &perr) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ẽ
		if !_accept(parser, _ẽAccepts, &pos, &perr) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// e
		if !_accept(parser, _eAccepts, &pos, &perr) {
			goto fail11
		}
		goto ok0
	fail11:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _E, start, pos, perr)
fail:
	return _memoize(parser, _E, start, -1, perr)
}

func _ENode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_E]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _E}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "E"}
	// ē/é/ě/ẻ/ê/è/ẽ/e
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// ē
		if !_node(parser, _ēNode, node, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// é
		if !_node(parser, _éNode, node, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ě
		if !_node(parser, _ěNode, node, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ẻ
		if !_node(parser, _ẻNode, node, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ê
		if !_node(parser, _êNode, node, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// è
		if !_node(parser, _èNode, node, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// ẽ
		if !_node(parser, _ẽNode, node, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// e
		if !_node(parser, _eNode, node, &pos) {
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

func _EFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _E, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "E",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _E}
	// ē/é/ě/ẻ/ê/è/ẽ/e
	{
		pos3 := pos
		// ē
		if !_fail(parser, _ēFail, errPos, failure, &pos) {
			goto fail4
		}
		goto ok0
	fail4:
		pos = pos3
		// é
		if !_fail(parser, _éFail, errPos, failure, &pos) {
			goto fail5
		}
		goto ok0
	fail5:
		pos = pos3
		// ě
		if !_fail(parser, _ěFail, errPos, failure, &pos) {
			goto fail6
		}
		goto ok0
	fail6:
		pos = pos3
		// ẻ
		if !_fail(parser, _ẻFail, errPos, failure, &pos) {
			goto fail7
		}
		goto ok0
	fail7:
		pos = pos3
		// ê
		if !_fail(parser, _êFail, errPos, failure, &pos) {
			goto fail8
		}
		goto ok0
	fail8:
		pos = pos3
		// è
		if !_fail(parser, _èFail, errPos, failure, &pos) {
			goto fail9
		}
		goto ok0
	fail9:
		pos = pos3
		// ẽ
		if !_fail(parser, _ẽFail, errPos, failure, &pos) {
			goto fail10
		}
		goto ok0
	fail10:
		pos = pos3
		// e
		if !_fail(parser, _eFail, errPos, failure, &pos) {
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

func _EAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_E]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _E}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// ē/é/ě/ẻ/ê/è/ẽ/e
	{
		pos3 := pos
		var node2 string
		// ē
		if p, n := _ēAction(parser, pos); n == nil {
			goto fail4
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// é
		if p, n := _éAction(parser, pos); n == nil {
			goto fail5
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail5:
		node = node2
		pos = pos3
		// ě
		if p, n := _ěAction(parser, pos); n == nil {
			goto fail6
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail6:
		node = node2
		pos = pos3
		// ẻ
		if p, n := _ẻAction(parser, pos); n == nil {
			goto fail7
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail7:
		node = node2
		pos = pos3
		// ê
		if p, n := _êAction(parser, pos); n == nil {
			goto fail8
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail8:
		node = node2
		pos = pos3
		// è
		if p, n := _èAction(parser, pos); n == nil {
			goto fail9
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail9:
		node = node2
		pos = pos3
		// ẽ
		if p, n := _ẽAction(parser, pos); n == nil {
			goto fail10
		} else {
			node = *n
			pos = p
		}
		goto ok0
	fail10:
		node = node2
		pos = pos3
		// e
		if p, n := _eAction(parser, pos); n == nil {
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

func _āAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ā, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [āĀ]/l:a (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [āĀ]
		if r, w := _next(parser, pos); r != 'ā' && r != 'Ā' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (macron_combiner/compound_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_accept(parser, _aAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_accept(parser, _macron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_accept(parser, _compound_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ā, start, pos, perr)
fail:
	return _memoize(parser, _ā, start, -1, perr)
}

func _āNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ā]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ā}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ā"}
	// [āĀ]/l:a (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [āĀ]
		if r, w := _next(parser, pos); r != 'ā' && r != 'Ā' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:a (macron_combiner/compound_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// macron_combiner/compound_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// macron_combiner
				if !_node(parser, _macron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// compound_tone
				if !_node(parser, _compound_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _āFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ā, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ā",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ā}
	// [āĀ]/l:a (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [āĀ]
		if r, w := _next(parser, pos); r != 'ā' && r != 'Ā' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[āĀ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (macron_combiner/compound_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_fail(parser, _aFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_fail(parser, _macron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_fail(parser, _compound_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _āAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ā]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ā}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [āĀ]/l:a (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [āĀ]
		if r, w := _next(parser, pos); r != 'ā' && r != 'Ā' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:a (macron_combiner/compound_tone)
			// l:a
			{
				pos8 := pos
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (macron_combiner/compound_tone)
			// macron_combiner/compound_tone
			{
				pos12 := pos
				// macron_combiner
				if p, n := _macron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// compound_tone
				if p, n := _compound_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['-'][l])
			}(
				start6, pos, label0)
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

func _ūAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ū, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ūŪ]/l:u (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [ūŪ]
		if r, w := _next(parser, pos); r != 'ū' && r != 'Ū' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (macron_combiner/compound_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_accept(parser, _uAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_accept(parser, _macron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_accept(parser, _compound_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ū, start, pos, perr)
fail:
	return _memoize(parser, _ū, start, -1, perr)
}

func _ūNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ū]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ū}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ū"}
	// [ūŪ]/l:u (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ūŪ]
		if r, w := _next(parser, pos); r != 'ū' && r != 'Ū' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:u (macron_combiner/compound_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_node(parser, _uNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// macron_combiner/compound_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// macron_combiner
				if !_node(parser, _macron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// compound_tone
				if !_node(parser, _compound_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ūFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ū, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ū",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ū}
	// [ūŪ]/l:u (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [ūŪ]
		if r, w := _next(parser, pos); r != 'ū' && r != 'Ū' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ūŪ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (macron_combiner/compound_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_fail(parser, _uFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_fail(parser, _macron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_fail(parser, _compound_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ūAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ū]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ū}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ūŪ]/l:u (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ūŪ]
		if r, w := _next(parser, pos); r != 'ū' && r != 'Ū' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:u (macron_combiner/compound_tone)
			// l:u
			{
				pos8 := pos
				// u
				if p, n := _uAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (macron_combiner/compound_tone)
			// macron_combiner/compound_tone
			{
				pos12 := pos
				// macron_combiner
				if p, n := _macron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// compound_tone
				if p, n := _compound_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['-'][l])
			}(
				start6, pos, label0)
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

func _īAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ī, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [īĪ]/l:i (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [īĪ]
		if r, w := _next(parser, pos); r != 'ī' && r != 'Ī' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (macron_combiner/compound_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_accept(parser, _iAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_accept(parser, _macron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_accept(parser, _compound_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ī, start, pos, perr)
fail:
	return _memoize(parser, _ī, start, -1, perr)
}

func _īNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ī]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ī}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ī"}
	// [īĪ]/l:i (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [īĪ]
		if r, w := _next(parser, pos); r != 'ī' && r != 'Ī' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:i (macron_combiner/compound_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// macron_combiner/compound_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// macron_combiner
				if !_node(parser, _macron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// compound_tone
				if !_node(parser, _compound_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _īFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ī, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ī",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ī}
	// [īĪ]/l:i (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [īĪ]
		if r, w := _next(parser, pos); r != 'ī' && r != 'Ī' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[īĪ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (macron_combiner/compound_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_fail(parser, _iFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_fail(parser, _macron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_fail(parser, _compound_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _īAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ī]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ī}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [īĪ]/l:i (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [īĪ]
		if r, w := _next(parser, pos); r != 'ī' && r != 'Ī' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:i (macron_combiner/compound_tone)
			// l:i
			{
				pos8 := pos
				// i
				if p, n := _iAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (macron_combiner/compound_tone)
			// macron_combiner/compound_tone
			{
				pos12 := pos
				// macron_combiner
				if p, n := _macron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// compound_tone
				if p, n := _compound_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['-'][l])
			}(
				start6, pos, label0)
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

func _ōAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ō, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ōŌ]/l:o (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [ōŌ]
		if r, w := _next(parser, pos); r != 'ō' && r != 'Ō' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (macron_combiner/compound_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_accept(parser, _oAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_accept(parser, _macron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_accept(parser, _compound_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ō, start, pos, perr)
fail:
	return _memoize(parser, _ō, start, -1, perr)
}

func _ōNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ō]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ō}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ō"}
	// [ōŌ]/l:o (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ōŌ]
		if r, w := _next(parser, pos); r != 'ō' && r != 'Ō' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:o (macron_combiner/compound_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// macron_combiner/compound_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// macron_combiner
				if !_node(parser, _macron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// compound_tone
				if !_node(parser, _compound_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ōFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ō, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ō",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ō}
	// [ōŌ]/l:o (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [ōŌ]
		if r, w := _next(parser, pos); r != 'ō' && r != 'Ō' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ōŌ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (macron_combiner/compound_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_fail(parser, _oFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_fail(parser, _macron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_fail(parser, _compound_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ōAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ō]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ō}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ōŌ]/l:o (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ōŌ]
		if r, w := _next(parser, pos); r != 'ō' && r != 'Ō' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:o (macron_combiner/compound_tone)
			// l:o
			{
				pos8 := pos
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (macron_combiner/compound_tone)
			// macron_combiner/compound_tone
			{
				pos12 := pos
				// macron_combiner
				if p, n := _macron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// compound_tone
				if p, n := _compound_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['-'][l])
			}(
				start6, pos, label0)
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

func _ēAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ē, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ēĒ]/l:e (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [ēĒ]
		if r, w := _next(parser, pos); r != 'ē' && r != 'Ē' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (macron_combiner/compound_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_accept(parser, _eAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_accept(parser, _macron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_accept(parser, _compound_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ē, start, pos, perr)
fail:
	return _memoize(parser, _ē, start, -1, perr)
}

func _ēNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ē]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ē}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ē"}
	// [ēĒ]/l:e (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ēĒ]
		if r, w := _next(parser, pos); r != 'ē' && r != 'Ē' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:e (macron_combiner/compound_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// macron_combiner/compound_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// macron_combiner
				if !_node(parser, _macron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// compound_tone
				if !_node(parser, _compound_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ēFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ē, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ē",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ē}
	// [ēĒ]/l:e (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		// [ēĒ]
		if r, w := _next(parser, pos); r != 'ē' && r != 'Ē' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ēĒ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (macron_combiner/compound_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_fail(parser, _eFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (macron_combiner/compound_tone)
		// macron_combiner/compound_tone
		{
			pos11 := pos
			// macron_combiner
			if !_fail(parser, _macron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// compound_tone
			if !_fail(parser, _compound_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ēAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ē]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ē}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ēĒ]/l:e (macron_combiner/compound_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ēĒ]
		if r, w := _next(parser, pos); r != 'ē' && r != 'Ē' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:e (macron_combiner/compound_tone)
			// l:e
			{
				pos8 := pos
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (macron_combiner/compound_tone)
			// macron_combiner/compound_tone
			{
				pos12 := pos
				// macron_combiner
				if p, n := _macron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// compound_tone
				if p, n := _compound_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['-'][l])
			}(
				start6, pos, label0)
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

func _macron_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _macron_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u0304"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̄" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _macron_combiner, start, pos, perr)
fail:
	return _memoize(parser, _macron_combiner, start, -1, perr)
}

func _macron_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_macron_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _macron_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "macron_combiner"}
	// "\u0304"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̄" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _macron_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _macron_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "macron_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _macron_combiner}
	// "\u0304"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̄" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u0304\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌̄"
	parser.fail[key] = failure
	return -1, failure
}

func _macron_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_macron_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _macron_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u0304"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̄" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _compound_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _compound_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [\-1]
	if r, w := _next(parser, pos); r != '-' && r != '1' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _compound_tone, start, pos, perr)
fail:
	return _memoize(parser, _compound_tone, start, -1, perr)
}

func _compound_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_compound_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _compound_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "compound_tone"}
	// [\-1]
	if r, w := _next(parser, pos); r != '-' && r != '1' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _compound_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _compound_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "compound_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _compound_tone}
	// [\-1]
	if r, w := _next(parser, pos); r != '-' && r != '1' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[\\-1]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _compound_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_compound_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _compound_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [\-1]
	if r, w := _next(parser, pos); r != '-' && r != '1' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _áAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _á, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [áÁ]/l:a (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [áÁ]
		if r, w := _next(parser, pos); r != 'á' && r != 'Á' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (acute_combiner/arg_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_accept(parser, _aAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_accept(parser, _acute_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_accept(parser, _arg_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _á, start, pos, perr)
fail:
	return _memoize(parser, _á, start, -1, perr)
}

func _áNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_á]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _á}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "á"}
	// [áÁ]/l:a (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [áÁ]
		if r, w := _next(parser, pos); r != 'á' && r != 'Á' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:a (acute_combiner/arg_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// acute_combiner/arg_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// acute_combiner
				if !_node(parser, _acute_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// arg_tone
				if !_node(parser, _arg_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _áFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _á, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "á",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _á}
	// [áÁ]/l:a (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [áÁ]
		if r, w := _next(parser, pos); r != 'á' && r != 'Á' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[áÁ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (acute_combiner/arg_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_fail(parser, _aFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_fail(parser, _acute_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_fail(parser, _arg_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _áAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_á]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _á}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [áÁ]/l:a (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [áÁ]
		if r, w := _next(parser, pos); r != 'á' && r != 'Á' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:a (acute_combiner/arg_tone)
			// l:a
			{
				pos8 := pos
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (acute_combiner/arg_tone)
			// acute_combiner/arg_tone
			{
				pos12 := pos
				// acute_combiner
				if p, n := _acute_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// arg_tone
				if p, n := _arg_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['/'][l])
			}(
				start6, pos, label0)
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

func _úAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ú, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [úÚ]/l:u (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [úÚ]
		if r, w := _next(parser, pos); r != 'ú' && r != 'Ú' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (acute_combiner/arg_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_accept(parser, _uAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_accept(parser, _acute_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_accept(parser, _arg_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ú, start, pos, perr)
fail:
	return _memoize(parser, _ú, start, -1, perr)
}

func _úNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ú]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ú}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ú"}
	// [úÚ]/l:u (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [úÚ]
		if r, w := _next(parser, pos); r != 'ú' && r != 'Ú' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:u (acute_combiner/arg_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_node(parser, _uNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// acute_combiner/arg_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// acute_combiner
				if !_node(parser, _acute_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// arg_tone
				if !_node(parser, _arg_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _úFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ú, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ú",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ú}
	// [úÚ]/l:u (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [úÚ]
		if r, w := _next(parser, pos); r != 'ú' && r != 'Ú' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[úÚ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (acute_combiner/arg_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_fail(parser, _uFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_fail(parser, _acute_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_fail(parser, _arg_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _úAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ú]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ú}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [úÚ]/l:u (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [úÚ]
		if r, w := _next(parser, pos); r != 'ú' && r != 'Ú' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:u (acute_combiner/arg_tone)
			// l:u
			{
				pos8 := pos
				// u
				if p, n := _uAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (acute_combiner/arg_tone)
			// acute_combiner/arg_tone
			{
				pos12 := pos
				// acute_combiner
				if p, n := _acute_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// arg_tone
				if p, n := _arg_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['/'][l])
			}(
				start6, pos, label0)
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

func _íAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _í, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [íÍ]/l:i (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [íÍ]
		if r, w := _next(parser, pos); r != 'í' && r != 'Í' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (acute_combiner/arg_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_accept(parser, _iAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_accept(parser, _acute_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_accept(parser, _arg_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _í, start, pos, perr)
fail:
	return _memoize(parser, _í, start, -1, perr)
}

func _íNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_í]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _í}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "í"}
	// [íÍ]/l:i (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [íÍ]
		if r, w := _next(parser, pos); r != 'í' && r != 'Í' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:i (acute_combiner/arg_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// acute_combiner/arg_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// acute_combiner
				if !_node(parser, _acute_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// arg_tone
				if !_node(parser, _arg_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _íFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _í, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "í",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _í}
	// [íÍ]/l:i (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [íÍ]
		if r, w := _next(parser, pos); r != 'í' && r != 'Í' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[íÍ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (acute_combiner/arg_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_fail(parser, _iFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_fail(parser, _acute_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_fail(parser, _arg_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _íAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_í]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _í}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [íÍ]/l:i (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [íÍ]
		if r, w := _next(parser, pos); r != 'í' && r != 'Í' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:i (acute_combiner/arg_tone)
			// l:i
			{
				pos8 := pos
				// i
				if p, n := _iAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (acute_combiner/arg_tone)
			// acute_combiner/arg_tone
			{
				pos12 := pos
				// acute_combiner
				if p, n := _acute_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// arg_tone
				if p, n := _arg_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['/'][l])
			}(
				start6, pos, label0)
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

func _óAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ó, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [óÓ]/l:o (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [óÓ]
		if r, w := _next(parser, pos); r != 'ó' && r != 'Ó' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (acute_combiner/arg_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_accept(parser, _oAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_accept(parser, _acute_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_accept(parser, _arg_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ó, start, pos, perr)
fail:
	return _memoize(parser, _ó, start, -1, perr)
}

func _óNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ó]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ó}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ó"}
	// [óÓ]/l:o (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [óÓ]
		if r, w := _next(parser, pos); r != 'ó' && r != 'Ó' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:o (acute_combiner/arg_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// acute_combiner/arg_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// acute_combiner
				if !_node(parser, _acute_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// arg_tone
				if !_node(parser, _arg_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _óFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ó, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ó",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ó}
	// [óÓ]/l:o (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [óÓ]
		if r, w := _next(parser, pos); r != 'ó' && r != 'Ó' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[óÓ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (acute_combiner/arg_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_fail(parser, _oFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_fail(parser, _acute_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_fail(parser, _arg_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _óAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ó]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ó}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [óÓ]/l:o (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [óÓ]
		if r, w := _next(parser, pos); r != 'ó' && r != 'Ó' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:o (acute_combiner/arg_tone)
			// l:o
			{
				pos8 := pos
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (acute_combiner/arg_tone)
			// acute_combiner/arg_tone
			{
				pos12 := pos
				// acute_combiner
				if p, n := _acute_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// arg_tone
				if p, n := _arg_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['/'][l])
			}(
				start6, pos, label0)
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

func _éAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _é, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [éÉ]/l:e (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [éÉ]
		if r, w := _next(parser, pos); r != 'é' && r != 'É' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (acute_combiner/arg_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_accept(parser, _eAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_accept(parser, _acute_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_accept(parser, _arg_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _é, start, pos, perr)
fail:
	return _memoize(parser, _é, start, -1, perr)
}

func _éNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_é]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _é}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "é"}
	// [éÉ]/l:e (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [éÉ]
		if r, w := _next(parser, pos); r != 'é' && r != 'É' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:e (acute_combiner/arg_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// acute_combiner/arg_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// acute_combiner
				if !_node(parser, _acute_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// arg_tone
				if !_node(parser, _arg_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _éFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _é, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "é",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _é}
	// [éÉ]/l:e (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		// [éÉ]
		if r, w := _next(parser, pos); r != 'é' && r != 'É' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[éÉ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (acute_combiner/arg_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_fail(parser, _eFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (acute_combiner/arg_tone)
		// acute_combiner/arg_tone
		{
			pos11 := pos
			// acute_combiner
			if !_fail(parser, _acute_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// arg_tone
			if !_fail(parser, _arg_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _éAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_é]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _é}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [éÉ]/l:e (acute_combiner/arg_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [éÉ]
		if r, w := _next(parser, pos); r != 'é' && r != 'É' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:e (acute_combiner/arg_tone)
			// l:e
			{
				pos8 := pos
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (acute_combiner/arg_tone)
			// acute_combiner/arg_tone
			{
				pos12 := pos
				// acute_combiner
				if p, n := _acute_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// arg_tone
				if p, n := _arg_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['/'][l])
			}(
				start6, pos, label0)
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

func _acute_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _acute_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u0301"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "́" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _acute_combiner, start, pos, perr)
fail:
	return _memoize(parser, _acute_combiner, start, -1, perr)
}

func _acute_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_acute_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _acute_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "acute_combiner"}
	// "\u0301"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "́" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _acute_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _acute_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "acute_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _acute_combiner}
	// "\u0301"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "́" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u0301\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌́"
	parser.fail[key] = failure
	return -1, failure
}

func _acute_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_acute_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _acute_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u0301"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "́" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _arg_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _arg_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [/2]
	if r, w := _next(parser, pos); r != '/' && r != '2' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _arg_tone, start, pos, perr)
fail:
	return _memoize(parser, _arg_tone, start, -1, perr)
}

func _arg_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "arg_tone"}
	// [/2]
	if r, w := _next(parser, pos); r != '/' && r != '2' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _arg_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _arg_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "arg_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _arg_tone}
	// [/2]
	if r, w := _next(parser, pos); r != '/' && r != '2' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[/2]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _arg_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_arg_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _arg_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [/2]
	if r, w := _next(parser, pos); r != '/' && r != '2' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _ǎAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ǎ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ǎǍ]/l:a (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǎǍ]
		if r, w := _next(parser, pos); r != 'ǎ' && r != 'Ǎ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (caron_combiner/breve_combiner/relative_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_accept(parser, _aAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_accept(parser, _caron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_accept(parser, _breve_combinerAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_accept(parser, _relative_toneAccepts, &pos, &perr) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ǎ, start, pos, perr)
fail:
	return _memoize(parser, _ǎ, start, -1, perr)
}

func _ǎNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ǎ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǎ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ǎ"}
	// [ǎǍ]/l:a (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ǎǍ]
		if r, w := _next(parser, pos); r != 'ǎ' && r != 'Ǎ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:a (caron_combiner/breve_combiner/relative_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// caron_combiner/breve_combiner/relative_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// caron_combiner
				if !_node(parser, _caron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// breve_combiner
				if !_node(parser, _breve_combinerNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// relative_tone
				if !_node(parser, _relative_toneNode, node, &pos) {
					goto fail16
				}
				goto ok10
			fail16:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ǎFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ǎ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ǎ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ǎ}
	// [ǎǍ]/l:a (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǎǍ]
		if r, w := _next(parser, pos); r != 'ǎ' && r != 'Ǎ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ǎǍ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (caron_combiner/breve_combiner/relative_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_fail(parser, _aFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_fail(parser, _caron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_fail(parser, _breve_combinerFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_fail(parser, _relative_toneFail, errPos, failure, &pos) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
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

func _ǎAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ǎ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǎ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ǎǍ]/l:a (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ǎǍ]
		if r, w := _next(parser, pos); r != 'ǎ' && r != 'Ǎ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:a (caron_combiner/breve_combiner/relative_tone)
			// l:a
			{
				pos8 := pos
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (caron_combiner/breve_combiner/relative_tone)
			// caron_combiner/breve_combiner/relative_tone
			{
				pos12 := pos
				// caron_combiner
				if p, n := _caron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// breve_combiner
				if p, n := _breve_combinerAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				// relative_tone
				if p, n := _relative_toneAction(parser, pos); n == nil {
					goto fail15
				} else {
					pos = p
				}
				goto ok9
			fail15:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['V'][l])
			}(
				start6, pos, label0)
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

func _ǔAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ǔ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ǔǓ]/l:u (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǔǓ]
		if r, w := _next(parser, pos); r != 'ǔ' && r != 'Ǔ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (caron_combiner/breve_combiner/relative_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_accept(parser, _uAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_accept(parser, _caron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_accept(parser, _breve_combinerAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_accept(parser, _relative_toneAccepts, &pos, &perr) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ǔ, start, pos, perr)
fail:
	return _memoize(parser, _ǔ, start, -1, perr)
}

func _ǔNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ǔ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǔ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ǔ"}
	// [ǔǓ]/l:u (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ǔǓ]
		if r, w := _next(parser, pos); r != 'ǔ' && r != 'Ǔ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:u (caron_combiner/breve_combiner/relative_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_node(parser, _uNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// caron_combiner/breve_combiner/relative_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// caron_combiner
				if !_node(parser, _caron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// breve_combiner
				if !_node(parser, _breve_combinerNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// relative_tone
				if !_node(parser, _relative_toneNode, node, &pos) {
					goto fail16
				}
				goto ok10
			fail16:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ǔFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ǔ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ǔ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ǔ}
	// [ǔǓ]/l:u (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǔǓ]
		if r, w := _next(parser, pos); r != 'ǔ' && r != 'Ǔ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ǔǓ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (caron_combiner/breve_combiner/relative_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_fail(parser, _uFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_fail(parser, _caron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_fail(parser, _breve_combinerFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_fail(parser, _relative_toneFail, errPos, failure, &pos) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
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

func _ǔAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ǔ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǔ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ǔǓ]/l:u (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ǔǓ]
		if r, w := _next(parser, pos); r != 'ǔ' && r != 'Ǔ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:u (caron_combiner/breve_combiner/relative_tone)
			// l:u
			{
				pos8 := pos
				// u
				if p, n := _uAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (caron_combiner/breve_combiner/relative_tone)
			// caron_combiner/breve_combiner/relative_tone
			{
				pos12 := pos
				// caron_combiner
				if p, n := _caron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// breve_combiner
				if p, n := _breve_combinerAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				// relative_tone
				if p, n := _relative_toneAction(parser, pos); n == nil {
					goto fail15
				} else {
					pos = p
				}
				goto ok9
			fail15:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['V'][l])
			}(
				start6, pos, label0)
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

func _ǐAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ǐ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ǐǏ]/l:i (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǐǏ]
		if r, w := _next(parser, pos); r != 'ǐ' && r != 'Ǐ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (caron_combiner/breve_combiner/relative_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_accept(parser, _iAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_accept(parser, _caron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_accept(parser, _breve_combinerAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_accept(parser, _relative_toneAccepts, &pos, &perr) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ǐ, start, pos, perr)
fail:
	return _memoize(parser, _ǐ, start, -1, perr)
}

func _ǐNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ǐ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǐ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ǐ"}
	// [ǐǏ]/l:i (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ǐǏ]
		if r, w := _next(parser, pos); r != 'ǐ' && r != 'Ǐ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:i (caron_combiner/breve_combiner/relative_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// caron_combiner/breve_combiner/relative_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// caron_combiner
				if !_node(parser, _caron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// breve_combiner
				if !_node(parser, _breve_combinerNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// relative_tone
				if !_node(parser, _relative_toneNode, node, &pos) {
					goto fail16
				}
				goto ok10
			fail16:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ǐFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ǐ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ǐ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ǐ}
	// [ǐǏ]/l:i (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǐǏ]
		if r, w := _next(parser, pos); r != 'ǐ' && r != 'Ǐ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ǐǏ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (caron_combiner/breve_combiner/relative_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_fail(parser, _iFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_fail(parser, _caron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_fail(parser, _breve_combinerFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_fail(parser, _relative_toneFail, errPos, failure, &pos) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
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

func _ǐAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ǐ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǐ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ǐǏ]/l:i (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ǐǏ]
		if r, w := _next(parser, pos); r != 'ǐ' && r != 'Ǐ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:i (caron_combiner/breve_combiner/relative_tone)
			// l:i
			{
				pos8 := pos
				// i
				if p, n := _iAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (caron_combiner/breve_combiner/relative_tone)
			// caron_combiner/breve_combiner/relative_tone
			{
				pos12 := pos
				// caron_combiner
				if p, n := _caron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// breve_combiner
				if p, n := _breve_combinerAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				// relative_tone
				if p, n := _relative_toneAction(parser, pos); n == nil {
					goto fail15
				} else {
					pos = p
				}
				goto ok9
			fail15:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['V'][l])
			}(
				start6, pos, label0)
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

func _ǒAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ǒ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ǒǑ]/l:o (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǒǑ]
		if r, w := _next(parser, pos); r != 'ǒ' && r != 'Ǒ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (caron_combiner/breve_combiner/relative_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_accept(parser, _oAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_accept(parser, _caron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_accept(parser, _breve_combinerAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_accept(parser, _relative_toneAccepts, &pos, &perr) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ǒ, start, pos, perr)
fail:
	return _memoize(parser, _ǒ, start, -1, perr)
}

func _ǒNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ǒ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǒ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ǒ"}
	// [ǒǑ]/l:o (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ǒǑ]
		if r, w := _next(parser, pos); r != 'ǒ' && r != 'Ǒ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:o (caron_combiner/breve_combiner/relative_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// caron_combiner/breve_combiner/relative_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// caron_combiner
				if !_node(parser, _caron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// breve_combiner
				if !_node(parser, _breve_combinerNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// relative_tone
				if !_node(parser, _relative_toneNode, node, &pos) {
					goto fail16
				}
				goto ok10
			fail16:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ǒFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ǒ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ǒ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ǒ}
	// [ǒǑ]/l:o (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ǒǑ]
		if r, w := _next(parser, pos); r != 'ǒ' && r != 'Ǒ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ǒǑ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (caron_combiner/breve_combiner/relative_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_fail(parser, _oFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_fail(parser, _caron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_fail(parser, _breve_combinerFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_fail(parser, _relative_toneFail, errPos, failure, &pos) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
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

func _ǒAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ǒ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ǒ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ǒǑ]/l:o (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ǒǑ]
		if r, w := _next(parser, pos); r != 'ǒ' && r != 'Ǒ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:o (caron_combiner/breve_combiner/relative_tone)
			// l:o
			{
				pos8 := pos
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (caron_combiner/breve_combiner/relative_tone)
			// caron_combiner/breve_combiner/relative_tone
			{
				pos12 := pos
				// caron_combiner
				if p, n := _caron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// breve_combiner
				if p, n := _breve_combinerAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				// relative_tone
				if p, n := _relative_toneAction(parser, pos); n == nil {
					goto fail15
				} else {
					pos = p
				}
				goto ok9
			fail15:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['V'][l])
			}(
				start6, pos, label0)
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

func _ěAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ě, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ěĚ]/l:e (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ěĚ]
		if r, w := _next(parser, pos); r != 'ě' && r != 'Ě' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (caron_combiner/breve_combiner/relative_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_accept(parser, _eAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_accept(parser, _caron_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_accept(parser, _breve_combinerAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_accept(parser, _relative_toneAccepts, &pos, &perr) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ě, start, pos, perr)
fail:
	return _memoize(parser, _ě, start, -1, perr)
}

func _ěNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ě]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ě}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ě"}
	// [ěĚ]/l:e (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ěĚ]
		if r, w := _next(parser, pos); r != 'ě' && r != 'Ě' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:e (caron_combiner/breve_combiner/relative_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// caron_combiner/breve_combiner/relative_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// caron_combiner
				if !_node(parser, _caron_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// breve_combiner
				if !_node(parser, _breve_combinerNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// relative_tone
				if !_node(parser, _relative_toneNode, node, &pos) {
					goto fail16
				}
				goto ok10
			fail16:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ěFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ě, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ě",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ě}
	// [ěĚ]/l:e (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		// [ěĚ]
		if r, w := _next(parser, pos); r != 'ě' && r != 'Ě' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ěĚ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (caron_combiner/breve_combiner/relative_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_fail(parser, _eFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (caron_combiner/breve_combiner/relative_tone)
		// caron_combiner/breve_combiner/relative_tone
		{
			pos11 := pos
			// caron_combiner
			if !_fail(parser, _caron_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// breve_combiner
			if !_fail(parser, _breve_combinerFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			// relative_tone
			if !_fail(parser, _relative_toneFail, errPos, failure, &pos) {
				goto fail14
			}
			goto ok8
		fail14:
			pos = pos11
			goto fail5
		ok8:
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

func _ěAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ě]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ě}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ěĚ]/l:e (caron_combiner/breve_combiner/relative_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ěĚ]
		if r, w := _next(parser, pos); r != 'ě' && r != 'Ě' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:e (caron_combiner/breve_combiner/relative_tone)
			// l:e
			{
				pos8 := pos
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (caron_combiner/breve_combiner/relative_tone)
			// caron_combiner/breve_combiner/relative_tone
			{
				pos12 := pos
				// caron_combiner
				if p, n := _caron_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// breve_combiner
				if p, n := _breve_combinerAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				// relative_tone
				if p, n := _relative_toneAction(parser, pos); n == nil {
					goto fail15
				} else {
					pos = p
				}
				goto ok9
			fail15:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['V'][l])
			}(
				start6, pos, label0)
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

func _caron_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _caron_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u030c"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̌" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _caron_combiner, start, pos, perr)
fail:
	return _memoize(parser, _caron_combiner, start, -1, perr)
}

func _caron_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_caron_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _caron_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "caron_combiner"}
	// "\u030c"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̌" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _caron_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _caron_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "caron_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _caron_combiner}
	// "\u030c"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̌" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u030c\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌̌"
	parser.fail[key] = failure
	return -1, failure
}

func _caron_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_caron_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _caron_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u030c"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̌" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _breve_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _breve_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u0306"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̆" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _breve_combiner, start, pos, perr)
fail:
	return _memoize(parser, _breve_combiner, start, -1, perr)
}

func _breve_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_breve_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _breve_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "breve_combiner"}
	// "\u0306"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̆" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _breve_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _breve_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "breve_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _breve_combiner}
	// "\u0306"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̆" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u0306\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌̆"
	parser.fail[key] = failure
	return -1, failure
}

func _breve_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_breve_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _breve_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u0306"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̆" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _relative_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _relative_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [V3]
	if r, w := _next(parser, pos); r != 'V' && r != '3' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _relative_tone, start, pos, perr)
fail:
	return _memoize(parser, _relative_tone, start, -1, perr)
}

func _relative_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_relative_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "relative_tone"}
	// [V3]
	if r, w := _next(parser, pos); r != 'V' && r != '3' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _relative_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _relative_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "relative_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _relative_tone}
	// [V3]
	if r, w := _next(parser, pos); r != 'V' && r != '3' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[V3]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _relative_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_relative_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _relative_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [V3]
	if r, w := _next(parser, pos); r != 'V' && r != '3' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _ảAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ả, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ảẢ]/l:a (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ảẢ]
		if r, w := _next(parser, pos); r != 'ả' && r != 'Ả' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (hook_combiner/verb_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_accept(parser, _aAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_accept(parser, _hook_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_accept(parser, _verb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ả, start, pos, perr)
fail:
	return _memoize(parser, _ả, start, -1, perr)
}

func _ảNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ả]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ả}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ả"}
	// [ảẢ]/l:a (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ảẢ]
		if r, w := _next(parser, pos); r != 'ả' && r != 'Ả' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:a (hook_combiner/verb_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// hook_combiner/verb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// hook_combiner
				if !_node(parser, _hook_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// verb_tone
				if !_node(parser, _verb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ảFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ả, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ả",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ả}
	// [ảẢ]/l:a (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ảẢ]
		if r, w := _next(parser, pos); r != 'ả' && r != 'Ả' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ảẢ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (hook_combiner/verb_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_fail(parser, _aFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_fail(parser, _hook_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_fail(parser, _verb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ảAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ả]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ả}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ảẢ]/l:a (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ảẢ]
		if r, w := _next(parser, pos); r != 'ả' && r != 'Ả' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:a (hook_combiner/verb_tone)
			// l:a
			{
				pos8 := pos
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (hook_combiner/verb_tone)
			// hook_combiner/verb_tone
			{
				pos12 := pos
				// hook_combiner
				if p, n := _hook_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// verb_tone
				if p, n := _verb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['?'][l])
			}(
				start6, pos, label0)
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

func _ủAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ủ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ủỦ]/l:u (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ủỦ]
		if r, w := _next(parser, pos); r != 'ủ' && r != 'Ủ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (hook_combiner/verb_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_accept(parser, _uAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_accept(parser, _hook_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_accept(parser, _verb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ủ, start, pos, perr)
fail:
	return _memoize(parser, _ủ, start, -1, perr)
}

func _ủNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ủ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ủ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ủ"}
	// [ủỦ]/l:u (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ủỦ]
		if r, w := _next(parser, pos); r != 'ủ' && r != 'Ủ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:u (hook_combiner/verb_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_node(parser, _uNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// hook_combiner/verb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// hook_combiner
				if !_node(parser, _hook_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// verb_tone
				if !_node(parser, _verb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ủFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ủ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ủ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ủ}
	// [ủỦ]/l:u (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ủỦ]
		if r, w := _next(parser, pos); r != 'ủ' && r != 'Ủ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ủỦ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (hook_combiner/verb_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_fail(parser, _uFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_fail(parser, _hook_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_fail(parser, _verb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ủAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ủ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ủ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ủỦ]/l:u (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ủỦ]
		if r, w := _next(parser, pos); r != 'ủ' && r != 'Ủ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:u (hook_combiner/verb_tone)
			// l:u
			{
				pos8 := pos
				// u
				if p, n := _uAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (hook_combiner/verb_tone)
			// hook_combiner/verb_tone
			{
				pos12 := pos
				// hook_combiner
				if p, n := _hook_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// verb_tone
				if p, n := _verb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['?'][l])
			}(
				start6, pos, label0)
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

func _ỉAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ỉ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ỉỈ]/l:i (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ỉỈ]
		if r, w := _next(parser, pos); r != 'ỉ' && r != 'Ỉ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (hook_combiner/verb_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_accept(parser, _iAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_accept(parser, _hook_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_accept(parser, _verb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ỉ, start, pos, perr)
fail:
	return _memoize(parser, _ỉ, start, -1, perr)
}

func _ỉNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ỉ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ỉ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ỉ"}
	// [ỉỈ]/l:i (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ỉỈ]
		if r, w := _next(parser, pos); r != 'ỉ' && r != 'Ỉ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:i (hook_combiner/verb_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// hook_combiner/verb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// hook_combiner
				if !_node(parser, _hook_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// verb_tone
				if !_node(parser, _verb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ỉFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ỉ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ỉ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ỉ}
	// [ỉỈ]/l:i (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ỉỈ]
		if r, w := _next(parser, pos); r != 'ỉ' && r != 'Ỉ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ỉỈ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (hook_combiner/verb_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_fail(parser, _iFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_fail(parser, _hook_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_fail(parser, _verb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ỉAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ỉ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ỉ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ỉỈ]/l:i (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ỉỈ]
		if r, w := _next(parser, pos); r != 'ỉ' && r != 'Ỉ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:i (hook_combiner/verb_tone)
			// l:i
			{
				pos8 := pos
				// i
				if p, n := _iAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (hook_combiner/verb_tone)
			// hook_combiner/verb_tone
			{
				pos12 := pos
				// hook_combiner
				if p, n := _hook_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// verb_tone
				if p, n := _verb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['?'][l])
			}(
				start6, pos, label0)
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

func _ỏAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ỏ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ỏỎ]/l:o (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ỏỎ]
		if r, w := _next(parser, pos); r != 'ỏ' && r != 'Ỏ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (hook_combiner/verb_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_accept(parser, _oAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_accept(parser, _hook_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_accept(parser, _verb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ỏ, start, pos, perr)
fail:
	return _memoize(parser, _ỏ, start, -1, perr)
}

func _ỏNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ỏ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ỏ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ỏ"}
	// [ỏỎ]/l:o (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ỏỎ]
		if r, w := _next(parser, pos); r != 'ỏ' && r != 'Ỏ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:o (hook_combiner/verb_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// hook_combiner/verb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// hook_combiner
				if !_node(parser, _hook_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// verb_tone
				if !_node(parser, _verb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ỏFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ỏ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ỏ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ỏ}
	// [ỏỎ]/l:o (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ỏỎ]
		if r, w := _next(parser, pos); r != 'ỏ' && r != 'Ỏ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ỏỎ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (hook_combiner/verb_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_fail(parser, _oFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_fail(parser, _hook_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_fail(parser, _verb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ỏAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ỏ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ỏ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ỏỎ]/l:o (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ỏỎ]
		if r, w := _next(parser, pos); r != 'ỏ' && r != 'Ỏ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:o (hook_combiner/verb_tone)
			// l:o
			{
				pos8 := pos
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (hook_combiner/verb_tone)
			// hook_combiner/verb_tone
			{
				pos12 := pos
				// hook_combiner
				if p, n := _hook_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// verb_tone
				if p, n := _verb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['?'][l])
			}(
				start6, pos, label0)
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

func _ẻAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ẻ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ẻẺ]/l:e (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ẻẺ]
		if r, w := _next(parser, pos); r != 'ẻ' && r != 'Ẻ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (hook_combiner/verb_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_accept(parser, _eAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_accept(parser, _hook_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_accept(parser, _verb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ẻ, start, pos, perr)
fail:
	return _memoize(parser, _ẻ, start, -1, perr)
}

func _ẻNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ẻ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ẻ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ẻ"}
	// [ẻẺ]/l:e (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ẻẺ]
		if r, w := _next(parser, pos); r != 'ẻ' && r != 'Ẻ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:e (hook_combiner/verb_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// hook_combiner/verb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// hook_combiner
				if !_node(parser, _hook_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// verb_tone
				if !_node(parser, _verb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ẻFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ẻ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ẻ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ẻ}
	// [ẻẺ]/l:e (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		// [ẻẺ]
		if r, w := _next(parser, pos); r != 'ẻ' && r != 'Ẻ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ẻẺ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (hook_combiner/verb_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_fail(parser, _eFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (hook_combiner/verb_tone)
		// hook_combiner/verb_tone
		{
			pos11 := pos
			// hook_combiner
			if !_fail(parser, _hook_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// verb_tone
			if !_fail(parser, _verb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ẻAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ẻ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ẻ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ẻẺ]/l:e (hook_combiner/verb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ẻẺ]
		if r, w := _next(parser, pos); r != 'ẻ' && r != 'Ẻ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:e (hook_combiner/verb_tone)
			// l:e
			{
				pos8 := pos
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (hook_combiner/verb_tone)
			// hook_combiner/verb_tone
			{
				pos12 := pos
				// hook_combiner
				if p, n := _hook_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// verb_tone
				if p, n := _verb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['?'][l])
			}(
				start6, pos, label0)
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

func _hook_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _hook_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u0309"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̉" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _hook_combiner, start, pos, perr)
fail:
	return _memoize(parser, _hook_combiner, start, -1, perr)
}

func _hook_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_hook_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _hook_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "hook_combiner"}
	// "\u0309"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̉" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _hook_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _hook_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "hook_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _hook_combiner}
	// "\u0309"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̉" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u0309\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌̉"
	parser.fail[key] = failure
	return -1, failure
}

func _hook_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_hook_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _hook_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u0309"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̉" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _verb_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _verb_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [?4]
	if r, w := _next(parser, pos); r != '?' && r != '4' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _verb_tone, start, pos, perr)
fail:
	return _memoize(parser, _verb_tone, start, -1, perr)
}

func _verb_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_verb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _verb_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "verb_tone"}
	// [?4]
	if r, w := _next(parser, pos); r != '?' && r != '4' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _verb_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _verb_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "verb_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _verb_tone}
	// [?4]
	if r, w := _next(parser, pos); r != '?' && r != '4' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[?4]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _verb_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_verb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _verb_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [?4]
	if r, w := _next(parser, pos); r != '?' && r != '4' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _âAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _â, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [âÂ]/l:a (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [âÂ]
		if r, w := _next(parser, pos); r != 'â' && r != 'Â' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (circumflex_combiner/content_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_accept(parser, _aAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_accept(parser, _circumflex_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_accept(parser, _content_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _â, start, pos, perr)
fail:
	return _memoize(parser, _â, start, -1, perr)
}

func _âNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_â]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _â}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "â"}
	// [âÂ]/l:a (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [âÂ]
		if r, w := _next(parser, pos); r != 'â' && r != 'Â' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:a (circumflex_combiner/content_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// circumflex_combiner/content_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// circumflex_combiner
				if !_node(parser, _circumflex_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// content_tone
				if !_node(parser, _content_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _âFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _â, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "â",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _â}
	// [âÂ]/l:a (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [âÂ]
		if r, w := _next(parser, pos); r != 'â' && r != 'Â' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[âÂ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (circumflex_combiner/content_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_fail(parser, _aFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_fail(parser, _circumflex_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_fail(parser, _content_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _âAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_â]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _â}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [âÂ]/l:a (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [âÂ]
		if r, w := _next(parser, pos); r != 'â' && r != 'Â' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:a (circumflex_combiner/content_tone)
			// l:a
			{
				pos8 := pos
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (circumflex_combiner/content_tone)
			// circumflex_combiner/content_tone
			{
				pos12 := pos
				// circumflex_combiner
				if p, n := _circumflex_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// content_tone
				if p, n := _content_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['^'][l])
			}(
				start6, pos, label0)
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

func _ûAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _û, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ûÛ]/l:u (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [ûÛ]
		if r, w := _next(parser, pos); r != 'û' && r != 'Û' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (circumflex_combiner/content_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_accept(parser, _uAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_accept(parser, _circumflex_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_accept(parser, _content_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _û, start, pos, perr)
fail:
	return _memoize(parser, _û, start, -1, perr)
}

func _ûNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_û]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _û}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "û"}
	// [ûÛ]/l:u (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ûÛ]
		if r, w := _next(parser, pos); r != 'û' && r != 'Û' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:u (circumflex_combiner/content_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_node(parser, _uNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// circumflex_combiner/content_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// circumflex_combiner
				if !_node(parser, _circumflex_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// content_tone
				if !_node(parser, _content_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ûFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _û, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "û",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _û}
	// [ûÛ]/l:u (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [ûÛ]
		if r, w := _next(parser, pos); r != 'û' && r != 'Û' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ûÛ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (circumflex_combiner/content_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_fail(parser, _uFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_fail(parser, _circumflex_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_fail(parser, _content_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ûAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_û]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _û}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ûÛ]/l:u (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ûÛ]
		if r, w := _next(parser, pos); r != 'û' && r != 'Û' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:u (circumflex_combiner/content_tone)
			// l:u
			{
				pos8 := pos
				// u
				if p, n := _uAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (circumflex_combiner/content_tone)
			// circumflex_combiner/content_tone
			{
				pos12 := pos
				// circumflex_combiner
				if p, n := _circumflex_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// content_tone
				if p, n := _content_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['^'][l])
			}(
				start6, pos, label0)
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

func _îAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _î, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [îÎ]/l:i (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [îÎ]
		if r, w := _next(parser, pos); r != 'î' && r != 'Î' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (circumflex_combiner/content_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_accept(parser, _iAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_accept(parser, _circumflex_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_accept(parser, _content_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _î, start, pos, perr)
fail:
	return _memoize(parser, _î, start, -1, perr)
}

func _îNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_î]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _î}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "î"}
	// [îÎ]/l:i (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [îÎ]
		if r, w := _next(parser, pos); r != 'î' && r != 'Î' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:i (circumflex_combiner/content_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// circumflex_combiner/content_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// circumflex_combiner
				if !_node(parser, _circumflex_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// content_tone
				if !_node(parser, _content_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _îFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _î, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "î",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _î}
	// [îÎ]/l:i (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [îÎ]
		if r, w := _next(parser, pos); r != 'î' && r != 'Î' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[îÎ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (circumflex_combiner/content_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_fail(parser, _iFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_fail(parser, _circumflex_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_fail(parser, _content_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _îAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_î]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _î}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [îÎ]/l:i (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [îÎ]
		if r, w := _next(parser, pos); r != 'î' && r != 'Î' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:i (circumflex_combiner/content_tone)
			// l:i
			{
				pos8 := pos
				// i
				if p, n := _iAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (circumflex_combiner/content_tone)
			// circumflex_combiner/content_tone
			{
				pos12 := pos
				// circumflex_combiner
				if p, n := _circumflex_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// content_tone
				if p, n := _content_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['^'][l])
			}(
				start6, pos, label0)
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

func _ôAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ô, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ôÔ]/l:o (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [ôÔ]
		if r, w := _next(parser, pos); r != 'ô' && r != 'Ô' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (circumflex_combiner/content_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_accept(parser, _oAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_accept(parser, _circumflex_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_accept(parser, _content_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ô, start, pos, perr)
fail:
	return _memoize(parser, _ô, start, -1, perr)
}

func _ôNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ô]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ô}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ô"}
	// [ôÔ]/l:o (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ôÔ]
		if r, w := _next(parser, pos); r != 'ô' && r != 'Ô' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:o (circumflex_combiner/content_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// circumflex_combiner/content_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// circumflex_combiner
				if !_node(parser, _circumflex_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// content_tone
				if !_node(parser, _content_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ôFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ô, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ô",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ô}
	// [ôÔ]/l:o (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [ôÔ]
		if r, w := _next(parser, pos); r != 'ô' && r != 'Ô' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ôÔ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (circumflex_combiner/content_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_fail(parser, _oFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_fail(parser, _circumflex_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_fail(parser, _content_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ôAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ô]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ô}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ôÔ]/l:o (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ôÔ]
		if r, w := _next(parser, pos); r != 'ô' && r != 'Ô' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:o (circumflex_combiner/content_tone)
			// l:o
			{
				pos8 := pos
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (circumflex_combiner/content_tone)
			// circumflex_combiner/content_tone
			{
				pos12 := pos
				// circumflex_combiner
				if p, n := _circumflex_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// content_tone
				if p, n := _content_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['^'][l])
			}(
				start6, pos, label0)
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

func _êAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ê, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [êÊ]/l:e (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [êÊ]
		if r, w := _next(parser, pos); r != 'ê' && r != 'Ê' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (circumflex_combiner/content_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_accept(parser, _eAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_accept(parser, _circumflex_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_accept(parser, _content_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ê, start, pos, perr)
fail:
	return _memoize(parser, _ê, start, -1, perr)
}

func _êNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ê]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ê}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ê"}
	// [êÊ]/l:e (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [êÊ]
		if r, w := _next(parser, pos); r != 'ê' && r != 'Ê' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:e (circumflex_combiner/content_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// circumflex_combiner/content_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// circumflex_combiner
				if !_node(parser, _circumflex_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// content_tone
				if !_node(parser, _content_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _êFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ê, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ê",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ê}
	// [êÊ]/l:e (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		// [êÊ]
		if r, w := _next(parser, pos); r != 'ê' && r != 'Ê' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[êÊ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (circumflex_combiner/content_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_fail(parser, _eFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (circumflex_combiner/content_tone)
		// circumflex_combiner/content_tone
		{
			pos11 := pos
			// circumflex_combiner
			if !_fail(parser, _circumflex_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// content_tone
			if !_fail(parser, _content_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _êAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ê]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ê}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [êÊ]/l:e (circumflex_combiner/content_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [êÊ]
		if r, w := _next(parser, pos); r != 'ê' && r != 'Ê' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:e (circumflex_combiner/content_tone)
			// l:e
			{
				pos8 := pos
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (circumflex_combiner/content_tone)
			// circumflex_combiner/content_tone
			{
				pos12 := pos
				// circumflex_combiner
				if p, n := _circumflex_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// content_tone
				if p, n := _content_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['^'][l])
			}(
				start6, pos, label0)
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

func _circumflex_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _circumflex_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u0302"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̂" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _circumflex_combiner, start, pos, perr)
fail:
	return _memoize(parser, _circumflex_combiner, start, -1, perr)
}

func _circumflex_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_circumflex_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _circumflex_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "circumflex_combiner"}
	// "\u0302"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̂" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _circumflex_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _circumflex_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "circumflex_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _circumflex_combiner}
	// "\u0302"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̂" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u0302\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌̂"
	parser.fail[key] = failure
	return -1, failure
}

func _circumflex_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_circumflex_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _circumflex_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u0302"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̂" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _content_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _content_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [5\^]
	if r, w := _next(parser, pos); r != '5' && r != '^' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _content_tone, start, pos, perr)
fail:
	return _memoize(parser, _content_tone, start, -1, perr)
}

func _content_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "content_tone"}
	// [5\^]
	if r, w := _next(parser, pos); r != '5' && r != '^' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _content_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _content_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "content_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _content_tone}
	// [5\^]
	if r, w := _next(parser, pos); r != '5' && r != '^' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[5\\^]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _content_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_content_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _content_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [5\^]
	if r, w := _next(parser, pos); r != '5' && r != '^' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _àAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _à, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [àÀ]/l:a (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [àÀ]
		if r, w := _next(parser, pos); r != 'à' && r != 'À' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (grave_combiner/preposition_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_accept(parser, _aAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_accept(parser, _grave_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_accept(parser, _preposition_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _à, start, pos, perr)
fail:
	return _memoize(parser, _à, start, -1, perr)
}

func _àNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_à]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _à}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "à"}
	// [àÀ]/l:a (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [àÀ]
		if r, w := _next(parser, pos); r != 'à' && r != 'À' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:a (grave_combiner/preposition_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// grave_combiner/preposition_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// grave_combiner
				if !_node(parser, _grave_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// preposition_tone
				if !_node(parser, _preposition_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _àFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _à, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "à",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _à}
	// [àÀ]/l:a (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [àÀ]
		if r, w := _next(parser, pos); r != 'à' && r != 'À' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[àÀ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (grave_combiner/preposition_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_fail(parser, _aFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_fail(parser, _grave_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_fail(parser, _preposition_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _àAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_à]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _à}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [àÀ]/l:a (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [àÀ]
		if r, w := _next(parser, pos); r != 'à' && r != 'À' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:a (grave_combiner/preposition_tone)
			// l:a
			{
				pos8 := pos
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (grave_combiner/preposition_tone)
			// grave_combiner/preposition_tone
			{
				pos12 := pos
				// grave_combiner
				if p, n := _grave_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// preposition_tone
				if p, n := _preposition_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['\\'][l])
			}(
				start6, pos, label0)
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

func _ùAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ù, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ùÙ]/l:u (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [ùÙ]
		if r, w := _next(parser, pos); r != 'ù' && r != 'Ù' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (grave_combiner/preposition_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_accept(parser, _uAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_accept(parser, _grave_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_accept(parser, _preposition_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ù, start, pos, perr)
fail:
	return _memoize(parser, _ù, start, -1, perr)
}

func _ùNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ù]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ù}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ù"}
	// [ùÙ]/l:u (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ùÙ]
		if r, w := _next(parser, pos); r != 'ù' && r != 'Ù' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:u (grave_combiner/preposition_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_node(parser, _uNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// grave_combiner/preposition_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// grave_combiner
				if !_node(parser, _grave_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// preposition_tone
				if !_node(parser, _preposition_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ùFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ù, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ù",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ù}
	// [ùÙ]/l:u (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [ùÙ]
		if r, w := _next(parser, pos); r != 'ù' && r != 'Ù' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ùÙ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (grave_combiner/preposition_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_fail(parser, _uFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_fail(parser, _grave_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_fail(parser, _preposition_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ùAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ù]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ù}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ùÙ]/l:u (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ùÙ]
		if r, w := _next(parser, pos); r != 'ù' && r != 'Ù' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:u (grave_combiner/preposition_tone)
			// l:u
			{
				pos8 := pos
				// u
				if p, n := _uAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (grave_combiner/preposition_tone)
			// grave_combiner/preposition_tone
			{
				pos12 := pos
				// grave_combiner
				if p, n := _grave_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// preposition_tone
				if p, n := _preposition_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['\\'][l])
			}(
				start6, pos, label0)
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

func _ìAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ì, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ìÌ]/l:i (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [ìÌ]
		if r, w := _next(parser, pos); r != 'ì' && r != 'Ì' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (grave_combiner/preposition_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_accept(parser, _iAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_accept(parser, _grave_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_accept(parser, _preposition_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ì, start, pos, perr)
fail:
	return _memoize(parser, _ì, start, -1, perr)
}

func _ìNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ì]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ì}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ì"}
	// [ìÌ]/l:i (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ìÌ]
		if r, w := _next(parser, pos); r != 'ì' && r != 'Ì' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:i (grave_combiner/preposition_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// grave_combiner/preposition_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// grave_combiner
				if !_node(parser, _grave_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// preposition_tone
				if !_node(parser, _preposition_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ìFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ì, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ì",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ì}
	// [ìÌ]/l:i (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [ìÌ]
		if r, w := _next(parser, pos); r != 'ì' && r != 'Ì' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ìÌ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (grave_combiner/preposition_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_fail(parser, _iFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_fail(parser, _grave_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_fail(parser, _preposition_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ìAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ì]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ì}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ìÌ]/l:i (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ìÌ]
		if r, w := _next(parser, pos); r != 'ì' && r != 'Ì' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:i (grave_combiner/preposition_tone)
			// l:i
			{
				pos8 := pos
				// i
				if p, n := _iAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (grave_combiner/preposition_tone)
			// grave_combiner/preposition_tone
			{
				pos12 := pos
				// grave_combiner
				if p, n := _grave_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// preposition_tone
				if p, n := _preposition_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['\\'][l])
			}(
				start6, pos, label0)
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

func _òAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ò, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [òÒ]/l:o (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [òÒ]
		if r, w := _next(parser, pos); r != 'ò' && r != 'Ò' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (grave_combiner/preposition_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_accept(parser, _oAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_accept(parser, _grave_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_accept(parser, _preposition_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ò, start, pos, perr)
fail:
	return _memoize(parser, _ò, start, -1, perr)
}

func _òNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ò]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ò}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ò"}
	// [òÒ]/l:o (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [òÒ]
		if r, w := _next(parser, pos); r != 'ò' && r != 'Ò' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:o (grave_combiner/preposition_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// grave_combiner/preposition_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// grave_combiner
				if !_node(parser, _grave_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// preposition_tone
				if !_node(parser, _preposition_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _òFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ò, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ò",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ò}
	// [òÒ]/l:o (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [òÒ]
		if r, w := _next(parser, pos); r != 'ò' && r != 'Ò' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[òÒ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (grave_combiner/preposition_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_fail(parser, _oFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_fail(parser, _grave_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_fail(parser, _preposition_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _òAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ò]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ò}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [òÒ]/l:o (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [òÒ]
		if r, w := _next(parser, pos); r != 'ò' && r != 'Ò' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:o (grave_combiner/preposition_tone)
			// l:o
			{
				pos8 := pos
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (grave_combiner/preposition_tone)
			// grave_combiner/preposition_tone
			{
				pos12 := pos
				// grave_combiner
				if p, n := _grave_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// preposition_tone
				if p, n := _preposition_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['\\'][l])
			}(
				start6, pos, label0)
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

func _èAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _è, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [èÈ]/l:e (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [èÈ]
		if r, w := _next(parser, pos); r != 'è' && r != 'È' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (grave_combiner/preposition_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_accept(parser, _eAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_accept(parser, _grave_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_accept(parser, _preposition_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _è, start, pos, perr)
fail:
	return _memoize(parser, _è, start, -1, perr)
}

func _èNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_è]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _è}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "è"}
	// [èÈ]/l:e (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [èÈ]
		if r, w := _next(parser, pos); r != 'è' && r != 'È' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:e (grave_combiner/preposition_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// grave_combiner/preposition_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// grave_combiner
				if !_node(parser, _grave_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// preposition_tone
				if !_node(parser, _preposition_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _èFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _è, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "è",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _è}
	// [èÈ]/l:e (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		// [èÈ]
		if r, w := _next(parser, pos); r != 'è' && r != 'È' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[èÈ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (grave_combiner/preposition_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_fail(parser, _eFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (grave_combiner/preposition_tone)
		// grave_combiner/preposition_tone
		{
			pos11 := pos
			// grave_combiner
			if !_fail(parser, _grave_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// preposition_tone
			if !_fail(parser, _preposition_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _èAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_è]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _è}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [èÈ]/l:e (grave_combiner/preposition_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [èÈ]
		if r, w := _next(parser, pos); r != 'è' && r != 'È' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:e (grave_combiner/preposition_tone)
			// l:e
			{
				pos8 := pos
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (grave_combiner/preposition_tone)
			// grave_combiner/preposition_tone
			{
				pos12 := pos
				// grave_combiner
				if p, n := _grave_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// preposition_tone
				if p, n := _preposition_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['\\'][l])
			}(
				start6, pos, label0)
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

func _grave_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _grave_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u0300"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̀" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _grave_combiner, start, pos, perr)
fail:
	return _memoize(parser, _grave_combiner, start, -1, perr)
}

func _grave_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_grave_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _grave_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "grave_combiner"}
	// "\u0300"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̀" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _grave_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _grave_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "grave_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _grave_combiner}
	// "\u0300"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̀" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u0300\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌̀"
	parser.fail[key] = failure
	return -1, failure
}

func _grave_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_grave_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _grave_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u0300"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̀" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _preposition_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _preposition_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [\\6]
	if r, w := _next(parser, pos); r != '\\' && r != '6' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _preposition_tone, start, pos, perr)
fail:
	return _memoize(parser, _preposition_tone, start, -1, perr)
}

func _preposition_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "preposition_tone"}
	// [\\6]
	if r, w := _next(parser, pos); r != '\\' && r != '6' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _preposition_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _preposition_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "preposition_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _preposition_tone}
	// [\\6]
	if r, w := _next(parser, pos); r != '\\' && r != '6' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[\\\\6]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _preposition_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_preposition_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _preposition_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [\\6]
	if r, w := _next(parser, pos); r != '\\' && r != '6' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _ãAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ã, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ãÃ]/l:a (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ãÃ]
		if r, w := _next(parser, pos); r != 'ã' && r != 'Ã' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (tilde_combiner/adverb_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_accept(parser, _aAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_accept(parser, _tilde_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_accept(parser, _adverb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ã, start, pos, perr)
fail:
	return _memoize(parser, _ã, start, -1, perr)
}

func _ãNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ã]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ã}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ã"}
	// [ãÃ]/l:a (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ãÃ]
		if r, w := _next(parser, pos); r != 'ã' && r != 'Ã' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:a (tilde_combiner/adverb_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_node(parser, _aNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// tilde_combiner/adverb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// tilde_combiner
				if !_node(parser, _tilde_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// adverb_tone
				if !_node(parser, _adverb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ãFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ã, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ã",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ã}
	// [ãÃ]/l:a (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ãÃ]
		if r, w := _next(parser, pos); r != 'ã' && r != 'Ã' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ãÃ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:a (tilde_combiner/adverb_tone)
		// l:a
		{
			pos7 := pos
			// a
			if !_fail(parser, _aFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_fail(parser, _tilde_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_fail(parser, _adverb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ãAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ã]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ã}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ãÃ]/l:a (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ãÃ]
		if r, w := _next(parser, pos); r != 'ã' && r != 'Ã' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:a (tilde_combiner/adverb_tone)
			// l:a
			{
				pos8 := pos
				// a
				if p, n := _aAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (tilde_combiner/adverb_tone)
			// tilde_combiner/adverb_tone
			{
				pos12 := pos
				// tilde_combiner
				if p, n := _tilde_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// adverb_tone
				if p, n := _adverb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['~'][l])
			}(
				start6, pos, label0)
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

func _ũAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ũ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ũŨ]/l:u (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ũŨ]
		if r, w := _next(parser, pos); r != 'ũ' && r != 'Ũ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (tilde_combiner/adverb_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_accept(parser, _uAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_accept(parser, _tilde_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_accept(parser, _adverb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ũ, start, pos, perr)
fail:
	return _memoize(parser, _ũ, start, -1, perr)
}

func _ũNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ũ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ũ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ũ"}
	// [ũŨ]/l:u (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ũŨ]
		if r, w := _next(parser, pos); r != 'ũ' && r != 'Ũ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:u (tilde_combiner/adverb_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_node(parser, _uNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// tilde_combiner/adverb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// tilde_combiner
				if !_node(parser, _tilde_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// adverb_tone
				if !_node(parser, _adverb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ũFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ũ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ũ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ũ}
	// [ũŨ]/l:u (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ũŨ]
		if r, w := _next(parser, pos); r != 'ũ' && r != 'Ũ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ũŨ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:u (tilde_combiner/adverb_tone)
		// l:u
		{
			pos7 := pos
			// u
			if !_fail(parser, _uFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_fail(parser, _tilde_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_fail(parser, _adverb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ũAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ũ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ũ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ũŨ]/l:u (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ũŨ]
		if r, w := _next(parser, pos); r != 'ũ' && r != 'Ũ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:u (tilde_combiner/adverb_tone)
			// l:u
			{
				pos8 := pos
				// u
				if p, n := _uAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (tilde_combiner/adverb_tone)
			// tilde_combiner/adverb_tone
			{
				pos12 := pos
				// tilde_combiner
				if p, n := _tilde_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// adverb_tone
				if p, n := _adverb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['~'][l])
			}(
				start6, pos, label0)
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

func _ĩAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ĩ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ĩĨ]/l:i (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ĩĨ]
		if r, w := _next(parser, pos); r != 'ĩ' && r != 'Ĩ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (tilde_combiner/adverb_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_accept(parser, _iAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_accept(parser, _tilde_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_accept(parser, _adverb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ĩ, start, pos, perr)
fail:
	return _memoize(parser, _ĩ, start, -1, perr)
}

func _ĩNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ĩ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ĩ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ĩ"}
	// [ĩĨ]/l:i (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ĩĨ]
		if r, w := _next(parser, pos); r != 'ĩ' && r != 'Ĩ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:i (tilde_combiner/adverb_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_node(parser, _iNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// tilde_combiner/adverb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// tilde_combiner
				if !_node(parser, _tilde_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// adverb_tone
				if !_node(parser, _adverb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ĩFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ĩ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ĩ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ĩ}
	// [ĩĨ]/l:i (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ĩĨ]
		if r, w := _next(parser, pos); r != 'ĩ' && r != 'Ĩ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ĩĨ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:i (tilde_combiner/adverb_tone)
		// l:i
		{
			pos7 := pos
			// i
			if !_fail(parser, _iFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_fail(parser, _tilde_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_fail(parser, _adverb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ĩAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ĩ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ĩ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ĩĨ]/l:i (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ĩĨ]
		if r, w := _next(parser, pos); r != 'ĩ' && r != 'Ĩ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:i (tilde_combiner/adverb_tone)
			// l:i
			{
				pos8 := pos
				// i
				if p, n := _iAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (tilde_combiner/adverb_tone)
			// tilde_combiner/adverb_tone
			{
				pos12 := pos
				// tilde_combiner
				if p, n := _tilde_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// adverb_tone
				if p, n := _adverb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['~'][l])
			}(
				start6, pos, label0)
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

func _õAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _õ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [õÕ]/l:o (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [õÕ]
		if r, w := _next(parser, pos); r != 'õ' && r != 'Õ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (tilde_combiner/adverb_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_accept(parser, _oAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_accept(parser, _tilde_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_accept(parser, _adverb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _õ, start, pos, perr)
fail:
	return _memoize(parser, _õ, start, -1, perr)
}

func _õNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_õ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _õ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "õ"}
	// [õÕ]/l:o (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [õÕ]
		if r, w := _next(parser, pos); r != 'õ' && r != 'Õ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:o (tilde_combiner/adverb_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_node(parser, _oNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// tilde_combiner/adverb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// tilde_combiner
				if !_node(parser, _tilde_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// adverb_tone
				if !_node(parser, _adverb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _õFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _õ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "õ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _õ}
	// [õÕ]/l:o (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [õÕ]
		if r, w := _next(parser, pos); r != 'õ' && r != 'Õ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[õÕ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:o (tilde_combiner/adverb_tone)
		// l:o
		{
			pos7 := pos
			// o
			if !_fail(parser, _oFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_fail(parser, _tilde_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_fail(parser, _adverb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _õAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_õ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _õ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [õÕ]/l:o (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [õÕ]
		if r, w := _next(parser, pos); r != 'õ' && r != 'Õ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:o (tilde_combiner/adverb_tone)
			// l:o
			{
				pos8 := pos
				// o
				if p, n := _oAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (tilde_combiner/adverb_tone)
			// tilde_combiner/adverb_tone
			{
				pos12 := pos
				// tilde_combiner
				if p, n := _tilde_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// adverb_tone
				if p, n := _adverb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['~'][l])
			}(
				start6, pos, label0)
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

func _ẽAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	var labels [1]string
	use(labels)
	if dp, de, ok := _memo(parser, _ẽ, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [ẽẼ]/l:e (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ẽẼ]
		if r, w := _next(parser, pos); r != 'ẽ' && r != 'Ẽ' {
			perr = _max(perr, pos)
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (tilde_combiner/adverb_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_accept(parser, _eAccepts, &pos, &perr) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_accept(parser, _tilde_combinerAccepts, &pos, &perr) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_accept(parser, _adverb_toneAccepts, &pos, &perr) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
		}
		goto ok0
	fail5:
		pos = pos3
		goto fail
	ok0:
	}
	return _memoize(parser, _ẽ, start, pos, perr)
fail:
	return _memoize(parser, _ẽ, start, -1, perr)
}

func _ẽNode(parser *_Parser, start int) (int, *peg.Node) {
	var labels [1]string
	use(labels)
	dp := parser.deltaPos[start][_ẽ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ẽ}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "ẽ"}
	// [ẽẼ]/l:e (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		nkids1 := len(node.Kids)
		// [ẽẼ]
		if r, w := _next(parser, pos); r != 'ẽ' && r != 'Ẽ' {
			goto fail4
		} else {
			node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
			pos += w
		}
		goto ok0
	fail4:
		node.Kids = node.Kids[:nkids1]
		pos = pos3
		// action
		// l:e (tilde_combiner/adverb_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_node(parser, _eNode, node, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		{
			nkids8 := len(node.Kids)
			pos09 := pos
			// tilde_combiner/adverb_tone
			{
				pos13 := pos
				nkids11 := len(node.Kids)
				// tilde_combiner
				if !_node(parser, _tilde_combinerNode, node, &pos) {
					goto fail14
				}
				goto ok10
			fail14:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				// adverb_tone
				if !_node(parser, _adverb_toneNode, node, &pos) {
					goto fail15
				}
				goto ok10
			fail15:
				node.Kids = node.Kids[:nkids11]
				pos = pos13
				goto fail5
			ok10:
			}
			sub := _sub(parser, pos09, pos, node.Kids[nkids8:])
			node.Kids = append(node.Kids[:nkids8], sub)
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

func _ẽFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	var labels [1]string
	use(labels)
	pos, failure := _failMemo(parser, _ẽ, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "ẽ",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _ẽ}
	// [ẽẼ]/l:e (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		// [ẽẼ]
		if r, w := _next(parser, pos); r != 'ẽ' && r != 'Ẽ' {
			if pos >= errPos {
				failure.Kids = append(failure.Kids, &peg.Fail{
					Pos:  int(pos),
					Want: "[ẽẼ]",
				})
			}
			goto fail4
		} else {
			pos += w
		}
		goto ok0
	fail4:
		pos = pos3
		// action
		// l:e (tilde_combiner/adverb_tone)
		// l:e
		{
			pos7 := pos
			// e
			if !_fail(parser, _eFail, errPos, failure, &pos) {
				goto fail5
			}
			labels[0] = parser.text[pos7:pos]
		}
		// (tilde_combiner/adverb_tone)
		// tilde_combiner/adverb_tone
		{
			pos11 := pos
			// tilde_combiner
			if !_fail(parser, _tilde_combinerFail, errPos, failure, &pos) {
				goto fail12
			}
			goto ok8
		fail12:
			pos = pos11
			// adverb_tone
			if !_fail(parser, _adverb_toneFail, errPos, failure, &pos) {
				goto fail13
			}
			goto ok8
		fail13:
			pos = pos11
			goto fail5
		ok8:
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

func _ẽAction(parser *_Parser, start int) (int, *string) {
	var labels [1]string
	use(labels)
	var label0 string
	dp := parser.deltaPos[start][_ẽ]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _ẽ}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [ẽẼ]/l:e (tilde_combiner/adverb_tone) {…}
	{
		pos3 := pos
		var node2 string
		// [ẽẼ]
		if r, w := _next(parser, pos); r != 'ẽ' && r != 'Ẽ' {
			goto fail4
		} else {
			node = parser.text[pos : pos+w]
			pos += w
		}
		goto ok0
	fail4:
		node = node2
		pos = pos3
		// action
		{
			start6 := pos
			// l:e (tilde_combiner/adverb_tone)
			// l:e
			{
				pos8 := pos
				// e
				if p, n := _eAction(parser, pos); n == nil {
					goto fail5
				} else {
					label0 = *n
					pos = p
				}
				labels[0] = parser.text[pos8:pos]
			}
			// (tilde_combiner/adverb_tone)
			// tilde_combiner/adverb_tone
			{
				pos12 := pos
				// tilde_combiner
				if p, n := _tilde_combinerAction(parser, pos); n == nil {
					goto fail13
				} else {
					pos = p
				}
				goto ok9
			fail13:
				pos = pos12
				// adverb_tone
				if p, n := _adverb_toneAction(parser, pos); n == nil {
					goto fail14
				} else {
					pos = p
				}
				goto ok9
			fail14:
				pos = pos12
				goto fail5
			ok9:
			}
			node = func(
				start, end int, l string) string {
				return string(tone.Diacritic['~'][l])
			}(
				start6, pos, label0)
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

func _tilde_combinerAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _tilde_combiner, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// "\u0303"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̃" {
		perr = _max(perr, pos)
		goto fail
	}
	pos += 2
	perr = start
	return _memoize(parser, _tilde_combiner, start, pos, perr)
fail:
	return _memoize(parser, _tilde_combiner, start, -1, perr)
}

func _tilde_combinerNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_tilde_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _tilde_combiner}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "tilde_combiner"}
	// "\u0303"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̃" {
		goto fail
	}
	node.Kids = append(node.Kids, _leaf(parser, pos, pos+2))
	pos += 2
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _tilde_combinerFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _tilde_combiner, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "tilde_combiner",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _tilde_combiner}
	// "\u0303"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̃" {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "\"\\u0303\"",
			})
		}
		goto fail
	}
	pos += 2
	failure.Kids = nil
	parser.fail[key] = failure
	return pos, failure
fail:
	failure.Kids = nil
	failure.Want = "◌̃"
	parser.fail[key] = failure
	return -1, failure
}

func _tilde_combinerAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_tilde_combiner]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _tilde_combiner}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// "\u0303"
	if len(parser.text[pos:]) < 2 || parser.text[pos:pos+2] != "̃" {
		goto fail
	}
	node = parser.text[pos : pos+2]
	pos += 2
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _adverb_toneAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _adverb_tone, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [~7]
	if r, w := _next(parser, pos); r != '~' && r != '7' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _adverb_tone, start, pos, perr)
fail:
	return _memoize(parser, _adverb_tone, start, -1, perr)
}

func _adverb_toneNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_tone}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "adverb_tone"}
	// [~7]
	if r, w := _next(parser, pos); r != '~' && r != '7' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _adverb_toneFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _adverb_tone, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "adverb_tone",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _adverb_tone}
	// [~7]
	if r, w := _next(parser, pos); r != '~' && r != '7' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[~7]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _adverb_toneAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_adverb_tone]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _adverb_tone}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [~7]
	if r, w := _next(parser, pos); r != '~' && r != '7' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _aAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _a, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [aA]
	if r, w := _next(parser, pos); r != 'a' && r != 'A' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _a, start, pos, perr)
fail:
	return _memoize(parser, _a, start, -1, perr)
}

func _aNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_a]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _a}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "a"}
	// [aA]
	if r, w := _next(parser, pos); r != 'a' && r != 'A' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _aFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _a, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "a",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _a}
	// [aA]
	if r, w := _next(parser, pos); r != 'a' && r != 'A' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[aA]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _aAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_a]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _a}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [aA]
	if r, w := _next(parser, pos); r != 'a' && r != 'A' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _bAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _b, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [bB]
	if r, w := _next(parser, pos); r != 'b' && r != 'B' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _b, start, pos, perr)
fail:
	return _memoize(parser, _b, start, -1, perr)
}

func _bNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_b]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _b}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "b"}
	// [bB]
	if r, w := _next(parser, pos); r != 'b' && r != 'B' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _bFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _b, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "b",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _b}
	// [bB]
	if r, w := _next(parser, pos); r != 'b' && r != 'B' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[bB]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _bAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_b]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _b}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [bB]
	if r, w := _next(parser, pos); r != 'b' && r != 'B' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _cAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _c, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [cC]
	if r, w := _next(parser, pos); r != 'c' && r != 'C' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
	return _memoize(parser, _c, start, pos, perr)
fail:
	return _memoize(parser, _c, start, -1, perr)
}

func _cNode(parser *_Parser, start int) (int, *peg.Node) {
	dp := parser.deltaPos[start][_c]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _c}
	node := parser.node[key]
	if node != nil {
		return start + int(dp-1), node
	}
	pos := start
	node = &peg.Node{Name: "c"}
	// [cC]
	if r, w := _next(parser, pos); r != 'c' && r != 'C' {
		goto fail
	} else {
		node.Kids = append(node.Kids, _leaf(parser, pos, pos+w))
		pos += w
	}
	node.Text = parser.text[start:pos]
	parser.node[key] = node
	return pos, node
fail:
	return -1, nil
}

func _cFail(parser *_Parser, start, errPos int) (int, *peg.Fail) {
	pos, failure := _failMemo(parser, _c, start, errPos)
	if failure != nil {
		return pos, failure
	}
	failure = &peg.Fail{
		Name: "c",
		Pos:  int(start),
	}
	key := _key{start: start, rule: _c}
	// [cC]
	if r, w := _next(parser, pos); r != 'c' && r != 'C' {
		if pos >= errPos {
			failure.Kids = append(failure.Kids, &peg.Fail{
				Pos:  int(pos),
				Want: "[cC]",
			})
		}
		goto fail
	} else {
		pos += w
	}
	parser.fail[key] = failure
	return pos, failure
fail:
	parser.fail[key] = failure
	return -1, failure
}

func _cAction(parser *_Parser, start int) (int, *string) {
	dp := parser.deltaPos[start][_c]
	if dp < 0 {
		return -1, nil
	}
	key := _key{start: start, rule: _c}
	n := parser.act[key]
	if n != nil {
		n := n.(string)
		return start + int(dp-1), &n
	}
	var node string
	pos := start
	// [cC]
	if r, w := _next(parser, pos); r != 'c' && r != 'C' {
		goto fail
	} else {
		node = parser.text[pos : pos+w]
		pos += w
	}
	parser.act[key] = node
	return pos, &node
fail:
	return -1, nil
}

func _dAccepts(parser *_Parser, start int) (deltaPos, deltaErr int) {
	if dp, de, ok := _memo(parser, _d, start); ok {
		return dp, de
	}
	pos, perr := start, -1
	// [dD]
	if r, w := _next(parser, pos); r != 'd' && r != 'D' {
		perr = _max(perr, pos)
		goto fail
	} else {
		pos += w
	}
fail:
}

	if dp < 0 {
}