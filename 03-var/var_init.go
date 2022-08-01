package main

import "fmt"

func main() {
	// 变量的初始化
	var v1 int
	fmt.Println(v1) // 0
	var f1 float64
	fmt.Println(f1) // 0
	var v2 string
	fmt.Println(v2) // ""
	var v3 map[string]int
	fmt.Println(v3) // map[]
	var v4 *map[string]int
	fmt.Println(v4) // nil

	// 数组
	v5 := [3]int{1, 2, 3}
	fmt.Println(v5) // [1 2 3]

	// 切片
	v6 := []int{1, 2, 3}
	fmt.Println(v6) // [1 2 3]
}
