package ast

type Token struct {
	// Using built in parsers, may be a []byte, int64, uint64, []Token or
	// None. An HAction may return any type.
	Value interface{}

	ByteOffset int64
	BitOffset  int8
}

// Represents a TT_NONE token
type None struct{}
