package main

import "strings"

func repeatedStringMatch(a string, b string) int {
	var cnt = 1
	var src = a
	if len(b) > len(a) {
		cnt = (len(b) + len(a) - 1) / len(a) // 取a的整数倍
		a = strings.Repeat(a, cnt)
	}
	if strings.Index(a, b) != -1 {
		return cnt
	}
	if strings.Index(a+src, b) != -1 {
		return cnt + 1
	}
	return -1
}
