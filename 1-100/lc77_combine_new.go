package main

import "fmt"

func combine(n int, k int) [][]int {
	res := make([][]int, 0, calcCap(n, k))

	tmp := make([]int, 0, k)
	var helper func(c int)

	helper = func(c int) {
		if len(tmp) == k {
			t := make([]int, k)
			copy(t, tmp)
			res = append(res, t)
			return
		}
		for i := c; i <= n; i++ {
			tmp = append(tmp, i)
			helper(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	helper(1)

	return res
}

func calcCap(a, b int) int {
	if a == b {
		return 1
	}
	if b > a-b {
		b = a - b
	}
	for i, t := a-b+1, a; i < t; i++ {
		a *= i
	}
	for i, t := 2, b; i < t; i++ {
		b *= i
	}
	return a / b
}

func main() {
	fmt.Println(combine(4, 2))
}
