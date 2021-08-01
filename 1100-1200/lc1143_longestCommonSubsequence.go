package main

func longestCommonSubsequence(text1 string, text2 string) int {

	//dp[i][j]表示的是text1[:i+1]和text2[:j+1]中的最长公共字串的长度
	var dp = make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				// 匹配的情况下, 直接从左上+1
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// 不匹配的情况下, 从上和左取一个最大值
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	// 压缩版

	//1. 选取较小的字符串构建DP数组

	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}
	var dp = make([]int, len(text1)+1)

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 和二维DP概念相同
	// 因为只需要 左, 上, 左上 三个方向, 所以可以进行dp的压缩
	var pre int
	for i := range text2 {
		// pre记录的是二维中左上方的值
		pre = 0
		for j := range text1 {
			// cur相当于记录的是二维中的正上方的值
			var cur = dp[j+1]
			if text2[i] == text1[j] {
				// 相等的情况下, 左上方+1
				dp[j+1] = pre + 1
			} else {
				// 不等的情况下, 上方和左方中取最大值
				dp[j+1] = max(dp[j], cur)
			}
			pre = cur
		}
	}

	return dp[len(text1)]
}
