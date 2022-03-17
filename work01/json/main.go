package main

import (
	"encoding/json"
	"fmt"
)

// 把GO中结构体-->json字符串
// json字符串-->go能识别的结构体

// 字段大写才能全局可见，才能被其他包使用
type person struct {
	Name string `json:"name"` //在json时用name代替字段名
	Age  int
}

func main() {
	p1 := person{
		Name: "dwadaw",
		Age:  23,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Print("mashal failed err:%v", err)
	}

	fmt.Printf("%#v", string(b))
	// 反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针修改本体，不然修改的是副本
	fmt.Printf("%T", p2)
}
