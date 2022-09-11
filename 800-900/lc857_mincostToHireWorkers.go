//go:build ignore

package main

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality, wage []int, k int) float64 {
	n := len(quality)
	h := make([]int, n)
	for i := range h {
		h[i] = i
	}
	sort.Slice(h, func(i, j int) bool {
		a, b := h[i], h[j]
		// 对比性价比
		// 实际上是 qa/wa > qb/wb
		// qa*wb > qb*wa
		return quality[a]*wage[b] > quality[b]*wage[a]
	})
	totalq := 0
	q := hp{}
	// 按照质量构建大根堆
	// 性价比最高的前k-1个工人对应的工作质量(?)
	for i := 0; i < k-1; i++ {
		totalq += quality[h[i]]
		heap.Push(&q, quality[h[i]])
	}
	ans := 1e9
	for i := k - 1; i < n; i++ {
		idx := h[i]
		totalq += quality[idx]
		heap.Push(&q, quality[idx])
		// 这么理解, 因为本身已经按照性价比进行了倒序排序, 那么此时只需要满足最新加入的员工的基础工资进行处理即可
		// 因为在他之前已选择的员工的性价比肯定都比他高, 那么只要能满足他的工资需求,
		// 其他人的需求也一定可以满足(新来的是下限, 吃白饭不干活的那种)
		// 使用大根堆, 是为了保证每次都可以移除工作质量最高的那一位, 这样可以有效地降低 totalq
		ans = math.Min(ans, float64(wage[idx])/float64(quality[idx])*float64(totalq))
		totalq -= heap.Pop(&q).(int)
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
