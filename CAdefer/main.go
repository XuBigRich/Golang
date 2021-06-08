package main

import "fmt"

/**
大纲：
	Go语言中defer语句会将其后面跟随的语句进行延迟处理，在defer归属的函数即将返回时，将延迟处理语句按defer定义的逆序进行执行，
	也就是说，先被defer定义的语句最后被执行，最后定义的defer语句先被执行，一个函数可以拥有多个defer
理解：
	你可以这样理解，当执行函数时，函数会正常执行，但遇到defer关键字，会将defer后面的代码块压入待执行栈
	当函数执行完毕之后，会执行栈中的代码，因为栈是先入后出，所以defer的执行顺序是倒着执行的
使用场景：
	当使用数据库连接时 要关闭连接，这样就可以将打开与关闭连接写在一起了
	当使用io流时，打开了文件句柄，还需要关闭文件句柄 ，也是可以将打开与关闭连接写在一起
	..........
defer：
	defer多用于函数结束之前的资源释放
汇编层面：
	在汇编层面 函数的return不是一个原子操作，它包含了对变量的赋值，与方法的返回。
	汇编实现：
		当执行return语句时。
		函数中return语句底层实现				defer语句执行的时机
			返回值=x								返回值=x
			RET指令								运行defer
												RET指令
	用语言描述return的汇编步骤：
		第一步：返回值赋值
		第二步：真正的RET指令
	如果函数种存在defer语句对已经赋值的返回值进行了更改，那么会对返回值产生影响，因为defer运行于 第一步与第二步之间。
*/

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿")
	fmt.Println("end")
	defer fmt.Println("呵呵呵")
}

//下面介绍几个刁钻的例子
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
	//分解步骤
	//result=5
	//x++
	//return result=5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
	//分解步骤
	//x=5
	//x++
	//return x=6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
	//分解步骤
	//y=5
	//x++
	//return y=5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
	//分解步骤
	//x=5
	//x拷贝一份值传给匿名方法 由匿名变量做++运算 被拷贝传入的变量值变为6
	//return x=5
}

func add(x int) int {
	return x + 1
}
func deferDemo1() (x int) {
	x = 1
	defer add(x)
	return x
}
func deferDemo2() int {
	x := 1
	defer add(x)
	return x
}
func main() {
	deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	deferDemo()
	fmt.Println(deferDemo1())
	fmt.Println(deferDemo2())
}
