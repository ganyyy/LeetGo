package main

import "math"

func coinChange(coins []int, amount int) int {
	// 无限背包, 每一个可以重复选择

	var dp = make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	// fmt.Println(dp)
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
