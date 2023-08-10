package main

func findSubstringInWraproundString(p string) (ans int) {
	var dp [26]int
	k := 0
	for i, ch := range p {
		// ab, b-a =   1 % 26 = 1
		// za, a-z = -25 % 26 = 1
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { // 字符之差为 1 或 -25
			k++
		} else {
			// 不是连续的, 重置长度从头计数
			k = 1
		}
		dp[ch-'a'] = max(dp[ch-'a'], k)
	}
	for _, v := range dp {
		ans += v
	}
	return
}
