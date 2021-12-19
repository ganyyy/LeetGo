//go:build ignore
// +build ignore

package main

import "sort"

func findRadiusBad(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	// 左端点从0开始, 为了处理不需要添加的情况
	var left, right = 0, max(houses[len(houses)-1], heaters[len(heaters)-1])

	var check = func(radio int) bool {
		// 如何快速判断是否可以覆盖?

		// WARNING 这里有很大的优化空间
		var p1, p2 int
		for p1 < len(houses) {
			var house = houses[p1]
			for ht := heaters[p2]; ht-radio > house || ht+radio < house; {
				p2++
				if p2 >= len(heaters) {
					return false
				}
				ht = heaters[p2]
			}
			p1++
		}
		return p2 < len(heaters)
	}

	for left <= right {
		var mid = left + (right-left)>>1

		if check(mid) {
			// 满足条件, 半径大了, 向左走
			right = mid - 1
		} else {
			// 不满足条件, 半径小了, 向右走
			left = mid + 1
		}
	}
	return left
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	var ret int
	var p2 int
	var ln = len(heaters)
	for _, house := range houses {
		// 查找到第一个大于当前房间的安置位置
		for p2 < ln && heaters[p2] < house {
			p2++
		}
		if p2 == 0 {
			// 开头
			ret = max(ret, heaters[p2]-house)
		} else if p2 == ln {
			// 结尾
			ret = max(ret, house-heaters[ln-1])
		} else {
			// 中间
			ret = max(ret, min(house-heaters[p2-1], heaters[p2]-house))
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
