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

func findCheapestPriceDP(n int, flights [][]int, src int, dst int, k int) int {
	// DP 做法来一套

	// 想一下状态转移方程

	// dp[t][i] 表示再t次转机后到达i点的最小消耗
	// dp[t][i] = min(dp[t-1][j] + cost(j, i)) j∈fights

	// 最终结果就是 min(dp[0][dst], ... dp[t][dst])

	// 显而易见的, dp[t]只和dp[t-1]相关, 所以只需要保留前一状态的dp转移后的结果即可

	// 初始情况下, dp[0][...]中, dp[0][src] = 0, 其余的都应该是一个非法值表示不可达

	const INVALID = math.MaxInt32

	var ret = INVALID

	// f表示前态, g表示现态
	var f, g = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = INVALID
		g[i] = INVALID
	}
	f[src] = 0

	for t := 1; t <= k+1; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			g[i] = min(g[i], f[j]+cost)
		}
		ret = min(ret, g[dst])
		f, g = g, f
		for i := range g {
			g[i] = INVALID
		}
	}

	if ret == INVALID {
		return -1
	}

	return ret
}
