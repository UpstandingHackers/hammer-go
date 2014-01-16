package main

import (
	"fmt"
	"log"

	"gopkg.upstandinghackers.com/hammer"
)

func base64_parser() hammer.Parser {
	digit := hammer.ChRange('0', '9')
	alpha := hammer.Choice(
		hammer.ChRange('A', 'Z'),
		hammer.ChRange('a', 'z'),
	)

	plus := hammer.Ch('+')
	slash := hammer.Ch('/')
	equals := hammer.Ch('=')

	bsfdig := hammer.Choice(alpha, digit, plus, slash)
	bsfdig_4bit := hammer.In([]byte("AEIMQUYcgkosw048"))
	bsfdig_2bit := hammer.In([]byte("AQgw"))

	base64_3 := hammer.RepeatN(bsfdig, 4)
	base64_2 := hammer.Sequence(bsfdig, bsfdig, bsfdig_4bit, equals)
	base64_1 := hammer.Sequence(bsfdig, bsfdig_2bit, equals, equals)

	base64 := hammer.Sequence(
		hammer.Many(base64_3),
		hammer.Optional(
			hammer.Choice(base64_2, base64_1),
		),
	)

	return base64
}

func main() {
	base64Parser := base64_parser()
	documentParser := hammer.Sequence(
		hammer.Whitespace(base64Parser),
		hammer.Whitespace(hammer.End()),
	)

	input := []byte("aGk=") // "hi"
	ast, err := hammer.Parse(documentParser, input)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%#v\n", ast)
}
