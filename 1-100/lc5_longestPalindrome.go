package main

func longestPalindrome(s string) string {
	// mark: 中心扩散法居然比动态规划快
	// 迭代次数更少, 内存分配更少

	var search = func(l, r int) string {
		if r > len(s) {
			return ""
		}
		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] != s[r] {
				break
			}
		}
		return s[l+1 : r]
	}

	var ret string
	var maxStr = func(s1, s2 string) string {
		if len(s1) < len(s2) {
			return s2
		}
		return s1
	}
	for i := 0; i < len(s); i++ {
		ret = maxStr(ret, maxStr(search(i, i), search(i, i+1)))
	}
	return ret
}

func longestPalindromeDP(s string) string {
	sl := len(s)
	if sl <= 1 {
		return s
	}
	dp := make([][]bool, sl)
	for i := 0; i < sl; i++ {
		dp[i] = make([]bool, sl)
	}
	start, length := 0, 1
	for r := 1; r < sl; r++ {
		for l := 0; l < r; l++ {
			// 相等
			// 偶数串(l,r)
			// 奇数串(l,x,r)
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				ll := r - l + 1
				if length < ll {
					length = ll
					start = l
				}
			}
		}
	}
	return s[start : start+length]
}
