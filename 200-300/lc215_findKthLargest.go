package main

import (
	"container/heap"
	"fmt"
)

type array []int

func (a array) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a array) Len() int {
	return len(a)
}

func (a array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *array) Push(x interface{}) {
	(*a) = append((*a), x.(int))
}

func (a *array) Pop() (v interface{}) {
	*a, v = (*a)[:a.Len()-1], (*a)[len(*a)-1]
	return v
}
func findKthLargest(nums []int, k int) int {
	// 用堆
	if len(nums) < k {
		return 0
	}
	var t array = nums[:k]
	heap.Init(&t)
	for _, v := range nums[k:] {
		if v > t[0] {
			heap.Pop(&t)
			heap.Push(&t, v)
		}
	}
	return t[0]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
