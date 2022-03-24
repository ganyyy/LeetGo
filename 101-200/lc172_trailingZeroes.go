package main

func trailingZeroes(n int) int {
	var cnt int

	for n != 0 {
		cnt += n / 5
		n /= 5
	}
	return cnt
}
