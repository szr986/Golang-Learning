package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("failed,err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	var tmp = make([]byte, 128)
	for {
		n, err := fileObj.Read(tmp)
		if err != nil {
			fmt.Printf("read failed,err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// 使用bufio
func readFromFile2() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("failed,err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	// 创建一个用来读文件的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		if err == io.EOF {
			return
		}

		fmt.Print(line)
	}
}

func readfromIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("failed,err:%v", err)
		return
	}
	fmt.Println(string(ret))
}

// 打开文件
func main() {
	// readFromFile2()
	readfromIoutil()
}
