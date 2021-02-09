package main

import "fmt"

/**
大纲：
	new与make区别
	1.make与new都是用来申请内存的
	2. new一般用于给基础数据类型申请内存,string int ....  返回对应类型的指针
	3.make 一般用于slice、map、chan 这些类型申请内存 返回对应类型的本身
	make:
		make申请内存空间,用于给slice、map、chan 初始化
	new:
		new是用来创建指针的，使用new(T type) 会返回一个指针类型 用于给基本数据类型申请内存
	point:
		指针分为取值和取值 &是取指符，*是取值符
			指针和值是指的同一个内存代码块。
				& 指针是描述位置的，
				* 值是描述位置上面的数据的
	nil:
		空值/零值
		按照Go语言规范，任何类型在未初始化时都对应一个零值：布尔类型是false，整型是0，字符串是""，
		而指针，函数，interface，slice，channel和map的零值都是nil。

*/

//point
func point() {
	//1.&:取地址
	n := 10
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	//2.*:根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Println(m == n)
}

//new
func useNew() {
	//错误示范  声明一个int类型的指针a
	var a *int
	//将100赋值给指针为a的地址上
	//*a = 100 这句代码会抛出 panic: runtime error: invalid memory address or nil pointer dereference
	//===========================
	//如过要使这行代码不再报错需要使用new去开辟内存空间
	//a=new(int)
	//===========================
	//若直接打印a 会到的一个nil值
	fmt.Println(a)
	*a = 100
	//取出指针地址为a的地址值
	fmt.Println(*a)

	//用于申请一个值类型内存地址 得到一个内存地址
	b := new(int)
	fmt.Println(b)


}

//make
func useMake() {
	//用来申请引用类型地址
	a := make([]int, 1, 2)
	//对于引用类型，只声明了，却没有初始化申请地址，所以这个值类型不可用
	var s1 []int
	//这样赋值是错的
	//s1[0]=100
	//这样是可以的
	s1 = append(s1, 100)
	//只有声明，并初始化后的引用类型才可以使用
	fmt.Println(a, s1)
}

//nil
func useNil() {
	//值类型
	var a int
	var b string
	var c bool
	var d float64
	fmt.Println(a, b, c, d)
	//引用类型 在未进行初始化的时候所对应的值是nil，在使用这些引用类型时一定别忘记初始化
	var e *int
	var f []int
	var g map[string]int
	fmt.Println(e, f, g)
	fmt.Println(e == nil, f == nil, g == nil)

}
func main() {
	//useMake()
	useNew()
	//point()
	//useNil()
}
