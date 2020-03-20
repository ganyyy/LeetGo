package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	// 构建大顶堆, 维持topK即可
	var makeHeap = func(arr []int) {
		if k <= 0 || len(arr) <= 1 {
			return
		}
		// 从最后一个开始, 看其根节点是否满足小顶堆
		ln := len(arr)
		begin := ln/2 - 1
		for begin >= 0 {
			v := arr[begin]
			lc := begin*2 + 1
			for lc < ln {
				// 取二者最小值
				if lc < ln-1 && arr[lc] < arr[lc+1] {
					lc++
				}
				// 最小的比根节点大, 返回即可
				if arr[lc] <= arr[begin] {
					break
				}
				// 交换根节点
				arr[(lc-1)/2] = arr[lc]
				// 到下一层
				lc = lc*2 + 1
			}
			arr[(lc-1)/2] = v
			begin--
		}
	}
	makeHeap(arr[:k])
	for i := k; i < len(arr); i++ {
		if arr[i] < arr[0] {
			arr[0] = arr[i]
			makeHeap(arr[:k])
		}
	}
	return arr[:k]
}

// 快排切分
func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return arr[:0]
	}
	return quickSearch(arr, 0, len(arr)-1, k-1)
}

// 快速查找
func quickSearch(arr []int, left, right, k int) []int {
	j := partition(arr, left, right)
	if j == k {
		return arr[:k+1]
	}
	if j > k {
		return quickSearch(arr, left, j-1, k)
	} else {
		return quickSearch(arr, j+1, right, k)
	}
}

// 位置交换
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	// 通过双指针处理, 保证i前边都比 pivot小, i >= pivot
	for j := left; j < right; j++ {
		if arr[j] > pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	j := partition(arr, left, right)
	quickSort(arr, left, j-1)
	quickSort(arr, j+1, right)
}

func main() {
	/**
	[0,0,1,2,4,2,2,3,1,4]
	8
	*/
	a := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	//fmt.Println(getLeastNumbers2(a, 8))
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
