package main

import "fmt"

func main() {
	// 变量赋值

	var v1 int
	fmt.Println("init value: ", v1)
	v1 = 1
	fmt.Println("assigned: ", v1)
	v2 := 100
	fmt.Println("init value: ", v1)
	v2 = 1
	fmt.Println("assigned: ", v2)

	i, j := 10, 100
	j, i = i, j // 多重赋值，快速交换i, j的值，不需要临时变量
	fmt.Printf("i = %d, j = %d \n", i, j)
}
