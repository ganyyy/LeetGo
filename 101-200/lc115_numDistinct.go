package main

func numDistinct(s string, t string) int {
	// DP 啊, 开始吧.
	// 从二维开始, 一步一步压缩到一维

	// 走起

	// dp[i][j]表示 s[:i]中包含t[:j]字串的数量

	// 预留一位表示空串
	var dp = make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	// 状态转移方程
	//              s[i] == t[j], dp[i-1][j-1](全都消耗掉) + dp[i-1][j](从前边找字串)
	// dp[i][j] =
	//              s[i] != t[j], dp[i-1][j](从前边找字串)

	// 初始化, 即 不管s的长度是多少, t为空串 就一定存在一个满足题意的子串
	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if j > i {
				break
			}
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
