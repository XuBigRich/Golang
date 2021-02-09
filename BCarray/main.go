package main

import "fmt"

/**
大纲：
	定义数组：数组必须在定义时就指定好类型与数组长度,当填值小于数组长度时，剩余数据将以0补全
	数组是值类型：一旦定义数组整体不可改变（数组长度、每个数据的长度），但数组上面的数据可以改变。
				且值类型还有一个特点就是当其他变量来接收时，其实是拷贝一份数据赋值给新的变量引用

*/
var (
	array = [5]int{0, 1, 2, 3, 4}
)

//定义数组
func define() {
	//先声明，后赋值
	var array1 [5]int
	var array2 [10]int
	array1 = [5]int{0, 1, 2, 3, 4}
	array2 = [10]int{0, 1, 2, 3, 4}
	fmt.Println(array1)
	fmt.Println(array2)
	//边声明，边赋值
	var array3 [5]int = [5]int{0, 1, 2, 3, 4}
	var array4 [10]int = [10]int{0, 1, 2, 3, 4}
	fmt.Println(array3)
	fmt.Println(array4)
	//边声明，边赋值，通过类型推断 让编译器根据值自己推断数组类型
	var array5 = [5]int{0, 1, 2, 3, 4}
	var array6 = [10]int{0, 1, 2, 3, 4}
	fmt.Println(array5)
	fmt.Println(array6)
	//使用...让编译器自己统计计算数组大小
	var array7 = [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(array7)
	//根据索引值进行初始化数组
	var array8 [20]int
	array8 = [20]int{19: 1}
	fmt.Println(array8)
}

//数组取值
func get() {
	fmt.Println(array[4])
}

//遍历数组
func each() {
	for i := len(array); i >= 0; i-- {
		fmt.Println(i)
	}
	for i := range array {
		fmt.Printf("索引为i:%v,数值为:%v \n", i, array[i])
	}
	for i, v := range array {
		fmt.Printf("索引为i:%v,数值为:%v \n", i, v)
	}
}

//数组是值类型
func price() {
	a := [2]int{1, 2}
	//将a的数据赋值给b，其实就是将a的数据拷贝一份 赋给变量b b是从内存中新申请的地址，与a不是同一个地址块
	b := a
	//改变b[0]
	b[0] = 100
	//打印变量a
	fmt.Println(a)
	//打印变量b   证明了a与b不是同一个地址变量
	fmt.Println(b)
}
func main() {
	define()
	get()
	each()
}
