package main

import "sort"

func minimizeArrayValue(nums []int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	return sort.Search(mx, func(limit int) bool {
		// 关键是怎么把这个限制给应用到数组上的
		// 这个limit相当于是设定了数组中的最大值
		// 比这个大的要减(nums[i]-limit)次并加到nums[i-1]上
		// 同理, 这个拥有传递性. 中断的前提是 在某个位置上, nums[i]+extra <= limit
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		// 最后所有的累加值都到了nums[0]上
		return nums[0]+extra <= limit
	})
}
