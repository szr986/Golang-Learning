package main

import "fmt"

//函数
func sum(x int, y int) (ret int) {
	return x + y
}

// defer
func deferdemmo() {
	fmt.Println("12")
	defer fmt.Println("23") //延迟执行
	defer fmt.Println("24") //多个defer倒叙执行,defer类似于压入栈中
	fmt.Println("234")
}

//defer与return的关系，看网页，defer执行在RET=x和return之间

//函数类型作为参数和返回值
func f2(x, y int) int {
	fmt.Println(x + y)
	return x + y
}

func f3(x func(m int, n int) int) int {
	ret := x(1, 2)
	return ret
}

func f4(x, y int) func(int, int) int {
	fmt.Println(23)
	return f2
}

func main() {
	// r := sum(10, 20)
	// fmt.Println(r)
	// deferdemmo()
	fmt.Println(f3(f4(1, 2)))
}
