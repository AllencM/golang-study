package main

import (
	"fmt"
	"strconv"
)

// 字符转int， 等于 parseInt
func testAtoi() {
	v1 := "100"
	v2, _ := strconv.Atoi(v1)
	fmt.Println(v2)

	// 整型转化成string
	v3 := 100
	fmt.Println(strconv.Itoa(v3))

	// string->bool
	v5 := "true"
	v6, _ := strconv.ParseBool(v5)
	fmt.Println(v6)

	v7 := "100"
	// string -> int, 第二个参数表示进制，第三个表示最大位数
	v8, _ := strconv.ParseInt(v7, 10, 64)
	fmt.Println(v8)

	var v9 int64 = 3
	// int64 -> string, 第二个参数是进制
	v10 := strconv.FormatInt(v9, 2)
	fmt.Println(v10)

	v11 := "99.999999111111"
	//  字符串转浮点型，第二个参数为精度 float32 / float64
	v12, _ := strconv.ParseFloat(v11, 64)
	fmt.Println(v12)

	// 超过float32 精度，结果就不对了
	v13 := "99.1111111111111111111111111111111111111111"
	v14, _ := strconv.ParseFloat(v13, 32)
	v15, _ := strconv.ParseFloat(v13, 64)
	fmt.Println(v14)
	fmt.Println(v15)

	q := "hello，world"
	fmt.Println(q)
	fmt.Println(strconv.Quote(q))
	// 字符串转 ascii
	fmt.Println(strconv.QuoteToASCII(q))

}
