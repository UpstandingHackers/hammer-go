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
type Parser *C.HParser
type ParserBackend C.HParserBackend

func byteToCArr(p []byte) (arr *C.uint8_t, n C.size_t) {
	arr = (*C.uint8_t)(unsafe.Pointer(&p[0]))
	n = C.size_t(len(p))

	return
}

func Bind_indirect(indirect Parser, inner Parser) { C.h_bind_indirect(indirect, inner) }

func Write_result_unamb(tok HParsedToken) string {
	return C.GoString(C.h_write_result_unamb(tok))
}

func Pprint(stream *os.File, tok HParsedToken, indent int, delta int) {
	cfile := C.fdopen(C.int(stream.Fd()), C.wFlag)
	C.h_pprint(cfile, tok, C.int(indent), C.int(delta))
}

func Compile(parser Parser, backend ParserBackend, params unsafe.Pointer) int {
	return int(C.h_compile(parser, C.HParserBackend(backend), params))
}
