package ast

type Token interface {
	Index() int64
	BitOffset() uint8
}

type Location struct {
	IndexF     int64
	BitOffsetF uint8
}

func (l Location) Index() int64     { return l.Index() }
func (l Location) BitOffset() uint8 { return l.BitOffsetF }

type NoneToken struct {
	Location
}

type BytesToken struct {
	Location
	Bytes []byte
}

type IntToken struct {
	Location
	Int int64
}

type UintToken struct {
	Location
	Uint uint64
}

type SequenceToken struct {
	Location
	Seq []Token
}
