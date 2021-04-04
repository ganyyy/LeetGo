package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 从后向前怼
	var l, r = m - 1, n - 1
	for i := m + n - 1; i >= 0; i-- {
		if l >= 0 && r >= 0 {
			if nums1[l] >= nums2[r] {
				nums1[i] = nums1[l]
				l--
			} else {
				nums1[i] = nums2[r]
				r--
			}
		} else if l >= 0 {
			nums1[i] = nums1[l]
			l--
		} else {
			nums1[i] = nums2[r]
			r--
		}
	}
}
