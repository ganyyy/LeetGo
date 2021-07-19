package main

import "sort"

func minPairSum(nums []int) int {
	// 先排序, 然后从头到尾两两组合

	sort.Ints(nums)

	var l, r = 0, len(nums) - 1

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var ret int
	for l < r {
		ret = max(ret, nums[l]+nums[r])
		l++
		r--
	}
	return ret
}
