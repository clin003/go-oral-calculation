package operator

import (
	"gen"
)

func NewSubtract() gen.IOperator {
	return &subtract{}
}

type subtract struct {
}

func (s *subtract) Calc(a gen.IFactor, b gen.IFactor) int {
	return a.Result() - b.Result()
}

func (s *subtract) ToString() string {
	return "-"
}
