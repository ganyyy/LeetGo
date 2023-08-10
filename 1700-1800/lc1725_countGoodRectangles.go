package main

func countGoodRectangles(rectangles [][]int) int {
	var cnt int
	var m int
	for _, r := range rectangles {
		var side = minSide(r)
		if side == m {
			cnt++
		} else if side > m {
			m = side
			cnt = 1
		}
	}
	return cnt
}

func minSide(r []int) int {
	var a, b = r[0], r[1]
	if a < b {
		return a
	}
	return b
}
