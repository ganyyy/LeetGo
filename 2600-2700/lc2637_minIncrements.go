package main

func minIncrements(_ int, cost []int) int {
	// 先找到最长的路径?

	var add int

	for i := len(cost) - 1; i > 0; i -= 2 {
		x := max(cost[i], cost[i-1])
		add += abs(cost[i] - cost[i-1]) // 差值, 将当前层的两个子节点设置为相同大小
		cost[(i-1)/2] += x              // 父节点累加上这个值
	}
	return add
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
