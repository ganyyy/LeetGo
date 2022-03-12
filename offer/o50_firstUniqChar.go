package main

import "math"

func firstUniqChar(s string) byte {
	const (
		SHIFT = 16
		MASK  = (1 << SHIFT) - 1
	)

	var getCnt = func(v int) (num, idx int) {
		return v >> SHIFT, v & MASK
	}

	var setCnt = func(num, idx int) int {
		return (num << SHIFT) | idx
	}

	var cnt [26]int

	for i := range s {
		var ch = s[i]
		var num, _ = getCnt(cnt[ch-'a'])
		cnt[ch-'a'] = setCnt(num+1, i)
	}

	var ret byte = ' '
	var minIdx = math.MaxInt32
	for i, v := range cnt {
		var num, idx = getCnt(v)
		if num != 1 {
			continue
		}
		if idx < minIdx {
			minIdx = idx
			ret = byte(i) + 'a'
		}
	}

	return ret
}
