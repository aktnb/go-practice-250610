//go:ignore build

package main

import "fmt"

func countChar(s string, char rune) int {
	count := 0
	for _, c := range s {
		if c == char {
			count++
		}
	}
	return count
}

func main() {
	s := "Hello, World!"
	c := 'o'
	fmt.Printf("文字列「%s」に文字「%c」が%d個含まれています\n", s, c, countChar(s, c))
}
