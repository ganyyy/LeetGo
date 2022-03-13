package main

import "math"

func maxProfit(prices []int) int {

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var profit int
	var mi = math.MaxInt32

	for _, price := range prices {
		profit = max(profit, price-mi)
		mi = min(mi, price)
	}

	return profit
}
