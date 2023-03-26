package main

func countSubstrings(s, t string) int {
	m, n := len(s), len(t)

	fill := func() (dp [][]int) {
		dp = make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		return
	}
	var dpl = fill()
	var dpr = fill()

	// dpl[i][j]: s[:i]和t[:j]中连续相等的字符串的长度(左往右)
	// dpr[i][j]: s[i:]和t[j:]中连续相等的字符串的长度(右往左)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				dpl[i+1][j+1] = dpl[i][j] + 1
			} else {
				dpl[i+1][j+1] = 0
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dpr[i][j] = dpr[i+1][j+1] + 1
			} else {
				dpr[i][j] = 0
			}
		}
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 左边连续相同 * 右边连续相同
			if s[i] != t[j] {
				ans += (dpl[i][j] + 1) * (dpr[i+1][j+1] + 1)
			}
		}
	}
	return ans
}
