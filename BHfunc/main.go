package main

import "fmt"

/**
大纲：
	函数的意义：
		函数时一段代码的封装，把一段逻辑抽象出来封装到一个函数中，给他取一个名字，每次用到他时直接使用函数名就可以调用了。
		使用函数可以让代码结构更清晰、更简洁，
		******特性**********
		go语言中所有的给方法传参都是拷贝类型，也就是Java中的值传递。
		Java语言值传递
			"和Java一个鸟样Java里面叫做值传递，就是将要传递的数值复制一份传给方法，但go语言比Java更绝
			因为Java中基本数据类型传递时，会传递数值，而对象数据类型与数组数据类型都是将 对象与数组的内存地址拷贝一份给函数
			注意是对象与数组的地址拷贝一份，地址拷贝一份，所以当在函数中更改传入的对象/数组时，外面的数组也会随着变化
			但根本上其实都是值传递"
	函数的定义：f 系列的方式，最原始的函数声明
	定义一个没有返回值的函数：
	定义一个没有参数的函数：
	定义一个没有参数 ,没有返回值的：
	多个返回值：
	参数类型简写：
	可变长参数：
	函数声明改进：返回值的参数省略，与含有返回值参数的简单使用
	双返回值改进：
	函数声明改进：
*/
//函数的定义: func声明是一个函数 sum为函数名，第一个()入参，第二个()返回值 没有时可省略
func f0(x int, y int) (ret int) {
	return x + y
}

//定义一个没有返回值的函数
func f1(x int, y int) {
	fmt.Println(x + y)
}

//定义一个没有参数的函数
func f2() (ret int) {
	return 1
}

//定义一个没有参数 ,没有返回值的
func f3() {
	fmt.Println("我干啥了？")
}

//多个返回值
func f4() (int, string) {
	return 1, "许大富"
}

//参数类型简写：
//	当参数中连续两个或两个以上参数的类型相同时，我们可以将非最后一个参数的类型省略
func f5(x, y int, z, h, m string) int {
	return x + y
}

//可变长参数   传入的是一个切片类，切与Java一样可变长参数必须放在函数参数最后
func f6(x ...int) {
	for _, v := range x {
		fmt.Print(v, " ")
	}
}

//函数声明改进 返回值的参数省略
//以f0为例：其实返回值的参数是可以省略的
func ff00(x int, y int) int {
	return x + y
}

//继续 改进 其实不建议返回值参数省略因为如果没有省略可以直接将参数赋值
func ff01(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以将return后面的代码省略  return ret
}

//双返回值改进
func ff40() (order int, name string) {
	order = 1
	name = "许大富"
	return
}

func main() {
	r := f0(1, 2)
	fmt.Println(r)
	f1(1, 2)
	fmt.Println(f2())
	f3()
	_, name := f4()
	fmt.Println(name)
	fmt.Println(f5(1, 2, "1", "2", "3"))
	f6(1, 2, 3, 4, 5, 6)

}
