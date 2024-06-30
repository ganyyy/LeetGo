package main

func maximalPathQuality(values []int, edges [][]int, maxTime int) (ans int) {
	n := len(values)
	type edge struct{ to, time int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, t})
		g[y] = append(g[y], edge{x, t})
	}

	vis := make([]bool, n)
	vis[0] = true
	var dfs func(int, int, int)
	dfs = func(x, sumTime, sumValue int) {
		if x == 0 {
			ans = max(ans, sumValue)
			// 注意这里没有 return，还可以继续走
		}
		for _, e := range g[x] {
			y, t := e.to, e.time
			if sumTime+t > maxTime {
				continue
			}
			if vis[y] {
				dfs(y, sumTime+t, sumValue)
			} else {
				vis[y] = true
				// 每个节点的价值至多算入价值总和中一次
				dfs(y, sumTime+t, sumValue+values[y])
				vis[y] = false // 恢复现场
			}
		}
	}
	dfs(0, 0, values[0])
	return ans
}
