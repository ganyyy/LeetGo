package main

func isCovered(ranges [][]int, left int, right int) bool {
	var diff [52]int

	// 差分数组的算法思想:

	// 差分数组的定义: diff[i] = nums[i]-nums[i-1]
	// 所以对nums[i:j+1]这个区间进行整体的增减操作时,
	// 不需要每一项都增/减, 只需要对diff[i], diff[j+1]进行操作即可

	for _, r := range ranges {
		// 这列意思就等同于把 r[0]:r[1]+1 区间内所有的值都+1
		diff[r[0]]++
		diff[r[1]+1]--
	}

	// 为啥要求前缀和呢?
	// 这个表示的是当前数字是否存在于区间内
	var cur int
	for i := 1; i <= right; i++ {
		cur += diff[i]
		if i >= left && cur <= 0 {
			return false
		}
	}

	return true
}

func isCoveredOld(ranges [][]int, left int, right int) bool {
	var count = make([]int, right-left+1)

	for _, r := range ranges {
		if r[1] < left || r[0] > right {
			continue
		}

		for i := max(r[0], left); i <= min(r[1], right); i++ {
			count[i-left] = 1
		}
	}

	for _, v := range count {
		if v == 0 {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
