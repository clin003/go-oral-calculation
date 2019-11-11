package main

import (
	"fmt"
	"gen"
	"gen/exporter"
	"gen/pool"
	"os"
	"path"
	"path/filepath"
	// "time"
)

func main() {
	ext, _ := os.Executable()

	// fmt.Printf("--- 共生成[%2d]道题 ---\n", len(result))
	// exp := exporter.NewConsole(result)
	// exp.Export(gen.PH_RAND, 4, false)

	for i := 0; i < 10; i++ {
		addition := pool.NewMix(pool.NewAddition(1, 10), pool.NewSubtraction(1, 10))
		result := addition.Rand(50)
		exp := exporter.NewCsv(result, path.Join(filepath.Dir(ext), fmt.Sprintf("result_l_%d.csv", i+1)))
		exp.Export(gen.PH_RIGHT, 2, false)

		addition = pool.NewMix(pool.NewAddition(1, 10), pool.NewSubtraction(1, 10))
		result = addition.Rand(50)
		exp = exporter.NewCsv(result, path.Join(filepath.Dir(ext), fmt.Sprintf("result_r_%d.csv", i+1)))
		exp.Export(gen.PH_RANDL, 2, false)
	}
}
