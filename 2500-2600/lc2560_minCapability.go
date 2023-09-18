package main

import (
	"math"
	"sort"
)

func minCapability(nums []int, k int) int {
	// 常见的最大化最小值问题(划重点: 常见的~)
	mx, mi := math.MinInt32, math.MaxInt32

	for _, n := range nums {
		if mx < n {
			mx = n
		}
		if mi > n {
			mi = n
		}
	}
	if mi == mx {
		return mi
	}
	// 不适合枚举所有的组合. 只需要先限定好当前迭代的最大金额
	n := len(nums)
	// 区间就是[0,mx-mi]
	return sort.Search(mx-mi, func(m int) bool {
		// 为什么二分出的结果, 一定是属于nums的呢?
		// 核心点在于nums[i] <= m这个限定. 因为如果存在若干个m不在nums[i]中
		// 那么最终一定可以缩进到某个nums[i]上.
		// 因为从[nums[i], m]之间返回的结果都是相同的!
		m += mi
		var cnt int
		for i := 0; i < n; i++ {
			if nums[i] <= m {
				cnt++
				i++
			}
		}
		// 如果在当前限定的mx偷取目标下, 可以偷取的房子数量
		// >= k, 说明mx还有下降空间
		//  < k, 说明mx需要变大才能做到偷取至少k个房子
		return cnt >= k
	}) + mi
}
