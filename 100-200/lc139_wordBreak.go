//go:build ignore

package main

// 这是一个失败的解决方式
// 以长度进行的区分, 果不其然超时了
func wordBreakFail(s string, wordDict []string) bool {
	keyMap := make(map[int][]string)
	for _, v := range wordDict {
		keyMap[len(v)] = append(keyMap[len(v)], v)
	}
	var check func(s string) bool
	check = func(s string) bool {
		ls := len(s)
		if ls == 0 {
			return true
		}
		// 每次只找一个单词
		for ln, words := range keyMap {
			if ln > ls {
				continue
			} else {
				ts := s[:ln]
				for _, word := range words {
					if ts == word {
						if check(s[ln:]) {
							return true
						}
						break
					}
				}
			}
		}
		return false
	}
	return check(s)
}

// 通过DP进行处理
func wordBreak(s string, wordDict []string) bool {
	// dp方式解决

	// 先记录所有出现的单词
	m := make(map[string]struct{}, len(wordDict))
	var empty = struct{}{}
	// 最大长度的字符串的长度
	var maxLn int
	for _, v := range wordDict {
		m[v] = empty
		if l := len(v); l > maxLn {
			maxLn = l
		}
	}
	// 初始化dp
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// 填充dp数组
	for i := 1; i <= len(s); i++ {
		for j := max(i-maxLn, 0); j < i; j++ {
			if !dp[j] {
				continue
			}
			if _, ok := m[s[j:i]]; ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
