package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	var ret int
	var left, right int
	var cur = 1
	for ; right < len(nums); right++ {
		cur *= nums[right]
		for left <= right && cur >= k {
			cur /= nums[left]
			left++
		}
		ret += right - left + 1
	}

	return ret
}
