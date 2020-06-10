package main

func isPalindrome(x int) bool {
	// 不成立的情况
	// 1. 小于0
	// 2. 不等于0且和10取余是0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 不需要全部反转, 只需要反转其中的一半即可
	var res int
	for x > res {
		res = res*10 + x%10
		x /= 10
	}

	// 如果是偶数, 那么x和res二者相同
	// 如果是奇数, 那么res比x多了一位数
	return x == res || x == res/10
}
