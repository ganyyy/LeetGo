package main

func guessNumber(n int) int {

	var guess = func(int2 int) int {
		return 0
	}

	var l, r = 1, n + 1

	for l < r {
		var mid = l + (r-l)/2

		if t := guess(mid); t == 0 {
			return mid
		} else if t == 1 {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return -1
}
