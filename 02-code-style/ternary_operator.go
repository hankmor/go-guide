package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Name string
	Age  int
}

func NewUser() *User {
	return &User{}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	age := rand.Intn(200)

	user := NewUser()
	user.Name = "huzhou"
	// go中没有三元运算符
	// user.Age = validAge(age) ? age: -1
	// 使用if语句代替三元运算符
	if validAge(age) {
		user.Age = age
	} else {
		user.Age = -1
	}
	// 自定义三元运算符
	ret := Choose(validAge(age), age, -1)
	user.Age = ret.(int)
	fmt.Printf("%v\n", user)

	user.Age = Choose1(validAge(age), age, -1)
	fmt.Printf("%v\n", user)
}

func validAge(age int) bool {
	return age > 0 && age < 130
}

func Choose(b bool, v1 any, v2 any) any {
	if b {
		return v1
	}
	return v2
}

func Choose1[T any](b bool, v1 T, v2 T) T {
	if b {
		return v1
	}
	return v2
}
