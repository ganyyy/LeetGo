package main

func match(word, pattern string) bool {
	var mp = make(map[byte]byte)
	for i := range word {
		x := word[i]
		// 首次映射即可
		y := pattern[i]
		if mp[x] == 0 {
			mp[x] = y
		} else if mp[x] != y { // word 中的同一字母必须映射到 pattern 中的同一字母上
			return false
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) (ans []string) {
	for _, word := range words {
		// 这一步才是重点啊
		if match(word, pattern) && match(pattern, word) {
			ans = append(ans, word)
		}
	}
	return
}
