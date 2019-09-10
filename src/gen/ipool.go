package gen

// IPool interface.
type IPool interface {
	// Rand generates num factors from the pool.
	Rand(num int) []IFactor

	// Items returns the whole items of the pool.
	Items() []IFactor
}
