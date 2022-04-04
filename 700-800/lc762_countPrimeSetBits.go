package main

import "math/bits"

func countPrimeSetBits(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		// 2,3,5,7,11,13,17,19: 为质数的统计个数
		if 1<<bits.OnesCount(uint(x))&665772 != 0 {
			ans++
		}
	}
	return
}
