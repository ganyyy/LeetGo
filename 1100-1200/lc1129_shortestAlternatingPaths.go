package main

func shortestAlternatingPaths(n int, redEdges, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	g := make([][]pair, n)
	// 整合红蓝两条路径
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 1})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	vis := make([][2]bool, n)
	vis[0] = [2]bool{true, true}
	// 0作为起点
	q := []pair{{0, 0}, {0, 1}}
	for level := 0; len(q) > 0; level++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			x := p.x
			if dis[x] < 0 {
				// 首次出现的肯定就是最短路径, 直接标记即可
				dis[x] = level
			}
			for _, to := range g[x] {
				// 找出和当前颜色不同, 并且未标记的节点作为候选节点
				if to.color != p.color && !vis[to.x][to.color] {
					vis[to.x][to.color] = true
					q = append(q, to)
				}
			}
		}
	}
	return dis
}
