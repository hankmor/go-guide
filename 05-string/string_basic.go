package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串基本操作

var s = "hello,中国"

func stringBasic() {
	// 取字符串长度
	var l = len(s)
	fmt.Println(l) // 12

	// 获取字符串中的第 i 个字符
	var c5 = s[4] // o
	var c6 = s[5] // ,
	fmt.Println(string(c5), string(c6))
	// 试图获取超过字符串长度的字符则 panic
	// fmt.Println(string(s[13])) // panic: runtime error: index out of range [13] with length 12

	// go 中字符串都以 UTF8 字符集存储，一个中文占3个字节
	// 单独获取第一个中文字符的第一个字节会乱码
	var c7 = s[6]
	fmt.Println(c7, string(c7)) // 228 ä
	// 可以用字符串切片来获取
	fmt.Println(s[6:9]) // 中
}

func stringCompare() {
	// 字符串是可以按照自然顺序进行比较的，即可以用 >、<、= 等来判断
	var s1 = "hello"
	var s2 = "world"
	var s3 = "he" + "llo"
	fmt.Println(s1 < s2)       // true
	fmt.Println(s3 == s1)      // true
	fmt.Println(s1 == "hello") // true
	// 其实s1、s3 和 hello 是相同的，底层公用相同的字节数组
}

func stringImmutable() {
	// 字符串一旦创建，其底层的字节数组就不可更改，这并不是说不能赋值给其他变量
	var (
		s1 = "hello"
		s2 = s1
		s3 = s1 + ",world"
	)
	fmt.Println(s1, s2, s3) // hello hello hello,world

	// 试图修改底层字节数组会导致 panic
	// s1[0] = 'w' // cannot assign to s1[0] (value of type byte)
}

func stringUTF8() {
	// go默认的字符集是 UTF8，这样可以很方便的处理中文
	var s = "中国加油"     // 原字符串
	var rs = []rune(s) // utf8编码的字符切片
	var bs = []byte(s) // 字节切片
	fmt.Printf("原 => %4s => %6s\n", "UTF8", "UNICODE")
	for i, v := range rs {
		var utf8Bytes []byte
		for j := i * 3; j < (i+1)*3; j++ {
			utf8Bytes = append(utf8Bytes, bs[j])
		}
		fmt.Printf("%s => %X => %X\n", string(v), v, utf8Bytes)
	}
	/*output:
	中 => 4E2D => E4B8AD
	国 => 56FD => E59BBD
	加 => 52A0 => E58AA0
	油 => 6CB9 => E6B2B9
	*/
}

func stringUTF8Char() {
	// 由于 go 是UTF8字符集，所以不需要判断字符集，可以直接处理诸如前缀、后缀、包含等
	var prefix = func(s string, prefix string) bool {
		return len(s) >= len(prefix) && s[:len(prefix)] == prefix
	}
	var suffix = func(s string, suffix string) bool {
		return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
	}
	var contains = func(s string, substr string) bool {
		// return strings.Index(s, substr) > 0
		for i := 0; i < len(s); i++ { // 按前缀判断
			if prefix(s[i:], substr) {
				return true
			}
		}
		return false
	}
	var s1 = "hello, world"
	fmt.Println("prefix: ", prefix(s1, "hello"))
	fmt.Println("suffix: ", suffix(s1, "world"))
	fmt.Println("contain: ", contains(s1, "o, w"))
	var s2 = "hello, china!你好，中国！"
	fmt.Println("prefix: ", prefix(s2, "hello"))
	fmt.Println("suffix: ", suffix(s2, "国！"))
	fmt.Println("contains: ", contains(s2, "你好"))
}

func stringUTF8Decode() {
	// len可以直接获取字符串的长度，但是这个长度是按字节计算的，如何需要获取utf8编码后的字符长度，即 []rune 长度?
	fmt.Println(s, " len: ", len(s))
	// 可以转为 []rune 后获取
	var rs = []rune(s)
	fmt.Println(len(rs))
	// go 还提供直接获取的方法
	fmt.Println(utf8.RuneCountInString(s))

	// 解码 UTF8 字符
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:]) // 解码给定字符串参数中的 rune，返回 rune 字符和其对应的字节长度
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// for ... range 包含隐式的 UTF8 字符解码操作，简化了手动解码过程
	for i, v := range s {
		fmt.Printf("%d\t%c\t%X\n", i, v, v)
	}
}

func main() {
	stringBasic()
	stringCompare()
	stringImmutable()
	stringUTF8()
	stringUTF8Char()
	stringUTF8Decode()
}
