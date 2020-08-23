package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	// 如果S 是由 n个s串组成的, 那么 S+S由 2n个s串组成
	// S 必然会在 S+S[1:len(S+S)-1]中存在 至少一次
	return strings.Index((s + s)[1:len(s)*2-1], s) != -1
}
