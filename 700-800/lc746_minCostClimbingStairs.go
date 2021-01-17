package main

func minCostClimbingStairs(cost []int) int {
	// dp. a 表示当前点的前2个位置的消耗, b表示当前点的前1个位置的消耗

	// 初始值是第0个点, 第1个点
	var a, b = cost[0], cost[1]

	// 一次变换指的是 将a变成b, 将b变成到当前点的最低消耗.
	// 可由原来的a, b跳到当前点, 所以跳到当前点的最低消耗就是 min(a, b) + cost[i]
	for i := 2; i < len(cost); i++ {
		a, b = b, min(a, b)+cost[i]
	}

	return min(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
