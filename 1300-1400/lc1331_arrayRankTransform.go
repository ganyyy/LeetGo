package main

import "sort"

func arrayRankTransform(arr []int) []int {
	var rank = make(map[int]int, len(arr))
	var tmp = make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(arr)

	var cnt = 1
	for _, v := range arr {
		if _, ok := rank[v]; !ok {
			rank[v] = cnt
			cnt++
		}
	}

	for i, v := range tmp {
		arr[i] = rank[v]
	}
	return arr
}
