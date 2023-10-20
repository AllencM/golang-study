package function

import "fmt"

// 匿名函数与闭包
// 所谓闭包指的是引用了自由变量（未绑定到特定对象的变量，通常在函数外定义）的函数，
// 被引用的自由变量将和这个函数一同存在，即使已经离开了创造它的上下文环境也不会被释放（比如传递到其他函数或对象中）。
// 简单来说，「闭」的意思是「封闭外部状态」，即使外部状态已经失效，闭包内部依然保留了一份从外部引用的变量。
// 显然，闭包只能通过匿名函数实现，我们可以把闭包看作是有状态的匿名函数，反过来，如果匿名函数引用了外部变量，就形成了一个闭包(Closure)
func anonymous() {
	j := 1
	f := func() {
		i := 1
		fmt.Printf("i, j: %d, %d\n", i, j)
	}
	f()
	j += 2
	f()

	// 上面的示例中，匿名函数引用了外部变量，所以同时也是个闭包，
	// 变量 f 指向的闭包引用了局部变量 i 和 j，i 在闭包内部定义，其值被隔离，不能从外部修改，而变量 j 在闭包外部定义，所以可以从外部修改，闭包持有的只是其引用
}

// 匿名函数作为函数参数
func anonymousParam() {
	add := func(a, b int) int {
		return a + b
	}

	func(call func(int, int) int) {
		fmt.Println(call(1, 2))
	}(add)
}

func anonymousParam2() {
	add := func(a, b int) int {
		return a + b
	}

	base := 10
	add2 := func(a, b int) int {
		return a*base + b
	}
	handleAdd(1, 2, add)
	handleAdd(1, 2, add2)

	// 第二个匿名函数 add2 引用了外部变量base，形成了一个闭包，在调用handleAdd 外部函数时传入了闭包add2作为参数，
	// add2 闭包在外部函数中执行时，虽然作用域离开了 anonymousParam2 函数，但是还是可以访问到变量base
	// 这样一来，就可以通过一个函数值执行多种不同算法，提升了代码的复用性

}

func handleAdd(a, b int, call func(int, int) int) {
	fmt.Println(call(a, b))
}

// 将匿名函数作为函数返回值
func deferAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}

func goDeferAdd() {
	addFunc := deferAdd(1, 2)
	// 这里才会真正执行对应操作
	fmt.Println(addFunc())
}
