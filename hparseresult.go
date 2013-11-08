package hammer

import (
	"runtime"
)

/*
	#cgo CFLAGS: -Ihammer/src
	#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
	#include <hammer.h>
*/
import "C"

type HParseResult struct {
	*hParseResult
}

// The extra level of indirection ensures the finalizer never frees the wrong
// data causing a double free.
type hParseResult struct {
	r *C.HParseResult
}

func Parse(parser HParser, input []byte) *HParseResult {
	arr, n := byteToCArr(input)
	return newHParseResult(C.h_parse(parser, arr, n))
}

func newHParseResult(r *C.HParseResult) *HParseResult {
	ret := &HParseResult{&hParseResult{r}}
	runtime.SetFinalizer(ret.hParseResult, (*hParseResult).free)
	return ret
}

//HAMMER_FN_DECL(void, h_parse_result_free, HParseResult *result);
func (p *HParseResult) Free() {
	if p == nil {
		return
	}

	p.hParseResult.free()
}

func (p *hParseResult) free() {
	if p.r == nil {
		return
	}

	C.h_parse_result_free(p.r)

	p.r = nil
	runtime.SetFinalizer(p, nil)
}
