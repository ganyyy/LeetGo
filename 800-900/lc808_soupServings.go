package main

func soupServings(n int) float64 {
	// 除25实现归一化(因为都是25的倍数), +24是实现向上取整(尽最大努力交付)
	n = (n + 24) / 25
	if n >= 179 {
		// 提前返回, 是因为 超过这个数量级之后, 概率无限接近于1
		return 1
	}
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}
	var dfs func(int, int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			// 同时倒完, A先倒完的概率是0.5
			return 0.5
		}
		if a <= 0 {
			// A先倒完, 概率为1
			return 1
		}
		if b <= 0 {
			// B先倒完, A先倒完的概率就是0
			return 0
		}
		dv := &dp[a][b]
		if *dv > 0 {
			return *dv
		}
		res := (dfs(a-4, b) + dfs(a-3, b-1) +
			dfs(a-2, b-2) + dfs(a-1, b-3)) / 4
		*dv = res
		return res
	}
	return dfs(n, n)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
