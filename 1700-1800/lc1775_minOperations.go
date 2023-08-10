//go:build ignore

package main

func help(h1 [7]int, h2 [7]int, diff int) (res int) {
	h := [7]int{}
	// h1 < h2
	for i := 1; i < 7; i++ {
		h[6-i] += h1[i] // h1[1] -> 5
		h[i-1] += h2[i] // h2[6] -> 5
	}

	for i := 5; i > 0 && diff > 0; i-- {
		// 向上取整(多算了也没事, 可以匀出去)
		t := min((diff+i-1)/i, h[i])
		// 减少了多少次
		res += t
		// diff的剩余量
		diff -= t * i
	}
	return res
}

func minOperations(nums1 []int, nums2 []int) (ans int) {
	n, m := len(nums1), len(nums2)
	if 6*n < m || 6*m < n {
		return -1
	}
	var cnt1, cnt2 [7]int
	// diff = sum(nums1)-sum(nums2)
	diff := 0
	for _, i := range nums1 {
		cnt1[i]++
		diff += i
	}
	for _, i := range nums2 {
		cnt2[i]++
		diff -= i
	}
	if diff == 0 {
		return 0
	}
	if diff > 0 {
		return help(cnt2, cnt1, diff)
	}
	return help(cnt1, cnt2, -diff)
}
