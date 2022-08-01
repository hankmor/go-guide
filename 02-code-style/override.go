package main

import (
	"errors"
	"fmt"
)

var data = make(map[int64]string)

func init() {
	data[1] = "zhangsan"
	data[2] = "lisi"
	data[3] = "huzhou"
}

func GetName(id int64) string {
	return data[id]
}

// 不能重载
// func GetName(id int64) (string, error) {
// }
// 不能重载
// func GetName(id int64, age int8) {
// }

func GetNameWithErr(id int64) (string, error) {
	name := data[id]
	if name != "" {
		return "", errors.New("not found")
	}
	return name, nil
}

func GetNameByIdAndAge(id int64, age int8) string {
	return ""
}

func main() {
	name := GetName(1)
	fmt.Println(name)
}
