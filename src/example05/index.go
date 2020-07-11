package example05

import (
	"fmt"
	"unicode/utf8"
)

func init() {
	println("======= example05 字符 =======")
	println("解释字符串 具有反斜杠转义 \\n \n")
	println(`非解释字符串 反斜杠不具有转义特性也可以直接换行
	\n `)
	// go 中的字符串对比是基于字节比较
	fmt.Printf("获取字符长度 len(wocao) => %d \n", len("wocao"))
	s := "hel" + "lo"
	s += "world!"
	println(s)
	// 使用 strings.Join() 拼接字符更方便
	// 使用下标可直接获得字符的纯字节
	// 练习
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"

	fmt.Printf("获取字节码长度: str1(%d) str2:[%d] \n", len(str1), len(str2))
	fmt.Printf("获取字符长度: str1(%d) str2:[%d] \n", utf8.RuneCountInString(str1), utf8.RuneCountInString(str2))
}
