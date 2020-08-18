package main

func countSubstrings(s string) int {
	// 每个字符都是一个回文子串, 看有几个
	var res int

	for i := 0; i < len(s); i++ {
		// 奇数个字符回文串
		res += check(s, i, i)
		// 偶数个字符回文串
		res += check(s, i, i+1)
	}
	return res
}

func check(s string, start, end int) int {
	var cnt int
	for start >= 0 && end < len(s) && s[start] == s[end] {
		cnt++
		start--
		end++
	}
	return cnt
}
