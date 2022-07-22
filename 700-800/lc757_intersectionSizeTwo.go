package main

import "sort"

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		var a, b = intervals[i], intervals[j]
		// 右边界升序, 左边界降序
		if a[1] != b[1] {
			return a[1] < b[1]
		}
		return a[0] > b[0]
	})

	// a/b代表当前集合中的最大值/次大值
	var a, b = -1, -1
	var cnt int
	for _, v := range intervals {
		var l, r = v[0], v[1]

		if l > b {
			// 没有任何重合的点
			// 贪心选择当前最大的两个点
			a, b = r-1, r
			cnt += 2
		} else if l > a {
			// 隐含的条件是: a < b <= l
			// b就是那个重合点(?) 此时就需要选择 b 作为次大值, r最为最大值
			// 这样做的好处就是, 只会选取一个额外的点
			a, b = b, r
			cnt++
		}

	}
	return cnt
}
