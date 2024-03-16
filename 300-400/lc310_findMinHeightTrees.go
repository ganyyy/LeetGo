package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
		degree[from]++
		degree[to]++
	}

	var queue []int
	for node, deg := range degree {
		if deg == 1 {
			queue = append(queue, node)
		}
	}

	remainNodes := n
	for remainNodes > 2 {
		// 移除叶子节点
		remainNodes -= len(queue)
		nextQueue := queue
		queue = nil
		for _, node := range nextQueue {
			for _, next := range graph[node] {
				degree[next]--
				if degree[next] == 1 {
					queue = append(queue, next)
				}
			}
		}
	}
	return queue
}
