package function

import (
	"fmt"
	"testing"
)

func TestNewAndMake(t *testing.T) {
	newAndMake()
}

func TestGoCalc(t *testing.T) {
	goCalc()
}

func TestVariableParams(t *testing.T) {
	variableParams()
}

func TestCustomPrintf(t *testing.T) {
	customPrintf(1, "1", [1]int{1}, true)
}

func TestCustomAdd(t *testing.T) {
	a, b := -1, 2
	result, err := customAdd(&a, &b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("add(%d,%d) = %d\n", a, b, result)
}
