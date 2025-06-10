//go:build ignore

package main

import (
	"fmt"
)

func minInt(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	minValue := minInt(nums)
	fmt.Println("The minimum value is:", minValue)
}
