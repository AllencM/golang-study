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
