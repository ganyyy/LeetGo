package main

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	if len(nums1) != len(nums2) {
		return nums1
	}
	ln := len(nums1)
	sort.Ints(nums1)
	idx := make([]int, ln)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] < nums2[idx[j]]
	})
	// fmt.Println(idx)
	// for _, id := range idx {
	//     fmt.Println(nums2[id])
	// }
	// fmt.Println(nums1)

	l, r := 0, ln-1
	ret := make([]int, ln)
	for _, v := range nums1 {
		// fmt.Println(v, idx[l], nums2[idx[l]])
		if v > nums2[idx[l]] {
			ret[idx[l]] = v
			l++
		} else {
			ret[idx[r]] = v
			r--
		}
		// fmt.Println(l, r)
	}
	return ret
}
