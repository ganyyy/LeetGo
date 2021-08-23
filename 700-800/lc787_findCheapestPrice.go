package main

import "math"

func findCheapestPriceOvertime(n int, flights [][]int, src int, dst int, k int) int {
	// DFS

	var tmp = make([][][]int, n)
	for _, f := range flights {
		tmp[f[0]] = append(tmp[f[0]], f[1:])
	}

	// fmt.Println(tmp)
	var dfs func(i int, step int)

	var ret, cur int
	ret = math.MaxInt32
	var visit = make([]bool, n)
	dfs = func(i int, step int) {
		// fmt.Println(i, step)
		if i == dst {
			if cur < ret {
				ret = cur
			}
			return
		}
		if step > k {
			return
		}
		visit[i] = true
		for _, f := range tmp[i] {
			if visit[f[0]] {
				continue
			}
			visit[f[0]] = true
			cur += f[1]
			dfs(f[0], step+1)
			cur -= f[1]
			visit[f[0]] = false
		}
		visit[i] = false
	}

	dfs(src, 0)

	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var dp = make([]int, n)
	var tmp = make([]int, n)

	for i := range dp {
		dp[i] = math.MaxInt32
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 起点的消耗为0
	dp[src] = 0

	// 依赖于k, 确定每一步的最小消耗
	for i := 0; i <= k; i++ {
		// tmp保存dp的前态
		copy(tmp, dp)
		for _, f := range flights {
			// dp[i] 表示的是第k步到i点的最小消耗
			dp[f[1]] = min(dp[f[1]], tmp[f[0]]+f[2])
		}
	}
	if dp[dst] != math.MaxInt32 {
		return dp[dst]
	}
	return -1
}
