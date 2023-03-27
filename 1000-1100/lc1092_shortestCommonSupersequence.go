package main

func shortestCommonSupersequence(str1 string, str2 string) string {

	// 正向迭代 LCS
	// dp[i][j] 表示 str1[:i] 和 str2[:j] 所组成的LCS
	n, m := len(str1), len(str2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// 计算二者的最长公共子序列, 因为结果一定是相同的部分+穿插不同的部分
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 此时要么使用str1[i-1], 要么使用str2[j-1]
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 举个例子
	// str1: acd
	// str2: abc
	// 此时构建的二元矩阵就是:
	/*
		    | a | b | c |
		  a | 1 | 1 | 1 |    -> "a"   & "abc"
		  c | 1 | 1 | 2 |    -> "ac"  & "abc"
		  d | 1 | 1 | 2 |    -> "acd" & "abc"

			迭代到[3][3]的时候, d ≠ c, 此时可选
			"ac"&"abc" 	 dp[i-1][j]
			和
			"acd" & "ab" dp[i][j-1]
			两种前缀, 很明显 "ac"&"abc"更合适. 所以, 对于LCS而言, str[:3]和str[:2]提供的贡献是相同的, 即没有影响
			此时相当于舍弃了 str1[3],
			所以, 在逆向迭代还原时, 如果 dp[i][j] == dp[i-1][j], 说明这个位置上需要补一下 str1[i]
	*/

	// 逆向构造 最终结果
	idx := len(str1) + len(str2) - 1 - dp[n][m]
	var buffer = make([]byte, idx+1)

	add := func(b byte) {
		buffer[idx] = b
		idx--
	}

	i, j := n, m
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			// 二者相等, 属于LCS的一部分, 直接加进去
			add(str1[i-1])
			i--
			j--
		} else if dp[i][j] == dp[i-1][j] {
			// LCS 选取了 str2[j], str1[i]作为独立部分加入
			add(str1[i-1])
			i--
		} else {
			// LCS 选取了 str1[i], str2[j]作为独立部分加入
			add(str2[j-1])
			j--
		}
	}

	for i > 0 {
		add(str1[i-1])
		i--
	}
	for j > 0 {
		add(str2[j-1])
		j--
	}

	return string(buffer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
