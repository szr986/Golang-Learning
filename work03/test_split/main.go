package main

import (
	"fmt"
	"work/work03/split_string"
)

func main() {
	ret := split_string.Split("bcbdbef", "b")
	fmt.Printf("%#v", ret)
}
