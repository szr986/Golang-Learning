package main

import "fmt"

//批量声明
var (
	name string
	age  int
	isok bool
)

//iota 常量计数器
//在const关键字出现时将被重置为0。
//const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
//使用iota能简化定义，在定义枚举时很有用
const (
	//iota类似const的行索引
	a1 = iota //0
	a2 = iota //1
	a3 = iota
)
const a4 = iota

func main() {
	name = "吉吉"
	age = 100
	isok = true
	fmt.Println("hello world")
	//Go中局部变量声明后就必须使用，不然无法编译
	fmt.Printf("name:%s", name) //%占位符
	fmt.Println(isok)
	//%在PrintLn无效？
	fmt.Println("age:%d", age)

	//类型推导(自动判断)
	var s1 = "daw"
	//简短声明,不能在全局声明
	s3 := "hahah"
	fmt.Printf(s1)
	fmt.Println(s3)
	// _  匿名变量，不占内存

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	var p1 float64 = (2.0 + 3.0) / 2
	fmt.Println(p1)
}
