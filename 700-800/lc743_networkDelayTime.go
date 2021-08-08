package main

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n, k int) (ans int) {

	// n个节点之间的连接耗时为times, 从k出发所需要的最短耗时

	// 迪杰斯特拉算法: 维护一个从起点到其他任意点距离的的数组
	// 点分为已确定最短距离的点和未确定最短距离的点两部分
	// 从未确定的点中取一个距离起点最近的点, 以此来更新和该点存在关系的其他点的距离

	type edge struct{ to, time int }

	// g[x]表示的是以x为起点, 所有可到达的点的集合
	g := make([][]edge, n)
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x] = append(g[x], edge{to: y, time: t[2]})
	}

	// 初始情况下, 每个点都是不可达的
	const inf int = math.MaxInt64 / 2

	//dist[i]表示的是从起点到i+1点的最短距离
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	// 标记起点
	dist[k-1] = 0

	// 构建初始堆
	h := &hp{{k - 1, 0}}

	// 正统的bfs
	for h.Len() > 0 {
		// 从堆中获取最小节点(最短节点)
		p := heap.Pop(h).(pair)
		x := p.x
		// 只需要关注距离较短的点.
		// 因为有向图可能会存在环, 如果一旦发生了多次遍历, 取最小值即可
		if dist[x] < p.d {
			continue
		}
		for _, e := range g[x] {
			// 遍历所有和x相关联的点
			y := e.to
			// 如果从x->y的耗时(x为从起点到x的最短距离)小于当前保存的最短距离, 就更新一下
			if d := dist[x] + e.time; d < dist[y] {
				dist[y] = d
				heap.Push(h, pair{y, d})
			}
		}
	}

	for _, d := range dist {
		// 如果存在不可达的点, 直接返回错误
		if d == inf {
			return -1
		}
		// 最终结果一定是所有点的集合的最大值
		ans = max(ans, d)
	}
	return
}

type pair struct{ x, d int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
