package main

import "math"

func numberOfSets(n, maxDistance int, roads [][]int) (ans int) {
	// 构建邻接表
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			if j != i { // graph[i][i] = 0
				graph[i][j] = math.MaxInt / 2 // 防止加法溢出
			}
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		graph[x][y] = min(graph[x][y], wt)
		graph[y][x] = min(graph[y][x], wt)
	}

	collect := make([][]int, n)
	for i := range collect {
		collect[i] = make([]int, n)
	}
next:
	for selected := 0; selected < 1<<n; selected++ { // 枚举子集
		// 将本次可选的点加入到f中, 通过Floyd算法找出最短路径
		for point, row := range graph {
			if selected>>point&1 == 0 {
				continue
			}
			copy(collect[point], row)
		}

		// Floyd
		// start -> end 经过某个中间节点 mid 可以使得整体耗时最短.
		for mid := range collect {
			if selected>>mid&1 == 0 {
				continue
			}
			for start := range collect {
				if selected>>start&1 == 0 {
					continue
				}
				for end := range collect {
					collect[start][end] = min(collect[start][end], collect[start][mid]+collect[mid][end])
				}
			}
		}

		// 记录任意两点之间的最短距离, 如果不超过 maxDistance 就可以算是一种解决方案
		for start, distances := range collect {
			if selected>>start&1 == 0 {
				continue
			}
			for end, distance := range distances {
				if selected>>end&1 > 0 && distance > maxDistance {
					continue next
				}
			}
		}
		ans++
	}
	return
}
