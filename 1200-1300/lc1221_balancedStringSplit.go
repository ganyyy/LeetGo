package main

func balancedStringSplit(s string) int {
	var r, cc int

	// 简简单单一个贪心, 搞起来
	var cnt int
	for r < len(s) {
		if s[r] == 'R' {
			cc++
		} else {
			cc--
		}
		if cc == 0 {
			cnt++
		}
		r++
	}
	return cnt
}
