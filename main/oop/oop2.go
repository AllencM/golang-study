package oop

import "fmt"

// 类的定义和初始化

// Student 类型 go不支持构造函数，也没有 class、extends、implements等关键字和概念
type Student struct {
	id    uint
	name  string
	male  int
	score float64
}

// NewStudent 定义一个初始化函数,返回一个指向该类的指针
func NewStudent(id uint, name string, male int, score float64) *Student {
	return &Student{id, name, male, score}
}

// NewStudent2 初始化指定字段
// 在 Go 语言中，未进行显式初始化的变量都会被初始化为该类型的零值，
// bool 类型的零值为 false，int 类型的零值为 0，string 类型的零值为空字符串，float 类型的零值为 0.0。
func NewStudent2(id uint, name string, male int) *Student {
	return &Student{
		id:   id,
		name: name,
		male: male,
	}
}

// GetName 定义成员方法
// 为 Go 类定义成员方法需要在 func 和方法名之间声明方法所属的类型（有的地方将其称之为接收者声明）
func (s Student) GetName() string {
	return s.name
}

func (s Student) GetScore() float64 {
	return s.score
}

// 指针方法
// 在类的成员方法中，可以通过声明的类型变量来访问类的属性和其他方法（Go 语言不支持隐藏的 this 指针，所有的东西都是显式声明）。
// GetName 是一个只读方法，如果我们要在外部通过 Student 类暴露的方法设置 name 值，可以这么做：

func (s *Student) SetName(name string) {
	s.name = name
}

// 这是因为 Go 语言面向对象编程不像 PHP、Java 那样支持隐式的 this 指针，所有的东西都是显式声明的，
// 在 GetXXX 方法中，由于不需要对类的成员变量进行修改，所以不需要传入指针，
// 而 SetXXX 方法需要在函数内部修改成员变量的值，并且该修改要作用到该函数作用域以外，所以需要传入指针类型（结构体是值类型，不是引用类型，所以需要显式传入指针）。
//
// 我们可以把接收者类型为指针的成员方法叫做指针方法，把接收者类型为非指针的成员方法叫做值方法，
// 二者的区别在于值方法传入的结构体变量是值类型（类型本身为指针类型除外），因此传入函数内部的是外部传入结构体实例的值拷贝，修改不会作用到外部传入的结构体实例。

// 另外，需要声明的是，在 Go 语言中，当我们将成员方法 SetName 所属的类型声明为指针类型时，
// 严格来说，该方法并不属于 Student 类，而是属于指向 Student 的指针类型，所以，归属于 Student 的成员方法只是 Student 类型下所有可用成员方法的子集，归属于 *Student 的成员方法才是 Student 类完整可用方法的集合。
//
// 在调用值方法和指针方法时，需要记住以下两条准则：
//
// 值方法可以通过指针和值类型实例调用，指针类型实例调用值方法时会自动解引用；
// 指针方法只能通过指针类型实例调用，
// 但有一个例外，如果某个值是可寻址的（或者说左值），那么编译器会在值类型实例调用指针方法时自动插入取地址符，使得在此情形下看起来像指针方法也可以通过值来调用。

func NewStudentV2(id uint, name string, score float64) Student {
	return Student{id: id, name: name, score: score}
}

func tst() {
	s := NewStudent(1, "test", 1, 100)
	s.SetName("test2")
	fmt.Println(s.GetName())

	s2 := NewStudentV2(2, "test1", 100)
	s2.SetName("SETNEame")
	fmt.Println(s2.GetName())
	NewStudent(3, "1111", 1, 80).SetName("asdasd") // ok 正常调用指针方法
	// NewStudentV2(4, "2222", 99).SetName("学院君4号")  // err 值类型调用指针方法
	// 所谓左值就是可以出现在赋值等号左边的值，而右值只能出现在赋值等号右边，比如函数返回值、字面量、常量值等。左值可寻址，右值不可寻址。
}

// 总结下来，就是一个自定义数据类型的方法集合中仅会包含它的所有「值方法」，
// 而该类型对应的指针类型包含的方法集合才囊括了该类型的所有方法，包括所有「值方法」和「指针方法」，指针方法可以修改所属类型的属性值，而值方法则不能。
// 定义值方法还是指针方法，通长情况有以下情况，考虑将类方法定义为指针方法
// 1、数据一致性，方法需要修改传入的类型实例本身
// 2、方法执行效率，如果是值方法，在方法调用时一定会产生值拷贝，而大对象拷贝的代价会很大

// go String, toString（java）
func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}

// #############################
// # 遇事不决请用 pointer method #
// #############################
