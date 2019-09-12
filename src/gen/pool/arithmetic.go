package pool

import (
	"gen"
)

type OP byte

const (
	ADD OP = 1
	SUB OP = 2
	MUL OP = 3
	DIV OP = 4
)

func NewArithmetic(level uint, op OP, ops ...OP) gen.IPool {
	return &arithmetic{level: level, ops: append(ops, op)}
}

type arithmetic struct {
	basepool
	level uint
	ops   []OP
}
