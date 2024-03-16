package main

func maxMoves(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	for _, row := range grid {
		row[0] *= -1 // 入队标记
	}
	for col := 0; col < cols-1; col++ {
		findNext := false
		for row := 0; row < rows; row++ {
			current := grid[row][col]
			if current > 0 { // 不在队列中
				continue
			}
			current = -current
			// (row-1, col+1), (row, col+1), (row+1, col+1) 三个位置尝试入队
			for nextRow := max(row-1, 0); nextRow < min(row+2, rows); nextRow++ {
				if grid[nextRow][col+1] > current {
					grid[nextRow][col+1] *= -1 // 入队标记
					findNext = true
				}
			}
		}
		if !findNext { // 无法再往右走了
			return col
		}
	}
	// 上线就是列数-1
	return cols - 1
}
