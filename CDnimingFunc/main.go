package main

import "fmt"

/**
大纲：
	匿名函数一般使用在函数的内部，因为函数内部不可以再次定义函数，所以函数内部使用其他函数都是以匿名函数的方式使用的
定义一个匿名函数：
使用匿名函数举例：

*/
//定义一个匿名函数
var f1 = func(x int) int {
	return x * 2
}
//使用匿名函数举例
func f2() {
	f3 := func() string {
		return "嘿嘿"
	}
	fmt.Println(f3())
}
//如果函数只调用一次，还可以简写成立即执行函数
func f3()  {
	func(x int) {
		fmt.Println(x)
	}(20)
}

func main() {
	fmt.Println(f1(2))
	f2()
	f3()
}
