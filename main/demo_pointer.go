package main

import (
	"fmt"
	"unsafe"
)

// 指针用来指向这些变量所在的内存地址的值（变量值所在内存地址的值不等于该内存地址存储的变量值）
// 指针变量再传值的时候可以节省空见，是因为指针指向的内存地址大小是固定的，32位机器上4字节，64位机器上8字节，与 指针指向内存地址存储的值类型无关
// 指针在 Go 语言中有两个典型的使用场景： 类型指针、切片
func initPointer() {
	a := 100
	// 声明一个指针，(没有指向任何变量内存时候，值是 nil)
	var ptr *int
	ptr2 := &a
	// 可以通过 new 声明
	ptr3 := new(int)

	// 初始化指针类型的变量为 a， &a 获得a 的变量所在的内存地址
	ptr = &a
	// ptr指向的是内存的地址，每次打印的ptr值可能不一样，因为存储变量a的内存地址在变动，不同的os也会有不同
	fmt.Println(ptr)
	// *ptr指向的是实际的值
	fmt.Println(*ptr)
	// 在格式化输出时，可以通过 %p 来标识指针类型。
	fmt.Printf("%p\n", ptr2)
	fmt.Printf("%d\n", *ptr2)

	*ptr3 = 100
	fmt.Println(ptr3)
}

// 不通过指针传值
func swap(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

// 通过指针传值
func swap2(a, b *int) {
	// 直接交换了内存地址
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func swapUse() {
	a := 1
	b := 2
	swap(a, b)
	a = 3
	b = 4
	swap2(&a, &b)
}

// unsafe.Pointer 是特别定义的一种指针类型，可以包含任意类型的变量地址（类似C中的void类型指针）
// 官方文档描述：
// 1、任何类型的指针都可以转化成unsafe.Pointer
// 2、unsafe.Pointer 可以转成任何类型的指针
// 3、uintptr 可以被转化为 unsafe.Pointer
// 4、unsafe.Pointer 可以被转化为 uintptr
func unsafePointer() {
	i := 10
	var p *int = &i
	// 将 *int 类型指针转化成了 unsafe.Pointer类型，再转成了 *float32
	var fp *float32 = (*float32)(unsafe.Pointer(p))
	*fp = *fp * 10
	fmt.Println(i)

	// 指针运算
	arr := [3]int{1, 2, 3}
	ap := &arr

	// sp := uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])
	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	*sp += 3
	fmt.Println(arr)

}
