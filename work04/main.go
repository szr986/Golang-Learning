package main

import "fmt"

func main() {
	result = combine(4, 2, 1)
	fmt.Println(result)

}

var path []int
var result [][]int

func combine(n int, k int, startindex int) [][]int {
	backtracking(n, k, 1)
	return result

}

func backtracking(n int, k int, startindex int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		result = append(result, temp)
		return
	}
	for i := startindex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i+1)
		path = path[:len(path)-1]
	}

}
