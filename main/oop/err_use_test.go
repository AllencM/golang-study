package oop

import "testing"

func TestCallDivide(t *testing.T) {
	callDivide()
}

func TestHandleError(t *testing.T) {
	args := []string{"100", "2", "3"}
	HandleError(args)

	args = args[1:]
	HandleError(args)

}
