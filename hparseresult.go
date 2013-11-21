package hammer

/*
	#cgo pkg-config: libhammer
	#include <hammer/hammer.h>
*/
import "C"

type hParseResult struct {
	r *C.HParseResult
}

// Like Parse() but returns an HParseResult instead of an AST. You problably shouldn't be using this.
func cParse(parser Parser, input []byte) *hParseResult {
	arr, n := byteToCArr(input)
	return &hParseResult{C.h_parse(parser, arr, n)}
}

func (p *hParseResult) free() {
	if p.r == nil {
		return
	}

	C.h_parse_result_free(p.r)
	p.r = nil
}
