package factor

import (
	"math/rand"
	"time"
)

func Poss(p float32) bool {
	return rand.Float32() < p
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
