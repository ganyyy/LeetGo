package main

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	// 构建连通图, 看剩余为连接的节点的数量以及多余的连接的数量
	var fa = make([]int, n)
	for i := range fa {
		fa[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	// 连通所有的点
	for _, connection := range connections {
		var a, b = connection[0], connection[1]
		a, b = find(a), find(b)
		if a != b {
			fa[a] = b
		}
	}

	var cnt int
	for i := range fa {
		if i == find(i) {
			cnt++
		}
	}

	return cnt - 1
}
