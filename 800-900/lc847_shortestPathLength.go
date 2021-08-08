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
			// 理论上, 每条路径最多只能经过两次, 多余的次数不应该算进去
			// 第一次 标记为 v->u进行标记, 然后 u->v也会进行标记
			//  v->u 和 v-u 的maskV的值是相同的
			if !seen[v][maskV] {
				q = append(q, tuple{u: v, mask: maskV, dist: t.dist + 1})
				seen[v][maskV] = true
			}
		}
	}
}
