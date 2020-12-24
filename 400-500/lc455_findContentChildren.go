package main

import "sort"

func findContentChildren(g []int, s []int) int {
	// 排序呗, 然后从小到大走双指针
	sort.Ints(g)
	sort.Ints(s)

	var i, j int

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}

	return i
}
