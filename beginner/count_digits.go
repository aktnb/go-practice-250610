//go:build ignore

package main

func countDigits(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func main() {
	num := 12345
	digitCount := countDigits(num)
	println("The number of digits in", num, "is:", digitCount)
}
