package main

var dir = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func getMaximumGold(grid [][]int) int {
	// 如何让一个点可以多次使用?

	// 如果一个点四周都不可达, 那他就是终点

	// 如果一个点有多条路径可达, 如何标识?

	// 大胆点呗...

	var lm, ln = len(grid), len(grid[0])

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		var val = grid[i][j]
		var n int
		grid[i][j] = 0
		for _, d := range dir {
			var ni, nj = i + d[0], j + d[1]
			if ni < 0 || ni >= lm {
				continue
			}
			if nj < 0 || nj >= ln {
				continue
			}
			if grid[ni][nj] == 0 {
				continue
			}
			n = max(n, dfs(ni, nj))
		}
		grid[i][j] = val
		return val + n
	}
	var ret int
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				continue
			}
			ret = max(ret, dfs(i, j))
		}
	}

	return ret
}
