package hammer

import (
	"errors"
	"reflect"
	"unsafe"

	"gopkg.upstandinghackers.com/hammer/ast"
)

/*
	#cgo pkg-config: libhammer
	#include <hammer/hammer.h>
	#include <stddef.h>

	int HParsedTokenUnionOffset();

	int HParsedTokenUnionOffset() {
		return offsetof(HParsedToken, sint);
	}
*/
import "C"

var parseFailed = errors.New("parse failed")

func Parse(parser Parser, input []byte) (token ast.Token, err error) {
	res := cParse(parser, input)
	defer res.free()

	if res.r == nil {
		return token, parseFailed
	}

	cacheMu.Lock()
	delete(tokenCache, res.r.arena)
	cacheMu.Unlock()

	return convertCToken(res.r.ast), nil
}

func convertCToken(ctoken *C.HParsedToken) ast.Token {
	if ctoken == nil {
		return ast.Token{}
	}

	token := ast.Token{
		ByteOffset: int64(ctoken.index),
		BitOffset:  int8(ctoken.bit_offset),
	}

	switch ctoken.token_type {
	case C.TT_NONE:
		token.Value = ast.None
	case C.TT_BYTES:
		token.Value = convertHBytes(ctoken)
	case C.TT_SINT:
		token.Value = *(*int64)(unionPointer(ctoken))
	case C.TT_UINT:
		token.Value = *(*uint64)(unionPointer(ctoken))
	case C.TT_SEQUENCE:
		token.Value = convertHCountedArray(ctoken)
	case goHParsedToken:
		token.Value = **(**interface{})(unionPointer(ctoken))
	}

	return token
}

var unionOffset = uintptr(C.HParsedTokenUnionOffset())

func unionPointer(ctoken *C.HParsedToken) unsafe.Pointer {
	// Conversion here is to ensure ctoken is in fact a pointer
	return unsafe.Pointer(uintptr(unsafe.Pointer(ctoken)) + unionOffset)
}

func convertHBytes(ctoken *C.HParsedToken) []byte {
	hbytes := *(*C.HBytes)(unionPointer(ctoken))
	return C.GoBytes(unsafe.Pointer(hbytes.token), C.int(hbytes.len))
}

func convertHCountedArray(ctoken *C.HParsedToken) []ast.Token {
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
		ret[i] = convertCToken(token)
	}

	return ret
}
