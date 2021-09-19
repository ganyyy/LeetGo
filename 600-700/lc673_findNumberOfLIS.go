package main

func findNumberOfLIS(nums []int) (ans int) {
	maxLen := 0
	n := len(nums)
	dp := make([]int, n)  // dp[i] 表示 nums[:i+1]中LIS的长度
	cnt := make([]int, n) // cnt[i] 表示 nums[:i+1]中以i为结尾的LIS的个数
	for i, x := range nums {
		// 基于i, x判断LIS
		// 默认情况下都是1
		dp[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			// 如果x>y, 那么 x, y 就可以组成一个递增序列
			if x > y {
				if dp[j]+1 > dp[i] {
					// dp[j]+1 > dp[i]说明 nums[:i+1]存在较长的递增子序列
					dp[i] = dp[j] + 1 // 更新dp[i]的值
					cnt[i] = cnt[j]   // 重置计数
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j] // 此时意味着nusm[:i+1]中存在不同组合的LIS
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i] // 跟新nums[:i+1]的LIS长度 和 组合计数
			ans = cnt[i]   // 重置LIS组合计数
		} else if dp[i] == maxLen {
			ans += cnt[i] // 没有出现更长的LIS, 直接进行递增
		}
	}
	return
}
