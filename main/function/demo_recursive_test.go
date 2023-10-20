package function

import (
	"fmt"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	n := 5
	num := fib(n)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n, num)

	goCatchFibExecTime()
}

func TestFib3(t *testing.T) {
	n := 5
	start := time.Now()
	num := fib3(n)
	end := time.Since(start)
	fmt.Printf("---fib(%d) = %d 执行耗时: %v ---\n", n, num, end)
}
