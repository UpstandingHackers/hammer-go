package hammer

import (
	/*
		#cgo CFLAGS: -Ihammer/src
		#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
		#include <hammer.h>
	*/
	"C"
	"unsafe"
)

/* There is no function in hammer that passes by value except
the input function/. Because of this, by defining go
"types" as the pointers needed as arguments, we satisfy
the requirements of the lib without having to instanciate
any of the hammer.structs in the go extension. \o/
[thnx qux!] */
type HParsedToken *C.HParsedToken
type HParseResult *C.HParseResult
type HAction C.HAction
type HPredicate C.HPredicate
type HParser *C.HParser
type HParserBackend C.HParserBackend
type File *C.FILE
type Void uintptr

// HAMMER_FN_DECL(HParseResult*, h_parse,  HParser* parser,  uint8_t* input, size_t length);
func Parse(parser HParser, input *uint8, length uintptr) HParseResult {
	return C.h_parse(parser, (*C.uint8_t)(input), (C.size_t)(length))
}

// HAMMER_FN_DECL(HParser*, h_token,  uint8_t *str,  size_t len);
func Token(str *uint8, length uintptr) HParser {
	return C.h_token((*C.uint8_t)(str), (C.size_t)(length))
}

// HAMMER_FN_DECL(HParser*, h_ch,  uint8_t c);
func Ch(c uint8) HParser { return C.h_ch((C.uint8_t)(c)) }

// HAMMER_FN_DECL(HParser*, h_ch_range,  uint8_t lower,  uint8_t upper);
func Ch_range(lower uint8, upper uint8) HParser {
	return C.h_ch_range((C.uint8_t)(lower), (C.uint8_t)(upper))
}

// HAMMER_FN_DECL(HParser*, h_int_range,  HParser *p,  int64_t lower,  int64_t upper);
func Int_range(p HParser, lower int64, upper int64) HParser {
	return C.h_int_range(p, (C.int64_t)(lower), (C.int64_t)(upper))
}

// HAMMER_FN_DECL(HParser*, h_bits, size_t len, bool sign);
func Bits(len uintptr, sign bool) HParser {
	var retVal C.bool

	// bool to C.bool conversion
	if sign {
		retVal = 1
	} else if sign == false {
		retVal = 0
	}
	return C.h_bits((C.size_t)(len), retVal)
}

// HAMMER_FN_DECL_NOARG(HParser*, h_int64);
func Int64() HParser { return C.h_int64() }

// HAMMER_FN_DECL_NOARG(HParser*, h_int32);
func Int32() HParser { return C.h_int32() }

// HAMMER_FN_DECL_NOARG(HParser*, h_int16);
func Int16() HParser { return C.h_int16() }

// HAMMER_FN_DECL_NOARG(HParser*, h_int8);
func Int8() HParser { return C.h_int8() }

// HAMMER_FN_DECL_NOARG(HParser*, h_uint64);
func Uint64() HParser { return C.h_uint64() }

// HAMMER_FN_DECL_NOARG(HParser*, h_uint32);
func Uint32() HParser { return C.h_uint32() }

// HAMMER_FN_DECL_NOARG(HParser*, h_uint16);
func Uint16() HParser { return C.h_uint16() }

// HAMMER_FN_DECL_NOARG(HParser*, h_uint8);
func Uint8() HParser { return C.h_uint8() }

// HAMMER_FN_DECL(HParser*, h_whitespace, const HParser* p);
func Whitespace(p HParser) HParser { return C.h_whitespace(p) }

//HAMMER_FN_DECL(HParser*, h_left,  HParser* p,  HParser* q);
func Left(p HParser, q HParser) HParser { return C.h_left(p, q) }

//HAMMER_FN_DECL(HParser*, h_right,  HParser* p,  HParser* q);
func Right(p HParser, q HParser) HParser { return C.h_right(p, q) }

//HAMMER_FN_DECL(HParser*, h_middle,  HParser* p,  HParser* x,  HParser* q);
func Middle(p HParser, x HParser, q HParser) HParser { return C.h_middle(p, x, q) }

//HAMMER_FN_DECL(HParser*, h_action,  HParser* p,  HAction a);
func Action(p HParser, a HAction) HParser { return C.h_action(p, (C.HAction)(a)) }

//HAMMER_FN_DECL(HParser*, h_in,  uint8_t *charset, size_t length);
func In(charset *uint8, length uintptr) HParser {
	return C.h_in((*C.uint8_t)(charset), (C.size_t)(length))
}

//HAMMER_FN_DECL(HParser*, h_not_in,  uint8_t *charset, size_t length);
func Not_in(charset *uint8, length uintptr) HParser {
	return C.h_not_in((*C.uint8_t)(charset), (C.size_t)(length))
}

//HAMMER_FN_DECL_NOARG(HParser*, h_end_p);
func End_p() HParser { return C.h_end_p() }

//HAMMER_FN_DECL_NOARG(HParser*, h_nothing_p);
func Nothing_p() HParser { return C.h_nothing_p() }




/* ok, a) this is the most nasty ball of hacked up ptr shit ever, b)
   the funcs called here are in src/parsers/choice.c, not hammer.h.
   cus the original one is variatic and cgo can't do va funcs as of
   1.2rc2 <_<... */

//HAMMER_FN_DECL_VARARGS_ATTR(__attribute__((sentinel)), HParser*, h_sequence, HParser* p);
func Sequence(p ...HParser) HParser {
	var cPtrs []unsafe.Pointer
	for _, cPtr := range(p) {
		cPtrs = append(cPtrs, (unsafe.Pointer)(cPtr))
	}
	var cPtr2cPtrs = &cPtrs[0]

	return C.h_sequence__a(cPtr2cPtrs)
}

//HAMMER_FN_DECL_VARARGS_ATTR(__attribute__((sentinel)), HParser*, h_choice, HParser* p);
func Choice(p ...HParser) HParser {
	var cPtrs []unsafe.Pointer
	for _, cPtr := range(p) {
		cPtrs = append(cPtrs, (unsafe.Pointer)(cPtr))
	}
	var cPtr2cPtrs = &cPtrs[0]

	return C.h_choice__a(cPtr2cPtrs)
}




//HAMMER_FN_DECL(HParser*, h_butnot,  HParser* p1,  HParser* p2);
func Butnot(p1 HParser, p2 HParser) HParser { return C.h_butnot(p1, p2) }

//HAMMER_FN_DECL(HParser*, h_difference,  HParser* p1,  HParser* p2);
func Difference(p1 HParser, p2 HParser) HParser { return C.h_difference(p1, p2) }

//HAMMER_FN_DECL(HParser*, h_xor,  HParser* p1,  HParser* p2);
func Xor(p1 HParser, p2 HParser) HParser { return C.h_xor(p1, p2) }

//HAMMER_FN_DECL(HParser*, h_many,  HParser* p);
func Many(p HParser) HParser { return C.h_many(p) }

//HAMMER_FN_DECL(HParser*, h_many1,  HParser* p);
func Many1(p HParser) HParser { return C.h_many1(p) }

//HAMMER_FN_DECL(HParser*, h_repeat_n,  HParser* p,  size_t n);
func Repeat_n(p HParser, n uintptr) HParser { return C.h_repeat_n(p, (C.size_t)(n)) }

//HAMMER_FN_DECL(HParser*, h_optional,  HParser* p);
func Optional(p HParser) HParser { return C.h_optional(p) }

//HAMMER_FN_DECL(HParser*, h_ignore,  HParser* p);
func Ignore(p HParser) HParser { return C.h_ignore(p) }

//HAMMER_FN_DECL(HParser*, h_sepBy,  HParser* p,  HParser* sep);
func SepBy(p HParser, sep HParser) HParser { return C.h_sepBy(p, sep) }

//HAMMER_FN_DECL(HParser*, h_sepBy1,  HParser* p,  HParser* sep);
func SepBy1(p HParser, sep HParser) HParser { return C.h_sepBy1(p, sep) }

//HAMMER_FN_DECL_NOARG(HParser*, h_epsilon_p);
func Epsilon_p() HParser { return C.h_epsilon_p() }

//HAMMER_FN_DECL(HParser*, h_length_value,  HParser* length,  HParser* value);
func Length_value(length HParser, value HParser) HParser {
	return C.h_length_value(length, value)
}

//HAMMER_FN_DECL(HParser*, h_attr_bool,  HParser* p, HPredicate pred);
func Attr_bool(p HParser, pred HPredicate) HParser { return C.h_attr_bool(p, pred) }

//HAMMER_FN_DECL(HParser*, h_and,  HParser* p);
func And(p HParser) HParser { return C.h_and(p) }

//HAMMER_FN_DECL(HParser*, h_not,  HParser* p);
func Not(p HParser) HParser { return C.h_not(p) }

//HAMMER_FN_DECL_NOARG(HParser*, h_indirect);
func Indirect() HParser { return C.h_indirect() }

//HAMMER_FN_DECL(void, h_bind_indirect, HParser* indirect,  HParser* inner);
func Bind_indirect(indirect HParser, inner HParser) { C.h_bind_indirect(indirect, inner) }

//HAMMER_FN_DECL(void, h_parse_result_free, HParseResult *result);
func Parse_result_free(result HParseResult) { C.h_parse_result_free(result) }

//HAMMER_FN_DECL(char*, h_write_result_unamb,  HParsedToken* tok);
func Write_result_unamb(tok HParsedToken) string {
	return C.GoString(C.h_write_result_unamb(tok))
}

// HAMMER_FN_DECL(void, h_pprint, FILE* stream,  HParsedToken* tok, int indent, int delta);
func Pprint(stream File, tok HParsedToken, indent int, delta int) {
	C.h_pprint(stream, tok, (C.int)(indent), (C.int)(delta))
}

//HAMMER_FN_DECL(int, h_compile, HParser* parser, HParserBackend backend,  void* params);
func Compile(parser HParser, backend HParserBackend, params Void) int {
	return int(C.h_compile(parser, (C.HParserBackend)(backend), (unsafe.Pointer)(params)))
}
