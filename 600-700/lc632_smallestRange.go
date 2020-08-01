package main

import (
	"container/heap"
	"fmt"
	"math"
)

func smallestRange(nums [][]int) []int {
	// 区间头, 区间尾y, 区间中的最大值
	minx, miny, maxInH := 0, math.MaxInt64, math.MinInt64
	// 构建一个小顶堆
	h := new(ItemHeap)
	totalNum := 0
	// 初始化, 把每一个数组中的 最小值 放入到堆里
	for i, num := range nums {
		heap.Push(h, &Item{
			value:   num[0],
			numsIdx: i,
			idx:     0,
		})
		// 更新最大值
		maxInH = max(maxInH, num[0])
		// 总计数
		totalNum += len(num)
	}
	// 遍历所有的值
	for i := 0; i < totalNum; i++ {
		// 取出最小区间的最小值
		minItem := heap.Pop(h).(*Item)
		// 如果区间差 大于 区间最大值和堆中最小值的差, 就更新一下区间
		if miny-minx > maxInH-minItem.value {
			minx = minItem.value
			miny = maxInH
		}
		// 如果某个数组走到头了, 就退出
		if minItem.idx+1 == len(nums[minItem.numsIdx]) {
			break
		}
		// 把取出的最小值的数组的下一个值放入到最小区间中
		heap.Push(h, &Item{
			value:   nums[minItem.numsIdx][minItem.idx+1],
			numsIdx: minItem.numsIdx,
			idx:     minItem.idx + 1,
		})
		// 更新区间最大值
		maxInH = max(maxInH, nums[minItem.numsIdx][minItem.idx+1])
	}

	return []int{minx, miny}
}

type Item struct {
	value   int // 对应的值
	numsIdx int // 对应第几个数组
	idx     int // 对应数组中的第几个
}

type ItemHeap []*Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ItemHeap) Less(i, j int) bool { return h[i].value < h[j].value }

func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}
func (h *ItemHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
	t := [][]int{
		{4, 10, 15, 24, 26},
		{0, 9, 12, 20},
		{5, 18, 22, 30},
	}
	fmt.Println(smallestRange(t))
}
