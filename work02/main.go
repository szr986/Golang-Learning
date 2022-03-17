package main

import (
	"fmt"
	"sort"
)

func findMax(l, r, u, d int) int {
	m := []int{l, r, u, d}
	sort.Ints(m)
	return m[0]
}

func main() {
	fmt.Println()
}
