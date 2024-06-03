package main

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	count := len(edges) + 1
	type edge struct{ to, wt int }
	graph := make([][]edge, count)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		graph[x] = append(graph[x], edge{y, wt})
		graph[y] = append(graph[y], edge{x, wt})
	}

	ret := make([]int, count)
	for idx, next := range graph {
		var cnt int
		var dfs func(int, int, int)
		dfs = func(cur, father, total int) {
			if total%signalSpeed == 0 {
				cnt++
			}
			for _, e := range graph[cur] {
				if e.to != father {
					dfs(e.to, cur, total+e.wt)
				}
			}
			return
		}
		// 为啥可以这么做捏?
		// 首先, 树中的每个节点都是独一无二的, 那么只要左右两边的节点不相等,
		// 那么就可以通过当前节点串联
		sum := 0
		for _, p := range next {
			cnt = 0
			dfs(p.to, idx, p.wt)
			// 先乘算, 再相加. 乘法原理.
			ret[idx] += cnt * sum
			sum += cnt
		}
	}
	return ret
}
