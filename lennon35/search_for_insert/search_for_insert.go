package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	var i, n int
	for i, n = range nums {
		if target <= n {
			return i
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 5, 6, 7, 8, 9}

	result := searchInsert(nums, 10)
	fmt.Println(result)
}
