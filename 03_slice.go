package main

import "fmt"

func main() {
	fmt.Println("======================== 创建切片 ========================")
	// 基于数组创建切片
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

	// 基于切片也可以创建切片
	// 取上半年的前三个月(0, 1, 2)
	baseSlice := firsthalf[:3]
	fmt.Println("baseSlice", baseSlice)

	// 也可以直接创建切片
	// 使用 make 创建一个初始长度为 5 的切片
	slice1 := make([]int, 5)
	// 创建一个初始长度为 5 容量为 10 的切片
	slice2 := make([]int, 5, 10)
	// 创建同时初始化内容
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)

	fmt.Println("======================== 遍历切片 ========================")
	fmt.Println("开始遍历summer")
	for i := 0; i < len(summer); i++ {
		fmt.Println(summer[i])
	}
	// 使用 range 遍历，代码更加简洁
	// range 返回两个值，第一个是索引，第二个是下标对应的值
	for i, v := range summer {
		fmt.Println("索引：", i, "值：", v)
	}

	fmt.Println("======================== 动态增加元素 ========================")
	// 计算切片实际长度和容量
	capAndLen := make([]int, 5, 10)
	fmt.Println("实际长度:", len(capAndLen)) // 5
	fmt.Println("容量:", cap(capAndLen))   // 10
	fmt.Println("切片内容：", capAndLen)
	// 向后面追加内容
	newSlice := append(capAndLen, 1, 2, 3)
	// 如果切片不扩容，就依然使用的原来的那份内存
	newSlice[0] = 888
	fmt.Println("追加内容并返回新切片：", newSlice)
	fmt.Println("新切片实际长度：", len(newSlice))
	fmt.Println("新切片容量：", cap(newSlice))
	// capAndLen 的零号元素也被同步更改了
	fmt.Println("追加内容前的切片：", capAndLen)
	// 自动扩容 再往新切片后面追加元素
	// 默认情况下扩容是原来的两倍，即 cap(slice) * 2 = cap(newSclie)
	// 当容量大于等于 1024 时，就扩容为原来的 1.25 倍
	autoCap := append(newSlice, 6, 7, 8)
	// 切片扩容后返回的新切片是把原来的数据复制一份，并不会影响原先的切片
	autoCap[0] = 666
	fmt.Println("再次追加内容自动扩容：", autoCap)
	fmt.Println("自动扩容后的实际长度：", len(autoCap))
	fmt.Println("自动扩容后的容量：", cap(autoCap))
	// 扩容前的 newSlice 并没有被扩容后的切片影响
	fmt.Println("自动扩容前的切片：", newSlice)
	// 所以进行了自动扩容了才会复制一份重新分配内存，不自动扩容的 append 还是用的原来的内存

	fmt.Println("======================== 内容复制 ========================")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8}
	fmt.Println("初始化 s1:", s1)
	fmt.Println("初始化 s2:", s2)
	// 复制 s1 内容到 s2 s2 只有三个位置，并不会扩容，同样只会复制 s1 的前三个给 s2
	copy(s2, s1)
	fmt.Println("复制 s1 内容到 s2")
	fmt.Println("复制后的 s1:", s1)
	fmt.Println("复制后的 s2:", s2)

	// 重新改下 s2 内容
	s2 = []int{8, 9, 10}
	// 复制 s2 内容到 s1 同样 s2 只有三个，也只会改变 s1 的前三个
	copy(s1, s2)
	fmt.Println("复制前 s1:", s1)
	fmt.Println("复制前 s2:", s2)
	fmt.Println("复制 s2 内容到 s1")
	fmt.Println("复制后的 s1:", s1)
	fmt.Println("复制后的 s2:", s2)
}
