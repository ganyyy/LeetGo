package main

import (
	"math"
	"sort"
)

func minCostConnectPoints3(points [][]int) int {
	// 通过堆优化一下

	return 0
}

func minCostConnectPoints2(points [][]int) int {
	// Prim算法
	var n = len(points)
	if n <= 1 {
		return 0
	}

	// 首先需要通过一个二维矩阵确定任意两点之间的距离
	var rank = make([][]int, n)
	for i := 0; i < n; i++ {
		rank[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank[i][j] = dist(points[i], points[j])
			rank[j][i] = rank[i][j]
		}
	}

	// n个点
	var dis = make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt32
	}
	// 第一个点的距离为0
	dis[0] = 0

	// 结果 和 当前未选取集合中选取的点
	var ans, from int

	var visited = make([]bool, n)

	for i := 0; i < n; i++ {
		var minDis = math.MaxInt32
		// 找出未访问的, 距离最近的点
		for j := 0; j < n; j++ {
			if !visited[j] && dis[j] < minDis {
				minDis = dis[j]
				from = j
			}
		}
		// 标记节点已访问
		visited[from] = true

		// 总距离加上当前选取点的权重
		ans += dis[from]

		// 更新一下其他节点的距离
		for j := 0; j < n; j++ {
			// 如果存在从当前点到其他未选取的点中 更小的距离, 就更新一下
			if !visited[j] && rank[from][j] < dis[j] {
				dis[j] = rank[from][j]
			}
		}
	}

	return ans
}

func minCostConnectPoints(points [][]int) int {
	type edge struct {
		// 起点, 终点, 长度
		v, w, dis int
	}
	// 最小生成树算法 Kruskal
	// 每次选取最小边进行联通
	var n = len(points)
	// 连通图
	var parent = make([]int, n)
	// 将整个图当成一个完全联通图
	// 首先统计所有点能连成的边
	var edges = make([]edge, 0, n*(n-1))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{
				i, j, dist(points[i], points[j]),
			})
		}
		// 初始化连通图
		parent[i] = i
	}
	// 按照最短距离排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})

	// find查找父节点
	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	// union 将j 合并到 i. 返回是否合并成功
	var union = func(i, j int) bool {
		i, j = find(i), find(j)
		if i == j {
			return false
		}
		parent[j] = i

		return true
	}

	var ans int
	var left = n - 1
	for _, e := range edges {
		if union(find(e.v), find(e.w)) {
			ans += e.dis
			left--
			if left == 0 {
				break
			}
		}
	}

	return ans
}

func dist(a, b []int) int {
	var x = a[0] - b[0]
	var y = a[1] - b[1]
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}
