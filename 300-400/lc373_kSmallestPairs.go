package main

import "container/heap"

func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		// 先将 nums1[0], nums2[0...n]入堆
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			// 此时不断地将 nums2[j+1]入堆
			// 因为最小的值一定在 nums1[i],nums2[j+1] / nums1[i+1],nums2[j] 中
			// nums1[i+1],nums2[j] 已在队列中, 所以此时只需要将 nums1[i],nums2[j+1] 入堆即可
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }
