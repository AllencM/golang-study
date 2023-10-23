package oop

import (
	"fmt"
	"testing"
)

func TestInteger_Equal(t *testing.T) {
	integerEqual()
}

func TestInteger_Oop(t *testing.T) {
	integerOop()
}

func TestNewStudent(t *testing.T) {
	student := NewStudent(1, "axi", 1, 88.8)
	fmt.Println(student)

	student2 := NewStudent2(2, "123z", 1)
	fmt.Println(student2)
}

func TestFoo(t *testing.T) {
	GoFoo()
}
