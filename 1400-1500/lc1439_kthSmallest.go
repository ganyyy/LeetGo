package main

import "container/heap"

func kthSmallest(mat [][]int, k int) int {
	// 逐级合并
	// 两两合并到最后, 一定可以获取全局的结果
	var buf = make([]pair, k)
	var cur = mat[0]
	for _, nxt := range mat[1:] {
		cur = kSmallestPairs(nxt, cur, k, buf)
	}
	return cur[k-1]
}

func kSmallestPairs(nums1, nums2 []int, k int, buf []pair) (ans []int) {
	m, n := len(nums1), len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	buf = buf[:0]
	h := hp{data: buf}
	for i := 0; i < m; i++ {
		h.data = append(h.data, pair{i, 0, nums1[i] + nums2[0]})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, p.sum)
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1, nums1[i] + nums2[j+1]})
		}
	}
	return
}

type pair struct{ i, j, sum int }
type hp struct {
	data []pair
}

func (h hp) Len() int            { return len(h.data) }
func (h hp) Less(i, j int) bool  { a, b := h.data[i], h.data[j]; return a.sum < b.sum }
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }
