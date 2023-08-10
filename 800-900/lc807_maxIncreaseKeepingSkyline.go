package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	// 统计每行/每列的最大值

	var maxRow, maxCol = make([]int, len(grid)), make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j, v := range grid[i] {
			maxRow[i] = max(maxRow[i], v)
			maxCol[j] = max(maxCol[j], v)
		}
	}
	var ret int
	for i := 0; i < len(grid); i++ {
		for j, v := range grid[i] {
			ret += min(maxRow[i], maxCol[j]) - v
		}
	}

	return ret
}
