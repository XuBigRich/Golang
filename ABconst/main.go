package main

import "fmt"

/**
大纲：
	常量定义了不一定 必须调用，这点与变量区分特性
	常量声明:使用关键子const
	批量声明&iota：使用批量声明时 必须给他赋值一个原始值,若没有写则默认与上一个相同，
				iota是go语言的常量计数器，只能在常量的表达式中使用。
				iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
				使用iota能简化定义，在定义枚举时很有用。
	iota妙用1：快速定义大小（根据iota特性灵活使用 ）
	iota妙用2:多个iota定义在一行
*/
//常量声明
const pi = 3.14

//批量声明&iota
const (
	n1 = 100
	n2
	n3 = iota
	n4
	_
	n6 = iota
)

//iota妙用1
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//iota妙用2
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	//常量声明
	fmt.Println(pi)
	//批量声明&iota：iota可以看作是一个行索引
	fmt.Println(n1, n2, n3, n4, n6)
	//iota妙用1
	fmt.Println(KB, MB, GB, TB, PB)
	//iota妙用2
	fmt.Println(a,f)

}
