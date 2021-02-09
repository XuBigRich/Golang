package main

import (
	"fmt"
	"strings"
)
/**
字符串：
		Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
		Go 语言里的字符串的内部实现使用UTF-8编码。
		字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符
		例如：
			s1 := "hello"
			s2 := "你好
		这里强行解释一波 经常在Java的字符串中可以看到后面跟着的ascii码，估计在go语言中就看不到了

		Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。
		转义符	含义
			\r	回车符（返回行首）
			\n	换行符（直接跳到下一行的同列位置）
			\t	制表符
			\'	单引号
			\"	双引号
			\\	反斜杠
		相应方法：
			len(str)	求长度
			+或fmt.Sprintf	拼接字符串
			strings.Split	分割
			strings.contains	判断是否包含
			strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
			strings.Index(),strings.LastIndex()	子串出现的位置
			strings.Join(a[]string, sep string)	join操作
 */
func main(){


	//求字符串长度
	s3 := "abc"
	fmt.Println(len(s3))

	//拼接字符串
	s4 := "def"
	fmt.Println(s3 + s4)
	s5 := fmt.Sprintf("%s===========%s", s3, s4)
	fmt.Println(s5)

	//以xx字符串分割
	s6 := "这个字符串要被分割，1，2"
	var ret []string = strings.Split(s6, ",")
	fmt.Println(ret)

	//判断是否包含某些字符串
	s7 := "判断这个字符串里面是否包含许大富"
	ret2 := strings.Contains(s7, "许大富.")
	fmt.Println(ret2)

	//判断是否以xxx开头xxx结尾
	s8 := "争渡争渡，惊起一滩鸥鹭"
	fmt.Printf(`
		判断"%s"这个字符串
		是否以"争渡"作为前缀:%t,
		是否以"许"作为前缀:%t,
		是否以"鸥鹭"作为后缀:%t,
		是否以"鸿志"作为后缀:%t
		`, s8, strings.HasPrefix(s8, "争渡"), strings.HasPrefix(s8, "许"), strings.HasSuffix(s8, "鸥鹭"), strings.HasSuffix(s8, "鸿志"))

	//判断字符串出现的位置 -1代表没找到
	s9 := "apple ,pen"
	fmt.Println(strings.Index(s9, "p"))
	fmt.Println(strings.LastIndex(s9, "p"))
	fmt.Println(strings.LastIndex(s9, "争渡"))
	fmt.Println(strings.LastIndex(s8, "争渡"))

	//join 将数组通过指定的字符连接起来
	a1 := []string{"Java", "C", "Golang"}
	fmt.Println(strings.Join(a1, "-"))
}


