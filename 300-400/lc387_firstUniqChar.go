package main

import "math"

const MAX = math.MaxInt32

func firstUniqChar(s string) int {
	var set [26]int

	var idx int
	for i := 0; i < len(s); i++ {
		idx = int(s[i] - 'a')
		if set[idx] == 0 {
			set[idx] = i + 1
		} else {
			set[idx] = MAX
		}
	}

	var min = 0
	for i := 0; i < 26; i++ {
		if set[i] != MAX && set[i] != 0 {
			if min == 0 || min > set[i] {
				min = set[i]
			}
		}
	}

	return min - 1
}
