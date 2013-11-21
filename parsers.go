package hammer

import (
	"unsafe"
)

/*
	#cgo pkg-config: libhammer
	#include <hammer/hammer.h>
*/
import "C"

func Token(str []byte) Parser {
	arr, n := byteToCArr(str)
	return C.h_token(arr, n)
}

func Ch(c uint8) Parser { return C.h_ch(C.uint8_t(c)) }

func ChRange(lower uint8, upper uint8) Parser {
	return C.h_ch_range(C.uint8_t(lower), C.uint8_t(upper))
}

func IntRange(p Parser, lower int64, upper int64) Parser {
	return C.h_int_range(p, C.int64_t(lower), C.int64_t(upper))
}

func Bits(len int, sign bool) Parser {
	// bool to C.bool conversion
	var retVal C.bool
	if sign {
		retVal = 1
	}

	return C.h_bits(C.size_t(len), retVal)
}

func Int64() Parser  { return C.h_int64() }
func Int32() Parser  { return C.h_int32() }
func Int16() Parser  { return C.h_int16() }
func Int8() Parser   { return C.h_int8() }
func Uint64() Parser { return C.h_uint64() }
func Uint32() Parser { return C.h_uint32() }
func Uint16() Parser { return C.h_uint16() }
func Uint8() Parser  { return C.h_uint8() }

func Whitespace(p Parser) Parser { return C.h_whitespace(p) }

func Left(p Parser, q Parser) Parser             { return C.h_left(p, q) }
func Right(p Parser, q Parser) Parser            { return C.h_right(p, q) }
func Middle(p Parser, x Parser, q Parser) Parser { return C.h_middle(p, x, q) }

func In(charset []byte) Parser {
	arr, n := byteToCArr(charset)
	return C.h_in(arr, n)
}

func NotIn(charset []byte) Parser {
	arr, n := byteToCArr(charset)
	return C.h_not_in(arr, n)
}

func End() Parser     { return C.h_end_p() }
func Nothing() Parser { return C.h_nothing_p() }

func Sequence(p ...Parser) Parser {
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

func Choice(p ...Parser) Parser {
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

func ButNot(p1 Parser, p2 Parser) Parser { return C.h_butnot(p1, p2) }

func Difference(p1 Parser, p2 Parser) Parser { return C.h_difference(p1, p2) }

func Xor(p1 Parser, p2 Parser) Parser { return C.h_xor(p1, p2) }

func Many(p Parser) Parser { return C.h_many(p) }

func Many1(p Parser) Parser { return C.h_many1(p) }

func RepeatN(p Parser, n uintptr) Parser { return C.h_repeat_n(p, C.size_t(n)) }

func Optional(p Parser) Parser { return C.h_optional(p) }

func Ignore(p Parser) Parser { return C.h_ignore(p) }

func SepBy(p Parser, sep Parser) Parser { return C.h_sepBy(p, sep) }

func SepBy1(p Parser, sep Parser) Parser { return C.h_sepBy1(p, sep) }

func Epsilon() Parser { return C.h_epsilon_p() }

func LengthValue(length Parser, value Parser) Parser {
	return C.h_length_value(length, value)
}

func And(p Parser) Parser { return C.h_and(p) }

func Not(p Parser) Parser { return C.h_not(p) }

func Indirect() Parser { return C.h_indirect() }
