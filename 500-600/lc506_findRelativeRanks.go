package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	const (
		SHIFT = 20
		MASK  = (1 << SHIFT) - 1
	)

	for i, v := range score {
		score[i] = (v << SHIFT) | i
	}

	sort.Ints(score)

	var ret = make([]string, len(score))

	var all = len(score) - 1
	for r, v := range score {
		ret[v&MASK] = GetMedal(all - r)
	}
	return ret
}

func GetMedal(rank int) string {
	switch rank {
	case 0:
		return "Gold Medal"
	case 1:
		return "Silver Medal"
	case 2:
		return "Bronze Medal"
	default:
		return strconv.Itoa(rank + 1)
	}
}
