package main

import "fmt"

type cat struct {
}

type dog struct {
}

type person struct{}

//接口
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}

func (p person) speak() {
	fmt.Println("aaaaa")
}

func (c cat) speak() {
	fmt.Println("miaomiao")
}

func (d dog) speak() {
	fmt.Println("wangwang")
}

// 不关心一个变量是什么类型，只关心调用什么方法
// 接口是一种特殊的类型，它规定了变量有哪些方法
func da(x speaker) {
	// 接收一个参数，传什么进来就打什么
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)
}
