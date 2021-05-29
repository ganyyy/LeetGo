package main

func isPowerOfTwo(x int) bool {
	if x <= 0 {
		return false
	}
	return x&(x-1) == 0
}
