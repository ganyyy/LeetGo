package main

func shortestBridge(grid [][]int) (step int) {
	// 第一次BFS确定第一座岛
	// 再次BFS确定距离第二座岛的距离
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	// 理是这个理, 写起来真麻烦
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}
			var island []pair
			grid[i][j] = -1
			q := []pair{{i, j}}
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				island = append(island, p)
				for _, d := range dirs {
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						grid[x][y] = -1
						q = append(q, pair{x, y})
					}
				}
			}

			q = island
			for {
				tmp := q
				q = nil
				for _, p := range tmp {
					for _, d := range dirs {
						x, y := p.x+d.x, p.y+d.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1
								q = append(q, pair{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}
	return
}
