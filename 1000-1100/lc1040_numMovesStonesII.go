package main

import "sort"

func numMovesStonesII(s []int) []int {
	sort.Ints(s)
	n := len(s)
	// 最大的移动次数, 等同于 所有的空位数
	e1 := s[n-2] - s[0] - n + 2
	e2 := s[n-1] - s[1] - n + 2 // 计算空位
	maxMove := max(e1, e2)
	if e1 == 0 || e2 == 0 { // 特殊情况：没有空位
		return []int{min(2, maxMove), maxMove}
	}
	// 最小的移动次数
	maxCnt, left := 0, 0
	for right, x := range s { // 滑动窗口：枚举右端点
		for s[left] <= x-n { // 窗口大小大于 n
			left++
		}
		maxCnt = max(maxCnt, right-left+1) // 维护窗口内的最大石子数
	}
	return []int{n - maxCnt, maxMove}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
