package function

import (
	"errors"
	"fmt"
	"reflect"
)

// go 中 没有public、protected 和 private  关键字
// 通过大小写字母来区分可见性
// 首字母小写的只能在同一个包中访问
// 首字母大写的可以在其他包中访问，全局变量也是
// go 内置函数参考：https://pkg.go.dev/builtin

//  new 函数作用于值类型，仅分配内存空间，返回的是指针
//  make 函数作用于引用类型，除了分配内存空间，还会对对应类型进行初始化，返回的是初始值
func newAndMake() {
	// int 指针，等价于 var p1 *int
	p1 := new(int)
	// string 指针
	p2 := new(string)
	// 数组类型指针
	p3 := new([3]int)

	type Student struct {
		id    int
		name  string
		grade string
	}

	// 对象类型的指针
	p4 := new(Student)

	println("p1:", p1)
	println("*p1:", *p1)
	println("p2:", p2)
	println("p3:", p3)
	println("p4:", p4)

	// 返回初始化后的切片类型值，即 []int{0, 0, 0}
	s1 := make([]int, 3)
	// 返回初始化的字典类型值，即散列化的 map 结构
	m1 := make(map[string]int, 2)

	println(s1)

	for i, v := range s1 {
		println(i, v)
	}

	println("len(m1)", len(m1))
	m1["test"] = 100

	for k, v := range m1 {
		println(k, v)
	}
}

// 参数传递
// go 使用的是按值来传递参数，也就是传递了一个副本
// 函数接收到传递进来的参数后，会将参数值拷贝给声明该参数的变量（也叫形式参数，简称形参），
// 如果在函数体中有对参数值做修改，实际上修改的是形参值，这不会影响到实际传递进来的参数值（也叫实际参数，简称实参）。
func calc(a, b int) int {
	a *= 2
	b *= 3
	return a + b
}

func goCalc() {
	x, y := 1, 2
	z := calc(x, y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)

	// &x &y 获得是x y 的变量地址，不是变量值
	x, y = 1, 2
	zz := calcPointer(&x, &y)
	fmt.Println("zz", zz)
}

// 引用传参
// 如果想在函数中修改的参数可以同步修改实际参数，需要通过引用传参来完成
func calcPointer(a, b *int) int {
	// *a *b 指向的是实际值
	*a *= 2
	*b *= 3
	return *a + *b
}

// 变长参数，只要在参数类型前加上 ...
func variableParams() {
	myFunc(1, 2, 3, 4, 5)
	fmt.Println("------")
	// 边长参数还支持通过 []int 类型的切片
	slice := []int{6, 7, 8, 9, 10}
	myFunc(slice...)
	fmt.Println("------")
	// 注：形如 ...type 格式的类型只能作为函数的参数类型存在，并且必须是函数的最后一个参数。
	myFunc(slice[1:3]...)
}

func myFunc(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

// 如果没有 ...type ，只能这么定义
func myFunc2(a []int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

// 任意类型的变长参数（泛型）
// go 并没有在语法层面提供对泛型的支持
// interface{} 是一个空的接口，可以用于标识任意类型，为了保证代码类型安全，需要在运行时通过反射对数据类型进行检查
func customPrintf(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type")
		default:
			fmt.Println(arg, "is an unknown type")
		}
	}
}

// 多返回值
func customAdd(a, b *int) (int, error) {
	// *a *b  指向的是实际值  &a &b 指向的是内存地址
	if *a < 0 || *b < 0 {
		err := errors.New("只支持非负整数相加")
		return 0, err
	}

	*a *= 2
	*b *= 3
	return *a + *b, nil
}

// 命名返回值
// 在设置多返回值时，还可以对返回值进行变量命名，这样，我们就可以在函数中直接对返回值变量进行赋值，而不必每次都按照指定的返回值格式返回多个变量
func customAdd2(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}
	*a *= 2
	*b *= 3
	c = *a + *b
	return
}
