package main

func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 100 -> 50 -> 25 -> 13/12 -> 6/7/6 -> 3/3/3/4 -> 1/2/1/2/1/2/2 -> 1/1/1/1/1/1/1/1

	var buf = make([]int, n+1)
	var ret int

	buf[0] = 0
	buf[1] = 1

	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			// 偶数
			buf[i] = buf[i>>1]
		} else {
			buf[i] = buf[i>>1] + buf[(i>>1)+1]
		}
		ret = max(ret, buf[i])
	}
	return ret
}
