//go:build ignore

package main

import (
	"fmt"
)

func mergeArrays(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	merged := mergeArrays(arr1, arr2)
	merged[2] = 100
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	fmt.Println("merged array: ", merged)
}
