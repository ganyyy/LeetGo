package main

import (
	"math"
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// Kruskal算法: 选边. 按照边的权重进行排序, 从小到大依次遍历, 如果两个端点不在一个连通图中就联通一下, 否则跳过

	var ln = len(edges)

	// 首先给每一条边编一下号
	for i := range edges {
		edges[i] = append(edges[i], i)
	}

	// 所有边根据边权重排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	var p = make([]int, n)
	var find func(i int) int
	find = func(i int) int {
		if i != p[i] {
			p[i] = find(p[i])
		}
		return p[i]
	}
	// 排除边获取最小权重值
	var work1 = func(exclude int) int {
		for i := 0; i < n; i++ {
			p[i] = i
		}
		// 统计边的数量 和最小路径
		var cur, cost int
		for _, edge := range edges {
			if edge[3] == exclude {
				continue
			}
			// 尝试合并a, b两个节点.
			var a, b = edge[0], edge[1]
			if fa, fb := find(a), find(b); fa != fb {
				// 如果成功的话, 就增加一条边并更新权重
				cur++
				cost += edge[2]
				p[fa] = fb
				if cur+1 == n {
					break
				}
			}
		}
		if cur != n-1 {
			return math.MaxInt32
		}
		return cost
	}

	// 指定边获取最小权重
	var work2 = func(include int) int {
		for i := 0; i < n; i++ {
			p[i] = i
		}
		// 统计边的数量 和最小路径
		var cur, cost int
		// 先把边加进去
		for _, edge := range edges {
			if edge[3] == include {
				cur++
				cost += edge[2]
				p[edge[0]] = edge[1]
				break
			}
		}

		for _, edge := range edges {
			// 尝试合并a, b两个节点.
			var a, b = edge[0], edge[1]
			if fa, fb := find(a), find(b); fa != fb {
				// 如果成功的话, 就增加一条边并更新权重
				cur++
				cost += edge[2]
				p[fa] = fb
				if cur+1 == n {
					break
				}
			}
		}
		if cur != n-1 {
			return math.MaxInt32
		}
		return cost
	}

	// 此时cost 就是最小路径
	var cost = work1(-1)

	var res = make([][]int, 2)
	for i := 0; i < ln; i++ {
		if work1(i) > cost {
			// 关键边, 排除后最短路径增加了
			res[0] = append(res[0], i)
		} else if work2(i) == cost {
			// 非关键边, 选择后最短路径不变
			res[1] = append(res[1], i)
		}
	}

	return res
}
