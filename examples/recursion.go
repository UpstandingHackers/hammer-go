package main

import (
	"fmt"
	"log"

	"gopkg.upstandinghackers.com/hammer"
)

var aParser hammer.Parser

func init() {
	aParser = hammer.Indirect()

	hammer.BindIndirect(aParser, hammer.Sequence(
		hammer.Ch('a'),
		hammer.Choice(aParser, hammer.End()),
	))
}

func main() {
	data := "aaaaa"
	ast, err := hammer.Parse(aParser, []byte(data))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", ast)
}
