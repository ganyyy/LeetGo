//go:build ignore

package main

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// dp, 从前到后累计成串的数量
	var ret int
	var dp = make([]int, len(words))
	for i := 1; i < len(words); i++ {
		// 自己不算, 末尾+1
		var b = words[i]
		for j := 0; j < i; j++ {
			var a = words[j]
			if distance(a, b) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(ret, dp[i])
	}

	return ret + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func distance(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}

	var diff int
	var ai int
	for bi := 0; ai < len(a) && bi < len(b); bi++ {
		if a[ai] == b[bi] {
			ai++
			continue
		}
		diff++
		if diff > 1 {
			return false
		}
	}
	return true
}
