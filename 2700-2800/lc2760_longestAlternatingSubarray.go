package main

func longestAlternatingSubarray(nums []int, threshold int) int {
	// 双指针
	ln := len(nums)
	var start = -1
	var ret int
	for l := 0; l < ln; l++ {
		lv := nums[l]
		if lv > threshold {
			start = -1
			continue
		}
		if start == -1 {
			if lv&1 == 0 {
				start = l
				ret = max(ret, 1)
			}
			continue
		}
		// 能到这里, l一定是 > 0的
		if nums[l-1]&1 == lv&1 {
			if lv&1 == 0 {
				// 偶数,
				start = l
			} else {
				start = -1
			}
			continue
		}
		ret = max(ret, l-start+1)
	}
	return ret
}
