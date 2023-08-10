package main

func balancedString(s string) int {
	var cnt [128]int
	for _, c := range s {
		cnt[byte(c)]++
	}
	partial := len(s) / 4
	check := func() bool {
		return !(cnt['Q'] > partial ||
			cnt['W'] > partial ||
			cnt['E'] > partial ||
			cnt['R'] > partial)
	}

	if check() {
		return 0
	}

	// 最长的结果, 就是整串替换
	res := len(s)
	r := 0
	for l, c := range s {
		// 右++, 直到不满足
		for r < len(s) && !check() {
			cnt[s[r]]--
			r += 1
		}
		if !check() {
			break
		}
		res = min(res, r-l)
		cnt[byte(c)]++
	}
	return res
}
