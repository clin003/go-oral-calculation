package factor

import (
	"fmt"
	"gen"
)

func NewDiv(a gen.IFactor, b gen.IFactor) gen.IFactor {
	return &div{a: a, b: b}
}

type div struct {
	a gen.IFactor
	b gen.IFactor
}

func (f *div) Result() int {
	return f.a.Result() / f.b.Result()
}

func (f *div) Format(t gen.FormatType, withResult bool) string {
	formula := ""

	t = gen.DecideFormat(t)
	if t == gen.PH_LEFT {
		if Poss(0.5) {
			formula += fmt.Sprintf("%s ÷ %s", f.a.Format(t, false), f.b.Format(gen.PH_NONE, false))
		} else {
			formula += fmt.Sprintf("%s ÷ %s", f.a.Format(gen.PH_NONE, false), f.b.Format(t, false))
		}
	} else {
		formula += fmt.Sprintf("%s ÷ %s", f.a.Format(gen.PH_NONE, false), f.b.Format(gen.PH_NONE, false))
	}

	if withResult {
		if t == gen.PH_RIGHT {
			formula += " ＝ (   )"
		} else {
			formula += fmt.Sprintf(" ＝ %d", f.Result())
		}
	}

	return formula
}
