package main

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	f := make([]int, n+1)
	// 画个图看看呗. 顺便
	f[0], f[1], f[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		f[i] = (f[i-1]*2 + f[i-3]) % (1e9 + 7)
	}
	return f[n]
}

func numTilings2(n int) int {
	// https://leetcode.cn/problems/domino-and-tromino-tiling/solutions/1968516/by-endlesscheng-umpp/
	if n == 1 {
		return 1
	}
	a, b, c := 1, 1, 2
	for i := 3; i <= n; i++ {
		a, b, c = b, c, (c*2+a)%(1e9+7)
	}
	return c
}
