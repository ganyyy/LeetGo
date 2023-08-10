package main

import (
	"container/heap"
	"math"
)

func reachableNodes(edges [][]int, maxMoves, n int) (ans int) {
	g := make([][]neighbor, n)
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		// cnt是增加的点的数量
		// cnt+1对应的就是对应的路径长(带权路径边)
		g[u] = append(g[u], neighbor{v, cnt + 1})
		g[v] = append(g[v], neighbor{u, cnt + 1}) // 建图
	}

	dist := dijkstra(g, 0) // 从 0 出发的最短路

	for _, d := range dist {
		if d <= maxMoves { // 这个点可以在 maxMoves 步内到达
			ans++
		}
	}
	for _, e := range edges {
		// start->e[0] 还剩余的步数
		// start->e[1] 还剩余的步数
		a := max(maxMoves-dist[e[0]], 0)
		b := max(maxMoves-dist[e[1]], 0)
		// 下限是多余出来的点的数量
		// 上限是中间增加的点的数量
		ans += min(a+b, e[2]) // 这条边上可以到达的节点数
	}
	return
}

// 以下为 Dijkstra 算法模板
type neighbor struct{ to, weight int }

// g: [0...n-1] 每个节点到起点的联通边
// start: 起点
// 返回值: 从start到[0...n-1]的最短路径
func dijkstra(g [][]neighbor, start int) []int {
	dist := make([]int, len(g))
	// 初始化各个点位距离起点的起始距离
	for i := range dist {
		dist[i] = math.MaxInt32 / 2
	}
	// 起点到起点的最短距离肯定是0
	dist[start] = 0
	// 起点入队
	h := hp{{start, 0}}
	for len(h) > 0 {
		// 获取一个距离起点最近的点
		p := heap.Pop(&h).(pair)
		x := p.x
		// 如果 start->x < start-> x (?)
		// 只有初始化时会这样(!)
		// 相当于同时承担了记录距离和是迭代标记
		if dist[x] < p.dist {
			continue
		}
		// 根据x点到其他点的距离, 更新其他点到起点的最短距离
		for _, e := range g[x] {
			y := e.to
			// start->x + x->y  < start->y
			// 更新 start->y 的距离
			// 将 {y, start->y}入队
			if d := dist[x] + e.weight; d < dist[y] {
				dist[y] = d
				heap.Push(&h, pair{y, d})
			}
		}
	}
	return dist
}

type pair struct{ x, dist int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dist < h[j].dist }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
