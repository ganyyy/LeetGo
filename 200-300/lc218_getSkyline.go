package main

import (
	"container/heap"
	"sort"
)

type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
	// building的左边缘本身就是有序的..!

	n := len(buildings)
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		// 线段的起点和终点. 关键点总是落在左右边缘上.
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	// idx是一个游标. 避免了指针的回退
	idx := 0
	h := hp{}
	// 针对当前边缘
	for _, boundary := range boundaries {
		// 找到所有有交集(建筑的左边缘小于该值)的建筑, 压入右边缘和高度
		// 这里边缘看的是左边界值
		// 这是一个大顶堆, 堆顶元素是高度最高的建筑
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		// 如果堆顶的右边缘(最高的那个建筑)小于当前边缘, 出堆
		// 很抽象..
		// 简单而言, 就是不停的出堆, 直到最高的的建筑的右边界大于当前的边界
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}

		// 记录当前边缘对应的最大高度
		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		// 如果高度和前一个边界值相等, 那么就能合并成一条边界线. 所以这里要特殊检查一下
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}
	return
}
