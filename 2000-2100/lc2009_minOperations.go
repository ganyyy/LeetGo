package main

import "slices"

func minOperations(nums []int) int {
	n := len(nums)
	slices.Sort(nums)
	a := slices.Compact(nums) // 原地去重
	// ans: 大小为n的窗口中, 包含多少个元素不需要替换的元素
	ans, left := 0, 0
	for i, x := range a {
		// [1,2,3,4,6]
		// (i, x)  a[l]  x-n+1  ans
		// (0, 1)   1   > -3     1
		// (1, 2)   1   > -2     2
		// (2, 3)   1   > -1     3
		// (3, 4)   1   >  0     4
		// (4, 6)   1   <  2     4
		for a[left] < x-n+1 { // a[left] 不在窗口内
			left++
		}
		ans = max(ans, i-left+1)
	}

	return n - ans
}
