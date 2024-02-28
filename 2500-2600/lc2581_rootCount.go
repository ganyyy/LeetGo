package main

func rootCount2(edges [][]int, guesses [][]int, k int) int {
	// 本方法是: 每次换根都从头构建一下父节点关系数组
	var nodeRoot = make([]int, len(edges)+1)
	var path = make([][]int, len(edges)+1)
	for _, edge := range edges {
		path[edge[0]] = append(path[edge[0]], edge[1])
		path[edge[1]] = append(path[edge[1]], edge[0])
	}

	var buildRoot func(root int, parent int)
	buildRoot = func(root, parent int) {
		nodeRoot[root] = parent
		for _, next := range path[root] {
			if next == parent {
				continue
			}
			buildRoot(next, root)
		}
	}

	var total int

	for i := 0; i <= len(edges); i++ {
		buildRoot(i, -1)
		var validNum int
		for _, guess := range guesses {
			parent, son := guess[0], guess[1]
			if nodeRoot[son] == parent {
				validNum++
			}
			if validNum >= k {
				total++
				break
			}
		}
	}
	return total
}

func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // 建图
	}

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
	}

	// 0为根的时候, 对应的满足条件的根的数量
	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根时，猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	// 换根DP
	var reRoot func(int, int, int)
	reRoot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				// 每次换根, 只会将0->1转变为1->0, 所以减去x->y的猜对次数, 加上y->x的猜对次数
				reRoot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reRoot(0, -1, cnt0)
	return
}
