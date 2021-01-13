package main

func prefixesDivBy5(A []int) []bool {
	// 对5取余的结果为0, 只需要个位数是 0或者5即可
	var res = make([]bool, len(A))

	var num int
	for i, v := range A {
		num <<= 1
		num += v
		num %= 10
		res[i] = num%5 == 0
	}

	return res
}
