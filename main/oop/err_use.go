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

/*
customError2
小写防止在package 外面直接声明，只能通过函数  NewCustomError2声明
*/
type customError2 struct {
	code int
	msg  string
}

// 实现了Error()方法， 就当作 CustomError2 实现了error接口
func (e customError2) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", e.code, e.msg)
}

func NewCustomError2(code int, msg string) error {
	return customError2{
		code: code,
		msg:  msg,
	}
}

func GetCode(err error) int {
	if e, ok := err.(customError2); ok {
		return e.code
	}
	return -1
}

func GetMsg(err error) string {
	var e customError2
	if errors.As(err, &e) {
		return e.msg
	}
	return ""
}

// 如果error变大写，也就是 在package 外可以直接声明，会存在问题

type MyError struct {
	code int
	msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", e.code, e.msg)
}

func HelloWorld() error {
	var err *MyError
	return err
}

func fly() {
	err := HelloWorld()

	fmt.Printf("%T,%v\n", err, err)
	fmt.Println(err == nil)
	// tm 竟然是false 原因是 error 是个接口，动态类型和动态值都是nil 才是nil，现在动态类型已经是 *MyError了,只是没有值而已
	var err2 *MyError
	fmt.Println(err2 == nil)
}

/*
上面 err == nil 是false的原因，和go 的interface设计有关
go的interface设计成了两部分：type 和value
其中：value 由一个任一的具体值标识，称为是 interface的dynamic value
而type 则对应value类，dynamic type
对于 var a int = 3 来说， 把 a 复制给interface 的时候，interface 是使用（int，3） 进行存储的
当判断interface 的值为nil时候，则必须其内部value和type均未设置的情况才是nil

上面的例子，var err *MyErr 的时候，interface 存的是 (*MyErr,nil)，这个时候 就不是nil了

参考：https://golang.design/go-questions/interface/dynamic-typing/
*/
