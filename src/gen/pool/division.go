package pool

import (
	"gen"
	"gen/factor"
)

func NewDivision(min int, max int) gen.IPool {
	m := &division{}
	m.init(min, max)

	return m
}

type division struct {
	basepool
}

func (d *division) init(min int, max int) {
	d.items = make([]gen.IFactor, 0)
	for i := min; i <= max; i++ {
		for j := min; j <= i; j++ {
			if i%j == 0 {
				d.items = append(d.items, factor.NewDiv(factor.NewInteger(i), factor.NewInteger(j)))
			}
		}
	}
}
