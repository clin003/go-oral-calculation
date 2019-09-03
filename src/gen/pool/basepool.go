package pool

import (
	"gen"
	"math/rand"
	"time"
)

type basepool struct {
	items []gen.IFactor
}

func (b *basepool) Rand(num int) []gen.IFactor {
	rand.Seed(time.Now().UnixNano())
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

func (b *basepool) Items() []gen.IFactor {
	return b.items
}
