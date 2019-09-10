package pool

import (
	"gen"
	"gen/factor"
)

func NewAddition(min int, max int) gen.IPool {
	a := &addition{}
	a.init(min, max)

	return a
}

type addition struct {
	basepool
}

func (a *addition) init(min int, max int) {
	a.items = make([]gen.IFactor, 0)
	for i := min; i <= max; i++ {
		for j := min; j <= (max - i); j++ {
			a.items = append(a.items, factor.NewAdd(factor.NewInteger(i), factor.NewInteger(j)))
		}
	}
}
