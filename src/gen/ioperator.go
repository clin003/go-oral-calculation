package gen

type IOperator interface {
	Calc(f1 IFactor, f2 IFactor) int
	ToString() string
}
