package main

func eventualSafeNodesTimeout(graph [][]int) []int {
	// 终点是出度为0的节点

	// 任何点所走的路径中不能有环

	// 需要有一种方法快速的判断是否该节点存在环

	// 超时了, 想想要怎么优化

	var loop = make([]bool, len(graph))
	var visited = make([]bool, len(graph))

	var ret []int
	var dfs func(i int) bool
	dfs = func(i int) bool {
		for _, p := range graph[i] {
			if visited[p] || loop[p] {
				loop[p] = true
				return false
			}
			visited[p] = true
			var ok = dfs(p)
			visited[p] = false
			if !ok {
				loop[p] = true
				return false
			}
		}
		// 到达终点了
		return true
	}

	for i := range graph {
		if loop[i] {
			continue
		}
		visited[i] = true
		if dfs(i) {
			ret = append(ret, i)
		} else {
			loop[i] = true
		}
		visited[i] = false
	}

	return ret
}

func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	// 三色标记
	// 0: 未访问
	// 1: 已访问, 非安全节点
	// 2: 已访问, 安全点
	color := make([]int, n)
	var safe func(int) bool
	safe = func(x int) bool {
		if color[x] > 0 {
			// 已访问的点, 要么是安全的, 要么是不安全的
			return color[x] == 2
		}
		// 否则, 一旦出现不安全的情况, 该点就会标记为不合法
		color[x] = 1
		for _, y := range graph[x] {
			if !safe(y) {
				return false
			}
		}
		// 安全点上的每一个回溯的路径都是安全的
		color[x] = 2
		return true
	}
	for i := 0; i < n; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return
}
