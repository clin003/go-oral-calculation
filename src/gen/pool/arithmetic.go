package pool

import (
	"gen"
)

type OpType byte

const (
	ADD OpType = 1
	SUB OpType = 2
	MUL OpType = 3
	DIV OpType = 4
)

func NewArithmetic(level int, op OpType, ops ...OpType) gen.IPool {
	a := &arithmetic{level: level, ops: append(ops, op)}
	a.init()

	return a
}

type arithmetic struct {
	basepool
	level int
	ops   []OpType
}

func (a *arithmetic) init() {
	if a.level <= 0 {
		a.level = 1
	}

	for i := 0; i < (a.level + 1); i++ {

	}
}
