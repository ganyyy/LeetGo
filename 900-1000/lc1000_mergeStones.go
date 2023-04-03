package main

import "math"

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) > 0 { // 无法合并成一堆
		return -1
	}

	s := make([]int, n+1)
	for i, x := range stones {
		s[i+1] = s[i] + x // 前缀和
	}

	f := make([][]int, n)
	// 区间DP
	for i := n - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			f[i][j] = math.MaxInt
			//
			for m := i; m < j; m += k - 1 {
				f[i][j] = min(f[i][j], f[i][m]+f[m+1][j])
			}
			if (j-i)%(k-1) == 0 { // 可以合并成一堆
				f[i][j] += s[j+1] - s[i]
			}
		}
	}
	return f[0][n-1]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
