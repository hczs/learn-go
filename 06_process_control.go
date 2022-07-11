package main

import "fmt"

func main() {
	fmt.Println("======================== 条件语句 ========================")
	condition := true
	// if
	if condition {
		// do something
	}
	// if...else...
	if condition {
		// do something
	} else {
		// do something
	}
	// if...else if...else...
	condition1 := true
	condition2 := true
	if condition1 {
		// do something
	} else if condition2 {
		// do something else
	} else {
		// catch-all or default
	}

	fmt.Println("======================== 分支语句 ========================")
	// 可以使用 switch 来替代 if...else if...else...
	// go 的 switch 不用显示 break
	score := 100
	// 注意：如果判断不是相等 就是直接switch就行，不用加变量
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80 && score < 90:
		fmt.Println("Grade: B")
	case score >= 70 && score < 80:
		fmt.Println("Grade: C")
	case score >= 60 && score < 70:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
	score1 := 100
	// 注意：如果 switch 后跟了变量了，就是判断与 case 后的值是否相等了
	switch score1 {
	case 90, 100:
		fmt.Println("Grade: A")
	case 80:
		fmt.Println("Grade: B")
	case 70:
		fmt.Println("Grade: C")
	case 60:
	case 65:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
	// 合并分支
	// 通过 fallthrough 关键字 如果 正好等于 60 了，还会执行后面分支的代码
	// 就相当于合并 case 60 和 case 65 这两个分支语句了 但是不会执行 case 100，它只会合并紧跟在后面的分支
	score2 := 60
	switch score2 {
	case 60:
		fallthrough
	case 65:
		fmt.Println("Grade: D")
	case 100:
		fmt.Println("100!")
	}

	fmt.Println("======================== 循环语句 ========================")
	// 注意 go 中循环只支持 for，不支持 while 和 do while 语句
	// 基本使用 计算 1 到 100 的数字和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("0 到 100 的数字和：", sum)
	sum = 0
	i := 1
	// 无限循环 计算 1 到 100 的数字和
	for {
		if i > 100 {
			break
		}
		sum += i
		i++
	}
	fmt.Println("0 到 100 的数字和：", sum)

	// for range 结构
	// 参考 slice 和 map 中遍历方式
	// slice
	summer := make([]int, 5)
	for i, v := range summer {
		fmt.Println("索引:", i, "值:", v)
	}
	// map
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

	// for 可以指定 break 哪个循环
outerLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break outerLoop
			}
			fmt.Println(i)
		}
	}

	fmt.Println("======================== 跳转语句 ========================")
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				// 会直接跳转到 EXIT 标签代码的位置
				goto EXIT
			}
			fmt.Println(num)
		}
	}
EXIT:
	fmt.Println("Exit.")

}
