package main

import "sort"

func numMatchingSubseq(s string, words []string) int {
	pos := [26][]int{}
	// 存储每个字符对应的下标切片
	for i, c := range s {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	ans := len(words)
	for _, w := range words {
		if len(w) > len(s) {
			ans--
			continue
		}
		p := -1
		for _, c := range w {
			// 二分查找字母的位置
			ps := pos[c-'a']
			// p+1所处的位置/插入的位置
			j := sort.SearchInts(ps, p+1)
			if j == len(ps) {
				ans--
				break
			}
			p = ps[j]
		}
	}
	return ans
}
