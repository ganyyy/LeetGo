package main

import "sort"

func minimizeMax(nums []int, p int) int {
	if len(nums) < 2 || p == 0 {
		return 0
	}
	sort.Ints(nums)

	return sort.Search(nums[len(nums)-1]-nums[0], func(m int) bool {
		// 怎么找出差值<=m的数对呢?

		var cnt int
		for i := 1; i < len(nums); i++ {
			if nums[i]-nums[i-1] <= m {
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}
