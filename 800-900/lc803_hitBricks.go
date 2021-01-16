package main

var dir = [4][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func hitBricks(grid [][]int, hits [][]int) []int {
	// 根节点X
	// 所有边缘节点都指向X
	// 连接边缘节点的值合并成同一个并查集

	// h行, w列. 那么格子可以使用 i(行)*w + j(第几个)的形式进行表示
	var h, w = len(grid), len(grid[0])
	// 根节点X映射到 h*w

	// fa表示每个节点的父节点
	var fa = make([]int, h*w+1)
	// size表示当i为某个联通分支祖先时才有效. 表示以i为祖先的连通分支节点数量
	// 合并两个连通分支时注意要合并这个值
	var size = make([]int, h*w+1)

	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	// 正向遍历hits数组, 将所有被消除的点标记为0

	var find func(int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	var union = func(from, to int) {
		var ff, ft = find(from), find(to)
		if ff == ft {
			return
		}
		// 增加稳定节点数量
		size[ft] += size[ff]
		// 连接ff和ft
		fa[ff] = ft
	}

	// 首先正向遍历一遍hits, 临时删除点, 标记为2
	for _, hit := range hits {
		if grid[hit[0]][hit[1]] != 0 {
			grid[hit[0]][hit[1]] = 2
		}
	}

	var root = h * w

	// 构建初始的连通图
	var cur int
	for i, row := range grid {
		for j, v := range row {
			if v&1 == 0 {
				continue
			}
			// 边缘节点直接和根节点联通
			if i == 0 {
				union(j, root)
			}
			cur = i*w + j
			// 如果上边是稳固的, 联通一下
			if i > 0 && grid[i-1][j] == 1 {
				union(cur, cur-w)
			}
			// 如果左边是稳固的, 联通一下
			if j > 0 && grid[i][j-1] == 1 {
				union(cur, cur-1)
			}
		}
	}

	var res = make([]int, len(hits))

	for i := len(hits) - 1; i >= 0; i-- {
		var hit = hits[i]
		var r, c = hit[0], hit[1]
		if grid[r][c] == 0 {
			continue
		}

		var preSize = size[find(root)]
		// 对于首列, 直接合并到root节点
		if r == 0 {
			union(c, root)
		}
		// 挨个方向查找
		for _, d := range dir {
			var newR, newC = r + d[0], c + d[1]
			if 0 > newR || newR >= h || 0 > newC || newC >= w {
				continue
			}
			// 合并新连接的稳固节点
			if grid[newR][newC] == 1 {
				union(r*w+c, newR*w+newC)
			}
		}

		var curSize = size[find(root)]
		// 如果新加入的节点使得总的连通节点数发生了变化
		// 就
		if cnt := curSize - preSize - 1; cnt > 0 {
			res[i] = cnt
		}
		grid[r][c] = 1
	}

	return res
}

func hitBricks2(grid [][]int, hits [][]int) []int {
	h, w := len(grid), len(grid[0])
	fa := make([]int, h*w+1)
	size := make([]int, h*w+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(from, to int) {
		from, to = find(from), find(to)
		if from != to {
			size[to] += size[from]
			fa[from] = to
		}
	}

	status := make([][]int, h)
	for i, row := range grid {
		status[i] = append([]int(nil), row...)
	}
	// 遍历 hits 得到最终状态
	for _, p := range hits {
		status[p[0]][p[1]] = 0
	}

	// 根据最终状态建立并查集
	root := h * w
	for i, row := range status {
		for j, v := range row {
			if v == 0 {
				continue
			}
			if i == 0 {
				union(i*w+j, root)
			}
			if i > 0 && status[i-1][j] == 1 {
				union(i*w+j, (i-1)*w+j)
			}
			if j > 0 && status[i][j-1] == 1 {
				union(i*w+j, i*w+j-1)
			}
		}
	}

	type pair struct{ x, y int }
	directions := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	ans := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		p := hits[i]
		r, c := p[0], p[1]
		if grid[r][c] == 0 {
			continue
		}

		preSize := size[find(root)]
		if r == 0 {
			union(c, root)
		}
		for _, d := range directions {
			if newR, newC := r+d.x, c+d.y; 0 <= newR && newR < h && 0 <= newC && newC < w && status[newR][newC] == 1 {
				union(r*w+c, newR*w+newC)
			}
		}
		curSize := size[find(root)]
		if cnt := curSize - preSize - 1; cnt > 0 {
			ans[i] = cnt
		}
		status[r][c] = 1
	}
	return ans
}
