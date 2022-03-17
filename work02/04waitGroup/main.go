package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func main() {
	f()
}
