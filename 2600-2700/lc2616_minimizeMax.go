package main

import "sort"

func minimizeMax(nums []int, p int) int {
	if len(nums) < 2 || p == 0 {
		return 0
	}
	sort.Ints(nums)

	return sort.Search(nums[len(nums)-1]-nums[0], func(diff int) bool {
		// 怎么找出差值<=m的数对呢?
		// diff就是限制相邻的两个数的差值的最大值
		// 在整体排序之后, 两两相邻的两个数之间的差值是最小的
		// 如果符合条件的数对的数量 >= p, 那么就说明diff还有下降空间

		var cnt int
		for i := 1; i < len(nums); i++ {
			if nums[i]-nums[i-1] <= diff {
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}
