package hammer

import (
	"os"
	"unsafe"
)

/*
	#cgo CFLAGS: -Ihammer/src
	#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
	#include <hammer.h>

	// used in go Pprint() function
	char* wFlag = "w";
*/
import "C"

type HParsedToken *C.HParsedToken
type HAction C.HAction
type HPredicate C.HPredicate
type HParser *C.HParser
type HParserBackend C.HParserBackend

func byteToCArr(p []byte) (arr *C.uint8_t, n C.size_t) {
	arr = (*C.uint8_t)(unsafe.Pointer(&p[0]))
	n = C.size_t(len(p))

	return
}

// HAMMER_FN_DECL(HParser*, h_token,  uint8_t *str,  size_t len);
func Token(str []byte) HParser {
	arr, n := byteToCArr(str)
	return C.h_token(arr, n)
}

// HAMMER_FN_DECL(HParser*, h_ch,  uint8_t c);
func Ch(c uint8) HParser { return C.h_ch(C.uint8_t(c)) }

// HAMMER_FN_DECL(HParser*, h_ch_range,  uint8_t lower,  uint8_t upper);
func Ch_range(lower uint8, upper uint8) HParser {
	return C.h_ch_range(C.uint8_t(lower), C.uint8_t(upper))
}

// HAMMER_FN_DECL(HParser*, h_int_range,  HParser *p,  int64_t lower,  int64_t upper);
func Int_range(p HParser, lower int64, upper int64) HParser {
	return C.h_int_range(p, C.int64_t(lower), C.int64_t(upper))
}

// HAMMER_FN_DECL(HParser*, h_bits, size_t len, bool sign);
func Bits(len uintptr, sign bool) HParser {
	// bool to C.bool conversion
	var retVal C.bool
	if sign {
		retVal = 1
	}

	return C.h_bits(C.size_t(len), retVal)
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
func Action(p HParser, a HAction) HParser { return C.h_action(p, C.HAction(a)) }

//HAMMER_FN_DECL(HParser*, h_in,  uint8_t *charset, size_t length);
func In(charset []byte) HParser {
	arr, n := byteToCArr(charset)
	return C.h_in(arr, n)
}

//HAMMER_FN_DECL(HParser*, h_not_in,  uint8_t *charset, size_t length);
func Not_in(charset []byte) HParser {
	arr, n := byteToCArr(charset)
	return C.h_not_in(arr, n)
}

//HAMMER_FN_DECL_NOARG(HParser*, h_end_p);
func End_p() HParser { return C.h_end_p() }

//HAMMER_FN_DECL_NOARG(HParser*, h_nothing_p);
func Nothing_p() HParser { return C.h_nothing_p() }

//HAMMER_FN_DECL_VARARGS_ATTR(__attribute__((sentinel)), HParser*, h_sequence, HParser* p);
func Sequence(p ...HParser) HParser {
	ptrs := make([]unsafe.Pointer, len(p)+1)
	for i, ptr := range p {
		ptrs[i] = unsafe.Pointer(ptr)
	}

	// ptrs is a null termated array
	ptrs[len(ptrs)-1] = nil

	// C.h_sequence__a is defined in src/parsers/choice.c and not hammer.h.
	// This function is equivlent to the sequence function in hammer.h but
	// accepts an array instead of a variatic.

	return C.h_sequence__a(&ptrs[0])
}

//HAMMER_FN_DECL_VARARGS_ATTR(__attribute__((sentinel)), HParser*, h_choice, HParser* p);
func Choice(p ...HParser) HParser {
	ptrs := make([]unsafe.Pointer, len(p)+1)
	for i, ptr := range p {
		ptrs[i] = unsafe.Pointer(ptr)
	}

	// ptrs is a null termated array
	ptrs[len(ptrs)-1] = nil

	// C.h_choice__a is defined in src/parsers/choice.c and not hammer.h.
	// This function is equivlent to the sequence function in hammer.h but
	// accepts an array instead of a variatic.

	return C.h_choice__a(&ptrs[0])
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
func Repeat_n(p HParser, n uintptr) HParser { return C.h_repeat_n(p, C.size_t(n)) }

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

//HAMMER_FN_DECL(char*, h_write_result_unamb,  HParsedToken* tok);
func Write_result_unamb(tok HParsedToken) string {
	return C.GoString(C.h_write_result_unamb(tok))
}

// HAMMER_FN_DECL(void, h_pprint, FILE* stream,  HParsedToken* tok, int indent, int delta);
func Pprint(stream *os.File, tok HParsedToken, indent int, delta int) {
	cfile := C.fdopen(C.int(stream.Fd()), C.wFlag)
	C.h_pprint(cfile, tok, C.int(indent), C.int(delta))
}

//HAMMER_FN_DECL(int, h_compile, HParser* parser, HParserBackend backend,  void* params);
func Compile(parser HParser, backend HParserBackend, params unsafe.Pointer) int {
	return int(C.h_compile(parser, C.HParserBackend(backend), params))
}
