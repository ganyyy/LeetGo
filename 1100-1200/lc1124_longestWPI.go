package main

func longestWPI(hours []int) int {
	var idx = make(map[int]int, len(hours))

	var sum int
	var ret int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, v := range hours {
		if v > 8 {
			sum++
		} else {
			sum--
		}
		if sum > 0 {
			ret = max(ret, i+1) // 累计大于0
		} else {
			id, ok := idx[sum-1] // 比如sum = -1, 那么只需要找到sum=-2的位置, 此时就是最长的正数队列
			if ok {
				ret = max(ret, i-id)
			}
		}
		if _, ok := idx[sum]; !ok {
			idx[sum] = i
		}
	}
	return ret
}
