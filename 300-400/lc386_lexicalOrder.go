package main

func lexicalOrder(n int) []int {
	var ret = make([]int, 0, n)

	var cur = 1

	for i := 0; i < n; i++ {
		ret = append(ret, cur)

		if cur*10 <= n {
			cur *= 10
			continue
		}
		if cur >= n {
			cur /= 10
		}
		cur++
		for cur%10 == 0 {
			cur /= 10
		}
	}

	return ret
}
