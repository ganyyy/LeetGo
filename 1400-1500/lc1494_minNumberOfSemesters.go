//go:build ignore

package main

import (
	"math"
	"math/bits"
)

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	need := make([]int, 1<<n)
	for _, edge := range relations {
		// 所有的前置课程对应的位都置为1
		need[1<<(edge[1]-1)] |= 1 << (edge[0] - 1)
	}
	dp[0] = 0
	for i := 1; i < (1 << n); i++ {
		// i&(i-1)
		// i&-i 代表?
		need[i] = need[i&(i-1)] | need[i&-i]
		if (need[i] | i) != i { // i 中有任意一门课程的前置课程没有完成学习
			continue
		}
		valid := i ^ need[i]                  // 当前学期可以进行学习的课程集合
		if bits.OnesCount(uint(valid)) <= k { // 如果个数小于 k，则可以全部学习，不再枚举子集
			dp[i] = min(dp[i], dp[i^valid]+1)
		} else {
			for sub := valid; sub > 0; sub = (sub - 1) & valid {
				if bits.OnesCount(uint(sub)) <= k {
					dp[i] = min(dp[i], dp[i^sub]+1)
				}
			}
		}
	}
	return dp[(1<<n)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
