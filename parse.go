package hammer

import (
	"reflect"
	"unsafe"

	"github.com/prevoty/hammer/ast"
)

/*
	#cgo CFLAGS: -Ihammer/src
	#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
	#include <hammer.h>
	#include <stddef.h>

	int HParsedTokenUnionOffset();

	int HParsedTokenUnionOffset() {
		return offsetof(HParsedToken, sint);
	}
*/
import "C"

func Parse(parser HParser, input []byte) ast.Token {
	res := CParse(parser, input)
	defer res.Free()

	if res.r.ast == nil {
		return nil
	}

	return convertToken(res.r.ast)
}

func convertToken(ctoken HParsedToken) ast.Token {
	loc := ast.Location{
		IndexF:     int64(ctoken.index),
		BitOffsetF: uint8(ctoken.bit_offset),
	}

	switch ctoken.token_type {
	case C.TT_NONE:
		return ast.NoneToken{loc}
	case C.TT_BYTES:
		return ast.BytesToken{loc, convertHBytes(ctoken)}
	case C.TT_SINT:
		i := *(*int64)(unionPointer(ctoken))
		return ast.IntToken{loc, i}
	case C.TT_UINT:
		i := *(*uint64)(unionPointer(ctoken))
		return ast.UintToken{loc, i}
	case C.TT_SEQUENCE:
		return ast.SequenceToken{loc, convertHCountedArray(ctoken)}
	}

	return nil
}

var unionOffset = uintptr(C.HParsedTokenUnionOffset())

func unionPointer(ctoken HParsedToken) unsafe.Pointer {
	// Conversion here is to ensure ctoken is in fact a pointer
	ptr := (*C.HParsedToken)(ctoken)
	return unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unionOffset)
}

func convertHBytes(ctoken HParsedToken) []byte {
	hbytes := *(*C.HBytes)(unionPointer(ctoken))
	return C.GoBytes(unsafe.Pointer(hbytes.token), C.int(hbytes.len))
}

func convertHCountedArray(ctoken HParsedToken) []ast.Token {
	hca := *(**C.HCountedArray)(unionPointer(ctoken))

	// elems is a []*C.HParsedToken using the hca.elements as a backing array
	shdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(hca.elements)),
		Len:  int(hca.used),
		Cap:  int(hca.used),
	}
	elems := *(*[]*C.HParsedToken)(unsafe.Pointer(&shdr))

	ret := make([]ast.Token, hca.used)
	for i, token := range elems {
		ret[i] = convertToken(token)
	}

	return ret
}
