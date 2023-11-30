package exporter

import (
	"github.com/clin003/kousuan/gen"
)

type baseexporter struct {
	items []gen.IFactor
}

func (e *baseexporter) Export(t gen.FormatType, cols int) {
}
