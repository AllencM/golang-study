package function

import (
	"fmt"
	"strconv"
)

func ageSum(users []map[string]string) int {
	var sum = 0
	for _, user := range users {
		num, _ := strconv.Atoi(user["age"])
		sum += num
	}
	return sum
}

func addAgeSum() {
	users := []map[string]string{
		{"name": "张1", "age": "18"},
		{"name": "张2", "age": "19"},
		{"name": "张3", "age": "20"},
	}

	sum := ageSum(users)
	fmt.Println("age is", sum)
}

func mapToString(items []map[string]string, getUserAge func(map[string]string) string) []string {
	newSlice := make([]string, len(items))
	for _, item := range items {
		newSlice = append(newSlice, getUserAge(item))
	}
	return newSlice
}

func fieldSum(items []string, f func(string) int) int {
	sum := 0
	for _, item := range items {
		sum += f(item)
	}
	return sum
}

func addAgeSum2() {
	users := []map[string]string{
		{"name": "张1", "age": "18"},
		{"name": "张2", "age": "19"},
		{"name": "张3", "age": "20"},
	}

	ageSlice := mapToString(users, func(user map[string]string) string {
		return user["age"]
	})

	sum := fieldSum(ageSlice, func(age string) int {
		intAge, _ := strconv.Atoi(age)
		return intAge
	})

	fmt.Println(sum)
}

// using filter
func itemsFilter(items []map[string]string, f func(map[string]string) bool) []map[string]string {
	newSlice := make([]map[string]string, len(items))
	for _, item := range items {
		if f(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

func callFilter() {
	var users = []map[string]string{
		{
			"name": "张三",
			"age":  "18",
		},
		{
			"name": "李四",
			"age":  "22",
		},
		{
			"name": "王五",
			"age":  "20",
		},
		{
			"name": "赵六",
			"age":  "-10",
		},
		{
			"name": "孙七",
			"age":  "60",
		},
		{
			"name": "周八",
			"age":  "10",
		},
	}

	validUsers := itemsFilter(users, func(user map[string]string) bool {
		age, ok := user["age"]
		if !ok {
			return false
		}

		intAge, err := strconv.Atoi(age)
		if err != nil {
			return false
		}

		return intAge >= 18 && intAge <= 35
	})

	ageSlice := mapToString(validUsers, func(user map[string]string) string {
		return user["age"]
	})

	sum := fieldSum(ageSlice, func(age string) int {
		i, _ := strconv.Atoi(age)
		return i
	})

	fmt.Println(sum)

}

//func filter(users []map[string]string) []map[string]string {
//}
