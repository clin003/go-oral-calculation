package factor

import (
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

func (f *add) ToString() string {
	return f.a.ToString() + " + " + f.b.ToString()
}
