package main

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	wb := make([]int, 0, len(words))
	for _, word := range words {
		wb = append(wb, f(word))
	}
	sort.Ints(wb)
	// fmt.Println(wb)
	qb := make([]int, 0, len(queries))
	for _, query := range queries {
		qb = append(qb, f(query))
	}
	ln := len(wb)
	// fmt.Println(qb)
	for i, v := range qb {
		// 二分查找大于 v的值的个数?
		idx := sort.SearchInts(wb, v+1)
		qb[i] = ln - idx
	}

	return qb
}

func f(s string) int {
	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
	}
	for _, v := range cnt {
		if v > 0 {
			return v
		}
	}
	return 0
}
