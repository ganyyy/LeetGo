package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	// 递归
	// 如何判断是不是已经存在了呢?

	// 可以排序?
	sort.Ints(nums)

	var res [][]int
	var tmp = make([]int, 0, len(nums))

	var dfs func(i int)
	dfs = func(i int) {
		var t = make([]int, len(tmp))
		copy(t, tmp)
		res = append(res, t)
		for j := i; j < len(nums); j++ {
			if j > i && nums[j-1] == nums[j] {
				continue
			}
			tmp = append(tmp, nums[j])
			dfs(j + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0)

	return res
}
