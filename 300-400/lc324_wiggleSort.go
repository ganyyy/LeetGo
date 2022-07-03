package main

import "sort"

func wiggleSort(nums []int) {
	// 三分快排, 不搞了...
	sort.Ints(nums)
	var n = len(nums)
	// 核心是找到中点
	var l, r = (n - 1) / 2, n - 1
	var ret = make([]int, 0, n)
	for i := 0; i < n; i++ {
		// 然后按照奇偶性, 从尾到头的一次迭代添加
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
