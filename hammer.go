package hammer

import (
	"errors"
	"unsafe"
)

/*
	#cgo CFLAGS: -Ihammer/src
	#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
	#include <hammer.h>
*/
import "C"

type Parser *C.HParser

type ParserBackend C.HParserBackend

const (
	Packrat ParserBackend = C.PB_PACKRAT // default
	Regular ParserBackend = C.PB_REGULAR
	LLk     ParserBackend = C.PB_LLk
	LALR    ParserBackend = C.PB_LALR
	GLR     ParserBackend = C.PB_GLR
)

func BindIndirect(indirect Parser, inner Parser) {
	C.h_bind_indirect(indirect, inner)
}

func Compile(parser Parser, backend ParserBackend, params unsafe.Pointer) error {
	ret := int(C.h_compile(parser, C.HParserBackend(backend), params))
	if ret != 0 {
		return errors.New("failed to compile")
	}

	return nil
}

// utility function to convert a byteslice to a C array
func byteToCArr(p []byte) (arr *C.uint8_t, n C.size_t) {
	arr = (*C.uint8_t)(unsafe.Pointer(&p[0]))
	n = C.size_t(len(p))

	return
}
