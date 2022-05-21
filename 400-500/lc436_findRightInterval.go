package main

import "sort"

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	type pair struct{ x, i int }
	starts := make([]pair, n)
	ends := make([]pair, n)
	for i, p := range intervals {
		starts[i] = pair{p[0], i}
		ends[i] = pair{p[1], i}
	}
	sort.Slice(starts, func(i, j int) bool { return starts[i].x < starts[j].x })
	sort.Slice(ends, func(i, j int) bool { return ends[i].x < ends[j].x })

	ans := make([]int, n)
	j := 0
	for _, p := range ends {
		for j < n && starts[j].x < p.x {
			j++
		}
		if j < n {
			ans[p.i] = starts[j].i
		} else {
			ans[p.i] = -1
		}
	}
	return ans
}
