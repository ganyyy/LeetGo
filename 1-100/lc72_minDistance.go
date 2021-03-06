package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// dp方法, 有印象, 但是忘了..
	// dp[i1][i2] 表示 word1[:i1+1] == word2[:i2+1] 所需要的步数
	// 如果需要插入一个字符才相等, dp[i1][i2] = dp[i1-1][i2] + 1
	// 如果需要删除一个字符才相等, dp[i1][i2] = dp[i1][i2-1] + 1
	// 如果需要替换一个字符才相等, dp[i1][i2] = dp[i1-1][i2-1] + 1
	// 如果什么都不用做, 那就dp[i1][i2] = dp[i1-1][i2-1]
	// ***每一步***操作都从以上几步中取最小值

	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	// 初始化第一列, 表示在 word2 == ""时, 需要删除多少
	for i := 1; i <= l1; i++ {
		dp[i][0] = i
	}
	// 初始化第一行, 表示在 word1 == ""时， 需要添加多少
	for j := 1; j <= l2; j++ {
		dp[0][j] = j
	}
	// 最后的判断
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			left, up, leftUp := dp[i][j-1], dp[i-1][j], dp[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				leftUp-- // 相等的话就就一样, 因为下边会+1, 所以这边先减去
			}
			dp[i][j] = min(left, min(up, leftUp)) + 1
		}
	}

	return dp[l1][l2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDistance("ab", "a"))
}
