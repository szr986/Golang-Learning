package main

import "fmt"

func f1() {
	fmt.Println("2323")
}

func f2() {
	f1()
}

func main() {
	f2()
}
