package example03

import (
	"fmt"
	"os"
	"runtime"
)

// 每个源文件都只能包含一个 init 函数
// 在每个包完成初始化后自动执行，并且执行优先级比 main 函数高

func init() {
	println("===== example 03 =====")
	// 自动推导的局限性,此时必须使用类型显式声明
	const n int64 = 2
	var (
		HOEM   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	println(HOEM, USER, GOROOT, 1)
	// 使用简写变量初始化声明,隐式推导
	goos := runtime.GOOS
	fmt.Printf("这个操作系统为 %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("path 为路径 %s\n", path)
	// 全局变量是允许声明但不使用
	a, b := 1, 2
	a, b = b, a
	println(a, b) // 2 1
}
