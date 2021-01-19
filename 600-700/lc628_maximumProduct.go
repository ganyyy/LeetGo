package main

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	// Top3 ?
	// 注意负数
	// 直接排序是 O(nlog(n))
	// 有没有O(n)级别的算法呢?

	sort.Ints(nums)

	var n = len(nums)
	if nums[0] > 0 {
		return nums[n-1] * nums[n-2] * nums[n-3]
	}
	// 有负数, 看一下怎么算更大

	// 有两个以上的负数, 取开头俩*末尾和末尾三的最大值
	if nums[1] < 0 {
		return max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])
	}

	// 这时候只能选这个了吧...
	return nums[n-1] * nums[n-2] * nums[n-3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumProduct2(nums []int) int {
	var min1, min2 = math.MaxInt32, math.MaxInt32
	var max1, max2, max3 = math.MinInt32, math.MinInt32, math.MinInt32

	// 求最小的两个和最大的三个数
	for _, v := range nums {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}

		if v > max1 {
			max2, max3 = max1, max2
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}

	var a, b = min1 * min2 * max1, max1 * max2 * max3
	if a > b {
		return a
	}
	return b
}
