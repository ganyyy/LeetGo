package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	ln := len(grid)
	if ln <= 0 {
		return 0
	}
	lm := len(grid[0])
	if lm <= 0 {
		return 0
	}
	// 行,列单独赋值
	for i := 1; i < lm; i++ {
		grid[0][i] = grid[0][i] + grid[0][i-1]
	}
	for i := 1; i < ln; i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}

	for i := 1; i < ln; i++ {
		for j := 1; j < lm; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[ln-1][lm-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 2, 5}, {3, 2, 1},
	}))
}
