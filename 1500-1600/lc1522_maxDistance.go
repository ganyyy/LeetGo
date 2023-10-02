package main

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)

	check := func(mp int) bool {
		var pre = position[0]
		var cnt = 1
		for _, cur := range position[1:] {
			if cur-pre >= mp {
				cnt++
				pre = cur
			}
		}
		return cnt >= m
	}

	// 核心区别在于: 满足条件后是往左边缩进还是往右边缩进
	// 往左边的是最小极大值, 往右边的是最大极小值

	// 极大最小值的区别在于: 满足条件后, left=mid+1....
	// 极小最大值是满足条件后, right=mid-1
	l, r := 1, position[len(position)-1]-position[0]
	var ret = -1
	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
