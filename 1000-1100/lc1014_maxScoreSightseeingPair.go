package main

import (
	"fmt"
	"math"
)

/*

给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
返回一对观光景点能取得的最高分。


*/

func maxScoreSightseeingPair(A []int) int {
	if len(A) < 1 {
		return 0
	}
	// left = max(left, A[i]+i)
	var left = A[0]
	// res = max(res, left+A[i]-i)
	var res = math.MinInt32
	for i, v := range A[1:] {
		if t := left + v - i - 1; t > res {
			res = t
		}
		if t := v + i + 1; t > left {
			left = t
		}
	}
	return res
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{1, 3, 5}))
}
