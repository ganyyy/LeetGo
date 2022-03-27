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

func findKthLargest2(nums []int, k int) int {
	var arr = make([]int, k)
	var heapArr = func() {
		// 这是一个小顶堆, 那么堆顶一定是最小值
		for i := k/2 - 1; i >= 0; i-- {
			var root = i
			for {
				var child = root*2 + 1
				if child >= k {
					break
				}
				if child+1 < k && arr[child+1] < arr[child] {
					child++
				}
				if arr[child] >= arr[root] {
					break
				}
				// 这一步是关键欸
				arr[root], arr[child] = arr[child], arr[root]
				root = child
			}
		}
	}

	copy(arr, nums[:k])
	heapArr()

	for _, v := range nums[k:] {
		if v <= arr[0] {
			continue
		}
		arr[0], arr[k-1] = arr[k-1], v
		heapArr()
	}

	return arr[0]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
