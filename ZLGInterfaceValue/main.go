package main

import (
	"fmt"
	"math"
)

/**
接口值:
	接口也是值。它们可以像其它值一样传递。
	接口值可以用作函数的参数或返回值。
	在内部，接口值可以看做包含值和具体类型的元组：
	(value, type)
	接口值保存了一个具体底层类型的具体值。
	接口值调用方法时会执行其底层类型的同名方法。

底层值为 nil 的接口值：
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
注意: 保存了 nil 具体值的接口其自身并不为 nil。
*/

//接口值

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func interfaceValue() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()
	i = F(math.Pi)
	describe(i)
	i.M()
}

//底层值为 nil 的接口值

type NI interface {
	NM()
}

type NT struct {
	S string
}

func (t *NT) NM() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func Ndescribe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceNull() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	Ndescribe(i)
	i.M()
}

func main() {
	//interfaceValue()
	interfaceNull()
}
