package main

import "sort"

func findLongestChain(pairs [][]int) int {
	// 又有点像DP?

	// 为啥排第二项呢?
	// (a, b), (c, b)
	// 因为前置要求是a<b, c<d
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	var ret = 1
	var pre = pairs[0]
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] <= pre[1] {
			continue
		}
		// 现在的数对满足 pre[1] < pairs[i][0], 可以组成完美的数对链
		ret++
		pre = pairs[i]
	}
	return ret
}
