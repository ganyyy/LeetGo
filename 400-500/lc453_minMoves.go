package main

import "math"

func minMoves(nums []int) int {
	// 其余数加1等同于自身减1

	// 所以, 计算将所有值减到最小值即可
	var min = math.MaxInt32

	var sum int
	for _, v := range nums {
		if min > v {
			min = v
		}
		sum += v
	}

	return sum - len(nums)*min

}

// func minMoves(nums []int) int {
// 	sum, minNum := 0, nums[0]
// 	for _, val := range nums {
// 		sum += val
// 		if val < minNum {
// 			minNum = val
// 		}
// 	}
// 	return sum - len(nums)*minNum
// }
