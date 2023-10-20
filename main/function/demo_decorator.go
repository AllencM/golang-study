package function

import (
	"fmt"
	"time"
)

// 装饰器模式
// go 不支持注解，无法像java一样通过注解来实现装饰器模式

// MultiplyFunc 为函数类型设置别名提高代码可读性
type MultiplyFunc func(int, int) int

// 乘法运算函数1（算术运算）
func multiply(a, b int) int {
	return a * b
}

// 乘法运算函数2（位运算）
func multiply2(a, b int) int {
	return a << b
}

func execTime(f MultiplyFunc) MultiplyFunc {
	return func(a int, b int) int {
		start := time.Now()
		c := f(a, b)
		end := time.Since(start)
		fmt.Println("exec cost:", end)
		return c
	}
}

func decoratorFunc() {
	a := 2
	b := 3
	decorator := execTime(multiply)
	c := decorator(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)
}

// 对比位运算和算术运算
func compareFunc() {
	a := 2
	b := 8
	fmt.Println("算术运算:")
	decorator1 := execTime(multiply)
	c := decorator1(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)

	fmt.Println("位运算")
	decorator2 := execTime(multiply2)
	a = 1
	b = 4
	c = decorator2(a, b)
	fmt.Printf("%d << %d = %d\n", a, b, c)
}
