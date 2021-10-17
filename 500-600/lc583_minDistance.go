package main

func minDistance(word1 string, word2 string) int {
	var l1, l2 = len(word1), len(word2)

	// dp[i][j] 表示 word1[i-1]和word2[j-1]中, 删除任意字符串使其相等的最大长度
	// 即转变为求两个字符串的最大公共子序列的问题
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

	// 只要保证最终结果保留的是最长的, 那么用总的长度减去这个值就是删除字符最少的结果
	return int(int16(l1+l2) - dp[l1][l2]*2)
}

func max(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
