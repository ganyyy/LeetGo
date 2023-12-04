package main

func minimumFuelCost(roads [][]int, seats int) int64 {
	// 建立邻接表(看清题目呀, 节点的个数是路的个数+1)
	var graph = make([][]int, len(roads)+1)
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], road[1])
		graph[road[1]] = append(graph[road[1]], road[0])
	}
	var oil int
	var dfs func(current, parent int) (car, total int)
	dfs = func(current, parent int) (car, total int) {
		for _, nxt := range graph[current] {
			if nxt == parent {
				continue
			}
			sc, st := dfs(nxt, current)
			oil += sc
			total += st
		}
		// 车的数量 (1+total+seats-1)/seats,
		// 总的节点个数 total+1
		return (total + seats) / seats, total + 1
	}

	dfs(0, -1)

	return int64(oil)
}
