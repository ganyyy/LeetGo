package main

import "sort"

func arrayPairSum(nums []int) int {
	// sort, 然后一个一个取呗...
	sort.Ints(nums)

	var res int
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}

	return res
}
