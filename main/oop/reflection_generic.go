package oop

import (
	"fmt"
	"reflect"
)

// 空接口、反射和泛型
// 空接口 interface{} 和nil值不是一个概念
func emptyInterface() {
	// 指向基本类型
	var v1 interface{} = 1
	var v2 interface{} = "test"
	var v3 interface{} = true

	// 也可以指向复合类型
	var v4 interface{} = &v2
	var v5 interface{} = []int{1, 2, 3}
	// 将 struct类型传给 interface{}
	var v6 interface{} = struct {
		id   int
		name string
	}{}

	// 空接口最典型的使用场景就是用于声明函数支持任意类型的参数
	// go 的 fmt 就是通过 interface{} 来实现的
	fmt.Println(v1, v2, v3, v4, v5, v6)
}

// 反射
// reflect 包提供的两个最常用、最重要的类型就是 reflect.Type 和 reflect.Value。
// 前者用于表示变量的类型:reflect.TypeOf
// 后者用于存储任何类型的值"reflect.ValueOf。
func reflectDemo() {
	animal := SmallDog{"smallholding"}
	dogType := reflect.TypeOf(animal)
	fmt.Println("type is", dogType)

	// 如果你想要获取 dog 值的结构体信息，并且动态调用其成员方法，使用反射的话需要先获取对应的 reflect.Value 类型值：
	dogValue := reflect.ValueOf(&animal).Elem()
	fmt.Println("dog value", dogValue)
	// 如果类中不包含指针方法的话，也可以返回dog值对应的reflect.Value 类型
	dogValue2 := reflect.ValueOf(animal)
	fmt.Println("value of :", dogValue2)

	str := "123"
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.ValueOf(str))
	fmt.Println("reflect.TypeOf(str).Kind() == reflect.Int:", reflect.TypeOf(str).Kind() == reflect.Int)

	strValue := reflect.ValueOf(&str).Elem()
	fmt.Println("reflect.ValueOf(&str).Elem():", strValue)
	fmt.Println("reflect.ValueOf(str):", strValue)

}

// getFiledAndMethod: 获取实例的所有属性和成员方法并调用
func getFiledAndMethod() {
	animal := SmallDog{"smallholding"}
	dogValue := reflect.ValueOf(&animal).Elem()

	// 获取所有属性
	for i := 0; i < dogValue.NumField(); i++ {
		// 获取属性名
		fmt.Println("name:", dogValue.Type().Field(i).Name)
		// 获取属性类型
		fmt.Println("type:", dogValue.Type().Field(i).Type)
		// 获取属性值
		fmt.Println("value:", dogValue.Field(i))
	}
	fmt.Println("--------")

	// 获取所有方法
	for i := 0; i < dogValue.NumMethod(); i++ {
		t := dogValue.Type()
		fmt.Println("name:", t.Method(i).Name)
		fmt.Println("type:", t.Method(i).Type)
		fmt.Println("exec result", dogValue.Method(i).Call([]reflect.Value{}))
	}
}

// 基于空接口和反射实现泛型

type Container struct {
	s reflect.Value
}

func NewContainer(t reflect.Type, size int) *Container {
	if size <= 0 {
		size = 64
	}

	// 基于切片类型实现这个容器，通过反射初始化这个底层切片
	return &Container{s: reflect.MakeSlice(reflect.SliceOf(t), 0, size)}
}

func (c *Container) Put(val interface{}) error {
	// 通过反射对实际传递进来的元素类型进行运行时检查，
	// 如果与容器初始化时设置的元素类型不同，则返回错误信息
	// c.s.Type() 对应的是切片类型，c.s.Type().Elem() 应的才是切片元素类型
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf("put error:cannot put a %T into a slice of %s", val, c.s.Type().Elem())
	}

	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}

// Get 从容器中读取元素，将返回结果赋值给 val，同样通过空接口指定元素类型
func (c *Container) Get(val interface{}) error {
	// 还是通过反射对元素类型进行检查，如果不通过则返回错误信息
	// Kind 与 Type 相比范围更大，表示类别，如指针，而 Type 则对应具体类型，如 *int
	// 由于 val 是指针类型，所以需要通过 reflect.ValueOf(val).Elem() 获取指针指向的类型
	if reflect.ValueOf(val).Kind() != reflect.Ptr ||
		reflect.ValueOf(val).Elem().Type() != c.s.Type().Elem() {
		return fmt.Errorf("get error: needs *%s but got %T", c.s.Type().Elem(), val)
	}
	// 将容器第一个索引位置值赋值给 val 指针
	reflect.ValueOf(val).Elem().Set(c.s.Index(0))
	c.s = c.s.Slice(1, c.s.Len())
	return nil
}

func TestPut() {
	nums := []int{1, 2, 3, 4, 5}
	c := NewContainer(reflect.TypeOf(nums[0]), 16)

	for _, n := range nums {
		err := c.Put(n)
		if err != nil {
			panic(err)
		}
	}

	// 从容器读取元素，将返回结果初始化为 0
	num := 0
	if err := c.Get(&num); err != nil {
		panic(err)
	}
	fmt.Printf("%v (%T) \n", num, num)

	if err := c.Put("S"); err != nil {
		panic(err)
	}

}
