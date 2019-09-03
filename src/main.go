package main

import (
	"fmt"
	"gen/exporter"
	"gen/pool"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	addition := pool.NewMix(pool.NewAddition(1, 20), pool.NewSubtraction(1, 20))
	result := addition.Rand(100)

	fmt.Printf("--- 共生成[%2d]道题 ---\n", len(result))
	exp := exporter.NewConsole(result)
	exp.Export(4)

	ext, _ := os.Executable()
	exp = exporter.NewCsv(result, path.Join(filepath.Dir(ext), fmt.Sprintf("result_%d.csv", time.Now().Unix())))
	exp.Export(4)
}
