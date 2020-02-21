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

func main() {

}
