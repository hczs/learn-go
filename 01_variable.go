package main

import "fmt"

func main() {
	// ======================== 变量声明 ========================
	// 纯粹的声明变量
	var v1 int

	// 声明多个变量
	var (
		v2 int
		v3 string
	)

	// 支持多种数据类型
	// 整型
	var v4 int
	// 字符串
	var v5 string
	// 布尔类型
	var v6 bool
	// int 数组
	var v7 [10]int
	// 结构体 有一个成员变量 f 是 64 位浮点型
	var v8 struct{
		f float64
	}
	// 指针 指向一个整型
	var v9 *int
	// map 字典 key 是字符串 value 是整型
	var v10 map[string]int
	// 函数 入参是整型 返回值也是整型
	var v11 func(a int) int

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)
	fmt.Println("v3:", v3)
	fmt.Println("v4:", v4)
	fmt.Println("v5:", v5)
	fmt.Println("v6:", v6)
	fmt.Println("v7:", v7)
	fmt.Println("v8:", v8)
	fmt.Println("v9:", v9)
	fmt.Println("v10:", v10)
	fmt.Println("v11:", v11)

	// ======================== 变量初始化 ========================
	// 常规初始化操作
	var v12 int = 12
	// 省略类型声明 编译器会自动推导类型
	var v13 = 13
	// 「一般都用这个」省略 var 和类型声明 编译器会自动推导类型
	// := 左侧的类型必须是未声明过的变量
	v14 := 14

	fmt.Println("v12:", v12)
	fmt.Println("v13:", v13)
	fmt.Println("v14:", v14)

	// ======================== 变量赋值和多重赋值 ========================
	// 赋值跟其他语言没啥区别
	var v15 int
	v15 = 15
	fmt.Println(v15)
	// 多重赋值 下面语句是交换 i 和 j 的值
	i := 100
	j := 50
	fmt.Printf("交换前：i: %d j: %d\n", i, j)
	i, j = j, i
	fmt.Printf("交换后：i: %d j: %d\n", i, j)

	// ======================== 匿名变量 ========================
	// 匿名变量通过下划线 _ 来声明，任何赋予给它的值都会被丢弃
	_, nickName := GetName()
	// fmt.Println("_:", _) // Cannot use '_' as a value
	fmt.Println("nickName:", nickName)
}

func GetName() (userName, nickName string) {
	return "sunnyc", "kbdwn"
}
