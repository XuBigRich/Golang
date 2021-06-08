package main

import (
	"fmt"
	"strings"
)

/**
大纲：
	学习了这个案例之后，我对比Java中的代理思考了很久，因为在我一开始看来，他和Java的代理非常像，因为都在方法中调用了另外一个方法。
	后来仔细思考后，其实闭包和代理有很大的不同，或者说完全不一样，代理是代理类执行目标方法时前后，增强目标方法。做一些其他工作
	而这个闭包则本质其实是在传入方法前，在f3方法中执行过一次f2方法，然后通过校准返回值，让f3的返回值与f1的入参达到一致。
	最终使调用f1方法时，也可以调用f2方法，其实这个操作也可以看作先调用了f2方法 后调用了f1方法，当然 也可以在f3中做一些其他操作从而影响f1与f2
	这样做还有一个更神奇的地方就是，外面的方法有可能会干预内部的方法，这是Java代理所不具备的

	闭包：闭包就是一个存在于函数内部的函数 ，这个内部函数除了引用了内部变量，还引用了其所在函数的外部变量，这样的函数叫做闭包，详情请看f3与f5与f6对闭包的探索。
		闭包是一个函数，这个函数包含了他外部作用域的一个变量,
	最新理解：
		闭包使用了方法内的变量,根据闭包传入的值 与方法内的值进行计算的到结果，闭包存在于一个方法的小世界当中
	闭包的底层原理：
		1.函数可以作为返回值
		2.函数内部查找顺序，先自己内部找，找不到外层找。
	闭包的增强变种与使用小技巧
*/

//定义一个函数
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

//定义一个函数 目的是让f1使用本函数当作参数
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数f3，对f2进行包装
func f3(x, y int) func() {
	//返回的方法将f2进行了包装，类似与Java中的代理模式，与切面编程，这里的temp引用了函数f3的x y变量所以temp是一个闭包
	temp := func() {
		f2(x, y)
	}
	return temp
}

//继续变 这是闭包的精髓之处
func f6() func(int) int {
	x := 100
	return func(i int) int {
		x += i
		return x
	}
}

//变种 ，闭包的操作使用 ，f4对标f2
func f4(x, y int) {
	fmt.Println(x + y)
}

//f5 其实就是f1与f2的结合，让传入的函数去处理m,n两个参数
func f5(f func(x, y int), m, n int) {
	//这个方法将使用外部传入的m，n变量，所以这里的tmp是一个闭包
	tmp := func() {
		f(m, n)
	}
	tmp()
}

//闭包进阶
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

//更优秀的闭包进阶
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	f1(f3(1, 2))
	f5(f4, 2, 5)
	tmp := f6()
	fmt.Println(tmp(200))
	txt := makeSuffixFunc(".txt")
	name := txt("ret")
	fmt.Println(name)
	file := txt("file.txt")
	fmt.Println(file)

	ff1, ff2 := calc(10)
	fmt.Println(ff1(1), ff2(2))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
