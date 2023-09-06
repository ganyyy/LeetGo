package main

import (
	"math"
	"sort"
)

func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	for _, r := range ranks {
		if r < minR {
			minR = r
		}
	}
	// 上界就是能力最差的人修所有的车
	return int64(sort.Search(minR*cars*cars, func(t int) bool {
		s := 0
		for _, r := range ranks {
			// 每个工人在时间t内, 最多可以修几辆车呢?
			// r*n^2 = t => n = sqrt(t/r)
			s += int(math.Sqrt(float64(t / r)))
		}
		// 累计每个工人的修车数, 二分查找答案.
		return s >= cars
	}))
}
