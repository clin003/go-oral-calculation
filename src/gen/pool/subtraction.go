package pool

import (
	"gen"
	"gen/factor"
)

func NewSubtraction(min int, max int) gen.IPool {
	s := &subtraction{}
	s.init(min, max)

	return s
}

type subtraction struct {
	basepool
}

func (s *subtraction) init(min int, max int) {
	s.items = make([]gen.IFactor, 0)
	for i := max; i >= min; i-- {
		for j := i - 1; j >= min; j-- {
			s.items = append(s.items, factor.NewSub(factor.NewInteger(i), factor.NewInteger(j)))
		}
	}
}
