package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening,err:%v", err)
		return
	}
	// write
	fileObj.Write([]byte("zhoulin\n"))
	// writestring
	fileObj.WriteString("dawdadad")
	fileObj.Close()
}

// bufio
func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening,err:%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("qiqiqiqqi") //写入缓存中
	wr.Flush()                  //将缓存中的写入
}

// ioutil
func writeDemo3() {
	str := "laopop"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("Error opening,err:%v", err)
		return
	}
}

// 写文件
func main() {

	writeDemo3()
}
