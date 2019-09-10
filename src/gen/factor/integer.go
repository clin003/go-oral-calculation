package factor

import (
	"gen"
	"strconv"
)

func NewInteger(v int) gen.IFactor {
	return &integer{v: v}
}

type integer struct {
	v int
}

func (f *integer) Result() int {
	return f.v
}

func (f *integer) Format(t gen.FormatType) string {
	return strconv.Itoa(f.v)
}
