package main

import "sort"

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return nil
	}

	var nums = make(map[int]int)

	for _, change := range changed {
		nums[change]++
	}
	sort.Ints(changed)
	var valid = changed[:0]
	for _, change := range changed {
		if nums[change] <= 0 {
			// 这个数字已经被消耗了
			continue
		}
		// 从小到大的
		nums[change]--
		double := change * 2
		if nums[double] <= 0 {
			// 不够了
			return nil
		}
		nums[double]--
		valid = append(valid, change)
	}
	return valid
}
