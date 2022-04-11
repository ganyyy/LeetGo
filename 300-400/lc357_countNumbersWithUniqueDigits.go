package main

func countNumbersWithUniqueDigits(n int) int {

	var ret int

	// 1位: 9
	// 2位: 9*9
	// 3位: 9*9*8
	// ...

	// 累加所有值即可
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	ret = 10
	var cur = 9
	for t := 2; t <= n; t++ {
		cur *= 11 - t
		ret += cur
	}

	return ret
}
