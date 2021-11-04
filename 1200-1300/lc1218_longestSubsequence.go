package main

import "math"

func longestSubsequenceBad(arr []int, difference int) int {
	var mm, mi = math.MinInt32, math.MaxInt32

	for _, v := range arr {
		mm = max(mm, v)
		mi = min(mi, v)
	}

	var sub = min(mi, mi-difference)

	var dp = make([]int, max(mm, mm-difference)-sub+1)
	var ret = 1
	for _, v := range arr {
		dp[v-sub] = dp[v-sub-difference] + 1
		ret = max(ret, dp[v-sub])
	}
	return ret
}

func longestSubsequence(arr []int, difference int) int {
	var dp = make(map[int]int)
	var ret = 1
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		ret = max(ret, dp[v])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
