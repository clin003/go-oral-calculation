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

func (f *integer) Format(t gen.FormatType, withResult bool) string {
	if t == gen.PH_NONE {
		return fmt.Sprintf("%2d", f.v)
	} else {
		return "(   )"
	}
}
