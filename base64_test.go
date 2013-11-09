package hammer

import (
	"fmt"
	"os"
	"testing"
)

/* different than the C example because go can't use nil as a
   const initilizer */
var document HParser = nil

func init_parser() {
	// CORE
	var digit HParser = Ch_range(0x30, 0x39)
	var alpha HParser = Choice(Ch_range(0x41, 0x5a), Ch_range(0x61, 0x7a), nil)

	// AUX.
	var plus HParser = Ch('+')
	var slash HParser = Ch('/')
	var equals HParser = Ch('=')

	var bsfdig HParser = Choice(alpha, digit, plus, slash, nil)
	var bsfdig_4bit HParser = In([]byte("AEIMQUYcgkosw048"))
	var bsfdig_2bit HParser = In([]byte("AQgw"))
	var base64_3 HParser = Repeat_n(bsfdig, 4)
	var base64_2 HParser = Sequence(bsfdig, bsfdig, bsfdig_4bit, equals, nil)
	var base64_1 HParser = Sequence(bsfdig, bsfdig_2bit, equals, equals, nil)
	var base64 HParser = Sequence(Many(base64_3), Optional(Choice(base64_2, base64_1, nil)), nil)
	document = Sequence(Whitespace(base64), Whitespace(End_p()), nil)
}

func TestBase64(t *testing.T) {
	init_parser()

	input := []byte("aGk=") // "hi"
	result := CParse(document, input)
	defer result.Free()

	// show result of parse actions
	if result.r != nil && result.r.ast != nil {
		Pprint(os.Stdout, result.r.ast, 0, 0)
	} else {
		t.Error("fail")
	}
}

func TestBase64_2(t *testing.T) {
	init_parser()

	input := []byte("aGk=") // "hi"
	ast := Parse(document, input)

	fmt.Printf("%#v\n", ast)
}
