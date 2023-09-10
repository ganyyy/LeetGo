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

func scheduleCourse2(courses [][]int) int {
	// 按照结束时间排序?

	// 首先, 剔除掉所有的 持续时间 > 截止时间的课程. 因为这些课程是不可能完成的
	var validCourses = courses[:0]
	for _, course := range courses {
		if course[0] > course[1] {
			continue
		}
		validCourses = append(validCourses, course)
	}
	courses = validCourses

	if len(courses) < 2 {
		return len(courses)
	}

	// 按照每门课程的截止时间排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	// 大顶堆, 堆顶是耗时最大的课程
	var h Heap
	h.IntSlice = make([]int, 0, len(courses)/2)
	var total int // 总耗时
	for _, course := range courses {
		if duration := course[0]; duration+total <= course[1] {
			// 如果当前已经消耗的时间 + 该门课程的持续时间 <= 该门课程的截至时间, 直接入堆
			total += duration
			// 为啥把时长入堆呢? 这样就可以优先弹出耗时最长的课程, 从而保证堆中的课程的持续时间之和最小
			heap.Push(&h, duration)
		} else if h.Len() > 0 && h.IntSlice[0] > duration {
			// 如果当前课程的持续时间要小于当前堆顶,
			// 首先可以确认的是: 当前的持续时间已经不满足同时学习堆顶课程和当前课程(中间的省略)
			// 那么, 如果堆顶课程的持续时间要 > 当前课程, 那么完全可以优先学习当前课程,
			// 从而降低整体的耗时, 方便未来学习更多的课程(?)
			total += duration - h.IntSlice[0]
			h.IntSlice[0] = duration
			heap.Fix(&h, 0)
		}
	}

	return h.Len()
}
