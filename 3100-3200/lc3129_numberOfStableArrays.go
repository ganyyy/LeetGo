package main

func numberOfStableArrays(zero, one, limit int) (ans int) {
	const mod = 1_000_000_007
	// i个0和j个1组成的稳定数组方案数, 并且最后一位填0/1
	f := make([][][2]int, zero+1)
	for i := range f {
		f[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(limit, zero); i++ {
		// 在一个1都没有的情况下, 最多连续limit个0可以组成有效方案
		f[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		// 在一个0都没有的情况下, 最多连续limit个1可以组成的有效方案
		f[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			// 当前位置填0, 上一个位置可以是0/1, 但是因为一共需要i个0, 所以前置状态就是[i-1]
			f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
			if i > limit {
				// + mod 保证答案非负
				// 因为f[i-1][j][0]包括了后续连续limit个0的情况, 此时就会导致再加上当前位置的0构成了limit+1个0
				// 所以需要保证不存在这种情况.
				// 末尾连续limit个0的方案有多少呢? 等同于 f[i-limit-1][j]这个位置上填的是1的有效个数!
				// 为啥呢? 因为i+j的个数是确定的! 如果j是确定的话, 那么i的数量也是确定, 只有后续是连续limit+1个0的情况下
				// 才能由 f[i-limit-1][j][1] 转移到 f[i][j][0]
				f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
			}
			// 当前位置填1, 上一个位置可以是0/1, 但是因为一共需要j个1, 所以前置状态就是[j-1]
			f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
			if j > limit {
				// 同理
				f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
			}
		}
	}
	return (f[zero][one][0] + f[zero][one][1]) % mod
}

func numberOfStableArrays2(zero int, one int, limit int) int {
	var dp = make([][][2]int, zero+1)
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}
	// 初始化连续0/连续1的情况
	for i := 1; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 1; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}

	const mod = 1_000_000_007

	safeAdd := func(a, b int) int {
		return (a + b) % mod
	}

	// 处理中间情况
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			dp[i][j][0] = safeAdd(dp[i-1][j][0], dp[i-1][j][1])
			if i > limit {
				// 存在末尾连续limit+1个0的可能
				dp[i][j][0] = safeAdd(dp[i][j][0], mod-dp[i-limit-1][j][1])
			}
			dp[i][j][1] = safeAdd(dp[i][j-1][0], dp[i][j-1][1])
			if j > limit {
				dp[i][j][1] = safeAdd(dp[i][j][1], mod-dp[i][j-limit-1][0])
			}
		}
	}
	return safeAdd(dp[zero][one][0], dp[zero][one][1])
}
