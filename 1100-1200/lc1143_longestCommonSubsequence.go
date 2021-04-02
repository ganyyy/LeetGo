package main

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}

	var preDP int
	var dp = make([]int, len(text2)+1)

	// 未压缩版
	for i := 1; i <= len(text1); i++ {
		preDP = 0
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = max(dp[j], preDP)
			}
			preDP = dp[j]
		}
	}

	return dp[len(text2)]

}

func longestCommonSubsequence2(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}

	var preDP int
	var dp = make([]int, len(text2)+1)

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
