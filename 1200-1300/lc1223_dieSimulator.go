package main

func dieSimulator(n int, rollMax []int) int {
	mod := int64(1_000_000_007)
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, 7)
	}

	for j := 1; j <= 6; j++ {
		dp[1][j] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= 6; j++ {
			ans := int64(0)
			for k := range dp[i-1] {
				ans += dp[i-1][k]
				ans %= mod
			}
			idx := i - 1 - rollMax[j-1]
			if idx >= 1 {
				cut := int64(0)
				for k := range dp[idx] {
					cut += dp[idx][k]
					cut %= mod
				}
				cut -= dp[idx][j]
				ans -= cut
			} else if idx == 0 {
				ans -= 1
			}
			// 取余后的结果可能为0, 所以上方的-需要重新计算
			if ans < 0 { // 取模有可能出现负数，加回来就得了，经验
				ans += mod
			}
			dp[i][j] = ans % mod
		}
	}

	res := int64(0)
	for k := range dp[n] {
		res += dp[n][k]
	}

	return int(res % mod)
}
