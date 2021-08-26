package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MedianFinderErr struct {
	nums []int
}

//ConstructorErr 超时太严重了...
func ConstructorErr() MedianFinderErr {
	return MedianFinderErr{}
}

func (f *MedianFinderErr) AddNum(num int) {
	var idx = sort.Search(len(f.nums), func(i int) bool {
		return f.nums[i] > num
	})
	f.nums = append(f.nums, num)
	if idx < len(f.nums) {
		copy(f.nums[idx+1:], f.nums[idx:])
		f.nums[idx] = num
	}
}

func (f *MedianFinderErr) FindMedian() float64 {
	var ln = len(f.nums)
	if ln&1 == 0 {
		// 偶数
		return float64(f.nums[ln>>1]+f.nums[(ln>>1)-1]) / 2
	} else {
		return float64(f.nums[ln>>1])
	}
}

// 使用堆搞一个

type IntNums []int

func (in IntNums) Len() int {
	return len(in)
}

func (in IntNums) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in *IntNums) Push(x interface{}) {
	*in = append(*in, x.(int))
}

func (in *IntNums) Pop() (ret interface{}) {
	ret, *in = (*in)[in.Len()-1], (*in)[:in.Len()-1]
	return
}

type MinHeap struct {
	IntNums
}

func (m MinHeap) Less(i, j int) bool {
	return m.IntNums[i] < m.IntNums[j]
}

type MaxHeap struct {
	IntNums
}

func (m MaxHeap) Less(i, j int) bool {
	return m.IntNums[i] > m.IntNums[j]
}

type MedianFinder struct {
	min *MinHeap // 存储后半部分
	max *MaxHeap // 存储前半部分
}

func Constructor() MedianFinder {
	return MedianFinder{
		min: &MinHeap{},
		max: &MaxHeap{},
	}
}

func (f *MedianFinder) AddNum(num int) {
	// 首先扔到小顶堆中
	if f.min.Len() == 0 {
		heap.Push(f.min, num)
		return
	}

	if f.min.Len() < f.max.Len() {
		// 如果后半部分的数量小于前半部分, 为了保证两者的平衡, 需要加入到后半部分中

		if num < f.max.IntNums[0] {
			// 如果num小于前半部分的最大值, 那么就将前半部分的最大值放入到后半部分中
			heap.Push(f.min, heap.Pop(f.max))
			heap.Push(f.max, num)
		} else {
			// 否则, 直接放入到后半部分中
			heap.Push(f.min, num)
		}
	} else {
		if num > f.min.IntNums[0] {
			// 如果num大于后半部分的最小值, 那么就将后半部分的最小值放入到前半部分中
			heap.Push(f.max, heap.Pop(f.min))
			heap.Push(f.min, num)
		} else {
			// 否则, 直接放入到前半部分中
			heap.Push(f.max, num)
		}
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.min.Len() == 0 && f.max.Len() == 0 {
		return -1
	}
	if f.max.Len() == f.min.Len() {
		// 偶数
		return float64(f.max.IntNums[0]+f.min.IntNums[0]) / 2
	} else if f.min.Len() < f.max.Len() { // 奇数的情况, 查找最长的那个
		return float64(f.max.IntNums[0])
	} else {
		return float64(f.min.IntNums[0])
	}
}

func main() {
	var minHeap = &MinHeap{}
	var maxHeap = &MaxHeap{}

	heap.Init(minHeap)
	heap.Init(maxHeap)
	for i := 0; i < 10; i++ {
		heap.Push(minHeap, i)
		heap.Push(maxHeap, i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(heap.Pop(minHeap), heap.Pop(maxHeap))
	}

	fmt.Println(minHeap, maxHeap)
}
