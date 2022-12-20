package main

import "sort"

func minimumSize(nums []int, maxOperations int) int {
	max := 0
	for _, x := range nums {
		if x > max {
			max = x
		}
	}
	// 左边界是0, 右边界是数组中元素的最大值(不进行任何分割)
	// 尝试将 所有球按照每份 y 个进行分割
	return sort.Search(max, func(y int) bool {
		if y == 0 {
			return false
		}
		ops := 0
		for _, x := range nums {
			// 将x分为y个/份 需要的最少操作
			ops += (x - 1) / y
		}
		return ops <= maxOperations
	})
}
