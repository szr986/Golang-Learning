package main

import "fmt"

//Go语言中如果标识符是大写的，就表示包对外部包可见

//Person 是一个结构体
type Person struct {
	gender, name string
}

//方法，作用于特定数据类型的函数
func (p Person) fuck() {
	fmt.Printf("fuck %v", p.gender)
}

//构造函数返回的是拷贝，所以尽量返回指针，减少内存消耗
func newPerson(gender, name string) *Person {
	return &Person{gender: gender,
		name: name}
}

// Go中函数参数永远是拷贝
func f1(x Person) {
	x.gender = "female"
}

func f2(x *Person) {
	(*x).gender = "female" //根据内存地址找到原变量
}

func main() {
	var p Person
	p.name = "John"
	p.gender = "male"
	f1(p)
	fmt.Println(p.gender) // male

	// 若要修改原变量，则需要指针
	f2(&p)
	fmt.Println(p.gender) //female

	var p2 = new(Person)
	fmt.Printf("%T\n", p2)
	fmt.Println(p2)

	pp := newPerson("dw", "dwa")
	fmt.Println(pp)
	//接受者表示调用该方法的具体变量
	p.fuck()
}
