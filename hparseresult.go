package hammer

/*
	#cgo CFLAGS: -Ihammer/src
	#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
	#include <hammer.h>
*/
import "C"

type HParseResult struct {
	result *C.HParseResult
}

// HAMMER_FN_DECL(HParseResult*, h_parse,  HParser* parser,  uint8_t* input, size_t length);
func Parse(parser HParser, input []byte) *HParseResult {
	arr, n := byteToCArr(input)
	return &HParseResult{C.h_parse(parser, arr, n)}
}

//HAMMER_FN_DECL(void, h_parse_result_free, HParseResult *result);
func (r *HParseResult) Free() {
	if r.result != nil {
		C.h_parse_result_free(r.result)
		r.result = nil
	}
}
