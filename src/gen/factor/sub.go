package factor

import (
	"fmt"
	"gen"
)

func NewSub(a gen.IFactor, b gen.IFactor) gen.IFactor {
	return &sub{a: a, b: b}
}

type sub struct {
	a gen.IFactor
	b gen.IFactor
}

func (f *sub) Result() int {
	return f.a.Result() - f.b.Result()
}

func (f *sub) Format(t gen.FormatType) string {
	switch t = gen.ToFixFormat(t); t {
	case gen.PH_LEFT:
		return fmt.Sprintf("(   ) - %2d = %2d", f.b.Result(), f.Result())
	case gen.PH_RIGHT:
		return fmt.Sprintf("%2d - (   ) = %2d", f.a.Result(), f.Result())
	case gen.PH_RESULT:
		return fmt.Sprintf("%2d - %2d = (   )", f.a.Result(), f.b.Result())
	}

	return ""
}
