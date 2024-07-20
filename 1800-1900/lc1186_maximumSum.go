package main

import "math"

func maximumSum(arr []int) int {
	ans := math.MinInt / 2

	// f0: 不删除值
	// f1: 必须要删除一个值(可以是自身, 也可以是之前的值)
	f0, f1 := ans, ans
	for _, x := range arr {
		// 删除值,
		// 如果删除的是x, 那么最终剩下的就是之前的f0
		// 如果删除的不是x, 那么剩下的就是之前删除过一个数的f1+x
		f1 = max(f1+x, f0)
		// 不删除x,
		// 如果不选前边的数, 那就是x
		// 如果选择了前边的数, 那就是f0+x
		f0 = max(f0, 0) + x
		ans = max(ans, f0, f1)
	}
	return ans
}
