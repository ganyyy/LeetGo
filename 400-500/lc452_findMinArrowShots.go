package main

import "sort"

func findMinArrowShotsLift(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	if len(points) < 2 {
		return 1
	}
	// 需要先拍一下顺序 按照左端点排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// 贪心算法

	var res int = 1
	// 当前能射穿的最远的距离
	var cur = points[0][1]
	for i := 1; i < len(points); i++ {
		// 这一箭射不到这个了
		if points[i][0] > cur {
			res++
			cur = points[i][1]
		} else if points[i][1] < cur {
			// 如果想要一箭射穿的话, 需要更新一下能射到的最远的距离
			cur = points[i][1]
		}
	}

	return res
}

func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	// 需要先拍一下顺序 按照右端点排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	var res = 1
	var cur = points[0][1]
	for i := 1; i < len(points); i++ {
		// 如果当前气球的左端点大于当前记录的最远右端点, 就需要射下一箭
		if points[i][0] > cur {
			res++
			cur = points[i][1]
		}
	}
	return res
}
