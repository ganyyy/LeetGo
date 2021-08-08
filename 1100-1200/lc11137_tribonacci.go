package main

func tribonacci(n int) int {
	var a0, a1, a2 = 0, 1, 1
	for t := 0; t < n; t++ {
		a0, a1, a2 = a1, a2, a0+a1+a2
	}
	return a0
}
