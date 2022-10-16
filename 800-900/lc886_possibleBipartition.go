package main

func possibleBipartition(n int, dislikes [][]int) bool {
	dislikeMap := make([][]int, n+1)

	for _, dislike := range dislikes {
		dislikeMap[dislike[0]] = append(dislikeMap[dislike[0]], dislike[1])
		dislikeMap[dislike[1]] = append(dislikeMap[dislike[1]], dislike[0])
	}

	const (
		S = 3
		A = 1
		B = 2
	)

	group := make([]int, n+1)

	var dfs func(p int, g int) bool

	dfs = func(p int, g int) bool {
		if group[p] != 0 {
			return group[p] == g
		}
		group[p] = g

		for _, op := range dislikeMap[p] {
			if !dfs(op, S^g) {
				return false
			}
		}
		return true
	}

	for p, dislike := range dislikeMap {
		if len(dislike) == 0 {
			continue
		}
		g := group[p]
		// 默认放到A组
		if g == 0 {
			g = A
		}
		if !dfs(p, g) {
			return false
		}
	}
	return true
}
