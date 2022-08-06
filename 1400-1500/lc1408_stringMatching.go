package main

import (
	"sort"
	"strings"
)

// kmp

func stringMatching(words []string) []string {
	var nextMap = make(map[string][]int, len(words))

	var buildNext = func(s string) []int {
		if next, ok := nextMap[s]; ok {
			return next
		}
		// build
		var ret = make([]int, len(s))
		ret[0] = -1
		var nv = -1
		// 第一位一定是-1,
		// 最后一位不用看
		// 综上所述, 取值区间是[0:len(s)-1]
		for i := 0; i < len(s)-1; {
			if nv == -1 || s[i] == s[nv] {
				// 如果nv == -1, 那么ret[i] == 0
				nv++
				i++
				ret[i] = nv
			} else {
				nv = ret[nv]
			}
		}
		nextMap[s] = ret
		return ret
	}

	var kmp = func(a, b string) bool {
		if len(a) < len(b) {
			return false
		}
		var next = buildNext(a)

		// compare
		var i, j int
		for i < len(a) && j < len(b) {
			if j == -1 || a[i] == b[j] {
				i++
				j++
			} else {
				j = next[j]
			}
		}
		return j == len(b)
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	var ret []string

	for j := len(words) - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if kmp(words[i], words[j]) {
				ret = append(ret, words[j])
				break
			}
		}
	}
	return ret
}

func stringMatching2(words []string) (ans []string) {
	for i, x := range words {
		for j, y := range words {
			if j != i && strings.Contains(y, x) {
				ans = append(ans, x)
				break
			}
		}
	}
	return
}

func main() {
	stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"})
}
