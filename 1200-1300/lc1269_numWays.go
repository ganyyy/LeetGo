package main

func numWays(steps int, arrLen int) int {
	// 猜测, dfs+枝减吧..

	// 回到0,
	//

	// 整体看成 +1, -1 +0 三种操作后, 在不大于 arrLen的范围的steps 步数内回到0

	// 这是一个排列问题...

	// dp 正义执行

	// 最多500步, 判断取值范围在 [1, 500]
	// 0, 501充当哨兵的作用
	var dp [502]int
	var tmp int
	dp[0] = 1
	// 外围是步数
	for i := 1; i <= steps; i++ {
		// 内围是长度
		tmp = 0
		// 必须保证剩余的步数可以回去
		//
		for j := 0; j < min(min(arrLen, i+1), steps-i+1); j++ {
			// tmp代表的是由后向前+1步
			// dp[j]代表的是原地不同+0步
			// dp[j+1]代表的是由前向后-1步

			// 这里替换之后, tmp指向了当前, 可以理解为对应j+1而言, tmp就相当于j
			dp[j], tmp = tmp, dp[j]
			dp[j] = mod(tmp + dp[j] + dp[j+1])
		}
	}
	return dp[0]
}

func mod(a int) int {
	const mod = 1000000007
	return a % mod
}
