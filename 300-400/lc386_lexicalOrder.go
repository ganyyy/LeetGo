package main

func lexicalOrder(n int) []int {
	var ret = make([]int, 0, n)

	var cur = 1

	for i := 0; i < n; i++ {
		ret = append(ret, cur)

		// 先计算1, 10, 100...
		if cur*10 <= n {
			cur *= 10
			continue
		}
		// 进制回退 100 -> 10
		if cur >= n {
			cur /= 10
		}
		//再计算11,12,13,14
		cur++
		// 如果加到10的倍数, 就需要先去掉0
		// 比如 110 -> 11
		// 20 -> 2
		for cur%10 == 0 {
			cur /= 10
		}
	}

	return ret
}

func main() {
	lexicalOrder(1000)
}
