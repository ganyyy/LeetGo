package main

import "math"

func minSubArrayLen(s int, nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}
	if ln == 1 {
		if nums[0] == s {
			return 1
		}
		return 0
	}
	var left, right int
	cur := nums[left]
	var res = math.MaxInt32
	for right < ln {
		if cur < s {
			right++
			if right < ln {
				cur += nums[right]
			}
		} else {
			if res > right-left+1 {
				res = right - left + 1
			}
			if right > left {
				cur -= nums[left]
				left++
			} else {
				right++
				if right < ln {
					cur += nums[right]
				}
			}

		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func minSubArrayLen2(target int, nums []int) int {
	var total int
	var minLength = math.MaxInt32
	for left, right := 0, 0; right < len(nums); right++ {
		total += nums[right]

		for total >= target && left <= right {
			minLength = min(minLength, right-left+1)
			total -= nums[left]
			left++
		}
	}
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}
