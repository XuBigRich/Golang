package main

import "fmt"

var (
	multArray = [3][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
	}
)

//定义二维数组
func define() {
	var multArray1 = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(multArray1)
	//多位数组除了第一次可以使用...，其他层都不可以使用
	var multArray2 = [...][3]int{4: {1, 2, 3}}
	fmt.Println(multArray2)
}

//获取数组
func get() {
	fmt.Printf("取出二维数组中的第一个数据:%v", multArray[0][0])
}

//遍历数组
func each() {
	for i := len(multArray)-1; i >= 0; i-- {
		for j := len(multArray[i])-1; j >= 0; j-- {
			fmt.Println(multArray[i][j])
		}
	}
	for i := range multArray {
		for j := range multArray[i] {
			fmt.Printf("数值为:%d \n", multArray[i][j])
		}
	}
	for _, v := range multArray {
		for _, v2 := range v {
			fmt.Printf("数值为:%d \n", v2)
		}
	}
}
func main() {
	//get()
	define()
	each()
}
