package main

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// left[i][j]: 第i行水平方向上到j的连续1的个数
	left := make([][]int, m+1)
	// up[i][j]: 第j列竖直方向上到i的连续1的个数
	up := make([][]int, m+1)
	for i := range left {
		left[i] = make([]int, n+1)
		up[i] = make([]int, n+1)
	}
	maxBorder := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				// 以grid[i][j]为右端点,
				left[i][j] = left[i][j-1] + 1
				up[i][j] = up[i-1][j] + 1
				border := min(left[i][j], up[i][j])
				// 看左上角横向连续1的个数, 和右下角纵向连续1的个数
				for left[i-border+1][j] < border || up[i][j-border+1] < border {
					border--
				}
				maxBorder = max(maxBorder, border)
			}
		}
	}
	return maxBorder * maxBorder
}
