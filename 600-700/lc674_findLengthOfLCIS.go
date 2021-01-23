package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 双指针走起
	var l, r, res = 0, 1, 1
	for r < len(nums) {
		if nums[r] <= nums[r-1] {
			if r-l > res {
				res = r - l
			}
			l = r
		}
		r++
	}

	if r-l > res {
		return r - l
	}
	return res
}
