package main

import "math"

func getMoneyAmount(n int) int {

	// f[i][j]表示从[i,j]之间确保胜利的最低金额
	// 对于猜的数字是x的情况,
	// 如果x猜大了, 就需要取f[i][x-1]; 猜小了就需要取f[x+1][j]
	// 想要胜利的前提, 就需要取二者中的最大值
	// f[i][j] = min(x + max(f[i][x-1], f[x+1][j])), x∈[i,j]
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	// 计算f[i][j], 需要用到 >i, <j. 所以要逆序迭代i, 正序迭代j
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				// 如果 k-1 <= i, 那么肯定是0, 不存在这样的区间
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}
	return f[1][n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func getMoneyAmount2(n int) int {
	var cache = make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}

	var dfs func(i, j int) int
	dfs = func(start, end int) int {
		if start >= end {
			return 0
		}
		if cache[start][end] != 0 {
			return cache[start][end]
		}
		var ans = math.MaxInt32
		for k := start; k <= end; k++ {
			var cur = max(dfs(start, k-1), dfs(k+1, end)) + k
			if cur < ans {
				ans = cur
			}
		}
		cache[start][end] = ans
		return ans
	}

	return dfs(1, n)
}

func main() {
	getMoneyAmount(10)
}
