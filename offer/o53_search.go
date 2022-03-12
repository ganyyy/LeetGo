package main

import "sort"

func search(nums []int, target int) int {
	// 左边界和右边界
	// 用search, true 向左, false 向右

	var left = sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	var right = sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})
	return right - left
}
