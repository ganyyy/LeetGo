package main

func numBusesToDestination(routes [][]int, source int, target int) int {
	// 还是BFS啊啊啊啊

	// 需要一个映射, 保留了每个公交车可以去的点

	if source == target {
		return 0
	}

	var mm = make(map[int][]int)

	for i, route := range routes {
		// 站点对应的公交路线
		for _, r := range route {
			mm[r] = append(mm[r], i)
		}
	}

	var step int

	var queue1, queue2 []int // 待探索的路线
	var visited = make(map[int]bool, len(mm))

	// fmt.Println(mm)

	queue1 = mm[source]

	for len(queue1) != 0 {

		for _, p := range queue1 {
			if visited[p] {
				continue
			}
			visited[p] = true
			for _, v := range routes[p] {
				if v == target {
					return step + 1
				} else {
					for _, pr := range mm[v] {
						if visited[pr] {
							continue
						}
						queue2 = append(queue2, pr)
					}
				}
			}
		}

		step++
		queue1, queue2 = queue2, queue1
		queue2 = queue2[:0]
	}

	return -1
}
