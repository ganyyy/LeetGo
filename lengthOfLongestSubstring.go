package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return sLen
	}
	m := make(map[int32]int)
	var start, mLen, l int
	for i, c := range s {
		if v, ok := m[c]; ok && v >= start {
			start = v + 1
		}
		m[c] = i
		l = i - start + 1
		if l > mLen {
			mLen = l
		}
	}
	return mLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}
