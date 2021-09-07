package main

func search(nums []int, target int) int {
	var l, r = 0, len(nums)
	for l < r {
		var mid = l + (r-l)/2
		if v := nums[mid]; v == target {
			return mid
		} else if v > target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return -1
}
