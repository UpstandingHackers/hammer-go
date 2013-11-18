package hammer

import (
	"unsafe"

	"hammer/ast"
)

/*
	#cgo CFLAGS: -Ihammer/src -std=gnu99
	#cgo LDFLAGS: hammer/build/opt/src/libhammer.a
	#include <hammer.h>

	#include "action.h"
*/
import "C"

const goHParsedToken = 200

// ActionFunc allows for transformation and validation of the parse tree
// represented by the input token. It returns a transformed parse tree as
// a token and whether or not the parse is valid as a bool.
type ActionFunc func(token ast.Token) (result ast.Token, ok bool)

func Action(p Parser, a ActionFunc) Parser {
	fp := &a

	cacheMu.Lock()
	funcCache = append(funcCache, fp)
	cacheMu.Unlock()

	return C.go_action(p, unsafe.Pointer(fp))
}

//export go_action_hook
func go_action_hook(action unsafe.Pointer, pr *C.HParseResult) C.GoActionResult {
	act := *(*ActionFunc)(action)
	token := convertCToken(pr.ast)
	result, ok := act(token)
	if !ok {
		return C.GoActionResult{
			valid: 0,
		}
	}

	resultCtoken := convertToken(result)

	if resultCtoken != nil {
		cacheMu.Lock()
		tokenCache[pr.arena] = append(tokenCache[pr.arena], resultCtoken)
		cacheMu.Unlock()
	}

	return C.GoActionResult{
		token: resultCtoken,
		valid: 1,
	}
}

//TODO: implement NONE, perhaps TT_BYTES and others
func convertToken(token ast.Token) *C.HParsedToken {
	switch v := token.Value.(type) {
	case nil:
		return nil
	case uint64:
		ctoken := &C.HParsedToken{
			token_type: C.TT_UINT,
			index:      C.size_t(token.ByteOffset),
			bit_offset: C.char(token.BitOffset),
		}
		C.assignUintValue(ctoken, C.uint64_t(v))
		return ctoken
	}

	// go token
	ctoken := &C.HParsedToken{
		token_type: goHParsedToken,
		index:      C.size_t(token.ByteOffset),
		bit_offset: C.char(token.BitOffset),
	}
	C.assignVoidValue(ctoken, unsafe.Pointer(&token))

	return ctoken
}
