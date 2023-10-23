package oop

import "fmt"

// go 里面分位 值方法和 指针方法
// 指针方法只能通过指针调用

type Foo struct {
	name string
}

func (f *Foo) pointerMethod() {
	fmt.Println("pointer method on", f.name)
}

func (f Foo) valueMethod() {
	fmt.Println("value method on", f.name)
}

func NewFoo() Foo {
	return Foo{
		name: "right value struct",
	}
}

func GoFoo() {
	f1 := Foo{name: "value struct"}
	// 编译期会自动插入取地址符号，变为 (&f1).pointerMethod()
	f1.pointerMethod()
	f1.valueMethod()

	f2 := &Foo{name: "pointer struct"}
	f2.pointerMethod()
	// 编译期会自动解引用，变为(*f2).valueMethod
	f2.valueMethod()

	NewFoo().valueMethod()
	// 编译不通过，会提示 cannot take the address of NewFoo()
	// NewFoo().pointerMethod()
	// 可以被寻址的是左值,既可以出现在赋值号左边也可以出现在右边
	// 不可以被寻址的为右值，只能出现在赋值号右边，比如函数返回值、字面值、常量值等等

	// f3 实际是左值了，所以可以调用 指针方法，编译期会自动插入 &
	f3 := NewFoo()
	f3.pointerMethod()
}
