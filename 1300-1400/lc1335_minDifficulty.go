//go:build ignore

package main

import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		// 至少需要d个任务
		return -1
	}
	// dp[i][j]
	// i: d, 要求再第i+1天完成任务
	// j: 在给定的期限内完成多少个任务
	dp := make([][]int, d)
	for i := range dp {
		ndp := make([]int, n)
		for i := range ndp {
			ndp[i] = math.MaxInt32
		}
		dp[i] = ndp
	}

	var ma int
	for i, jd := range jobDifficulty {
		ma = max(ma, jd)
		// 第一天完成任务的话, 就对应着迭代到某一个任务的最大值
		dp[0][i] = ma
	}

	for i := 1; i < d; i++ {
		// 先迭代天数, 前置的天数作为后置的前提条件
		// [i,n)表示的是这一天可选的任务
		for j := i; j < n; j++ {
			// 这里真么理解呢?
			// j代表的是可选的任务, 从第(i+1)天开始, 前边至少已经选择了
			// i个任务, 所以下界是i, 上界是任务的数量n
			ma = 0
			// 为什么要这样迭代k呢?
			// 可以理解为: 前N天完成[0,j]的最小值已经计算出来了
			// 此时要从[i,j]这段区间内再找到一个分割点, 使得
			// 使得 今天 + 前N天 的累计 工作难度最低
			// 这个分割点就是k, 前面i-1天负责[0,k-1], 今天负责[k,j]
			// 为啥要从后向前迭代呢?  这里有一个很有意思的点
			// 如果较大值本身在后边, 那么是无可厚非的, 因为较大值执行的时候无法和前边的更大值聚合到一天
			// 比如[6,5,4,3,2,1]和[6,4,3,2,1,5]. 假设现在是第二天, 需要从数组下标为1的地方开始迭代
			// 如果第一种情况从前往后迭代, 那么很明显这段区间的最大值从一开始就确定了
			// 如果是从后往前迭代, 就会优先选取相对较小的值作为分割点
			// 对比情况2也是一样的
			for k := j; k >= i; k-- {
				ma = max(ma, jobDifficulty[k])
				dp[i][j] = min(dp[i][j], ma+dp[i-1][k-1])
			}
		}
	}
	return dp[d-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
