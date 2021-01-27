package main

func pivotIndex(nums []int) int {
	// 求和
	var sum int
	for _, v := range nums {
		sum += v
	}

	// 从头到尾求结果
	var cur int
	for i, v := range nums {
		if sum-v == cur<<1 {
			return i
		}
		cur += v
	}
	return -1
}
