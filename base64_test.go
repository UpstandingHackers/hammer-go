package hammer

import (
	"fmt"
	"testing"
)

func base64_parser() Parser {
	digit := ChRange('0', '9')
	alpha := Choice(
		ChRange('A', 'Z'),
		ChRange('a', 'z'),
	)

	plus := Ch('+')
	slash := Ch('/')
	equals := Ch('=')

	bsfdig := Choice(alpha, digit, plus, slash)
	bsfdig_4bit := In([]byte("AEIMQUYcgkosw048"))
	bsfdig_2bit := In([]byte("AQgw"))

	base64_3 := RepeatN(bsfdig, 4)
	base64_2 := Sequence(bsfdig, bsfdig, bsfdig_4bit, equals)
	base64_1 := Sequence(bsfdig, bsfdig_2bit, equals, equals)

	base64 := Sequence(
		Many(base64_3),
		Optional(
			Choice(base64_2, base64_1),
		),
	)

	return base64
}

func TestBase64(t *testing.T) {
	base64Parser := base64_parser()
	documentParser := Sequence(
		Whitespace(base64Parser),
		Whitespace(End()),
	)

	input := []byte("aGk=") // "hi"
	ast, err := Parse(documentParser, input)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", ast)
}
