package main

import "fmt"

/**
大纲：
	特性：
		go语言的变量只要声明了就必须要使用
		var 是一个关键字，代表着变量， var作为前缀表示 告诉系统我要声明一个变量，申请一块内存地址
		go语言声明的变量如果没指定值 他们都是有默认值的

类型推导：go语言还支持类型推导，go语言的类型推导，无需明确的指明数据类型，go语言会直接根据值判断数据类型
批量声明：go语言还支持批量声明
短变量声明：go语言支持短变量声明 但只能在方法块中使用
占位符替换：go语言也支持c语言那种的占位符替换方法
默认值：go语言变量的默认值
匿名变量：go语言支持匿名变量生成 匿名变量其实相当于一个垃圾桶 匿名变量用 '_' 表示，匿名变量可以不被调用,也不可以被调用
*/
//批量声明
var (
	//默认值
	a string
	b int
	c bool
	d,f=20,21
)

//匿名变量所调用的函数 定义一个有两个返回值的函数
func foo() (string, int) {
	return "许大富", 24
}

func main() {

	//最原始变量声明方式
	var name string = "许鸿志"
	fmt.Println(name)

	//类型推导
	var age = 20
	fmt.Println(age)

	//短变量声明
	sex := "男" //var sex string = "男"
	fmt.Println(sex)

	//默认值
	fmt.Println(a, b, c,d,f)

	//占位符替换
	fmt.Printf("我叫%s\n", name)

	//匿名变量：使用匿名变量接收一个不需要的值
	username, _ := foo()
	fmt.Println(username)



}
