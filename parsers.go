package hammer

import (
	"unsafe"
)

/*
	#cgo CFLAGS: -Ihammer/src
	#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
	#include <hammer.h>
*/
import "C"

func Token(str []byte) HParser {
	arr, n := byteToCArr(str)
	return C.h_token(arr, n)
}

func Ch(c uint8) HParser { return C.h_ch(C.uint8_t(c)) }

func Ch_range(lower uint8, upper uint8) HParser {
	return C.h_ch_range(C.uint8_t(lower), C.uint8_t(upper))
}

func Int_range(p HParser, lower int64, upper int64) HParser {
	return C.h_int_range(p, C.int64_t(lower), C.int64_t(upper))
}

func Bits(len uintptr, sign bool) HParser {
	// bool to C.bool conversion
	var retVal C.bool
	if sign {
		retVal = 1
	}

	return C.h_bits(C.size_t(len), retVal)
}

func Int64() HParser  { return C.h_int64() }
func Int32() HParser  { return C.h_int32() }
func Int16() HParser  { return C.h_int16() }
func Int8() HParser   { return C.h_int8() }
func Uint64() HParser { return C.h_uint64() }
func Uint32() HParser { return C.h_uint32() }
func Uint16() HParser { return C.h_uint16() }
func Uint8() HParser  { return C.h_uint8() }

func Whitespace(p HParser) HParser { return C.h_whitespace(p) }

func Left(p HParser, q HParser) HParser              { return C.h_left(p, q) }
func Right(p HParser, q HParser) HParser             { return C.h_right(p, q) }
func Middle(p HParser, x HParser, q HParser) HParser { return C.h_middle(p, x, q) }

func Action(p HParser, a HAction) HParser { return C.h_action(p, C.HAction(a)) }

func In(charset []byte) HParser {
	arr, n := byteToCArr(charset)
	return C.h_in(arr, n)
}

func Not_in(charset []byte) HParser {
	arr, n := byteToCArr(charset)
	return C.h_not_in(arr, n)
}

func End_p() HParser     { return C.h_end_p() }
func Nothing_p() HParser { return C.h_nothing_p() }

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

func Butnot(p1 HParser, p2 HParser) HParser { return C.h_butnot(p1, p2) }

func Difference(p1 HParser, p2 HParser) HParser { return C.h_difference(p1, p2) }

func Xor(p1 HParser, p2 HParser) HParser { return C.h_xor(p1, p2) }

func Many(p HParser) HParser { return C.h_many(p) }

func Many1(p HParser) HParser { return C.h_many1(p) }

func Repeat_n(p HParser, n uintptr) HParser { return C.h_repeat_n(p, C.size_t(n)) }

func Optional(p HParser) HParser { return C.h_optional(p) }

func Ignore(p HParser) HParser { return C.h_ignore(p) }

func SepBy(p HParser, sep HParser) HParser { return C.h_sepBy(p, sep) }

func SepBy1(p HParser, sep HParser) HParser { return C.h_sepBy1(p, sep) }

func Epsilon_p() HParser { return C.h_epsilon_p() }

func Length_value(length HParser, value HParser) HParser {
	return C.h_length_value(length, value)
}

func Attr_bool(p HParser, pred HPredicate) HParser { return C.h_attr_bool(p, pred) }

func And(p HParser) HParser { return C.h_and(p) }

func Not(p HParser) HParser { return C.h_not(p) }

func Indirect() HParser { return C.h_indirect() }
