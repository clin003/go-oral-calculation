package gen

import (
	"math/rand"
	"time"
)

type FormatType byte

// Where to use placeholder instead.
const (
	PH_LEFT   FormatType = 0
	PH_RIGHT  FormatType = 1
	PH_RESULT FormatType = 2
	PH_RAND   FormatType = 3
)

// FoFixFormat returns the fixed format type in PH_LEFT, PH_RIGHT or PH_RESULT.
func ToFixFormat(t FormatType) FormatType {
	if t != PH_RAND {
		return t
	}

	return FormatType(rand.Intn(3))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
