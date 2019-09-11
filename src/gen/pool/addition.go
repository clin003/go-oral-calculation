package pool

import (
	"gen"
	"gen/factor"
)

func NewAddition(min int, bound int) gen.IPool {
	a := &addition{}
	a.init(min, bound)

	return a
}

type addition struct {
	basepool
}

func (a *addition) init(min int, bound int) {
	a.items = make([]gen.IFactor, 0)
	for i := min; i <= bound; i++ {
		for j := min; j <= (bound - i); j++ {
			a.items = append(a.items, factor.NewAdd(factor.NewInteger(i), factor.NewInteger(j)))
		}
	}
}
