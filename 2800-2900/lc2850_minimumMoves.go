package main

import "math"

type pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
	var from, to []pair
	for i, row := range grid {
		for j, cnt := range row {
			if cnt > 1 {
				for k := 1; k < cnt; k++ {
					from = append(from, pair{i, j})
				}
			} else if cnt == 0 {
				to = append(to, pair{i, j})
			}
		}
	}

	ans := math.MaxInt
	permute(from, func() {
		total := 0
		for i, f := range from {
			total += abs(f.x-to[i].x) + abs(f.y-to[i].y)
		}
		ans = min(ans, total)
	})
	return ans
}

func permute(a []pair, do func()) {
	if len(a) == 0 {
		do()
		return
	}
	for i := 0; i < len(a); i++ {
		a[0], a[i] = a[i], a[0]
		permute(a[1:], do)
		a[0], a[i] = a[i], a[0]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
