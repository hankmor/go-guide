package main

import "fmt"

func main() {
	// 分号是多余的
	fmt.Println("Redundant semicolon") // ;
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered")
		}
	}()
}
