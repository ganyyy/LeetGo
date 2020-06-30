package _00_800

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
