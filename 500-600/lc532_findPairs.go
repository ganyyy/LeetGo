package main

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	var cnt int

	for left, right := 0, 1; right < len(nums); right++ {
		// 去掉相同的数字
		if right+1 < len(nums) && nums[right] == nums[right+1] {
			continue
		}
		var sub = nums[right] - nums[left]
		if sub < k {
			continue
		}
		for sub > k {
			left++
			sub = nums[right] - nums[left]
		}
		if sub != k || left == right {
			continue
		}
		cnt++
		// 去掉相同的数字
		for left < right && nums[left] == nums[left+1] {
			left++
		}
	}

	return cnt
}
