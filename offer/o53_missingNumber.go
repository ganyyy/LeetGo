package main

func missingNumber(nums []int) int {
	// 可以用二分

	var left, right = 0, len(nums)

	for left < right {
		var mid = left + (right-left)/2
		if mid == nums[mid] {
			// 在右边
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
