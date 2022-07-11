package main

import "fmt"

func main() {
	// ======================== 字典常用方式 ========================
	// 中括号里面是 key 类型 后面那个是 value 类型
	var testMap map[string]int
	testMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("The element of key %s; %d\n", k, v)
	} else {
		fmt.Println("Not found")
	}

	// ======================== 字典声明和初始化 ========================
	// 先声明变量 再复制内容
	var m1 map[string]int
	m1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// 声明同时初始化内容
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	// 也可以用 make 来创建一个字典
	var m3 = make(map[string]int)
	// 创建时也可以指定容量
	var m4 = make(map[string]int, 10)
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)
	fmt.Println("m3:", m3)
	fmt.Println("m4:", m4)

	// ======================== 添加/修改元素 ========================
	// 和 py 差不多 直接添加 key value 键值对
	m3["one"] = 1
	// 对已有的 key 赋值就是修改了
	m3["one"] = 2
	fmt.Println("m3:", m3)
	// 特别注意，以下代码是有问题的，这种只声明，但是没有初始化的map，是没法赋值的
	//var m5 map[string]int
	//m5["a"] = 1
	//fmt.Println("m5:", m5)

	// ======================== 查找元素 ========================
	m6 := map[string]int{
		"one": 1,
		"two": 2,
	}
	// 从 m6 中查找 one 的值 返回的第一个参数是具体值 第二个参数是是否找到了
	value, ok := m6["one"]
	if ok {
		fmt.Println("m6[\"one\"] value:", value)
		fmt.Println("m6[\"one\"] ok:", ok)
	}
	value, ok = m6["three"]
	fmt.Println("m6[\"three\"] value:", value)
	fmt.Println("m6[\"three\"] ok:", ok)

	// ======================== 删除元素 ========================
	// 删除 m3 中 key 为 one 的键值对
	fmt.Println("删除one之前的m6:", m6)
	delete(m6, "one")
	fmt.Println("删除one之后的m6:", m6)
	// 删除 m3 中 key 为 three 的键值对 其实没有 three 这个 key，所以执行没啥影响 也不会报错
	fmt.Println("删除three之前的m6:", m6)
	delete(m6, "three")
	fmt.Println("删除three之后的m6:", m6)

	// ======================== 遍历字典 ========================
	// 先初始化一个字典
	var m7 map[int]string
	m7 = map[int]string{
		12345: "aa",
		22:    "bb",
		666:   "cc",
	}
	// 遍历 m7
	for key, value := range m7 {
		fmt.Println(key, value)
	}
	// 只获取字典的所有 key
	for key := range m7 {
		fmt.Println(key)
	}
	// 借助匿名变量，只获取字典的值
	for _, value := range m7 {
		fmt.Println(value)
	}
}
