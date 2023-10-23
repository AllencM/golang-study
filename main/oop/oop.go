package oop

import "fmt"

// go 的面向对象，并没有类的继承、接口的实现、构造函数等，也没有 private、protected、private之类的访问修饰符
// 与java相比，go中的大多数类型都是值语义，包括：
// 基本类型：bool、int、float、string
// 复合类型：array、struct（slice、map、pointer、pipeline）
// 这里的值语义和引用语义等价于之前介绍类型时提到的值类型和引用类型。

// Integer 为值类型定义成员方法，将基本类型通过 type 关键字设置为新的类型
type Integer int

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func integerEqual() {
	var x Integer
	var y Integer
	x, y = 10, 15
	fmt.Println(x.Equal(y))
}

//  给 Integer 类型添加更多方法, 这样一来就将基本的数字类型转成了面向对象类型
func (a Integer) Add(b Integer) Integer {
	return a + b
}
func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

func integerOop() {
	var x Integer
	var y Integer
	x, y = 10, 15

	fmt.Println(x.Equal(y))
	fmt.Println(x.Add(y))
	fmt.Println(x.Multiply(y))
}

// 接口实现
// go 不支持传统面向对象编程中的继承和实现语法
// 如果有下面这接口的存在，则认为Integer 实现了math interface 无需显示声明
type Math interface {
	Add(i Integer) Integer
	Multiply(i Integer) Integer
}
