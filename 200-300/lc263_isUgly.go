package main

func isUgly(n int) bool {
	if n < 1 {
		return false
	}

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
