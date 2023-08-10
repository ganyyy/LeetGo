package main

import "math"

func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
	type edge struct{ to, eid int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, i})
		g[y] = append(g[y], edge{x, i}) // 建图，额外记录边的编号
	}

	// dijkstra: 从起点到任意地点的最短路径长度
	var delta int
	dis := make([][2]int, n)
	for i := range dis {
		dis[i] = [2]int{math.MaxInt, math.MaxInt}
	}
	dis[source] = [2]int{}
	dijkstra := func(k int) { // 这里 k 表示第一次/第二次
		// 最短路径啊
		vis := make([]bool, n)
		for {
			// 找到当前最短路，去更新它的邻居的最短路
			// 根据数学归纳法，dis[x][k] 一定是最短路长度
			x := -1
			for y, b := range vis {
				// 从没有访问过的路径中查找距离已访问节点最近的点
				if !b && (x < 0 || dis[y][k] < dis[x][k]) {
					x = y
				}
			}
			if x == destination { // 起点 source 到终点 destination 的最短路已确定
				return
			}
			vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
			for _, e := range g[x] {
				// 基于当前选中的点x, 更新其后续节点相关的路径权重
				y, wt := e.to, edges[e.eid][2]
				if vis[y] {
					// 已经迭代过的节点, 无需处理
					continue // ?
				}
				// 负权重节点, 首先尝试更新成1
				if wt == -1 {
					wt = 1 // -1 改成 1
				}
				if k == 1 && edges[e.eid][2] == -1 {
					// 第二次 Dijkstra，改成 w
					// 首先:
					// 路径分为 source - x - y - destination 三部分
					// source - x       : 已确定(本次迭代筛选出来的最短路径点)
					// x - y            : w
					// y - destination  : 后续迭代
					// 其中:
					// source - x   = dis[x][1]
					// destination-y= dis[destination][0]-dis[y][0], 第一轮dijkstra计算出的结果
					// 那么:
					// 如果 y 在 最短路径上, 并且需要保证最短路径为 target, 那么
					// w + dis[x][1] + (dis[destination][0]-dis[y][0]) = target
					// w = target - dis[destination][0] + dis[y][0] - dis[x][1]
					// target - dis[destination][0] 这个答案可以在第一轮dijkstra之后计算出来(delta)
					// 结论:
					// 为啥要直接怼满呢...? 是因为第一轮计算出来的结果中, 已经默认所有的负权重节点变成1啦
					// 这样一来, 是可能影响到后续的最短路径的, 因为第一轮计算出来的 dis[destination][0] - dis[y][0] 中途径的节点里可能包含了负权重节点
					// 不过这个不影响, 因为这里就是通过贪心的方式, 保证了恰好一次修改后, 最短路径就是 target, 即使最短路径上不包含y
					// 所有路径的最小值都是1,
					w := delta + dis[y][0] - dis[x][1]
					if w > wt {
						wt = w
						edges[e.eid][2] = w // 直接在 edges 上修改
					}
				}
				// 更新最短路
				// 需要注意的是: 负权重节点不一定在最短路径上欸(?)
				dis[y][k] = min(dis[y][k], dis[x][k]+wt)
			}
		}
	}

	dijkstra(0)
	delta = target - dis[destination][0]
	if delta < 0 { // -1 全改为 1 时，最短路比 target 还大. 1是负权重节点所能变化的最小值, 不能再小了
		return nil
	}

	dijkstra(1)
	if dis[destination][1] < target { // 最短路无法再变大，无法达到 target
		return nil
	}

	for _, e := range edges {
		if e[2] == -1 { // 剩余没修改的边全部改成 1
			e[2] = 1
		}
	}
	return edges
}
