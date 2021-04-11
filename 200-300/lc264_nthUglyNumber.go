package main

func nthUglyNumber(n int) int {
	var tmp = make([]int, 0, n)
	tmp = append(tmp, 1)
	var i2, i3, i5 = 0, 0, 0
	var a, b, c, next int
	for i := 1; i <= n; i++ {
		a, b, c = tmp[i2]*2, tmp[i3]*3, tmp[i5]*5
		next = min(a, min(b, c))
		if next == a {
			i2++
		}
		if next == b {
			i3++
		}
		if next == c {
			i5++
		}
		tmp = append(tmp, next)
	}
	return tmp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
