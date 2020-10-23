package main

import (
	"fmt"
	"math"
)

func videoStitching(clips [][]int, T int) int {
	if len(clips) == 0 {
		return -1
	}

	// dp[i] 是可以到 i 所用的最少的剪辑片段次数
	var dp = make([]int, T+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for _, c := range clips {
			// 起点必须比当前值要小, 终点必须要比当前值大, 不然没有计算的必要
			if c[0] <= i && c[1] >= i {
				// 从当前区间中寻找最小值
				dp[i] = min(dp[i], dp[c[0]]+1)
			}
		}
	}

	if dp[T] == math.MaxInt32 {
		return -1
	}
	return dp[T]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var clips = [][]int{
		{0, 2},
		{4, 6},
		{8, 10},
		{1, 9},
		{1, 5},
		{5, 9},
	}

	fmt.Println(videoStitching(clips, 10))
}
