package main

import (
	"hammer"
)

func jsonParser() hammer.HParser {
	// string
	jsonString := hammer.Sequence(
		hammer.Ch('"'),
		hammer.Many1(
			hammer.Choice(
				hammer.Ch_range('\x20', '\x21'), // skip "
				hammer.Ch_range('\x23', '\x2E'), // skip /
				hammer.Ch_range('\x30', '\x5B'), // skip \
				hammer.Ch_range('\x5D', '\x7E'), // skip DEL
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
				hammer.Ch_range('1', '9'),
				hammer.Many1(hammer.Ch_range('0', '9')),
			),
		),
		hammer.Optional(hammer.Sequence(
			hammer.Ch('.'),
			hammer.Many1(hammer.Ch_range(0, 9)),
		)),
		hammer.Optional(hammer.Sequence(
			hammer.Choice(hammer.Ch('e'), hammer.Ch('E')),
			hammer.Optional(hammer.Choice(hammer.Ch('+'), hammer.Ch('-'))),
			hammer.Many1(hammer.Ch_range(0, 9)),
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
	hammer.Bind_indirect(jsonArray,
		hammer.Sequence(
			hammer.Ch('['),
			hammer.SepBy1(jsonValue, hammer.Ch(',')),
			hammer.Ch(']'),
		),
	)

	// object
	hammer.Bind_indirect(jsonObject,
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
	// init parser
	jsonParser()

	// read input

}
