package main

func countPairs(n int, edges [][]int) int64 {
	// 标记节点是否访问过
	var visited = make([]bool, n)

	// 用邻接表存储图
	var nextEdges = make([][]int, n)
	for _, edge := range edges {
		f, t := edge[0], edge[1]
		nextEdges[f] = append(nextEdges[f], t)
		nextEdges[t] = append(nextEdges[t], f)
	}

	var countAndVisit func(int) int
	// p: 来源节点
	countAndVisit = func(point int) int {
		visited[point] = true
		var groupCount = 1
		for _, next := range nextEdges[point] {
			if visited[next] {
				continue
			}
			groupCount += countAndVisit(next)
		}
		return groupCount
	}

	var total int
	// 迭代所有点位, 计数
	for point, visit := range visited {
		if visit {
			continue
		}
		count := countAndVisit(point)
		total += count * (n - count)
		n -= count
	}

	return int64(total)
}
