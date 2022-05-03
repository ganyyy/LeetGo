package main

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	var isDigit = func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	var checkIsDigitLog = func(s string) bool {
		var idx = strings.Index(s, " ")
		if idx == -1 {
			return false
		}
		return isDigit(s[idx+1])
	}

	var split = func(s string) (string, string) {
		var idx = strings.Index(s, " ")
		if idx == -1 {
			return s, ""
		}
		return s[:idx], s[idx+1:]
	}

	var left int

	for right := 0; right < len(logs); right++ {
		if !checkIsDigitLog(logs[right]) {
			continue
		}
		logs[left], logs[right] = logs[right], logs[left]
		left++
	}

	sort.Slice(logs[left:], func(i, j int) bool {
		var a, b = logs[left+i], logs[left+j]

		var h1, c1 = split(a)
		var h2, c2 = split(b)

		if c1 != c2 {
			return c1 < c2
		}
		return h1 < h2
	})

	return append(logs[left:], logs[:left]...)
}
