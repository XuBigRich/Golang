package main

import "fmt"

//结构体定义
type Vertex struct {
	x int
	y int
}

//打印声明的结构体
func print() {
	fmt.Println(Vertex{1, 2})
}

//结构体声明并打印出值
func useStruct() {
	v := Vertex{1, 2}
	v.y = 4
	fmt.Println(v.y)
}

//结构体指针
func pointStruct() {
	v := Vertex{1, 2}
	p := &v
	//指针的显式调用
	(*p).x = 1e2
	fmt.Println(v)
	//指针的隐式调用
	p.x = 1e9
	fmt.Println(v)
}

//结构体声明
func statementStruct() {
	v1 := Vertex{1, 2} // 创建一个 Vertex 类型的结构体
	v2 := Vertex{x: 1}
	v3 := Vertex{}
	p := &Vertex{1, 2}
	fmt.Println(v1, v2, v3, p)
}
func main() {
	print()
	useStruct()
	pointStruct()
	statementStruct()
}
