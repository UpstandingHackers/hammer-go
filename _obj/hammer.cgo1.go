// Created by cgo - DO NOT EDIT

//line /Users/cvincent/go/src/hammer/hammer.go:1
package hammer
//line /Users/cvincent/go/src/hammer/hammer.go:14

//line /Users/cvincent/go/src/hammer/hammer.go:13
type HParsedToken *[0]byte
type HParseResult *[0]byte
type HAction func(*[0]byte,) *[0]byte
type HPredicate func(*[0]byte,) *[0]byte
type HParser *[0]byte
type HCaseResult *[0]byte
type File *_Ctype_FILE
type Void uintptr
//line /Users/cvincent/go/src/hammer/hammer.go:24

//line /Users/cvincent/go/src/hammer/hammer.go:23
func Parse(parser HParser, input *uint8, length uintptr) ParseResult {
//line /Users/cvincent/go/src/hammer/hammer.go:23
	return _Cfunc_h_parse(parser, input, length)
//line /Users/cvincent/go/src/hammer/hammer.go:23
}
//line /Users/cvincent/go/src/hammer/hammer.go:27

//line /Users/cvincent/go/src/hammer/hammer.go:26
func Token(str *uint8, length uintptr) HParser {
	var cStr string = _Cfunc_CString(str)
	return _Cfunc_h_token(cStr, length)
}
//line /Users/cvincent/go/src/hammer/hammer.go:33

//line /Users/cvincent/go/src/hammer/hammer.go:32
func Ch(c uint8) HParser	{ return _Cfunc_h_ch(c) }
//line /Users/cvincent/go/src/hammer/hammer.go:36

//line /Users/cvincent/go/src/hammer/hammer.go:35
func Ch_range(lower uint8, upper uint8) HParser	{ return _Cfunc_h_ch_range(lower, upper) }
//line /Users/cvincent/go/src/hammer/hammer.go:39

//line /Users/cvincent/go/src/hammer/hammer.go:38
func Int_range(p HParser, lower int64, upper int64) HParser {
	return _Cfunc_h_int_range(p, lower, upper)
}
//line /Users/cvincent/go/src/hammer/hammer.go:44

//line /Users/cvincent/go/src/hammer/hammer.go:43
func Bits(length uintptr, sign bool) HParser	{ return _Cfunc_h_bits(length, sign) }
//line /Users/cvincent/go/src/hammer/hammer.go:47

//line /Users/cvincent/go/src/hammer/hammer.go:46
func Int64() HParser	{ return _Cfunc_h_int64() }
//line /Users/cvincent/go/src/hammer/hammer.go:50

//line /Users/cvincent/go/src/hammer/hammer.go:49
func Int32() HParser	{ return _Cfunc_h_int32() }
//line /Users/cvincent/go/src/hammer/hammer.go:53

//line /Users/cvincent/go/src/hammer/hammer.go:52
func Int16() HParser	{ return _Cfunc_h_int16() }
//line /Users/cvincent/go/src/hammer/hammer.go:56

//line /Users/cvincent/go/src/hammer/hammer.go:55
func Int8() HParser	{ return _Cfunc_h_int8() }
//line /Users/cvincent/go/src/hammer/hammer.go:59

//line /Users/cvincent/go/src/hammer/hammer.go:58
func Uint64() HParser	{ return _Cfunc_h_uint64() }
//line /Users/cvincent/go/src/hammer/hammer.go:62

//line /Users/cvincent/go/src/hammer/hammer.go:61
func Uint32() HParser	{ return _Cfunc_h_uint32() }
//line /Users/cvincent/go/src/hammer/hammer.go:65

//line /Users/cvincent/go/src/hammer/hammer.go:64
func Uint16() HParser	{ return _Cfunc_h_uint16() }
//line /Users/cvincent/go/src/hammer/hammer.go:68

//line /Users/cvincent/go/src/hammer/hammer.go:67
func Uint8() HParser	{ return _Cfunc_h_uint8() }
//line /Users/cvincent/go/src/hammer/hammer.go:71

//line /Users/cvincent/go/src/hammer/hammer.go:70
func Whitespace(p HParser) HParser	{ return _Cfunc_h_whitespace(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:74

//line /Users/cvincent/go/src/hammer/hammer.go:73
func Left(p HParser, q HParser) HParser	{ return _Cfunc_h_left(p, q) }
//line /Users/cvincent/go/src/hammer/hammer.go:77

//line /Users/cvincent/go/src/hammer/hammer.go:76
func Right(p HParser, q HParser) HParser	{ return _Cfunc_h_right(p, q) }
//line /Users/cvincent/go/src/hammer/hammer.go:80

//line /Users/cvincent/go/src/hammer/hammer.go:79
func Middle(p HParser, x HParser, q HParser) HParser	{ return _Cfunc_h_middle(p, x, q) }
//line /Users/cvincent/go/src/hammer/hammer.go:83

//line /Users/cvincent/go/src/hammer/hammer.go:82
func Action(p HParser, q HAction)	{ return _Cfunc_h_action(p, a) }
//line /Users/cvincent/go/src/hammer/hammer.go:86

//line /Users/cvincent/go/src/hammer/hammer.go:85
func In(charset *uint8, length uintptr) HParser	{ return _Cfunc_h_in(charset, length) }
//line /Users/cvincent/go/src/hammer/hammer.go:89

//line /Users/cvincent/go/src/hammer/hammer.go:88
func Not_in(charset *uint8, length uintptr) HParser	{ _Cfunc_h_not_in(charser, length) }
//line /Users/cvincent/go/src/hammer/hammer.go:92

//line /Users/cvincent/go/src/hammer/hammer.go:91
func End_p()	{ return _Cfunc_h_end_p() }
//line /Users/cvincent/go/src/hammer/hammer.go:95

//line /Users/cvincent/go/src/hammer/hammer.go:94
func Nothing_p()	{ return _Cfunc_h_nothing_p() }
//line /Users/cvincent/go/src/hammer/hammer.go:98

//line /Users/cvincent/go/src/hammer/hammer.go:97
func Sequence(p HParser) HParser	{ return _Cfunc_h_sequence(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:101

//line /Users/cvincent/go/src/hammer/hammer.go:100
func Choice(p HParser) HParser	{ return _Cfunc_h_choice(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:104

//line /Users/cvincent/go/src/hammer/hammer.go:103
func Butnot(p1 HParser, p2 HParser) HParser	{ return _Cfunc_h_butnot(p1, p2) }
//line /Users/cvincent/go/src/hammer/hammer.go:107

//line /Users/cvincent/go/src/hammer/hammer.go:106
func Difference(p1 HParser, p2 HParser) HParser	{ return _Cfunc_h_difference(p1, p2) }
//line /Users/cvincent/go/src/hammer/hammer.go:110

//line /Users/cvincent/go/src/hammer/hammer.go:109
func Xor() HParser	{ return _Cfunc_h_xor() }
//line /Users/cvincent/go/src/hammer/hammer.go:113

//line /Users/cvincent/go/src/hammer/hammer.go:112
func Many(p HParser) HParser	{ return _Cfunc_h_many(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:116

//line /Users/cvincent/go/src/hammer/hammer.go:115
func Many1(p HParser) HParser	{ return _Cfunc_h_many1(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:119

//line /Users/cvincent/go/src/hammer/hammer.go:118
func Repeat_n(p HParser, n uintptr) HParser	{ return _Cfunc_h_repeat_n(p, n) }
//line /Users/cvincent/go/src/hammer/hammer.go:122

//line /Users/cvincent/go/src/hammer/hammer.go:121
func Optional(p HParser) HParser	{ return _Cfunc_h_optional(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:125

//line /Users/cvincent/go/src/hammer/hammer.go:124
func Ignore(p HParser) HParser	{ return _Cfunc_h_ignore(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:128

//line /Users/cvincent/go/src/hammer/hammer.go:127
func SepBy(p HParser, sep HParser) HParser	{ return _Cfunc_h_sepBy(p, sep) }
//line /Users/cvincent/go/src/hammer/hammer.go:131

//line /Users/cvincent/go/src/hammer/hammer.go:130
func SepBy1(p HParser, sep HParser) HParser	{ return _Cfunc_h_sepBy1(p, sep) }
//line /Users/cvincent/go/src/hammer/hammer.go:134

//line /Users/cvincent/go/src/hammer/hammer.go:133
func Epsilon_p() HParser	{ return _Cfunc_h_epsilon_p() }
//line /Users/cvincent/go/src/hammer/hammer.go:137

//line /Users/cvincent/go/src/hammer/hammer.go:136
func Length_value(h_length HParser, length HParser, value HParser) HParser {
//line /Users/cvincent/go/src/hammer/hammer.go:136
	return _Cfunc_h_length_value(h_length, length, value)
//line /Users/cvincent/go/src/hammer/hammer.go:136
}
//line /Users/cvincent/go/src/hammer/hammer.go:140

//line /Users/cvincent/go/src/hammer/hammer.go:139
func Attr_bool(p HParser, pred HPredicate) HParser	{ return _Cfunc_h_attr_bool(p, pred) }
//line /Users/cvincent/go/src/hammer/hammer.go:143

//line /Users/cvincent/go/src/hammer/hammer.go:142
func And(p HParser) HParser	{ return _Cfunc_h_and(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:146

//line /Users/cvincent/go/src/hammer/hammer.go:145
func Not(p HParser) HParser	{ return _Cfunc_h_not(p) }
//line /Users/cvincent/go/src/hammer/hammer.go:149

//line /Users/cvincent/go/src/hammer/hammer.go:148
func Indirect() HParser	{ return _Cfunc_h_indirect() }
//line /Users/cvincent/go/src/hammer/hammer.go:152

//line /Users/cvincent/go/src/hammer/hammer.go:151
func Bind_indirect(indirect HParser, inner HParser) HParser	{ _Cfunc_h_bind_indirect(indirect, inner) }
//line /Users/cvincent/go/src/hammer/hammer.go:155

//line /Users/cvincent/go/src/hammer/hammer.go:154
func Parse_result_free(result HParseResult)	{ _Cfunc_h_parse_result_free(result) }
//line /Users/cvincent/go/src/hammer/hammer.go:158

//line /Users/cvincent/go/src/hammer/hammer.go:157
func Write_result_unamb(tok HParsedToken) string {
	return _Cfunc_GoString(_Cfunc_h_write_result_unamb(tok))
}
//line /Users/cvincent/go/src/hammer/hammer.go:163

//line /Users/cvincent/go/src/hammer/hammer.go:162
func Pprint(stream File, tok HParsedToken, indent int, delta int) {
//line /Users/cvincent/go/src/hammer/hammer.go:162
	_Cfunc_h_pprint(stream, tok, indent, delta)
//line /Users/cvincent/go/src/hammer/hammer.go:162
}
//line /Users/cvincent/go/src/hammer/hammer.go:166

//line /Users/cvincent/go/src/hammer/hammer.go:165
func Compile(parser HParser, backend HParserBackend, params Void) int {
//line /Users/cvincent/go/src/hammer/hammer.go:165
	return _Cfunc_h_compile(parser, backend, params)
//line /Users/cvincent/go/src/hammer/hammer.go:165
}
