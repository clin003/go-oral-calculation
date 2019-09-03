package exporter

import (
	"fmt"
	"gen"
)

func NewConsole(items []gen.IFactor) gen.IExporter {
	c := &console{}
	c.items = items

	return c
}

type console struct {
	baseexporter
}

func (c *console) Export(cols int) {
	for i, it := range c.items {
		fmt.Printf("(%d) %s =\t", i+1, it.ToString())
		if cols <= 1 || (i+1)%cols == 0 {
			fmt.Print("\n")
		}
	}
}
