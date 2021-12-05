package main

const MOD = 1337
const PHI = 1140 // 小于MOD, 并且与MOD互质的数的个数

func superPow(a int, b []int) int {
	// 在 c是一个质数的情况下,
	// a^b % c = a^(b%(PHI(c)))%c

	if a == 1 || len(b) == 0 {
		return a
	}

	var exp int
	for _, v := range b {
		exp = (exp*10 + v) % PHI
	}

	return quickPow(a, exp)
}

func quickPow(a, b int) int {
	var res = 1
	for b > 0 {
		if b&1 == 1 {
			res = (a * res) % MOD
		}
		b /= 2
		a = (a * a) % MOD
	}
	return res
}

func superPowNormal(a int, b []int) int {
	ans := 1
	// 降幂
	// a^1111
	// -> a^{(111 * 10) + 1}
	// -> a^{111 * 10} * a^1
	// -> {a^111}^10 * a
	for i := len(b) - 1; i >= 0; i-- {
		// 先乘以尾数
		ans = ans * quickPow(a, b[i]) % MOD
		// 再乘以10次幂
		a = quickPow(a, 10)
	}
	return ans
}
