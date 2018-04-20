{
	package parser
	import (
		"github.com/eaburns/peggy/peg"
		"github.com/eaburns/toaq/ast"
		"github.com/eaburns/toaq/tone"
	)
}

full_text <- t:text EOF { return ast.Text(t) }
text <- s:_ d:discourse? EOF? { return ast.Text(text(s, d)) }

#
# Discourse
#

discourse <- (
		x0:sentence s0:_ { return ast.Node(x0.ModSentence(s0)) } /
		x1:fragment s1:_ { return ast.Node(x1.ModFragment(s1)) }
	)+

#
# Fragment, sentence, and statement
#

fragment <- f:(
	x0:afterthought_cop<relative_clause_1, relative_clause_1> { return ast.Fragment((*ast.CoPRelative)(x0)) } /
	x1:relative_clause  { return ast.Fragment(x1) } /
	x2:prenex  { return ast.Fragment(x2) } /
	x3:terms { return ast.Fragment(x3) }
	) s:_ { return ast.Fragment(f.ModFragment(s)) }

sentence <-
	x:afterthought_cop<sentence_1, sentence> { return ast.Sentence((*ast.CoPSentence)(x)) } /
	sentence_1
sentence_1 <-
	x:forethought_cop<sentence, sentence> { return ast.Sentence((*ast.CoPSentence)(x)) } /
	sentence_2
sentence_2 <-
	je:sentence_prefix_spaces? se:sentence_3 { se.JE = je; return ast.Sentence(se) }
sentence_prefix_spaces <-
	je:sentence_prefix s:_ { return ast.Word(*je.Mod(s)) }
sentence_3 <-
	st:statement s:_ da:illocutionary? { return &ast.StatementSentence{Statement: st.ModStatement(s), DA: da} }

statement <- p:prenex_spaces? s:statement_1 { return ast.Statement(prenexStmt(p, s)) }
statement_1 <-
	n:afterthought_cop<statement_2, statement> { return ast.Statement((*ast.CoPStatement)(n)) } /
	statement_2
statement_2 <-
	n:forethought_cop<statement, statement> { return ast.Statement((*ast.CoPStatement)(n)) } /
	statement_3
statement_3 <-
	p:predication s:_ na:end_statement? { return ast.Statement(endPredication(p, s, na)) }
predication <-
	p:predicate s:_ ts:terms? { return ast.Predication(predication(p, s, ts)) }
prenex_spaces <-
	 p:prenex s:_ { return ast.Prenex(*p.Mod(s)) }
prenex <-
	ts:terms s:_ bi:end_prenex { return (*ast.Prenex)(prenex(ts, s, bi)) }

#
# Predicates
#

predicate <- serial_predicate
predicate_1 <-
	x:afterthought_cop<predicate_2, predicate_1> { return ast.Predicate((*ast.CoPPredicate)(x)) } /
	predicate_2
predicate_2 <-
	forethought_cop_pred<predicate> /
	LU_predicate<verb_syllable> /
	MI_phrase<verb_syllable> /
	PO_phrase<verb_syllable> /
	MO_phrase<verb_syllable> /
	prefixed_predicate /
	predicate_3
prefixed_predicate <-
	x:prefix s:_ p:predicate_2 { return ast.Predicate(prefixed(x, s, p)) }
predicate_3 <-
	p:predicate_word<verb_syllable> s:_ { return ast.Predicate(p.ModPredicate(s)) }
serial_predicate <-
	serial<predicate_1> /
	predicate_1
serial<left> <-
	l:left s:_ r:serial_predicate { return ast.Predicate(serial(l, s, r)) }

#
# Terms
#

terms <-
	terms_2 /
	t:term { return &ast.Terms{Term: t} }
terms_2 <-
	t:term s:_ ts:terms { return &ast.Terms{Term: t.ModTerm(s), Terms: ts} }
term <-
	x0:linked_argument { return ast.Term(x0) } /
	x1:adverb { return ast.Term(x1) } /
	x2:termset { return ast.Term(x2) } /
	x3:prepositional_phrase { return ast.Term(x3) }
linked_argument <-
	g:linking_word_spaces? a:argument { return ast.Term(linkedTerm(g, a)) }
linking_word_spaces <-
	w:linking_word s:_ { return (*ast.Word)(w.Mod(s)) }
termset <-
	x0:forethought_cop<terms_V, terms_V> { return (*ast.TermSet)(x0) } /
	x1:forethought_cop<terms_IV, terms_IV> { return (*ast.TermSet)(x1) } /
	x2:forethought_cop<terms_III, terms_III> { return (*ast.TermSet)(x2) } /
	x3:forethought_cop<terms_II, terms_II> { return (*ast.TermSet)(x3) } /
	x4:forethought_cop<terms_V, terms_V> { return (*ast.TermSet)(x4) }
terms_II <-
	l:term s:_ r:term { return &ast.Terms{Term: l.ModTerm(s), Terms: &ast.Terms{Term: r}} }
terms_III <-
	l:term s:_ r:terms_II { return &ast.Terms{Term: l.ModTerm(s), Terms: r} }
terms_IV <-
	l:term s:_ r:terms_III { return &ast.Terms{Term: l.ModTerm(s), Terms: r} }
terms_V <-
	l:term s:_ r:terms_IV { return &ast.Terms{Term: l.ModTerm(s), Terms: r} }

#
# Arguments
#

argument <-
	x:afterthought_cop<arg_1, argument> { return ast.Argument((*ast.CoPArgument)(x)) } /
	arg_1
arg_1 <-
	x:forethought_cop<argument, argument> { return ast.Argument((*ast.CoPArgument)(x)) } /
	arg_2
arg_2 <-
	ku:focus_spaces? a:arg_3 { return ast.Argument(focusedArgument(ku, a)) }
focus_spaces <-
	ku:focus s:_ { return ast.Word(*ku.Mod(s)) }
arg_3 <-
	q:quantifier_spaces? a:arg_4 { return ast.PredicateArgument(quantifiedArgument(q, a)) }
quantifier_spaces <-
	q:quantifier s:_ { return ast.Word(*q.Mod(s)) }
arg_4 <-
	a:arg_5 s:_ rc:relative_clause? { return ast.PredicateArgument(argumentRelative(*a, s, rc)) }
arg_5 <-
	p:( serial<arg_6> / arg_6 ) { return &ast.PredicateArgument{Predicate: p} }
arg_6 <-
	x0:content_clause { return ast.Predicate(x0) } /
	afterthought_cop_pred<arg_7> /
	# NOTE: this differs from the official grammar, which uses argument instead of arg_5.
	forethought_cop_pred<arg_5> /
	LU_predicate<arg_syllable> /
	MI_phrase<arg_syllable> /
	PO_phrase<arg_syllable> /
	MO_phrase<arg_syllable> /
	arg_7
arg_7 <- predicate_word<arg_syllable>

#
# Relative clauses
#

relative_clause <-
	# BUG? should the right branch be relative_clause instead of relative_clause_1?
	x0:afterthought_cop<relative_clause_1, relative_clause_1> { return ast.Relative((*ast.CoPRelative)(x0)) } /
	x1:LU_phrase<relative_syllable> { return ast.Relative((*ast.LURelative)(x1)) } /
	relative_clause_1
relative_clause_1 <-
	x:forethought_cop<relative_clause, relative_clause> { return ast.Relative((*ast.CoPRelative)(x)) } /
	relative_clause_2
relative_clause_2 <-
	x0:afterthought_cop<relative_clause_3, statement> { return ast.Relative((*ast.CoPRelative)(x0)) } /
	x1:relative_clause_3 { return ast.Relative(x1) }
relative_clause_3 <-
	p:relative_predication s:_ na:end_statement? { return &ast.PredicationRelative{Predication: *endPredication(p, s, na)} }
relative_predication <-
	p:relative_predicate s:_ ts:terms? { return ast.Predication(predication(p, s, ts))  }

relative_predicate <-
	serial<relative_predicate_1> / relative_predicate_1
relative_predicate_1 <-
	afterthought_cop_pred<relative_predicate_2> /
	forethought_cop_pred<relative_predicate> /
	relative_predicate_2
relative_predicate_2 <-
	MI_phrase<relative_syllable> /
	PO_phrase<relative_syllable> /
	MO_phrase<relative_syllable> /
	predicate_word<relative_syllable>

#
# Adverbs
#

adverb <-
	x:afterthought_cop<adverb_1, adverb> { return ast.Adverb((*ast.CoPAdverb)(x)) } /
	adverb_1
adverb_1 <-
	x:forethought_cop<adverb, adverb> { return ast.Adverb((*ast.CoPAdverb)(x)) } /
	adverb_2
adverb_2 <-
	p:( serial<adverb_3> / adverb_3 ) { return ast.Adverb(&ast.PredicateAdverb{Predicate: p}) }
adverb_3 <-
	afterthought_cop_pred<adverb_4> /
	# NOTE: this differs from the original grammar, which uses adverb instead of adverb_3.
	forethought_cop_pred<adverb_3> /
	LU_predicate<adverb_syllable> /
	MI_phrase<adverb_syllable> /
	PO_phrase<adverb_syllable> /
	MO_phrase<adverb_syllable> /
	adverb_4
adverb_4 <- predicate_word<adverb_syllable>

#
# Prepositional phrases
#

prepositional_phrase <-
	x:afterthought_cop<prepositional_phrase_1, prepositional_phrase_1> { return ast.Preposition((*ast.CoPPreposition)(x)) } /
	prepositional_phrase_1
prepositional_phrase_1 <-
	x:forethought_cop<prepositional_phrase, prepositional_phrase> { return ast.Preposition((*ast.CoPPreposition)(x)) } /
	prepositional_phrase_2
prepositional_phrase_2 <-
	p:preposition s:_ a:argument { return ast.Preposition(prepPhrase(p, s, a)) }

preposition <-
	x:afterthought_cop<preposition_1, preposition> { return ast.Predicate((*ast.CoPPredicate)(x)) } /
	preposition_1
preposition_1 <-
	x:forethought_cop<preposition, preposition> { return ast.Predicate((*ast.CoPPredicate)(x)) } /
	preposition_2
preposition_2 <-
	serial<preposition_3> / preposition_3
preposition_3 <-
	afterthought_cop_pred<preposition_4> /
	forethought_cop_pred<preposition> /
	LU_predicate<preposition_syllable> /
	MI_phrase<preposition_syllable> /
	PO_phrase<preposition_syllable> /
	MO_phrase<preposition_syllable> /
	preposition_4
preposition_4 <- predicate_word<preposition_syllable>

#
# Content clauses
#

content_clause <-
	x:afterthought_cop<content_clause_1, statement> { return ast.Content((*ast.CoPContent)(x)) } /
	content_clause_1
content_clause_1 <-
	x0:content_statement { return ast.Content(x0) } /
	x1:LU_phrase<content_syllable> { return ast.Content((*ast.LUContent)(x1)) }
content_statement <-
	p:content_predication s:_ na:end_statement? { return &ast.PredicationContent{Predication: *endPredication(p, s, na)} }
content_predication <-
	p:content_predicate s:_ ts:terms? { return ast.Predication(predication(p, s, ts)) }

content_predicate <-
	serial<content_predicate_1> / content_predicate_1
content_predicate_1 <-
	afterthought_cop_pred<content_predicate_2> /
	forethought_cop_pred<content_predicate_2> /
	content_predicate_2
content_predicate_2 <-
	MI_phrase<content_syllable> /
	PO_phrase<content_syllable> /
	MO_phrase<content_syllable> /
	predicate_word<content_syllable>

#
# CoP
#

afterthought_cop_pred<left> <-
	x:afterthought_cop<left, predicate> { return ast.Predicate((*ast.CoPPredicate)(x)) }
forethought_cop_pred<left> <-
	x:forethought_cop<left, predicate> { return ast.Predicate((*ast.CoPPredicate)(x)) }

afterthought_cop<left, right> <-
	l:left s:_ bar:cop_bar<right> { return (*ast.CoP)(copLeft(l, s, bar)) }
cop_bar<right> <-
	ru:connective s:_ r:right { return ast.CoP(copRight(ru, s, r)) }

forethought_cop<left, right> <-
	con:forethought_connective s:_ cop:forethought_cop_1<left, right> { return (*ast.CoP)(foreCoP(con, s, cop)) }
forethought_cop_1<left, right> <-
	l:left s:_ bar:forethought_cop_bar<right> { return ast.CoP(foreCoPLeft(l, s, bar)) }
forethought_cop_bar<right> <-
	to:forethought_marker s:_ r:right { return ast.CoP(foreCoPRight(to, s, r)) }
forethought_connective <-
	to:forethought_marker s:_ ru:connective { return ast.CoP(connector(to, s, ru)) }

#
# LU, MI, PO, and MO phrases
#

LU_predicate<syllable> <-
	x:LU_phrase<syllable> { return ast.Predicate((*ast.LUPredicate)(x)) }
LU_phrase<syllable> <-
	lu:LU_word<syllable> s:_ st:statement { return &ast.LUPhrase{LU: *lu.Mod(s), Statement: st} }
LU_word<syllable> <-
	&LU w:syllable { return &ast.Word{S: start, E: end, T: w} }
	

MI_phrase<syllable> <-
	mi:MI_phrase_1<syllable> s:_ ga:end_predicatizer? { return ast.Predicate(miEnd(mi, s, ga)) }
MI_phrase_1<syllable> <-
	mi:MI_word<syllable> s:_ p:MI_complement { return ast.MIPredicate{MI: *mi.Mod(s), Phrase: p} }
MI_word<syllable> <-
	&MI w:syllable { return &ast.Word{S: start, E: end, T: w} }
MI_complement <-
	x0:predicate { return ast.Phrase(x0) } /
	x1:argument { return ast.Phrase(x1) } /
	x2:adverb { return ast.Phrase(x2) }/
	x3:prepositional_phrase { return ast.Phrase(x3) }

PO_phrase<syllable> <-
	po:PO_phrase_1<syllable> s:_ ga:end_predicatizer? { return ast.Predicate(poEnd(po, s, ga)) }
PO_phrase_1<syllable> <-
	po:PO_word<syllable> s:_ a:argument { return ast.POPredicate{PO: *po.Mod(s), Argument: a} }
PO_word<syllable> <-
	&PO w:syllable { return &ast.Word{S: start, E: end, T: w} }

MO_phrase<syllable> <-
	mo:MO_phrase_1<syllable> s:_ teo:end_quote { return ast.Predicate(moEnd(mo, s, teo)) }
MO_phrase_1<syllable> <-
	mo:MO_word<syllable> s:_ d:discourse { return ast.MOPredicate{MO: *mo.Mod(s), Discourse: d} }
MO_word<syllable> <-
	&MO w:syllable { return &ast.Word{S: start, E: end, T: w} }

#
# Space and free modifiers
#

_ <- space_or_freemod?
space_or_freemod <-
	m:( freemod / space_word ) ms:space_or_freemod? { return ast.Mod(m.Mod(ms)) }
space_word <-
	s:space+ { return ast.Mod(&ast.Space{S: start, E: end, T: s}) }
freemod <-
	x0:interjection { return ast.Mod(x0) } /
	x1:parenthetical { return ast.Mod(x1) }/
	x2:incidental { return ast.Mod(x2) } /
	x3:vocative { return ast.Mod(x3) }
parenthetical <-
	p:parenthetical_1 s:_ ki:end_parenthetical { return (*ast.Parenthetical)(endParen(p, s, ki)) }
parenthetical_1 <-
	kio:start_parenthetical s:_ d:discourse { return ast.Parenthetical{KIO: *kio.Mod(s), Discourse: d} }
incidental <-
	ju:start_incidental s:_ st:statement { return &ast.Incidental{JU: *ju.Mod(s), Statement: st} }
vocative <-
	hu:vocative_marker s:_ a:argument { return &ast.Vocative{HU: *hu.Mod(s), Argument: a} }

#
# Morphology
#

word <- w:(
		non_function_word /
		toned_function_word /
		neutral_function_word
	) s:space_word? { return (*ast.Word)(ast.Word{S: start, E: end, T: w}.Mod(s)) }
predicate_word<syllable> <- w:(
		syllable compound_syllable+ /
		!function_word syllable
	){ return ast.Predicate(&ast.WordPredicate{S: start, E: end, T: w}) }
non_function_word <-
	initial_syllable compound_syllable+ / !function_word initial_syllable
neutral_function_word <-
	!toned_function_word &function_word neutral_syllable
toned_function_word <-
	!neutral_syllable ( LU / MI / MO / PO )
initial_syllable <-
	arg_syllable /
	relative_syllable /
	verb_syllable /
	content_syllable /
	preposition_syllable /
	adverb_syllable

prefix <- &MU w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
focus <- &KU w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
end_quote <- &TEO w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
end_predicatizer <- &GA w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
end_statement <- &NA w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
sentence_prefix <- &KEO w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
end_prenex <- &BI w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
start_incidental <- &JU w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
start_parenthetical <- &KIO w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
end_parenthetical <- &KI w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
vocative_marker <- &HU w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
linking_word <- &GO w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
connective <- &RA w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
illocutionary <- &DA w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
quantifier <- &TU w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }
interjection <- &HA w:neutral_syllable { return &ast.Interjection{S: start, E: end, T: w} }
forethought_marker <- &TO w:neutral_syllable { return ast.Word{S: start, E: end, T: w} }

# The following function word rules match any tone.
# The intended use is as the subexpression of a ! or &,
# followed by the desired syllable tone.
function_word <- BI / DA / GA / GO / HA / HU / JU / KU / KI / KIO / KEO / LU / MU / MI / MO / NA / PO / RA / TU / TO / TEO
BI <- ( b I / p A ) &( tone? boundary )
DA <- ( d A / m O q / k A / s O / b A ) &( tone? boundary )
GA <- g A &( tone? boundary )
GO <- ( g O / f I / c U / k E ) &( tone? boundary )
HA <- ( h A / h U e / h I a / m ) &( tone? boundary )
HU <- h U &( tone? boundary )
JU <- j U &( tone? boundary )
KU <- k U &( tone? boundary )
KI <- k I &( tone? boundary )
KIO <- k I o &( tone? boundary )
KEO <- ( k E o / j E ) &( tone? boundary )
LU <- ( l U / l I / l O / m A / t I o ) &( tone? boundary )
MU <- m U &( tone? boundary )
MI <- ( m I / s h U ) &( tone? boundary )
MO <- m O &( tone? boundary )
NA <- n A &( tone? boundary )
PO <- ( p O / p E ) &( tone? boundary )
RA <- ( r A / r U / r I / r O i / r O / r E ) &( tone? boundary )
TU <- ( t A / t U / s I a / s A / h I / j A ) &( tone? boundary )
TO <- t O &boundary
TEO <- t E o &( tone? boundary )

syllable<desinence, tone> <- initial ( desinence / d:neutral_desinence t:tone { return string(addTone(d, t)) } ) &boundary
neutral_syllable <- initial neutral_desinence &boundary
compound_syllable <- syllable<compound_desinence, compound_tone>
arg_syllable <- syllable<arg_desinence, arg_tone>
relative_syllable <- syllable<relative_desinence, relative_tone>
verb_syllable <- syllable<verb_desinence, verb_tone>
content_syllable <- syllable<content_desinence, content_tone>
preposition_syllable <- syllable<preposition_desinence, preposition_tone>
adverb_syllable <- syllable<adverb_desinence, adverb_tone>

boundary <- ( initial / space / EOF ) { return "" } # Return string just for type checking.

initial <- b / c h / c / d / f / g / h / j / k / l / m / n / p / r / s h / s / t

desinence<ā,ū,ī,ō,ē> <-
	ā o / ā i / ā q / ā /
	ū a o / ū a i / ū e q / ū e / ū o q / ū o / ū i / ū a q / ū a / ū q / ū /
	ī a o / ī a i / ī e q / ī e / ī o q / ī o / ī u / ī a q / ī a / ī q / ī /
	ō e q / ō e / ō i / ō a q / ō a / ō q / ō /
	ē o q / ē o / ē i / ē a q / ē a / ē q / ē
neutral_desinence <- desinence<a,u,i,o,e>
compound_desinence <- desinence<ā,ū,ī,ō,ē>
arg_desinence <- desinence<á,ú,í,ó,é>
relative_desinence <- desinence<ǎ,ǔ,ǐ,ǒ,ě>
verb_desinence <- desinence<ả,ủ,ỉ,ỏ,ẻ>
content_desinence <- desinence<â,û,î,ô,ê>
preposition_desinence <- desinence<à,ù,ì,ò,è>
adverb_desinence <- desinence<ã,ũ,ĩ,õ,ẽ>

tone <- compound_tone / arg_tone / verb_tone / relative_tone / adverb_tone / preposition_tone / content_tone

A <- ā / á / ǎ / ả / â / à / ã / a
U <- ū / ú / ǔ / ủ / û / ù / ũ / u
I <- ī / í / ǐ / ỉ / î / ì / ĩ / i
O <- ō / ó / ǒ / ỏ / ô / ò / õ / o
E <- ē / é / ě / ẻ / ê / è / ẽ / e

ā <- [āĀ] / l:a ( macron_combiner / compound_tone ) { return string(tone.Diacritic['-'][l]) }
ū <- [ūŪ] / l:u ( macron_combiner / compound_tone ) { return string(tone.Diacritic['-'][l]) }
ī <- [īĪ] / l:i ( macron_combiner / compound_tone ) { return string(tone.Diacritic['-'][l]) }
ō <- [ōŌ] / l:o ( macron_combiner / compound_tone ) { return string(tone.Diacritic['-'][l]) }
ē <- [ēĒ] / l:e ( macron_combiner / compound_tone ) { return string(tone.Diacritic['-'][l]) }
macron_combiner "\u25CC\u0304" <- "\u0304"
compound_tone <- [\-1]

á <- [áÁ] / l:a ( acute_combiner / arg_tone ) { return string(tone.Diacritic['/'][l]) }
ú <- [úÚ] / l:u ( acute_combiner / arg_tone ) { return string(tone.Diacritic['/'][l]) }
í <- [íÍ] / l:i ( acute_combiner / arg_tone ) { return string(tone.Diacritic['/'][l]) }
ó <- [óÓ] / l:o ( acute_combiner / arg_tone ) { return string(tone.Diacritic['/'][l]) }
é <- [éÉ] / l:e ( acute_combiner / arg_tone ) { return string(tone.Diacritic['/'][l]) }
acute_combiner "\u25CC\u0301" <- "\u0301"
arg_tone <- [/2]

# This also includes the combining breve forms ( \u0306 ),
# because la jelca is too lazy to fix his keyboard.
ǎ <- [ǎǍ] / l:a ( caron_combiner / breve_combiner / relative_tone ) { return string(tone.Diacritic['V'][l]) }
ǔ <- [ǔǓ] / l:u ( caron_combiner / breve_combiner / relative_tone ) { return string(tone.Diacritic['V'][l]) }
ǐ <- [ǐǏ] / l:i ( caron_combiner / breve_combiner / relative_tone ) { return string(tone.Diacritic['V'][l]) }
ǒ <- [ǒǑ] / l:o ( caron_combiner / breve_combiner / relative_tone ) { return string(tone.Diacritic['V'][l]) }
ě <- [ěĚ] / l:e ( caron_combiner / breve_combiner / relative_tone ) { return string(tone.Diacritic['V'][l]) }
caron_combiner "\u25CC\u030C" <- "\u030C"
breve_combiner "\u25CC\u0306" <- "\u0306"
relative_tone <- [V3]

ả <- [ảẢ] / l:a ( hook_combiner / verb_tone ) { return string(tone.Diacritic['?'][l]) }
ủ <- [ủỦ] / l:u ( hook_combiner / verb_tone ) { return string(tone.Diacritic['?'][l]) }
ỉ <- [ỉỈ] / l:i ( hook_combiner / verb_tone ) { return string(tone.Diacritic['?'][l]) }
ỏ <- [ỏỎ] / l:o ( hook_combiner / verb_tone ) { return string(tone.Diacritic['?'][l]) }
ẻ <- [ẻẺ] / l:e ( hook_combiner / verb_tone ) { return string(tone.Diacritic['?'][l]) }
hook_combiner "\u25CC\u0309" <- "\u0309"
verb_tone <- [?4]

â <- [âÂ] / l:a ( circumflex_combiner / content_tone ) { return string(tone.Diacritic['^'][l]) }
û <- [ûÛ] / l:u ( circumflex_combiner / content_tone ) { return string(tone.Diacritic['^'][l]) }
î <- [îÎ] / l:i ( circumflex_combiner / content_tone ) { return string(tone.Diacritic['^'][l]) }
ô <- [ôÔ] / l:o ( circumflex_combiner / content_tone ) { return string(tone.Diacritic['^'][l]) }
ê <- [êÊ] / l:e ( circumflex_combiner / content_tone ) { return string(tone.Diacritic['^'][l]) }
circumflex_combiner "\u25CC\u0302" <- "\u0302"
content_tone <- [5\^]

à <- [àÀ] / l:a ( grave_combiner / preposition_tone ) { return string(tone.Diacritic['\\'][l]) }
ù <- [ùÙ] / l:u ( grave_combiner / preposition_tone ) { return string(tone.Diacritic['\\'][l]) }
ì <- [ìÌ] / l:i ( grave_combiner / preposition_tone ) { return string(tone.Diacritic['\\'][l]) }
ò <- [òÒ] / l:o ( grave_combiner / preposition_tone ) { return string(tone.Diacritic['\\'][l]) }
è <- [èÈ] / l:e ( grave_combiner / preposition_tone ) { return string(tone.Diacritic['\\'][l]) }
grave_combiner "\u25CC\u0300" <- "\u0300"
preposition_tone <- [\\6]

ã <- [ãÃ] / l:a ( tilde_combiner / adverb_tone ) { return string(tone.Diacritic['~'][l]) }
ũ <- [ũŨ] / l:u ( tilde_combiner / adverb_tone ) { return string(tone.Diacritic['~'][l]) }
ĩ <- [ĩĨ] / l:i ( tilde_combiner / adverb_tone ) { return string(tone.Diacritic['~'][l]) }
õ <- [õÕ] / l:o ( tilde_combiner / adverb_tone ) { return string(tone.Diacritic['~'][l]) }
ẽ <- [ẽẼ] / l:e ( tilde_combiner / adverb_tone ) { return string(tone.Diacritic['~'][l]) }
tilde_combiner "\u25CC\u0303" <- "\u0303"
adverb_tone <- [~7]

a <- [aA]
b <- [bB]
c <- [cC]
d <- [dD]
e <- [eE]
f <- [fF]
g <- [gG]
h <- [hH]
i <- l:[ıiI] { if l == "i" { return "ı" }; return string(l) }
j <- [jJ]
k <- [kK]
l <- [lL]
m <- [mM]
n <- [nN]
o <- [oO]
p <- [pP]
q <- [qQ]
r <- [rR]
s <- [sS]
t <- [tT]
u <- [uU]
w <- [wW]
y <- [yY]

space "space" <- s:[\t\n\r\u0020.,:!]+
EOF "EOF" <- !.
