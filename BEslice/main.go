package main

import (
	"fmt"
)

/**
大纲：
	切片是拥有相同元素的可变长度的序列，他是基于数组又封装了一层，他非常灵活，支持自动扩容。
	切片是一个引用类型，它内部包含地址、大小、容量。切片一般用于快速的操作一块数据集合。
	记住切片三要素：
		地址：切片第一个元素指向的内存空间地址  ，
			也是这个切片引用数组的内存地址。
		大小：切片中目前元素的个数		len()
			这个大小是指切片所在的引用数组的内存地址使用量
		容量：底层数组最大能存放的元素个数（不够时将扩容） cap（）
			这个容量是指，切片使用了所在数组剩余的全部容量，若超过了这个容量，那么切片就会重新生成一个数组。
=========================
	这里单独讲一下值类型和引用类型 在指针方面的区别：
			查看区别
=========================
	引用类型：
			引用类型指向的其实是一个内存地址
	值类型：
			值类型指向的是一个真正的数据
	切片定义：
			切片定义与数组基本一致 这里就不过多演示了,这里演示切片增加的定义功能
			直接声明切片、通过数组得到切片、通过切片得到切片
	切片取值：
			切片取值与数组也基本一致
	切片的容量与大小：
			分为从数组中得到的切片，与直接声明的切片
	切片的扩容，与append内置函数：
			切片是基于数组的，当切片的个数大于了切片的容量时，切片就会扩容，扩容时，切片就会重新指向一片新的内存,
			扩容的原则是每次都是上次的两倍数
	切片指针：
			切片的指针注意区分他们的指向指针与值指针（取名不准确）
	copy函数:深拷贝
	从切片中删除元素： 切片的删除是通过拼接切片实现的，并展示了如何拆开切片，获取到切片中每个元素
	关于扩容:
		从上面的代码可以看出以下内容：
可以通	过查看$GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：
		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
		否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
		否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
		需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。

*/
var (
	slice = []int{1, 2, 3, 4}
)

//切片定义
func define() {
	//直接声明切片
	var a []int = []int{123, 321, 231}
	var b [3]int = [3]int{123, 321, 231}
	fmt.Printf("a:%T   ,   b:%T \n", a, b)
	fmt.Println(a)
	//使用赋值的方式建立切片
	c := []int{2: 1}
	fmt.Println(c)
	//从数组建立切片	数组的全部建立切片
	d := b[:]
	fmt.Printf("输出数据类型: %T ,切片值: %v \n", d, d)
	//以数组下标1->2(不包含)建立切片
	e := b[1:2]
	fmt.Printf("输出数据类型: %T ,切片值: %v \n", e, e)
	//以数组下标0->1(不包含)建立切片
	f := b[0:1]
	fmt.Printf("输出数据类型: %T ,切片值: %v \n", f, f)
	//从数组开始处取到下标为1(不包含) 建立切片
	g := b[:1]
	fmt.Printf("输出数据类型: %T ,切片值: %v \n", g, g)
	//从数组下标为1处开始 取到数组结尾 建立切片
	h := b[1:]
	fmt.Printf("输出数据类型: %T ,切片值: %v \n", h, h)
	//通过切片得到切片
	i := a[:]
	fmt.Printf("输出数据类型: %T ,切片值: %v \n", i, i)
}

//切片取值
func get() {
	fmt.Println(slice[1])
}

//计算当前切片的大小,与容量  分为从数组中得到的切片，与直接声明的切片
func length() {
	//分为从数组中得到的切片
	city := [...]string{"北京", "上海", "深圳", "杭州", "广州", "济南"}
	sliceCity := city[2:4]
	//切片的大小(切片从数组切到的大小)
	fmt.Println(len(sliceCity))
	//容量（底层数组最大能放多少元素）
	fmt.Println(cap(sliceCity))
	//与直接声明的切片
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

//切片的扩容，与append
func dilatation() {
	var dila = []int{} //空切片 数组空间为0
	fmt.Printf("dila:%v ,len:%d ,cap:%d ,ptr:%p \n", dila, len(dila), cap(dila), dila)
	//追加
	dila = append(dila, 1) //第一次扩容 数组空间由0 ->1 数组重新内存选址
	fmt.Printf("dila:%v ,len:%d ,cap:%d ,ptr:%p \n", dila, len(dila), cap(dila), dila)
	//追加
	dila = append(dila, 1) //第二次扩容  数组空间由1 ->2 数组重新内存选址
	fmt.Printf("dila:%v ,len:%d ,cap:%d ,ptr:%p \n", dila, len(dila), cap(dila), dila)
	//追加
	dila = append(dila, 1) //第三次扩容	数组空间由2 ->4 数组重新内存选址
	fmt.Printf("dila:%v ,len:%d ,cap:%d ,ptr:%p \n", dila, len(dila), cap(dila), dila)
	//追加
	dila = append(dila, 1)
	fmt.Printf("dila:%v ,len:%d ,cap:%d ,ptr:%p \n", dila, len(dila), cap(dila), dila)
	//追加
	dila = append(dila, 1) //第四次扩容	数组空间由4 ->8 数组重新内存选址
	fmt.Printf("dila:%v ,len:%d ,cap:%d ,ptr:%p \n", dila, len(dila), cap(dila), dila)
}

//切片指针
func point() {
	//切片为空
	var dila = []int{}
	println(&dila)
	fmt.Printf("%p\n", dila)
	fmt.Println("========================")

	//	第一次扩容   切片进行扩容重新选址
	dila = append(dila, 1)
	println(&dila)
	//切片的内存地址是切片元素的第一个元素所在内存地址的起始位置
	println(&dila[0])
	fmt.Printf("%p\n", dila)
	fmt.Println("========================")

	//第二次扩容
	dila = append(dila, 1)
	println(&dila)
	//切片的内存地址是切片元素的第一个元素所在内存地址的起始位置
	println(&dila[0])
	fmt.Printf("%p\n", dila)
	//有个疑问就是两个&dila 的地址为什么一样，我目前猜测因为他们取的是切片索引的地址

	//但事实上可以确定的是，切片的内存地址是切片元素的第一个元素所在内存地址的起始位置

	fmt.Println("============这一系列操作讲述了切片的功能原理============")
	//探寻根据数组建立的切片内存地址是否不同，然后更改一个切片地址之后再打印
	fmt.Println("先依照现有的数组建立一个切片，查看他们的内存地址")
	array := [...]int{1, 2}
	sliceByArray := array[:1]
	fmt.Println(&array[0], &sliceByArray[0])
	fmt.Println(array, sliceByArray)
	fmt.Println("发现根据数组建立的切片他们的内存地址相同，也就是说切片使用底层数组是array")
	fmt.Println("========================")
	fmt.Println("更改切片的下标为0的数据，看数组值是否会发生变化")
	sliceByArray[0] = 0
	fmt.Println(&array[0], &sliceByArray[0])
	fmt.Println(array, sliceByArray)
	fmt.Println("更改后发现数组的值发生了变化，说明切片是引用类型，引用了数组的地址")
	fmt.Println("========================")
	//主动让他扩容一次
	fmt.Println("主动让他扩容一次，但这次扩容没有超过数组大小，查看数组是否会因为 切片的改变，而数组产生改变")
	sliceByArray = append(sliceByArray, 3)
	sliceByArray[0] = 0
	fmt.Println(&array[0], &sliceByArray[0])
	fmt.Println(array, sliceByArray)
	fmt.Println("数组产生了改变，说明切片增加了数据是再原来数组上增加的，造成了数组原来数字的丢失")
	fmt.Println("========================")
	//主动让他扩容一次
	fmt.Println("再让切片扩容一次，这次扩容超过了数组大小，查看数组与切片的地址变化")
	sliceByArray = append(sliceByArray, 3)
	sliceByArray[0] = 0
	fmt.Println(&array[0], &sliceByArray[0])
	fmt.Println(array, sliceByArray)
	fmt.Println("此时此刻可以看到，原来的数组已经容不下切片了，于是切片又重新生成了一个数组，将指针指向了新的数组")

}

//深拷贝
func copySlice() {
	a := []int{1, 2, 3}
	b := a //直接赋值,因为切片是引用变量，所以赋值给b是 直接将a指向的内存地址赋值给b

	//先声明，然后初始化
	var c []int //声明一个c变量 还没有此时还没有申请内存地址
	fmt.Println(c)
	//给引用类型初始化（申请内存空间）
	c = make([]int, 2, 3) //申请一个新的内存地址 地址属性为， int切片类型，长度为3，大小为3的

	//声明并且初始化
	d := []int{}
	fmt.Println(c)
	//将a指向的内存地址的数据 拷贝一份给c
	copy(c, a)
	copy(d, a)
	//更改b[0] 查看a师否会跟着改变
	b[0] = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

//从切片中删除元素
func del() {
	//删除切片上的深圳
	city := []string{"北京", "上海", "深圳", "杭州", "广州", "济南"}
	city = append(city[:2], city[3:]...)
	fmt.Println(city)
}

//这里单独讲一下值类型和引用类型 在指针方面的区别
func pointReferenced() {
	//声明一个数组
	array := [...]int{1, 2, 3, 4}
	//声明一个切片
	slice1 := array[:]
	fmt.Printf("数组的地址:%p \n", &array)   //去寻找array这个变量对应值的内存地址
	fmt.Printf("数组的地址上的值:%v \n", *&array)
	fmt.Printf("切片的地址:%p \n", slice1)   //取slice1指向的内存地址
	fmt.Printf("切片的地址:%p \n", *&slice1) //取出指向slice1的指针的值

	fmt.Println(`
	由上面输出可知，值类型必须通过变量取指，才能找到变量指向的值，
	而引用类型，只需要输出变量就可以了，因为引用类型他们本身就是指向的内存地址。
	`)
	/**
	array 的内存地址是A   [...]int{1,2,3,4} 的内存地址是B   =array[0]的地址
	&array是取B
	slice1 本身就是存的内存地址   ==array[0]  等于B
	&slice1 是指向slice1的指针地址
	*&slice1 是从slice1的指针地址

	*&slice1 = slice1 = &array = array[0]的内存地址
	这更加印证了 切片存储的是一个内存地址

	&slice1 其实就类似于 &&array[0] ,可以说没有意义

	*/

}

//关于扩容
func dilata() {
	//请自己查看$GOROOT/src/runtime/slice.go源码
	//需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
}

func main() {
	//define()
	//get()
	//length()
	//dilatation()
	//point()
	//copySlice()
	//del()
	pointReferenced()
	//dilata()

}
