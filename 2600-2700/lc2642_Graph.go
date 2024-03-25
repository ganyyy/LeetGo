package main

import "math"

const inf = math.MaxInt / 2 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n) // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // 添加一条边（题目保证没有重边）
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2] // 添加一条边（题目保证这条边之前不存在）
}

func (g Graph) ShortestPath(start, end int) int {
	n := len(g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
	for i := range dis {
		dis[i] = inf
	}
	// 首次选出来的一定是start
	// dijkstra算法找到的是从start到各个点的最短路径, 如果end可达, 那么最终dis[end]就是最短路径
	dis[start] = 0
	vis := make([]bool, n)
	for {
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		// x < 0 表示找不到下一个可以迭代的点
		// dis[x] == inf 表示找到的最短路径点不可达
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == end { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
		}
	}
}
