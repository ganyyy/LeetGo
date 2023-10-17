package main

import "container/heap"

func maxKelements(nums []int, k int) int64 {
	// 用一个堆? 每次都用最大值作为计算的对象?
	var h intHeap = nums
	heap.Init(&h)

	var ret int
	for i := 0; i < k; i++ {
		cur := h[0]
		ret += cur
		h[0] = (cur + 2) / 3
		heap.Fix(&h, 0)
	}
	return int64(ret)
}

// 由int切片组成的大顶堆
type intHeap []int

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 往堆中放入元素
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

// Pop 从堆中取出元素
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
