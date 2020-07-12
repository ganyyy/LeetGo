package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	// 让1是较短的的那一个
	if len(nums1) > len(nums2) {
		nums2, nums1 = nums1, nums2
	}
	l1, l2 := len(nums1), len(nums2)
	var k int
	for i, j := 0, 0; i < l1 && j < l2; {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
