package factor

import (
	"fmt"
	"gen"
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

func (f *integer) ToString() string {
	return fmt.Sprintf("%d", f.v)
}
