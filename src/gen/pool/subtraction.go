package pool

import (
	"github.com/clin003/kousuan/gen"
	"github.com/clin003/kousuan/gen/factor"
)

func NewSubtraction(min int, bound int) gen.IPool {
	s := &subtraction{}
	s.init(min, bound)

	return s
}

type subtraction struct {
	basepool
}

func (s *subtraction) init(min int, bound int) {
	s.items = make([]gen.IFactor, 0)
	for i := bound; i >= min; i-- {
		for j := i - 1; j >= min; j-- {
			s.items = append(s.items, factor.NewSub(factor.NewInteger(i), factor.NewInteger(j)))
		}
	}
}
