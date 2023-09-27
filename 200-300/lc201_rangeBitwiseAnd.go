package main

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	// 这是在找 n和m 二进制字符串的的最小公共前缀(从左向右)
	// 因为是相与, 所以最小公共前缀之后的位都会被消去
	for n > m {
		// 消去n 最右边的1,
		n &= n - 1
	}

	return n
}
