package main

func isPowerOfTwo(x int) bool {
	// 一句话的事, 建议尽量简化
	return x > 0 && x&(x-1) == 0
}
