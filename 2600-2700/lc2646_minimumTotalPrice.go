package main

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {

	// 所有边的集合
	var graph = make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		from, to = to, from
		graph[from] = append(graph[from], to)
	}

	// 基于每条路径, 计算所有节点经过的最少次数
	var tripCount = make([]int, n)
	var end int

	// 找到一条最短路径. 因为路径数量越少, 花费的开销一定是最低的
	var tripDFS func(current, father int) bool
	tripDFS = func(current, father int) bool {
		if current == end {
			tripCount[current]++
			return true
		}
		for _, next := range graph[current] {
			if next == father {
				continue
			}
			if !tripDFS(next, current) {
				continue
			}
			// 作为中间节点, 路径+1
			tripCount[current]++
			return true
		}
		return false
	}

	for _, trip := range trips {
		end = trip[1]
		tripDFS(trip[0], -1)
	}

	// 聪明的小偷Ⅲ: 树形DP
	// 相邻节点不允许同时half, 所以需要记录两个状态
	var costDFS func(current, father int) (half, noHalf int)
	costDFS = func(current, father int) (half, noHalf int) {
		noHalf = tripCount[current] * price[current]
		half = noHalf / 2
		for _, next := range graph[current] {
			if next == father {
				continue
			}
			nextHalf, nextNoHalf := costDFS(next, current)
			// 相邻的节点不能同时half
			noHalf += min(nextHalf, nextNoHalf)
			half += nextNoHalf
		}
		return
	}

	// 从任意节点出发, 都可以得到最优解, 因为整个树的最小路径不管从哪个节点出发, 都是一样的
	return min(costDFS(0, -1))
}
