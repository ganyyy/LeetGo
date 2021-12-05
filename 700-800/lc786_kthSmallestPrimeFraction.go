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
	// 值二分法, 给定一个实数, 求得所有小于这个实数的分数的个数是否恰好等于k

	// 如果数量大于k, 说明猜测的实数大了, 向左偏移
	// 如果数量小于k, 说明猜测的实数小了, 向右偏移
	n := len(arr)
	for {
		mid := (left + right) / 2
		i, count := -1, 0
		// 记录最大的分数
		x, y := 0, 1

		for j := 1; j < n; j++ {
			// i为啥不需要回退呢?
			// 可以这样理解:
			// i是一定不可能大于j的. 那么相对的, 在小于i的任意值中, arr[n]/arr[j](n <= i) <= arr[i]/arr[j]
			// 随着j的逐渐变大, arr[i]/arr[j]会变小
			for float64(arr[i+1])/float64(arr[j]) < mid {
				i++
				if arr[i]*y > arr[j]*x {
					x, y = arr[i], arr[j]
				}
			}
			// 所以直接加上去就行. 这里也可以理解为是一种特殊的归并
			// 如果 i1/j1 < mid, i2/j1可能大于mid.
			// 如果 i1/j1 < mid, i1/j2也是一定小于mid的
			count += i + 1
		}

		// 统计小于mid的分数的个数, 如果恰好相等, 答案就是这个
		if count == k {
			return []int{x, y}
		}
		if count < k {
			left = mid
		} else {
			right = mid
		}
	}
}

func kthSmallestPrimeFractionHeap(arr []int, k int) []int {
	n := len(arr)
	h := make(fracHeap, n-1)

	// 多路归并: 求多个有序数组的第K大的数
	// 类似的题目:

	// 将整体划分为 n 个子序列(n = len(arr))
	// 每个子序列形如 {arr[0], ..., arr[i]}, i∈[1,n)
	// 对于每个子序列而言, 从头开始依次和末尾元素相除, 整体的结果也是单调递增的
	// 初始情况下先将每个子序列的头部入队

	// 首先入队相对较小的数, 即使用arr[0]为分子, 一路除过来
	for j := 1; j < n; j++ {
		// 很显然, 每个子序列的最小值就是 arr[0]/arr[j]
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
