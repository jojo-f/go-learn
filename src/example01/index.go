package example01

// 类型定义
type (
	IZ   int
	FLAG bool
)

var a = 1

// go 程序的结构,包导入->变量定义->init函数->main->其他函数
func init() {

	println("======= Example01=========")
	// 函数的返回值
	var a int
	var b int
	a, b = sum(1, 2)
	println("1 + 2 =", a, b)
	// init 函数会在 mai 函数运行前运行,即在包被导入时运行
	println("这是 init 函数")
	// println 函数仅用于调试阶段的打印,部署阶段使用 fmt 包中的函数
	print("print")
	println("println")
	// 类型可以是基本类型，如：int、float、bool、string；结构化的（复合的），如：struct、array、slice、map、channel；只描述类型的行为的，如：interface。
	// 结构化的类型没有真正的值，它使用 nil 作为默认值
	// IZ 底层使用 int 类型,所以可以互转,go只允许显示转换
	var c = IZ(a)
	println("c IZ = ", c)
	// 在go 中使用大小写的命名规范
}

func private() {
	println("小写 私有函数 private")
}

func Public() {
	println("大写 公用函数 Public")
}

// 返回变量不可与参数变量同名
func sum(a, b int) (x, y int) { return (a + b), 4 }
