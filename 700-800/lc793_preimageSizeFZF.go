package main

import "sort"

func zeta(n int) (res int) {
	// 返回阶乘结果中, 0的个数
	for n > 0 {
		n /= 5
		res += n
	}
	return
}

// 至少k个0的阶乘数的个数
func nx(k int) int {
	// 可以这样理解,
	// zeta(x) = sum(x/5, x/25, x/125, x/625...) >= x/5
	// zeta(5x) >= x (两边同时乘以5)
	return sort.Search(5*k, func(x int) bool { return zeta(x) >= k })
}

func preimageSizeFZF(k int) int {
	return nx(k+1) - nx(k)
}
