package factor

import (
	"fmt"

	"github.com/clin003/kousuan/gen"
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
		return fmt.Sprintf("%d", f.v)
	} else {
		return "(   )"
	}
}
