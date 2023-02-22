package main

func circularPermutation(n int, start int) []int {
	// start 一定是 2^n - 1
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = (i >> 1) ^ i ^ start
	}
	return ans
}
