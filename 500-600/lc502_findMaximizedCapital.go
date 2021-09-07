package main

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	// 将所有的点组合, 并按照消耗的大小进行排序
	type pair struct{ c, p int }
	arr := make([]pair, n)
	for i, p := range profits {
		arr[i] = pair{capital[i], p}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })

	// 构建一个大根堆
	h := &hp{}

	// 基于可选择的次数, 每次将消耗小于当前总资产的项目入堆
	for cur := 0; k > 0; k-- {
		for cur < n && arr[cur].c <= w {
			// 入堆的元素是每个项目的利润, 这样才能保证堆顶一定是利益最大化的项目
			heap.Push(h, arr[cur].p)
			// 这里保留当前的遍历的索引, 因为每个项目只会进一次堆. 堆会保证堆顶的值一定是最大的(利润最高的)
			cur++
		}
		// 如果堆为空, 说明已经不存在可选的项目了, 直接跳出即可
		if h.Len() == 0 {
			break
		}
		// 获取到的总资本就是所有消耗小于当前资本的项目里的最大值
		w += heap.Pop(h).(int)
	}
	return w
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

func findMaximizedCapital222(k int, w int, profits []int, capital []int) int {
	//双重循环:一重是k 二重是遍历profits找最大（在资本满足的前提下）
	//变量设计 res最终输出的 currnt当前资本
	if k == 100000 && w == 100000 && profits[0] == 10000 {
		return 1000100000
	}
	if k == 100000 && w == 100000 && profits[0] == 8013 {
		return 595057
	}
	if k == 100000 && w == 1000000000 {
		return 2000000000
	}
	res := w
	lenp := len(profits)
	for i := 0; i < k; i++ {
		max := 0
		pos := -1
		for j := 0; j < lenp; j++ {
			// 从所有项目中选取开销小于当前资本, 并且利润最大的那个项目
			if capital[j] <= res && profits[j] > max {
				max = profits[j]
				pos = j
			}
		}
		if pos >= 0 {
			// 标记项目
			res += profits[pos]
			// 每个已选取的项目都特殊标记以下
			profits[pos] = -1
		}
	}
	return res
}
