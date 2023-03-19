package main

import "strconv"

func numDupDigitsAtMostN(n int) (ans int) {
	s := strconv.Itoa(n)
	// m上限是9
	m := len(s)
	// 我们这里的数位dp, 求取的是不存在相同位的数字的个数
	// 所以需要通过 一个 mask 来表示当前已选取的位置
	// 如果只看选不选, 一位数就够了, 但是因为每一种组合都对应一个解, 所以必须要采用一个数组存储结果
	dp := make([][1 << 10]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1 // -1 表示没有计算过
		}
	}
	// i:       当前是多少位
	// mask:    当前这位上, 所有选取的数字(避免重复计算)
	// isLimit: 这一位是否存在上限? 比如, n = 100, 那么在第2/3位上可选的是 [0, 9], 但是第1位上只能选[0,1]
	//          注意: 这里是从高到低迭代的数字, 所以首位limit = true, 后续的limit = false
	// isNum  : 这是不是一个合法的数字? 起始状态下, 因为没有选取任何数字, 所以是false
	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1 // 得到了一个合法数字, 并且它不包含重复的数字
			}
			return
		}
		if !isLimit && isNum {
			// 记忆化搜索+枝减
			// 对于第i位上, 选取了 mask 里的数字的情况下, 有效答案是 res
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return n - f(0, 0, true, false)
}
