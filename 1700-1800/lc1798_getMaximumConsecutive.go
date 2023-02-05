package main

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	// 这个算法很有意思
	// 简单而言, 就是拼
	// 1, 1, 1, 1 -> 1, 2, 3, 4
	// 1, 2, 3, 4 -> 1, 3, 6, 10 -> 1, 2, 3, 4, (1+4), (2+4), (3+4), (1+3+4), (2+3+4), (1+2+3+4)
	var res = 1
	for _, v := range coins {
		if v > res {
			return res
		}
		res += v
	}
	return res
}
