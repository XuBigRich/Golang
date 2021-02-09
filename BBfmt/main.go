package main

import "fmt"

/**
fmt输出类型
通用：
		%v	值的默认格式表示
		%+v	类似%v，但输出结构体时会添加字段名
		%#v	值的Go语法表示
		%T	值的类型的Go语法表示
		%%	百分号
布尔值：
		%t	单词true或false
整数：

		%b	表示为二进制
		%-5d 对齐与补位 指定最小宽度为5  左对齐
		%5d  指定最小宽度为5 右对齐
		%05d 指定最小宽度为5 前面补0
		%c	该值对应的unicode码值
		%d	表示为十进制
		%o	表示为八进制
		%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
		%x	表示为十六进制，使用a-f
		%X	表示为十六进制，使用A-F
		%U	表示为Unicode格式：U+1234，等价于"U+%04X"
		%%	百分号转译

浮点数与复数的两个组分：
		%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
		%e	科学计数法，如-1234.456e+78
		%E	科学计数法，如-1234.456E+78
		%f	有小数部分但无指数部分，如123.456
		%F	等价于%f
		%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

字符串和[]byte：
		%s	直接输出字符串或者[]byte
		%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
		%x	每个字节用两字符十六进制数表示（使用a-f）
		%X	每个字节用两字符十六进制数表示（使用A-F）
指针：
		%p	带ox的指针
		%#p	不带ox的指针
*/

var (
	a = 120
	b = 'x'
	c = '许'
	d = "许鸿志"
	e = false
	f = 3.141592654
	g = &a
	//数组
	h = [...]byte{'x', 'h', 'z'}
	//切片
	i = []byte{'x', 'h', 'z'}
)

//通用
func base() {
	//按照默认的类型输出
	fmt.Println(a, b, c, d, e, f, g, h)
	//如果不知道是那种类型那就用%v 效果与 上面一样
	fmt.Printf("a=%v,b=%v,c=%v,d=%v,e=%v,f=%v,g=%v,h=%v,i=%v \n", a, b, c, d, e, f, g, h, i)
	//%+v 效果与 上面一样 但输出结构体时会添加字段名
	fmt.Printf("a=%+v,b=%+v,c=%+v,d=%+v,e=%+v,f=%+v,g=%+v,h=%+v,i=%+v \n", a, b, c, d, e, f, g, h, i)
	//%#v	值的Go语法表示 字符串添加了双引号
	fmt.Printf("a=%#v,b=%#v,c=%#v,d=%#v,e=%#v,f=%#v,g=%#v,h=%#v,i=%#v \n", a, b, c, d, e, f, g, h, i)
	//输出相应的类型
	fmt.Printf("a=%T,b=%T,c=%T,d=%T,e=%T,f=%T,g=%T,h=%T,i=%T \n", a, b, c, d, e, f, g, h, i)
}

//布尔
func boolean() {
	fmt.Printf("e=%t \n", e)
}

//整数
func integer() {
	//输出二进制
	fmt.Printf("a=%b \n", a)
	//对齐与补位 左对齐
	fmt.Printf("|%-5d| \n", a)
	//右对齐
	fmt.Printf("|%5d| \n", a)
	//前面补0
	fmt.Printf("|%05d| \n", a)
	//%c	该值对应的unicode码值  120对应的unicode 是字符x
	fmt.Printf("a=%c \n", a)
	//%d	表示为十进制
	fmt.Printf("a=%d \n", a)
	//%o	表示为八进制
	fmt.Printf("a=%o \n", a)
	//%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("a=%q \n", a)
	//%x	表示为十六进制，使用a-f
	fmt.Printf("a=%x \n", a)
	//%X	表示为十六进制，使用A-F
	fmt.Printf("a=%X \n", a)
	//%U	表示为Unicode格式：U+1234，等价于"U+%04X"
	fmt.Printf("a=%U \n", a)
	//%%	描述一个百分之120 这里的%好不可以使用\%进行转译
	fmt.Printf("a=%b%% \n", a)
}

//浮点数与复数的两个组分：
func float() {
	//%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
	fmt.Printf("f=%b \n", f)
	//%e	科学计数法，如-1234.456e+78
	fmt.Printf("f=%e \n", f)
	//%E	科学计数法，如-1234.456E+78
	fmt.Printf("f=%E \n", f)
	//%f	有小数部分但无指数部分，如123.456
	fmt.Printf("f=%f \n", f)
	//%2f	保留小数点后两位
	fmt.Printf("f=%.2f \n", f)
	//%F	等价于%f
	fmt.Printf("f=%F \n", f)
	//%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	fmt.Printf("f=%g \n", f)
	//%.2g	输出总共2个数字
	fmt.Printf("f=%.2g \n", f)
	//%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	fmt.Printf("f=%G \n", f)
}

//字符串和[]byte
func chars() {
	//%s	直接输出字符串或者[]byte
	fmt.Printf("h=%s \n", h)
	fmt.Printf("d=%s \n", d)
	//最小宽度1最大宽度2
	fmt.Printf("d=%1.2s \n", d)
	//%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	fmt.Printf("h=%q \n", h)
	fmt.Printf("d=%q \n", d)
	//%x	每个字节用两字符十六进制数表示（使用a-f）
	fmt.Printf("h=%x \n", h)
	fmt.Printf("d=%x \n", d)
	//%X	每个字节用两字符十六进制数表示（使用A-F）
	fmt.Printf("h=%X \n", h)
	fmt.Printf("d=%X \n", d)
}

//指针
func point() {
	fmt.Printf("g=%p \n", g)
	fmt.Printf("g=%#p \n", g)
}
func main() {
	base()
	//boolean()
	//integer()
	//float()
	//chars()
	//point()
}
