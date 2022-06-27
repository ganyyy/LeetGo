package main

import "sort"

func wiggleSort(nums []int) {
	sort.Ints(nums)
	var n = len(nums)
	var l, r = (n - 1) / 2, n - 1
	var ret = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			ret = append(ret, nums[l])
			l--
		} else {
			ret = append(ret, nums[r])
			r--
		}
	}
	copy(nums, ret)
}
