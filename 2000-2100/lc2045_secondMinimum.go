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
	// 记录到达点位i的耗时
	// TODO 记录这个信息的作用还是没搞太明白
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
				// TODO 时间还会变短吗..?
				if t <= log[step] || visited[step] == 0 {
					// 过滤掉时间更短的点位. 如果一个点位多次访问. 时间只会越来越长
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
