package pool

import (
	"gen"
)

func NewMix(pools ...gen.IPool) gen.IPool {
	m := &mix{}
	m.init(pools...)

	return m
}

type mix struct {
	basepool
}

func (m *mix) init(pools ...gen.IPool) {
	for _, p := range pools {
		m.items = append(m.items, p.Items()...)
	}
}
