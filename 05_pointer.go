package main

import "fmt"

func main() {
	// ======================== 指针入门简单使用 ========================
	a := 100
	var ptr *int
	ptr = &a
	// a 变量实际的内存地址
	fmt.Println(ptr)
	// 加上 * 之后 就是打印这个指针具体指向的值是多少
	fmt.Println(*ptr)

	// ======================== 指针声明和初始化 ========================
	// 可以像上面那样先声明，再初始化 不过需要注意，如果声明的是 *int 即指向 int 类型的指针，就不能指向其他类型了
	b := "aa"
	// 这样赋值会报错 Cannot use '&b' (type *string) as the type *int
	//ptr = &b
	// 可以使用 := 声明指针并初始化 会自动推断类型
	ptr1 := &b
	fmt.Println("ptr1:", ptr1)
	fmt.Println("*ptr1:", *ptr1)
	// 也可以通过内置函数 new 声明指针
	ptr2 := new(int)
	c := 99
	ptr2 = &c
	fmt.Println("ptr2:", ptr2)
	fmt.Println("*ptr2:", *ptr2)

	// ======================== 通过指针传值 ========================
	t1, t2 := 1, 2
	swap(t1, t2)
	// 实际是没有换值的 想要真正交换这两个值需要传这两个值的引用
	fmt.Println(t1, t2)
	// 使用传递指针（值的引用）
	t3, t4 := 3, 4
	swapByPointer(&t3, &t4)
	fmt.Println(t3, t4)
}

func swap(a, b int) {
	// 这个交换的只是这个方法的局部变量 a b
	a, b = b, a
	fmt.Println("swap方法内：", a, b)
}

func swapByPointer(a, b *int) {
	// 改变指针具体指向的值
	*a, *b = *b, *a
	fmt.Println("swapByPointer方法内：", *a, *b)
}
