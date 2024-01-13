package main

// dp核心是递推关系
// 如果能满足 新状态 = 旧状态+变换, 可以尝试使用动态规划的方式解决

func findLength(A []int, B []int) int {
	// 转换为 字符串的最长公共子串
	// 动态规划... 怎么就是不长脑子呢
	// dp[i][j]表示 A中第i个结尾和 B中第j个结尾的最长公共长度
	dp := make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, len(B)+1)
	}
	var maxln int
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				t := dp[i][j] + 1
				if t > maxln {
					maxln = t
				}
				dp[i+1][j+1] = t
			}
		}
	}
	return maxln
}

func findLength2(nums1 []int, nums2 []int) int {
	var l1, l2 = len(nums1), len(nums2)

	// dp := make([][]int, l1+1)
	// for i := range dp {
	//     dp[i] = make([]int, l2+1)
	// }

	dp2 := make([]int, l2+1)

	// dp := make

	var ret int
	for i := 1; i <= l1; i++ {
		var leftTop2 int
		for j := 1; j <= l2; j++ {
			var top2 = dp2[j]
			// var left = dp[i][j-1]
			if nums1[i-1] == nums2[j-1] {
				dp2[j] = leftTop2 + 1
				ret = max(ret, dp2[j])
			} else {
				// 这种情况的dp压缩, 得需要考虑重置的逻辑
				dp2[j] = 0
			}
			leftTop2 = top2
		}
	}
	return ret
}
