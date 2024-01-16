package main

func count(num1 string, num2 string, min_sum int, max_sum int) int {

	var memo [23][401]int
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	l1, l2 := len(num1), len(num2)
	bytes1 := make([]byte, l2)
	for i := range bytes1[:l2-l1] {
		bytes1[i] = '0'
	}
	copy(bytes1[l2-l1:], num1)
	bytes2 := []byte(num2)
	var dfs func(total, idx int, lLimit, rLimit bool) int
	dfs = func(total, idx int, lLimit, rLimit bool) int {
		if idx >= l2 {
			if total >= min_sum && total <= max_sum {
				return 1
			}
			return 0
		}
		// 只有在没有超过边界的情况下, 可以使用cache
		useMemo := !(lLimit || rLimit)
		if useMemo && memo[idx][total] != -1 {
			return memo[idx][total]
		}
		var ret int
		const MOD = int(1e9 + 7)
		low := 0
		high := 9
		if lLimit {
			low = int(bytes1[idx] - '0')
		}
		if rLimit {
			high = int(bytes2[idx] - '0')
		}
		for i := low; i <= high; i++ {
			ret = ret + dfs(total+i, idx+1, lLimit && i == low, rLimit && i == high)
		}
		ret %= MOD
		if useMemo {
			memo[idx][total] = ret
		}
		return ret
	}
	return dfs(0, 0, true, true)
}
