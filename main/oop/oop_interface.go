package oop

import "fmt"

// go 没有提供 implement 关键字
// 一个类只要实现了某个接口要求的所有方法，就可以当作这个类实现了该接口

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func main1() {
	// 编译通过
	var num1 Number = 1
	var num2 Number1 = num1
	var num3 Number2 = num2
	fmt.Println(num3)

	// 任何实现了 Number1 接口的类，也实现了 Number2
	// 任何实现了 Number1 接口的类实例都可以赋值给 Number2，反之亦然；
	// 在任何地方使用 Number1 接口与使用 Number2 并无差异
}

// 方法子集
// 接口赋值并不要求两个接口完全等价（方法完全相同）。如果接口 A 的方法列表是接口 B 的方法列表的子集，那么接口 B 也可以赋值给接口 A
// Number1 Number2 的方法列表是Number3的方法列表的子集
// Number3 可以复制给 Number1 2

type Number3 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
	Add(i int)
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}
