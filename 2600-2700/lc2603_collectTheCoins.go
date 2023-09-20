package main

func collectTheCoins(coins []int, edges [][]int) int {
	var coinCount = len(coins)
	// 每个节点的度
	var degree = make([]int, coinCount)
	// 构建邻接图
	var graph = make([][]int, coinCount)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		degree[from]++
		degree[to]++
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	// 1 找出所有没有金币的叶子节点
	var removeQueue []int
	for id, deg := range degree {
		if deg == 1 && coins[id] == 0 {
			removeQueue = append(removeQueue, id)
		}
	}
	// 剩余的边的数量
	var remainEdge = coinCount - 1

	// 2 递归的删除没有金币的叶子节点(如果一个叶子节点被删除后, 其上级节点也变成了叶子节点且没有金币, 也要删除)

	for len(removeQueue) > 0 {
		node := removeQueue[len(removeQueue)-1]
		removeQueue = removeQueue[:len(removeQueue)-1]
		remainEdge--
		for _, next := range graph[node] {
			degree[next]--
			// 上级节点去掉这条边之后也变成了叶子节点, 并且这个位置上没有硬币
			if degree[next] == 1 && coins[next] == 0 {
				removeQueue = append(removeQueue, next)
			}
		}
	}

	// 3 删除所有有金币的节点和其父节点
	for id, deg := range degree {
		if deg == 1 && coins[id] > 0 {
			removeQueue = append(removeQueue, id)
		}
	}
	remainEdge -= len(removeQueue)
	for _, node := range removeQueue {
		for _, next := range graph[node] {
			degree[next]--
			// 如果带有金币叶子节点的的父节点也变成了叶子节点, 也可以直接删除了
			// 注意: 这里不用管这个父节点是否有金币, 因为收集金币的范围是两格
			if degree[next] == 1 {
				remainEdge--
			}
		}
	}
	// 这个判断是为什么呢? 当所有节点都需要删除时, 最后剩余的两个点之间的边会被删除两次
	if remainEdge > 0 {
		// 剩下还存在的边, 就是需要走的路径数
		return remainEdge << 1
	}
	return 0
}
