package exporter

import (
	"gen"
)

type baseexporter struct {
	items []gen.IFactor
}

func (e *baseexporter) Export(t gen.FormatType, cols int) {
}
