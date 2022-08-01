package main

import (
	"fmt"
	"github.com/huzhouv/go-guide/user"
)

func main() {
	name := user.Name
	fmt.Println(name)

	// name = user.getName()
	// fmt.Println(name)

	user.SayHello("huzhou")
}
