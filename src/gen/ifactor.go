package gen

type IFactor interface {
	Append(factor IFactor, operator IOperator) IFactor
	Result() int
	ToString() string
}
