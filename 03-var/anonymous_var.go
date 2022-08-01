package main

import "fmt"

func main() {
	// 使用匿名变量

	a := 10
	b := 2
	mul, div, mod := calc(a, b)
	fmt.Printf("multiply = %d, divide = %d, mod = %d \n", mul, div, mod)

	_, _, mod1 := calc(a, b) // 使用匿名变量忽略不需要的返回值
	fmt.Printf("mod = %d \n", mod1)
}

func calc(a int, b int) (int, int, int) {
	return a * b, a / b, a % b
}
