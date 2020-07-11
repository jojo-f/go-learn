package example06

import (
	"fmt"
	"strings"
)

func init() {
	println("======== strins 和 strconv 包==========")
	fmt.Printf("ef 首字母包含 e %t \n", strings.HasPrefix("ef", "e"))
	fmt.Printf("ef 尾字母包含 f %t \n", strings.HasSuffix("ef", "f"))
	fmt.Printf("abc 字符串中包含 b %t \n", strings.Contains("abc", "b"))
	fmt.Printf("abcabc 字符串中首次出现b的索引 %d \n", strings.Index("abcabc", "b"))
	fmt.Printf("abcabc 字符串中最后一次出现b的索引 %d \n", strings.LastIndex("abcabc", "b"))
	fmt.Printf("阿里巴巴四十大盗 字符串中查询非ASCII码字符 %d \n", strings.LastIndex("阿里巴巴四十大盗", "巴"))
	// 替换匹配到的前1个,-1为所有
	fmt.Printf("阿里巴巴四十大盗四十大盗 字符串替换[四十大盗] %s \n", strings.Replace("阿里巴巴四十大盗四十大盗", "四十大盗", "", 1))
	fmt.Printf("阿里巴巴四十大盗四十大盗 字符串计数[四十]出现的次数 %d次 \n", strings.Count("阿里巴巴四十大盗四十大盗", "四十"))
	fmt.Printf("阿里巴巴 拼接字符串[阿里巴巴]3次 %s \n", strings.Repeat("阿里巴巴", 3))
	// 大小写
	// strings.ToLower("ABC")
	// strings.ToUpper("abc")
	// 剔除首尾字符
	// strings.TrimSpace(" abc ")
	// strings.Trim("_abc_", "_")
	// strings.TrimLeft(" abc ", " ")
	// strings.TrimRight(" abc ", " ")
	fmt.Printf("woshishui 分隔字符串为切片 %s \n", strings.Fields("woshishui?"))
	fmt.Printf("wo,shi,shui 按逗号分隔字符串为切片 %s \n", strings.Split("wo,shi,shui?", ","))
	fmt.Printf("拼接string类型的切片为字符串 %s \n", strings.Join(strings.Split("wo,shi,shui?", ","), "|"))
	// 使用 Reader 读取字符串中的内容
	// strings.Reader("ff") // 返回一个 Reader 指针
}
