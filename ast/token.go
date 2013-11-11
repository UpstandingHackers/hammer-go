package ast

type Token struct {
	Value interface{}

	ByteOffset int64
	BitOffset  int8
}

type NoneType struct{}

// The value of a TT_NONE token
var None = NoneType{}
