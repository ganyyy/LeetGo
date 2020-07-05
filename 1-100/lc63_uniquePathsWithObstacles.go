package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m <= 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n <= 0 {
		return 0
	}
	// 最后是障碍物, 直接返回
	if obstacleGrid[m-1][n-1] != 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] != 0 {
				continue
			}
			if i == 0 || j == 0 {
				if i != 0 {
					dp[i][j] = dp[i-1][j]
				} else if j != 0 {
					dp[i][j] = dp[i][j-1]
				} else {
					// 0,0的情况
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 代码不一定是越少越好
func uniquePathsWithObstaclesNew(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	if row == 0 {
		return 0
	}
	col := len(obstacleGrid[0])
	if col == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 || obstacleGrid[row-1][col-1] == 1 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	// 起点
	dp[0][0] = 1
	// 第一行, 一路走到尾, 只有一种
	for i := 1; i < col; i++ {
		if obstacleGrid[0][i] != 1 {
			dp[0][i] = dp[0][i-1]
		}
	}
	// 第一列, 一路走到尾, 只有一种
	for i := 1; i < row; i++ {
		if obstacleGrid[i][0] != 1 {
			dp[i][0] = dp[i-1][0]
		}
	}
	// 开始走路
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] != 1 {
				// 左边+上边
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[row-1][col-1]
}

func main() {

}
