package main

func distinctSubseqII(s string) (total int) {

	// 对于位置i而言, 相当于求和f[0..i]并加上1
	// 但是如果存在f[j] == f[k], 那么就会出现重复计算的情况(s[0:j]+s[k] == s[:j+1], 此时f[j]就会被重复计算)
	// 所以需要计算到每个字符上一次出现的位置即可

	const mod int = 1e9 + 7
	g := make([]int, 26)
	// total 表示累加和
	for _, c := range s {
		oi := c - 'a'
		// g[oi]表示f[j]
		prev := g[oi]
		// f[i] = total+1
		g[oi] = (total + 1) % mod
		// 整体需要去掉 prev 的部分, 因为对于s[i]而言, 这部分被重复计算了
		// 为什么又加了一次mod呢? 因为会出现溢出的问题
		total = ((total+g[oi]-prev)%mod + mod) % mod
	}
	return
}
