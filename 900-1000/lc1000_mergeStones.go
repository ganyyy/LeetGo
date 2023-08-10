package main

import "math"

func mergeStones(stones []int, k int) int {
	n := len(stones)
	// 关于这个的推算:
	// n -> 1: 需要消去 n-1 堆
	// 每次合并会减少 k-1 堆
	// 所以 n-1 必须要是 k-1 的整数倍
	if (n-1)%(k-1) > 0 { // 无法合并成一堆
		return -1
	}

	/*

		dp(i,j,k): 将[i, j]合并成k堆所需要的最小开销

		假设 i: 0, j: 6, k: 3, sum(i,j)表示[i,j]前缀和

		dp(0,6,1): |x x x x x x x|
		=
		dp(0,6,3) + sum(0,6)
		=
		因为每次合并, 都需要3堆石头, 所以第一堆只能是以下三种情况:
		dp(0,0,1) + dp(1,6,2) // |x|x x x x x x| (第一堆 0 次合并)
		dp(0,2,1) + dp(3,6,2) // |x x x|x x x x| (第一堆 1 次合并)
		dp(0,4,1) + dp(5,6,2) // |x x x x x|x x| (第一堆 2 次合并)


		一般化:
		- 递归入口 dp(0, n-1, 1) 合并到只剩1堆
		- 特殊值   dp(i,i,1) = 0 自己和自己无需合并


		- dp(i, j, 1) = dp(i,j,k) + sum(i, j)
		- dp(i, j, p) = min[dp(i, m, 1) + dp(m+1, j, p-1)]
							m = i+(k-1)*X, X∈1, 2, ...
							p >= 2

		- dp(i, j, 1) = [dp(i, m, 1) + dp(m+1, j, k-1)] + sum(i, j)
							m = i+(k-1)*X, X∈1, 2, ...
		- 可以通过 (j-i)%(k-1) == 0 来判断是否可以合并成一批

		- ndp(i, j): 将[i, j]合并到不能合并的最小成本. 可以合并成1堆, 那么最终 ndp[0][n-1]就是最终求解的结果
		- 此时忽略了 p, 是因为 [i,j]合并到不能合并时, p的值是固定的((j−i)%(k−1)+1)!
	*/

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
