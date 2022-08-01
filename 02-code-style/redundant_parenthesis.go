package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 大括号不能换行
	// {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(100)
	// 括号是多余的
	// if (a%2 == 0) {
	if a%2 == 0 {
		fmt.Println(a, " is an even number")
	} else {
		fmt.Println(a, " is an odd number")
	}
}
