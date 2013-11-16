package main

import (
	"fmt"
	"log"
	"strconv"

	"hammer"
	"hammer/ast"
)

var uint_10 hammer.HParser

func init() {
	digit := hammer.Ch_range('0', '9')
	uint_10 = hammer.Action(hammer.Many1(digit), func(token ast.Token) ast.Token {
		numstr := charSequenceString(token.Value.([]ast.Token))

		num, err := strconv.ParseUint(numstr, 10, 64)
		if err != nil {
			token.Value = nil
			return token
		}

		token.Value = num
		return token
	})
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
	data := "20"
	ast, err := hammer.Parse(uint_10, []byte(data))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", ast)
}
