package main

import (
	"fmt"
	"strconv"
)

// strconv
// Go语言不能强制类型转换，要转需要用到strconv

func main() {
	str := "1000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	fmt.Println(ret1)
}
