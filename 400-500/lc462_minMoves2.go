package main

import "sort"

func minMoves2(nums []int) int {
	// 1, 10, 2, 9
	// 8+7+1
	// 22

	// 中位数最优解
	sort.Ints(nums)

	var ret int
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		ret += nums[j] - nums[i]
	}
	return ret
}
