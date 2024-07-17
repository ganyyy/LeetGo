package main

import "container/heap"

func minimumTime(n int, edges [][]int, disappear []int) []int {
	// Dijkstra算法, 有源最短路径
	var graph = make([][]Item, n)
	for _, edge := range edges {
		from, to, time := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], Item{to, time})
		from, to = to, from
		graph[from] = append(graph[from], Item{to, time})
	}

	// fmt.Println(graph)

	var dist = make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0
	var pq PriorityQueue
	heap.Push(&pq, &Item{0, 0})

	for len(pq) > 0 {
		top := heap.Pop(&pq).(*Item)
		current := top.dist
		u := top.node
		if current > dist[u] {
			continue
		}
		// fmt.Println("current", u)
		for _, next := range graph[u] {
			to := next.node
			// u -> v的时间
			newDist := next.dist + current
			if newDist < disappear[to] && (newDist < dist[to] || dist[to] < 0) {
				// fmt.Println(u, "to", to, "current", newDist, "disappear", disappear[to])
				dist[to] = newDist
				heap.Push(&pq, &Item{to, newDist})
			}
		}
	}
	return dist
}

// PriorityQueue implements a priority queue for Dijkstra's algorithm
type PriorityQueue []*Item

// Item represents an item in the priority queue
type Item struct {
	node, dist int
}

// Len implements heap.Interface.Len
func (pq PriorityQueue) Len() int { return len(pq) }

// Less implements heap.Interface.Less
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }

// Swap implements heap.Interface.Swap
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push implements heap.Interface.Push
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }

// Pop implements heap.Interface.Pop
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
