package main

func minDays(n int) int {
	memo := map[int]int{0: 0, 1: 1}

	var helper func(n int) int
	helper = func(n int) int {
		if days, ok := memo[n]; ok {
			return days
		}
		// n%2 n%3 得到的结果, 等同于在进行2次缩减和3次缩减时需要执行的额外的-1的次数.
		memo[n] = min(n%2+1+helper(n/2), n%3+1+helper(n/3))
		return memo[n]
	}
	return helper(n)
}
