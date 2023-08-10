package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var sum float64
	var kf = float64(k)
	var res float64 = math.MinInt32
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	res = max((sum)/kf, res)

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i] - nums[i-k])
		res = max((sum)/(kf), res)
	}

	return res
}
