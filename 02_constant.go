package main

import "fmt"

func main() {
	// ======================== 常量使用 ========================
	// 特别注意：大写字母开头的常量包外可见 小写字母开头的常量只能在包内访问 函数体内声明的常量仅在函数体中有效
	// 使用 const 来定义常量
	// 常量只可以是数值类型（包括整型、 浮点型和复数类型）、布尔类型、字符串类型等标量类型
	const Pi float64 = 3.14159265358979323846
	// 无类型浮点常量
	const zero = 0.0
	// 一个 const 关键字定义多个常量
	const (
		size float64 = 1024
		// 也可以不写类型 自动推导
		eof = -1
	)
	// 常量多重赋值
	const u, v float32 = 0, 3
	// 也可以不写类型 自动推导
	const a, b, c = 1, 2, "ccc"

	// ======================== 预定义常量 ========================
	// go 预定义了三个常量 true false 和 iota
	// true 和 false 就是布尔类型的真假值
	// iota 比较特殊，可以被认为是一个可被编译器修改的常量，在每一个 const 关键字出现时被重置为 0
	// 然后在下一个 const 出现之前，每出现一次 iota，其所代表的数字会自动增 1
	const (    // iota 被重置为 0
		c0 = iota   // c0 = 0
		c1 = iota   // c1 = 1
		c2 = iota   // c2 = 2
	)
	const (
		l = iota * 2  // l = 0
		m = iota * 2  // m = 2
		n = iota * 2  // n = 4
	)
	const x = iota  // x = 0
	const y = iota  // y = 0
	// 如果两个 const 的赋值语句的表达式是一样的，那么还可以省略后一个赋值表达式
	const (
		// 0
		a1 = iota * 3
		// 1 * 3 = 3
		a2
		// 2 * 3 = 6
		a3
	)

	fmt.Println("c0:", c0)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("l:", l)
	fmt.Println("m:", m)
	fmt.Println("n:", n)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	// ======================== 常量还可用于枚举 ========================
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println("Sunday", Sunday)
	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Friday", Friday)
	fmt.Println("Saturday", Saturday)
	fmt.Println("numberOfDays", numberOfDays)

}
