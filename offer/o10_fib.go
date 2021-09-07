package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	var a, b = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return b % (1e9 + 7)
}
