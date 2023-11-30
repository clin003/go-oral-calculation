package pool

import (
	"math/rand"
	"time"

	"github.com/clin003/kousuan/gen"
)

type basepool struct {
	items []gen.IFactor
}

func (b *basepool) Rand(num int) []gen.IFactor {
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
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
