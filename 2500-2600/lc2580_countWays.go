package main

import "slices"

func countWays(ranges [][]int) int {
	if len(ranges) <= 1 {
		// 0 -> 0
		// 1 -> 2
		return len(ranges) * 2
	}
	// 按照起始区间的大小合并,
	// 小的在前, 大的在后
	slices.SortFunc(ranges, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	const MOD = 1e9 + 7

	var curEnd = ranges[0][1]
	var ret = 2
	for _, r := range ranges[1:] {
		if curEnd >= r[0] {
			curEnd = max(curEnd, r[1])
		} else {
			ret = ret * 2 % MOD
			curEnd = r[1]
		}
	}

	// 最后一位数
	// 貌似只需要看几组就行..?

	return ret
}
