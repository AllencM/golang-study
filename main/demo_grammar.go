package main

import (
	"fmt"
)

func main() {
	// testForLoop()
	// testType()
	// testFloat()
	// testFor()
	testAtoi()
}

func testForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func testType() {
	var intValue int8
	intValue2 := 8
	intValue = int8(intValue2)
	fmt.Println(intValue)
}

func testFloat() {
	// float32 等于java 的float类型,单精度浮点数，小数点后7位
	var f1 float32 = 1.4
	// float64 等于java 的 double 类型，双精度浮点数，小数点后15位
	var f2 = 2.123123123123
	var f3 = 1.1e-10
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)

	var f4 = 0.3
	var f5 = 0.7
	var f6 = int(f4 + f5)
	fmt.Println(f6)

	p1 := 1.000000000001
	p2 := 1.00000000001
	fmt.Println(p1 == p2)
}

func testString() {
	str := "what r u fucking"
	fmt.Println(str)
	fmt.Printf("the length of \"%s\" is %d\n", str, len(str))

	// 多行字符
	label := `Search result For "Golang":
	- GO
	- GOLANG
     GOLANG programing
	`
	fmt.Println(label)

	label2 := "123" + "123"
	fmt.Println(label2)

}

func testFirstChar() {
	str := "\u0000what r u fuck talking about"
	fmt.Printf("tht first character of %s is %c \n", str, str[0])

	runes := []rune(str)
	fmt.Printf("tht first character of %s is %c", str, runes[0])
}

func testSlice() {
	str := "testGoSlice"
	// 前0-5(不含)个字符
	str1 := str[:5]
	// 第7个(包含)到最后
	str2 := str[7:]
	// 第5个到第7个（不含）
	str3 := str[5:7]
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str[:])
}

func testFor() {
	str := "hello the fucking world"
	// 字符串的字节长度
	n := len(str)
	// 字节数组遍历方式
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	// unicode 遍历字符方式，每个字符的类型是 rune，不是byte
	// rune 代表单个unicode 字符（int32的别名）
	// unicode 是一种字符集，包含了所有语言的所有字符，类似的还有 ascii、iso8859-1等,广义的 unicode 包含了字符集也包含了编码规则
	// uft-8是unicode 字符集的实现方式之一，将unicode字符以某种方式进行编码，utf-8 是一种变长的编码规则，1-4个字节不等
	for j, ch := range str {
		fmt.Println(j, ch)
	}

	// unicode 编码转化成可打印的字符
	// utf-8 编码不能这样转化，中文字符集会出现乱码，一个中文字符编码须要3个字节
	// range 关键字遍历字符串，是从 unicode 字符角度
	str1 := "hello the fucking 世界"
	for k, ch := range str1 {
		fmt.Println(k, string(ch))
	}

	for i := 0; i < len(str1); i++ {
		ch := str1[i]
		// 这样处理，会乱码，utf8下，一个中文是3个字节
		fmt.Println(i, string(ch))
	}

}
