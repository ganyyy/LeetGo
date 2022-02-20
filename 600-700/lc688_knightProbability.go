package main

var dx, dy = []int{-2, -2, -1, -1, 2, 2, 1, 1}, []int{-1, 1, -2, 2, -1, 1, -2, 2}

func knightProbability(N int, K int, r int, c int) float64 {
	// 求移动结束后，“马” 仍留在棋盘上的概率。
	// 每移动一步是8种可能
	// dp[k][i][j], 第k次，落在（i，j）的概率
	// dp[k][i][j] = forloop (sum(dp[k-1][dx][dy]/8))
	// 最后求sum(dp[k])

	// 初始化第0步的各种情况
	dp := make([][][]float64, K+1)
	dp[0] = make([][]float64, N)
	for i := range dp[0] {
		dp[0][i] = make([]float64, N)
	}
	// 初始位置情况下, 落在(r, c)中的概览是100%
	dp[0][r][c] = 1
	for k := 1; k <= K; k++ {
		// 计算第i步的情况
		dp[k] = make([][]float64, N)
		for i := range dp[k] {
			dp[k][i] = make([]float64, N)
			for j := range dp[k][i] {
				for l := range dx {
					// 迭代整个棋盘上的坐标, 计算所有可以到达的位置的
					px, py := i+dx[l], j+dy[l]
					// 这里计算的其实是 从k-1步到k步, 所需要的点位信息
					// 如果越界了, 或者前一步对应的左边中没有落脚, 那么就不需要统计对应的概率
					if px < 0 || px >= N || py < 0 || py >= N || dp[k-1][px][py] == 0 {
						continue
					}
					// 整合所有的概率
					dp[k][i][j] += dp[k-1][px][py] / 8
				}
			}
		}
	}
	// 最终结果是所有点位上存在的概率和
	sum := float64(0)
	for i := range dp[K] {
		for j := range dp[K][i] {
			sum += dp[K][i][j]
		}
	}
	return sum
}
