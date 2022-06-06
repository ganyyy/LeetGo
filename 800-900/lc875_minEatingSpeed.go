package main

func minEatingSpeed(piles []int, h int) int {
	// 值二分法?

	var maxP int
	for _, p := range piles {
		if p > maxP {
			maxP = p
		}
	}

	var left, right = 1, maxP

	for left < right {
		var mid = left + (right-left)/2
		var cost = 0
		for _, v := range piles {
			// 向上取整
			cost += (v + mid - 1) / mid
		}
		if cost > h {
			// 耗时较高, 需要吃更多的香蕉
			left = mid + 1
		} else {
			// 耗时低, 可以少吃一点
			right = mid
		}
	}
	return left
}
