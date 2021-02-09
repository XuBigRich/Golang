package main

import (
	"fmt"
	"math"
)

/**
基本数据类型大纲：
	整型：
		有符号 int int8 int16 int32 int64    节省内存，查找速度也会快
		无符号 uint uint8 uint16 uint32 uint64

		 int与uint是特殊的，他会根据你操作系统的位数自动判断选择，如果你是32位操作系统那么int=int32 如果你是64位操作系统那么int=int64
	浮点：
		浮点数永远是不精确的，
	复数：
		complex64 与 complex128
	布尔：
		bool
	byte和rune:
		英文字符用byte(ASCII)码表示
		rune(UTF-8)
	十进制：什么也不写
	八进制：数字前面添加0
	十六进制：数字前面添加Ox
	正常以十进制输出：默认输出为十进制（人类所默认的）
	按照指定进制输出：通过%b（二进制） 与%o（八进制） 与%x（十六进制） 指定按照相应的进制输出
	变量内存地址:%p 首先使用&取指符 取出a的内存地址，然后通过 %p输出替换
	float：分别有float32与float64两种类型 可以使用math包输出float32与float64的精确大小
	布尔型：	布尔类型变量的默认值为false。
			Go 语言中不允许将整型强制转换为布尔型.
			布尔型无法参与数值运算，也无法与其他类型进行转换。
	字符串：
		一个英文字符长度为1，一个中文字符长度为3
	字符:
		在Java中为char类型，在Golang中被分为了两种 uint8 与rune
		他们的区别是:
			uint8:	uint8代表了一个ASCII码字符，或叫byte型
			rune:	代表utf-8字符串
		当需要处理中文，日文或者其他复合字符时，则需要用到rune类型。rune类型类型实际是一个int32。
		Go使用了特殊的rune类型来处理Unicode，让基于Unicode的文本处理更方便，也可以使用byte型进行默认字符串处理
	字符串转字符数组：



*/

func main() {
	//十进制
	var a int = 10

	//八进制：八进制77 等于十进制的63
	var b int = 077

	//十六进制：十六进制的FF 等于十进制的255
	var c int = 0xFF

	//正常以十进制输出
	fmt.Println(a, b, c)

	//按照指定进制输出
	fmt.Printf("这是一个二进制数字%b \r\n", a)
	fmt.Printf("这是一个八进制数字%o \r\n", b)
	fmt.Printf("这是一个八进制数字%x \r\n", c)

	//变量内存地址：&a为取指符
	fmt.Printf("这是a变量的内存地址%p \r\n", &a)

	//float：输出float32与float64的精确大小
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	var d float32 = 3.2 //这就是Java里面的float类型
	var e float64 = 2.2 //默认是float64 也就是Java里面的Double类型
	fmt.Println(d, e)

	//布尔型
	var f bool
	fmt.Println(f)

	//字符串
	//一个汉字的长度为3，一个英文字符的长度为1
	fmt.Println(len("争"))
	//一个英文字母的长度
	fmt.Println(len("a"))
	//中文标点符号长度与英文标点符号长度
	fmt.Println(len("，,"))
	fmt.Println("这是一个字符串，它主要用于测试\"转译符\", \r\n\\并且我发现这里的转译符和Java差不多,就是换行有点不同。\\")
	var s1 = "单行文本"
	//照格式输出
	var s2 = `
	这是 
	一个多行文本
	他会原样输出" "  \
	类似于groovy字符串
	`
	fmt.Println(s1)
	fmt.Println(s2)

	//字符
	//按照数组位去遍历字符串 当使用这种方法遍历字符时，遇到中文会碰到乱码，因为他们每次只读一个位8个bit 255 (他是按照ASCII码去打印输出的)
	char := "i am 许鸿志"
	for i := 0; i < len(char); i++ {
		//%c代表输出单个字符
		fmt.Printf("该字符在go语言中存储的方式是:%v ,该值对应的unicode码值:%c \n", char[i], char[i])
	}
	fmt.Println()
	//按照rune类型遍历字符串（或许是ASCII码或许是UTF-8 由range自动识别） range会自动识别一个字符 并返回索引与字符 我们使用 i 接收索引，r接收字符
	//range循环可以循环字符串也可以循环数组 ,可以发现i其实依然是从5 8 11 这个位置开始汉字长度计算的
	for i, r := range char {
		fmt.Printf("索引位置:%d,该字符在go语言中存储的方式是:%v ,该值对应的unicode码值:%c ;\n", i, r, r)
	}
	fmt.Println()

	//字符串强转字符数组
	target := "我是将要被替换的数组" //10个字符

	chars := []rune(target) //强转[]rune类型数组 chars的数组大小为10
	chars[2] = '已'
	fmt.Println(string(chars))

	bytes := []byte(target) //强转[]byte类型数组后 bytes的数组大小为30
	fmt.Println(bytes)

	//字符串反转

	//待改进（错误）
	//将字符串转为字符数组，然后计算这个字符数组的长度，然后声明一个新的数组，这个数组大小暂定与字符串长度一致，
	//然后最后用过for循环从数组最后一个字符开始取出，赋值新生成的数组从0开始一直到最后。
	//最终将字节数组强转为字符串，最后发现这个数组输出的竟然是一堆数字，侧面的表明，rune其实是int32
	/*reverse := "要反转这个字符串"
	operation := []rune(reverse)
	length := len(operation)
	var result []rune
	index := 0
	for i := length - 1; i >= 0; i-- {
		result[index] = operation[i]
		index++
	}
	fmt.Println(string(result))*/

	//将字符串转为字符数组，然后计算这个字符数组的长度，最后用过for循环从数组最后一个字符开始取出，通过强转为string类型追加给result字符串
	reverse := "要反转这个字符串"
	operation := []rune(reverse)
	length := len(operation)
	result := ""
	for i := length - 1; i >= 0; i-- {
		result += string(operation[i])
	}
	fmt.Println(result)

	//方法二
	resultChars := []rune(result)
	for i := 0; i < length/2; i++ {
		//高级语言互换位置的语法
		resultChars[i], resultChars[length-1-i] = resultChars[length-1-i], resultChars[i]
	}
	fmt.Println(string(resultChars))
}
