package pool

import (
	"gen"
	"math/rand"
)

type basepool struct {
	items []gen.IFactor
}

func (b *basepool) Rand(num int) []gen.IFactor {
	result := make([]gen.IFactor, 0)
	p := b.items[:]
	for i := 0; i < num; i++ {
		if len(p) == 0 {
			break
		}

		index := rand.Intn(len(p))
		result = append(result, p[index])
		p = append(p[:index], p[index+1:]...)
	}

	return result
}
