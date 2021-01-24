package main

func regionsBySlashes(grid []string) int {
	// 日你奶奶的, 全TM阴间题目

	var ln = len(grid)

	// 将一个小方块划分 按照顺时针划分为4个小方块, 根据不同情况合并
	/**
	  \ 0/
	  3\/
	   /\1
	  /2 \
	*/
	var fa = make([]int, ln*ln*4)
	for i := range fa {
		fa[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if i != fa[i] {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	var union = func(from, to int) {
		fa[find(from)] = find(to)
	}

	for i := 0; i < ln; i++ {
		for j := 0; j < ln; j++ {
			// 小方块对应的起点坐标
			var start = 4 * ((i * ln) + j)
			switch grid[i][j] {
			case ' ':
				// 全合并
				union(start, start+1)
				union(start+1, start+2)
				union(start+2, start+3)
			case '/':
				// 0, 3 合并
				// 1, 2 合并
				union(start, start+3)
				union(start+1, start+2)
			case '\\':
				// 0, 1合并
				// 2, 3合并
				union(start, start+1)
				union(start+2, start+3)
			}
			if i > 0 {
				// 如果过了一行, 当前行的0和上一行的2合并
				union(start, start-4*ln+2)
			}
			if j > 0 {
				// 如果过了一列, 当前列的3和上一列的1合并
				union(start+3, start-3)
			}
		}
	}

	// 统计联通区域的数量
	var cnt int
	for i := range fa {
		// 如果父节点是自己, 说明这是一个根节点?
		// 这里写成 i == fa[i]和 i == find(i)是一个意思
		// 本质上只有在父节点是自己的情况下才会相等
		if i == find(i) {
			cnt++
		}
	}
	return cnt
}
