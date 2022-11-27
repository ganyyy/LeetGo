package main

import "math"

func largestSumOfAverages2(nums []int, k int) float64 {
	n := len(nums)
	// 前缀和
	prefix := make([]float64, n+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + float64(x)
	}
	// dp[i][j]: nums[0:i]分成j个子数组的最大平均和
	// i >= j
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}
	// 单独处理 当 j == 1的情况
	for i := 1; i <= n; i++ {
		dp[i][1] = prefix[i] / float64(i)
	}
	// 可以进行压缩, 因为只用到了 j-1 的状态
	// 迭代所有的分组可能
	for j := 2; j <= k; j++ {
		// 迭代数组的长度
		for i := j; i <= n; i++ {
			// 迭代分组x的位置
			for x := j - 1; x < i; x++ {
				// nums[0, x], nums[x, i]
				// 从x开始进行分组, 意味着
				// x点分j-1组的最大值, 加上 nums[x,i]对应的平均值
				dp[i][j] = math.Max(dp[i][j], dp[x][j-1]+(prefix[i]-prefix[x])/float64(i-x))
			}
		}
	}
	return dp[n][k]
}

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	prefix := make([]float64, n+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + float64(x)
	}
	dp := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = prefix[i] / float64(i)
	}
	for j := 2; j <= k; j++ {
		// 为啥要倒着算i呢?
		// 因为后依赖于前,
		// 如果从前向后迭代的话, 前边就先更新了, 导致后边的计算不准确
		for i := n; i >= j; i-- {
			for x := j - 1; x < i; x++ {
				dp[i] = math.Max(dp[i], dp[x]+(prefix[i]-prefix[x])/float64(i-x))
			}
			if j == k {
				break
			}
		}
	}
	return dp[n]
}
