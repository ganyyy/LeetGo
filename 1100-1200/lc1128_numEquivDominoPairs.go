package main

func numEquivDominoPairs(dominoes [][]int) int {
	var res int

	var cp = [100]int{}

	// 题目没读太懂,
	for _, dominoe := range dominoes {
		// 这个等价的意思是, 和前边的每一个都等价吗?
		var v = getVal(dominoe)
		res += cp[v]
		cp[v]++
	}

	return res
}

func getVal(dominoe []int) int {
	var a, b = dominoe[0], dominoe[1]
	if a > b {
		a, b = b, a
	}
	return a*10 + b
}
