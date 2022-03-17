package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1.与server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed err:", err)
		return
	}
	// 2.发送数据
	var msg string
	// if len(os.Args) < 2 {
	// 	msg = "hello wangye"
	// } else {
	// 	msg = os.Args[1]
	// }
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("请说话：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(text))
	}

	conn.Close()
}
