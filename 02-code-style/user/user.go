package user

import "fmt"

var Name = "huzhou"

func getName() string {
	return Name
}

func SayHello(name string) {
	fmt.Println("hello, ", name)
}
