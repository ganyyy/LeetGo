package main

func islandPerimeter(grid [][]int) int {
	var res int

	var ln = len(grid)
	if ln == 0 {
		return res
	}
	var lm = len(grid[0])
	if lm == 0 {
		return res
	}

	for i := 0; i < ln; i++ {
		for j := 0; j < lm; j++ {
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}
