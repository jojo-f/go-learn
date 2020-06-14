package example02

func init() {
	// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
	// 常量必须为声明时赋值,编译时推导
	// 在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len()
	// 数字常量可以使用任意精度而不用担心溢出问题
	const PI = 3.14159
	// go 中的变量类型可以被隐式推导
	// 常量的并行赋值
	const a, b, c = 1, 2, 3
	const (
		x, y = 1, 2
		z    = 3
	)
	// iota  在const 关键字出现时被重置为 0 ，而后每增加一个常量 iota 会自增 1
	// 下一个未赋值常量将会继承上一个iota表达式,或是常规值
	const (
		unknow = iota      // 0		iota = 0
		male               // 1
		female             // 2
		d      = 2         // 2
		e                  // 2
		f      = 1 << iota // 32 1*2^5		iota = 5
		g      = iota      // 6		iota = 6
		h                  // 7
	)
	println("====== Example 02 =========")
	println(unknow, male, female, d, e, f, g, h, 1<<3)
}
