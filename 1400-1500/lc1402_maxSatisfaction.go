package main

import "sort"

func maxSatisfaction(satisfaction []int) int {
	// [left][right] 区间内的最大价值
	// 对于 [mid] 这道菜, 可以选择做或者不做(?)
	// 或则, 先排个序(?)
	// 首先可以确定的是: 尽可能地将 s 值最大的放到最后, s 值较小的选择做或者不做
	sort.Ints(satisfaction)

	var ret, cur, sum int
	for i := len(satisfaction) - 1; i >= 0; i-- {
		cur += satisfaction[i] + sum
		if cur <= 0 || cur < ret {
			break
		}
		ret = cur
		sum += satisfaction[i]
	}

	return ret
}
