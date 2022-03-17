package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func addNum(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func getNum2(ch2 chan int, ch1 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(2)
	go addNum(ch1)
	go getNum2(ch2, ch1)
	wg.Wait()
	for i := range ch2 {
		fmt.Println(i)
	}
}
