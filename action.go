package hammer

import (
	"sync"
	"unsafe"

	"github.com/prevoty/hammer/ast"
)

/*
	#cgo pkg-config: libhammer
	#include <hammer/hammer.h>

	#include "action.h"
*/
import "C"

var goHParsedToken = C.getGoTokenId()

var (
	// mutex for all variables in this group
	cacheMu = new(sync.Mutex)
	// tokens pinned to arenas
	tokenCache = make(map[*C.HArena][]*C.HParsedToken)
	// values of go tokens
	gotokenCache = make(map[*C.HArena][]*interface{})
	// functions cached forever
	funcCache = []*ActionFunc{}
)

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
			valid: false,
		}
	}

	resultCtoken, govalue := convertToken(result)

	if resultCtoken != nil {
		cacheMu.Lock()
		tokenCache[pr.arena] = append(tokenCache[pr.arena], resultCtoken)
		if govalue != nil {
			gotokenCache[pr.arena] = append(gotokenCache[pr.arena], govalue)
		}
		cacheMu.Unlock()
	}

	return C.GoActionResult{
		token: resultCtoken,
		valid: true,
	}
}

//TODO: implement TT_BYTES and other tokens
func convertToken(token ast.Token) (ctoken *C.HParsedToken, govalue *interface{}) {
	switch v := token.Value.(type) {
	case nil:
		return nil, nil
	case ast.NoneType:
		ctoken := &C.HParsedToken{
			token_type: C.TT_NONE,
			index:      C.size_t(token.ByteOffset),
			bit_offset: C.char(token.BitOffset),
		}
		return ctoken, nil
	case uint64:
		ctoken := &C.HParsedToken{
			token_type: C.TT_UINT,
			index:      C.size_t(token.ByteOffset),
			bit_offset: C.char(token.BitOffset),
		}
		C.assignUintValue(ctoken, C.uint64_t(v))
		return ctoken, nil
	}

	// go token
	ctoken = &C.HParsedToken{
		token_type: goHParsedToken,
		index:      C.size_t(token.ByteOffset),
		bit_offset: C.char(token.BitOffset),
	}
	union := (**interface{})(unionPointer(ctoken))
	*union = &token.Value

	return ctoken, &token.Value
}
