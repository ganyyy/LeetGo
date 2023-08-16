package main

func ways(pizza []string, k int) int {
	m := len(pizza)
	n := len(pizza[0])
	const mod = 1_000_000_007
	// apples[i][j]表示以pizza[i][j]为左上角的矩形种, 苹果的数量
	// 相当于是一个二维的前缀和
	apples := make([][]int, m+1)
	for i := range apples {
		apples[i] = make([]int, n+1)
	}

	// dp[k][i][j]: 以[i,j]作为左上角坐标的矩形, 切成k块pizza的方案数
	// 结果就是dp[k][0][0]
	dp := make([][][]int, k+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	// 预处理:
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			apples[i][j] = apples[i+1][j] + apples[i][j+1] - apples[i+1][j+1]
			if pizza[i][j] == 'A' {
				// 端点是个苹果, 得需要+1
				apples[i][j] += 1
			}
			if apples[i][j] > 0 {
				// 切一刀就能保证至少有一个苹果, 不管横着切还是竖着切
				dp[1][i][j] = 1
			}
		}
	}

	for ki := 2; ki <= k; ki++ {
		// 计算切成k块pizza的方法
		// 不管在哪个方向上, 切一刀的前提是: 大矩形一定要比小矩形拥有更多的苹果数量!
		// 这样才能保证在大矩形多出来的那一行/列切完后, 新出现的那块pizza也至少含有一个苹果!
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				// 水平方向切
				for i2 := i + 1; i2 < m; i2++ {
					if apples[i][j] > apples[i2][j] {
						dp[ki][i][j] = (dp[ki][i][j] + dp[ki-1][i2][j]) % mod
					}
				}
				// 垂直方向切
				for j2 := j + 1; j2 < n; j2++ {
					if apples[i][j] > apples[i][j2] {
						dp[ki][i][j] = (dp[ki][i][j] + dp[ki-1][i][j2]) % mod
					}
				}
			}
		}
	}
	return dp[k][0][0]
}

func ways2(pizza []string, k int) int {
	m := len(pizza)
	n := len(pizza[0])
	const mod = 1_000_000_007
	// apples[i][j]表示以pizza[i][j]为左上角的矩形种, 苹果的数量
	// 相当于是一个二维的前缀和
	apples := make([][]int, m+1)
	for i := range apples {
		apples[i] = make([]int, n+1)
	}

	// dp[i][j]: 以[i,j]作为左上角坐标的矩形, 切成k块pizza的方案数
	// 结果就是dp[k][0][0]
	dp := make([][]int, m+1)
	next := make([][]int, m+1)
	for j := range dp {
		dp[j] = make([]int, n+1)
		next[j] = make([]int, n+1)
	}

	// 预处理:
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			apples[i][j] = apples[i+1][j] + apples[i][j+1] - apples[i+1][j+1]
			if pizza[i][j] == 'A' {
				// 端点是个苹果, 得需要+1
				apples[i][j] += 1
			}
			if apples[i][j] > 0 {
				// 切一刀就能保证至少有一个苹果, 不管横着切还是竖着切
				dp[i][j] = 1
			}
		}
	}

	for ki := 2; ki <= k; ki++ {
		// 计算切成k块pizza的方法
		// 不管在哪个方向上, 切一刀的前提是: 大矩形一定要比小矩形拥有更多的苹果数量!
		// 这样才能保证在大矩形多出来的那一行/列切完后, 新出现的那块pizza也至少含有一个苹果!
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				next[i][j] = 0
				// 水平方向切
				for i2 := i + 1; i2 < m; i2++ {
					if apples[i][j] > apples[i2][j] {
						next[i][j] = (next[i][j] + dp[i2][j]) % mod
					}
				}
				// 垂直方向切
				for j2 := j + 1; j2 < n; j2++ {
					if apples[i][j] > apples[i][j2] {
						next[i][j] = (next[i][j] + dp[i][j2]) % mod
					}
				}
			}
		}
		next, dp = dp, next
	}
	return dp[0][0]
}
