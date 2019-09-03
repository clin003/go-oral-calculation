package operator

import (
	"gen"
)

func NewAdd() gen.IOperator {
	return &add{}
}

type add struct {
}

func (a *add) Calc(f1 gen.IFactor, f2 gen.IFactor) int {
	return f1.Result() + f2.Result()
}

func (a *add) ToString() string {
	return "+"
}
