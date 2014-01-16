package main

import (
	"fmt"
	"log"

	"gopkg.upstandinghackers.com/hammer"
)

func jsonParser_init() hammer.Parser {
	// string
	jsonString := hammer.Sequence(
		hammer.Ch('"'),
		hammer.Many1(
			hammer.Choice(
				hammer.ChRange('\x20', '\x21'), // skip "
				hammer.ChRange('\x23', '\x2E'), // skip /
				hammer.ChRange('\x30', '\x5B'), // skip \
				hammer.ChRange('\x5D', '\x7E'), // skip DEL
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('"')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('\\')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('/')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('b')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('f')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('n')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('r')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('t')),
				hammer.Sequence(hammer.Ch('\\'), hammer.Ch('u')),
			),
		),
		hammer.Ch('"'),
	)

	// number
	jsonNumber := hammer.Sequence(
		hammer.Optional(hammer.Ch('-')),
		hammer.Choice(
			hammer.Ch('0'),
			hammer.Sequence(
				hammer.ChRange('1', '9'),
				hammer.Many1(hammer.ChRange('0', '9')),
			),
		),
		hammer.Optional(hammer.Sequence(
			hammer.Ch('.'),
			hammer.Many1(hammer.ChRange(0, 9)),
		)),
		hammer.Optional(hammer.Sequence(
			hammer.Choice(hammer.Ch('e'), hammer.Ch('E')),
			hammer.Optional(hammer.Choice(hammer.Ch('+'), hammer.Ch('-'))),
			hammer.Many1(hammer.ChRange(0, 9)),
		)),
	)

	// indirects for use in value
	jsonArray := hammer.Indirect()
	jsonObject := hammer.Indirect()

	// value
	jsonValue := hammer.Choice(
		jsonString,
		jsonNumber,
		jsonObject,
		jsonArray,
		hammer.Token([]byte("true")),
		hammer.Token([]byte("false")),
		hammer.Token([]byte("null")),
	)

	// array
	hammer.BindIndirect(jsonArray,
		hammer.Sequence(
			hammer.Ch('['),
			hammer.SepBy1(jsonValue, hammer.Ch(',')),
			hammer.Ch(']'),
		),
	)

	// object
	hammer.BindIndirect(jsonObject,
		hammer.Sequence(
			hammer.Ch('{'),
			jsonString,
			hammer.Ch(':'),
			jsonValue,
			hammer.Ch('}'),
		),
	)

	return jsonObject
}

func main() {
	jsonParser := jsonParser_init()
	documentParser := hammer.Sequence(
		hammer.Whitespace(jsonParser),
		hammer.Whitespace(hammer.End()),
	)

	input := []byte("{\"name\":\"foo\",\"num\":100,\"balance\":1000.21,\"is_vip\":true,\"nickname\":null}")

	ast, err := hammer.Parse(documentParser, input)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%#v\n", ast)
}
