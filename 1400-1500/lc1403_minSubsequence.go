package main

import "sort"

func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var sum int
	for _, v := range nums {
		sum += v
	}
	// 有个坑, 如何判断过半?
	// 2*v > sum √
	// v > sum/2 ×
	var cur int
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if cur*2 > sum {
			return nums[:i+1]
		}
	}
	return nums
}
