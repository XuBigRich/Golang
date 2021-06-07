package main

import "fmt"

/**
大纲：
	在go语言中函数也是一种类型，函数的类型可以通过fmt %T打印出来

	当作函数的参数:函数还可以当作函数的参数传给函数这是一个go语言的新特性
	函数作为返回值:函数还可以当作函数的返回值
*/
//声明一个无参无返回值的函数
func f1() {

}

//声明一个无参有返回值的函数
func f2() int {
	return 100
}

//声明一个有参有返回值的函数
func f3(x int) int {
	return x
}

//打印
func f4() {
	fmt.Printf("f1的类型%T \n", f1)
	fmt.Printf("f1的类型%T \n", f2)
	fmt.Printf("f1的类型%T \n", f3)
	var t1 = f1 //还可以将函数赋值给变量
	fmt.Println(t1)
}

//当作函数的参数，声明一个函数，接收一个入参为int 返回值为int的函数
func f5(f func(int) int) int {
	return f(20)
}

//返回一个匿名函数类型
func f6() func(int) int {
	return func(i int) int {
		return i * 20
	}
}

func main() {
	f4()
	fmt.Println(f5(f3))
	//使用()表示执行函数
	fmt.Println(f6()(2))
}
