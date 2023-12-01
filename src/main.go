package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/clin003/kousuan/gen"
	"github.com/clin003/kousuan/gen/exporter"
	"github.com/clin003/kousuan/gen/pool"
)

func main() {
	printFun()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入标号（按回车继续，输入 'q' 退出）：")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if strings.ToLower(userInput) == "q" {
			fmt.Println("程序结束.")
			return
		}
		if len(userInput) > 0 {
			// 尝试将用户输入的内容转换为 uint64 类型
			number, err := strconv.ParseUint(userInput, 10, 8)
			if err != nil {
				fmt.Println("输入内容不是有效的 0 到 5 之间的数字:", err)
				continue
			}
			if number > 5 {
				fmt.Println("输入的数字超出范围，请输入 0 到 5 之间的数字")
				continue
			}
			formatType := uint8(number)
			genFun(gen.FormatType(formatType))
			fmt.Println("本次任务完成")
		} else {
			printFun()
		}

	}

}

func printFun() {
	fmt.Println("生成格式:")
	fmt.Println("0 无留空")
	fmt.Println("1 左侧留空")
	fmt.Println("2 右侧留空")
	fmt.Println("3 随机留空(左5:右5)")
	fmt.Println("4 随机留空(左8:右2)")
	fmt.Println("5 随机留空(左2:右8)")
}

func genFun(formatType gen.FormatType) {
	ext, _ := os.Executable()

	addition := pool.NewMix(pool.NewAddition(1, 20), pool.NewSubtraction(1, 20))
	result := addition.Rand(100)
	fmt.Printf("--- 共生成[%d]道题 ---\n", len(result))
	// exp := exporter.NewConsole(result)
	// exp.Export(formatType, 2, false)
	fileName := path.Join(
		filepath.Dir(ext),
		fmt.Sprintf("result_%s.xlsx", time.Now().Format(time.DateOnly)),
	)
	exp := exporter.NewExcel(
		result,
		fileName,
	)
	exp.Export(formatType, 4, false)
	fmt.Println("生成内容存储到:", fileName)
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
