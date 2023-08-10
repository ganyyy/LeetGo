package main

import "math"

func strangePrinter(s string) int {
	// DP啊, 想不出来状态转移方程都白搭

	// 0 <= i, j <= len(s)-1
	// DP[i][j]表示打印s[i:j+1]所需要的最小次数
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		// s[i:i+1]必须要通过一次打印
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// 如果s[i] == s[j], 那么就可以理解为 s[j]是通过s[i]一次打印过来的
				// 此时需要的总的打印次数就是 dp[i][j-1]所需要的次数
				f[i][j] = f[i][j-1]
			} else {
				// 此时可以理解为需要将s[i:j+1]分为两部分进行打印
				// 取一个中间值K, 需要满足s[i:k+1], s[k+1:j+1]一次打印
				// 通过遍历的形式获取满足最小值的k
				f[i][j] = math.MaxInt64
				for k := i; k < j; k++ {
					// 因为 i <= k < j, 所以为了满足计算时k对应的值已计算
					// 需要后序遍历i, 前序遍历j
					f[i][j] = min(f[i][j], f[i][k]+f[k+1][j])
				}
			}
		}
	}
	return f[0][n-1]
}
