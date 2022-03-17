package main

import "fmt"

func main() {
	for i := 0; i < 19; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
