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

func minDistance2(word1 string, word2 string) int {
	// 经典DP - 二位压缩版本

	var m, n = len(word1), len(word2)
	var cost = make([]int, n+1)
	// cost[i][j], word1[:i] -> word2[:j] 所需要的最小的消耗、
	// cost[0][X]和cost[X][0]的开销, 就是两个字符串对应的长度: 只能删除

	for i := 1; i <= n; i++ {
		// 初始化第一行, 代表着在word1 == ""时, 需要添加多少
		cost[i] = i
	}

	for i := 1; i <= m; i++ {
		cost[0] = i // 每次循环开始, cost[0]都是i, 这代表着word1[:i] -> word2[:0]需要删除多少
		// pre代表左上角的值, cur代表上, cost[j-1]代表左
		var pre = i - 1
		for j := 1; j <= n; j++ {
			cur := cost[j]
			if word1[i-1] == word2[j-1] {
				// 相等直接继承
				cost[j] = pre
			} else {
				// 删除, 增加, 替换的三者成本的最小值
				// 感觉可以压缩啊..?
				// 上, 左, 左上
				cost[j] = min(min(cur, cost[j-1]), pre) + 1
			}
			pre = cur
		}
	}
	return cost[n]
}

func minDistance3(word1 string, word2 string) int {
	/*
	   cost[i+1][j+1] = word1[:i+1] -> word2[:j+1]所需要的最小步骤
	   针对 word1[i]和word2[j]而言
	   相等:
	       不需要任何步骤
	   不相等: 需要从以下步骤中找出最小的开销+1
	       如果是插入的话, 相当于是由word1[:i] -> word2[:j+1] (cost[i][j+1])
	       如果是删除的话, 相当于是由word1[:i+1] -> word2[:j] (cost[i+1][j])
	       如果是替换的话, 相当于是由word1[:i] -> word2[:j]   (cost[i][j])

	   由上边的dp可得: 只和上一行有关. 但是需要关注 leftTop, top, left 三个方向
	*/
	if word1 == word2 {
		return 0
	}
	cost := make([]int, len(word2)+1)

	// 特殊情况1: cost[0][...]的处理, 相当于将 "" -> word2[:j]的开销, 都是插入
	for i := range cost {
		cost[i] = i
	}

	for w1Idx := 1; w1Idx <= len(word1); w1Idx++ {
		// word1[:w1Idx] -> word2[:0], 需要删除多少个字符
		cost[0] = w1Idx
		// 左上方, 初始值对应的就是dp[w1Idx-1][0]
		leftTop := w1Idx - 1
		for w2Idx := 1; w2Idx <= len(word2); w2Idx++ {
			// 左方
			left := cost[w2Idx-1]
			// 正上方
			top := cost[w2Idx]

			if word1[w1Idx-1] == word2[w2Idx-1] {
				cost[w2Idx] = leftTop
			} else {
				cost[w2Idx] = min(left, min(top, leftTop)) + 1
			}
			// 下一列的 leftTop 就是这一列的top
			leftTop = top
		}
	}
	return cost[len(word2)]
}

func minDistance4(word1 string, word2 string) int {
	/*
	   dp[i+1][j+1] = word1[:i+1] -> word2[:j+1]所需要的最小步骤
	   针对 word1[i]和word2[j]而言
	   相等:
	       不需要任何步骤
	   不相等: 需要从以下步骤中找出最小的开销+1
	       如果是插入的话, 相当于是由word1[:i] -> word2[:j+1] (dp[i][j+1])
	       如果是删除的话, 相当于是由word1[:i+1] -> word2[:j] (dp[i+1][j])
	       如果是替换的话, 相当于是由word1[:i] -> word2[:j]   (dp[i][j])

	   由上边的dp可得: 只和上一行有关. 但是需要关注 leftTop, top, left 三个方向
	*/
	if word1 == word2 {
		return 0
	}
	lenWord1 := len(word1)
	lenWord2 := len(word2)
	// dp[i][j] -> word1[:i] -> word2[:j]的开销
	dp := make([][]int, lenWord1+1)
	for i := range dp {
		dp[i] = make([]int, lenWord2+1)
		// word1[:i] -> word2[:0], 需要删除i个字符
		dp[i][0] = i
	}
	for i := range dp[0] {
		// word1[:0] -> word2[:i], 需要添加i个字符
		dp[0][i] = i
	}

	for idx1 := 1; idx1 <= lenWord1; idx1++ {
		for idx2 := 1; idx2 <= lenWord2; idx2++ {
			if word1[idx1-1] == word2[idx2-1] {
				dp[idx1][idx2] = dp[idx1-1][idx2-1]
			} else {
				// 添加
				// 删除
				// 替换
				dp[idx1][idx2] = min(dp[idx1-1][idx2], dp[idx1][idx2-1], dp[idx1-1][idx2-1]) + 1
			}
		}
	}
	return dp[lenWord1][lenWord2]
}

func main() {
	fmt.Println(minDistance("ab", "a"))
}
