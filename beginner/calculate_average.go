//go:build ignore

package main

import "fmt"

func calc_average(nums []int) float64 {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return float64(sum) / float64(len(nums))
}

func main() {
	nums := []int{10, 20, 30, 40, 50}
	average := calc_average(nums)
	fmt.Printf("The average is: %v\n", average)
}
