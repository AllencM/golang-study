package function

import "log"

type user struct {
	name string
	age  int
}

// filterAge 返回类型声明成 interface类型，为了便于统一管道声明filter 和 map 函数
// interface 类型可以用于表示返回任何类型
func filterAge(users []user) interface{} {
	var slice []user
	for _, u := range users {
		if u.age >= 18 && u.age <= 35 {
			slice = append(slice, u)
		}
	}
	return slice
}

func mapAgeToSlice(users []user) interface{} {
	var slice []int
	for _, u := range users {
		slice = append(slice, u.age)
	}
	return slice
}

func sumAge(users []user, pipes ...func([]user) interface{}) int {
	var ages []int
	var sum int
	for _, f := range pipes {
		result := f(users)
		switch result.(type) {
		case []user:
			users = result.([]user)
		case []int:
			ages = result.([]int)
		}
	}

	if len(ages) == 0 {
		log.Fatalln("没有在管道中加入 mapAgeToSlice 方法")
	}

	for _, age := range ages {
		sum += age
	}
	return sum

}
