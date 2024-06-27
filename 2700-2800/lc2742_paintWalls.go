package main

import (
	"math"
)

func paintWalls(cost, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2 // 防止加法溢出
	}
	// 付费刷墙时间之和 >= 免费刷墙的个数
	// 付费刷墙个数+免费刷墙个数 = n
	// 付费刷墙时间之和 >= n-付费刷墙个数
	// 付费刷墙时间之和 +  付费刷墙个数 >= n
	// 等同于在每一个位置上, sum(付费刷墙时间+1) >= n
	// 将每一面墙的时间+1看成体积, 然后通过这个体积选取n个物品对应的最小价值
	// f[i]: i件物品经过各种组合之后的最低开销
	// 因为后态只和前态相关, 所以可以状态压缩

	// cost[i] 是所有可选的物品的价值
	for i, c := range cost {
		// t 是这个物品对应的体积
		// 简单而言, 就是找出最大的体积+最小的价值, 这样就可以填入更多的免费刷墙的数量, 保证最终开销最低
		t := time[i] + 1 // 注意这里加一了
		for j := n; j > 0; j-- {
			// 为啥最终要到0呢?因为最少需要选一面墙付费刷, 当体积不足时, 那也得刷一下
			f[j] = min(f[j], f[max(j-t, 0)]+c)
		}
		// fmt.Println(f)
	}
	return f[n]
}

func paintWalls2(cost, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 没有计算过
		}
	}
	// i: 剩余的墙的数量
	// j: 付费的墙的时间之和-免费的墙的数量(也是免费墙的时间之和)
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j > i { // 剩余的墙都可以免费刷
			return 0
		}
		if i < 0 { // 上面 if 不成立，意味着 j < 0，不符合题目要求
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j+n] // 加上偏移量 n，防止出现负数
		if *p != -1 {      // 之前计算过
			return *p
		}
		// dfs(i-1, j+time[i])+cost[i]: 付费刷这一面墙
		// dfs(i-1, j-1): 免费刷这一面墙
		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		return *p
	}
	return dfs(n-1, 0)
}
