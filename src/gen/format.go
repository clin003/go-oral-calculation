package gen

import (
	"math/rand"
	"time"
)

type FormatType byte

// Where to use placeholder to replace the number.
const (
	PH_NONE  FormatType = 0 // No placeholder in the formula.
	PH_LEFT  FormatType = 1 // Replace one in left side of the formula with a placeholder.
	PH_RIGHT FormatType = 2 // Replace the formula result with a placeholder.
	PH_RAND  FormatType = 3 // Random the placeholder between left or right side.
)

// RandFormat returns a random format type in PH_LEFT, PH_RIGHT.
func RandFormat() FormatType {
	return FormatType(1 + rand.Intn(2))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
