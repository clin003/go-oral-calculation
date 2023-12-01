package exporter

import (
	"fmt"
	"time"

	"github.com/clin003/kousuan/gen"
	"github.com/xuri/excelize/v2"
)

func NewExcel(items []gen.IFactor, file string) gen.IExporter {
	c := &excel{file: file}
	c.items = items

	return c
}

type excel struct {
	baseexporter
	file string
}

func (c *excel) Export(t gen.FormatType, cols int, indexed bool) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 设置A4纸
	var (
		size          = 10
		blackAndWhite = true
	)
	if err := f.SetPageLayout("Sheet1", &excelize.PageLayoutOptions{
		Size:          &size,
		BlackAndWhite: &blackAndWhite,
	}); err != nil {
		fmt.Println(err)
	}

	if cols > 0 {
		// 设置单元格宽度
		if err := f.SetColWidth("Sheet1", "B", fmt.Sprintf("%c", cols+65), 20); err != nil {
			fmt.Println(err)
		}
		// 合并单元格，用于抬头
		if err := f.MergeCell("Sheet1", "B2", fmt.Sprintf("%c", cols+65)+"2"); err != nil {
			fmt.Println(err)
		}
		f.SetCellValue(
			"Sheet1", "B2",
			fmt.Sprintf(
				"%d 年___月___日 \t\t 姓名_________ 班级_________\n\n\n",
				time.Now().Year(),
			),
		)
		if err := f.SetRowHeight("Sheet1", 2, 30); err != nil {
			fmt.Println(err)
		}
	} else {
		return
	}

	// 隐藏网格线
	var (
		showGridLines = false
	)
	if err := f.SetSheetView("Sheet1", 0, &excelize.ViewOptions{
		ShowGridLines: &showGridLines,
	}); err != nil {
		fmt.Println(err)
	}

	cellIndex := 3  //起始单元格行（2）
	asciiIndex := 1 //起始单元格列（B）
	rowIndex := 2   //行索引
	for i, it := range c.items {
		if indexed {
			f.SetCellValue(
				"Sheet1",
				fmt.Sprintf("%c", asciiIndex+65)+fmt.Sprintf("%d", cellIndex),
				fmt.Sprintf("[%d]%s", i+1, it.Format(t, true)),
			)
		} else {
			f.SetCellValue(
				"Sheet1",
				fmt.Sprintf("%c", asciiIndex+65)+fmt.Sprintf("%d", cellIndex),
				fmt.Sprintf("%s", it.Format(t, true)),
			)
		}

		if cols <= 1 || (i+1)%cols == 0 {
			cellIndex = cellIndex + 1
			asciiIndex = 1
			rowIndex = rowIndex + 1
			if err := f.SetRowHeight("Sheet1", rowIndex, 28); err != nil {
				fmt.Println(err)
			}
		} else {
			asciiIndex = asciiIndex + 1
		}
	}

	if err := f.SaveAs(c.file); err != nil {
		fmt.Println(err)
	}
}
