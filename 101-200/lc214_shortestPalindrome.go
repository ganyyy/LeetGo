package main

func shortestPalindrome(s string) string {

	/*
	   abcd
	   dcba


	*/

	tmp := reverse(s)
	ln := len(s)
	if ln == 0 {
		return ""
	}
	var i int
	for {
		// 因为要在前边添加 最少字符
		// 使其满足回文
		// 找到一个点 i, 满足 tmp[i:] == s[:ln-i](这两段可能是重复的字符串, 可能是一个字符, 可能是一个逆序串)
		// 因为 tmp是 s的逆序字符串, 索引 tmp[:i] 和 s[ln-i:]是满足逆序关系的
		// 索引 tmp[:i] + s 就是一个回文串, 且满足 从前边添加最少字符这一条件
		if tmp[i:] == s[:ln-i] {
			break
		}
		i++
	}
	return tmp[:i] + s
}

func reverse(s string) string {
	res := make([]byte, len(s))
	for i, j := 0, len(s)-1; i <= j; {
		res[i] = s[j]
		res[j] = s[i]
		i++
		j--
	}
	return string(res)
}
