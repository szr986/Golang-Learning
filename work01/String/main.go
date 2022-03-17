package main

import (
	"fmt"
	"strings"
)

func main() {
	//多行字符串,路径也可直接用不用转义/
	s1 := `商务第三方
		大青蛙多
	`
	fmt.Println(s1)
	//字符串拼接
	name := "卢本伟"
	world := "NB"
	s2 := name + world
	fmt.Println(s2)
	s3 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(s3)
	//字符串分割
	s4 := `F:\Download\Villain0.2.8`
	ret := strings.Split(s4, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(s1, "NB"))
	//前后缀
	fmt.Println(strings.HasPrefix(s1, "商务"))
	fmt.Println(strings.HasSuffix(s1, "商务"))
	//判断子串
	s5 := `abcdefgc`
	fmt.Println(strings.Index(s5, "c"))
	fmt.Println(strings.LastIndex(s5, "def"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
	//字符串修改
	a1 := "白萝卜"
	a2 := []rune(a1)
	a2[0] = '红' //改成切片之后 里面是字符，不是字符串
	fmt.Println(string(a2))

	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//for range
	s := "dawdaw你好"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
