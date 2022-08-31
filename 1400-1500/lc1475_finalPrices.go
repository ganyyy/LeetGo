package main

func finalPrices(prices []int) []int {
	// 单调栈
	// 递增

	var stack []int
	ret := make([]int, len(prices))
	for i, v := range prices {
		for len(stack) != 0 && prices[stack[len(stack)-1]] >= v {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[idx] = prices[idx] - v
		}
		stack = append(stack, i)
	}

	for _, idx := range stack {
		ret[idx] = prices[idx]
	}

	return ret
}
