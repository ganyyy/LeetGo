package main

func checkValidString(str string) bool {
	/*
	   简单堆栈, 完全可以通过数字表示

	   这道题本质上, 是一个拥有上下界的栈.

	   lo表示 * 充当右括号时左括号的剩余数量
	   hi表示 * 充当左括号时左括号的剩余数量
	*/
	var lo, hi int
	for _, v := range str {
		switch v {
		case '(':
			lo++
			hi++
		case ')':
			lo = max(0, lo-1)
			hi--
			if hi < 0 {
				return false
			}
		case '*':
			// * 表示, 当左括号大于0时, 充当右括号; 否则就充当左括号
			lo = max(lo-1, 0)
			hi++
		}
	}
	return lo <= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
