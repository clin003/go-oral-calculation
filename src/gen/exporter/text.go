package exporter

import (
	"fmt"
	"os"

	"github.com/clin003/kousuan/gen"
)

func NewText(items []gen.IFactor, file string) gen.IExporter {
	c := &text{file: file}
	c.items = items

	return c
}

type text struct {
	baseexporter
	file string
}

func (c *text) Export(t gen.FormatType, cols int, indexed bool) {
	f, err := os.OpenFile(c.file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()

	for i, it := range c.items {
		if indexed {
			f.WriteString(fmt.Sprintf("[%d] %s\t", i+1, it.Format(t, true)))
		} else {
			f.WriteString(fmt.Sprintf("   %s\t", it.Format(t, true)))
		}

		if cols <= 1 || (i+1)%cols == 0 {
			f.WriteString(fmt.Sprintf("\n"))
		}
	}
}
