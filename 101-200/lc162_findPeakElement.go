package main

func findPeakElement(nums []int) int {
	// 二分吧..

	// 核心是找到一个拐点

	/*
	   这样理解:
	       如果nums[mid] < nums[mid+1], 那么nums[mid+2]存在两种可能:
	           1. nums[mid+2]>nums[mid+1], 那么就继续向右查询, 直到最右边,
	              因为nums[len(nums)-1] > -∞ 且nums[len(nums)-1] >nums[len(nums)-2]
	           2. nums[mid+2]<nums[mid+1], 那么mid+1就是一个峰值
	       同理, 如果nums[mid]>nums[mid+1], 就向左边查询
	*/

	// 注意: 这里右边界使用len(nums)-1 是为了防止nums[mid+1]越界
	var l, r = 0, len(nums) - 1
	for l < r {
		var mid = l + (r-l)/2

		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
