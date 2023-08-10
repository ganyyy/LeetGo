//go:build ignore

package main

import (
	"container/heap"
	"sort"
)

func findCrossingTime(n, k int, time [][]int) (cur int) {

	// 整体流程:
	/*
	   1. 从左岸过桥到右岸 [0]
	   2. 在右岸的仓库选择一个箱子 [1]
	   3. 从右岸过桥到左岸 [2]
	   4. 将箱子放到左岸的仓库 [3]
	*/

	// 稳定排序: 按照效率高低进行排序, 当相同的时候, 保持原有的顺序
	sort.SliceStable(time, func(i, j int) bool {
		a, b := time[i], time[j]
		return a[0]+a[2] < b[0]+b[2]
	})

	// waitL: 左边等待过桥的. 所有人的起始状态都是等待从左边过桥, 优先级越高, 下标越小. time已经按照效率排序了, 所以下标越小, 效率越高
	// waitR: 右边等待过桥的
	waitL, waitR := make(hp, k), hp{}
	for i := range waitL {
		waitL[i].i = k - 1 - i // 下标越大效率越低
	}
	// workL: 左边放箱子的
	// workR: 右边放箱子的
	workL, workR := hp2{}, hp2{}

	// cur: 当前时间
	for n > 0 {
		// 1. 如果左边工作的人在当前时间完成了, 就把他们放左边等待过桥的队列中
		for len(workL) > 0 && workL[0].t <= cur {
			heap.Push(&waitL, heap.Pop(&workL)) // 左边完成放箱
		}
		// 2. 如果右边工作的人在当前时间完成了, 就把他们放右边等待过桥的队列中
		for len(workR) > 0 && workR[0].t <= cur {
			heap.Push(&waitR, heap.Pop(&workR)) // 右边完成搬箱
		}
		// 右边过桥的优先级高于左边过桥的
		if len(waitR) > 0 { // 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
			// 3. 如果右边等待过桥的人数大于0, 就把他们过桥, 并且把他们放到左边工作的队列中
			p := heap.Pop(&waitR).(pair)
			cur += time[p.i][2]
			heap.Push(&workL, pair{p.i, cur + time[p.i][3]}) // 放箱，记录完成时间
		} else if len(waitL) > 0 { // 左边过桥
			// 4. 如果左边等待过桥的人数大于0, 就把他们过桥, 并且把他们放到右边工作的队列中
			p := heap.Pop(&waitL).(pair)
			cur += time[p.i][0]
			heap.Push(&workR, pair{p.i, cur + time[p.i][1]}) // 搬箱，记录完成时间
			n--
		} else {
			// 5. 如果左边和右边都没有人过桥, 就把当前时间设置为左边和右边工作的人中最早完成的时间
			if len(workL) == 0 { // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
				cur = workR[0].t
			} else if len(workR) == 0 {
				cur = workL[0].t
			} else {
				cur = min(workL[0].t, workR[0].t)
			}
		}
	}
	for len(workR) > 0 {
		p := heap.Pop(&workR).(pair) // 右边完成搬箱
		// 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
		cur = max(p.t, cur) + time[p.i][2]
	}
	return cur // 最后一个过桥的时间
}

// pair 是个二元组，i 表示下标，t 表示完成时间
type pair struct{ i, t int }

// hp 是个大根堆，存放等待过桥的人，按照效率从高到低排序. 等待过桥的人只看效率, 不看完成时间
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// hp2 是个小根堆，存放放/搬箱的人，按照完成时间从小到大排序. 放/搬箱的人只看完成时间, 不看效率
type hp2 []pair

func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
