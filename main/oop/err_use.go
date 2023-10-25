package oop

import (
	"errors"
	"fmt"
	"strconv"
)

/*
go 为错误处理定义了一个标准的error接口
type error interface {
	Error() string
}
*/

func divide(a, b int) (n int, err error) {
	if b == 0 {
		err = errors.New("divisor is zero")
		return
	}

	return a / b, nil
}

func callDivide() {
	a, b := 100, 0
	result, err := divide(a, b)
	if err != nil {
		// 错误处理
		panic(err)
		return
	}

	fmt.Println("result is:", result)
}

func HandleError(args []string) {
	if len(args) != 3 {
		fmt.Println("args len isn't 3!,input args len is", len(args))
		return
	}

	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	z, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("there is no error occurred,result is:", z)
	}
}

// CustomError 自定义异常
type CustomError struct {
	Message string
	Err     error
}

func (c *CustomError) Error() string {
	c.Message = "出现了 custom error"
	return c.Message
}

func NewCustomError(text string) *CustomError {
	return &CustomError{Message: text, Err: errors.New(text)}
}

func CustomErrors(a, b int) (c int, err CustomError) {
	if b == 0 {
		err = *NewCustomError("divisor is zero")
	}
	return a / b, err
}
