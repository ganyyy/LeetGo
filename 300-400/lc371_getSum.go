package main

func getSum(a int, b int) int {
	for a&b != 0 {
		a, b = a&b, a^b
		a <<= 1
	}
	return a | b
}
