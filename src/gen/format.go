package gen

import (
	"math/rand"
	"time"
)

type FormatType byte

// Where to use placeholder to replace the number.
const (
	PH_NONE     FormatType = 0 // No placeholder in the formula.
	PH_LEFT     FormatType = 1 // Replace one in left side of the formula with a placeholder.
	PH_RIGHT    FormatType = 2 // Replace the formula result with a placeholder.
	PH_RANDL5R5 FormatType = 3 // Random the placeholder between left or right side.
	PH_RANDL8R2 FormatType = 4 //
	PH_RANDL2R8 FormatType = 5 //
)

// DecideFormat returns a const format type in PH_NONE, PH_LEFT and PH_RIGHT.
func DecideFormat(t FormatType) FormatType {
	switch t {
	case PH_NONE, PH_LEFT, PH_RIGHT:
		return t
	case PH_RANDL5R5:
		return FormatType(1 + rand.Intn(2))
	case PH_RANDL8R2:
		if rand.Float32() < 0.8 {
			return PH_LEFT
		} else {
			return PH_RIGHT
		}
	case PH_RANDL2R8:
		if rand.Float32() < 0.8 {
			return PH_RIGHT
		} else {
			return PH_LEFT
		}
	}

	return PH_NONE
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
