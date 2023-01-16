package main

import "fmt"

func declareConst() {
	// 单行声明常量
	const c1 = 1
	const c2 = -100
	const c3 = "hello"

	// 组合声明多个常量
	const (
		c4 = 3.14
		c5 = 1.2 + 12i // 复数
	)
	// c1 = 2 // 编译失败，常量不可能更改：cannot assign to c1 (untyped int constant 1)

	const (
		m = 1     // 1
		n         // 1
		k         // 1
		l = m + 1 // 2
		i         // 2
	)
	fmt.Println(m, n, k, l, i)
}

func add() {
	var x int = 1
	var y int32 = 2
	// fmt.Println(x + y) // 编译失败：invalid operation: x + y (mismatched types int and int32)
	fmt.Println(x + int(y))
}

type MyInt int

func add2() {
	var x int = 1
	var y MyInt = 2
	// fmt.Println(x + y) // 编译失败
	fmt.Println(MyInt(x) + y)
}

func area() {
	const pi float32 = 3.1415926
	// 有类型常量
	const r int = 4
	// const area = PI * r * r // 编译失败: invalid operation: PI * r (mismatched types float32 and int)
	const area = pi * float32(r) * float32(r)
	fmt.Println("r = ", r, ", area = ", area)
}

func area2() {
	const pi = 3.1415926
	const r = 4
	const area = pi * r * r
	fmt.Printf("r type: %T, pi type: %T, area type: %T\n", r, pi, area)
	fmt.Println("r = ", r, ", area = ", area)
}

// 定义枚举
const (
	// Monday    = 0
	// Tuesday   = 1
	// Wednesday = 2
	// Thursday  = 3
	// Friday    = 4
	// Saturday  = 5
	// Sunday    = 6
	// weekDays  = 7

	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	weekDays
)

func printWeek() {

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Println("week day number: ", weekDays)
}

func iotaBasic() {
	const (
		x = -1   // -1, iota = 0
		a = 1    // 1, iota = 1
		b        // 1, iota = 2
		c = iota // iota = 3，第一个表达是 x = -1 开始计算 iota 的值，故此时 iota 为 3
		d        // iota = 4
	)
	fmt.Println(x, a, b, c, d)

	// 每次使用 const 时，iota 的值会重置为0
	const c1 = iota
	const c2 = iota
	fmt.Println(c1, c2) // 0 0

	const (
		e, f = iota, iota + 10 // 0, 10, 同一行，无论 iota 重复多少次，其值都是一样的
		g, h                   // 1, 11
	)
	fmt.Println(e, f, g, h)
}

func main() {
	declareConst()
	add()
	add2()
	area()
	area2()
	printWeek()

	iotaBasic()

	checkWorkday(1)
	checkWorkday(10)

	// checkWorkday(Monday)
}

// 参数为 int，不能限制调用者

func checkWorkday(day int) {
	if day == Saturday || day == Sunday {
		fmt.Println("(〃'▽'〃), good weekend!")
	} else {
		fmt.Println("(灬ꈍ ꈍ灬)，it's workday!")
	}
}

// 定义类型安全的枚举
//
// type WeekDay int
//
// const (
// 	Monday WeekDay = iota
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// 	Sunday
// 	weekDays
// )
//
// func checkWorkday(day WeekDay) {
// 	if day == Saturday || day == Sunday {
// 		fmt.Println("(〃'▽'〃), good weekend!")
// 	} else {
// 		fmt.Println("(灬ꈍ ꈍ灬)，it's workday!")
// 	}
// }
