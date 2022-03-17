package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端
func processConn(conn net.Conn) {
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read failed err:", err)
		}
		fmt.Println(string(tmp[:n]))

		fmt.Println("请回复：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
	}
}

func main() {
	// 1.本地端口启动服务
	listner, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp failed err:", err)
	}
	// 2.等待别人来跟我建立连接
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Accept failed err:", err)
		}
		// 3.与客户通信
		go processConn(conn)
	}

}
