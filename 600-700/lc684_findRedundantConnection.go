package main

func findRedundantConnection(edges [][]int) []int {
	// 还是并查集

	var p = make([]int, len(edges))
	for i := range p {
		p[i] = i + 1
	}

	var find func(i int) int
	find = func(i int) int {
		if p[i-1] != i {
			p[i-1] = find(p[i-1])
		}
		return p[i-1]
	}

	var merge = func(a, b int) bool {
		// 父节点相同说明组成了环
		var pa, pb = find(a), find(b)
		if pa == pb {
			return false
		}
		// 否则正常合并即可
		p[pa-1] = pb
		return true
	}

	for _, e := range edges {
		// 如果出现了相同的父节点, 说明这个边就是构成环的最后一条边
		if !merge(e[0], e[1]) {
			return e
		}
	}

	return nil
}
