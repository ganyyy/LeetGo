package main

import (
	"container/heap"
	"sort"
)

type MinIntHeap struct{ sort.IntSlice }

func (m *MinIntHeap) Push(x any) {
	v, _ := x.(int)
	m.IntSlice = append(m.IntSlice, v)
}

func (m *MinIntHeap) Pop() any {
	a := m.IntSlice
	v := a[len(a)-1]
	m.IntSlice = a[:len(a)-1]
	return v
}

type MaxIntHeap struct{ MinIntHeap }

func (m *MaxIntHeap) Less(i, j int) bool { return m.IntSlice[i] > m.IntSlice[j] }

func minStoneSum(piles []int, k int) int {
	var h MaxIntHeap
	h.IntSlice = piles
	heap.Init(&h)
	var ret int
	for k > 0 {
		k--
		var half = h.IntSlice[0] / 2
		h.IntSlice[0] -= half
		heap.Fix(&h, 0)
	}
	for _, v := range h.IntSlice {
		ret += v
	}
	return ret
}
