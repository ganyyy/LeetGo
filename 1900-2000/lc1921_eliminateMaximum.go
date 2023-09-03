package main

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	// 时间
	var times = make([]float64, len(dist))
	for i := range dist {
		times[i] = float64(dist[i]) / float64(speed[i])
	}
	// 按照时间排序?
	sort.Float64s(times)

	// 贪心
	var i int
	for i = 1; i < len(times); i++ {
		if float64(i) >= times[i] {
			break
		}
	}

	return i
}
