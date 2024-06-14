package main

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	k *= 2

	// n + 2*k
	var ret int
	var left int
	for i, rv := range nums {
		rightLimit := rv - k
		for ; nums[left] < rightLimit; left++ {
		}
		ret = max(ret, i-left+1)
	}
	return ret
}
