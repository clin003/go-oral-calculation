package exporter

import (
	"fmt"

	"github.com/clin003/kousuan/gen"
)

func NewConsole(items []gen.IFactor) gen.IExporter {
	c := &console{}
	c.items = items

	return c
}

type console struct {
	baseexporter
}

func (c *console) Export(t gen.FormatType, cols int, indexed bool) {
	for i, it := range c.items {
		if indexed {
			fmt.Printf("[%d] %s\t", i+1, it.Format(t, true))
		} else {
			fmt.Printf("   %s\t", it.Format(t, true))
		}

		if cols <= 1 || (i+1)%cols == 0 {
			fmt.Print("\n")
		}
	}
}
