package main

func isUgly(n int) bool {
	if n < 1 {
		return false
	}

	// 从大到小, 依次的除过来. 最终剩下的如果不是1就说明有问题

	for n%5 == 0 {
		n /= 5
	}
	for n%3 == 0 {
		n /= 3
	}

	for n&1 == 0 {
		n >>= 1
	}

	return n == 1
}
