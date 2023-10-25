package main

import (
	"fmt"
	"testing"
)

// 单元测试的方法一定要 Test开头
func TestSliceInit(t *testing.T) {
	initSlice()
}

func TestSliceMake(t *testing.T) {
	sliceMake()
}

func TestSliceIterate(t *testing.T) {
	iterateSlice()
}

func TestModifySlice(t *testing.T) {
	modifySlice()
}

func TestSliceArrayShare(t *testing.T) {
	sliceArrayShare()
}

// slice 特别篇

func TestSliceParam(t *testing.T) {
	slice := make([]int, 0, 4)
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3)
	sliceMod(slice)
	fmt.Println(slice)
}

func sliceMod(slice []int) {
	slice = append(slice, 4)
}

func TestSlice2(t *testing.T) {
	slice := make([]int, 0, 4)
	slice = append(slice, 1, 2, 3)
	sliceMod2(slice)
	fmt.Println(slice)
}

func sliceMod2(slice []int) {
	slice = append(slice, 4)
	slice[0] = 10
}
