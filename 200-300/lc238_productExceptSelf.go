package main

/**
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

包括但不限于 排除数组中指定位置 进行其他操作, 理论上都可以通过这个解决
 */

func productExceptSelf(nums []int) []int {
	ln := len(nums)
	output := make([]int, ln)
	left, right := 1, 1
	// 先左边, 在右边, 注意不要乘以当前的值
	for i := 0; i < ln; i++ {
		// 此时应该是等于, 因为第一遍走的时候是0
		output[i] = left
		left *= nums[i]
	}
	for i := ln-1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}
	return output
}

func main() {
	
}
