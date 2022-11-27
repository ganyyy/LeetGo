package main

const mod int = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
	// 涉及到大数运算, 首先想一下可不可以二分
	l := min(a, b)
	r := n * l
	// c是a/b的最小公倍数
	c := a / gcd(a, b) * b
	// 左右边界, 标准二分
	for l <= r {
		mid := (l + r) / 2
		// a的公倍数
		// b的公倍数
		// a/b共同的公倍数
		cnt := mid/a + mid/b - mid/c
		if cnt >= n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return (r + 1) % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
