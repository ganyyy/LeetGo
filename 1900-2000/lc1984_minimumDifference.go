package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	// 还不行, 因为某个区间内的值可能会差值更小

	// 这题有点意思欸

	// 滑动窗口走一波?

	sort.Ints(nums)

	var ret = math.MaxInt32
	for i := 0; i <= len(nums)-k; i++ {
		var cur = nums[i+k-1] - nums[i]
		if cur < ret {
			ret = cur
		}
	}
	return ret
}
