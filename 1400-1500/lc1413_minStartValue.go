package main

import "math"

func minStartValue(nums []int) int {
	// 最小前缀和?

	var m = math.MaxInt32

	var sum int
	for _, v := range nums {
		sum += v
		m = min(m, sum)
	}

	if m < 0 {
		return -m + 1
	}
	return 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
