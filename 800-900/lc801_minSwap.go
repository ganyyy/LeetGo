//go:build ignore

package main

func minSwap(nums1, nums2 []int) int {
	n := len(nums1)
	// a: 不交换
	// b: 交换
	a, b := 0, 1
	for i := 1; i < n; i++ {
		at, bt := a, b
		a, b = n, n
		// [1, 2] [3, 4]
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			// 不交换的话, a不变(at)
			a = min(a, at)
			// 交换的话, b+1(bt+1)
			b = min(b, bt+1)
		}
		// [3, 5] [4, 4]
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			// i不交换, 那么就需要i-1交换(bt)
			a = min(a, bt)
			// i交换, i-1不交换+1(at+1)
			b = min(b, at+1)
		}
	}
	return min(a, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
