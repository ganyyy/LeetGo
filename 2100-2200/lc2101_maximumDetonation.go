package main

func maximumDetonation(bombs [][]int) (ans int) {
	n := len(bombs)
	g := make([][]int, n)
	for i, p := range bombs {
		x, y, r := p[0], p[1], p[2]
		for j, q := range bombs {
			dx := x - q[0]
			dy := y - q[1]
			if dx*dx+dy*dy <= r*r {
				// 有向图, 因为引爆是不可逆的.
				g[i] = append(g[i], j) // i 可以引爆 j
			}
		}
	}

	// 依次记录以每个节点为起点, dfs看引爆的数量
	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(x int) int {
		vis[x] = true
		cnt := 1
		for _, y := range g[x] {
			if !vis[y] {
				cnt += dfs(y)
			}
		}
		return cnt
	}
	for i := range g {
		clear(vis)
		ans = max(ans, dfs(i))
	}
	return
}
