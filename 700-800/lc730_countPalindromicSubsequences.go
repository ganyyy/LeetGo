package main

func countPalindromicSubsequences(s string) (ans int) {
	const mod int = 1e9 + 7
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	// a, aa, a[...]a = 2 + 2*dp[i+1][j-1]

	for sz := 2; sz <= n; sz++ {
		for i, j := 0, sz-1; j < n; i++ {
			if s[i] == s[j] {
				// a..[a_low]..[a_high]..a
				// 针对 [a_low+1][a_high-1]这段区间内的数据进行去重
				//
				low, high := i+1, j-1
				for low <= high && s[low] != s[i] {
					low++
				}
				for high >= low && s[high] != s[j] {
					high--
				}
				if low > high {
					// 没有相似的区间, 直接全加
					dp[i][j] = (2 + dp[i+1][j-1]*2) % mod
				} else if low == high {
					// 存在, 但是只有一个数, 只需要减去一次
					dp[i][j] = (1 + dp[i+1][j-1]*2) % mod
				} else {
					// 存在相似的区间, 减去对应的个数
					dp[i][j] = (dp[i+1][j-1]*2 - dp[low+1][high-1] + mod) % mod
				}
			} else {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + mod) % mod
			}
			j++
		}
	}

	return dp[0][n-1]
}
