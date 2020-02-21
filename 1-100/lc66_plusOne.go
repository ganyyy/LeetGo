package main

import "fmt"

func plusOne(digits []int) []int {
	add := 1
	l := len(digits)

	for i := l - 1; i >= 0; i-- {
		digits[i] += add
		if digits[i] < 10 {
			return digits
		} else {
			add = 1
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}
