package main

import (
	"flag"
	"fmt"
)

// flag 获取命令行参数
func main() {
	name := flag.String("name", "wangzhi", "请输入名字")
	flag.Parse()
	fmt.Println(*name)
}
