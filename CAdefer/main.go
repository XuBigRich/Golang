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
*/

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿")
	fmt.Println("end")
	defer fmt.Println("呵呵呵")
}
func main() {
	deferDemo()
}
