package main

import (
	"fmt"
	"unsafe"
)

var empty = struct{}{}

func wordBreak(s string, wordDict []string) []string {
	var m = make(map[string]struct{}, len(wordDict))
	for _, v := range wordDict {
		m[v] = empty
	}
	var h = make(map[string]struct{}, 100)

	var helper func(s string)

	var res []string
	var tmp []byte
	helper = func(s string) {
		for i := range s {
			var ts = s[:i+1]
			if _, ok := m[ts]; ok {
				tmp = append(tmp, ts...)
				if i != len(s)-1 {
					tmp = append(tmp, ' ')
					helper(s[i+1:])
					tmp = tmp[:len(tmp)-len(ts)-1]
				} else {
					var t = make([]byte, len(tmp))
					copy(t, tmp)
					tmp = tmp[:len(tmp)-len(ts)]
					var ss = toString(t)
					if _, ok := h[ss]; ok {
						continue
					}
					h[ss] = empty
					res = append(res, ss)
				}
			}
		}
	}
	helper(s)
	return res
}

func wordBreak2(s string, wordDict []string) []string {
	// 遍历字母是行不通了, 那就遍历单词吧

	// 加一个缓存
	var m = map[string][]string{}

	// 要不要先排个序?

	var helper func(s string) []string
	helper = func(s string) []string {
		if v, ok := m[s]; ok {
			return v
		}
		if len(s) == 0 {
			// 方便结尾的处理
			return []string{""}
		}
		var r []string
		for _, word := range wordDict {
			if len(s) < len(word) {
				continue
			}
			if s[:len(word)] != word {
				continue
			}
			var t = helper(s[len(word):])
			for _, ts := range t {
				if len(ts) != 0 {
					r = append(r, word+" "+ts)
				} else {
					r = append(r, word)
				}
			}
		}
		// 缓存一下结果
		m[s] = r
		return r
	}

	return helper(s)
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	/*
		"catsanddog"
		["cat","cats","and","sand","dog"]
	*/
	var t = []string{"a", "aa", "aaa", "aaaa"}
	for _, v := range wordBreak2("aaaa", t) {
		fmt.Println(v)
	}
}
