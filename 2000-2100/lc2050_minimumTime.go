//go:build ignore

package main

func minimumTime(n int, relations [][]int, time []int) int {
	// 每个课程的入度
	var deps = make([]int, n)
	// 每个课程的后继课程
	var forward = make([][]int, n)
	// 每个课程对应的耗时
	var cost = make([]int, n)

	// 通过关系整理课程的后继和入度
	for _, relation := range relations {
		pre, next := relation[0]-1, relation[1]-1
		forward[pre] = append(forward[pre], next)
		deps[next]++
	}

	var ret int

	// 首先, 找出所有的起点
	var cur, next []int
	for course, dep := range deps {
		if dep == 0 {
			// 没有前置课程
			cur = append(cur, course)
			// 当前课程的消耗时间
			cost[course] = time[course]
			ret = max(ret, time[course])
		}
	}

	// 然后, 迭代所有课程, 并更新其后继节点
	for len(cur) != 0 {
		for _, course := range cur {
			cc := cost[course]
			for _, nc := range forward[course] {
				// 入度-1
				deps[nc]--
				// 更新学完nc所需要的最大时间, 总时间是所有前驱课程中的最大值
				// 为啥不能预处理呢? 因为前驱课程的前驱课程是不确定的, 只能通过迭代的方式一步一步的递增
				cost[nc] = max(cost[nc], cc+time[nc])
				if deps[nc] == 0 {
					// nc的所有前置都已经完成, 将nc加入到待学习队列中
					ret = max(ret, cost[nc])
					next = append(next, nc)
				}
			}
		}
		cur, next = next, cur[:0]
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
