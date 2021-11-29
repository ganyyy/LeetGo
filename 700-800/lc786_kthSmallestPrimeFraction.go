package main

import (
	"container/heap"
	"sort"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	// 暴力解法, C{2, len(arr)}
	var tmp = make([]int, 0, len(arr)*(len(arr)-1)/2)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			tmp = append(tmp, (arr[i]<<32)|arr[j])
		}
	}

	const Mask = (1 << 32) - 1
	var getVal = func(i int) (a, b int) {
		return i >> 32, i & Mask
	}

	sort.Slice(tmp, func(i, j int) bool {
		// 整数运算的性能相对于浮点运算能提升16%
		// 704 -> 592
		var i1, i2 = getVal(tmp[i])
		var i3, i4 = getVal(tmp[j])
		return i1*i4 < i2*i3
	})

	var ret = tmp[k-1]

	return []int{ret >> 32, ret & Mask}
}

func kthSmallestPrimeFractionBinarySearch(arr []int, k int) []int {
	// 分数的上下界就是0, 1
	left, right := 0.0, 1.0
	p, q := 0, 1
	// 值二分法, 给定一个实数, 求得所有小于这个实数的分数
	n := len(arr)
	for {
		p = 0
		count := 0

		// 取一个实数
		mid := (left + right) / 2
		// 双指针获取小于该实数的分数的个数

		// i为分子, 向右移动
		// j为分母, 向左移动
		// TODO 周末好好看看
		for i, j := 0, n-1; i < n; i++ {
			// 获取最后一个满足 arr[i]/arr[j] < mid 的j
			for j >= 0 && float64(arr[i]) > mid*float64(arr[n-1-j]) {
				j--
			}
			// count表示满足 小于 mid 的数字的个数(?)
			count += j + 1
			if j >= 0 && p*arr[n-1-j] < q*arr[i] {
				p, q = arr[i], arr[n-1-j]
			}
		}

		if count < k {
			left = mid
		} else if count > k {
			right = mid
		} else {
			return []int{p, q}
		}
	}

}

func kthSmallestPrimeFractionHeap(arr []int, k int) []int {
	n := len(arr)
	h := make(fracHeap, n-1)
	// 首先入队相对较小的数, 即使用arr[0]为分子, 一路除过来
	for j := 1; j < n; j++ {
		h[j-1] = frac{arr[0], arr[j], 0, j}
	}
	heap.Init(&h)

	// 循环出队k次, 剩下的队首就是第k小的数
	for loop := k - 1; loop > 0; loop-- {
		f := heap.Pop(&h).(frac)
		// 如果分子的索引+1小于分母, 那么也可以入队
		if f.i+1 < f.j {
			heap.Push(&h, frac{arr[f.i+1], f.y, f.i + 1, f.j})
		}
	}
	return []int{h[0].x, h[0].y}
}

type frac struct{ x, y, i, j int }
type fracHeap []frac

func (h fracHeap) Len() int            { return len(h) }
func (h fracHeap) Less(i, j int) bool  { return h[i].x*h[j].y < h[i].y*h[j].x }
func (h fracHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *fracHeap) Push(v interface{}) { *h = append(*h, v.(frac)) }
func (h *fracHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
