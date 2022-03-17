package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

func f(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("323")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	// 如何优雅的通知子goroutine退出
}
