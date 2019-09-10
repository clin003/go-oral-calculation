package exporter

import (
	"fmt"
	"gen"
	"os"
)

func NewCsv(items []gen.IFactor, file string) gen.IExporter {
	c := &csv{file: file}
	c.items = items

	return c
}

type csv struct {
	baseexporter
	file string
}

func (c *csv) Export(t gen.FormatType, cols int) {
	f, err := os.OpenFile(c.file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()

	for i, it := range c.items {
		f.WriteString(fmt.Sprintf("[%d] %s,", i+1, it.Format(t)))
		if cols <= 1 || (i+1)%cols == 0 {
			f.WriteString("\n")
		}
	}
}
