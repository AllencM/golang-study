package condition

import "fmt"

// Go 语言的流程控制和其他编程语言类似，支持如下几种流程控制语句：
//
//条件语句：用于条件判断，对应的关键字有 if、else 和 else if；
//分支语句：用于分支选择，对应的关键字有 switch、case 和 select（用于通道，后面介绍协程时会提到）；
//循环语句：用于循环迭代，对应的关键字有 for 和 range；
//跳转语句：用于代码跳转，对应的关键字有 goto。
func ifCondition(score float64) {
	// 条件语句不需要使用()
	// 无论语句体内有几条语句，花括号 {} 都是必须存在的；
	// 左花括号 { 必须与 if 或者 else 处于同一行；
	if score > 90 {
		fmt.Println("A")
	} else if score > 80 {
		fmt.Println("B")
	} else if score > 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}

// switch...case...  和条件语句if一样，左花括号 { 必须与 switch 处于同一行；
// 单个 case 中，可以出现多个结果选项（通过逗号分隔）；
// 与其它语言不同，Go 语言不需要用 break 来明确退出一个 case；
// 只有在 case 中明确添加 fallthrough 关键字，才会继续执行紧跟的下一个 case；
// 可以不设定 switch 之后的条件表达式，在这种情况下，整个 switch 结构与多个 if...else... 的逻辑作用等同。
func switchCondition(score float64) {
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80 && score < 90:
		fmt.Println("B")
	case 70 <= score && score < 80:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// 如果 score 放到 switch 后面 这时候只会判等， 90,100 即 等于90 或等于100
	switch score {
	case 90, 100:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// 合并分支， go 的case 会自动退出，所以想要继续执行后续的分支，则需要通过  fallthrough 来处理
	// 通过 fallthrough 合并 case 70 <= score && score < 80 和  case score < 70
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80 && score < 90:
		fmt.Println("B")
	case 70 <= score && score < 80:
		fallthrough
	case score < 70:
		fmt.Println("score is <= 70")
	}

}

// 和条件语句、分支语句一样，左花括号 { 必须与 for 处于同一行；
// 不支持 while 和 do-while 结构的循环语句；
// 可以通过 for-range 结构对可迭代集合进行遍历；
// 支持基于条件判断进行循环迭代；
// 允许在循环条件中定义和初始化变量，且支持多重赋值
func testFor() {
	// 无限循环
	sum := 0
	i := 0
	for {
		i++
		if i > 100 {
			break
		}
		sum += i
	}
	fmt.Println(sum)

	// 基于条件判断循环
	sum = 0
	i = 0
	for i < 100 {
		i++
		sum += i
	}
	fmt.Println(sum)
}

// go 同样支持跳出外层循环
func testBreak() {
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > 5 {
				break outer
			}
			fmt.Println(i)
		}
	}
}

// go 支持 goto 跳转指定位置，跳转到本函数内的某个标签
// goto 容易造成代码逻辑混乱，导致新的bug
func testGoto() {
	arr := []string{"1", "2", "3", "4"}
	for i := 0; i < len(arr); i++ {
		// arr[i] 的值是3的时候，跳转到 exit 标签
		if arr[i] == "3" {
			goto exit
		}
		fmt.Println(arr[i])
	}

exit:
	fmt.Println("跳转到了exit")
}
