package main

func minSteps(n int) int {
	// 快速查询

	// 质数就是就是本身(C+P*(n-1))
	// 合数需要分解质因数, 然后判断最小的合成次数

	// 想法如此, 做起来有点麻烦, 有更好的办法吗?

	var res int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			res += i
			n /= i
		}
	}

	return res
}
