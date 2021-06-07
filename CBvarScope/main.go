package main

import "fmt"

/**
	大纲：
		作用域分为全局作用域与局部作用域
		函数中查找变量的顺序是
			1.现在函数内部查找
			2.找不到在函数外部查找，一直找到全局变量
 */
//全局作用域
var x=100
//定义一个函数
func f1()  {
	fmt.Println(x)
}
//演示if的局部变量作用域
func f2()  {
	if i:=0;i<10{
		fmt.Println(i)
	}

	//fmt.Println(i)
}

func main() {
	f1()
}
