package main

import "fmt"

func main() {
	// 单行注释
	fmt.Println("hello")
	/*
		多行注释1
		多行注释2
	*/
	sayHello("huzhou")
}

// sayHello is a function to say hello with the giving name.
func sayHello(name string) {
	fmt.Println("hello, ", name)
}
