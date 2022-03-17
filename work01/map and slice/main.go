package main

import (
	"fmt"
	"strings"
)

func main() {
	// 元素为map类型的切片
	var slice1 = make([]map[int]string, 5, 10) // make(类型，长度，容量)
	// 要对内部的map做初始化
	slice1[0] = make(map[int]string, 2)
	slice1[0][2] = "wda"
	slice1[0][6] = "dadaw"
	fmt.Println(slice1)

	str := "how do you do think you about"
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)

	countMap := make(map[string]int, 10)
	fmt.Println(countMap)
	for _, key := range strSlice {

		val, isReal := countMap[key]
		fmt.Println(countMap)
		fmt.Println(val)
		fmt.Println(isReal)
		if !isReal {
			countMap[key] = 1
		} else {
			countMap[key] += 1
		}
	}
	fmt.Println(countMap)
}
