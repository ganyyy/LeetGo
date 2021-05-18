package main

import (
	"container/heap"
)

type IntArray []int

func (a IntArray) Len() int {
	return len(a)
}

func (a IntArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *IntArray) Push(x interface{}) {
	*a = append(*a, x.(int))
}

func (a *IntArray) Pop() (val interface{}) {
	*a, val = (*a)[:a.Len()-1], (*a)[a.Len()-1]
	return
}

type LessHeap struct {
	IntArray
}

func (h LessHeap) Less(i, j int) bool {
	return h.IntArray[i] < h.IntArray[j]
}

type MoreHeap struct {
	IntArray
}

func (h MoreHeap) Less(i, j int) bool {
	return h.IntArray[i] > h.IntArray[j]
}

func kthLargestValue(matrix [][]int, k int) int {
	// 使用heap & 前缀和
	var row, col = len(matrix), len(matrix[0])
	var push func(v int)
	var get func() int
	if k > row*col {
		k = row*col - k
		var hh = &MoreHeap{IntArray: make([]int, 0, k)}
		push = func(v int) {
			if hh.Len() >= k {
				if hh.IntArray[0] > v {
					heap.Pop(hh)
				} else {
					return
				}
			}
			heap.Push(hh, v)
		}
		get = func() int {
			return hh.IntArray[0]
		}
	} else {
		var hh = &LessHeap{IntArray: make([]int, 0, k)}
		push = func(v int) {
			if hh.Len() >= k {
				if hh.IntArray[0] < v {
					heap.Pop(hh)
				} else {
					return
				}
			}
			heap.Push(hh, v)
		}
		get = func() int {
			return hh.IntArray[0]
		}
	}

	push(matrix[0][0])

	// 第一行, 第一列先全算一遍
	for i := 1; i < col; i++ {
		matrix[0][i] ^= matrix[0][i-1]
		push(matrix[0][i])
	}
	for i := 1; i < row; i++ {
		matrix[i][0] ^= matrix[i-1][0]
		push(matrix[i][0])
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			matrix[i][j] ^= matrix[i-1][j-1] ^ matrix[i][j-1] ^ matrix[i-1][j]
			push(matrix[i][j])
		}
	}
	return get()
}

func main() {
	var arr = []int{8, 10, 5, 8, 5, 7, 6, 0, 1, 4, 10, 6, 4, 3, 6, 8, 7, 9, 4, 2}
	for i := 1; i <= len(arr); i++ {
		var tmp = make([]int, len(arr))
		copy(tmp, arr)
		println(kthLargestValue([][]int{
			tmp,
		}, 2))
	}
}
