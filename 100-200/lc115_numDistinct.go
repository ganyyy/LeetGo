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

func numDistinctSingleDP(s string, t string) int {
	// len(s) >= len(t)

	var dp = make([]int, len(t)+1)
	// 空串自带一个字串
	dp[0] = 1
	// 由二维降低到一维, 需要考虑的是遍历的次数和需要保存的数据分别是什么
	// 这道题目中, 需要遍历len(s)次,
	// 遍历到第i次的dp[j]表示 t[:j+1]在s[:i+1]中子序列的个数(j <= i)
	// 状态转移方程不变.
	var lt = len(t)
	for i := 1; i <= len(s); i++ {
		// 这里是个重点. dp压缩中, 什么时候需要倒序? 什么时候不需要倒序?
		// 具体通过DP转移方程来进行判断
		// 如果转移方程中, 数据全部依赖于前态, 那么就需要进行倒序处理
		// 如果部分依赖于前态, 部分依赖于当前态, 那么就可以正序处理
		for j := min(i, lt); j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[len(t)]
}
