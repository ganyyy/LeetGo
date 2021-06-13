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

func change2(amount int, coins []int) int {
	// 无限背包

	// dp[i] 表示到 i有几种换法. 0代表只有一种换法, 即什么都不换
	var dp = make([]int, amount+1)

	dp[0] = 1
	for _, c := range coins {
		// 这里执行前序遍历, 是因为当前值需要通过前边的值来确定
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}
	return dp[amount]
}
