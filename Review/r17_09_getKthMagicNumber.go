package main

import "math"

var val = [3]int{3, 5, 7}

func getKthMagicNumber(k int) int {
	var dp = [3]int{0, 0, 0}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	result := make([]int, 1, k)
	result[0] = 1

	for i := 2; i <= k; i++ {
		rn := math.MaxInt32
		for idx := 0; idx < 3; idx++ {
			rn = min(val[idx]*result[dp[idx]], rn)
		}

		if rn%3 == 0 {
			dp[0]++
		}
		if rn%5 == 0 {
			dp[1]++
		}
		if rn%7 == 0 {
			dp[2]++
		}
		result = append(result, rn)
	}
	return result[k-1]
}
