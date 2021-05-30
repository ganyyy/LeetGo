package main

func isPowerOfFour(n int) bool {
	// 首先必须是大于0的2的整数次幂
	// 其次, 要么模3==1
	// 要么1在奇数位上
	return n >= 1 && n&(n-1) == 0 && 0x55555555&n != 0
	// return n >= 1 && n&(n-1) == 0 && n % 3 == 1
}
