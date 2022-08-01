package main

import "fmt"

func main() {
	// 变量声明过后，必须被使用，否则编译错误

	// 变量声明：var 变量名 变量类型
	var v0 rune                                  // 声明字符类型的变量 v0
	var v1 int                                   // 声明整型变量 v1
	var v2 string                                // 声明字符串类型的变量 v2
	v0 = 'A'                                     // 给变量赋初始值
	fmt.Println("v0=", v0, "v1=", v1, "v2=", v2) // ~: v0= 65 v1= 0 v2=

	var i1 = 1 // 自动类型推导
	fmt.Printf("i1 = %d, type: %T \n", i1, i1)

	// 可以将多个变量合并声明
	var (
		v3 [2]int          // 数组
		v4 []int           // 数组切片
		v5 struct{ i int } // 结构体
		v6 *int            // 指针
		v7 map[string]int  // map
		// v8 func(a int, b int) int // 匿名函数
	)
	fmt.Println("v3=", v3, "v4=", v4, "v5=", v5, "v6=", v6, "v7=", v7, "v8=")
	// ~: v3=[0 0] v4=[] v5={0} v6=<nil> v7=map[] v8=<nil>

	v9 := 10 // 声明变量 v9 并初始化，自动推导类型
	fmt.Printf("v9 = %d, type: %T \n", v9, v9)

	a, b := func1() // 直接声明变量并接收方法的返回值
	fmt.Println(a, b)

	var v10 = "hello"
	// v10 := "go" // 编译失败：no new variables
	fmt.Println(v10)
}

func func1() (int, int) {
	return 1, 2
}
