package main

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	// 按照课程截止时间进行排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	// h是一个大顶堆, 表示的是当前课程列表中, 耗时最长的课程
	// 同时, h的长度也是当前期限内, 能学习的最多的课程的最优解
	h := &Heap{}
	total := 0 // 优先队列中所有课程的总时间
	for _, course := range courses {
		if t := course[0]; total+t <= course[1] {
			// 如果允许学习当前课程, 就将当前课程入队
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			// 如果当前剩余时间不足以学习这门课程, 并且该门课程的耗时小于耗时最大的课程
			// 那么就可以进行替换, 这样可以保证有更多充裕的时间来学习
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}
