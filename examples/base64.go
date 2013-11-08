package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/prevoty/hammer"
)

/* different than the C example because go can't use nil as a
   const initilizer */
var document hammer.HParser = nil

func init_parser() {
	// CORE
	var digit hammer.HParser = hammer.Ch_range(0x30, 0x39)
	var alpha hammer.HParser = hammer.Choice(hammer.Ch_range(0x41, 0x5a), hammer.Ch_range(0x61, 0x7a), nil)

	// AUX.
	var plus hammer.HParser = hammer.Ch('+')
	var slash hammer.HParser = hammer.Ch('/')
	var equals hammer.HParser = hammer.Ch('=')

	var bsfdig hammer.HParser = hammer.Choice(alpha, digit, plus, slash, nil)
	var bsfdig_4bit hammer.HParser = hammer.In([]byte("AEIMQUYcgkosw048"))
	var bsfdig_2bit hammer.HParser = hammer.In([]byte("AQgw"))
	var base64_3 hammer.HParser = hammer.Repeat_n(bsfdig, 4)
	var base64_2 hammer.HParser = hammer.Sequence(bsfdig, bsfdig, bsfdig_4bit, equals, nil)
	var base64_1 hammer.HParser = hammer.Sequence(bsfdig, bsfdig_2bit, equals, equals, nil)
	var base64 hammer.HParser = hammer.Sequence(hammer.Many(base64_3), hammer.Optional(hammer.Choice(base64_2, base64_1, nil)), nil)
	document = hammer.Sequence(hammer.Whitespace(base64), hammer.Whitespace(hammer.End_p()), nil)
}

func main() {
	var input []byte
	var result hammer.HParseResult

	// deffine grammer
	init_parser()

	// get user input
	io.Copy(bytes.NewBuffer(input), os.Stdin)

	// show what we got
	inputsize := len(input)
	fmt.Printf("inputsize=%d\ninput=%s\n", inputsize, input)

	// parse input
	result = hammer.Parse(document, input)

	// show result of parse actions
	if result {
		fmt.Printf("parsed=%d bytes\n", result.bit_length/8)
		hammer.Pprint(stdout, result.ast, 0, 0)
		return 0
	} else {
		return 1
	}
}
