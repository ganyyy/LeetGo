package main

import "sort"

func change(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	// 先排个序

	// 这是一个组合问题, 所以需要选项在外边
	sort.Ints(coins)
	dp[0] = 1
	for _, c := range coins {
		if c > amount {
			break
		}
		for j := c; j <= amount; j++ {
			dp[j] += dp[j-c]
		}
	}
	return dp[amount]
}
