package main

func numberOfMatches(n int) int {
	var cnt int

	for n != 1 {
		cnt += n >> 1
		if n&1 != 0 {
			n += 1
		}
		n = n >> 1
	}

	return cnt
}
