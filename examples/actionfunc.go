package main

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.upstandinghackers.com/hammer"
	"gopkg.upstandinghackers.com/hammer/ast"
)

var uint_10 hammer.Parser

func init() {
	digit := hammer.ChRange('0', '9')
	uint_10 = hammer.Action(hammer.Many1(digit), func(token ast.Token) (ast.Token, bool) {
		numstr := charSequenceString(token.Value.([]ast.Token))

		num, err := strconv.ParseUint(numstr, 10, 64)
		if err != nil {
			return ast.Token{}, false
		}

		token.Value = num
		return token, true
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
