package main

import "fmt"

// go 的数组变量一旦长度声明后就是固定了，无法动态添加元素，如果要添加元素，就只能新建一个更大长度的数组，然后copy过去
// go 的数组类型是值传递，作为参数传递到func的时候，传递是数组的值的拷贝，当函数内对数组进行修改的时候，对应的改动不会影响原来的数组

// 数组初始化
func testArrayInit() {
	// 长度为8的数组，每个元素为1个字节
	// 如果没有填满，则空位会通过对应的元素类型零值填充
	var a [8]byte
	// 二维数组
	var b [3][3]int
	// 三维数组
	var c [3][3][3]float64
	// 声明时候初始化
	var d = [3]int{1, 2, 3}
	// 通过new初始化
	var e = new([3]string)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// := 进行一次性声明和初始化，所有数组元素通过 {} 包裹，然后通过逗号分隔多个元素
	// var array =  [capacity]data_type{element_values}
	f := [5]int{1, 2, 3, 4, 5}
	fmt.Println(f)

	// 省略数组长度的声明, go 会在编译的时候自动计算array长度
	g := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(g))

	// 指定下标填充, 1下标填充9，3下标填充8 --> [0 9 0 8 0]
	h := [5]int{1: 9, 3: 8}
	fmt.Println(h)
}

// 数组的访问设置和遍历
func arrayRange() {
	arr := [5]int{1, 2, 3, 4, 5}
	// 数组的长度通过 len 方法获取
	a1, a2 := arr[0], arr[len(arr)-1]
	fmt.Println(a1, a2)

	arr[0] = 100
	fmt.Println(arr[0])

	for i := 0; i < len(arr); i++ {
		fmt.Println("Element", i, "of arr is", arr[i])
	}

	// range 方式,range返回了两个值，第一个是数组下标index，第二个是对应数组元素的值
	for i, v := range arr {
		fmt.Println("Element", i, "of arr is", v)
	}

	// 只想获取元素，不要索引
	for _, v := range arr {
		fmt.Println("value", v)
	}

	// 只想要索引，不要值,实际是 for i, _ := range arr,  ,_  可以省略
	for i := range arr {
		fmt.Println("Element", i)
	}
}

// 多维数组
// 打印一个99乘法表
func multiArr() {
	var multi [9][9]string
	// i 是横轴
	for i := 0; i < 9; i++ {
		// j 是纵轴
		for j := 0; j < 9; j++ {
			v1 := j + 1
			v2 := i + 1
			// 横轴大于纵轴的时候，不打印，只要对角线内的即可
			if v2 > v1 {
				continue
			}
			// Sprintf : 根据格式说明符进行格式化并返回结果字符串
			multi[j][i] = fmt.Sprintf("%dx%d=%d", v2, v1, v1*v2)
		}
	}

	for _, v1 := range multi {
		for _, v2 := range v1 {
			// 位宽是8，左对齐
			fmt.Printf("%-8s", v2)
		}
		fmt.Println()
	}

}
