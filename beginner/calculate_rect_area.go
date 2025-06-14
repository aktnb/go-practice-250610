//go:build ignore

package main

import "fmt"

func CalculateRectArea(width float64, height float64) float64 {
	return width * height
}

func main() {
	fmt.Println(CalculateRectArea(10.5, 20.5))
}
