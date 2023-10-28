package main

import (
	"container/heap"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	var maxInterval = func(end int, nums []int) int {
		sort.Ints(nums)
		var left int
		var ret int

		for _, right := range nums {
			ret = max(ret, right-left)
			left = right
		}
		return max(ret, end-left)
	}
	return (maxInterval(h, horizontalCuts) * maxInterval(w, verticalCuts)) % (10e8 + 7)
}

type BigHeap struct{ sort.IntSlice }

// Less 重写 Less 方法, 使其变成大顶堆
func (h *BigHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }

// Push 往堆中放入元素
func (h *BigHeap) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }

// Pop 从堆中取出元素
func (h *BigHeap) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[:n-1]
	return x
}

func NewBigHeap(nums []int) *BigHeap {
	h := &BigHeap{nums}
	heap.Init(h)
	return h
}

type HeapArray[T any] struct{ elements []T }

// Len returns the length of the array.
func (h *HeapArray[T]) Len() int { return len(h.elements) }

// Swap swaps the elements with indexes i and j.
func (h *HeapArray[T]) Swap(i, j int) { h.elements[i], h.elements[j] = h.elements[j], h.elements[i] }

// Push pushes the element x HeapArray.
func (h *HeapArray[T]) Push(x interface{}) { h.elements = append(h.elements, x.(T)) }

// Pop pops the element from HeapArray.
func (h *HeapArray[T]) Pop() interface{} {
	n := h.Len()
	x := h.elements[n-1]
	h.elements = h.elements[:n-1]
	return x
}

func (h *HeapArray[T]) Peek() T { return h.elements[0] }

// At returns the element at index i.
func (h *HeapArray[T]) At(i int) T { return h.elements[i] }

type BigIntHeap struct{ HeapArray[int] }

// Less 重写 Less 方法, 使其变成大顶堆
func (h *BigIntHeap) Less(i, j int) bool { return h.elements[i] > h.elements[j] }

// NewBigIntHeap 新建一个大顶堆
func NewBigIntHeap(nums []int) *BigIntHeap {
	h := &BigIntHeap{HeapArray[int]{nums}}
	heap.Init(h)
	return h
}
