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

type Parser *C.HParser
type ParserBackend C.HParserBackend

func Bind_indirect(indirect Parser, inner Parser) {
	C.h_bind_indirect(indirect, inner)
}

func Compile(parser Parser, backend ParserBackend, params unsafe.Pointer) int {
	return int(C.h_compile(parser, C.HParserBackend(backend), params))
}

// utility function to convert a byteslice to a C array
func byteToCArr(p []byte) (arr *C.uint8_t, n C.size_t) {
	arr = (*C.uint8_t)(unsafe.Pointer(&p[0]))
	n = C.size_t(len(p))

	return
}
