package main

import (
	"fmt"
)

/**
方法：

你可以为指针接收者声明方法。

这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）

指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
*/

type Vertex struct {
	X, Y float64
}

//值类型接收者改值

func (v Vertex) Abs(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

/**
函数：
改值

*/

func Abs(v Vertex, f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

/**
方法和函数的区别
*/
func function(v *Vertex, f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}
func (v *Vertex) method(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

func main() {
	v := Vertex{1, 1}

	//方法调用
	v.Abs(1)
	fmt.Println(v)
	v.Scale(1)
	fmt.Println(v)

	//函数调用
	Abs(v, 1)
	fmt.Println(v)
	Scale(&v, 1)
	fmt.Println(v)

	//方法和函数在调用方式上的区别
	//方法的调用者可以是指针也可以是值
	v.method(1)
	fmt.Println(v)
	p := &v
	p.method(1)
	fmt.Println(v)
	//函数的调用者只能为指针
	function(p, 1)
	fmt.Println(v)
}
