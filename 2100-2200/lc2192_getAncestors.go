package main

import "slices"

func getAncestors(n int, edges [][]int) [][]int {
	// 每个节点的入度
	indegree := make([]int, n)

	// 每个节点的子节点
	to := make([][]int, n)
	for _, e := range edges {
		indegree[e[1]]++
		to[e[0]] = append(to[e[0]], e[1])
	}

	// BFS, 广度迭代
	q := make([]int, 0, n)

	// 收集所有入度为0的根节点
	for i, v := range indegree {
		if v == 0 {
			q = append(q, i)
		}
	}
	res := make([][]int, n)
	for ; len(q) > 0; q = q[1:] {
		p := q[0]
		for _, t := range to[p] {
			// p 是 t 的父节点
			res[t] = append(res[t], p)
			// p 的 父节点 也是 t的父节点
			// 得益于先后关系, res[p]肯定已经求出来了
			res[t] = append(res[t], res[p]...)
			// 先排序, 再去重
			slices.Sort(res[t])
			res[t] = slices.Compact(res[t])

			// t的入度--, 如果入度减到0, 就加入到队列中
			indegree[t]--
			if indegree[t] == 0 {
				q = append(q, t)
			}
		}
	}
	for i, r := range res {
		slices.Sort(r)
		// 合并重复的数字..?
		res[i] = slices.Compact(r)
	}
	return res
}
