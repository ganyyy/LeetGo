package main

import "sort"

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	pre := getLISArray(nums)
	suf := getLISArray(reversed(nums))
	suf = reversed(suf)

	ans := 0
	for i := 0; i < n; i++ {
		// 不能作为端点
		if pre[i] > 1 && suf[i] > 1 {
			ans = max(ans, pre[i]+suf[i]-1)
		}
	}
	return n - ans
}

func getLISArray(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	var seq []int
	for i := 0; i < n; i++ {
		it := sort.SearchInts(seq, nums[i])
		if it == len(seq) {
			seq = append(seq, nums[i])
			dp[i] = len(seq)
		} else {
			seq[it] = nums[i]
			dp[i] = it + 1
		}
	}
	return dp
}

func reversed(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[n-1-i]
	}
	return ans
}
