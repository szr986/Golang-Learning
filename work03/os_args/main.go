package main

import (
	"fmt"
	"os"
)

// 获取命令行参数
func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%T", os.Args)
}
