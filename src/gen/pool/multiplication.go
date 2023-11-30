package pool

import (
	"github.com/clin003/kousuan/gen"
	"github.com/clin003/kousuan/gen/factor"
)

func NewMultiplication(min int, max int) gen.IPool {
	m := &multiplication{}
	m.init(min, max)

	return m
}

type multiplication struct {
	basepool
}

func (m *multiplication) init(min int, max int) {
	m.items = make([]gen.IFactor, 0)
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			m.items = append(m.items, factor.NewMul(factor.NewInteger(i), factor.NewInteger(j)))
		}
	}
}
