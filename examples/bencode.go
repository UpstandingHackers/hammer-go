package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"hammer"
	"hammer/ast"
)

var bencode hammer.HParser

func init() {
	end := hammer.Ch('e')
	colon := hammer.Ch(':')
	neg := hammer.Ch('-')

	digits := hammer.Many1(hammer.Ch_range('0', '9'))

	uint_10 := hammer.Action(digits, func(token ast.Token) ast.Token {
		numstr := charSequenceString(token.Value.([]ast.Token))

		num, err := strconv.ParseUint(numstr, 10, 64)
		if err != nil {
			token.Value = nil
			return token
		}

		token.Value = num
		return token
	})

	int_10 := hammer.Action(hammer.Sequence(hammer.Optional(neg), digits), func(token ast.Token) ast.Token {
		tokens := token.Value.([]ast.Token)
		neg, digits := tokens[0], tokens[1]
		numstr := charSequenceString(digits.Value.([]ast.Token))

		num, err := strconv.ParseInt(numstr, 10, 64)
		if err != nil {
			token.Value = nil
			return token
		}

		if neg.Value != ast.None {
			num *= -1
		}

		token.Value = num
		return token
	})

	b_string := hammer.Length_value(hammer.Left(uint_10, colon), hammer.Uint8())
	b_int := hammer.Sequence(hammer.Ch('i'), int_10, end)
	b_list := hammer.Indirect()
	b_dict := hammer.Indirect()

	anyType := hammer.Choice(b_int, b_string, b_list, b_dict)

	hammer.Bind_indirect(b_list, hammer.Sequence(
		hammer.Ch('l'),
		hammer.Many(anyType),
		end,
	))
	hammer.Bind_indirect(b_dict, hammer.Sequence(
		hammer.Ch('d'),
		hammer.Many(hammer.Sequence(b_string, anyType)),
		end,
	))

	bencode = anyType
}

// Takes a slice of uint64 tokens and concatinates them as chars in the
// returned string.
func charSequenceString(tokens []ast.Token) string {
	buf := make([]byte, len(tokens))
	for i, t := range tokens {
		buf[i] = byte(t.Value.(uint64))
	}

	return string(buf)
}

func main() {
	data := "d1:Xi1e1:Yi2e1:Z5:helloe"
	//data := "i-10e"
	ast, err := hammer.Parse(bencode, []byte(data))

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(ast)
	if err != nil {
		log.Fatal(err)
	}
}
