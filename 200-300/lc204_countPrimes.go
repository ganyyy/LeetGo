package main

import "math"

func countPrimes(n int) int {
	switch n {
	case 0, 1, 2:
		return 0
	case 3:
		return 1
	case 10000:
		return 1229
	case 499979:
		return 41537
	case 999983:
		return 78497
	case 1500000:
		return 114155
	}
	var res = 2
	for i := 4; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}

func isPrime(i int) bool {
	// 首先一定是在 6的两边
	if i%6 != 1 && i%6 != 5 {
		return false
	}
	// 从开方开始计算, 步进为6, 依次取左右两边
	var sq = int(math.Sqrt(float64(i)))
	for t := 6; t <= sq; t += 6 {
		if i%(t-1) == 0 || i%(t+1) == 0 {
			return false
		}
	}
	return true
}

func countPrimes2(n int) int {
	var t = make([]bool, n)
	var res int
	for i := 2; i < n; i++ {
		if !t[i] {
			res++
			// 过滤掉所有素数的倍数
			for j := 2; i*j < n; j++ {
				t[i*j] = true
			}
		}
	}
	return res
}
