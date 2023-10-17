package main

import "testing"

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
