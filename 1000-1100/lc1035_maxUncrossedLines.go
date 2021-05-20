package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {

	// 感觉可以搞成DP...
	// dp[i][j] 表示nums1[1:i+1]和nums2[1:j+1]之间组成的最多线段数
	var dp = make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	// dp可以压缩欸
	// 数据也可以压缩
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				// 在原来的基础上加一条
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// 选nums1的, 或者选nums2的
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[len(nums1)][len(nums2)]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
