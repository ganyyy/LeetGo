package main

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := make([][]bool, n)
	for i := range seen {
		// 这里使用 1<<n 个bool 组成的数组是因为 每个位置上有 n种可能
		// 这个可以压缩一下, 可以再次按位压缩
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, tuple{i, 1 << i, 0})
	}

	for {
		t := q[0]
		q = q[1:]
		if t.mask == 1<<n-1 {
			// 表示的是所有点都经过了
			// 首次出现的一定是距离最短的.
			return t.dist
		}
		// 搜索相邻的节点
		for _, v := range graph[t.u] {
			// 标记该点已经过
			maskV := t.mask | 1<<v
			// 如果该点已经检查过了, 就不用二次处理?
			if !seen[v][maskV] {
				q = append(q, tuple{v, maskV, t.dist + 1})
				seen[v][maskV] = true
			}
		}
	}
}
