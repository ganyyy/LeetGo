package main

import "math"

func minScoreTriangulation(values []int) int {
	memo := make(map[int]int)
	n := len(values)
	var dp func(int, int) int
	dp = func(i int, j int) int {
		// i, j 是从顺时针走的两个端点
		// 最小的三角形, 至少是 i,j,k 之间相差1
		if i+2 > j {
			return 0
		}
		if i+2 == j {
			// i, i+1, i+2
			return values[i] * values[i+1] * values[j]
		}
		// [i....k....j]
		// 一个[i:j] 围成的凸包, 可以经过中间的k划分为三部分
		// [i...k]: 一个凸包
		// (i, j, k): 一个三角形
		// [k...j]: 一个凸包
		key := i*n + j
		if _, ok := memo[key]; !ok {
			minScore := math.MaxInt32
			// 找到可以使 dp[i,j]最小的凸包
			for k := i + 1; k < j; k++ {
				minScore = min(minScore, values[i]*values[k]*values[j]+dp(i, k)+dp(k, j))
			}
			memo[key] = minScore
		}
		return memo[key]
	}
	return dp(0, n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
