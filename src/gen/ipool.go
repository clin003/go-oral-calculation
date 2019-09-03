package gen

type IPool interface {
	Rand(num int) []IFactor
	Items() []IFactor
}
