package main

func maxArrayValue(nums []int) int64 {
	// 从需求上考虑, 就是两两合并非逆序对, 也就是说, 从后往前遍历, 逐步合并, 才可能做到尽可能多的合并
	sum := int64(nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		if int64(nums[i]) <= sum {
			sum = int64(nums[i]) + sum
		} else {
			sum = int64(nums[i])
		}
	}
	return sum
}
