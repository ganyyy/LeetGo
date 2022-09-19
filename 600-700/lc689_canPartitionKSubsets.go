package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	all := 0
	for _, v := range nums {
		all += v
	}
	if all%k > 0 {
		return false
	}
	per := all / k
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > per {
		return false
	}

	// 状态压缩
	dp := make([]bool, 1<<n)
	var dfs func(int, int) bool
	dfs = func(s, p int) bool {
		// 迭代所有的可能
		if s == 0 {
			// 如果存在一种可能, 此时选取了全部的数字, 那么就可以认为存在满足条件的组合
			return true
		}
		if dp[s] {
			return false
		}
		dp[s] = true
		for i, num := range nums {
			// 整体是有序的
			if num+p > per {
				break
			}
			// 如果这个数字还没选取, 就选一下
			// 使用%是为了保证当 p+nums[i] == pre时, 自动归0
			if s>>i&1 > 0 && dfs(s^1<<i, (p+nums[i])%per) {
				return true
			}
		}
		return false
	}
	return dfs(1<<n-1, 0)
}
