package main

import "slices"

func minCost(nums []int, x int) int64 {
	n := len(nums)
	// 最终结果是操作次数*x+购买的开销
	// 成本处于[0, n)之间, 此时的s是仅移动的成本
	s := make([]int64, n) // s[k] 统计操作 k 次的总成本
	for i := range s {
		s[i] = int64(i) * int64(x)
	}
	for i, mn := range nums { // 子数组左端点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形的）
			// 从i移动到j时, 对应的最小开销
			// 实际上, mn的值会在降低到最小值后, 一直保持不变
			mn = min(mn, nums[j%n]) // 维护从 nums[i] 到 nums[j] 的最小值
			// j-i表示的就是移动了多少次
			s[j-i] += int64(mn) // 累加操作 j-i 次的花费
		}
	}
	return slices.Min(s)
}
