package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello go!")
	for i, arg := range os.Args { // os.Args 获取命令行传递的参数，第一个始终为程序的名称
		fmt.Println(i, "=", arg)
	}
}
