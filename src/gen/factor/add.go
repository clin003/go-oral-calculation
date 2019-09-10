package factor

import (
	"fmt"
	"gen"
)

func NewAdd(a gen.IFactor, b gen.IFactor) gen.IFactor {
	return &add{a: a, b: b}
}

type add struct {
	a gen.IFactor
	b gen.IFactor
}

func (f *add) Result() int {
	return f.a.Result() + f.b.Result()
}

func (f *add) Format(t gen.FormatType) string {
	switch t = gen.ToFixFormat(t); t {
	case gen.PH_LEFT:
		return fmt.Sprintf("(   ) + %2d = %2d", f.b.Result(), f.Result())
	case gen.PH_RIGHT:
		return fmt.Sprintf("%2d + (   ) = %2d", f.a.Result(), f.Result())
	case gen.PH_RESULT:
		return fmt.Sprintf("%2d + %2d = (   )", f.a.Result(), f.b.Result())
	}

	return ""
}
