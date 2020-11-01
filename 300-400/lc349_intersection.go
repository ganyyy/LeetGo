package main

var empty = struct{}{}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	var m = make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		m[v] = empty
	}
	var res = make([]int, 0, len(nums1))

	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			delete(m, v)
			res = append(res, v)
		}
	}
	return res

}
