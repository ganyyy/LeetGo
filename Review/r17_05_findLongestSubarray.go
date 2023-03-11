package main

import "unicode"

func findLongestSubarray(array []string) []string {
	// 还是前缀和?

	var m = make(map[int]int)

	var cur int
	var ret []string
	m[0] = -1 // 表示一个没有
	for i, c := range array {
		if unicode.IsNumber(rune(c[0])) {
			cur++
		} else {
			cur--
		}
		if old, ok := m[cur]; ok {
			if i-old > len(ret) {
				ret = array[old+1 : i+1] // 包括i, 不包括old
			}
		} else {
			m[cur] = i
		}
	}
	return ret
}
