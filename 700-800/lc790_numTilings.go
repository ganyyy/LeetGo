package main

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	f := make([]int, n+1)
	f[0], f[1], f[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		f[i] = (f[i-1]*2 + f[i-3]) % (1e9 + 7)
	}
	return f[n]
}
