package main

import "sort"

func unequalTriplets(nums []int) (ans int) {
	sort.Ints(nums)
	start, n := 0, len(nums)
	// 多指针, 定位到分界点
	// 1,1,1,2,2,3,3,3
	//	   ⬆   ⬆
	// s     s   s
	// 前边的数字 * 后边的数字 = 组合数
	for i, x := range nums[:n-1] {
		if x != nums[i+1] {
			ans += start * (i - start + 1) * (n - 1 - i)
			start = i + 1
		}
	}
	return
}
