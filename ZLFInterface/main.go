package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MyFloat float64
type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f<0{
		return float64(-f)
	}
	return float64(f)
}
//如果是指针类型只可以由指针类型进行调用,如果是值类型，那么可以被指针或值调用
func (v *Vertex) Abs()float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}


	a = f  // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser
	fmt.Println(a.Abs())
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}
