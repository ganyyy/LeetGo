package main

func maxSum(nums1 []int, nums2 []int) int {
	// 序列DP
	// F[i] = 到达nums1[i]时的最大值
	// G[j] = 到达nums2[j]时的最大值
	var l1, l2 = len(nums1), len(nums2)
	var F, G = make([]int, l1+1), make([]int, l2+1)
	var i, j = 1, 1

	for i <= l1 || j <= l2 {
		if i <= l1 && j <= l2 {
			if nums1[i-1] < nums2[j-1] {
				F[i] = F[i-1] + nums1[i-1]
				i++
			} else if nums1[i-1] > nums2[j-1] {
				G[j] = G[j-1] + nums2[j-1]
				j++
			} else {
				// nums1[i-1] == nums2[j-1]
				F[i] = max(F[i-1], G[j-1]) + nums1[i-1]
				G[j] = F[i]
				i++
				j++
			}
		} else if i <= l1 {
			F[i] = F[i-1] + nums1[i-1]
			i++
		} else {
			// j < l2
			G[j] = G[j-1] + nums2[j-1]
			j++
		}
	}

	return max(F[l1], G[l2]) % (1e9 + 7)
}
