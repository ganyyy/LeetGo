package main

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp = make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	// 未压缩版
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]

}

func longestCommonSubsequence2(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}

	var preDP int
	var dp = make([]int, len(text2)+1)

	// dp压缩的新思路:
	/*
		DP的压缩通常要看到具体要使用的数据的来源的依赖关系
		本题中, DP主要依赖于 左上, 上, 左 三个方向
		压缩成一维意味着取消 上 这个维度, 所以需要每次都保留当前状态的初始值留给下一状态使用
	*/

	// 压缩版
	for i := 1; i <= len(text1); i++ {
		preDP = 0
		for j := 1; j <= len(text2); j++ {
			var tmp = dp[j] // 理解为正上方
			if text1[i-1] == text2[j-1] {
				dp[j] = preDP + 1 // 理解为斜上方
			} else {
				dp[j] = max(dp[j-1], tmp)
			}
			preDP = tmp
		}
	}

	return dp[len(text2)]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
