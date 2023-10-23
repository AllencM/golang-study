package function

import (
	"log"
	"testing"
)

func TestSumAge(t *testing.T) {
	var users = []user{
		{
			name: "张三",
			age:  18,
		},
		{
			name: "李四",
			age:  22,
		},
		{
			name: "王五",
			age:  20,
		},
		{
			name: "赵六",
			age:  10,
		},
		{
			name: "孙七",
			age:  60,
		},
		{
			name: "周八",
			age:  10,
		},
	}

	sum := sumAge(users, filterAge, mapAgeToSlice)
	log.Printf("user age sum result is %d", sum)
}
