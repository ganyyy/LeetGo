package main

import "sort"

func hIndex3(citations []int) int {

	for i, v := range citations {
		var cnt = len(citations) - i
		if v >= cnt {
			return cnt
		}
	}
	return 0
}

func hIndexGood(citations []int) int {
	// lh, rh := 0, len(citations)
	ln := len(citations)

	return ln - sort.Search(ln, func(midh int) bool {
		return citations[midh] >= ln-midh
	})
}
