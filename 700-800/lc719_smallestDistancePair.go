package main

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	// 值二分法, 如何快速找到满足条件的对数呢?

	var left, right = 0, nums[len(nums)-1] - nums[0]

	for left < right {
		var mid = (right + left) / 2
		var cnt int

		// 双指针计算区间个数
		// 计算的是距离小于mid的个数
		for i, j := 0, 1; j < len(nums); j++ {
			for nums[j]-nums[i] > mid {
				i++
			}
			cnt += j - i
		}

		if cnt >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
