package main

import (
	"slices"
	"sort"
)

func maxTaxiEarnings(n int, rides [][]int) int64 {
	slices.SortFunc(rides, func(a, b []int) int {
		return a[1] - b[1]
	})
	m := len(rides)
	dp := make([]int, m+1)
	for i, r := range rides {
		j := sort.Search(i+1, func(i int) bool {
			return rides[i][1] >= r[0]+1
		})
		dp[i+1] = max(dp[i], dp[j]+r[1]-r[0]+r[2])
	}
	return int64(dp[m])
}

func maxTaxiEarnings2(n int, rides [][]int) int64 {
	f := make([]int, n+1)
	// 以终点为下标, 保存所有终点为该下标的乘客的起点和小费
	groups := make([][][2]int, n+1)
	for _, r := range rides {
		start, end, tip := r[0], r[1], r[2]
		groups[end] = append(groups[end], [2]int{start, tip}) // 按终点位置分组
	}
	for end := 1; end <= n; end++ { // 从前往后枚举终点
		f[end] = f[end-1]
		for _, r := range groups[end] {
			// 从所有终点为 end 的乘客中选择收益最大的
			start, tip := r[0], r[1]
			f[end] = max(f[end], f[start]+end-start+tip) // 接所有终点为 end 的乘客中收益最大的
		}
	}
	return int64(f[n])
}
