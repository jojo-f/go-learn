package example04

import (
	"fmt"
	"math"
)

func init() {
	println("========= example04 =========")
	// go 中仅可以比较相同类型的值
	// println(1 == true) err
	// ! 一元运算符仅需要一个值
	// != == && || 为二元运算符至少需要两个值
	// %t 格式化输出 bol 值
	fmt.Printf("bol:%t \n", true)

	// go 的数字类型支持整数,浮点,复数
	// Go 也有基于架构的类型，例如：int、uint 和 uintptr。
	// int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）
	// uintptr 的长度被设定为足够存放一个指针即可。

	// go 中计算最快的类型就是 int 型
	i8 := int8(127)                   // -128 ~ 127
	i16 := int16(32767)               // -32768  ~ 32767
	i32 := int32(2147483647)          // -2,147,483,648 ~ 2,147,483,647
	i64 := int64(9223372036854775807) // -9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807
	fmt.Printf("i8:%d i16:%d i32:%d i64:%d \n", i8, i16, i32, i64)

	// 无符号整数 能表述的数值范围为 int 类型对应的两倍数 0 ~ 255 共 256位 为128*2
	u8 := uint8(255)                    // 0 ~ 255
	u16 := uint16(65535)                // 0 ~ 65,535
	u32 := uint32(4294967295)           // 0 ~ 4294967295
	u64 := uint64(18446744073709551615) // 0 ~ 18446744073709551615
	fmt.Printf("u8:%d u16:%d u32:%d u64:%d \n", u8, u16, u32, u64)

	// %x 用于格式化16进制表示的数字
	fmt.Printf("0xffffff:%x \n", 0xffffff)

	// float32 精确到小数点后 7 位，float64 精确到小数点后 15 位。
	// 尽可能使用 float64 因为 math 包中所有有关数学运算的函数都会要求接收这个类型
	f32 := float32(3.4e+38)                 // +- 1e-45 -> +- 3.4 * 1e38
	f64 := float64(1.7976931348623157e+308) // +- 5 * 1e-324 -> 107 * 1e308
	fmt.Printf("格式化浮点数 f32:%g f64:%g \n", f32, f64)
	fmt.Printf("浮点数 f32:%f f64:%f \n", f32, f64)
	fmt.Printf("科学计数法浮点数 f32:%e f64:%e \n", f32, f64)
	fmt.Printf("精确计数位 f-5:%0.2e f-5:%0.3g f-5:%0.2f \n", f32, f64, f64) // $m.n[e,f,g] // m为任意常数,n为小数位长度 e,f,g 为输出格式

	// go 中部分常量可以不转换类型直接计算
	a := float32(3.2) + 5
	fmt.Printf("常量计算 %f \n", a)

	// 二元运算
	// 按位与 &
	/*
		1 & 1 -> 1
		1 & 0 -> 0
		0 & 1 -> 0
		0 & 0 -> 0
	*/

	// 按位或 |
	/*
		1 | 1 -> 1
		1 | 0 -> 1
		0 | 1 -> 1
		0 | 0 -> 0
	*/

	// 按位异或 ^
	/*
		1 ^ 1 -> 0
		1 ^ 0 -> 1
		0 ^ 1 -> 1
		0 ^ 0 -> 0
	*/

	fmt.Printf("%d \n", 010^101)
	fmt.Printf("%d \n", 100^111)
	fmt.Printf("%d \n", 101^110)

}

// 数值的安全转换
func Uint8FromtInt(u int) (uint8, error) {
	if 0 <= u && u <= math.MaxUint8 {
		return uint8(u), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", u)
}
