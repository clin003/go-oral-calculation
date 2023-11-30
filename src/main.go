package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/clin003/kousuan/gen"
	"github.com/clin003/kousuan/gen/exporter"
	"github.com/clin003/kousuan/gen/pool"
)

func main() {
	printFun()
	var formatType uint8
	for {

		if n, err := fmt.Scanln(&formatType); err != nil {
			fmt.Println(n, "err:", err)
			fmt.Print("重新输入:")
			continue
		} else {
			if formatType < 6 && formatType >= 0 {
				genFun(gen.FormatType(formatType))
				fmt.Println("生成完成")

				printFun()
				continue
			} else {
				fmt.Print("输入数字非法,重新输入:")
			}
		}
	}

}

func printFun() {
	fmt.Println("生成格式:")
	fmt.Println("0 无留空")
	fmt.Println("1 左侧留空")
	fmt.Println("2 右侧留空")
	fmt.Println("3 随机留空")
	fmt.Println("4 随机留空")
	fmt.Println("5 随机留空")
	fmt.Println("请输入生成格式:")
}

func genFun(formatType gen.FormatType) {
	ext, _ := os.Executable()

	addition := pool.NewMix(pool.NewAddition(1, 10), pool.NewSubtraction(1, 10))
	result := addition.Rand(90)
	fmt.Printf("--- 共生成[%d]道题 ---\n", len(result))
	exp := exporter.NewConsole(result)
	exp.Export(formatType, 3, false)
	exp = exporter.NewText(
		result, path.Join(
			filepath.Dir(ext),
			fmt.Sprintf("result_r_%s.text", time.Now().Format(time.DateOnly)),
		),
	)
	exp.Export(formatType, 3, false)
}

func genForFun(formatType gen.FormatType) {
	ext, _ := os.Executable()
	for i := 0; i < 30; i++ {
		addition := pool.NewMix(pool.NewAddition(1, 10), pool.NewSubtraction(1, 10))
		result := addition.Rand(100)
		exp := exporter.NewText(
			result, path.Join(
				filepath.Dir(ext), fmt.Sprintf("result_l_%d.txt", i+1),
			),
		)
		exp.Export(gen.PH_RIGHT, 3, false)

		addition = pool.NewMix(pool.NewAddition(1, 10), pool.NewSubtraction(1, 10))
		result = addition.Rand(100)
		exp = exporter.NewText(
			result, path.Join(
				filepath.Dir(ext), fmt.Sprintf("result_r_%d.txt", i+1),
			),
		)
		exp.Export(gen.PH_LEFT, 3, false)

		addition = pool.NewMix(pool.NewAddition(1, 10), pool.NewSubtraction(1, 10))
		result = addition.Rand(100)
		exp = exporter.NewText(
			result, path.Join(
				filepath.Dir(ext), fmt.Sprintf("result_m_%d.txt", i+1),
			),
		)
		exp.Export(gen.PH_LEFT, 3, false)
	}
}
