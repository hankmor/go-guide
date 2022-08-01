package main

import "fmt"

func main() {
	// 变量err未使用，编译错误
	// name, err := getName()
	// 使用_忽略err
	name, _ := getName()
	fmt.Println(name)
}

func getName() (string, error) {
	return "huzhou", nil
}
