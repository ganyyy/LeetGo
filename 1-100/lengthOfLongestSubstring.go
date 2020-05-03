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
	// pos[v] = t 第几个位置出现的
	pos := make([]int, 128)
	// 最大值和左指针
	var max, start int
	for i, v := range s {
		// 存在 且 在出现在左指针的后边, 计算最大长度, 更新左值针位置
		if t := pos[v]; 0 != t && start < t {
			// 计算长度
			if l := i - start; l > max {
				max = l
			}
			// 重置位置
			start = t
		}
		// 更新一下当前字母的出现位置
		pos[v] = i + 1
	}
	// 可能会出现最后一个字母 不在pos中, 做一个判断
	if v := len(s) - start; v > max {
		return v
	} else {
		return max
	}
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("aab"))
}
