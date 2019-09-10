package factor

import (
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

func (f *sub) ToString() string {
	return f.a.ToString() + " - " + f.b.ToString()
}
