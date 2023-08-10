package main

func minTaps(n int, ranges []int) int {
	rightMost := make([]int, n+1)
	for i := range rightMost {
		rightMost[i] = i
	}
	for i, r := range ranges {
		// start ... end
		// 0  .......  n
		// start 可以到达的最远的位置
		start := max(0, i-r)
		end := min(n, i+r)
		rightMost[start] = max(rightMost[start], end)
	}

	last, ret, pre := 0, 0, 0
	for i := 0; i < n; i++ {
		last = max(last, rightMost[i])
		// 这是为啥呢? 可以理解为: 初始值是这样的. 如果相等了, 就说明上一个的end无法包含当前
		// 这也就意味着无法全部包含
		if i == last {
			return -1
		}
		// 根据贪心原则, 每一步都走最远的距离, 那么总体的步数就会越小(?)
		// 为啥可行呢, 可以这么理解: 在走到 pre 之前, last 可以达到的最远距离 就是下一步的 pre
		//                         因为下一步可以在到达pre之前的任意一步走出来
		if i == pre {
			// 有点类似于买股票(跳箱子?)
			ret++
			pre = last
		}
	}
	return ret
}
