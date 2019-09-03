package gen

import (
	"container/list"
	"fmt"
)

func NewFactor(value int) IFactor {
	return &Factor{value: value, result: value, stack: list.New()}
}

type Factor struct {
	value  int
	result int
	stack  *list.List
}

func (f *Factor) Append(factor IFactor, operator IOperator) IFactor {
	f.result = operator.Calc(f, factor)
	f.stack.PushBack(operator)
	f.stack.PushBack(factor)

	return f
}

func (f *Factor) Result() int {
	return f.result
}

func (f *Factor) ToString() string {
	format := fmt.Sprintf("%d", f.value)
	for el := f.stack.Front(); el != nil; el = el.Next() {
		format += el.Value.(IFormator).ToString()
	}

	return format
}
