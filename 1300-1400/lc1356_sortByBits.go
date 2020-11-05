package main

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		var l, r = cntBits(arr[i]), cntBits(arr[j])
		if l != r {
			return l < r
		}
		return arr[i] < arr[j]
	})
	return arr
}

func cntBits(n int) int {
	var r int
	for n != 0 {
		n &= n - 1
		r++
	}
	return r
}
