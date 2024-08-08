package main

import "slices"

func minimumAddedInteger(nums1, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	// 因为移除两个, 那么前三个中必然保留下来一个!
	// 枚举保留 nums1[2] 或者 nums1[1] 或者 nums1[0]
	// 倒着枚举是因为 nums1[i] 越大答案越小，第一个满足的就是答案
	for i := 2; i > 0; i-- {
		x := nums2[0] - nums1[i]
		// 在 {nums1[i] + x} 中找子序列 nums2
		j := 0
		for _, v := range nums1[i:] {
			if nums2[j] == v+x {
				j++
				// nums2 是 {nums1[i] + x} 的子序列
				if j == len(nums2) {
					return x
				}
			}
		}
	}
	// 题目保证答案一定存在
	return nums2[0] - nums1[0]
}
