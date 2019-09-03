package exporter

import (
	"gen"
)

type baseexporter struct {
	items []gen.IFactor
}

func (e *baseexporter) Export(cols int) {
}
