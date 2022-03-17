package main

import (
	"fmt"
)

func main() {
	var s string = "wadpd"
	// var m = make(map[string]int)
	// pp := strings.Split(s, "")
	// m[pp[2]] = 2
	// m[pp[3]] = 1
	// fmt.Println(m)
	// fmt.Printf("%T", pp[2])
	for _, v := range s {
		// fmt.Println(k)
		fmt.Println(v - 'a')
	}
}
