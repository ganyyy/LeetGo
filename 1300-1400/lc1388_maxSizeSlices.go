package main

func calculate(slices []int) int {
	N, n := len(slices), (len(slices)+1)/3
	// TODO 可以压缩
	// dp[i][j]: 在前(i+1)个数中选取了j个不相邻的最大值
	// j的最大值就是n
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -0x3f3f3f3f
		}
	}
	// dp[0][0]: 一个数, 一个都不选
	// dp[0][1]: 就一个数, 直接选了
	// dp[1][0]: 两个数, 一个都不选
	// dp[1][1]: 两个数, 选取较大的那个
	dp[0][0], dp[0][1], dp[1][0], dp[1][1] = 0, slices[0], 0, max(slices[0], slices[1])
	for i := 2; i < N; i++ {
		// 任意个候选数字, 一个都不选的最大就是0
		dp[i][0] = 0
		// 迭代所有选取的可能.. but, 实际可选取的数量一定不会超过i的1/3吧...?
		// 但是, 这次选取的可以是前(i+1)个数中的任意一个!
		mx := min(i, n)
		for j := 1; j <= mx; j++ {
			// 比如, 针对第i个数,
			//      如果选了, 就不能选第i-1个数
			//      如果不选, 就可以选第i-1个数
			dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[i])
		}
	}
	return dp[N-1][n]
}

func maxSizeSlices(slices []int) int {
	return max(calculate(slices[1:]), calculate(slices[:len(slices)-1]))
}
