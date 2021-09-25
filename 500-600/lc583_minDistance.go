package main

func minDistance(word1 string, word2 string) int {
	var l1, l2 = len(word1), len(word2)

	var dp = make([][]int16, l1+1)
	for i := range dp {
		dp[i] = make([]int16, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return int(int16(l1+l2) - dp[l1][l2]*2)
}

func max(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
