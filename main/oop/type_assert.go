package oop

import (
	"fmt"
	"reflect"
)

// 接口类型断言
func assert() {
	var num1 Number = 1
	var num2 Number2 = &num1
	// num2.(Number1) 这个表达式断言 num2 是否是 Number1 类型的实例，如果是，ok 值为 true，然后执行 if 语句块中的代码；否则 ok 值为 false，不执行 if 语句块中的代码。
	if num3, ok := num2.(Number1); ok {
		fmt.Println(num3.Equal(1))
	}
	// 类型断言是否成功要在运行期才能够确定，它不像接口赋值，编译器只需要通过静态类型检查即可判断赋值是否可行
}

// struct 类型的断言

type IAnimal interface {
	GetName() string
	Call() string
	FavouriteFood() string
}

type SmallDog struct {
	name string
}

type BigDog struct {
	name string
}

func (s SmallDog) GetName() string {
	fmt.Println("my name is small dog")
	return "small dog" + s.name
}

func (s SmallDog) Call() string {
	return "small dog wang wang wang"
}

func (s SmallDog) FavouriteFood() string {
	return "骨头啊肯定"
}

func (b BigDog) GetName() string {
	return "big dog is here" + b.name
}

func (b BigDog) Call() string {
	return "small dog where are you going?"
}

func (b BigDog) FavouriteFood() string {
	return "my favourite food is small dog!"
}

// 需要注意的是，在 Go 语言结构体类型断言时，子类的实例并不归属于父类，即使子类和父类属性名和成员方法列表完全一致，
// 因为类与类之间的「继承」是通过组合实现的，并不是 Java/PHP 中的那种父子继承关系，这是新手需要注意的地方。
// 同理，父类实现了某个接口，不代表组合类它的子类也实现了这个接口
func assertStruct() {
	// var smallDog = SmallDog{"small dog"}
	// var bigDog = BigDog{"big dog"}

	//if dog,ok := I
	//
}

// 基本数据类型，不需要通过反射，直接用 variable.(type) 即可获取variable变量对应的类型值
func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case bool:
			fmt.Println(arg, "is a bool value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 基于反射在运行时动态进行类型断言，使用 reflect 包提供的 TypeOf 函数
func myPrintf2(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value.")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
