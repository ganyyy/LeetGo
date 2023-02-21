package main

func stoneGameII(piles []int) int {
	ln := len(piles)
	var sum int

	// dp[i][j]表示剩余[i,ln-1]堆时，M = j的情况下，先取的人能获得的最多石子数
	// 易得M的上限就是ln
	dp := make([][]int, ln)
	for i := range dp {
		dp[i] = make([]int, ln+1)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 因为前依赖于后, 所以倒叙查找
	for i := ln - 1; i >= 0; i-- {
		sum += piles[i]
		for M := 1; M <= ln; M++ {
			if i+2*M >= ln {
				// 可以一次性全部拿走
				dp[i][M] = sum
			} else {
				// 无法一次性全部拿走, 就尽量让对手拿的少
				for x := 1; x <= 2*M; x++ {
					dp[i][M] = max(
						dp[i][M], sum-dp[i+x][max(M, x)],
					)
				}
			}
		}
	}
	return dp[0][1]
}
