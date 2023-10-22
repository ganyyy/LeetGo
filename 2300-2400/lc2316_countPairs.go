package main

func countPairs(n int, edges [][]int) int64 {
	// 标记节点是否访问过
	var visited = make([]bool, n)

	// 用邻接表存储图
	var nextEdges = make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		nextEdges[from] = append(nextEdges[from], to)
		nextEdges[to] = append(nextEdges[to], from)
	}

	var countAndVisit func(int) int
	// from: 来源节点
	// 返回值: 从from开始的相关联且未被访问过的节点数量
	countAndVisit = func(from int) int {
		visited[from] = true
		var groupCount = 1
		for _, next := range nextEdges[from] {
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
		// 因为最终结果是求解任意两个分组的组合数, 所以这里是乘法
		// 类似于阶乘, 但是步进不是1, 而是当前分组的节点数量
		total += count * (n - count)
		n -= count
	}

	return int64(total)
}
