package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type intArr []int

func (ia intArr) Len() int {
	return len(ia)
}

func (ia intArr) Less(i, j int) bool {
	return ia[i] > ia[j]
}

func (ia intArr) Swap(i, j int) {
	ia[i], ia[j] = ia[j], ia[i]
}

func (ia *intArr) Push(x interface{}) {
	*ia = append(*ia, x.(int))
}

func (ia *intArr) Pop() (v interface{}) {
	*ia, v = (*ia)[:ia.Len()-1], (*ia)[ia.Len()-1]
	return
}

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	// 大顶堆? 二分快排?
	if len(arr) < k {
		return arr
	}
	var tmp intArr = arr[:k]
	heap.Init(&tmp)

	for _, v := range arr[k:] {
		if v < tmp[0] {
			heap.Pop(&tmp)
			heap.Push(&tmp, v)
		}
	}
	return tmp
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(smallestK(arr, 5))
}
