package main

import "math"

func maxSubArray(nums []int) int {
	var sum, ret = 0, math.MinInt32
	for _, v := range nums {
		sum += v
		if sum > ret {
			ret = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return ret
}
