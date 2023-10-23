package oop

import "fmt"

// GO 没有直接提供继承相关的语法实现，但是可以通过组合的方式间接实现
// 组合就是将一个类型嵌入到到另外一个类型，从而构建新的类型

type Animal struct {
	name string
}

func (a Animal) call() string {
	return "有 animal 在叫"
}

func (a Animal) favouriteFood() string {
	return "favouriteFood"
}

func (a Animal) getName() string {
	return a.name
}

// Dog dog中嵌入了Animal类型，这样一来，可以在Dog实例上访问所有Animal类型的包含的属性和方法
type Dog struct {
	Animal
}

func compositeExtend() {
	animal := Animal{name: "田园dog"}
	dog := Dog{animal}
	// dog 可以访问animal类型包含的属性和方法，相当于变向实现了继承
	fmt.Println(dog.getName())
	fmt.Println(dog.call())
	fmt.Println(dog.favouriteFood())
}

// 多态
// 我们还可以通过在子类中定义同名方法来覆盖父类方法的实现，在面向对象编程中这一术语叫做方法重写
// 比如在上述 Dog 类型中，我们可以重写 Call 方法和 FavoriteFood 方法的实现如下
func (d Dog) favouriteFood() string {
	return "骨头"
}

func (d Dog) call() string {
	return "wang wang wang!"
}

func multi() {
	animal := Animal{name: "田园dog"}
	dog := Dog{animal}

	fmt.Println(dog.Animal.call())
	fmt.Println(dog.call())
	fmt.Println(dog.Animal.favouriteFood())
	fmt.Println(dog.favouriteFood())
}

// 多继承同名方法冲突

type Pet struct {
	name string
}

func (p Pet) getName() string {
	return p.name
}

type Dog2 struct {
	Animal
	Pet
}

func main2() {
	animal := Animal{name: "123"}
	pet := Pet{name: "pets"}
	dog2 := Dog2{animal, pet}
	// 编译会报错，提示 Ambiguous reference 'name'
	//fmt.Println(dog2.name)
	// 需要显式指定
	fmt.Println(dog2.Pet.getName())
	fmt.Println(dog2.Animal.getName())
}

type DogWithPointer struct {
	*Animal
	Pet
}

func main3() {
	animal := Animal{name: "animal"}
	pet := Pet{name: "pets"}

	dog := DogWithPointer{
		Animal: &animal,
		Pet:    pet,
	}

	fmt.Println(dog.Animal.getName())
	fmt.Println(dog.Animal.call())
	fmt.Println(dog.call())
	fmt.Println(dog.Animal.favouriteFood())
	fmt.Println(dog.favouriteFood())
}
