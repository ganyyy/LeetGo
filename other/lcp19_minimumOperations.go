package main

import "math"

func minimumOperations(leaves string) int {
	if len(leaves) < 3 {
		return 0
	}
	// dp[i] 表示 leave[0:i] 之间的调整操作, 我们把每一个位置分成三个部分进行讨论
	// i[0]: 表示调整到第一个红的最少次数
	// 		 当且仅当 i-1[0]时, i[0]是一个合法的变换, 且需要满足变换的条件 isYellow(i)
	// i[1]: 表示调整到第二个黄的最少次数
	//		 当且仅当 i-1[0]或者i-1[1]时, i[1]是一个合法的变换, 且需要满足条件 isRed(i)。取 i-1[0]和i-1[1]之间的最小值
	// i[2]: 表示调整到第三个红的最少次数
	//		 当且仅当 i-1[1]或者i-1[2]时, i[2]是一个合法的变化, 且需要满足 isYellow(i)。取 i-1[1]和i-1[2]之间的最小值

	// 最终的答案就是 dp[len(leaves)-1][2]
	// 需要注意的是: 叶子的数量必须大于状态的数量
	dp := make([][3]int, len(leaves))
	// 初始化起始条件
	dp[0][0] = isYellow(leaves[0])
	dp[0][1], dp[0][2], dp[1][2] = math.MaxInt32, math.MaxInt32, math.MaxInt32

	for i := 1; i < len(leaves); i++ {
		ar, ay := getAddRedYellow(leaves[i])
		dp[i][0] = dp[i-1][0] + ay
		dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + ar
		// 因为每种状态最少需要1个叶子, 所以说 对于状态2来说, 必须要在 i >= 2的时候才开始计算
		if i >= 2 {
			dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + ay
		}
	}
	return dp[len(leaves)-1][2]
}

func isYellow(s byte) int {
	if s == 'y' {
		return 1
	}
	return 0
}

func getAddRedYellow(s byte) (red, yellow int) {
	if s == 'r' {
		return 1, 0
	} else {
		return 0, 1
	}
}

func isRed(s byte) int {
	if s == 'r' {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
