package main

import (
	"fmt"
	"sort"
)

// 字典(无序的)
func initMap() {
	// key type is string, value type is int, 可以理解为Java中的HashMap<String,Integer>
	var testMap map[string]int
	fmt.Println(testMap)

	testMap2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// map[one:1 three:3 two:2]
	fmt.Println(testMap2)

	// 通过make
	var testMap3 = make(map[string]int)
	testMap3["one"] = 1
	testMap3["two"] = 2
	testMap3["three"] = 3
	fmt.Println(testMap3)

	// 通过make指定初始容量（超出会自动扩容）
	var testMap4 = make(map[string]int, 2)
	testMap4["one"] = 1
	testMap4["two"] = 2
	testMap4["three"] = 3
	fmt.Println("len(testMap4)", len(testMap4))

}

// map 要初始化后才能进行赋值操作，如果只是声明，则map 的值为 nil，在nil上操作会提示 panic
// map 的底层还是 hash，同样会存在 hash冲突问题，这个时候go还会判断原始键的值是否相等，因此，在声明map的key 的类型的时候
// 要求 数据类型必须是支持通过 == 或 != 进行判等操作的类型， 比如int、string等
func modifyMap() {
	var testMap3 = make(map[string]int)
	testMap3["one"] = 1
	testMap3["two"] = 2
	testMap3["three"] = 3

	// 查找元素, 第一个返回值是存储的value，第二个bool类型的值是是否找到
	value, contains := testMap3["one"]
	if contains {
		fmt.Println("testMap3[\"one\"] value is", value)
	} else {
		fmt.Println("testMap3 is not contains one")
	}

	// delete
	delete(testMap3, "one")
	// map未包含，这个调用无任何影响
	delete(testMap3, "abc")
}

func iterateMap() {
	var testMap3 = make(map[string]int)
	testMap3["one"] = 1
	testMap3["two"] = 2
	testMap3["three"] = 3

	for key, value := range testMap3 {
		fmt.Println(key, value)
	}

	for _, value := range testMap3 {
		fmt.Println(value)
	}

	for key := range testMap3 {
		fmt.Println(key)
	}

	// kv 键值对调
	invMap := make(map[string]string, 3)
	invMap["one"] = "1"
	invMap["two"] = "2"
	invMap["three"] = "3"
	for k, v := range invMap {
		invMap[v] = k
	}
	for k, v := range invMap {
		fmt.Println(k, v)
	}
}

// 对map进行排序，通过为 key 和 value 创建切片，然后对切片进行排序
func sortMap() {
	invMap := make(map[string]int, 3)
	invMap["one"] = 1
	invMap["two"] = 2
	invMap["three"] = 3

	keys := make([]string, 0)
	for k := range invMap {
		keys = append(keys, k)
	}

	// 对key 进行排序
	sort.Strings(keys)
	fmt.Println("sort by keys:", keys)

	values := make([]int, 0)
	for _, val := range invMap {
		values = append(values, val)
	}

	sort.Ints(values)
	fmt.Println("sort by values:", values)

}
