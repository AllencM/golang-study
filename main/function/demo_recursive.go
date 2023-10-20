package function

import (
	"fmt"
	"time"
)

// 递归函数
// 某个满足以下条件就可以通过递归函数来解决：
// 1、一个问题的解可以被拆分成多个子问题的解
// 2、拆分前的原问题与拆分后的子问题除了数据规模不同，求解思路完全一样
// 3、子问题存在递归终止条件
// 所以我们可以归纳出递归函数的编写思路：抽象出递归模型（可以被复用到子问题的模式/公式），同时找到终止条件。

type FibFunc func(int) int

// 实现一个基本的斐波那契数列
func fib(i int) int {
	if i == 1 {
		return 0
	}
	if i == 2 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func catchFibExecTime(f FibFunc) FibFunc {
	return func(i int) int {
		start := time.Now()
		num := f(i)
		end := time.Since(start)
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return num
	}
}

func goCatchFibExecTime() {
	n1 := 5
	f := catchFibExecTime(fib)
	result1 := f(n1)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, result1)

	//n1 = 50
	//result2 := f(n1)
	//fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, result2)

	f2 := catchFibExecTime(fib2)
	result3 := f2(n1)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, result3)

}

// 进一步优化
// 通过一个数组来保存之前的计算结果，在下一次递归的时候直接拿数组里面的值即可
var fibs [100]int

func fib2(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	index := n - 1
	if fibs[index] != 0 {
		return fibs[index]
	}
	num := fib2(n-1) + fib2(n-2)
	fibs[index] = num
	return num
}

// 通过尾递归优化递归函数性能
// 计算机科学里，尾调用是指一个函数的最后一个动作是调用一个函数（只能是一个函数调用，不能有其他操作，比如函数相加、乘以常量等）
//func f(x int) int {
//  ...
//  return g(x);
//}
// 尾调用的一个重要特性是它不是在函数调用栈上添加一个新的堆栈帧 —— 而是更新它，尾递归自然也继承了这一特性，
// 这就使得原来层层递进的调用栈变成了线性结构，因而可以极大优化内存占用，提升程序性能，这就是尾递归优化技术。
// 优化一下上面的fib2 不保存中间状态

// n：斐波那契数列中的当前位置
func fibTail(n, first, second int) int {
	fmt.Println("n", n, "first", first, "second", second)
	if n < 2 {
		return first
	}
	// 当前的 first + second 的和赋值给下次调用的 second函数
	// 当前second值复制给下次调用的 first ，等同于实现了 fib(n) = fib(n-1) + fib(n-2)
	// 循环往复，直到fib(1) = 0
	return fibTail(n-1, second, first+second)
}

func fib3(n int) int {
	return fibTail(n, 0, 1)
}
