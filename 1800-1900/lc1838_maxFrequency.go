package main

import "sort"

func maxFrequency(nums []int, k int) int {

	// 滑动窗口?

	sort.Ints(nums)

	var ret int
	var left, right = 0, 1

	for right < len(nums) {
		// 计算 使得 nums[left:right+1] == nums[right]所需要的消耗
		// 假设 nums[left:right+1] == nums[right], 那么要使得 nums[left:right+2] == nums[right+1]\
		// 所需要的开销就是 (nums[right+1]-nums[right])*(right+1-left)
		if sub := (nums[right] - nums[right-1]) * (right - left); k >= sub {
			k -= sub
			right++
		} else {
			ret = max(ret, right-left)
			// 注意这里, 因为不满足将nums[right] = nums[right-1]
			// 所以需要恢复将nums[right-1]变成nums[right-2]所需要的消耗
			k += nums[right-1] - nums[left]
			left++
			// 注意: 此时不能直接跳过right. 因为可能需要多次回缩才能满足 nums[right]==nums[right-1]
		}
	}
	// 末尾再计算一下最大值
	return max(ret, right-left)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
