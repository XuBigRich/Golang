package main

import (
	"fmt"
	"math"
)

/**
接口的显示调用：
	声明一个接口类型，然后声明一个普通结构体（这个结构体要是某接口方法的接收者）
	然后将声明好的结构体赋值给接口（接口可以接收 声明的结构体接收者），接口就可以调用，对应的方法了
接口的隐式调用:

*/

//接口的显示调用
type Abser interface {
	Abs() float64
}
type MyFloat float64
type Vertex struct {
	X, Y float64
}

//这是一个不可以被接口接收的类型
type NoInterfaceType int

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//如果是指针类型只可以由指针类型进行调用,如果是值类型，那么可以被指针或值调用
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//方法隐式调用
type HiddenStruct bool

type hiddenInterface interface {
	hiddenMonth() bool
}

func (h HiddenStruct) hiddenMonth() bool {
	return bool(h)
}

func (f MyFloat) hiddenMonth() bool {
	return false
}

func obvious() {
	//接口方法的显示调用
	var a Abser
	f := MyFloat(-math.Sqrt2)
	fmt.Printf("f的类型%T\n", f)
	v := Vertex{3, 4}
	fmt.Printf("v的类型%T\n", v)
	a = f // a MyFloat 实现了 Abser
	fmt.Printf("a的类型%T\n", a)
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser
	fmt.Println(a.Abs())
	//Abser这个接口没有实现NoInterfaceType类型的方法
	//a = NoInterfaceType
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v
	fmt.Println(a.Abs())
}

func hide() {
	//接口方法的显示/隐式调用
	b := HiddenStruct(true)
	var i hiddenInterface
	i = &b
	fmt.Printf("i的类型%T\n", i)
	fmt.Println(i.hiddenMonth())
	var c hiddenInterface = HiddenStruct(false)
	fmt.Printf("c的类型%T\n", c)
	fmt.Println(c.hiddenMonth())
}

func main() {
	//obvious()
	//hide()

}
