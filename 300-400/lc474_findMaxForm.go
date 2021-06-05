package main

func findMaxForm(strs []string, m int, n int) int {
	// var count = func(s string) (ret [2]int) {
	//     for i := range s {
	//         if s[i] == '1' {
	//             ret[1]++
	//         } else {
	//             ret[0]++
	//         }
	//     }
	// }

	// 先整体排个序, 因为数量少的肯定有最多的可能
	// sort.Slice(strs, func(i, j int) bool {
	//     return len(strs[i]) < len(strs[j])
	// })

	// 然后正常的双指针, 尽量求最大的
	// 求和的最大值

	// 好家伙, 背包问题

	// 选or不选

	// 背包的总容量是?

	// 如何把一维扩展到两个维度?

	// 多维背包就是多个维度的匹配
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		var zero, one int
		for i := range s {
			if s[i] == '1' {
				one++
			} else {
				zero++
			}
		}
		//状态转移方程：dp[i][j] = Math.max(dp[i][j],1+dp[i-numZero][j-numOne])
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				if t := dp[i-zero][j-one] + 1; t > dp[i][j] {
					dp[i][j] = t
				}
			}
		}
	}

	return dp[m][n]
}
