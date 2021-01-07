package main

func findCircleNum(isConnected [][]int) int {
	// 并查集 a 和 b 联通, 那么 a -> b
	// 最终结果就是并查集中根不相同的数量

	var fa = make([]int, len(isConnected))

	for i := range fa {
		fa[i] = i
	}

	var find func(c int) int

	find = func(c int) int {
		if fa[c] != c {
			fa[c] = find(fa[c])
		}
		return fa[c]
	}

	var res = len(isConnected)
	for i, row := range isConnected {
		for j, c := range row {
			if c == 1 {
				// 如果是连接的, 但是不具有相同的父节点, 那就说明这是一个省份的
				if pa, pb := find(i), find(j); pa != pb {
					fa[pa] = pb
					// 总的省份数量-1
					res--
				}
			}
		}
	}
	return res
}

func findCircleNumBfs(isConnected [][]int) int {
	// 这里用dfs 尝试处理
	var ln = len(isConnected)
	var visit = make([]bool, ln)
	var ret int

	var dfs func(i int)

	dfs = func(i int) {
		for j := 0; j < ln; j++ {
			// 找到所有和当前城市连接的城市, 并上一个标记
			if isConnected[i][j] == 1 && !visit[j] {
				visit[j] = true
				dfs(j)
			}
		}
	}

	for i := range isConnected {
		if !visit[i] {
			// 如果是一个未访问过的城市, 就找出和这个城市相关联的城市组成一个省
			dfs(i)
			ret++
		}
	}

	return ret
}
