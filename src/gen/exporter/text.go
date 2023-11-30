package exporter

import (
	"fmt"
	"os"
	"time"

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
	f.WriteString(
		fmt.Sprintf(
			"%d 年___月___日 \t\t 姓名_________ 班级_________\n\n\n",
			time.Now().Year(),
		),
	)

	for i, it := range c.items {

		if cols <= 1 || (i+1)%cols == 0 {
			if indexed {
				f.WriteString(fmt.Sprintf("[%d]%s\n", i+1, it.Format(t, true)))
			} else {
				f.WriteString(fmt.Sprintf("%s\n", it.Format(t, true)))
			}
			// f.WriteString(fmt.Sprintf("\n"))
		} else {
			if indexed {
				f.WriteString(fmt.Sprintf("[%d]%s\t\t", i+1, it.Format(t, true)))
			} else {
				f.WriteString(fmt.Sprintf("%s\t\t", it.Format(t, true)))
			}
		}
	}
}
