package main

import (
	"fmt"
	"gen/pool"
)

func main() {
	addition := pool.NewSubtraction(1, 20)
	result := addition.Rand(20)
	fmt.Printf("--- 共生成[%2d]道题 ---\n", len(result))
	for _, v := range result {
		fmt.Println(v.ToString())
	}
}
