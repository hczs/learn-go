package main

import "fmt"

func main() {
	// ======================== 创建切片 ========================
	// 基于数组切片
	// ... 的方式其实是语法糖，可以省略数组长度的声明，会自动计算数组长度
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	// 取第 3 4 5 个元素，左闭右开
	q2 := months[3:6]
	// 夏季 取第 5 6 7 个元素
	summer := months[5:8]
	// 基于 months 所有元素创建切片
	all := months[:]
	// 基于 months 前 6 个元素创建切片 前面的 0 可以省略
	firsthalf := months[:6]
	// 基于 months 后 6 个元素创建切片 最末尾也可以省略
	secondhalf := months[6:]
	fmt.Println("第二季度：", q2)
	fmt.Println("第二季度：", summer)
	fmt.Println("基于 months 所有元素创建切片（全年）：", all)
	fmt.Println("上半年：", firsthalf)
	fmt.Println("下半年：", secondhalf)

}
