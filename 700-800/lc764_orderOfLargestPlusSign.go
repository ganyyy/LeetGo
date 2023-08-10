package main

func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}

	// 一个二维的前缀和
	banned := map[int]bool{}
	for _, p := range mines {
		banned[p[0]*n+p[1]] = true
	}
	var count int
	check := func(i, j int) int {
		if banned[i*n+j] {
			count = 0
		} else {
			count++
		}
		dp[i][j] = min(dp[i][j], count)
		return dp[i][j]
	}

	// 记录四个维度的最长连续空间
	for i := 0; i < n; i++ {
		count = 0
		/* left */
		for j := 0; j < n; j++ {
			check(i, j)
		}
		count = 0
		/* right */
		for j := n - 1; j >= 0; j-- {
			check(i, j)
		}
	}
	for i := 0; i < n; i++ {
		count = 0
		/* up */
		for j := 0; j < n; j++ {
			check(j, i)
		}
		count = 0
		/* down */
		for j := n - 1; j >= 0; j-- {
			ans = max(ans, check(j, i))
		}
	}
	return ans
}
