package main

func networkBecomesIdle(edges [][]int, patience []int) int {
	// 最短路径
	// Djakarta
	// 需要确认每个点到原点的最短路径
	// 然后计算每个点的最小公倍数中的

	var relations = make([][]int, len(patience))
	// 统计点之间的关系
	for _, edge := range edges {
		var a, b = edge[0], edge[1]
		relations[a] = append(relations[a], b)
		relations[b] = append(relations[b], a)
	}

	// 起始距离
	var dist = make([]int, len(patience))
	for i := range dist {
		dist[i] = -1
	}

	dist[0] = 0

	var queue = []int{0}

	for len(queue) != 0 {
		var from = queue[0]
		queue = queue[1:]
		for _, p := range relations[from] {
			if dist[p] != -1 {
				continue
			}
			dist[p] = dist[from] + 1
			queue = append(queue, p)
		}
	}

	var ret int
	for i := 1; i < len(patience); i++ {
		var length = dist[i] * 2 // 往返时间
		var cost = length + 1    // 加1s是检测的时间
		var p = patience[i]
		// 因为整体是流水线形式的, 想要算出总共的时长,
		// 等同于首个包抵达后, 最近发出的一个包距离终点的距离
		// 需要计算当首个包返回时, 最新的包还需要多久才可以返回
		if length >= p {
			if length%p == 0 {
				cost += length - p // 因为可以整除, 所以相当于少发了一个包
			} else {
				cost += length - (length % p) //
			}
		}
		ret = max(ret, cost)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
