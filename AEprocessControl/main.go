package main

import "fmt"

/**
大纲：
	if判断：if语句与Java中的if语句整体相同，但是他增加了一个新的特性，就是在if代码块后面定义一个作用域只属于if语句块的变量
	for循环: for循环是go语言中唯一的一种循环，go语言取消了while do-while循环
			for循环整体与Java一致，但他比Java更强大的是go语言的for循环可以一起且在省略时必须一起省略初始语句与结束语句
			或者同时省略所有语句 构成for循环的死循环
			for循环可以通过break goto return panic语句退出循环
	for range(键值循环)：
			go语言可以使用for range遍历数组，切片，字符串，map及channel(通道),
			遍历他们返回有以下规律
				数组、切片、字符串			索引、数值
				map						键、值
				channel					只返回通道内的值
	switch ：流程控制语句和Java中有些不同
			Java中 当符合case后如果让流程不继续执行必须使用break跳出switch
			go中 当执行完符合条件的case后会自动跳出 switch 省略了break语句
			go中还可以写多个case相等的值
			go中更可以在case上面直接写if判断
			fallthrough 为了兼容C语言 fallthrough会执行下一个case语句

	break：跳出单层循环	支持标签
	goto:跳出所有循环		支持标签
	continue:继续 循环  支持标签

*/
func main() {
	//if判断
	age := 10
	if age > 18 {
		fmt.Println("年龄大于18岁")
	} else if age < 18 {
		fmt.Println("warning....")
	} else {
		fmt.Println("成年了！")
	}

	if age2 := 2; age2 > 20 {
		fmt.Println("age2大于20")
	} else if age2 > 1 {
		fmt.Println("age2大于1")
	}

	//for循环

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	index := 0
	for index < 10 {
		index++
		fmt.Println(index)

	}
	//死循环 类似于Java的while(true)
	//for {
	//fmt.Println("我是死循环")
	//}

	nuns := []int{1, 2, 3, 4}
	for i, v := range nuns {
		fmt.Printf("数组的下标为:%v,数组的值:%v \n", i, v)
	}
	//switch
	switchDemo()
	switchDemo2()
	switchDemo3()
	//break
	useBreak()
	useBreakTag()
	//continue
	useContinue()
	continueDemo()
	//goto
	useBreakBreak()
	useGoto()
}

//switch
func switchDemo() {
	//finger := 3
	finger := 1
	switch finger {
	case 1:
		println("finger等于1")
		break
	case 2:
		println("finger等于2")
	case 3:
		println("finger等于3")

	case 4:
		println("finger等于4")
	default:
		println("无效输入")
	}
}

//switch
func switchDemo2() {
	switch finger := 3; finger {
	case 1, 3, 5, 7:
		println("finger in 1 3 5 7")
		break
	case 2, 4, 6, 8:
		println("finger in 2 4 6 8")
	default:
		println("无效输入")
	}

}

//演示在case中使用if判读，并且使用fallthrough去执行紧挨着的下一个case
func switchDemo3() {
	age := 24
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
		fallthrough
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

//beak
func useBreak() {
	i := 0
	for j := 0; j < 3; j++ {
		for i <= 3 {
			if i == 2 {
				//break会跳出当前循环终止本次循环
				break
			}
			i++

		}
		fmt.Printf("j执行了:%d次\n", j)
	}
	fmt.Printf("最终i的值等于：%d\n", i)
}

func useBreakTag() {
	i := 0
tag:
	for j := 0; j < 3; j++ {
		for i <= 3 {
			if i == 2 {
				//break会跳出当前循环终止本次循环
				break tag
			}
			i++
		}
		fmt.Printf("j执行了:%d次\n", j)
	}
	fmt.Printf("最终i的值等于：%d\n", i)
}

//useContinue
func useContinue() {
	i := 0
	for j := 0; j < 3; j++ {
		for i <= 3 {
			i++
			fmt.Println("我能执行三次")
			if i >= 2 {
				//continue会继续下一次循环，与Java一样
				continue
			}
			i++
			fmt.Println("我最多执行一次")

		}
		fmt.Printf("j执行了:%d次\n", j)
	}
	fmt.Printf("最终i的值等于：%d\n", i)
}

//使用标签可直接继续执行外层循环
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

//在没有goto关键字之前，如果要跳出双层循环需要使用一个标志位，两个break才能跳出双层循环
func useBreakBreak() {
	flag := false
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				flag = true
				break
			}
			fmt.Printf("%d----%d\n", i, j)
		}
		if flag {
			break
		}
	}
	fmt.Println("双层for循环结束")
}
func useGoto() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				goto breakTag
			}
			fmt.Printf("%d----%d\n", i, j)
		}
	}
breakTag: //定义一个标签
	fmt.Println("双层for循环结束")
}
