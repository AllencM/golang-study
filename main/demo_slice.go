package main

import "fmt"

// 切片是一个新的数据类型，与数组最大的不同是 切片的长度可以变化的。
// 切片是一个可变长度，同一类型的元素集合
func initSlice() {
	var slice = []string{"a", "b", "c"}
	fmt.Println(slice)

	// 基于数组创建slice
	months := [...]string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	// 第二季度 Go 语言支持通过 array[start:end] 这样的方式基于数组生成一个切片，
	// start 表示切片在数组中的下标起点，end 表示切片在数组中的下标终点，两者之间的元素就是切片初始化后的元素集合
	// 通常等于：[start:end), 例外情况： month[6:] 这个时候表示下标6 到最后，前后都包含
	q2 := months[3:6]
	q3 := months[6:9]

	fmt.Println(q2)
	fmt.Println(q3)

	all := months[:]
	fmt.Println(all)

	// 上半年
	firstHalf := months[:6]
	// 下班年
	secondHalf := months[6:]
	fmt.Println(firstHalf, secondHalf)

	// slice 底层引用了一个数组，三部分构成：指针、长度、容量
	// 指针指向数组的起始下标，长度对应数组中的元素，容量则是切片起始位置到底层数组结尾的位置
	// 切片长度不能超过容量，比如说上面的 q2 ，指向的是 month 下标为3 的位置，长度是3，容量是9 （下标3 到下标11结束，可容纳9个元素）

	// 切片长度用 len
	fmt.Println(len(q2))
	// 切片容量
	fmt.Println("cap(q2) is", cap(q2))
	fmt.Println("cap(q3) is", cap(q3))
	fmt.Println("cap(firstHalf) is", cap(firstHalf))
	fmt.Println("cap(secondHalf) is", cap(secondHalf))

	// 切片修改，会影响原先数组的元素
	//firstHalf[0] = "test"
	//fmt.Println(months)
	//fmt.Println(firstHalf)

	// 基于切片创建切片
	q1 := firstHalf[:3]
	fmt.Println(q1)
	fmt.Println("cap(q1) is", cap(q1))
}

// 并非一定要准备一个array 才能创建切片，可以直接通过make创建
func sliceMake() {
	// 初始长度是2 的切片，未指定容量，这个时候容量是2
	slice1 := make([]int, 2)
	fmt.Println("slice1 len is", len(slice1), "cap(slice1) is", cap(slice1))
	// 初始长度是5，容量是10
	slice2 := make([]int, 5, 10)
	fmt.Println("slice2 len is", len(slice2), "cap(slice2) is", cap(slice2))
	// 直接创建并初始化包含5个元素的数组切片（len cap 均是5）
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice3 len is", len(slice3), "cap(slice3) is", cap(slice3))

	// 和array 一样，未初始化的slice，元素均为0, slice1: [0 0]
	fmt.Println(slice1)
}

// 切片的遍历 和数组一致，并无区别
func iterateSlice() {
	slice1 := []string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	for i := 0; i < len(slice1); i++ {
		fmt.Println("slice[", i, "] is ", slice1[i])
	}

	for i, v := range slice1 {
		fmt.Println("slice[", i, "] = ", v)
	}

}

// 切片比数组更强大的地方在于动态新增元素，甚至可以自动扩容
func modifySlice() {
	// 一个切片的容量初始值根据创建方式的不同而不同：
	//
	//对于基于数组和切片创建的切片而言，默认容量是从切片起始索引到对应底层数组的结尾索引；
	//对于通过内置 make 函数创建的切片而言，在没有指定容量参数的情况下，默认容量和切片长度一致。
	//所以，通常一个切片的长度值小于等于其容量值，我们可以通过 Go 语言内置的 cap() 函数和 len() 函数来获取某个切片的容量和实际长度：

	var oldSlice = make([]int, 5, 10)
	fmt.Println("len(oldSlice)", len(oldSlice))
	fmt.Println("cap(oldSlice)", cap(oldSlice))

	// 通过append 函数向切片心新增元素
	newSlice := append(oldSlice, 1, 2, 3)
	fmt.Println(newSlice)
	fmt.Println("len(newSlice)", len(newSlice))
	fmt.Println("cap(newSlice)", cap(newSlice))

	// append 整个slice,append 函数并不会修改原先的切片，而是会生成一个容量更大的切片，把原先的元素和新元素 一并copy进去
	appendSlice := []int{1, 2, 3}

	appendSlice2 := append(oldSlice, appendSlice...) // 末尾的 ...不能少
	fmt.Println(appendSlice2)

	// slice 会自动扩容，默认是原来的2倍，如果还不足以容纳新元素，则按照同样的操作继续扩容，直到新容量不小于原长度与要追加的元素数量之和。
	// 但是，当原切片的长度大于或等于 1024 时，Go 语言将会以原容量的 1.25 倍作为新容量的基准。
	// 因此，如果事先能预估切片的容量并在初始化时合理地设置容量值，可以大幅降低切片内部重新分配内存和搬送内存块的操作次数，从而提高程序性能

	// copy 函数,返回拷贝成功的元素的个数
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7, 8}
	slice3 := []int{4, 5, 6, 7, 8}
	// slice1 拷贝到 slice2 中，只会复制slice1 的前3个元素
	succNum := copy(slice2, slice1)
	fmt.Println("copy(slice2, slice1) is", slice2, "succNum:", succNum)

	// slice 3 到 slice1 ,只会复制slice3的前3个到slice1
	copy(slice1, slice3)
	fmt.Println("copy(slice1, slice3)", slice1)

	// 不同类型切片无法 copy
	//slice4 := []string{"123", "456"}
	//copy(slice1,slice4)

	// 动态删除元素(其实就是通过切片实现的伪删除)
	slice4 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 删除slice4 最后5个元素，删除不会改变容量，会改变len
	slice4 = slice4[:len(slice4)-5]
	fmt.Println(slice4)
	fmt.Println(len(slice4), cap(slice4))

	// via append or copy
	// 删除开头三个
	slice4 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice4[:0])
	slice5 := append(slice4[:0], slice4[3:]...)
	fmt.Println(slice5)

	slice6 := append(slice4[:1], slice4[4:]...)
	fmt.Println(slice6)

	// copy 之所以可以用于删除元素，是因为其返回值是拷贝成功的元素个数，我们可以根据这个值完成新切片的设置从而达到「删除」元素的效果。
	slice4 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	// [4 5 6 7 8]
	fmt.Println(slice4[3:])
	// 此时 slice4 : [4 5 6 7 8 6 7 8],copyElements = 5
	copyElements := copy(slice4, slice4[3:])
	fmt.Println("copy(slice4, slice4[3:])", copyElements, "slice4 is", slice4)
	// 通过copy删除开头3个元素
	slice7 := slice4[:copyElements]
	fmt.Println(slice7)
}

// slice 底层是通过 array 实现的，其array 是共享的
func sliceArrayShare() {
	slice := []int{1, 2, 3, 4, 5}
	// slice2 是基于slice 创建的，他们的数组指针指向了同一个数组，因此，修改修改了slice2，会同步slice，因为他们修改的是同一份数据
	slice2 := slice[1:3]
	slice2[0] = 99
	fmt.Println(slice, "\n", slice2)

	// 如果想不互相影响
	slice3 := make([]int, 4)
	slice4 := slice3[1:3]
	// 通过append 申请了新的内存，slice4 和新的 slice3 并不共享底层的array数组
	// slice3 容量是4，append后，会扩容，等于新申请了一个内存，再指向slice3
	// 如果 slice3 的定义修改为：  slice3 := make([]int, 4, 5) 这个时候通过append 并不会扩容，原因是原先slice3容量够，不会申请新的内存
	// 所以解决array共享的问题关键是扩容，重新分配新的内存空见，否则不行
	slice3 = append(slice3, 0)
	slice3[1] = 99
	slice4[1] = 88
	fmt.Println(slice3)
	fmt.Println(slice4)

	oldSlice := make([]int, 4, 5)
	newSlice := oldSlice[1:4]
	// append 没有触发扩容，还是同一个内存
	oldSlice = append(oldSlice, 0)
	oldSlice[1] = 999
	newSlice[1] = 888
	fmt.Println(oldSlice)
	fmt.Println(newSlice)

}
