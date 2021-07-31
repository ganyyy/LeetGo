package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {

	const (
		POW  = 8
		BASE = (1 << POW) - 1
	)

	var tmp = make([]int, 0, len(mat))

	for i, row := range mat {
		var sum int
		for _, v := range row {
			sum += v
		}
		tmp = append(tmp, (sum<<POW)+i)
	}

	sort.Ints(tmp)

	var ret = make([]int, k)
	for i, v := range tmp[:k] {
		ret[i] = v & BASE
	}
	return ret
}
