package main

// 一般牵扯到图的问题, 要么深度优先遍历, 要么广度优先遍历
func isBipartite(graph [][]int) bool {

	// 用来存储每一个点的颜色数据
	color := make([]int, len(graph))

	// 是否符合题目要求
	// l是第几条边, c是当前边首节点的颜色
	// return 当前边是否满足二分
	var dfs func(l, c int) bool

	// 这里采用深度优先遍历
	dfs = func(l, c int) bool {
		color[l] = c
		// 遍历指定点所有的邻接点, 看看是否有相同的颜色
		for column, i := graph[l], 0; i < len(column); i++ {
			cc := column[i]
			// 如果没有上过标记
			if color[cc] == 0 {
				// 这里只要求有两种颜色, 所以对于当前边, 每一个未访问过的点的颜色都要和当前颜色不相同
				if !dfs(cc, -c) {
					return false
				}
			} else {
				// 如果同一边中出现了颜色相等边, 也返回错误
				if color[cc] == c {
					return false
				}
			}
		}
		return true
	}

	// 遍历每一个头节点
	for i := 0; i < len(graph); i++ {
		if color[i] == 0 {
			if !dfs(i, -1) {
				return false
			}
		}
	}
	return true
}
