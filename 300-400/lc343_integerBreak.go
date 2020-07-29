package main

func integerBreak(n int) int {
	// 看有几个 2和 3
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	var a = 1
	// 等于4的情况, 2*2 > 3 * 1, 剩下的就看有几个3 就完事了
	for n > 4 {
		n -= 3
		a *= 3
	}
	return a * n
}
