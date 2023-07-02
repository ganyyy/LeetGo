//go:build ignore

package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func connectTwoGroups(cost [][]int) int {
	// 数组维护的是任意两点之间的连接成本. 全相联映射
	size1, size2, m := len(cost), len(cost[0]), 1<<len(cost[0])
	// 一维是 size1, 二维是 m,
	// dp[i][s]表示前i个点位和size2点集s的最小联通成本
	// dp[size1][m-1]就是最终的结果
	dp := make([][]int, size1+1)
	for i := 0; i <= size1; i++ {
		dp[i] = make([]int, m)
	}
	// 初始状态:
	// dp[0][X], dp[X][0]: 都是无解的. 因为无法联通
	for s := 1; s < m; s++ {
		dp[0][s] = 0x3f3f3f3f
	}

	for i := 1; i <= size1; i++ {
		for s := 0; s < m; s++ {
			dp[i][s] = 0x3f3f3f3f // 假定当前无解
			for k := 0; k < size2; k++ {
				// 从头开始迭代s中包含的点位
				// s是一个二进制数, 从低位到高位, 依次判断是否包含k
				// k是size2中的点位索引, 从0开始
				if (s & (1 << k)) == 0 {
					continue
				}

				// 这种DP, 怎么可能做的出来...

				// k锁定(i-1). 因为k之前可能已经被锁定了, 所以需要异或消除之前k带来的连接的开销
				// dp[i][s^(k-1)]表示的是包含(i-1)这个点时, 点集为(s^(1<<k))的最小成本
				dp[i][s] = min(dp[i][s], dp[i][s^(1<<k)]+cost[i-1][k])
				// (i-1)锁定k. 需要在保持s的基础上, 将(i-1)连接到k上.
				// dp[i-1][s]表示的是不包含(i-1)这个点时, 点集为s的最小成本
				dp[i][s] = min(dp[i][s], dp[i-1][s]+cost[i-1][k])
				// (i-1)&k 相互锁定. 采用不包含(i-1)和k的子集, 将(i-1)和k相互连接.
				// dp[i-1][s^(k-1)]表示的是不包含(i-1)这个点时, 点集为(s^(1<<k))的最小成本
				dp[i][s] = min(dp[i][s], dp[i-1][s^(1<<k)]+cost[i-1][k])
			}
		}
	}
	// 感觉, 有些类似于状态转移方程如 dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + cost[i][j] 这种形式耶
	// 这里也是分别选取了三种情况, 然后取最小值. 包括了选i&(s^(k-1)), 选(i-1)&s, 选(i-1)&(s^(k-1))这三种情况
	return dp[size1][m-1]
}
