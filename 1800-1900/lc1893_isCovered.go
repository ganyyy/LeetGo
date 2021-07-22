package main

func isCovered(ranges [][]int, left int, right int) bool {
	var diff [52]int

	for _, r := range ranges {
		diff[r[0]]++
		diff[r[1]+1]--
	}

	var cur int
	for i, v := range diff {
		cur += v
		if i >= left && i <= right && cur <= 0 {
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
