package main

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	var w1, w2 = math.MinInt32, math.MinInt32
	var abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var ret = math.MaxInt32
	for idx, word := range words {
		if word == word1 {
			w1 = idx
		} else if word == word2 {
			w2 = idx
		}
		if w1 != math.MinInt32 && w2 != math.MinInt32 {
			ret = min(ret, abs(w1-w2))
		}
	}
	return ret
}
