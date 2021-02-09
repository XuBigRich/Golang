package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func define() {
	var m1 map[string]int
	//使用map要先进行初始化，申请一段内存空间，map同样支持自动扩容，但建议提前估算好容量
	m1 = make(map[string]int, 20)
	m1["许鸿志"] = 18
	fmt.Println(m1)
}

//map的取值
func get() {
	var m1 map[string]int
	m1 = make(map[string]int, 20)
	m1["许鸿志"] = 18
	//如果取一个存在的值可以直接打印出来
	fmt.Println(m1["许鸿志"])
	//如果取一个不存在的值会打印出对应类型的0值
	fmt.Println(m1["许大富"])
	//显然 打印出0我们无法判断是否这个值存在 ，所以我们通常取值是使用 下面这个方法
	v, ok := m1["许大富"]
	if ok {
		fmt.Printf("值为%v\n", v)
	} else {
		fmt.Println("该索引中没有值")
	}
}

//map的遍历
func mapRange() {
	var m1 map[string]int
	m1 = make(map[string]int, 20)
	m1["许鸿志"] = 18
	m1["许渣渣"] = 21
	m1["许大富"] = 23
	//只遍历key
	for k := range m1 {
		fmt.Println("遍历map的Key:", k)
	}
	//遍历key与value
	for k, v := range m1 {
		fmt.Println("遍历map的Key:", k, "遍历map的Value:", v)
	}
}

//删除map的某个元素
func deleteMap() {
	var m1 map[string]int
	m1 = make(map[string]int, 20)
	m1["许鸿志"] = 18
	m1["许渣渣"] = 21
	m1["许大富"] = 23
	//只遍历key
	for k := range m1 {
		fmt.Println("遍历map的Key:", k)
	}
	//删除许渣渣
	delete(m1, "许渣渣")
	for k := range m1 {
		fmt.Println("遍历map的Key:", k)
	}
	//尝试删除一个不存在的元素  成功执行 ，并没有报错
	delete(m1, "许")
	for k := range m1 {
		fmt.Println("遍历map的Key:", k)
	}
}

//小技巧 按照指定的顺序遍历所有，
//可以先将所有的key取出，然后通过指定的规则对key进行排序，然后通过排序好的key逐个取出value
//这样愚蠢的方法 来达到 实现 按照指定顺序遍历map的需求
func specifiedSore() {
	//预先准备好要操作的map
	rand.Seed(time.Now().UnixNano()) //初始化伪随机的数种子，如果种子一定，那么每次生成的随机数也一定 可以查看randSeed方法
	var scareMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", rand.Intn(100))
		val := rand.Intn(100)
		scareMap[key] = val
	}
	for k, v := range scareMap {
		fmt.Println(k, v)
	}
	fmt.Println("===排序后===")
	//将map中的key全部取出，放入切片keys中
	keys := make([]string, 0, 200)
	for k := range scareMap {
		keys = append(keys, k)
	}
	//对切片进行排序 这个内置函数好
	sort.Strings(keys)
	//从排序好的切片中挨个取出map的值
	for _, v := range keys {
		fmt.Println(v, scareMap[v])
	}

}

//map组合与slice
func MapGroupSlice() {
	a := make(map[string][]int, 10)
	a["你好"]=make([]int,2,2)
	a["世界"]=[]int{123,321}
	fmt.Println(a)
}

//slice与map组合
func SliceGroupMap() {
	//声明一个类型为map的切片,实际使用大小为0，容量为10
	a := make([]map[string]int, 0, 10)
	//建立一个map
	b := make(map[string]int, 10)
	b["你好"] = 200
	//将map追加到切片当中，使用append是一点一点追加占用切片位置
	a = append(a, b)
	fmt.Println(a)

	//下面代码有两个错误首先，先仔细观察下面代码哪里有错误
	//第一：建立切片容量和大小，容量虽然是10，但大小为0，代表建立的这个切片c的大小一点也没使用 ，这时候使用c[0]就会造成数组越界
	//c:=make([]map[string]int,0,10)
	//第二：在map没有初始化的情况下 就直接对map进行赋值，会抛出nil异常
	//c[0]["你好"]=100
	//改进版  一建立就使用了10个位置的切片位置，后续使用就需要通过替换的方式去使用
	c := make([]map[string]int, 10, 10)
	d := make(map[string]int, 10)
	//使用a[0]是直接替换切片的[0]位置
	c[0] = d
	c[0]["你好"] = 100
	fmt.Println(c)
}

//获取当前时间
func getTime() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Day())
}

/**
所谓假Random，是指所返回的随机数字其实是一个稳定算法所得出的稳定结果序列，而不是真正意义上的随机序列。
Seed就是这个算法开始计算的第一个值。所以就会出现只要seed是一样的，那么后续所有“随机”结果和顺序也都是完全一致的。
通常情况下，你可以用 DateTime.Now.Millisecend() 也就是当前始终的毫秒来做Seed .因为毫秒对你来说是一个1000以内的随即数字。
这样可以大大改善保准库的Random结果的随机性。 不过这仍然算不上是完全随机，因为重复的概率还是千分之一。

另外需要注意的是，如果一直调用标准库Random，那么在调用了N次以后，输出结果就会循环最开始的序列了。
也就是说，标准库Random所能生成的不同结果的个数也是有限的。32位系统一般也就是几万次以后就会出现重复。


*/
func randSeed() {
	//当调用到一个定周期之后，
	//随机数也将会出现机械性重复类似于123123123
	rand.Seed(1)
	for i := 0; i < 5; i++ {
		println(rand.Int())
	}
	//每次生成的随机数都一样
	for i := 0; i < 5; i++ {
		rand.Seed(1)
		println(rand.Int())
	}
}

//小作业 写一个程序统计一个字符串中每个单词出现次数比如"how do you do"中how = 1 、do = 2 、you = 1
func assignment() {
	target := "how do you do"
	operation := strings.Split(target, " ")
	count := make(map[string]int, 20)
	for _, key := range operation {
		value, ok := count[key]
		if(ok){
			value=count[key]+1
			count[key]=value
		}else {
			count[key]=1
		}
	}
	fmt.Println(count)
}
//用Java写
/*public void assignment(){
	String target="how do you do";
	String[] operation=target.split(" ");
	Map count=new HashMap();
	for(String key:operation){
		if(count.containsKey(key)){
			int value=map.get(key)+1;
			count.put(key,value)
		}else{
			count.put(key,1)
        }
    }
	System.out.println(map);
}*/
func main() {
	//define()
	//get()
	//mapRange()
	//deleteMap()
	//specifiedSore()
	//MapGroupSlice()
	//SliceGroupMap()
	//getTime()
	//randSeed()
	assignment()
}
