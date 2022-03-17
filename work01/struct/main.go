package main

import "fmt"

//数组
//存放元素的容器，必须指定元素的类型和容量
// 数组的长度是数组类型的一部分
func main() {
	a := [...]int{4, 6, 7, 9, 3, 5, 1}
	fmt.Println(len(a))
	//基于数组切割，左闭右开
	b := a[1:4]
	c := a[2:4]
	fmt.Println(len(b))
	fmt.Println(len(c))
	//切片的容量是指切片第一个元素到底层数组最后一个元素的容量！！
	fmt.Println(cap(b))
	fmt.Println(cap(c))
	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
	//一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	// 但是我们不能说一个长度和容量都是0的切片一定是nil

	//删除切片元素
	a1 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素，2左边切片，2右边切片，再拼接到一起
	a1 = append(a[:2], a[3:]...)
	fmt.Println(a1) //[30 31 33 34 35 36 37]

	//指针
	// &取地址，*根据地址取值

	//map，hash表。
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["利息"] = 18
	m1["wdq"] = 35

	fmt.Println(m1)
	// 判断是否存在该key,约定成俗用ok接受bool值
	v, ok := m1["naza"]
	fmt.Println(v)
	fmt.Println(ok)
	// 遍历map
	for key, value := range m1 {
		fmt.Println(key)
		fmt.Println(value)
	}

	for _, value := range m1 {
		fmt.Println(value)
	}
	// 删除
	// delete(m1,"wdq")

}
