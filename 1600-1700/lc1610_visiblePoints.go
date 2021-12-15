package main

import (
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	// 所有和坐标点相同的点的数量
	sameCnt := 0
	var polarDegrees []float64
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			// 获取每一个点相较于所在位置对应的弧度
			polarDegrees = append(polarDegrees, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		// 将所有点重复加入, 避免了边界值(360°)问题
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	right := 0
	// 可视角度转变为弧度. 将问题转变为滑动窗口去解决
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		// 依次迭代每个点, 计算以该点为边界值, 整个弧度内能看到的最多的点的数量
		for right < n*2 && polarDegrees[right] <= deg+degree {
			right++
		}
		if right-i > maxCnt {
			maxCnt = right - i
		}
	}
	return sameCnt + maxCnt
}
