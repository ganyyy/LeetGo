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

func lengthOfLongestSubstring2(s string) int {
	pos := make([]int, 128)
	for i := 0; i < 128; i++ {
		pos[i] = -1
	}
	var max, start int
	for i, v := range s {
		if t := pos[v]; -1 != t && start <= t {
			// 计算长度
			if l := i - start; l > max {
				max = l
			}
			// 重置位置
			start = t + 1
		}
		pos[v] = i
	}
	if v := len(s) - start; v > max {
		return v
	} else {
		return max
	}
}

func main() {
	fmt.Println(lengthOfLongestSubstring2(""))
}
