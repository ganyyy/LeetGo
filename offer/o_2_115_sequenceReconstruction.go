package main

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	var n = len(nums)
	var graph = make([][]int, n+1)
	var depth = make([]int, n+1) // 入度

	for _, s := range sequences {
		for i := 1; i < len(s); i++ {
			graph[s[i-1]] = append(graph[s[i-1]], s[i])
			depth[s[i]]++
		}
	}

	var queue []int
	for i := 1; i <= n; i++ {
		if depth[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 必须要保证拓扑排序的结果是唯一的, 否则就是不满足需求

	for len(queue) != 0 {
		if len(queue) > 1 {
			return false
		}
		head := queue[0]
		queue = queue[1:]
		for _, p := range graph[head] {
			depth[p]--
			if depth[p] == 0 {
				queue = append(queue, p)
			}
		}
	}

	return true
}
