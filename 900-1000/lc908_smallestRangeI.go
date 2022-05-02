package main

import "math"

func smallestRangeI(nums []int, k int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var mx = math.MinInt32
	var mi = math.MaxInt32

	for _, v := range nums {
		mx = max(v, mx)
		mi = min(v, mi)
	}

	if mx-mi <= 2*k {
		return 0
	}
	return (mx - mi) - 2*k
}
