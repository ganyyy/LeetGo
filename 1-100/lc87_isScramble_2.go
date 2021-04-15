package main

func isScramble(s1 string, s2 string) bool {
	// 这里需要使用map进行缓存
	m := make(map[string]bool)
	return dfs(s1, s2, m)
}

func dfs(s1 string, s2 string, m map[string]bool) bool {
	if res, ok := m[getConcatString(s1, s2)]; ok {
		return res
	}
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) || !hasSameChar(s1, s2) {
		return false
	}
	if len(s1) <= 3 {
		return true
	}

	n := len(s1)
	for i := 1; i <= n-1; i++ {
		if dfs(s1[0:i], s2[0:i], m) && dfs(s1[i:], s2[i:], m) {
			m[getConcatString(s1, s2)] = true
			return true
		}

		if dfs(s1[0:i], s2[n-i:], m) && dfs(s1[i:], s2[0:n-i], m) {
			m[getConcatString(s1, s2)] = true
			return true
		}
	}
	m[getConcatString(s1, s2)] = false
	return false
}

func hasSameChar(s1, s2 string) bool {
	letters := [26]int{}
	for i := range s1 {
		letters[s1[i]-'a']++
		letters[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if letters[i] != 0 {
			return false
		}
	}
	return true
}

func getConcatString(s1, s2 string) string {
	return s1 + "_" + s2
}
