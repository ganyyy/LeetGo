package main

func maximumDifference(nums []int) int {
	var ret = -1
	var mi = nums[0]

	for _, v := range nums[1:] {
		if v > mi {
			ret = max(v-mi, ret)
		}
		mi = min(mi, v)
	}
	return ret
}
