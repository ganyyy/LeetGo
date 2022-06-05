package main

import "sort"

func makesquare(matchsticks []int) bool {
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}
	// 从大到小排序
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // 减少搜索量

	var edges [4]int
	// 平均值
	var avg = totalLen >> 2
	// 依次尝试将边放入进去, 综合计算出一个可行的结果
	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return true
		}
		for i := range edges {
			edges[i] += matchsticks[idx]
			if edges[i] <= avg && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}
	return dfs(0)
}
