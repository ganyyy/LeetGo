package main

func longestPalindromeSubseq(s string) int {
	// dp

	// dp[i][j] 表示的是从 s[i:j+1] 之间的最长子序列
	// dp[i][j]:
	// 1) i == j: 1
	// 2) s[i] != s[j]: max(dp[i+1][j], dp[i][j-1])
	// 3) s[i] == s[j]: (j-i > 1 ? dp[i+1][j-1] : 0) + 2

	var dp = make([][]int, len(s))

	for i := range dp {
		dp[i] = make([]int, len(s))

		// 初始化长度, 单个元素的回文长度都是1
		dp[i][i] = 1
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s)-1]
}

func longestPalindromeSubseq2(s string) int {
	// dp[i][j]:
	// 1) i == j: 1
	// 2) s[i] != s[j]: max(dp[i+1][j], dp[i][j-1])
	// 3) s[i] == s[j]: (j-i > 1 ? dp[i+1][j-1] : 0) + 2

	// dp[i][j]表示的是s[i:j+1]中的最大长度(? 个数?)
	n := len(s)
	if n <= 1 {
		return n
	}
	// 压缩i对应的维度
	// 1. 需要注意i== j 的时候应该置为1
	// 2. 这里用到了 dp[i+1][j-1], 所以需要保留一下前态的原始值用来做后续的计算
	// 3. 因为dp[i][j]依赖于dp[i+1][j]和dp[i][j-1], 所以需要逆序遍历i, 正序遍历j
	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		var pre int
		for j := i + 1; j < n; j++ {
			temp := dp[j]
			if j == i {
				dp[j] = 1
			} else if s[i] != s[j] {
				dp[j] = max(dp[j], dp[j-1])
			} else if s[i] == s[j] {
				if j-i == 1 {
					dp[j] = 2
				} else {
					dp[j] = pre + 2
				}

			}
			pre = temp
		}
	}
	return dp[n-1]
}
