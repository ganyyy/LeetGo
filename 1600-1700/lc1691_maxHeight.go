//go:build ignore

package main

import "sort"

func maxHeight(cuboids [][]int) (ret int) {
	// 收益最大的是方块的高
	// 这个要排到最后面
	for _, cuboid := range cuboids {
		sort.Ints(cuboid)
	}
	sort.Slice(cuboids, func(i, j int) bool {
		return sum(cuboids[i]) < sum(cuboids[j])
	})
	dp := make([]int, len(cuboids))

	for i, ic := range cuboids {
		dp[i] = ic[2]
		for j, jc := range cuboids[:i] {
			if jc[0] <= ic[0] && jc[1] <= ic[1] && jc[2] <= ic[2] {
				dp[i] = max(dp[i], dp[j]+ic[2])
			}
		}
		ret = max(ret, dp[i])
	}
	return
}

func sum(nums []int) (ret int) {
	for _, v := range nums {
		ret += v
	}
	return
}
