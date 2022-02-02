package main

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// i->j, j->i
	var next = make([][]int, n)
	for _, e := range edges {
		next[e[0]-1] = append(next[e[0]-1], e[1]-1)
		next[e[1]-1] = append(next[e[1]-1], e[0]-1)
	}
	// 经过的位置计数
	// 最短路径是每个节点最多只访问一次,
	// 那么次短路径每个节点最多访问两次
	var visited = make([]int, n)
	for i := range visited {
		visited[i] = 2
	}
	// 到达节点i最新轮所需要的耗时
	var log = make([]int, n)
	for i := range log {
		log[i] = -1
	}

	// BFS
	var queue = make([]int, 0, n)
	queue = append(queue, 0)

	var t int
	var first = -1 // 首次到达的时间
	for len(queue) != 0 {
		var idx = len(queue)

		if (t/change)%2 == 1 {
			// 红灯期间, 需要等待红灯结束再前往下一个节点
			t += change - (t % change)
		}
		// 本次耗时
		t += time

		for i := 0; i < idx; i++ {
			var p = queue[i]
			var pn = next[p]
			for _, step := range pn {
				// 每一轮, 耗时都是递增的, 但是在同层次中, 下一轮的目标节点是可以重复选取的
				// 此时 t == log[step], 没有必要多次迭代, 所以可以直接过滤掉
				// 同样的, 多次添加该节点不会改变总的耗时, 所以没必要二次添加
				if t == log[step] || visited[step] == 0 {
					// 过滤掉消耗时间更短的点位.
					// 过滤掉剩余步数为0的点位
					continue
				}
				// 更新统计
				visited[step]--
				log[step] = t
				// 加入到备选节点
				queue = append(queue, step)

				if step == n-1 {
					if first == -1 {
						// 如果是首次到达, 就更新一下时间
						first = t
						continue
					} else {
						// 第二次到达, 一定是次短的时间
						return t
					}
				}
			}
		}
		queue = queue[idx:]
	}

	return 0
}

func secondMinimumNew(n int, edges [][]int, time int, change int) int {
	// 构建映射关系
	var graph = make([][]int, n)
	for _, e := range edges {
		var s, e = e[0] - 1, e[1] - 1
		graph[s] = append(graph[s], e)
		graph[e] = append(graph[e], s)
	}

	// 记录节点的最新访问时间
	var log = make([]int, n)
	// 记录每个节点的最多的访问次数
	// 每个节点最多访问一次: 最短路径, 每个节点最多访问两次: 次短路径
	var visited = make([]int, n)

	for i := range log {
		log[i] = -1
		visited[i] = 2
	}

	var queue = make([]int, 0, n)
	queue = append(queue, 0)

	var t int      // 整个路程的总耗时
	var first = -1 // 首次到达终点的耗时

	for len(queue) > 0 {
		var ln = len(queue)

		// 计算当前轮所需要的耗时
		if (t/change)%2 == 1 {
			t += change - (t % change) // 等红灯的耗时
		}
		t += time // 前进的耗时

		// 迭代当前轮次所有的目标节点

		for i := 0; i < ln; i++ {
			var cur = queue[i]
			var next = graph[cur]

			for _, step := range next {
				if t == log[step] || visited[step] == 0 {
					// 每一轮只迭代一次
					// 只迭代不超过两次的节点
					continue
				}

				// 更新统计信息
				visited[step]--
				log[step] = t
				// 加入到BFS
				queue = append(queue, step)

				// 终点判定
				if step != n-1 {
					continue
				}
				if first == -1 {
					first = t
				} else {
					return t
				}
			}

		}

		queue = queue[ln:]
	}

	return 0
}
