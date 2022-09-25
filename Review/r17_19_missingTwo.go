package main

import "math/bits"

func missingTwo(nums []int) []int {
	var diff int
	n := len(nums) + 2
	for _, v := range nums {
		diff ^= v
	}
	for i := 1; i <= n; i++ {
		diff ^= i
	}

	// fmt.Println(diff)

	var a, b int

	// 首个不为0的位置
	dv := 1 << bits.TrailingZeros(uint(diff))

	for i := 1; i <= n; i++ {
		if i&dv != 0 {
			a ^= i
		}
	}
	for _, v := range nums {
		if v&dv != 0 {
			a ^= v
		}
	}
	b = diff ^ a
	return []int{a, b}
}
