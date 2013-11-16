package main

import (
	"fmt"
	"log"

	"hammer"
)

var aParser hammer.HParser

func init() {
	aParser = hammer.Indirect()

	hammer.Bind_indirect(aParser, hammer.Sequence(
		hammer.Ch('a'),
		hammer.Choice(aParser, hammer.End_p()),
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
