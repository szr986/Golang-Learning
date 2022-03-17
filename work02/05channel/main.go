package main

import (
	"fmt"
	"sync"
)

var b chan int
var wg sync.WaitGroup

// 无缓冲区的channel
func noBuffChannel() {
	b = make(chan int) //通道初始化，无缓冲区
	// 通道的操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("取到了", x)
	}()
	b <- 10 //无缓冲区必须要有接受者才能放进去，不然会卡主
	fmt.Println("10进入了")
	wg.Wait()

}

func buffChannel() {
	b = make(chan int, 16) //通道初始化，16为缓冲区大小
	b <- 100               //有缓冲区就能直接放入，不会卡主
	fmt.Println("执行")
	x := <-b
	fmt.Println("取到了", x)
}

func main() {
	buffChannel()
}
