//go:build ignore

package main

import "math"

func minSideJumps(obstacles []int) int {
	d := [...]int{1, 0, 1}
	// dp
	// 统计三条道路的最小值
	for _, x := range obstacles[1:] {
		minCnt := math.MaxInt / 2
		for j := 0; j < 3; j++ {
			if j == x-1 {
				// 被堵住了
				d[j] = math.MaxInt / 2
			} else {
				// 可以前进
				minCnt = min(minCnt, d[j])
			}
		}
		for j := 0; j < 3; j++ {
			if j != x-1 {
				// 如果不是被堵住的位置, 就更新一波前进所需要的消耗
				// 此时, 这个位置上没有被堵住的路的开销是相同的
				d[j] = min(d[j], minCnt+1)
			}
		}
	}
	return min(min(d[0], d[1]), d[2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
