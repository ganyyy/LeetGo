package main

func cherryPickup(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	current := make([][]int, col)
	next := make([][]int, col)
	for i := range current {
		current[i] = make([]int, col)
		next[i] = make([]int, col)
		for j := range current[i] {
			current[i][j] = -1
			next[i][j] = -1
		}
	}

	current[0][col-1] = grid[0][0] + grid[0][col-1]
	// 一共只能走row-1层
	for layer := 1; layer < row; layer++ {
		for j1 := 0; j1 < col; j1++ {
			for j2 := 0; j2 < col; j2++ {
				best := -1
				// 迭代所有可能的组合
				for dj1 := j1 - 1; dj1 <= j1+1; dj1++ {
					for dj2 := j2 - 1; dj2 <= j2+1; dj2++ {
						if dj1 >= 0 && dj1 < col && dj2 >= 0 && dj2 < col && current[dj1][dj2] != -1 {
							if j1 == j2 {
								// 相等的情况下, 只能取一个
								best = max(best, current[dj1][dj2]+grid[layer][j1])
							} else {
								best = max(best, current[dj1][dj2]+grid[layer][j1]+grid[layer][j2])
							}
						}
					}
				}
				next[j1][j2] = best
			}
		}
		// 当前层需要前一层的状态
		// 每一层的组合方式, 为 col*n种.
		// current[x1][x2]分别表示为当机器人1处于当前层的x1的位置, 机器人2处于当前层x2的位置时, 带来的最大收益
		current, next = next, current
	}

	// 迭代最终的结果, 取所有组合的最大值
	ans := 0
	for _, row := range current {
		for _, v := range row {
			ans = max(ans, v)
		}
	}
	return ans
}
